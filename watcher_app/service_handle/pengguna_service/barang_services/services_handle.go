package barang_pengguna_handle

import (
	"context"
	"fmt"

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

func CreateLikesBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateLikesBarang"
	var Objek sot_models.BarangDisukai
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BarangDisukai = cass_models.BarangDisukai{
		ID:            Objek.ID,
		IdPengguna:    Objek.IdPengguna,
		IdBarangInduk: Objek.IdBarangInduk,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasuakan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.IdPengguna)
	return nil
}

func DeleteUnlikesBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteUnlikesBarang"
	var Objek sot_models.BarangDisukai
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BarangDisukai = cass_models.BarangDisukai{
		ID:            Objek.ID,
		IdPengguna:    Objek.IdPengguna,
		IdBarangInduk: Objek.IdBarangInduk,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedDataLikesBarang map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedDataLikesBarang)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedDataLikesBarang); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.IdPengguna)
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
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateEditKomentarBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatedEditKomentarBarang"
	var Objek cass_models.Komentar
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
		DeletedAt:     Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
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

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
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
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateMentionChildKomentar(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateMentionChildKomentar"
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
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
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

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
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

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}
	return nil
}

func CreateTambahKeranjangBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahKeranjangBarang"
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

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}
	return nil
}

func UpdateEditKeranjangBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditKeranjangBarang"
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

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}
	return nil
}

func DeleteHapusKeranjangBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusKeranjangBarang"
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

func CreateBerikanReviewBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateBerikanReviewBarang"
	var Objek sot_models.Review
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	var ObjekCass cass_models.Review = cass_models.Review{
		ID:            Objek.ID,
		IdPengguna:    Objek.IdPengguna,
		IdBarangInduk: Objek.IdBarangInduk,
		Rating:        Objek.Rating,
		Ulasan:        Objek.Ulasan,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}
	return nil
}

func UpdateBerikanReviewBarangIIUpdateTransaksi(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdateberikanReviewBarangiiUpdateTransaksi"
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

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
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

	if task_info, err := se_index.TransaksiIndex.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Println("Berhasil memasukan data ke dalam search engine dengan antrean UID %s", task_info.IndexUID)
	}

	return nil
}

func CreateLikeReviewBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdateLikeReviewBarang"
	var Objek sot_models.ReviewLike
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.ReviewLike = cass_models.ReviewLike{
		ID:         Objek.ID,
		IdPengguna: Objek.IdPengguna,
		IdReview:   Objek.IdReview,
		CreatedAt:  Objek.CreatedAt,
		UpdatedAt:  Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}
	return nil
}

func DeleteUnlikeReviewBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "DeleteUnlikeReviewBarang"
	var Objek sot_models.ReviewLike
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.ReviewLike = cass_models.ReviewLike{
		ID:         Objek.ID,
		IdPengguna: Objek.IdPengguna,
		IdReview:   Objek.IdReview,
		CreatedAt:  Objek.CreatedAt,
		UpdatedAt:  Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}
	return nil
}
