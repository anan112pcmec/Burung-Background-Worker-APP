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
	media_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/media_services"
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
		case models.MediaSellerProfilFoto.TableName(models.MediaSellerProfilFoto{}):
			if err := media_seller_handle.DeleteHapusFotoProfilSeller(d); err != nil {
				return err
			}
		case models.MediaSellerBannerFoto.TableName(models.MediaSellerBannerFoto{}):
			if err := media_seller_handle.DeleteHapusFotoBannerSeller(d); err != nil {
				return err
			}
		case models.MediaSellerBannerFoto.TableName(models.MediaSellerBannerFoto{}):
			if err := media_seller_handle.DeleteHapusFotoBannerSeller(d); err != nil {
				return err
			}
		case models.MediaEtalaseFoto.TableName(models.MediaEtalaseFoto{}):
			if err := media_seller_handle.DeleteHapusFotoEtalaseSeller(d); err != nil {
				return err
			}
		case models.MediaBarangIndukFoto.TableName(models.MediaBarangIndukFoto{}):
			if err := media_seller_handle.DeleteHapusMediaBarangIndukFoto(d); err != nil {
				return err
			}
		case models.MediaBarangIndukVideo.TableName(models.MediaBarangIndukVideo{}):
			if err := media_seller_handle.DeleteHapusBarangIndukVideo(d); err != nil {
				return err
			}
		case models.MediaKategoriBarangFoto.TableName(models.MediaKategoriBarangFoto{}):
			if err := media_seller_handle.DeleteHapusKategoriBarangFoto(d); err != nil {
				return err
			}
		case models.MediaDistributorDataDokumen.TableName(models.MediaDistributorDataDokumen{}):
			if err := media_seller_handle.DeleteHapusMediaDistributorDataDokumen(d); err != nil {
				return err
			}
		case models.MediaDistributorDataNPWPFoto.TableName(models.MediaDistributorDataNPWPFoto{}):
			if err := media_seller_handle.DeleteHapusMediaDistributorDataNPWPFoto(d); err != nil {
				return err
			}
		case models.MediaDistributorDataNIBFoto.TableName(models.MediaDistributorDataNIBFoto{}):
			if err := media_seller_handle.DeleteHapusDistributorDataNIBFoto(d); err != nil {
				return err
			}
		case models.MediaDistributorDataSuratKerjasamaDokumen.TableName(models.MediaDistributorDataSuratKerjasamaDokumen{}):
			if err := media_seller_handle.DeleteHapusMediaDistributorDataDokumen(d); err != nil {
				return err
			}
		case models.MediaBrandDataPerwakilanDokumen.TableName(models.MediaBrandDataPerwakilanDokumen{}):
			if err := media_seller_handle.DeleteHapusBrandDataPerwakilanDokumen(d); err != nil {
				return err
			}
		case models.MediaBrandDataPerwakilanDokumen.TableName(models.MediaBrandDataPerwakilanDokumen{}):
			if err := media_seller_handle.DeleteHapusBrandDataPerwakilanDokumen(d); err != nil {
				return err
			}
		case models.MediaBrandDataSertifikatFoto.TableName(models.MediaBrandDataSertifikatFoto{}):
			if err := media_seller_handle.DeleteHapusMediaBrandDataSertifikatFoto(d); err != nil {
				return err
			}
		case models.MediaBrandDataNIBFoto.TableName(models.MediaBrandDataNIBFoto{}):
			if err := media_seller_handle.DeleteHapusMediaBrandDataNIBFoto(d); err != nil {
				return err
			}
		case models.MediaBrandDataNPWPFoto.TableName(models.MediaBrandDataNPWPFoto{}):
			if err := media_seller_handle.DeleteHapusMediaBrandNPWPFoto(d); err != nil {
				return err
			}
		case models.MediaBrandDataLogoFoto.TableName(models.MediaBrandDataLogoFoto{}):
			if err := media_seller_handle.DeleteHapusMediaBrandDataLogoFoto(d); err != nil {
				return err
			}
		case models.MediaBrandDataSuratKerjasamaDokumen.TableName(models.MediaBrandDataSuratKerjasamaDokumen{}):
			if err := media_seller_handle.DeleteHapusBrandDataSuratKerjasamaDokumen(d); err != nil {
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
