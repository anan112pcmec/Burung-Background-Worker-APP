package consume_pengguna_dispatcher

import (
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	alamat_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service/pengguna_service/alamat_services"
	barang_pengguna_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service/pengguna_service/barang_services"

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
