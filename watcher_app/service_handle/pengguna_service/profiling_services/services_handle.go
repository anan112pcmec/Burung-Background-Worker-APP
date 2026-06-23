package profiling_pengguna_handle

import (
	"context"
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

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

func UpdateUbahPersonalProfilingPengguna(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services string = "UpdateUbahPersonalProfilingPengguna"
	var Objek sot_models.Pengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Pengguna = cass_models.Pengguna{
		ID:             Objek.ID,
		Username:       Objek.Username,
		Nama:           Objek.Nama,
		Email:          Objek.Email,
		PasswordHash:   Objek.PasswordHash,
		PinHash:        Objek.PinHash,
		StatusPengguna: Objek.StatusPengguna,
		CreatedAt:      Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Pengguna = se_models.Pengguna{
		ID:             Objek.ID,
		Username:       Objek.Username,
		Nama:           Objek.Nama,
		Email:          Objek.Email,
		PasswordHash:   Objek.PasswordHash,
		PinHash:        Objek.PinHash,
		StatusPengguna: Objek.StatusPengguna,
		CreatedAt:      Objek.CreatedAt,
	}

	if task_info, err := se_index.PenggunaIndex.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Println("berhasil memasukan data ke dalam search engine dengan UID %s", task_info.IndexUID)
	}

	if err := cache_db_function.UpdateSessionData[sot_models.Pengguna](ctx, *rds_session, cache_db_session.GetSessionKey[*sot_models.Pengguna](&Objek), Objek); err != nil {
		return fmt.Errorf("gagal memasukan mengubah data sesi pengguna")
	}

	return nil
}
