package consume_pengguna_dispatcher

import (
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
	social_media_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/social_media_services"
	transaction_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/transaction_services"
	wishlist_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/wishlist_services"
)

func PenggunaCreateServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](data *T, read *gorm.DB, redis_authentication, redis_session redis.Client, cass_historcal, cass_sot_replica *gocql.Session, se se_models.IndexWrapper) error {

	var d mb_cud_serializer.ParsedDataMessage
	switch v := any(data).(type) {
	case mb_cud_serializer.ConsumeDataJson:
		d = v.Parse()
	case mb_cud_serializer.ConsumeDataProto:
		d = v.Parse()
	default:
		return fmt.Errorf("unsupported data type")
	}

	switch d.TableName {
	case sot_models.Pengguna{}.TableName():
		if err := auth_handle.CreateValidatePenggunaRegistration(d); err != nil {
			return err
		}
	case sot_models.AlamatPengguna.TableName(sot_models.AlamatPengguna{}):
		if err := alamat_pengguna_handle.CreateAlamatPub(d); err != nil {
			return err
		}
	case sot_models.BarangDisukai.TableName(sot_models.BarangDisukai{}):
		if err := barang_pengguna_handle.CreateLikesBarang(d); err != nil {
			return err
		}
	case sot_models.Komentar.TableName(sot_models.Komentar{}):
		if err := barang_pengguna_handle.CreateMasukanKomentarBarang(d); err != nil {
			return err
		}
	case sot_models.KomentarChild.TableName(sot_models.KomentarChild{}):
		if err := barang_pengguna_handle.CreateMasukanChildKomentar(d); err != nil {
			return err
		}
	case sot_models.Keranjang.TableName(sot_models.Keranjang{}):
		if err := barang_pengguna_handle.CreateTambahKeranjangBarang(d); err != nil {
			return err
		}
	case sot_models.Review.TableName(sot_models.Review{}):
		if err := barang_pengguna_handle.CreateBerikanReviewBarang(d); err != nil {
			return err
		}
	case "MembuatSecretPinPengguna":
		if err := credential_pengguna_handle.CreateMembuatSecretPinPengguna(d); err != nil {
			return err
		}
	case sot_models.MediaReviewFoto.TableName(sot_models.MediaReviewFoto{}):
		if err := media_pengguna_handle.CreateTambahMediaReviewFoto(d); err != nil {
			return err
		}
	case sot_models.MediaReviewVideo.TableName(sot_models.MediaReviewVideo{}):
		if err := media_pengguna_handle.CreateTambahMediaReviewVideo(d); err != nil {
			return err
		}
	case sot_models.EntitySocialMedia.TableName(sot_models.EntitySocialMedia{}):
		if err := social_media_pengguna_handle.CreateEngageTautkanSocialMediaPengguna(d); err != nil {
			return err
		}
	case sot_models.Follower.TableName(sot_models.Follower{}):
		if err := social_media_pengguna_handle.CreateFollowSeller(d); err != nil {
			return err
		}
	case sot_models.Wishlist.TableName(sot_models.Wishlist{}):
		if err := wishlist_pengguna_handle.CreateTambahBarangKeWishlist(d); err != nil {
			return err
		}
	case "LockTransaksiVa":
		if err := transaction_pengguna_handle.CreateLockTransaksiVa(d); err != nil {
			return err
		}
	case "LockTransaksiWallet":
		if err := transaction_pengguna_handle.CreateLockTransaksiWallet(d); err != nil {
			return err
		}
	case "LockTransaksiGerai":
		if err := transaction_pengguna_handle.CreateLockTransaksiGerai(d); err != nil {
			return err
		}

	}

	return nil
}
