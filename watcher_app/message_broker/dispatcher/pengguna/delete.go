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
	alamat_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/alamat_services"
	barang_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/barang_services"
	media_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/media_services"
	social_media_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/social_media_services"
	transaction_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/transaction_services"
	wishlist_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/wishlist_services"
)

func PenggunaDeleteServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](ctx context.Context, data *T, read *gorm.DB, redis_authentication, redis_session *redis.Client, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {

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
	case sot_models.AlamatPengguna.TableName(sot_models.AlamatPengguna{}):
		if err := alamat_pengguna_handle.DeleteAlamatPub(d, ctx, cass_historcal, cass_sot_replica, se_index); err != nil {
			return err
		}
	case sot_models.BarangDisukai.TableName(sot_models.BarangDisukai{}):
		if err := barang_pengguna_handle.DeleteUnlikesBarang(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.Komentar.TableName(sot_models.Komentar{}):
		if err := barang_pengguna_handle.DeleteHapusKomentarBarang(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.KomentarChild.TableName(sot_models.KomentarChild{}):
		if err := barang_pengguna_handle.DeleteHapusChildKomentar(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.Keranjang.TableName(sot_models.Keranjang{}):
		if err := barang_pengguna_handle.DeleteHapusKeranjangBarang(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaPenggunaProfilFoto.TableName(sot_models.MediaPenggunaProfilFoto{}):
		if err := media_pengguna_handle.DeleteHapusFotoProfilPengguna(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.Follower.TableName(sot_models.Follower{}):
		if err := social_media_pengguna_handle.DeleteUnfollowSeller(d, read, ctx, cass_historcal, cass_sot_replica, se_index); err != nil {
			return err
		}
	case sot_models.Wishlist.TableName(sot_models.Wishlist{}):
		if err := wishlist_pengguna_handle.DeleteHapusBarangDariWishlist(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.Keranjang.TableName(sot_models.Keranjang{}):
		if err := transaction_pengguna_handle.DeleteCheckoutBarangUser(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.ReviewLike{}.TableName():
		if err := barang_pengguna_handle.DeleteUnlikeReviewBarang(d, ctx, cass_historcal, cass_sot_replica, se_index); err != nil {
			return err
		}
	}

	return nil
}
