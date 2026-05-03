package consume_pengguna_dispatcher

import (
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	alamat_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/alamat_services"
	barang_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/barang_services"
	media_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/media_services"
	social_media_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/social_media_services"
	transaction_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/transaction_services"
	wishlist_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/pengguna_service/wishlist_services"
)

func PenggunaDeleteServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](data *T) error {

	var d mb_cud_serializer.ParsedDataMessage
	switch v := any(data).(type) {

	case mb_cud_serializer.ConsumeDataJson:

		d = v.Parse()

		switch d.TableName {
		case models.AlamatPengguna.TableName(models.AlamatPengguna{}):
			if err := alamat_pengguna_handle.DeleteAlamatPub(d); err != nil {
				return err
			}
		case models.BarangDisukai.TableName(models.BarangDisukai{}):
			if err := barang_pengguna_handle.DeleteUnlikesBarang(d); err != nil {
				return err
			}
		case models.Komentar.TableName(models.Komentar{}):
			if err := barang_pengguna_handle.DeleteHapusKomentarBarang(d); err != nil {
				return err
			}
		case models.KomentarChild.TableName(models.KomentarChild{}):
			if err := barang_pengguna_handle.DeleteHapusChildKomentar(d); err != nil {
				return err
			}
		case models.Keranjang.TableName(models.Keranjang{}):
			if err := barang_pengguna_handle.DeleteHapusKeranjangBarang(d); err != nil {
				return err
			}
		case models.MediaPenggunaProfilFoto.TableName(models.MediaPenggunaProfilFoto{}):
			if err := media_pengguna_handle.DeleteHapusFotoProfilPengguna(d); err != nil {
				return err
			}
		case models.Follower.TableName(models.Follower{}):
			if err := social_media_pengguna_handle.DeleteUnfollowSeller(d); err != nil {
				return err
			}
		case models.Wishlist.TableName(models.Wishlist{}):
			if err := wishlist_pengguna_handle.DeleteHapusBarangDariWishlist(d); err != nil {
				return err
			}
		case models.Keranjang.TableName(models.Keranjang{}):
			if err := transaction_pengguna_handle.DeleteCheckoutBarangUser(d); err != nil {
				return err
			}
		}

	case mb_cud_serializer.ConsumeDataProto:
		d = v.Parse()

		switch d.TableName {
		case models.AlamatPengguna.TableName(models.AlamatPengguna{}):
			if err := alamat_pengguna_handle.DeleteAlamatPub(d); err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("unsupported data type")
	}

	return nil
}
