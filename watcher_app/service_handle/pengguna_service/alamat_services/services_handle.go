package alamat_pengguna_handle

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/meilisearch/meilisearch-go"

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

// data body yang diinput merupakan model relasi alamat pengguna
func CreateAlamatPub(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateAlamatPub"
	var Objek sot_models.AlamatPengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.AlamatPengguna = cass_models.AlamatPengguna{
		ID:              Objek.ID,
		IDPengguna:      Objek.IDPengguna,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodePos,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica sync %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historica db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.AlamatPengguna = se_models.AlamatPengguna{
		ID:              Objek.ID,
		IDPengguna:      Objek.IDPengguna,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodePos,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.AlamatPenggunaIndex.AddDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("Berhasil memasukan data ke dalam search engine dengan antrean UID %s", task_info.IndexUID)
	}
	fmt.Println("Berhasil mendapatkan data", Objek.ID)

	// Skenario penulisan pesan biar lebih engaging dan dinamis
	var pesanNotifikasi string
	panggilanAlamat := strings.ToLower(Objek.PanggilanAlamat) // Contoh: rumah, kantor, kosan

	switch {
	case strings.Contains(panggilanAlamat, "rumah"):
		pesanNotifikasi = fmt.Sprintf("ðŸ  Hore! Alamat rumah baru lu ('%s') udah berhasil disimpan. Paket belanjaan siap meluncur langsung ke depan pintu!", Objek.NamaAlamat)
	case strings.Contains(panggilanAlamat, "kantor") || strings.Contains(panggilanAlamat, "kerja"):
		pesanNotifikasi = fmt.Sprintf("ðŸ’¼ Alamat kantor baru ('%s') aman terkendali. Jangan lupa set jadi alamat utama pas jam kerja ya!", Objek.NamaAlamat)
	default:
		pesanNotifikasi = fmt.Sprintf("ðŸ“ Alamat baru ('%s' - %s) sukses ditambahkan! Sekarang kirim-kirim barang ke sini jadi makin gampang.", Objek.PanggilanAlamat, Objek.NamaAlamat)
	}

	var Notification notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: Objek.IDPengguna, // ðŸ‘€ NOTE: Tadi lu pake Objek.ID (ID Alamat), kudunya IDPengguna kan?
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸ“ Alamat Baru Berhasil Disimpan!",
		Activity:   true,
		Inbox:      false,
		Archive:    true,
		Pesan:      pesanNotifikasi,
		CreatedAt:  time.Now().Format(time.RFC3339),                  // Pake RFC3339 biar standar ISO & gampang di-parse front-end, jangan .GoString()
		ExpiredAt:  time.Now().AddDate(0, 1, 0).Format(time.RFC3339), // Expired otomatis 1 bulan ke depan (jangan hardcode nama bulan wkwk)
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			// Metadata diisi info general buat kebutuhan tracking/analytics di FE
			Metadata: map[string]interface{}{
				"alamat_id":   Objek.ID,
				"action_type": "create_alamat",
				"platform":    "mobile_app",
			},
			// Special bisa diisi payload koordinat biar kalau notifikasinya diklik, FE bisa langsung buka map
			Special: map[string]interface{}{
				"latitude":     Objek.Latitude,
				"longitude":    Objek.Longitude,
				"click_action": "OPEN_MAPS_DIRECTION",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, Notification, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk); err != nil {
		return err
	}
	return nil
}

func UpdateAlamatPub(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdateAlamatPub"
	var Objek sot_models.AlamatPengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.AlamatPengguna = cass_models.AlamatPengguna{
		ID:              Objek.ID,
		IDPengguna:      Objek.IDPengguna,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodePos,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica sync %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historica db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.AlamatPengguna = se_models.AlamatPengguna{
		ID:              Objek.ID,
		IDPengguna:      Objek.IDPengguna,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodePos,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.AlamatPenggunaIndex.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("Berhasil memasukan data ke dalam search engine dengan antrean UID %s", task_info.IndexUID)
	}

	panggilanAlamatUpdate := strings.ToLower(Objek.PanggilanAlamat)
	var pesanUpdate string

	if strings.Contains(panggilanAlamatUpdate, "rumah") {
		pesanUpdate = fmt.Sprintf("ðŸ”„ Alamat rumah lu ('%s') barusan di-update. Data paling gress udah tersimpan rapi ya!", Objek.NamaAlamat)
	} else {
		pesanUpdate = fmt.Sprintf("âœ¨ Perubahan berhasil disimpan! Alamat '%s' (%s) lu sekarang udah pakai data yang paling baru.", Objek.PanggilanAlamat, Objek.NamaAlamat)
	}

	var NotificationUpdate notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: Objek.IDPengguna, // Pastikan pakai IDPengguna, bukan Objek.ID
		Pengirim:   notification_seeders.Sistem,
		Judul:      "âš™ï¸ Perubahan Alamat Disimpan",
		Pesan:      pesanUpdate,
		Activity:   true,
		Inbox:      false,
		Archive:    true,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 1, 0).Format(time.RFC3339), // Expired 1 bulan ke depan
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"alamat_id":   Objek.ID,
				"action_type": "update_alamat",
				"platform":    "mobile_app",
			},
			Special: map[string]interface{}{
				"latitude":     Objek.Latitude,
				"longitude":    Objek.Longitude,
				"click_action": "REFRESH_ADDRESS_DETAIL", // Perintah ke FE buat refresh view detail
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotificationUpdate, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk); err != nil {
		return fmt.Errorf("gagal mengirim notifikasi update: %w", err)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteAlamatPub(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "DeleteAlamatPub"
	var Objek sot_models.AlamatPengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.AlamatPengguna = cass_models.AlamatPengguna{
		ID:              Objek.ID,
		IDPengguna:      Objek.IDPengguna,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodePos,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica sync %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historica db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.AlamatPengguna = se_models.AlamatPengguna{
		ID:              Objek.ID,
		IDPengguna:      Objek.IDPengguna,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodePos,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	idStr := strconv.FormatInt(ObjekSearchEngine.ID, 10)

	if task_info, err := se_index.AlamatPenggunaIndex.DeleteDocumentWithContext(ctx, idStr, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("Berhasil memasukan data ke dalam search engine dengan antrean UID %s", task_info.IndexUID)
	}

	pesanDelete := fmt.Sprintf("ðŸ—‘ï¸ Alamat '%s' (%s) resmi dihapus dari daftar lu. Tenang, kalau butuh lagi tinggal tambahin baru kok!", Objek.PanggilanAlamat, Objek.NamaAlamat)

	var NotificationDelete notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: Objek.IDPengguna,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸ—‘ï¸ Alamat Berhasil Dihapus",
		Activity:   true,
		Inbox:      false,
		Archive:    true,
		Pesan:      pesanDelete,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339), // Khusus delete, simpan historinya 7 hari aja di notif tab
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"alamat_id":   Objek.ID,
				"action_type": "delete_alamat",
				"platform":    "mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REDIRECT_TO_ADDRESS_LIST", // Perintah ke FE buat balik ke halaman list alamat
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotificationDelete, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk); err != nil {
		return fmt.Errorf("gagal mengirim notifikasi delete: %w", err)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
