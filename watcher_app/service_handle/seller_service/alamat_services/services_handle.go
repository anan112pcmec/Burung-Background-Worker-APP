package alamat_seller_handle

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
	return nil
}
