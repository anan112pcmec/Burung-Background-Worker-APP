package profiling_kurir_handle

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

func UpdatePersonalProfilingKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services string = "UpdatePersonalProfilingKurir"
	var Objek sot_models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.Kurir = cass_models.Kurir{
		ID:            Objek.ID,
		Nama:          Objek.Nama,
		Username:      Objek.Username,
		Email:         Objek.Email,
		Jenis:         Objek.Jenis,
		PasswordHash:  Objek.PasswordHash,
		Deskripsi:     Objek.Deskripsi,
		StatusKurir:   Objek.StatusKurir,
		StatusBid:     Objek.StatusBid,
		VerifiedKurir: Objek.VerifiedKurir,
		Rating:        Objek.Rating,
		TipeKendaraan: Objek.TipeKendaraan,
		CreatedAt:     Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Kurir = se_models.Kurir{
		ID:            Objek.ID,
		Nama:          Objek.Nama,
		Username:      Objek.Username,
		Email:         Objek.Email,
		Jenis:         Objek.Jenis,
		PasswordHash:  Objek.PasswordHash,
		Deskripsi:     Objek.Deskripsi,
		StatusKurir:   Objek.StatusKurir,
		StatusBid:     Objek.StatusBid,
		VerifiedKurir: Objek.VerifiedKurir,
		Rating:        Objek.Rating,
		TipeKendaraan: Objek.TipeKendaraan,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
		DeletedAt:     &Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.KurirIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke search engine dengan info: %s ", task_info.IndexUID)
	}

	if err := cache_db_function.UpdateSessionData(ctx, *rds_session, cache_db_session.GetSessionKey(&Objek), Objek); err != nil {
		return fmt.Errorf("gagal mengupdate session data %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Kurir: Sinkronisasi data personal profiling ke perangkat lokal
	if Objek.ID != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.ID,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Profil Personal Diperbarui",
			Pesan:     "Data personal profil kurir Anda berhasil diperbarui di sistem.",
			Pop:       0, // Background sync tanpa mengganggu kurir
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"kurir_id": Objek.ID, "sync_type": "PERSONAL_PROFILING"},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_PROFILE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateGeneralProfilingKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services string = "UpdateGeneralProfilingKurir"
	var Objek sot_models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.Kurir = cass_models.Kurir{
		ID:            Objek.ID,
		Nama:          Objek.Nama,
		Username:      Objek.Username,
		Email:         Objek.Email,
		Jenis:         Objek.Jenis,
		PasswordHash:  Objek.PasswordHash,
		Deskripsi:     Objek.Deskripsi,
		StatusKurir:   Objek.StatusKurir,
		StatusBid:     Objek.StatusBid,
		VerifiedKurir: Objek.VerifiedKurir,
		Rating:        Objek.Rating,
		TipeKendaraan: Objek.TipeKendaraan,
		CreatedAt:     Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Kurir = se_models.Kurir{
		ID:            Objek.ID,
		Nama:          Objek.Nama,
		Username:      Objek.Username,
		Email:         Objek.Email,
		Jenis:         Objek.Jenis,
		PasswordHash:  Objek.PasswordHash,
		Deskripsi:     Objek.Deskripsi,
		StatusKurir:   Objek.StatusKurir,
		StatusBid:     Objek.StatusBid,
		VerifiedKurir: Objek.VerifiedKurir,
		Rating:        Objek.Rating,
		TipeKendaraan: Objek.TipeKendaraan,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
		DeletedAt:     &Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.KurirIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke search engine dengan info: %s ", task_info.IndexUID)
	}

	if err := cache_db_function.UpdateSessionData(ctx, *rds_session, cache_db_session.GetSessionKey(&Objek), Objek); err != nil {
		return fmt.Errorf("gagal mengupdate session data %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Kurir: Sinkronisasi data general profiling ke perangkat lokal
	if Objek.ID != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.ID,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Profil Umum Diperbarui",
			Pesan:     "Data umum profil kurir Anda berhasil diperbarui di sistem.",
			Pop:       0, // Background sync tanpa mengganggu kurir
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"kurir_id": Objek.ID, "sync_type": "GENERAL_PROFILING"},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_PROFILE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

