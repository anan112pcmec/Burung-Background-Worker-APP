package consume_pengguna_dispatcher

import (
	"context"
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	se_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/search_engine/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	auth_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/auth"
	alamat_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/alamat_services"
	barang_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/barang_services"
	credential_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/credential_services"
	media_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/media_services"
	profiling_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/profiling_services"
	social_media_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/social_media_services"
)

func PenggunaUpdateServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](ctx context.Context, data *T, read *gorm.DB, redis_authentication, redis_session *redis.Client, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {

	var d mb_cud_serializer.ParsedDataMessage
	switch v := any(data).(type) {
	case *mb_cud_serializer.ConsumeDataJson:
		d = v.Parse()
	case *mb_cud_serializer.ConsumeDataProto:
		d = v.Parse()
	default:
		return fmt.Errorf("unsupported data type")
	}

	switch d.TableName {
	case "PenggunaLogin":
		if err := auth_handle.UpdatePenggunaLogin(d, ctx, cass_historcal, cass_sot_replica, redis_session); err != nil {
			return err
		}
	case sot_models.AlamatPengguna.TableName(sot_models.AlamatPengguna{}):
		if err := alamat_pengguna_handle.UpdateAlamatPub(d, ctx, cass_historcal, cass_sot_replica, se_index); err != nil {
			return err
		}
	case sot_models.Komentar.TableName(sot_models.Komentar{}):
		if err := barang_pengguna_handle.UpdateEditKomentarBarang(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.KomentarChild.TableName(sot_models.KomentarChild{}):
		if err := barang_pengguna_handle.UpdateEditChildKomentar(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.Keranjang.TableName(sot_models.Keranjang{}):
		if err := barang_pengguna_handle.UpdateEditKeranjangBarang(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "UpdateBerikanReviewBarangIIUpdateTransaksi":
		if err := barang_pengguna_handle.UpdateBerikanReviewBarangIIUpdateTransaksi(d, ctx, cass_historcal, cass_sot_replica, se_index); err != nil {
			return err
		}
	case "ValidateUbahPasswordPenggunaViaOtp":
		if err := credential_pengguna_handle.UpdateValidateUbahPasswordPenggunaViaOtp(d, ctx, read, cass_historcal, cass_sot_replica, se_index, redis_session); err != nil {
			return err
		}
	case "ValidateUbahPasswordPenggunaViaPin":
		if err := credential_pengguna_handle.UpdateValidateUbahPasswordPenggunaViaPin(d, ctx, read, cass_historcal, cass_sot_replica, se_index, redis_session); err != nil {
			return err
		}
	case "UpdateSecretPinPengguna":
		if err := credential_pengguna_handle.UpdateSecretPinPengguna(d, ctx, read, cass_historcal, cass_sot_replica, se_index, redis_session); err != nil {
			return err
		}
	case sot_models.MediaPenggunaProfilFoto.TableName(sot_models.MediaPenggunaProfilFoto{}):
		if err := media_pengguna_handle.UpdateUbahFotoProfilPengguna(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.Pengguna.TableName(sot_models.Pengguna{}):
		if err := profiling_pengguna_handle.UpdateUbahPersonalProfilingPengguna(d, ctx, read, cass_historcal, cass_sot_replica, se_index, redis_session); err != nil {
			return err
		}
	case "EngageTautkanSocialMediaPengguna":
		if err := social_media_pengguna_handle.UpdateEngageTautkanSocialMediaPengguna(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "EngageHapusSocialMedia":
		if err := social_media_pengguna_handle.UpdateEngageHapusSocialMedia(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}

	}

	return nil
}
