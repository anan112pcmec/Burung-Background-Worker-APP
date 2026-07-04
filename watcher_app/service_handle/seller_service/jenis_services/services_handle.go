package jenis_seller_handle

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

func CreateMasukanDataDistributor(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateMasukanDataDistributor"
	var Objek sot_models.DistributorData
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.DistributorData = cass_models.DistributorData{
		ID:                        Objek.ID,
		SellerId:                  Objek.SellerId,
		NamaPerusahaan:            Objek.NamaPerusahaan,
		NIB:                       Objek.NIB,
		NPWP:                      Objek.NPWP,
		DokumenIzinDistributorUrl: Objek.DokumenIzinDistributorUrl,
		Alasan:                    Objek.Alasan,
		Status:                    Objek.Status,
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
			Judul:     "ðŸ“„ Data Distributor Diajukan",
			Pesan:     fmt.Sprintf("Dokumen formalitas untuk perusahaan %s berhasil dikirim. Menunggu proses verifikasi.", Objek.NamaPerusahaan),
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
				Metadata: map[string]interface{}{"seller_id": Objek.SellerId, "distributor_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "DISTRIBUTOR_DATA_SUBMITTED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func UpdateEditDataDistributor(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditDataDistributor"
	var Objek sot_models.DistributorData
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.DistributorData = cass_models.DistributorData{
		ID:                        Objek.ID,
		SellerId:                  Objek.SellerId,
		NamaPerusahaan:            Objek.NamaPerusahaan,
		NIB:                       Objek.NIB,
		NPWP:                      Objek.NPWP,
		DokumenIzinDistributorUrl: Objek.DokumenIzinDistributorUrl,
		Alasan:                    Objek.Alasan,
		Status:                    Objek.Status,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
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
			Judul:     "âœï¸ Data Distributor Diperbarui",
			Pesan:     fmt.Sprintf("Perubahan data formalitas untuk perusahaan %s telah disimpan dan akan diperiksa kembali.", Objek.NamaPerusahaan),
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
				Metadata: map[string]interface{}{"seller_id": Objek.SellerId, "distributor_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "DISTRIBUTOR_DATA_UPDATED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func DeleteHapusDataDistributor(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusDataDistributor"
	var Objek sot_models.DistributorData
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.DistributorData = cass_models.DistributorData{
		ID:                        Objek.ID,
		SellerId:                  Objek.SellerId,
		NamaPerusahaan:            Objek.NamaPerusahaan,
		NIB:                       Objek.NIB,
		NPWP:                      Objek.NPWP,
		DokumenIzinDistributorUrl: Objek.DokumenIzinDistributorUrl,
		Alasan:                    Objek.Alasan,
		Status:                    Objek.Status,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
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
			Judul:     "ðŸ—‘ï¸ Data Distributor Dihapus",
			Pesan:     fmt.Sprintf("Data pengajuan formalitas distributor untuk %s telah dihapus.", Objek.NamaPerusahaan),
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
				Metadata: map[string]interface{}{"seller_id": Objek.SellerId, "distributor_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "DISTRIBUTOR_DATA_REMOVED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func CreateMasukanDataBrand(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateMasukanDataBrand"
	var Objek sot_models.BrandData
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BrandData = cass_models.BrandData{
		ID:                    Objek.ID,
		SellerId:              Objek.SellerId,
		NamaPerusahaan:        Objek.NamaPerusahaan,
		NegaraAsal:            Objek.NegaraAsal,
		LembagaPendaftaran:    Objek.LembagaPendaftaran,
		NomorPendaftaranMerek: Objek.NomorPendaftaranMerek,
		SertifikatMerekUrl:    Objek.SertifikatMerekUrl,
		DokumenPerwakilanUrl:  Objek.DokumenPerwakilanUrl,
		NIB:                   Objek.NIB,
		NPWP:                  Objek.NPWP,
		Alasan:                Objek.Alasan,
		Status:                Objek.Status,
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
			Judul:     "ðŸ“œ Data Merek/Brand Diajukan",
			Pesan:     fmt.Sprintf("Dokumen kepemilikan hak merek untuk %s berhasil dikirim. Menunggu proses verifikasi.", Objek.NamaPerusahaan),
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
				Metadata: map[string]interface{}{"seller_id": Objek.SellerId, "brand_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "BRAND_DATA_SUBMITTED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func UpdateEditDataBrand(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditDataBrand"
	var Objek sot_models.BrandData
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BrandData = cass_models.BrandData{
		ID:                    Objek.ID,
		SellerId:              Objek.SellerId,
		NamaPerusahaan:        Objek.NamaPerusahaan,
		NegaraAsal:            Objek.NegaraAsal,
		LembagaPendaftaran:    Objek.LembagaPendaftaran,
		NomorPendaftaranMerek: Objek.NomorPendaftaranMerek,
		SertifikatMerekUrl:    Objek.SertifikatMerekUrl,
		DokumenPerwakilanUrl:  Objek.DokumenPerwakilanUrl,
		NIB:                   Objek.NIB,
		NPWP:                  Objek.NPWP,
		Alasan:                Objek.Alasan,
		Status:                Objek.Status,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
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
			Judul:     "âœï¸ Data Merek/Brand Diperbarui",
			Pesan:     fmt.Sprintf("Perubahan berkas merek/brand untuk %s telah disimpan dan masuk antrean pengecekan.", Objek.NamaPerusahaan),
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
				Metadata: map[string]interface{}{"seller_id": Objek.SellerId, "brand_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "BRAND_DATA_UPDATED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func DeleteHapusDataBrand(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusDataBrand"
	var Objek sot_models.BrandData
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BrandData = cass_models.BrandData{
		ID:                    Objek.ID,
		SellerId:              Objek.SellerId,
		NamaPerusahaan:        Objek.NamaPerusahaan,
		NegaraAsal:            Objek.NegaraAsal,
		LembagaPendaftaran:    Objek.LembagaPendaftaran,
		NomorPendaftaranMerek: Objek.NomorPendaftaranMerek,
		SertifikatMerekUrl:    Objek.SertifikatMerekUrl,
		DokumenPerwakilanUrl:  Objek.DokumenPerwakilanUrl,
		NIB:                   Objek.NIB,
		NPWP:                  Objek.NPWP,
		Alasan:                Objek.Alasan,
		Status:                Objek.Status,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
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
			Judul:     "ðŸ—‘ï¸ Data Merek/Brand Dihapus",
			Pesan:     fmt.Sprintf("Data pengajuan registrasi merek untuk %s telah dihapus.", Objek.NamaPerusahaan),
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
				Metadata: map[string]interface{}{"seller_id": Objek.SellerId, "brand_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "BRAND_DATA_REMOVED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

