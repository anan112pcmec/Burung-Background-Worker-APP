// INFOO: INI KODE YANG DI KASIH NOTIFIKASI

package media_seller_handle

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

func CreateTambahFotoProfilSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahFotoProfilSeller" // Disesuaikan agar konsisten dengan kata kerja Create
	var Objek sot_models.MediaSellerProfilFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaSellerProfilFoto = cass_models.MediaSellerProfilFoto{
		ID:        Objek.ID,
		IdSeller:  Objek.IdSeller,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
		UpdatedAt: Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Notifikasi ke Seller (Foto Profil Toko Berhasil Ditambahkan)
	if Objek.IdSeller != 0 {
		var NotifSeller = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.IdSeller),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ“¸ Foto Profil Toko Berhasil Diunggah!",
			Pesan:     "Foto profil toko kamu udah aktif dan tampil di halaman toko. Tampilan makin meyakinkan buat calon pembeli, nih!",
			Pop:       3.0,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_media": Objek.ID, "id_seller": Objek.IdSeller},
				Special:  map[string]interface{}{"click_action": "OPEN_SELLER_PROFILE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func UpdateUbahFotoProfilSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahFotoProfilSeller"
	var Objek sot_models.MediaSellerProfilFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaSellerProfilFoto = cass_models.MediaSellerProfilFoto{
		ID:        Objek.ID,
		IdSeller:  Objek.IdSeller,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
		UpdatedAt: Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Notifikasi ke Seller (Foto Profil Toko Berhasil Diperbarui)
	if Objek.IdSeller != 0 {
		var NotifSeller = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.IdSeller),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Foto Profil Toko Diperbarui!",
			Pesan:     "Foto profil toko kamu berhasil diganti dengan yang baru. Update kece, biar makin dikenal pembeli!",
			Pop:       3.0,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_media": Objek.ID, "id_seller": Objek.IdSeller},
				Special:  map[string]interface{}{"click_action": "OPEN_SELLER_PROFILE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func DeleteHapusFotoProfilSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusFotoProfilSeller"
	var Objek sot_models.MediaSellerProfilFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaSellerProfilFoto = cass_models.MediaSellerProfilFoto{
		ID:        Objek.ID,
		IdSeller:  Objek.IdSeller,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
		UpdatedAt: Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateTambahFotoBannerSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahFotoBannerSeller" // Disesuaikan agar seragam dengan kata kerja Create
	var Objek sot_models.MediaSellerBannerFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaSellerBannerFoto = cass_models.MediaSellerBannerFoto{
		ID:        Objek.ID,
		IdSeller:  Objek.IdSeller,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
		UpdatedAt: Objek.UpdatedAt,
		// Bertipe time.Time sesuai draf bawaanmu
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahFotoBannerSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahFotoBannerSeller"
	var Objek sot_models.MediaSellerBannerFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaSellerBannerFoto = cass_models.MediaSellerBannerFoto{
		ID:        Objek.ID,
		IdSeller:  Objek.IdSeller,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
		UpdatedAt: Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusFotoBannerSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusFotoBannerSeller"
	var Objek sot_models.MediaSellerBannerFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaSellerBannerFoto = cass_models.MediaSellerBannerFoto{
		ID:        Objek.ID,
		IdSeller:  Objek.IdSeller,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
		UpdatedAt: Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}
func CreateTambahkanFotoTokoFisikSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahkanFotoTokoFisikSeller"
	var Objek sot_models.MediaSellerTokoFisikFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaSellerTokoFisikFoto = cass_models.MediaSellerTokoFisikFoto{
		ID:        Objek.ID,
		IdSeller:  Objek.IdSeller,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
		// Bertipe time.Time sesuai draf bawaanmu
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusFotoTokoFisikSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusFotoTokoFisikSeller"
	var Objek sot_models.MediaSellerTokoFisikFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaSellerTokoFisikFoto = cass_models.MediaSellerTokoFisikFoto{
		ID:        Objek.ID,
		IdSeller:  Objek.IdSeller,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}
func CreateTambahFotoEtalaseSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahFotoEtalaseSeller" // Disesuaikan dengan kata kerja Create
	var Objek sot_models.MediaEtalaseFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaEtalaseFoto = cass_models.MediaEtalaseFoto{
		ID:        Objek.ID,
		IdEtalase: Objek.IdEtalase,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
		UpdatedAt: Objek.UpdatedAt,
		// Bertipe time.Time sesuai draf bawaanmu
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahFotoEtalaseSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahFotoEtalaseSeller"
	var Objek sot_models.MediaEtalaseFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaEtalaseFoto = cass_models.MediaEtalaseFoto{
		ID:        Objek.ID,
		IdEtalase: Objek.IdEtalase,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
		UpdatedAt: Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusFotoEtalaseSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahFotoEtalaseSeller"
	var Objek sot_models.MediaEtalaseFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaEtalaseFoto = cass_models.MediaEtalaseFoto{
		ID:        Objek.ID,
		IdEtalase: Objek.IdEtalase,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
		UpdatedAt: Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateTambahkanMediaBarangIndukFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahkanMediaBarangIndukFoto"
	var Objek sot_models.MediaBarangIndukFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBarangIndukFoto = cass_models.MediaBarangIndukFoto{
		ID:            Objek.ID,
		IdBarangInduk: Objek.IdBarangInduk,
		Key:           Objek.Key,
		Format:        Objek.Format,
		CreatedAt:     Objek.CreatedAt,
		// Bertipe time.Time sesuai draf bawaanmu
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusMediaBarangIndukFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusMediaBarangIndukFoto"
	var Objek sot_models.MediaBarangIndukFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBarangIndukFoto = cass_models.MediaBarangIndukFoto{
		ID:            Objek.ID,
		IdBarangInduk: Objek.IdBarangInduk,
		Key:           Objek.Key,
		Format:        Objek.Format,
		CreatedAt:     Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateTambahBarangIndukVideo(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahBarangIndukVideo" // Disesuaikan dengan kata kerja Create
	var Objek sot_models.MediaBarangIndukVideo
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBarangIndukVideo = cass_models.MediaBarangIndukVideo{
		ID:            Objek.ID,
		IdBarangInduk: Objek.IdBarangInduk,
		Key:           Objek.Key,
		Format:        Objek.Format,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
		// Konversi gorm.DeletedAt ke time.Time untuk Cassandra
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahBarangIndukVideo(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahBarangIndukVideo"
	var Objek sot_models.MediaBarangIndukVideo
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBarangIndukVideo = cass_models.MediaBarangIndukVideo{
		ID:            Objek.ID,
		IdBarangInduk: Objek.IdBarangInduk,
		Key:           Objek.Key,
		Format:        Objek.Format,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusBarangIndukVideo(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusBarangIndukVideo"
	var Objek sot_models.MediaBarangIndukVideo
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBarangIndukVideo = cass_models.MediaBarangIndukVideo{
		ID:            Objek.ID,
		IdBarangInduk: Objek.IdBarangInduk,
		Key:           Objek.Key,
		Format:        Objek.Format,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}
func CreateTambahKategoriBarangFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahKategoriBarangFoto" // Diselaraskan agar konsisten dengan kata kerja Create
	var Objek sot_models.MediaKategoriBarangFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaKategoriBarangFoto = cass_models.MediaKategoriBarangFoto{
		ID:               Objek.ID,
		IdKategoriBarang: Objek.IdKategoriBarang,
		IdBarangInduk:    Objek.IdBarangInduk,
		Key:              Objek.Key,
		Format:           Objek.Format,
		CreatedAt:        Objek.CreatedAt,
		UpdatedAt:        Objek.UpdatedAt,
		// Bertipe time.Time sesuai draf bawaanmu
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahKategoriBarangFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahKategoriBarangFoto"
	var Objek sot_models.MediaKategoriBarangFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaKategoriBarangFoto = cass_models.MediaKategoriBarangFoto{
		ID:               Objek.ID,
		IdKategoriBarang: Objek.IdKategoriBarang,
		IdBarangInduk:    Objek.IdBarangInduk,
		Key:              Objek.Key,
		Format:           Objek.Format,
		CreatedAt:        Objek.CreatedAt,
		UpdatedAt:        Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusKategoriBarangFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusKategoriBarangFoto"
	var Objek sot_models.MediaKategoriBarangFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaKategoriBarangFoto = cass_models.MediaKategoriBarangFoto{
		ID:               Objek.ID,
		IdKategoriBarang: Objek.IdKategoriBarang,
		IdBarangInduk:    Objek.IdBarangInduk,
		Key:              Objek.Key,
		Format:           Objek.Format,
		CreatedAt:        Objek.CreatedAt,
		UpdatedAt:        Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateTambahDistributorDataDokumen(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahDistributorDataDokumen"
	var Objek sot_models.MediaDistributorDataDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaDistributorDataDokumen = cass_models.MediaDistributorDataDokumen{
		ID:                Objek.ID,
		IdDistributorData: Objek.IdDistributorData,
		Key:               Objek.Key,
		Format:            Objek.Format,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		// Bertipe time.Time sesuai draf bawaanmu
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahDistributorDataDokumen(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahDistributorDataDokumen" // Diselaraskan dari UpdateTambah...
	var Objek sot_models.MediaDistributorDataDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaDistributorDataDokumen = cass_models.MediaDistributorDataDokumen{
		ID:                Objek.ID,
		IdDistributorData: Objek.IdDistributorData,
		Key:               Objek.Key,
		Format:            Objek.Format,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusMediaDistributorDataDokumen(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusMediaDistributorDataDokumen"
	var Objek sot_models.MediaDistributorDataDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaDistributorDataDokumen = cass_models.MediaDistributorDataDokumen{
		ID:                Objek.ID,
		IdDistributorData: Objek.IdDistributorData,
		Key:               Objek.Key,
		Format:            Objek.Format,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateTambahMediaDistributorDataNPWPFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaDistributorDataNPWPFoto"
	var Objek sot_models.MediaDistributorDataNPWPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaDistributorDataNPWPFoto = cass_models.MediaDistributorDataNPWPFoto{
		ID:                Objek.ID,
		IdDistributorData: Objek.IdDistributorData,
		Key:               Objek.Key,
		Format:            Objek.Format,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		// Bertipe time.Time sesuai draf bawaanmu
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahMediaDistributorDataNPWPFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahMediaDistributorDataNPWPFoto" // Diselaraskan dari UpdateTambah...
	var Objek sot_models.MediaDistributorDataNPWPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaDistributorDataNPWPFoto = cass_models.MediaDistributorDataNPWPFoto{
		ID:                Objek.ID,
		IdDistributorData: Objek.IdDistributorData,
		Key:               Objek.Key,
		Format:            Objek.Format,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusMediaDistributorDataNPWPFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusMediaDistributorDataNPWPFoto"
	var Objek sot_models.MediaDistributorDataNPWPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaDistributorDataNPWPFoto = cass_models.MediaDistributorDataNPWPFoto{
		ID:                Objek.ID,
		IdDistributorData: Objek.IdDistributorData,
		Key:               Objek.Key,
		Format:            Objek.Format,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}
func CreateTambahDistributorDataNIBFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahDistributorDataNIBFoto"
	var Objek sot_models.MediaDistributorDataNIBFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaDistributorDataNIBFoto = cass_models.MediaDistributorDataNIBFoto{
		ID:                Objek.ID,
		IdDistributorData: Objek.IdDistributorData,
		Key:               Objek.Key,
		Format:            Objek.Format,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		// Bertipe time.Time sesuai draf bawaanmu
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahDistributorDataNIBFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahDistributorDataNIBFoto" // Diselaraskan dari UpdateTambah...
	var Objek sot_models.MediaDistributorDataNIBFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaDistributorDataNIBFoto = cass_models.MediaDistributorDataNIBFoto{
		ID:                Objek.ID,
		IdDistributorData: Objek.IdDistributorData,
		Key:               Objek.Key,
		Format:            Objek.Format,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusDistributorDataNIBFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusDistributorDataNIBFoto"
	var Objek sot_models.MediaDistributorDataNIBFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaDistributorDataNIBFoto = cass_models.MediaDistributorDataNIBFoto{
		ID:                Objek.ID,
		IdDistributorData: Objek.IdDistributorData,
		Key:               Objek.Key,
		Format:            Objek.Format,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateTambahDistributorDataSuratKerjasamaDokumen(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahDistributorDataSuratKerjasamaDokumen"
	var Objek sot_models.MediaDistributorDataSuratKerjasamaDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaDistributorDataSuratKerjasamaDokumen = cass_models.MediaDistributorDataSuratKerjasamaDokumen{
		ID:                Objek.ID,
		IdDistributorData: Objek.IdDistributorData,
		Key:               Objek.Key,
		Format:            Objek.Format,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		// Bertipe time.Time sesuai draf bawaanmu
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahDistributorDataSuratKerjasamaDokumen(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahDistributorDataSuratKerjasamaDokumen" // Diselaraskan dari UpdateTambah...
	var Objek sot_models.MediaDistributorDataSuratKerjasamaDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaDistributorDataSuratKerjasamaDokumen = cass_models.MediaDistributorDataSuratKerjasamaDokumen{
		ID:                Objek.ID,
		IdDistributorData: Objek.IdDistributorData,
		Key:               Objek.Key,
		Format:            Objek.Format,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusDistributorDataSuratKerjasamaDokumen(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusDistributorDataSuratKerjasamaDokumen"
	var Objek sot_models.MediaDistributorDataSuratKerjasamaDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaDistributorDataSuratKerjasamaDokumen = cass_models.MediaDistributorDataSuratKerjasamaDokumen{
		ID:                Objek.ID,
		IdDistributorData: Objek.IdDistributorData,
		Key:               Objek.Key,
		Format:            Objek.Format,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateTambahBrandDataPerwakilanDokumen(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahBrandDataPerwakilanDokumen"
	var Objek sot_models.MediaBrandDataPerwakilanDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataPerwakilanDokumen = cass_models.MediaBrandDataPerwakilanDokumen{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahBrandDataPerwakilanDokumen(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahBrandDataPerwakilanDokumen" // Diselaraskan dari UpdateTambah...
	var Objek sot_models.MediaBrandDataPerwakilanDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataPerwakilanDokumen = cass_models.MediaBrandDataPerwakilanDokumen{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusBrandDataPerwakilanDokumen(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusBrandDataPerwakilanDokumen"
	var Objek sot_models.MediaBrandDataPerwakilanDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataPerwakilanDokumen = cass_models.MediaBrandDataPerwakilanDokumen{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateTambahMediaBrandDataSertifikatFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaBrandDataSertifikatFoto"
	var Objek sot_models.MediaBrandDataSertifikatFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataSertifikatFoto = cass_models.MediaBrandDataSertifikatFoto{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahMediaBrandDataSertifikatFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahMediaBrandDataSertifikatFoto" // Diselaraskan dari UpdateTambah...
	var Objek sot_models.MediaBrandDataSertifikatFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataSertifikatFoto = cass_models.MediaBrandDataSertifikatFoto{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusMediaBrandDataSertifikatFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusMediaBrandDataSertifikatFoto"
	var Objek sot_models.MediaBrandDataSertifikatFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataSertifikatFoto = cass_models.MediaBrandDataSertifikatFoto{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}
func CreateTambahMediaBrandDataNIBFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaBrandDataNIBFoto"
	var Objek sot_models.MediaBrandDataNIBFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataNIBFoto = cass_models.MediaBrandDataNIBFoto{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahMediaBrandDataNIBFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahMediaBrandDataNIBFoto" // Diselaraskan dari UpdateTambah...
	var Objek sot_models.MediaBrandDataNIBFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataNIBFoto = cass_models.MediaBrandDataNIBFoto{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusMediaBrandDataNIBFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusMediaBrandDataNIBFoto"
	var Objek sot_models.MediaBrandDataNIBFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataNIBFoto = cass_models.MediaBrandDataNIBFoto{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateTambahMediaBrandNPWPFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaBrandNPWPFoto"
	var Objek sot_models.MediaBrandDataNPWPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataNPWPFoto = cass_models.MediaBrandDataNPWPFoto{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahMediaBrandNPWPFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahMediaBrandNPWPFoto" // Diselaraskan dari UpdateTambah...
	var Objek sot_models.MediaBrandDataNPWPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataNPWPFoto = cass_models.MediaBrandDataNPWPFoto{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusMediaBrandNPWPFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusMediaBrandNPWPFoto"
	var Objek sot_models.MediaBrandDataNPWPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataNPWPFoto = cass_models.MediaBrandDataNPWPFoto{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}
func CreateTambahMediaBrandDataLogoFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaBrandDataLogoFoto"
	var Objek sot_models.MediaBrandDataLogoFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataLogoFoto = cass_models.MediaBrandDataLogoFoto{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahMediaBrandDataLogoFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahMediaBrandDataLogoFoto" // Diselaraskan dari UpdateTambah...
	var Objek sot_models.MediaBrandDataLogoFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataLogoFoto = cass_models.MediaBrandDataLogoFoto{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusMediaBrandDataLogoFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusMediaBrandDataLogoFoto"
	var Objek sot_models.MediaBrandDataLogoFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataLogoFoto = cass_models.MediaBrandDataLogoFoto{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateTambahBrandDataSuratKerjasamaDokumen(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahBrandDataSuratKerjasamaDokumen"
	var Objek sot_models.MediaBrandDataSuratKerjasamaDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataSuratKerjasamaDokumen = cass_models.MediaBrandDataSuratKerjasamaDokumen{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData, // Dipetakan agar ID relasi tersimpan rapi di Cassandra
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahBrandDataSuratKerjasamaDokumen(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahBrandDataSuratKerjasamaDokumen" // Diselaraskan dari UpdateTambah...
	var Objek sot_models.MediaBrandDataSuratKerjasamaDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataSuratKerjasamaDokumen = cass_models.MediaBrandDataSuratKerjasamaDokumen{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusBrandDataSuratKerjasamaDokumen(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusBrandDataSuratKerjasamaDokumen"
	var Objek sot_models.MediaBrandDataSuratKerjasamaDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaBrandDataSuratKerjasamaDokumen = cass_models.MediaBrandDataSuratKerjasamaDokumen{
		ID:          Objek.ID,
		IdBrandData: Objek.IdBrandData,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateTambahMediaTransaksiApprovedFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaTransaksiApprovedFoto"
	var Objek sot_models.MediaTransaksiApprovedFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaTransaksiApprovedFoto = cass_models.MediaTransaksiApprovedFoto{
		ID:          Objek.ID,
		IdTransaksi: Objek.IdTransaksi,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateTambahMediaTransaksiApprovedVideo(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaTransaksiApprovedVideo"
	var Objek sot_models.MediaTransaksiApprovedVideo
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	// Mengikuti struktur field yang sama namun dengan tipe model Video
	var ObjekCass cass_models.MediaTransaksiApprovedVideo = cass_models.MediaTransaksiApprovedVideo{
		ID:          Objek.ID,
		IdTransaksi: Objek.IdTransaksi,
		Key:         Objek.Key,
		Format:      Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}
