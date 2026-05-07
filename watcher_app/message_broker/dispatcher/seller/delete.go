package consume_seller_dispatcher

import (
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	alamat_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/alamat_services"
	barang_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/barang_services"
	credential_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/credential_services"
	diskon_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/diskon_services"
	etalase_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/etalase_services"
	jenis_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/jenis_services"
)

func SellerDeleteServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](data T) error {
	var d mb_cud_serializer.ParsedDataMessage
	switch v := any(data).(type) {
	case mb_cud_serializer.ConsumeDataJson:
		d = v.Parse()

		switch d.TableName {
		case models.AlamatGudang.TableName(models.AlamatGudang{}):
			if err := alamat_seller_handle.DeleteHapusAlamatGudang(d); err != nil {
				return err
			}

		case models.BarangInduk.TableName(models.BarangInduk{}):
			if err := barang_seller_handle.DeleteHapusBarangInduk(d); err != nil {
				return err
			}
		case models.KategoriBarang.TableName(models.KategoriBarang{}):
			if err := barang_seller_handle.DeleteHapusBarangKategori(d); err != nil {
				return err
			}
		case models.Komentar.TableName(models.Komentar{}):
			if err := barang_seller_handle.DeleteHapusKomentarBarang(d); err != nil {
				return err
			}
		case models.KomentarChild.TableName(models.KomentarChild{}):
			if err := barang_seller_handle.DeleteHapusChildKomentar(d); err != nil {
				return err
			}
		case models.RekeningSeller.TableName(models.RekeningSeller{}):
			if err := credential_seller_handle.DeleteHapusRekeningSeller(d); err != nil {
				return err
			}
		case models.DiskonProduk.TableName(models.DiskonProduk{}):
			if err := diskon_seller_handle.DeleteHapusDiskonProduk(d); err != nil {
				return err
			}
		case models.BarangDiDiskon.TableName(models.BarangDiDiskon{}):
			if err := diskon_seller_handle.DeleteHapusDiskonPadaBarang(d); err != nil {
				return err
			}
		case models.Etalase.TableName(models.Etalase{}):
			if err := etalase_seller_handle.DeleteHapusEtalaseSeller(d); err != nil {
				return err
			}
		case models.BarangKeEtalase.TableName(models.BarangKeEtalase{}):
			if err := etalase_seller_handle.DeleteHapusBarangDariEtalase(d); err != nil {
				return err
			}
		case models.DistributorData.TableName(models.DistributorData{}):
			if err := jenis_seller_handle.DeleteHapusDataDistributor(d); err != nil {
				return err
			}
		case models.BrandData.TableName(models.BrandData{}):
			if err := jenis_seller_handle.DeleteHapusDataBrand(d); err != nil {
				return err
			}

		}

	case mb_cud_serializer.ConsumeDataProto:
		d = v.Parse()

		switch d.TableName {
		case models.AlamatGudang.TableName(models.AlamatGudang{}):
			if err := alamat_seller_handle.DeleteHapusAlamatGudang(d); err != nil {
				return err
			}

		}
	}

	return nil
}
