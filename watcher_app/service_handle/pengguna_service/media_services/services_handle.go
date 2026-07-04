package media_pengguna_handle

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"gorm.io/gorm"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/cache"
	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	notification_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/models"
	notification_request "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/request"
	notification_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/seeders"
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
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	// Perbaikan Bug: Memasukkan parsedData ke parameter UpdateData
	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
			return fmt.Errorf("gagal memasukan data ke dalam sot replica sync %s dalam %s", err, handle_services)
		}
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historica db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Pengguna: Refresh URL / State Foto Profil di UI Aplikasi
	if Objek.IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: Objek.IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸ”„ Foto Profil Diperbarui",
			Pesan:      "Foto profil Anda berhasil disinkronisasi.",
			Pop:        0, // Silent sync
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"pengguna_id": Objek.IdPengguna, "media_id": Objek.ID, "key": Objek.Key},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_PROFILE_IMAGE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
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
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica sync %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historica db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Pengguna: Hapus avatar/kembalikan ke default placeholder di UI
	if Objek.IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: Objek.IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸ—‘ï¸ Foto Profil Dihapus",
			Pesan:      "Foto profil Anda telah dihapus.",
			Pop:        0, // Silent sync
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"pengguna_id": Objek.IdPengguna},
				Special:  map[string]interface{}{"click_action": "SILENT_REMOVE_PROFILE_IMAGE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
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

