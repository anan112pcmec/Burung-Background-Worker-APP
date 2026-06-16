package consume_seller_dispatcher

import (
	"fmt"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
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
	case mb_cud_serializer.ConsumeDataProto:
		d = v.Parse()
	default:
		return fmt.Errorf("unsupported data type")
	}

	switch d.TableName {
	case sot_models.AlamatGudang{}.TableName(): // 1
		if err := alamat_seller_handle.DeleteHapusAlamatGudang(d); err != nil {
			return err
		}
	case sot_models.BarangInduk{}.TableName(): // 2
		if err := barang_seller_handle.DeleteHapusBarangInduk(d); err != nil {
			return err
		}
	case sot_models.KategoriBarang{}.TableName(): // 3
		if err := barang_seller_handle.DeleteHapusBarangKategori(d); err != nil {
			return err
		}
	case sot_models.Komentar{}.TableName(): // 4
		if err := barang_seller_handle.DeleteHapusKomentarBarang(d); err != nil {
			return err
		}
	case sot_models.KomentarChild{}.TableName(): // 5
		if err := barang_seller_handle.DeleteHapusChildKomentar(d); err != nil {
			return err
		}
	case sot_models.RekeningSeller{}.TableName(): // 6
		if err := credential_seller_handle.DeleteHapusRekeningSeller(d); err != nil {
			return err
		}
	case sot_models.DiskonProduk{}.TableName(): // 7
		if err := diskon_seller_handle.DeleteHapusDiskonProduk(d); err != nil {
			return err
		}
	case sot_models.BarangDiDiskon{}.TableName(): // 8
		if err := diskon_seller_handle.DeleteHapusDiskonPadaBarang(d); err != nil {
			return err
		}
	case sot_models.Etalase{}.TableName(): // 9
		if err := etalase_seller_handle.DeleteHapusEtalaseSeller(d); err != nil {
			return err
		}
	case sot_models.BarangKeEtalase{}.TableName(): // 10
		if err := etalase_seller_handle.DeleteHapusBarangDariEtalase(d); err != nil {
			return err
		}
	case sot_models.DistributorData{}.TableName(): // 11
		if err := jenis_seller_handle.DeleteHapusDataDistributor(d); err != nil {
			return err
		}
	case sot_models.BrandData{}.TableName(): // 12
		if err := jenis_seller_handle.DeleteHapusDataBrand(d); err != nil {
			return err
		}
	case sot_models.MediaSellerProfilFoto{}.TableName(): // 13
		if err := media_seller_handle.DeleteHapusFotoProfilSeller(d); err != nil {
			return err
		}
	case sot_models.MediaSellerBannerFoto{}.TableName(): // 14
		if err := media_seller_handle.DeleteHapusFotoBannerSeller(d); err != nil {
			return err
		}
	case sot_models.MediaSellerBannerFoto{}.TableName(): // 15 (Duplikat bawaan asli)
		if err := media_seller_handle.DeleteHapusFotoBannerSeller(d); err != nil {
			return err
		}
	case sot_models.MediaEtalaseFoto{}.TableName(): // 16
		if err := media_seller_handle.DeleteHapusFotoEtalaseSeller(d); err != nil {
			return err
		}
	case sot_models.MediaBarangIndukFoto{}.TableName(): // 17
		if err := media_seller_handle.DeleteHapusMediaBarangIndukFoto(d); err != nil {
			return err
		}
	case sot_models.MediaBarangIndukVideo{}.TableName(): // 18
		if err := media_seller_handle.DeleteHapusBarangIndukVideo(d); err != nil {
			return err
		}
	case sot_models.MediaKategoriBarangFoto{}.TableName(): // 19
		if err := media_seller_handle.DeleteHapusKategoriBarangFoto(d); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataDokumen{}.TableName(): // 20
		if err := media_seller_handle.DeleteHapusMediaDistributorDataDokumen(d); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataNPWPFoto{}.TableName(): // 21
		if err := media_seller_handle.DeleteHapusMediaDistributorDataNPWPFoto(d); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataNIBFoto{}.TableName(): // 22
		if err := media_seller_handle.DeleteHapusDistributorDataNIBFoto(d); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataSuratKerjasamaDokumen{}.TableName(): // 23
		if err := media_seller_handle.DeleteHapusMediaDistributorDataDokumen(d); err != nil {
			return err
		}
	case sot_models.MediaBrandDataPerwakilanDokumen{}.TableName(): // 24
		if err := media_seller_handle.DeleteHapusBrandDataPerwakilanDokumen(d); err != nil {
			return err
		}
	case sot_models.MediaBrandDataPerwakilanDokumen{}.TableName(): // 25 (Duplikat bawaan asli)
		if err := media_seller_handle.DeleteHapusBrandDataPerwakilanDokumen(d); err != nil {
			return err
		}
	case sot_models.MediaBrandDataSertifikatFoto{}.TableName(): // 26
		if err := media_seller_handle.DeleteHapusMediaBrandDataSertifikatFoto(d); err != nil {
			return err
		}
	case sot_models.MediaBrandDataNIBFoto{}.TableName(): // 27
		if err := media_seller_handle.DeleteHapusMediaBrandDataNIBFoto(d); err != nil {
			return err
		}
	case sot_models.MediaBrandDataNPWPFoto{}.TableName(): // 28
		if err := media_seller_handle.DeleteHapusMediaBrandNPWPFoto(d); err != nil {
			return err
		}
	case sot_models.MediaBrandDataLogoFoto{}.TableName(): // 29
		if err := media_seller_handle.DeleteHapusMediaBrandDataLogoFoto(d); err != nil {
			return err
		}
	case sot_models.MediaBrandDataSuratKerjasamaDokumen{}.TableName(): // 30
		if err := media_seller_handle.DeleteHapusBrandDataSuratKerjasamaDokumen(d); err != nil {
			return err
		}
	}

	return nil
}
