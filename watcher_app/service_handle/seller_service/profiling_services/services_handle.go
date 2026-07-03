package profiling_seller_handle

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/cache"
	cache_db_function "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cache_db/function"
	cache_db_session "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cache_db/session"
	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	se_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/search_engine/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	notification_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/models"
	notification_request "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/request"
	notification_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/seeders"
)

func UpdateUpdatePersonalSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services string = "UpdateUpdatePersonalSeller"
	var Objek sot_models.Seller
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Seller = cass_models.Seller{
		ID:               int(Objek.ID),
		Username:         Objek.Username,
		Nama:             Objek.Nama,
		Email:            Objek.Email,
		Jenis:            Objek.Jenis,
		SellerDedication: Objek.SellerDedication,
		JamOperasional:   Objek.JamOperasional,
		Punchline:        Objek.Punchline,
		Password:         Objek.Password,
		Deskripsi:        Objek.Deskripsi,
		StatusSeller:     Objek.StatusSeller,
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

	var ObjekSearchEngine se_models.Seller = se_models.Seller{
		ID:               Objek.ID,
		Username:         Objek.Username,
		Nama:             Objek.Nama,
		Email:            Objek.Email,
		Jenis:            Objek.Jenis,
		SellerDedication: Objek.SellerDedication,
		JamOperasional:   Objek.JamOperasional,
		Punchline:        Objek.Punchline,
		Password:         Objek.Password,
		Deskripsi:        Objek.Deskripsi,
		StatusSeller:     Objek.StatusSeller,
		CreatedAt:        Objek.CreatedAt,
	}

	if task_info, err := se_index.SellerIndex.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke dalam search engine dengan uid: %s\n", task_info.IndexUID)
	}

	if err := cache_db_function.UpdateSessionData[sot_models.Seller](ctx, *rds_session, cache_db_session.GetSessionKey[*sot_models.Seller](&Objek), Objek); err != nil {
		return fmt.Errorf("gagal mengupdate session data %s dalam %s", err, handle_services)
	}

	// ðŸ”” SISTEM NOTIFIKASI
	if Objek.ID != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.ID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”’ Data Personal Diperbarui",
			Pesan:     "Informasi kredensial dan data personal akun toko Anda telah berhasil diperbarui.",
			Pop:       1,
			Archive:   false,
			Inbox:     true,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.ID, "username": Objek.Username},
				Special:  map[string]interface{}{"click_action": "SELLER_PERSONAL_DATA_UPDATED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func UpdateUpdateInfoGeneralPublic(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services string = "UpdateUpdateInfoGeneralPublic"
	var Objek sot_models.Seller
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Seller = cass_models.Seller{
		ID:               int(Objek.ID),
		Username:         Objek.Username,
		Nama:             Objek.Nama,
		Email:            Objek.Email,
		Jenis:            Objek.Jenis,
		SellerDedication: Objek.SellerDedication,
		JamOperasional:   Objek.JamOperasional,
		Punchline:        Objek.Punchline,
		Password:         Objek.Password,
		Deskripsi:        Objek.Deskripsi,
		StatusSeller:     Objek.StatusSeller,
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

	var ObjekSearchEngine se_models.Seller = se_models.Seller{
		ID:               Objek.ID,
		Username:         Objek.Username,
		Nama:             Objek.Nama,
		Email:            Objek.Email,
		Jenis:            Objek.Jenis,
		SellerDedication: Objek.SellerDedication,
		JamOperasional:   Objek.JamOperasional,
		Punchline:        Objek.Punchline,
		Password:         Objek.Password,
		Deskripsi:        Objek.Deskripsi,
		StatusSeller:     Objek.StatusSeller,
		CreatedAt:        Objek.CreatedAt,
	}

	if task_info, err := se_index.SellerIndex.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke dalam search engine dengan uid: %s\n", task_info.IndexUID)
	}

	if err := cache_db_function.UpdateSessionData[sot_models.Seller](ctx, *rds_session, cache_db_session.GetSessionKey[*sot_models.Seller](&Objek), Objek); err != nil {
		return fmt.Errorf("gagal mengupdate session data %s dalam %s", err, handle_services)
	}

	// ðŸ”” SISTEM NOTIFIKASI
	if Objek.ID != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.ID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸª Profil Toko Diperbarui",
			Pesan:     fmt.Sprintf("Informasi umum toko Anda yang terlihat oleh publik resmi diubah."),
			Pop:       1,
			Archive:   false,
			Inbox:     true,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.ID, "nama_toko": Objek.Nama},
				Special:  map[string]interface{}{"click_action": "SELLER_PUBLIC_INFO_UPDATED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}
