package consume_pengguna_dispatcher

import (
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	alamat_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/alamat_services"
	barang_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/barang_services"
	credential_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/credential_services"
	media_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/media_services"
	social_media_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/social_media_services"
	wishlist_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/wishlist_services"
)

func PenggunaCreateServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](data *T) error {

	var d mb_cud_serializer.ParsedDataMessage
	switch v := any(data).(type) {

	case mb_cud_serializer.ConsumeDataJson:

		d = v.Parse()

		switch d.TableName {
		case models.AlamatPengguna.TableName(models.AlamatPengguna{}):
			if err := alamat_pengguna_handle.CreateAlamatPub(d); err != nil {
				return err
			}
		case models.BarangDisukai.TableName(models.BarangDisukai{}):
			if err := barang_pengguna_handle.CreateLikesBarang(d); err != nil {
				return err
			}
		case models.Komentar.TableName(models.Komentar{}):
			if err := barang_pengguna_handle.CreateMasukanKomentarBarang(d); err != nil {
				return err
			}
		case models.KomentarChild.TableName(models.KomentarChild{}):
			if err := barang_pengguna_handle.CreateMasukanChildKomentar(d); err != nil {
				return err
			}
		case models.Keranjang.TableName(models.Keranjang{}):
			if err := barang_pengguna_handle.CreateTambahKeranjangBarang(d); err != nil {
				return err
			}
		case models.Review.TableName(models.Review{}):
			if err := barang_pengguna_handle.CreateBerikanReviewBarang(d); err != nil {
				return err
			}
		case "MembuatSecretPinPengguna":
			if err := credential_pengguna_handle.CreateMembuatSecretPinPengguna(d); err != nil {
				return err
			}
		case models.MediaReviewFoto.TableName(models.MediaReviewFoto{}):
			if err := media_pengguna_handle.CreateTambahMediaReviewFoto(d); err != nil {
				return err
			}
		case models.MediaReviewVideo.TableName(models.MediaReviewVideo{}):
			if err := media_pengguna_handle.CreateTambahMediaReviewVideo(d); err != nil {
				return err
			}
		case models.EntitySocialMedia.TableName(models.EntitySocialMedia{}):
			if err := social_media_pengguna_handle.CreateEngageTautkanSocialMediaPengguna(d); err != nil {
				return err
			}
		case models.Follower.TableName(models.Follower{}):
			if err := social_media_pengguna_handle.CreateFollowSeller(d); err != nil {
				return err
			}
		case models.Wishlist.TableName(models.Wishlist{}):
			if err := wishlist_pengguna_handle.CreateTambahBarangKeWishlist(d); err != nil {
				return err
			}

		}

	case mb_cud_serializer.ConsumeDataProto:
		d = v.Parse()

		switch d.TableName {
		case models.AlamatPengguna.TableName(models.AlamatPengguna{}):
			if err := alamat_pengguna_handle.CreateAlamatPub(d); err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("unsupported data type")
	}

	return nil
}
