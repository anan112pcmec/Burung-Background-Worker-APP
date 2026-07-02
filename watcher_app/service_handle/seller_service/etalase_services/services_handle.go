package etalase_seller_handle

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/environment"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	notification_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/models"
	notification_request "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/request"
	notification_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/seeders"
)

func CreateTambahEtalaseSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahEtalaseSeller"
	var Objek sot_models.Etalase
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Etalase = cass_models.Etalase{
		ID:           Objek.ID,
		SellerID:     Objek.SellerID,
		Nama:         Objek.Nama,
		Deskripsi:    Objek.Deskripsi,
		JumlahBarang: Objek.JumlahBarang,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 🔔 SISTEM NOTIFIKASI
	if Objek.SellerID != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "📦 Etalase Baru Dibuat",
			Pesan:     fmt.Sprintf("Etalase '%s' telah berhasil dibuat.", Objek.Nama),
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
				Metadata: map[string]interface{}{"seller_id": Objek.SellerID, "etalase_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SELLER_SHOWCASE_CREATED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, environment.HostRunningAPIInNotifikasi, environment.PortRunningAPIInNotifikasi, environment.SellerPathNotifikasiMasuk)
	}

	return nil
}

func UpdateEditEtalaseSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditEtalaseSeller"
	var Objek sot_models.Etalase
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Etalase = cass_models.Etalase{
		ID:           Objek.ID,
		SellerID:     Objek.SellerID,
		Nama:         Objek.Nama,
		Deskripsi:    Objek.Deskripsi,
		JumlahBarang: Objek.JumlahBarang,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	// ID langsung dioper tanpa casting karena sudah berformat int64
	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 🔔 SISTEM NOTIFIKASI
	if Objek.SellerID != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "✏️ Etalase Diperbarui",
			Pesan:     fmt.Sprintf("Informasi etalase '%s' berhasil diubah.", Objek.Nama),
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
				Metadata: map[string]interface{}{"seller_id": Objek.SellerID, "etalase_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SELLER_SHOWCASE_UPDATED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, environment.HostRunningAPIInNotifikasi, environment.PortRunningAPIInNotifikasi, environment.SellerPathNotifikasiMasuk)
	}

	return nil
}

func DeleteHapusEtalaseSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusEtalaseSeller"
	var Objek sot_models.Etalase
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Etalase = cass_models.Etalase{
		ID:           Objek.ID,
		SellerID:     Objek.SellerID,
		Nama:         Objek.Nama,
		Deskripsi:    Objek.Deskripsi,
		JumlahBarang: Objek.JumlahBarang,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	// Aksi delete tetap melakukan INSERT baru ke historical db sesuai prinsip append-only
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 🔔 SISTEM NOTIFIKASI
	if Objek.SellerID != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "🗑️ Etalase Dihapus",
			Pesan:     fmt.Sprintf("Etalase '%s' telah berhasil dihapus.", Objek.Nama),
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
				Metadata: map[string]interface{}{"seller_id": Objek.SellerID, "etalase_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SELLER_SHOWCASE_REMOVED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, environment.HostRunningAPIInNotifikasi, environment.PortRunningAPIInNotifikasi, environment.SellerPathNotifikasiMasuk)
	}

	return nil
}

func CreateTambahkanBarangKeEtalase(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahkanBarangKeEtalase"
	var Objek sot_models.BarangKeEtalase
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BarangKeEtalase = cass_models.BarangKeEtalase{
		ID:            Objek.ID,
		IdEtalase:     Objek.IdEtalase,
		IdBarangInduk: Objek.IdBarangInduk,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 🔔 SISTEM NOTIFIKASI (Silent Update UI Lokal Seller)
	var Notifikasi = notification_models.NotificationSeller{
		IDSeller:  0, // Handler service akan memproses lewat data payload map/special jika ditangkap di konsumen API internal
		Pengirim:  notification_seeders.Sistem,
		Judul:     "🔄 Produk Ditambahkan ke Etalase",
		Pesan:     "Item produk berhasil dikaitkan ke dalam susunan etalase.",
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
			Metadata: map[string]interface{}{"barang_ke_etalase_id": Objek.ID, "etalase_id": Objek.IdEtalase, "barang_induk_id": Objek.IdBarangInduk},
			Special:  map[string]interface{}{"click_action": "SILENT_ADD_PRODUCT_TO_SHOWCASE"},
		},
	}
	_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, environment.HostRunningAPIInNotifikasi, environment.PortRunningAPIInNotifikasi, environment.SellerPathNotifikasiMasuk)

	return nil
}

func DeleteHapusBarangDariEtalase(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusBarangDariEtalase"
	var Objek sot_models.BarangKeEtalase
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BarangKeEtalase = cass_models.BarangKeEtalase{
		ID:            Objek.ID,
		IdEtalase:     Objek.IdEtalase,
		IdBarangInduk: Objek.IdBarangInduk,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	// Aksi delete tetap melakukan INSERT baru ke historical db demi kelengkapan audit log
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 🔔 SISTEM NOTIFIKASI (Silent Update UI Lokal Seller)
	var Notifikasi = notification_models.NotificationSeller{
		IDSeller:  0,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "🔄 Produk Dilepas dari Etalase",
		Pesan:     "Kaitan produk pada etalase pilihan telah berhasil dilepas.",
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
			Metadata: map[string]interface{}{"barang_ke_etalase_id": Objek.ID, "etalase_id": Objek.IdEtalase, "barang_induk_id": Objek.IdBarangInduk},
			Special:  map[string]interface{}{"click_action": "SILENT_REMOVE_PRODUCT_FROM_SHOWCASE"},
		},
	}
	_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, environment.HostRunningAPIInNotifikasi, environment.PortRunningAPIInNotifikasi, environment.SellerPathNotifikasiMasuk)

	return nil
}
