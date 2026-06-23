package alamat_pengguna_handle

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

// data body yang diinput merupakan model relasi alamat pengguna
func CreateAlamatPub(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateAlamatPub"
	var Objek sot_models.AlamatPengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.AlamatPengguna = cass_models.AlamatPengguna{
		ID:              Objek.ID,
		IDPengguna:      Objek.IDPengguna,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodePos,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica sync %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historica db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.AlamatPengguna = se_models.AlamatPengguna{
		ID:              Objek.ID,
		IDPengguna:      Objek.IDPengguna,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodePos,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.AlamatPenggunaIndex.AddDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Println("Berhasil memasukan data ke dalam search engine dengan antrean UID %s", task_info.IndexUID)
	}
	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAlamatPub(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdateAlamatPub"
	var Objek sot_models.AlamatPengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.AlamatPengguna = cass_models.AlamatPengguna{
		ID:              Objek.ID,
		IDPengguna:      Objek.IDPengguna,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodePos,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica sync %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historica db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.AlamatPengguna = se_models.AlamatPengguna{
		ID:              Objek.ID,
		IDPengguna:      Objek.IDPengguna,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodePos,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.AlamatPenggunaIndex.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Println("Berhasil memasukan data ke dalam search engine dengan antrean UID %s", task_info.IndexUID)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteAlamatPub(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "DeleteAlamatPub"
	var Objek sot_models.AlamatPengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.AlamatPengguna = cass_models.AlamatPengguna{
		ID:              Objek.ID,
		IDPengguna:      Objek.IDPengguna,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodePos,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica sync %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historica db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.AlamatPengguna = se_models.AlamatPengguna{
		ID:              Objek.ID,
		IDPengguna:      Objek.IDPengguna,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodePos:         Objek.KodePos,
		KodeNegara:      Objek.KodeNegara,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	idStr := strconv.FormatInt(ObjekSearchEngine.ID, 10)

	if task_info, err := se_index.AlamatPenggunaIndex.DeleteDocumentWithContext(ctx, idStr, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Println("Berhasil memasukan data ke dalam search engine dengan antrean UID %s", task_info.IndexUID)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
