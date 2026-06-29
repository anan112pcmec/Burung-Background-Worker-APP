package barang_seller_handle

import (
	"context"
	"fmt"
	"strconv"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/meilisearch/meilisearch-go"

	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	se_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/search_engine/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"

)

func CreateMasukanBarangInduk(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateMasukanBarangInduk"
	var Objek sot_models.BarangInduk
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BarangInduk = cass_models.BarangInduk{
		ID:               Objek.ID,
		SellerID:         Objek.SellerID,
		IdDiskon:         Objek.IdDiskon,
		NamaBarang:       Objek.NamaBarang,
		JenisBarang:      Objek.JenisBarang,
		Deskripsi:        Objek.Deskripsi,
		OriginalKategori: Objek.OriginalKategori,
		HargaKategoris:   Objek.HargaKategoris,
		CreatedAt:        Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.BarangInduk = se_models.BarangInduk{
		ID:               Objek.ID,
		SellerID:         Objek.SellerID,
		IdDiskon:         Objek.IdDiskon,
		NamaBarang:       Objek.NamaBarang,
		JenisBarang:      Objek.JenisBarang,
		Deskripsi:        Objek.Deskripsi,
		OriginalKategori: Objek.OriginalKategori,
		HargaKategoris:   Objek.HargaKategoris,
		CreatedAt:        Objek.CreatedAt,
	}

	if task_info, err := se_index.BarangIndukIndex.AddDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke dalam search engine dengan uid task %s", task_info.IndexUID)
	}
	return nil
}

func UpdateEditBarangInduk(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdateEditBarangInduk"
	var Objek sot_models.BarangInduk
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BarangInduk = cass_models.BarangInduk{
		ID:               Objek.ID,
		SellerID:         Objek.SellerID,
		IdDiskon:         Objek.IdDiskon,
		NamaBarang:       Objek.NamaBarang,
		JenisBarang:      Objek.JenisBarang,
		Deskripsi:        Objek.Deskripsi,
		OriginalKategori: Objek.OriginalKategori,
		HargaKategoris:   Objek.HargaKategoris,
		CreatedAt:        Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.BarangInduk = se_models.BarangInduk{
		ID:               Objek.ID,
		SellerID:         Objek.SellerID,
		IdDiskon:         Objek.IdDiskon,
		NamaBarang:       Objek.NamaBarang,
		JenisBarang:      Objek.JenisBarang,
		Deskripsi:        Objek.Deskripsi,
		OriginalKategori: Objek.OriginalKategori,
		HargaKategoris:   Objek.HargaKategoris,
		CreatedAt:        Objek.CreatedAt,
	}

	if task_info, err := se_index.BarangIndukIndex.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke dalam search engine dengan uid task %s", task_info.IndexUID)
	}
	return nil
}

func DeleteHapusBarangInduk(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "DeleteHapusBarangInduk"
	var Objek sot_models.BarangInduk
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BarangInduk = cass_models.BarangInduk{
		ID:               Objek.ID,
		SellerID:         Objek.SellerID,
		IdDiskon:         Objek.IdDiskon,
		NamaBarang:       Objek.NamaBarang,
		JenisBarang:      Objek.JenisBarang,
		Deskripsi:        Objek.Deskripsi,
		OriginalKategori: Objek.OriginalKategori,
		HargaKategoris:   Objek.HargaKategoris,
		CreatedAt:        Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID)); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	idStr := strconv.FormatInt(int64(Objek.ID), 10)

	if task_info, err := se_index.BarangIndukIndex.DeleteDocumentWithContext(ctx, idStr, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke dalam search engine dengan uid task %s", task_info.IndexUID)
	}
	return nil
}

func CreateMasukanKategoriBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateMasukanKategoriBarang"
	var Objek sot_models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KategoriBarang = cass_models.KategoriBarang{
		ID:             Objek.ID,
		SellerID:       int32(Objek.SellerID),
		IdBarangInduk:  int32(Objek.IdBarangInduk),
		IDAlamat:       Objek.IDAlamat,
		IDRekening:     Objek.IDRekening,
		Nama:           Objek.Nama,
		Deskripsi:      Objek.Deskripsi,
		Warna:          Objek.Warna,
		Stok:           int32(Objek.Stok),
		Harga:          int32(Objek.Harga),
		PotonganDiskon: int32(Objek.PotonganDiskon),
		BeratGram:      Objek.BeratGram,
		DimensiPanjang: Objek.DimensiPanjang,
		DimensiLebar:   Objek.DimensiLebar,
		Sku:            Objek.Sku,
		IsOriginal:     Objek.IsOriginal,
		CreatedAt:      Objek.CreatedAt,
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

func UpdateEditKategoriBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditKategoriBarang"
	var Objek sot_models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KategoriBarang = cass_models.KategoriBarang{
		ID:             Objek.ID,
		SellerID:       int32(Objek.SellerID),
		IdBarangInduk:  int32(Objek.IdBarangInduk),
		IDAlamat:       Objek.IDAlamat,
		IDRekening:     Objek.IDRekening,
		Nama:           Objek.Nama,
		Deskripsi:      Objek.Deskripsi,
		Warna:          Objek.Warna,
		Stok:           int32(Objek.Stok),
		Harga:          int32(Objek.Harga),
		PotonganDiskon: int32(Objek.PotonganDiskon),
		BeratGram:      Objek.BeratGram,
		DimensiPanjang: Objek.DimensiPanjang,
		DimensiLebar:   Objek.DimensiLebar,
		Sku:            Objek.Sku,
		IsOriginal:     Objek.IsOriginal,
		CreatedAt:      Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateUbahHargaKategoriBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahHargaKategoriBarang"
	var Objek sot_models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KategoriBarang = cass_models.KategoriBarang{
		ID:             Objek.ID,
		SellerID:       int32(Objek.SellerID),
		IdBarangInduk:  int32(Objek.IdBarangInduk),
		IDAlamat:       Objek.IDAlamat,
		IDRekening:     Objek.IDRekening,
		Nama:           Objek.Nama,
		Deskripsi:      Objek.Deskripsi,
		Warna:          Objek.Warna,
		Stok:           int32(Objek.Stok),
		Harga:          int32(Objek.Harga),
		PotonganDiskon: int32(Objek.PotonganDiskon),
		BeratGram:      Objek.BeratGram,
		DimensiPanjang: Objek.DimensiPanjang,
		DimensiLebar:   Objek.DimensiLebar,
		Sku:            Objek.Sku,
		IsOriginal:     Objek.IsOriginal,
		CreatedAt:      Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusBarangKategori(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusBarangKategori"
	var Objek sot_models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KategoriBarang = cass_models.KategoriBarang{
		ID:             Objek.ID,
		SellerID:       int32(Objek.SellerID),
		IdBarangInduk:  int32(Objek.IdBarangInduk),
		IDAlamat:       Objek.IDAlamat,
		IDRekening:     Objek.IDRekening,
		Nama:           Objek.Nama,
		Deskripsi:      Objek.Deskripsi,
		Warna:          Objek.Warna,
		Stok:           int32(Objek.Stok),
		Harga:          int32(Objek.Harga),
		PotonganDiskon: int32(Objek.PotonganDiskon),
		BeratGram:      Objek.BeratGram,
		DimensiPanjang: Objek.DimensiPanjang,
		DimensiLebar:   Objek.DimensiLebar,
		Sku:            Objek.Sku,
		IsOriginal:     Objek.IsOriginal,
		CreatedAt:      Objek.CreatedAt,
	}
	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID)); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateDownStokBarangInduk(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditKategoriBarang"
	var Objek sot_models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KategoriBarang = cass_models.KategoriBarang{
		ID:             Objek.ID,
		SellerID:       int32(Objek.SellerID),
		IdBarangInduk:  int32(Objek.IdBarangInduk),
		IDAlamat:       Objek.IDAlamat,
		IDRekening:     Objek.IDRekening,
		Nama:           Objek.Nama,
		Deskripsi:      Objek.Deskripsi,
		Warna:          Objek.Warna,
		Stok:           int32(Objek.Stok),
		Harga:          int32(Objek.Harga),
		PotonganDiskon: int32(Objek.PotonganDiskon),
		BeratGram:      Objek.BeratGram,
		DimensiPanjang: Objek.DimensiPanjang,
		DimensiLebar:   Objek.DimensiLebar,
		Sku:            Objek.Sku,
		IsOriginal:     Objek.IsOriginal,
		CreatedAt:      Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateDownKategoriBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditKategoriBarang"
	var Objek sot_models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KategoriBarang = cass_models.KategoriBarang{
		ID:             Objek.ID,
		SellerID:       int32(Objek.SellerID),
		IdBarangInduk:  int32(Objek.IdBarangInduk),
		IDAlamat:       Objek.IDAlamat,
		IDRekening:     Objek.IDRekening,
		Nama:           Objek.Nama,
		Deskripsi:      Objek.Deskripsi,
		Warna:          Objek.Warna,
		Stok:           int32(Objek.Stok),
		Harga:          int32(Objek.Harga),
		PotonganDiskon: int32(Objek.PotonganDiskon),
		BeratGram:      Objek.BeratGram,
		DimensiPanjang: Objek.DimensiPanjang,
		DimensiLebar:   Objek.DimensiLebar,
		Sku:            Objek.Sku,
		IsOriginal:     Objek.IsOriginal,
		CreatedAt:      Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateMasukanKomentarBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateMasukanKomentarBarang"
	var Objek sot_models.Komentar
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Komentar = cass_models.Komentar{
		ID:            Objek.ID,
		IdBarangInduk: Objek.IdBarangInduk,
		IdEntity:      Objek.IdEntity,
		JenisEntity:   Objek.JenisEntity,
		Komentar:      Objek.Komentar,
		IsSeller:      Objek.IsSeller,
		Dibalas:       Objek.Dibalas,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
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

func UpdateEditKomentarBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditKomentarBarang"
	var Objek sot_models.Komentar
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Komentar = cass_models.Komentar{
		ID:            Objek.ID,
		IdBarangInduk: Objek.IdBarangInduk,
		IdEntity:      Objek.IdEntity,
		JenisEntity:   Objek.JenisEntity,
		Komentar:      Objek.Komentar,
		IsSeller:      Objek.IsSeller,
		Dibalas:       Objek.Dibalas,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusKomentarBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusKomentarBarang"
	var Objek sot_models.Komentar
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Komentar = cass_models.Komentar{
		ID:            Objek.ID,
		IdBarangInduk: Objek.IdBarangInduk,
		IdEntity:      Objek.IdEntity,
		JenisEntity:   Objek.JenisEntity,
		Komentar:      Objek.Komentar,
		IsSeller:      Objek.IsSeller,
		Dibalas:       Objek.Dibalas,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID)); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}
func CreateMasukanChildKomentar(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateMasukanChildKomentar"
	var Objek sot_models.KomentarChild
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KomentarChild = cass_models.KomentarChild{
		ID:          Objek.ID,
		IdKomentar:  Objek.IdKomentar,
		IdEntity:    Objek.IdEntity,
		JenisEntity: Objek.JenisEntity,
		IsiKomentar: Objek.IsiKomentar,
		IsSeller:    Objek.IsSeller,
		Mention:     Objek.Mention,
		CreatedAt:   Objek.CreatedAt,
		UpdatedAt:   Objek.UpdatedAt,
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

func UpdateEditChildKomentar(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditChildKomentar"
	var Objek sot_models.KomentarChild
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KomentarChild = cass_models.KomentarChild{
		ID:          Objek.ID,
		IdKomentar:  Objek.IdKomentar,
		IdEntity:    Objek.IdEntity,
		JenisEntity: Objek.JenisEntity,
		IsiKomentar: Objek.IsiKomentar,
		IsSeller:    Objek.IsSeller,
		Mention:     Objek.Mention,
		CreatedAt:   Objek.CreatedAt,
		UpdatedAt:   Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusChildKomentar(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusChildKomentar"
	var Objek sot_models.KomentarChild
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KomentarChild = cass_models.KomentarChild{
		ID:          Objek.ID,
		IdKomentar:  Objek.IdKomentar,
		IdEntity:    Objek.IdEntity,
		JenisEntity: Objek.JenisEntity,
		IsiKomentar: Objek.IsiKomentar,
		IsSeller:    Objek.IsSeller,
		Mention:     Objek.Mention,
		CreatedAt:   Objek.CreatedAt,
		UpdatedAt:   Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID)); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	// Sesuai prinsip append-only, penghapusan tetap mencatat state baru ke historical db
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}
