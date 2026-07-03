package alamat_kurir_handle

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/cache"
	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	se_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/search_engine/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	notification_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/models"
	notification_request "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/request"
	notification_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/seeders"
)

func CreateMasukanAlamatKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	var Objek sot_models.AlamatKurir
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.AlamatKurir = cass_models.AlamatKurir{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodeNegara:      Objek.KodeNegara,
		KodePos:         Objek.KodePos,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
	}

	parsedData := ObjekCass.ParseToCUDType()

	// Gunakan cass_sot_replica untuk tabel replica (Bukan cass_historcal)
	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return err
	}

	pencatatan := historical_format.Sekarang()
	parsedData["tahun_update"] = pencatatan.TahunUpdate
	parsedData["bulan_update"] = pencatatan.BulanUpdate
	parsedData["event_time"] = pencatatan.EventTime

	// Perbaikan Bug: parsedData WAJIB dipassing ke parameter ke-4 agar data historical tersimpan
	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return err
	}

	var ObjekSearchEngine se_models.AlamatKurir = se_models.AlamatKurir{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodeNegara:      Objek.KodeNegara,
		KodePos:         Objek.KodePos,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
	}

	// Perbaikan Bug: `strconv.FormatInt` butuh basis angka 10 (desimal), bukan 0 (nol akan bikin panic/error)
	// Gunakan nil pada parameter kedua AddDocuments jika meilisearch-go Anda versi baru
	if task_info, err := se_index.AlamatKurir.AddDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("task dengan id:" + task_info.IndexUID + " diproses")
	}

	var pesanNotifikasi string
	panggilanAlamat := strings.ToLower(Objek.PanggilanAlamat)

	switch {
	case strings.Contains(panggilanAlamat, "hub") || strings.Contains(panggilanAlamat, "pool"):
		pesanNotifikasi = fmt.Sprintf("ðŸ¢ Hub/Pool baru ('%s') berhasil didaftarkan. Koordinat operasional siap digunakan untuk pickup barang!", Objek.NamaAlamat)
	case strings.Contains(panggilanAlamat, "rumah") || strings.Contains(panggilanAlamat, "basecamp"):
		pesanNotifikasi = fmt.Sprintf("ðŸ  Basecamp kurir lu ('%s') udah disimpan. Sekarang titik istirahat atau standby jadi lebih presisi!", Objek.NamaAlamat)
	default:
		pesanNotifikasi = fmt.Sprintf("ðŸ“ Alamat operasional baru ('%s' - %s) sukses ditambahkan ke profil kurir lu.", Objek.PanggilanAlamat, Objek.NamaAlamat)
	}

	var Notification notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   Objek.IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸš€ Alamat Operasional Kurir Disimpan!",
		Pesan:     pesanNotifikasi,
		Pop:       5.0, // Muncul sebagai pop-up selama 5 detik di aplikasi kurir
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 1, 0).Format(time.RFC3339), // Expired 1 bulan
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"alamat_kurir_id": Objek.ID,
				"action_type":     "create_alamat_kurir",
				"platform":        "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"latitude":     Objek.Latitude,
				"longitude":    Objek.Longitude,
				"click_action": "OPEN_MAPS_ROUTE",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, Notification, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk); err != nil {
		return fmt.Errorf("gagal mengirim notifikasi create kurir: %w", err)
	}
	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatedEditAlamatKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdateEditAlamatKurir"
	var Objek sot_models.AlamatKurir
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.AlamatKurir = cass_models.AlamatKurir{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodeNegara:      Objek.KodeNegara,
		KodePos:         Objek.KodePos,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
	}

	parsedData := ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return err
	}

	pencatatan := historical_format.Sekarang()
	parsedData["tahun_update"] = pencatatan.TahunUpdate
	parsedData["bulan_update"] = pencatatan.BulanUpdate
	parsedData["event_time"] = pencatatan.EventTime

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return err
	}

	var ObjekSearchEngine se_models.AlamatKurir = se_models.AlamatKurir{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodeNegara:      Objek.KodeNegara,
		KodePos:         Objek.KodePos,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
	}

	if task_info, err := se_index.AlamatKurir.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("task dengan id:" + task_info.IndexUID + " diproses")
	}

	var KataKataNotif string = ""
	var JudulNotif string = strings.Trim(handle_services, "Update")

	if strings.Contains(Objek.NamaAlamat, "rumah") || strings.Contains(Objek.NamaAlamat, "Rumah") {
		KataKataNotif = fmt.Sprintf("Kamu Baru aja Ngeupdate alamat rumah mu ya")
	} else if strings.Contains(Objek.NamaAlamat, "basecamp") || strings.Contains(Objek.NamaAlamat, "Basecamp") {
		KataKataNotif = fmt.Sprintf("Kamu Baru aja Ngeupdate alamat basecamp mu ya")
	}

	var Notification notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   Objek.IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     JudulNotif,
		Pesan:     KataKataNotif,
		Pop:       0.8,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 2).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"alamat_kurir_id": Objek.ID,
				"action_type":     "update_alamat_kurir",
				"platform":        "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"latitude":     Objek.Latitude,
				"longitude":    Objek.Longitude,
				"click_action": "OPEN_MAPS_ROUTE",
			},
		},
	}

	if err := notification_request.PostToNotification[notification_models.NotificationKurir](ctx, Notification, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		return err
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusAlamatKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	var Objek sot_models.AlamatKurir
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.AlamatKurir = cass_models.AlamatKurir{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodeNegara:      Objek.KodeNegara,
		KodePos:         Objek.KodePos,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
	}

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return err
	}

	parsedData := ObjekCass.ParseToCUDType()
	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return err
	}

	idStr := strconv.FormatInt(ObjekCass.ID, 10)

	if task_info, err := se_index.AlamatKurir.DeleteDocumentWithContext(ctx, idStr, nil); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("task dengan id:" + task_info.IndexUID + " diproses")
	}

	// ðŸ”” LOGIK NOTIFIKASI DELETE
	pesanDelete := fmt.Sprintf("ðŸ—‘ï¸ Alamat '%s' (%s) dicabut dari rute operasional lu. Inbox ini otomatis hilang dalam 3 hari.", Objek.PanggilanAlamat, Objek.NamaAlamat)

	var NotificationDelete notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:  Objek.IdKurir,
		Pengirim: notification_seeders.Sistem,
		Judul:    "ðŸ—‘ï¸ Alamat Kurir Dihapus",
		Pesan:    pesanDelete,
		// Pop sengaja dikosongin (default 0) -> Artinya gak bakal muncul pop-up/toast mengganggu di layar kurir, tapi diem-diem langsung masuk inbox tab.
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 3).Format(time.RFC3339), // Cukup 3 hari aja buat history delete kurir
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"alamat_kurir_id": Objek.ID,
				"action_type":     "delete_alamat_kurir",
				"platform":        "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REDIRECT_TO_DASHBOARD",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotificationDelete, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk); err != nil {
		return fmt.Errorf("gagal mengirim notifikasi delete kurir: %w", err)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
