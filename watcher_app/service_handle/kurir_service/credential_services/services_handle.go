package credential_kurir_handle

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

func UpdateValidateUbahPasswordKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	var Objek sot_models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	if err := cache_db_function.UpdateSessionData(ctx, *rds_session, cache_db_session.GetSessionKey(&Objek), Objek); err != nil {
		fmt.Println("gagal update data session kurir ke cache")
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

	parsedData := ObjekCass.ParseToCUDType()

	// Gunakan cass_sot_replica untuk tabel replica (Bukan cass_historcal)
	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return err
	}

	pencatatan := historical_format.Sekarang()
	parsedData["tahun_update"] = pencatatan.TahunUpdate
	parsedData["bulan_update"] = pencatatan.BulanUpdate
	parsedData["event_time"] = pencatatan.EventTime

	// Perbaikan Bug: parsedData WAJIB dipassing ke parameter ke-4 agar data historical tersimpan
	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return err
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
	}

	// Perbaikan Bug: `strconv.FormatInt` butuh basis angka 10 (desimal), bukan 0 (nol akan bikin panic/error)
	// Gunakan nil pada parameter kedua AddDocuments jika meilisearch-go Anda versi baru
	if task_info, err := se_index.AlamatKurir.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("task dengan id:" + task_info.IndexUID + " diproses")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil

}
