package social_media_kurir_handle

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/environment"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	notification_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/models"
	notification_request "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/request"
	notification_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/seeders"
)

func CreateEngagementSocialMediaKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateEngagementSocialMediaKurir"
	var Objek sot_models.EntitySocialMedia

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
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
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 🔔 Silent Update Kurir: Sinkronisasi awal pembuatan link media sosial kurir
	if Objek.EntityId != 0 && Objek.EntityType == "kurir" {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.EntityId,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "🔄 Media Sosial Ditautkan",
			Pesan:     "Akun media sosial Anda berhasil diintegrasikan ke dalam sistem.",
			Pop:       0, // Silent sync background
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"entity_id": Objek.EntityId, "social_media_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_SOCIAL_LINKS"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, environment.HostRunningAPIInNotifikasi, environment.PortRunningAPIInNotifikasi, environment.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateEngagementSocialMediaKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEngagementSocialMediaKurir"
	var Objek sot_models.EntitySocialMedia

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
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
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 🔔 Silent Update Kurir: Sinkronisasi perubahan link media sosial kurir
	if Objek.EntityId != 0 && Objek.EntityType == "kurir" {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.EntityId,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "🔄 Kontak Media Sosial Diperbarui",
			Pesan:     "Perubahan data akun media sosial Anda berhasil disinkronisasi.",
			Pop:       0, // Silent sync background
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"entity_id": Objek.EntityId, "social_media_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_SOCIAL_LINKS"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, environment.HostRunningAPIInNotifikasi, environment.PortRunningAPIInNotifikasi, environment.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
