package alamat_seller_handle

import (
	"context"
	"fmt"
	"strconv"
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

func CreateTambahAlamatGudang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateTambahAlamatGudang"
	var Objek sot_models.AlamatGudang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.AlamatGudang = cass_models.AlamatGudang{
		ID:              Objek.ID,
		IDSeller:        Objek.IDSeller,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodeNegara, // Catatan: Jika ini field KodePos, pastikan Objek.KodePos yang dimasukkan jika tersedia di struct asal
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.AlamatGudang = se_models.AlamatGudang{
		ID:              Objek.ID,
		IDSeller:        Objek.IDSeller,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodeNegara,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       &Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.AlamatGudang.AddDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke dalam search engine dengan uid: %s", task_info.IndexUID)
	}

	// ðŸ”” Alert Notifikasi: Pop tinggi untuk alert keamanan alamat baru
	var Notifikasi notification_models.NotificationSeller = notification_models.NotificationSeller{
		IDSeller:  int64(Objek.IDSeller),
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ  Penambahan Alamat Gudang Baru",
		Pesan:     "Halo, kami mendeteksi penambahan alamat gudang baru di akun Anda.",
		Pop:       0.9,
		Archive:   true,
		Inbox:     false,
		Activity:  true,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 2).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"seller_id": Objek.IDSeller, "alamat_gudang_id": Objek.ID},
			Special:  map[string]interface{}{"click_action": "REFRESH_ALAMAT_GUDANG"},
		},
	}

	if err := notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk); err != nil {
		return err
	}

	return nil
}

func UpdateEditAlamatGudang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdateEditAlamatGudang"
	var Objek sot_models.AlamatGudang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.AlamatGudang = cass_models.AlamatGudang{
		ID:              Objek.ID,
		IDSeller:        Objek.IDSeller,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodeNegara,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.AlamatGudang = se_models.AlamatGudang{
		ID:              Objek.ID,
		IDSeller:        Objek.IDSeller,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodeNegara,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       &Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.AlamatGudang.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke dalam search engine dengan uid: %s", task_info.IndexUID)
	}

	// ðŸ”” Silent Update Seller: Sinkronisasi data perubahan alamat di background app
	if Objek.IDSeller != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.IDSeller),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Alamat Gudang Diperbarui",
			Pesan:     "Perubahan data alamat gudang Anda telah berhasil disinkronisasi.",
			Pop:       0, // Background sync, gausah di-pop
			Archive:   true,
			Inbox:     false,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.IDSeller, "alamat_gudang_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_ALAMAT_GUDANG"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func DeleteHapusAlamatGudang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "DeleteHapusAlamatGudang"
	var Objek sot_models.AlamatGudang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.AlamatGudang = cass_models.AlamatGudang{
		ID:              Objek.ID,
		IDSeller:        Objek.IDSeller,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodeNegara,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	idStr := strconv.FormatInt(ObjekCass.ID, 10)

	if task_info, err := se_index.AlamatGudang.DeleteDocumentWithContext(ctx, idStr, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke dalam search engine dengan uid: %s", task_info.IndexUID)
	}

	// ðŸ”” Silent Update Seller: Hapus data alamat dari list UI lokal perangkat seller
	if Objek.IDSeller != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.IDSeller),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ—‘ï¸ Alamat Gudang Dihapus",
			Pesan:     "Alamat gudang Anda telah berhasil dihapus dari sistem.",
			Pop:       0, // Silent sync background, gausah di-pop
			Archive:   true,
			Inbox:     false,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.IDSeller, "alamat_gudang_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_REMOVE_ALAMAT_GUDANG"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}
