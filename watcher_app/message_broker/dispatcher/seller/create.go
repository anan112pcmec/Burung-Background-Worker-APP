package consume_seller_dispatcher

import (
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	auth_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/auth"
	alamat_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/alamat_services"
	barang_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/barang_services"
	credential_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/credential_services"
	diskon_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/diskon_services"
	etalase_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/etalase_services"
	jenis_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/jenis_services"
	media_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/media_services"
	social_media_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/social_media_services"
	transaksi_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/transaksi_services"
)

func SellerCreateServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](data T) error {
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
	case models.Seller{}.TableName():
		if err := auth_handle.CreateValidateSellerRegistration(d); err != nil {
			return err
		}
	case models.AlamatGudang{}.TableName():
		if err := alamat_seller_handle.CreateTambahAlamatGudang(d); err != nil {
			return err
		}
	case models.BarangInduk{}.TableName():
		if err := barang_seller_handle.CreateMasukanBarangInduk(d); err != nil {
			return err
		}
	case models.KategoriBarang{}.TableName():
		if err := barang_seller_handle.CreateMasukanKategoriBarang(d); err != nil {
			return err
		}
	case models.Komentar{}.TableName():
		if err := barang_seller_handle.CreateMasukanKomentarBarang(d); err != nil {
			return err
		}
	case models.KomentarChild{}.TableName():
		if err := barang_seller_handle.CreateMasukanChildKomentar(d); err != nil {
			return err
		}
	case models.RekeningSeller{}.TableName():
		if err := credential_seller_handle.CreateTambahRekeningSeller(d); err != nil {
			return err
		}
	case models.DiskonProduk{}.TableName():
		if err := diskon_seller_handle.CreateTambahDiskonProduk(d); err != nil {
			return err
		}
	case models.BarangDiDiskon{}.TableName():
		if err := diskon_seller_handle.CreateTetapkanDiskonPadaBarang(d); err != nil {
			return err
		}
	case models.Etalase{}.TableName():
		if err := etalase_seller_handle.CreateTambahEtalaseSeller(d); err != nil {
			return err
		}
	case models.BarangKeEtalase{}.TableName():
		if err := etalase_seller_handle.CreateTambahkanBarangKeEtalase(d); err != nil {
			return err
		}
	case models.DistributorData{}.TableName():
		if err := jenis_seller_handle.CreateMasukanDataDistributor(d); err != nil {
			return err
		}
	case models.BrandData{}.TableName():
		if err := jenis_seller_handle.CreateMasukanDataBrand(d); err != nil {
			return err
		}
	case models.MediaSellerProfilFoto{}.TableName():
		if err := media_seller_handle.CreateUbahFotoProfilSeller(d); err != nil {
			return err
		}
	case models.MediaSellerBannerFoto{}.TableName():
		if err := media_seller_handle.CreateUbahFotoBannerSeller(d); err != nil {
			return err
		}
	case models.MediaSellerTokoFisikFoto{}.TableName():
		if err := media_seller_handle.CreateTambahkanFotoTokoFisikSeller(d); err != nil {
			return err
		}
	case models.MediaEtalaseFoto{}.TableName():
		if err := media_seller_handle.CreateUbahFotoEtalaseSeller(d); err != nil {
			return err
		}
	case models.MediaBarangIndukFoto{}.TableName():
		if err := media_seller_handle.CreateTambahkanMediaBarangIndukFoto(d); err != nil {
			return err
		}
	case models.MediaBarangIndukVideo{}.TableName():
		if err := media_seller_handle.CreateUbahBarangIndukVideo(d); err != nil {
			return err
		}
	case models.MediaKategoriBarangFoto{}.TableName():
		if err := media_seller_handle.CreateUbahKategoriBarangFoto(d); err != nil {
			return err
		}
	case models.MediaDistributorDataDokumen{}.TableName():
		if err := media_seller_handle.CreateTambahDistributorDataDokumen(d); err != nil {
			return err
		}
	case models.MediaDistributorDataNPWPFoto{}.TableName():
		if err := media_seller_handle.CreateTambahMediaDistributorDataNPWPFoto(d); err != nil {
			return err
		}
	case models.MediaDistributorDataNIBFoto{}.TableName():
		if err := media_seller_handle.CreateTambahDistributorDataNIBFoto(d); err != nil {
			return err
		}
	case models.MediaDistributorDataSuratKerjasamaDokumen{}.TableName():
		if err := media_seller_handle.CreateTambahDistributorDataDokumen(d); err != nil {
			return err
		}
	case models.MediaBrandDataPerwakilanDokumen{}.TableName():
		if err := media_seller_handle.CreateTambahBrandDataPerwakilanDokumen(d); err != nil {
			return err
		}
	case models.MediaBrandDataSertifikatFoto{}.TableName():
		if err := media_seller_handle.CreateTambahMediaBrandDataSertifikatFoto(d); err != nil {
			return err
		}
	case models.MediaBrandDataNIBFoto{}.TableName():
		if err := media_seller_handle.CreateTambahMediaBrandDataNIBFoto(d); err != nil {
			return err
		}
	case models.MediaBrandDataNPWPFoto{}.TableName():
		if err := media_seller_handle.CreateTambahMediaBrandNPWPFoto(d); err != nil {
			return err
		}
	case models.MediaBrandDataLogoFoto{}.TableName():
		if err := media_seller_handle.CreateTambahMediaBrandDataLogoFoto(d); err != nil {
			return err
		}
	case models.MediaBrandDataSuratKerjasamaDokumen{}.TableName():
		if err := media_seller_handle.CreateTambahBrandDataSuratKerjasamaDokumen(d); err != nil {
			return err
		}
	case models.MediaTransaksiApprovedFoto{}.TableName():
		if err := media_seller_handle.CreateTambahMediaTransaksiApprovedFoto(d); err != nil {
			return err
		}
	case models.MediaTransaksiApprovedVideo{}.TableName():
		if err := media_seller_handle.CreateTambahMediaTransaksiApprovedVideo(d); err != nil {
			return err
		}
	case models.EntitySocialMedia{}.TableName():
		if err := social_media_seller_handle.CreateEngageSocialMediaSeller(d); err != nil {
			return err
		}
	case models.PengirimanEkspedisi{}.TableName():
		if err := transaksi_seller_handle.CreateKirimOrderTransaksiEkspedisi(d); err != nil {
			return err
		}
	case models.Pengiriman{}.TableName():
		if err := transaksi_seller_handle.CreateKirimOrderTransaksiBiasa(d); err != nil {
			return err
		}
	}

	return nil
}
