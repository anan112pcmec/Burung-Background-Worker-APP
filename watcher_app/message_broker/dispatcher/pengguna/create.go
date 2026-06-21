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
	social_media_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/pengguna_se_indexrvice/social_media_se_indexrvices"
	transaction_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/pengguna_se_indexrvice/transaction_se_indexrvices"
	wishlist_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/pengguna_se_indexrvice/wishlist_se_indexrvices"
)

func PenggunaCreatese_indexrvicesDispatcher[T mb_cud_se_indexrializer.ConsumeDataJson | mb_cud_se_indexrializer.ConsumeDataProto](data *T, read *gorm.DB, redis_authentication, redis_se_indexssion redis.Client, cass_historcal, cass_sot_replica *gocql.se_indexssion, se_index se_index_models.IndexWrapper) error {

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
	case_index sot_models.Pengguna{}.TableName():
		if err := auth_handle.CreateValidatePenggunaRegistration(d); err != nil {
			return err
		}
	case_index sot_models.AlamatPengguna.TableName(sot_models.AlamatPengguna{}):
		if err := alamat_pengguna_handle.CreateAlamatPub(d); err != nil {
			return err
		}
	case_index sot_models.BarangDisukai.TableName(sot_models.BarangDisukai{}):
		if err := barang_pengguna_handle.CreateLikesBarang(d); err != nil {
			return err
		}
	case_index sot_models.Komentar.TableName(sot_models.Komentar{}):
		if err := barang_pengguna_handle.CreateMasukanKomentarBarang(d); err != nil {
			return err
		}
	case_index sot_models.KomentarChild.TableName(sot_models.KomentarChild{}):
		if err := barang_pengguna_handle.CreateMasukanChildKomentar(d); err != nil {
			return err
		}
	case_index sot_models.Keranjang.TableName(sot_models.Keranjang{}):
		if err := barang_pengguna_handle.CreateTambahKeranjangBarang(d); err != nil {
			return err
		}
	case_index sot_models.Review.TableName(sot_models.Review{}):
		if err := barang_pengguna_handle.CreateBerikanReviewBarang(d); err != nil {
			return err
		}
	case_index "Membuatse_indexcretPinPengguna":
		if err := credential_pengguna_handle.CreateMembuatse_indexcretPinPengguna(d); err != nil {
			return err
		}
	case_index sot_models.MediaReviewFoto.TableName(sot_models.MediaReviewFoto{}):
		if err := media_pengguna_handle.CreateTambahMediaReviewFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaReviewVideo.TableName(sot_models.MediaReviewVideo{}):
		if err := media_pengguna_handle.CreateTambahMediaReviewVideo(d); err != nil {
			return err
		}
	case_index sot_models.EntitySocialMedia.TableName(sot_models.EntitySocialMedia{}):
		if err := social_media_pengguna_handle.CreateEngageTautkanSocialMediaPengguna(d); err != nil {
			return err
		}
	case_index sot_models.Follower.TableName(sot_models.Follower{}):
		if err := social_media_pengguna_handle.CreateFollowse_indexller(d); err != nil {
			return err
		}
	case_index sot_models.Wishlist.TableName(sot_models.Wishlist{}):
		if err := wishlist_pengguna_handle.CreateTambahBarangKeWishlist(d); err != nil {
			return err
		}
	case_index "LockTransaksiVa":
		if err := transaction_pengguna_handle.CreateLockTransaksiVa(d); err != nil {
			return err
		}
	case_index "LockTransaksiWallet":
		if err := transaction_pengguna_handle.CreateLockTransaksiWallet(d); err != nil {
			return err
		}
	case_index "LockTransaksiGerai":
		if err := transaction_pengguna_handle.CreateLockTransaksiGerai(d); err != nil {
			return err
		}

	}

	return nil
}



