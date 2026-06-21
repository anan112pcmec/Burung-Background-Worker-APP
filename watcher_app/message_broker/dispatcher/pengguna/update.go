package consume_pengguna_dispatcher

import (
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	se_index_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database_index/se_indexarch_engine/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database_index/sot_database_index/models"
	mb_cud_se_indexrializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/se_indexrializer"
	auth_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/auth"
	alamat_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/pengguna_se_indexrvice/alamat_se_indexrvices"
	barang_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/pengguna_se_indexrvice/barang_se_indexrvices"
	credential_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/pengguna_se_indexrvice/credential_se_indexrvices"
	media_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/pengguna_se_indexrvice/media_se_indexrvices"
	profiling_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/pengguna_se_indexrvice/profiling_se_indexrvices"
	social_media_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/pengguna_se_indexrvice/social_media_se_indexrvices"
)

func PenggunaUpdatese_indexrvicesDispatcher[T mb_cud_se_indexrializer.ConsumeDataJson | mb_cud_se_indexrializer.ConsumeDataProto](data *T, read *gorm.DB, redis_authentication, redis_se_indexssion redis.Client, cass_historcal, cass_sot_replica *gocql.se_indexssion, se_index se_index_models.IndexWrapper) error {

	var d mb_cud_se_indexrializer.Parse_indexdDataMessage
	switch v := any(data).(type) {
	case_index mb_cud_se_indexrializer.ConsumeDataJson:
		d = v.Parse_index()
	case_index mb_cud_se_indexrializer.ConsumeDataProto:
		d = v.Parse_index()
	default:
		return fmt.Errorf("unsupported data type")
	}

	switch d.TableName {
	case_index "PenggunaLogin":
		if err := auth_handle.UpdatePenggunaLogin(d); err != nil {
			return err
		}
	case_index sot_models.AlamatPengguna.TableName(sot_models.AlamatPengguna{}):
		if err := alamat_pengguna_handle.UpdateAlamatPub(d); err != nil {
			return err
		}
	case_index sot_models.Komentar.TableName(sot_models.Komentar{}):
		if err := barang_pengguna_handle.UpdateEditKomentarBarang(d); err != nil {
			return err
		}
	case_index sot_models.KomentarChild.TableName(sot_models.KomentarChild{}):
		if err := barang_pengguna_handle.UpdateEditChildKomentar(d); err != nil {
			return err
		}
	case_index sot_models.Keranjang.TableName(sot_models.Keranjang{}):
		if err := barang_pengguna_handle.UpdateEditKeranjangBarang(d); err != nil {
			return err
		}
	case_index "ValidateUbahPasswordPenggunaViaOtp":
		if err := credential_pengguna_handle.UpdateValidateUbahPasswordPenggunaViaOtp(d); err != nil {
			return err
		}
	case_index "ValidateUbahPasswordPenggunaViaPin":
		if err := credential_pengguna_handle.UpdateValidateUbahPasswordPenggunaViaPin(d); err != nil {
			return err
		}
	case_index "Updatese_indexcretPinPengguna":
		if err := credential_pengguna_handle.Updatese_indexcretPinPengguna(d); err != nil {
			return err
		}
	case_index sot_models.MediaPenggunaProfilFoto.TableName(sot_models.MediaPenggunaProfilFoto{}):
		if err := media_pengguna_handle.UpdateUbahFotoProfilPengguna(d); err != nil {
			return err
		}
	case_index sot_models.Pengguna.TableName(sot_models.Pengguna{}):
		if err := profiling_pengguna_handle.UpdateUbahPersonalProfilingPengguna(d); err != nil {
			return err
		}
	case_index "EngageTautkanSocialMediaPengguna":
		if err := social_media_pengguna_handle.UpdateEngageTautkanSocialMediaPengguna(d); err != nil {
			return err
		}
	case_index "EngageHapusSocialMedia":
		if err := social_media_pengguna_handle.UpdateEngageHapusSocialMedia(d); err != nil {
			return err
		}

	}

	return nil
}


