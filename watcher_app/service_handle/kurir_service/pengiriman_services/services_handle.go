package pengiriman_kurir_handle

import (
	"context"
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"

	cache_db_function "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cache_db/function"
	cache_db_session "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cache_db/session"
	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	se_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/search_engine/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func CreateAktifkanBidKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateAktifkanBidKurir"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         Objek.Selesai, // Dikembalikan asli sesuai properti bawaanmu
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt, // Ekstraksi gorm.DeletedAt ke time.Time untuk Cassandra
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateUbahBidKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahBidKurir"
	var Objek sot_models.BidKurirData // Diselaraskan pakai BidKurirData sesuai pasangannya

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePosisiBidKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatePosisiBidKurir"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data posisi ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data posisi ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateAmbilPengirimanNonEksManualRegulerIIBidKurirNonEksSchedulerCreatePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateAmbilPengirimanNonEksManualRegulerIIBidKurirNonEksSchedulerCreatePublish"
	var Objek sot_models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirNonEksScheduler = cass_models.BidKurirNonEksScheduler{
		ID:           Objek.ID,
		IdBid:        Objek.IdBid,
		IdKurir:      Objek.IdKurir,
		Urutan:       Objek.Urutan,
		IdPengiriman: Objek.IdPengiriman,
		Status:       Objek.Status,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAmbilPengirimanNonEksManualRegulerIIpengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateAmbilPengirimanNonEksManualRegulerIIpengirimanUpdatedPublish"
	var Objek sot_models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.Pengiriman = cass_models.Pengiriman{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatPengguna:  Objek.IdAlamatPengguna,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataAmbilPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataAmbilPengirimanUpdatedPublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataStatusUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataStatusUpdatedPublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateAmbilPengirimanEksManualRegulerIIbidKurirEksSchedulerCreatePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateAmbilPengirimanEksManualRegulerIIbidKurirEksSchedulerCreatePublish"
	var Objek sot_models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirEksScheduler = cass_models.BidKurirEksScheduler{
		ID:              Objek.ID,
		IdBid:           Objek.IdBid,
		IdKurir:         Objek.IdKurir,
		Urutan:          Objek.Urutan,
		IdPengirimanEks: Objek.IdPengirimanEks,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAmbilPengirimanEksManualRegulerIIpengirimanEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateAmbilPengirimanEksManualRegulerIIpengirimanEksUpdatedPublish"
	var Objek sot_models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.Pengiriman = cass_models.Pengiriman{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatPengguna:  Objek.IdAlamatPengguna,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataUpdatedPublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataAmbilPengirimanEksStatusUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataAmbilPengirimanEksStatusUpdatedPublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateLockSiapAntarBidKurirIIEksScheduler(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateLockSiapAntarBidKurirIIEksScheduler"
	var Objek sot_models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirEksScheduler = cass_models.BidKurirEksScheduler{
		ID:              Objek.ID,
		IdBid:           Objek.IdBid,
		IdKurir:         Objek.IdKurir,
		Urutan:          Objek.Urutan,
		IdPengirimanEks: Objek.IdPengirimanEks,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateLockSiapAntarBidKurirIINonEksScheduler(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateLockSiapAntarBidKurirIINonEksScheduler"
	var Objek sot_models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirNonEksScheduler = cass_models.BidKurirNonEksScheduler{
		ID:           Objek.ID,
		IdBid:        Objek.IdBid,
		IdKurir:      Objek.IdKurir,
		Urutan:       Objek.Urutan,
		IdPengiriman: Objek.IdPengiriman,
		Status:       Objek.Status,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateLockSiapAntarBidKurirIIbidKurirDataLockSiapAntarUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateLockSiapAntarBidKurirIIbidKurirDataLockSiapAntarUpdatedPublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
func UpdateLockSiapAntarBidKurirIIkurirLockSiapAntarUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services = "UpdateLockSiapAntarBidKurirIIKurirLockSiapAntarUpdatedPublish"
	var Objek sot_models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.Kurir = cass_models.Kurir{
		ID:            Objek.ID,
		Nama:          Objek.Nama,
		Username:      Objek.Username,
		Email:         Objek.Email,
		Jenis:         Objek.Jenis,
		PasswordHash:  Objek.PasswordHash,
		Deskripsi:     Objek.Deskripsi,
		StatusKurir:   Objek.StatusKurir,
		StatusBid:     Objek.StatusBid,
		VerifiedKurir: Objek.VerifiedKurir,
		Rating:        Objek.Rating,
		TipeKendaraan: Objek.TipeKendaraan,
		CreatedAt:     Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}
	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Kurir = se_models.Kurir{
		ID:            Objek.ID,
		Nama:          Objek.Nama,
		Username:      Objek.Username,
		Email:         Objek.Email,
		Jenis:         Objek.Jenis,
		PasswordHash:  Objek.PasswordHash,
		Deskripsi:     Objek.Deskripsi,
		StatusKurir:   Objek.StatusKurir,
		StatusBid:     Objek.StatusBid,
		VerifiedKurir: Objek.VerifiedKurir,
		Rating:        Objek.Rating,
		TipeKendaraan: Objek.TipeKendaraan,
		CreatedAt:     Objek.CreatedAt,
	}

	if task_info, err := se_index.KurirIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke search engine dengan info: %s ", task_info.IndexUID)
	}

	if err := cache_db_function.UpdateSessionData(ctx, *rds_session, cache_db_session.GetSessionKey(&Objek), Objek); err != nil {
		return fmt.Errorf("gagal memperbarui data di cache %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreatePickedUpPengirimanNonEksIIjejakPengirimanCreatePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreatePickedUpPengirimanNonEksIIjejakPengirimanCreatePublish"
	var Objek sot_models.JejakPengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengiriman = cass_models.JejakPengiriman{
		ID:           Objek.ID,
		IdPengiriman: Objek.IdPengiriman,
		Lokasi:       Objek.Lokasi,
		Keterangan:   Objek.Keterangan,
		Latitude:     Objek.Latitude,
		Longtitude:   Objek.Longtitude,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanNonEksIIschedulerPickedUpNonEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatePickedUpPengirimanNonEksIIschedulerPickedUpNonEksUpdatedPublish"
	var Objek sot_models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirNonEksScheduler = cass_models.BidKurirNonEksScheduler{
		ID:           Objek.ID,
		IdBid:        Objek.IdBid,
		IdKurir:      Objek.IdKurir,
		Urutan:       Objek.Urutan,
		IdPengiriman: Objek.IdPengiriman,
		Status:       Objek.Status,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanNonEksIIpengirimanPickedUpNonEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatePickedUpPengirimanNonEksIIpengirimanPickedUpNonEksUpdatedPublish" // Dinormalisasi namanya
	var Objek sot_models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.Pengiriman = cass_models.Pengiriman{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatPengguna:  Objek.IdAlamatPengguna,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanNonEksIItransaksiPickedUpNonEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdatePickedUpPengirimanNonEksIItransaksiPickedUpNonEksUpdatedPublish"
	var Objek sot_models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	// 1. Mapping ke Cassandra Model (Flat Primitif)
	var ObjekCass cass_models.Transaksi = cass_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
		UpdatedAt:           Objek.UpdatedAt,
		DeletedAt:           Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	// 2. Pipeline Cassandra (SOT Replica)
	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	// 3. Pipeline Cassandra (Historical DB)
	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 4. Mapping & Push ke Search Engine (Meilisearch)
	var ObjekSearchEngine se_models.Transaksi = se_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
	}

	if task_info, err := se_index.TransaksiIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data transaksi ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data transaksi ke search engine dengan info: %s ", task_info.IndexUID)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanNonEksIIbidKurirPengirimanNonEksSchedulerUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateKirimPengirimanNonEksIIbidKurirPengirimanNonEksSchedulerUpdatedPublish"
	var Objek sot_models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirNonEksScheduler = cass_models.BidKurirNonEksScheduler{
		ID:           Objek.ID,
		IdBid:        Objek.IdBid,
		IdKurir:      Objek.IdKurir,
		Urutan:       Objek.Urutan,
		IdPengiriman: Objek.IdPengiriman,
		Status:       Objek.Status,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanNonEksIIpengirimanPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateKirimPengirimanNonEksIIpengirimanPengirimanUpdatedPublish"
	var Objek sot_models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.Pengiriman = cass_models.Pengiriman{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatPengguna:  Objek.IdAlamatPengguna,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanNonEksIIjejakpengirimanPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateKirimPengirimanNonEksIIjejakpengirimanPengirimanUpdatedPublish"
	var Objek sot_models.JejakPengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengiriman = cass_models.JejakPengiriman{
		ID:           Objek.ID,
		IdPengiriman: Objek.IdPengiriman,
		Lokasi:       Objek.Lokasi,
		Keterangan:   Objek.Keterangan,
		Latitude:     Objek.Latitude,
		Longtitude:   Objek.Longtitude,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateInformasiPerjalananPengirimanNonEks(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateInformasiPerjalananPengirimanNonEks"
	var Objek sot_models.JejakPengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengiriman = cass_models.JejakPengiriman{
		ID:           Objek.ID,
		IdPengiriman: Objek.IdPengiriman,
		Lokasi:       Objek.Lokasi,
		Keterangan:   Objek.Keterangan,
		Latitude:     Objek.Latitude,
		Longtitude:   Objek.Longtitude,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteSampaiPengirimanNonEksIIbidKurirNonEksDeletePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteSampaiPengirimanNonEksIIbidKurirNonEksDeletePublish"
	var Objek sot_models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirNonEksScheduler = cass_models.BidKurirNonEksScheduler{
		ID:           Objek.ID,
		IdBid:        Objek.IdBid,
		IdKurir:      Objek.IdKurir,
		Urutan:       Objek.Urutan,
		IdPengiriman: Objek.IdPengiriman,
		Status:       Objek.Status,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan log hapus ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanNonEksIIpengirimanSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateSampaiPengirimanNonEksIIpengirimanSampaiUpdatedPublish"
	var Objek sot_models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.Pengiriman = cass_models.Pengiriman{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatPengguna:  Objek.IdAlamatPengguna,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanNonEksIIbidKurirDataSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateSampaiPengirimanNonEksIIbidKurirDataSampaiUpdatedPublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanNonEksIIjejakPengirimanSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateSampaiPengirimanNonEksIIjejakPengirimanSampaiUpdatedPublish"
	var Objek sot_models.JejakPengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengiriman = cass_models.JejakPengiriman{
		ID:           Objek.ID,
		IdPengiriman: Objek.IdPengiriman,
		Lokasi:       Objek.Lokasi,
		Keterangan:   Objek.Keterangan,
		Latitude:     Objek.Latitude,
		Longtitude:   Objek.Longtitude,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanNonEksIItransaksiSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdateSampaiPengirimanNonEksIItransaksiSampaiUpdatedPublish"
	var Objek sot_models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	// 1. Mapping ke Cassandra Model
	var ObjekCass cass_models.Transaksi = cass_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
		UpdatedAt:           Objek.UpdatedAt,
		DeletedAt:           Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	// 2. Update ke Cassandra Sot Replica
	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	// 3. Catat ke Cassandra Historical DB
	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 4. Sinkronisasi ke Search Engine (Meilisearch)
	var ObjekSearchEngine se_models.Transaksi = se_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
	}

	if task_info, err := se_index.TransaksiIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data transaksi ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data transaksi ke search engine dengan info: %s ", task_info.IndexUID)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateSampaiPengirimanNonEksIIpayOutSellerCreatePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateSampaiPengirimanNonEksIIpayOutSellerCreatePublish"
	var Objek sot_models.PayOutSeller

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.PayOutSeller = cass_models.PayOutSeller{
		ID:               Objek.ID,
		IdSeller:         Objek.IdSeller,
		IdDisbursment:    Objek.IdDisbursment,
		UserId:           Objek.UserId,
		Amount:           Objek.Amount,
		Status:           Objek.Status,
		Reason:           Objek.Reason,
		Timestamp:        Objek.Timestamp,
		BankCode:         Objek.BankCode,
		AccountNumber:    Objek.AccountNumber,
		RecipientName:    Objek.RecipientName,
		SenderBank:       Objek.SenderBank,
		Remark:           Objek.Remark,
		Receipt:          Objek.Receipt,
		TimeServed:       Objek.TimeServed,
		BundleId:         Objek.BundleId,
		CompanyId:        Objek.CompanyId,
		RecipientCity:    Objek.RecipientCity,
		CreatedFrom:      Objek.CreatedFrom,
		Direction:        Objek.Direction,
		Sender:           Objek.Sender,
		Fee:              Objek.Fee,
		BeneficiaryEmail: Objek.BeneficiaryEmail,
		IdempotencyKey:   Objek.IdempotencyKey,
		IsVirtualAccount: Objek.IsVirtualAccount,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateSampaiPengirimanNonEksIIpayOutKurirCreatePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateSampaiPengirimanNonEksIIpayOutKurirCreatePublish"
	var Objek sot_models.PayOutKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.PayOutKurir = cass_models.PayOutKurir{
		ID:               Objek.ID,
		IdKurir:          Objek.IdKurir,
		IdDisbursment:    Objek.IdDisbursment,
		UserId:           Objek.UserId,
		Amount:           Objek.Amount,
		Status:           Objek.Status,
		Reason:           Objek.Reason,
		Timestamp:        Objek.Timestamp,
		BankCode:         Objek.BankCode,
		AccountNumber:    Objek.AccountNumber,
		RecipientName:    Objek.RecipientName,
		SenderBank:       Objek.SenderBank,
		Remark:           Objek.Remark,
		Receipt:          Objek.Receipt,
		TimeServed:       Objek.TimeServed,
		BundleId:         Objek.BundleId,
		CompanyId:        Objek.CompanyId,
		RecipientCity:    Objek.RecipientCity,
		CreatedFrom:      Objek.CreatedFrom,
		Direction:        Objek.Direction,
		Sender:           Objek.Sender,
		Fee:              Objek.Fee,
		BeneficiaryEmail: Objek.BeneficiaryEmail,
		IdempotencyKey:   Objek.IdempotencyKey,
		IsVirtualAccount: Objek.IsVirtualAccount,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreatePickedUpPengirimanEksIIjejakPengirimanEksCreatePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreatePickedUpPengirimanEksIIjejakPengirimanEksCreatePublish"
	var Objek sot_models.JejakPengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengirimanEkspedisi = cass_models.JejakPengirimanEkspedisi{
		ID:                    Objek.ID,
		IdPengirimanEkspedisi: Objek.IdPengirimanEkspedisi,
		Lokasi:                Objek.Lokasi,
		Keterangan:            Objek.Keterangan,
		Latitude:              Objek.Latitude,
		Longitude:             Objek.Longitude,
		CreatedAt:             Objek.CreatedAt,
		UpdatedAt:             Objek.UpdatedAt,
		DeletedAt:             Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanEksIIschedulerEksPickedUpUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatePickedUpPengirimanEksIIschedulerEksPickedUpUpdatedPublish"
	var Objek sot_models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirEksScheduler = cass_models.BidKurirEksScheduler{
		ID:              Objek.ID,
		IdBid:           Objek.IdBid,
		IdKurir:         Objek.IdKurir,
		Urutan:          Objek.Urutan,
		IdPengirimanEks: Objek.IdPengirimanEks,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanEksIIpengirimanEksPickedUpUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatePickedUpPengirimanEksIIpengirimanEksPickedUpUpdatedPublish"
	var Objek sot_models.PengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.PengirimanEkspedisi = cass_models.PengirimanEkspedisi{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatEkspedisi: Objek.IdAlamatEkspedisi,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanEksIItransaksiPickedUpUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdatePickedUpPengirimanEksIItransaksiPickedUpUpdatedPublish"
	var Objek sot_models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	// 1. Mapping ke Cassandra Model
	var ObjekCass cass_models.Transaksi = cass_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
		UpdatedAt:           Objek.UpdatedAt,
		DeletedAt:           Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	// 2. Update ke Cassandra Sot Replica
	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	// 3. Catat ke Cassandra Historical DB
	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 4. Sinkronisasi ke Search Engine (Meilisearch)
	var ObjekSearchEngine se_models.Transaksi = se_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
	}

	if task_info, err := se_index.TransaksiIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data transaksi ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data transaksi ke search engine dengan info: %s ", task_info.IndexUID)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanEksIIschedulerPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateKirimPengirimanEksIIschedulerPengirimanUpdatedPublish"
	var Objek sot_models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirEksScheduler = cass_models.BidKurirEksScheduler{
		ID:              Objek.ID,
		IdBid:           Objek.IdBid,
		IdKurir:         Objek.IdKurir,
		Urutan:          Objek.Urutan,
		IdPengirimanEks: Objek.IdPengirimanEks,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanEksIIpengirimanEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateKirimPengirimanEksIIpengirimanEksUpdatedPublish"
	var Objek sot_models.PengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.PengirimanEkspedisi = cass_models.PengirimanEkspedisi{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatEkspedisi: Objek.IdAlamatEkspedisi,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanEksIIpengirimanPengirimanEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateKirimPengirimanEksIIpengirimanPengirimanEksUpdatedPublish"
	var Objek sot_models.JejakPengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengirimanEkspedisi = cass_models.JejakPengirimanEkspedisi{
		ID:                    Objek.ID,
		IdPengirimanEkspedisi: Objek.IdPengirimanEkspedisi,
		Lokasi:                Objek.Lokasi,
		Keterangan:            Objek.Keterangan,
		Latitude:              Objek.Latitude,
		Longitude:             Objek.Longitude,
		CreatedAt:             Objek.CreatedAt,
		UpdatedAt:             Objek.UpdatedAt,
		DeletedAt:             Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateInformasiPerjalananPengirimanEks(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateInformasiPerjalananPengirimanEks"
	var Objek sot_models.JejakPengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengirimanEkspedisi = cass_models.JejakPengirimanEkspedisi{
		ID:                    Objek.ID,
		IdPengirimanEkspedisi: Objek.IdPengirimanEkspedisi,
		Lokasi:                Objek.Lokasi,
		Keterangan:            Objek.Keterangan,
		Latitude:              Objek.Latitude,
		Longitude:             Objek.Longitude,
		CreatedAt:             Objek.CreatedAt,
		UpdatedAt:             Objek.UpdatedAt,
		DeletedAt:             Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteSampaipengirimanEksIIbidKurirEksDeletePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteSampaipengirimanEksIIbidKurirEksDeletePublish"
	var Objek sot_models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirEksScheduler = cass_models.BidKurirEksScheduler{
		ID:              Objek.ID,
		IdBid:           Objek.IdBid,
		IdKurir:         Objek.IdKurir,
		Urutan:          Objek.Urutan,
		IdPengirimanEks: Objek.IdPengirimanEks,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan log hapus ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIpengirimanSampaiEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateSampaiPengirimanEksIIpengirimanSampaiEksUpdatedPublish"
	var Objek sot_models.PengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.PengirimanEkspedisi = cass_models.PengirimanEkspedisi{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatEkspedisi: Objek.IdAlamatEkspedisi,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIbidKurirDataEksSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateSampaiPengirimanEksIIbidKurirDataEksSampaiUpdatedPublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIjejakPengirimanEksSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateSampaiPengirimanEksIIjejakPengirimanEksSampaiUpdatedPublish"
	var Objek sot_models.JejakPengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengirimanEkspedisi = cass_models.JejakPengirimanEkspedisi{
		ID:                    Objek.ID,
		IdPengirimanEkspedisi: Objek.IdPengirimanEkspedisi,
		Lokasi:                Objek.Lokasi,
		Keterangan:            Objek.Keterangan,
		Latitude:              Objek.Latitude,
		Longitude:             Objek.Longitude,
		CreatedAt:             Objek.CreatedAt,
		UpdatedAt:             Objek.UpdatedAt,
		DeletedAt:             Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIItransaksiSampaiEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdateSampaiPengirimanEksIItransaksiSampaiEksUpdatedPublish"
	var Objek sot_models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	// 1. Mapping ke Cassandra Model
	var ObjekCass cass_models.Transaksi = cass_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
		UpdatedAt:           Objek.UpdatedAt,
		DeletedAt:           Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	// 2. Update ke Cassandra Sot Replica
	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	// 3. Catat ke Cassandra Historical DB
	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 4. Sinkronisasi ke Search Engine (Meilisearch)
	var ObjekSearchEngine se_models.Transaksi = se_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
		UpdatedAt:           Objek.UpdatedAt,
		DeletedAt:           &ObjekCass.Pengguna.DeletedAt,
	}

	if task_info, err := se_index.TransaksiIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data transaksi ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data transaksi ke search engine dengan info: %s ", task_info.IndexUID)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateSampaiPengirimanEksIIpayOutKurirEksCreatePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateSampaiPengirimanEksIIpayOutKurirEksCreatePublish"
	var Objek sot_models.PayOutKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.PayOutKurir = cass_models.PayOutKurir{
		ID:               Objek.ID,
		IdKurir:          Objek.IdKurir,
		IdDisbursment:    Objek.IdDisbursment,
		UserId:           Objek.UserId,
		Amount:           Objek.Amount,
		Status:           Objek.Status,
		Reason:           Objek.Reason,
		Timestamp:        Objek.Timestamp,
		BankCode:         Objek.BankCode,
		AccountNumber:    Objek.AccountNumber,
		RecipientName:    Objek.RecipientName,
		SenderBank:       Objek.SenderBank,
		Remark:           Objek.Remark,
		Receipt:          Objek.Receipt,
		TimeServed:       Objek.TimeServed,
		BundleId:         Objek.BundleId,
		CompanyId:        Objek.CompanyId,
		RecipientCity:    Objek.RecipientCity,
		CreatedFrom:      Objek.CreatedFrom,
		Direction:        Objek.Direction,
		Sender:           Objek.Sender,
		Fee:              Objek.Fee,
		BeneficiaryEmail: Objek.BeneficiaryEmail,
		IdempotencyKey:   Objek.IdempotencyKey,
		IsVirtualAccount: Objek.IsVirtualAccount,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIkurirUpdatedSampaiEksPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services string = "UpdateSampaiPengirimanEksIIkurirUpdatedSampaiEksPublish"
	var Objek sot_models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	// 1. Update ke Cassandra Sot Replica & Historical
	var ObjekCass cass_models.Kurir = cass_models.Kurir{
		ID:            Objek.ID,
		Nama:          Objek.Nama,
		Username:      Objek.Username,
		Email:         Objek.Email,
		Jenis:         Objek.Jenis,
		PasswordHash:  Objek.PasswordHash,
		Deskripsi:     Objek.Deskripsi,
		StatusKurir:   Objek.StatusKurir,
		StatusBid:     Objek.StatusBid,
		VerifiedKurir: Objek.VerifiedKurir,
		Rating:        Objek.Rating,
		TipeKendaraan: Objek.TipeKendaraan,
		CreatedAt:     Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 2. Sinkronisasi ke Search Engine (Meilisearch)
	var ObjekSearchEngine se_models.Kurir = se_models.Kurir{
		ID:            Objek.ID,
		Nama:          Objek.Nama,
		Username:      Objek.Username,
		Email:         Objek.Email,
		Jenis:         Objek.Jenis,
		PasswordHash:  Objek.PasswordHash,
		Deskripsi:     Objek.Deskripsi,
		StatusKurir:   Objek.StatusKurir,
		StatusBid:     Objek.StatusBid,
		VerifiedKurir: Objek.VerifiedKurir,
		Rating:        Objek.Rating,
		TipeKendaraan: Objek.TipeKendaraan,
		CreatedAt:     Objek.CreatedAt,
	}

	if task_info, err := se_index.KurirIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data kurir ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data kurir ke search engine dengan info: %s ", task_info.IndexUID)
	}

	// 3. Update Session Data di Cache Redis
	if err := cache_db_function.UpdateSessionData(ctx, *rds_session, cache_db_session.GetSessionKey(&Objek), Objek); err != nil {
		return fmt.Errorf("gagal memperbarui data kurir di cache %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteNonaktifkanBidKurirIIbidKurirDataDeletePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteNonaktifkanBidKurirIIbidKurirDataDeletePublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan log hapus ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateNonaktifkanBidKurirIIkurirNonaktifkanBidUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services string = "UpdateNonaktifkanBidKurirIIkurirNonaktifkanBidUpdatedPublish"
	var Objek sot_models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	// 1. Update ke Cassandra Sot Replica & Historical
	var ObjekCass cass_models.Kurir = cass_models.Kurir{
		ID:            Objek.ID,
		Nama:          Objek.Nama,
		Username:      Objek.Username,
		Email:         Objek.Email,
		Jenis:         Objek.Jenis,
		PasswordHash:  Objek.PasswordHash,
		Deskripsi:     Objek.Deskripsi,
		StatusKurir:   Objek.StatusKurir,
		StatusBid:     Objek.StatusBid,
		VerifiedKurir: Objek.VerifiedKurir,
		Rating:        Objek.Rating,
		TipeKendaraan: Objek.TipeKendaraan,
		CreatedAt:     Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 2. Sinkronisasi ke Search Engine (Meilisearch)
	var ObjekSearchEngine se_models.Kurir = se_models.Kurir{
		ID:            Objek.ID,
		Nama:          Objek.Nama,
		Username:      Objek.Username,
		Email:         Objek.Email,
		Jenis:         Objek.Jenis,
		PasswordHash:  Objek.PasswordHash,
		Deskripsi:     Objek.Deskripsi,
		StatusKurir:   Objek.StatusKurir,
		StatusBid:     Objek.StatusBid,
		VerifiedKurir: Objek.VerifiedKurir,
		Rating:        Objek.Rating,
		TipeKendaraan: Objek.TipeKendaraan,
		CreatedAt:     Objek.CreatedAt,
	}

	if task_info, err := se_index.KurirIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data kurir ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data kurir ke search engine dengan info: %s ", task_info.IndexUID)
	}

	// 3. Update Session Data di Cache Redis
	if err := cache_db_function.UpdateSessionData(ctx, *rds_session, cache_db_session.GetSessionKey(&Objek), Objek); err != nil {
		return fmt.Errorf("gagal memperbarui data kurir di cache %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
