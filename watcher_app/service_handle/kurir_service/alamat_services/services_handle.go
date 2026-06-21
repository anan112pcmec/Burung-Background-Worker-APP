package alamat_kurir_handle

import (
	"context"
	"fmt"
	"strconv"

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

func CreateMasukanAlamatKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	var Objek sot_models.AlamatKurir
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.AlamatKurir = cass_models.AlamatKurir{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodeNegara:      Objek.KodeNegara,
		KodePos:         Objek.KodePos,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	parsedData := ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return err
	}

	pencatatan := historical_format.Sekarang()
	parsedData["tahun_update"] = pencatatan.TahunUpdate
	parsedData["bulan_update"] = pencatatan.BulanUpdate
	parsedData["event_time"] = pencatatan.EventTime

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical()); err != nil {
		return err
	}

	var ObjekSearchEngine se_models.AlamatKurir = se_models.AlamatKurir{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodeNegara:      Objek.KodeNegara,
		KodePos:         Objek.KodePos,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	if task_info, err := se_index.AlamatKurir.AddDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr(strconv.FormatInt(ObjekSearchEngine.ID, 0)),
	}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("task dengan id:" + task_info.IndexUID + "diproses")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatedEditAlamatKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	var Objek sot_models.AlamatKurir
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.AlamatKurir = cass_models.AlamatKurir{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodeNegara:      Objek.KodeNegara,
		KodePos:         Objek.KodePos,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	parsedData := ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_historcal, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return err
	}

	pencatatan := historical_format.Sekarang()
	parsedData["tahun_update"] = pencatatan.TahunUpdate
	parsedData["bulan_update"] = pencatatan.BulanUpdate
	parsedData["event_time"] = pencatatan.EventTime

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return err
	}

	var ObjekSearchEngine se_models.AlamatKurir = se_models.AlamatKurir{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodeNegara:      Objek.KodeNegara,
		KodePos:         Objek.KodePos,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	if task_info, err := se_index.AlamatKurir.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr(strconv.FormatInt(ObjekSearchEngine.ID, 0)),
	}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("task dengan id:" + task_info.IndexUID + "diproses")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusAlamatKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	var Objek sot_models.AlamatKurir
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.AlamatKurir = cass_models.AlamatKurir{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodeNegara:      Objek.KodeNegara,
		KodePos:         Objek.KodePos,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	if err := cass_cud.DeleteData(ctx, cass_historcal, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return err
	}

	parsedData := ObjekCass.ParseToCUDType()
	pencatatan := historical_format.Sekarang()
	parsedData["tahun_update"] = pencatatan.TahunUpdate
	parsedData["bulan_update"] = pencatatan.BulanUpdate
	parsedData["event_time"] = pencatatan.EventTime

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return err
	}

	var ObjekSearchEngine se_models.AlamatKurir = se_models.AlamatKurir{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		PanggilanAlamat: Objek.PanggilanAlamat,
		NomorTelephone:  Objek.NomorTelephone,
		NamaAlamat:      Objek.NamaAlamat,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		KodeNegara:      Objek.KodeNegara,
		KodePos:         Objek.KodePos,
		Deskripsi:       Objek.Deskripsi,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt,
	}

	if task_info, err := se_index.AlamatKurir.DeleteDocumentWithContext(ctx, strconv.FormatInt(ObjekSearchEngine.ID, 0), nil); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("task dengan id:" + task_info.IndexUID + "diproses")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
