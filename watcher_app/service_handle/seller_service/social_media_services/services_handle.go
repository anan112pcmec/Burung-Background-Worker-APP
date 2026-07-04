package social_media_seller_handle

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

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

func CreateEngageSocialMediaSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateEngageSocialMediaSeller"
	var Objek sot_models.EntitySocialMedia
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.EntitySocialMedia = cass_models.EntitySocialMedia{
		ID:         Objek.ID,
		EntityId:   Objek.EntityId,
		Whatsapp:   Objek.Whatsapp, // Jika nama field asal typo, mengikuti model bawaan Anda
		Facebook:   Objek.Facebook,
		TikTok:     Objek.TikTok,
		Instagram:  Objek.Instagram,
		Metadata:   Objek.Metadata,
		EntityType: Objek.EntityType,
		CreatedAt:  Objek.CreatedAt,
		UpdatedAt:  Objek.UpdatedAt,
	}

	// Menyesuaikan jika ada typo penamaan internal pada mapping manual Anda

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” SISTEM NOTIFIKASI
	if Objek.EntityId != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.EntityId),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸŒ Media Sosial Ditautkan",
			Pesan:     "Akun media sosial baru berhasil dihubungkan ke profil toko Anda.",
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
				Metadata: map[string]interface{}{"seller_id": Objek.EntityId, "social_media_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SOCIAL_MEDIA_CONNECTED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func UpdateEngageSocialMediaSeller(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEngageSocialMediaSeller"
	var Objek sot_models.EntitySocialMedia
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
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
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” SISTEM NOTIFIKASI
	if Objek.EntityId != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.EntityId),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "âœï¸ Media Sosial Diperbarui",
			Pesan:     "Perubahan tautan atau informasi saluran media sosial toko telah berhasil disimpan.",
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
				Metadata: map[string]interface{}{"seller_id": Objek.EntityId, "social_media_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SOCIAL_MEDIA_UPDATED"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}


