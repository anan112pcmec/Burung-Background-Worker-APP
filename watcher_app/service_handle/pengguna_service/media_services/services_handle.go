package media_pengguna_handle

import (
	"context"
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"gorm.io/gorm"

	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func UpdateUbahFotoProfilPengguna(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahFotoProfilPengguna"
	var Objek sot_models.MediaPenggunaProfilFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaPenggunaProfilFoto = cass_models.MediaPenggunaProfilFoto{
		ID:         Objek.ID,
		IdPengguna: Objek.IdPengguna,
		Key:        Objek.Key,
		Format:     Objek.Format,
		CreatedAt:  Objek.CreatedAt,
		UpdatedAt:  Objek.UpdatedAt,
		DeletedAt:  Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
			return fmt.Errorf("gagal memasukan data ke dalam sot replica sync %s dalam %s", err, handle_services)
		}

	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historica db %s dalam %s", err, handle_services)
	}

	return nil
}

func DeleteHapusFotoProfilPengguna(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteUbahFotoProfilPengguna"
	var Objek sot_models.MediaPenggunaProfilFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaPenggunaProfilFoto = cass_models.MediaPenggunaProfilFoto{
		ID:         Objek.ID,
		IdPengguna: Objek.IdPengguna,
		Key:        Objek.Key,
		Format:     Objek.Format,
		CreatedAt:  Objek.CreatedAt,
		UpdatedAt:  Objek.UpdatedAt,
		DeletedAt:  Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica sync %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historica db %s dalam %s", err, handle_services)
	}
	return nil
}

func CreateTambahMediaReviewFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaReviewFoto"
	var Objek sot_models.MediaReviewFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaReviewFoto = cass_models.MediaReviewFoto{
		ID:       Objek.ID,
		IdReview: Objek.IdReview,
		Key:      Objek.Key,
		Format:   Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica sync %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historica db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateTambahMediaReviewVideo(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaReviewVideo"
	var Objek sot_models.MediaReviewVideo
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.MediaReviewVideo = cass_models.MediaReviewVideo{
		ID:       Objek.ID,
		IdReview: Objek.IdReview,
		Key:      Objek.Key,
		Format:   Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica sync %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historica db %s dalam %s", err, handle_services)
	}

	return nil
}
