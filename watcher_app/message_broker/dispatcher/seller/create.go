package consume_seller_dispatcher

import (
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	alamat_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/alamat_services"
	barang_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/barang_services"

)

func SellerCreateServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](data T) error {
	var d mb_cud_serializer.ParsedDataMessage
	switch v := any(data).(type) {
	case mb_cud_serializer.ConsumeDataJson:
		d = v.Parse()

		switch d.TableName {
		case models.AlamatGudang.TableName(models.AlamatGudang{}):
			if err := alamat_seller_handle.CreateTambahAlamatGudang(d); err != nil {
				return err
			}
		case models.BarangInduk.TableName(models.BarangInduk{}):
			if err := barang_seller_handle.CreateMasukanBarangInduk(d); err != nil {
				return err
			}
		case models.KategoriBarang.TableName(models.KategoriBarang{}):
			if err := barang_seller_handle.CreateMasukanKategoriBarang(d); err != nil {
				return err
			}
		case models.Komentar.TableName(models.Komentar{}):
			if err := barang_seller_handle.CreateMasukanKomentarBarang(d); err != nil {
				return err
			}
		case models.KomentarChild.TableName(models.KomentarChild{}):
			if err := barang_seller_handle.CreateMasukanChildKomentar(d); err != nil {
				return err
			}

		}

	case mb_cud_serializer.ConsumeDataProto:
		d = v.Parse()

		switch d.TableName {
		case models.AlamatGudang.TableName(models.AlamatGudang{}):
			if err := alamat_seller_handle.CreateTambahAlamatGudang(d); err != nil {
				return err
			}

		}
	}

	return nil
}
