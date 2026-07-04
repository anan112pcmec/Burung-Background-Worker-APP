package diskon_seller_handle

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/cache"
	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	notification_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/models"
	notification_request "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/request"
	notification_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/seeders"
)

func CreateTambahDiskonProduk(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahDiskonProduk"
	var Objek sot_models.DiskonProduk
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.DiskonProduk = cass_models.DiskonProduk{
		ID:            Objek.ID,
		SellerId:      Objek.SellerId,
		Nama:          Objek.Nama,
		Deskripsi:     Objek.Deskripsi,
		DiskonPersen:  Objek.DiskonPersen,
		BerlakuMulai:  Objek.BerlakuMulai,
		BerlakuSampai: Objek.BerlakuSampai,
		Status:        Objek.Status,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
		DeletedAt:     Objek.DeletedAt, // Mengambil nilai Time dari gorm.DeletedAt jika tipenya gorm
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” SISTEM NOTIFIKASI
	if Objek.SellerId != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerId),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ·ï¸ Diskon Baru Dibuat",
			Pesan:     fmt.Sprintf("Promo diskon '%s' sebesar %d%% berhasil dibuat.", Objek.Nama, Objek.DiskonPersen),
			Pop:       1,
			Archive:   false,
			Inbox:     true,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.SellerId, "diskon_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "PROMO_DISCOUNT_CREATED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func UpdateEditDiskonProduk(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditDiskonProduk"
	var Objek sot_models.DiskonProduk
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.DiskonProduk = cass_models.DiskonProduk{
		ID:            Objek.ID,
		SellerId:      Objek.SellerId,
		Nama:          Objek.Nama,
		Deskripsi:     Objek.Deskripsi,
		DiskonPersen:  Objek.DiskonPersen,
		BerlakuMulai:  Objek.BerlakuMulai,
		BerlakuSampai: Objek.BerlakuSampai,
		Status:        Objek.Status,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
		DeletedAt:     Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	// ID tidak perlu dicasting karena sudah berformat int64
	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” SISTEM NOTIFIKASI
	if Objek.SellerId != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerId),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "âœï¸ Promo Diskon Diubah",
			Pesan:     fmt.Sprintf("Informasi promo diskon '%s' telah berhasil diperbarui.", Objek.Nama),
			Pop:       1,
			Archive:   false,
			Inbox:     true,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.SellerId, "diskon_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "PROMO_DISCOUNT_EDITED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func DeleteHapusDiskonProduk(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusDiskonProduk"
	var Objek sot_models.DiskonProduk
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.DiskonProduk = cass_models.DiskonProduk{
		ID:            Objek.ID,
		SellerId:      Objek.SellerId,
		Nama:          Objek.Nama,
		Deskripsi:     Objek.Deskripsi,
		DiskonPersen:  Objek.DiskonPersen,
		BerlakuMulai:  Objek.BerlakuMulai,
		BerlakuSampai: Objek.BerlakuSampai,
		Status:        Objek.Status,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
		DeletedAt:     Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	// Aksi delete tetap melakukan INSERT ke historical db sesuai prinsip append-only kita
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” SISTEM NOTIFIKASI
	if Objek.SellerId != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerId),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ—‘ï¸ Promo Diskon Dihapus",
			Pesan:     fmt.Sprintf("Promo diskon '%s' telah berhasil dihapus.", Objek.Nama),
			Pop:       1,
			Archive:   false,
			Inbox:     true,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.SellerId, "diskon_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "PROMO_DISCOUNT_REMOVED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func CreateTetapkanDiskonPadaBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTetapkanDiskonPadaBarang"
	var Objek sot_models.BarangDiDiskon
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BarangDiDiskon = cass_models.BarangDiDiskon{
		ID:               Objek.ID,
		SellerId:         Objek.SellerId,
		IdDiskon:         Objek.IdDiskon,
		IdBarangInduk:    Objek.IdBarangInduk,
		IdKategoriBarang: Objek.IdKategoriBarang,
		Status:           Objek.Status,
		CreatedAt:        Objek.CreatedAt,
		UpdatedAt:        Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” SISTEM NOTIFIKASI
	if Objek.SellerId != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerId),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Promo Diterapkan pada Produk",
			Pesan:     "Diskon baru telah sukses diterapkan ke item produk pilihan.",
			Pop:       0,
			Archive:   true,
			Inbox:     false,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.SellerId, "barang_di_diskon_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_APPLY_PROMO_PRODUCT"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func DeleteHapusDiskonPadaBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusDiskonPadaBarang"
	var Objek sot_models.BarangDiDiskon
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BarangDiDiskon = cass_models.BarangDiDiskon{
		ID:               Objek.ID,
		SellerId:         Objek.SellerId,
		IdDiskon:         Objek.IdDiskon,
		IdBarangInduk:    Objek.IdBarangInduk,
		IdKategoriBarang: Objek.IdKategoriBarang,
		Status:           Objek.Status,
		CreatedAt:        Objek.CreatedAt,
		UpdatedAt:        Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	// Aksi delete tetap melakukan INSERT baru ke historical db
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” SISTEM NOTIFIKASI
	if Objek.SellerId != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerId),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Promo Dilepas dari Produk",
			Pesan:     "Kaitan diskon pada item produk pilihan telah berhasil dilepas.",
			Pop:       0,
			Archive:   true,
			Inbox:     false,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.SellerId, "barang_di_diskon_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_REMOVE_PROMO_PRODUCT"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

