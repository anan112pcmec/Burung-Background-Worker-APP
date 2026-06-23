package social_media_pengguna_handle

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

func CreateEngageTautkanSocialMediaPengguna(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateEngageTautkanSocialMediaPengguna"
	var Objek sot_models.EntitySocialMedia
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.EntitySocialMedia = cass_models.EntitySocialMedia{
		ID:         Objek.ID,
		EntityId:   Objek.EntityId,
		Whatsapp:   Objek.Whatsapp,
		Facebook:   Objek.Facebook,
		TikTok:     Objek.TikTok,
		Instagram:  Objek.Instagram,
		Metadata:   Objek.Metadata,
		EntityType: Objek.EntityType,
		CreatedAt:  Objek.CreatedAt,
		UpdatedAt:  Objek.UpdatedAt,
		DeletedAt:  Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasuakan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateEngageTautkanSocialMediaPengguna(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEngageTautkanSocialMediaPengguna"
	var Objek sot_models.EntitySocialMedia
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.EntitySocialMedia = cass_models.EntitySocialMedia{
		ID:         Objek.ID,
		EntityId:   Objek.EntityId,
		Whatsapp:   Objek.Whatsapp,
		Facebook:   Objek.Facebook,
		TikTok:     Objek.TikTok,
		Instagram:  Objek.Instagram,
		Metadata:   Objek.Metadata,
		EntityType: Objek.EntityType,
		CreatedAt:  Objek.CreatedAt,
		UpdatedAt:  Objek.UpdatedAt,
		DeletedAt:  Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasuakan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateEngageHapusSocialMedia(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEngageHapusSocialMedia"
	var Objek sot_models.EntitySocialMedia
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.EntitySocialMedia = cass_models.EntitySocialMedia{
		ID:         Objek.ID,
		EntityId:   Objek.EntityId,
		Whatsapp:   Objek.Whatsapp,
		Facebook:   Objek.Facebook,
		TikTok:     Objek.TikTok,
		Instagram:  Objek.Instagram,
		Metadata:   Objek.Metadata,
		EntityType: Objek.EntityType,
		CreatedAt:  Objek.CreatedAt,
		UpdatedAt:  Objek.UpdatedAt,
		DeletedAt:  Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasuakan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateFollowSeller(Data mb_cud_serializer.ParsedDataMessage, Read *gorm.DB, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateFollowSeller"
	var Objek sot_models.Follower
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Follower = cass_models.Follower{
		ID:         Objek.ID,
		IdFollower: Objek.IdFollower,
		IdFollowed: Objek.IdFollowed,
		CreatedAt:  Objek.CreatedAt,
		UpdatedAt:  Objek.UpdatedAt,
		DeletedAt:  Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasuakan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Follower = se_models.Follower{
		ID:         Objek.ID,
		IdFollower: Objek.IdFollower,
		IdFollowed: Objek.IdFollowed,
		CreatedAt:  Objek.CreatedAt,
		UpdatedAt:  Objek.UpdatedAt,
		DeletedAt:  Objek.DeletedAt,
	}

	if task_info, err := se_index.FollowerSeller.AddDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke dalam search engine dengan UID %s", task_info.IndexUID)
	}

	return nil
}

func DeleteUnfollowSeller(Data mb_cud_serializer.ParsedDataMessage, Read *gorm.DB, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "DeleteUnfollowSeller"
	var Objek sot_models.Follower
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Follower = cass_models.Follower{
		ID:         Objek.ID,
		IdFollower: Objek.IdFollower,
		IdFollowed: Objek.IdFollowed,
		CreatedAt:  Objek.CreatedAt,
		UpdatedAt:  Objek.UpdatedAt,
		DeletedAt:  Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasuakan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Follower = se_models.Follower{
		ID:         Objek.ID,
		IdFollower: Objek.IdFollower,
		IdFollowed: Objek.IdFollowed,
		CreatedAt:  Objek.CreatedAt,
		UpdatedAt:  Objek.UpdatedAt,
		DeletedAt:  Objek.DeletedAt,
	}

	idStr := strconv.FormatInt(ObjekSearchEngine.ID, 10)

	if task_info, err := se_index.FollowerSeller.DeleteDocumentWithContext(ctx, idStr, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memproses data ke dalam search engine dengan UID %s", task_info.IndexUID)
	}

	return nil
}
