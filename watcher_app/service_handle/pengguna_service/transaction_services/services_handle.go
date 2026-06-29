package transaction_pengguna_handle

import (
	"context"
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"

	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	se_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/search_engine/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func DeleteCheckoutBarangUser(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteCheckoutBarangUser"
	var Objek sot_models.Keranjang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	var ObjekCass cass_models.Keranjang = cass_models.Keranjang{
		ID:            Objek.ID,
		IdPengguna:    Objek.IdPengguna,
		IdSeller:      Objek.IdSeller,
		IdBarangInduk: Objek.IdBarangInduk,
		IdKategori:    Objek.IdKategori,
		Jumlah:        Objek.Jumlah,
		Status:        Objek.Status,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateLockTransaksiVa(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateLockTransaksiVa"
	var Objek sot_models.Transaksi
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

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

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal menghapus data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

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
		DeletedAt:           &Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.TransaksiIndex.AddDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("Berhasil memasukan data ke dalam search engine dengan antrean UID %s", task_info.IndexUID)
	}

	return nil
}

func CreateLockTransaksiWallet(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateLockTransaksiWallet"
	var Objek sot_models.Transaksi
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

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

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal menghapus data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

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
		DeletedAt:           &Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.TransaksiIndex.AddDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("Berhasil memasukan data ke dalam search engine dengan antrean UID %s", task_info.IndexUID)
	}

	return nil
}

func CreateLockTransaksiGerai(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateLockTransaksiGerai"
	var Objek sot_models.Transaksi
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

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

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal menghapus data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

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
		DeletedAt:           &Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.TransaksiIndex.AddDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("Berhasil memasukan data ke dalam search engine dengan antrean UID %s", task_info.IndexUID)
	}

	return nil
}
