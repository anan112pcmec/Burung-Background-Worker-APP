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

func SellerUpdateServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](data T) error {
	var d mb_cud_serializer.ParsedDataMessage
	switch v := any(data).(type) {
	case mb_cud_serializer.ConsumeDataJson:
		d = v.Parse()

		switch d.TableName {
		case models.AlamatGudang.TableName(models.AlamatGudang{}):
			if err := alamat_seller_handle.UpdateEditAlamatGudang(d); err != nil {
				return err
			}
		case models.BarangInduk.TableName(models.BarangInduk{}):
			if err := barang_seller_handle.UpdateEditBarangInduk(d); err != nil {
				return err
			}
		case models.KategoriBarang.TableName(models.KategoriBarang{}):
			if err := barang_seller_handle.UpdateEditKategoriBarang(d); err != nil {
				return err
			}
		case models.Komentar.TableName(models.Komentar{}):
			if err := barang_seller_handle.UpdateEditKomentarBarang(d); err != nil {
				return err
			}
		case models.KomentarChild.TableName(models.KomentarChild{}):
			if err := barang_seller_handle.UpdateEditChildKomentar(d); err != nil {
				return err
			}
		case "ValidateUbahPasswordSeller":
			if err := credential_seller_handle.UpdateValidateUbahPasswordSeller(d); err != nil {
				return err
			}
		case "EditRekeningSeller":
			if err := credential_seller_handle.UpdateEditRekeningSeller(d); err != nil {
				return err
			}
		case "SetDefaultRekeningSeller":
			if err := credential_seller_handle.UpdateSetDefaultRekeningSeller(d); err != nil {
				return err
			}
		case models.DiskonProduk.TableName(models.DiskonProduk{}):
			if err := diskon_seller_handle.UpdateEditDiskonProduk(d); err != nil {
				return err
			}
		case models.Etalase.TableName(models.Etalase{}):
			if err := etalase_seller_handle.UpdateEditEtalaseSeller(d); err != nil {
				return err
			}
		case models.DistributorData.TableName(models.DistributorData{}):
			if err := jenis_seller_handle.UpdateEditDataDistributor(d); err != nil {
				return err
			}
		case models.BrandData.TableName(models.BrandData{}):
			if err := jenis_seller_handle.UpdateEditDataBrand(d); err != nil {
				return err
			}
		case models.MediaSellerProfilFoto.TableName(models.MediaSellerProfilFoto{}):
			if err := media_seller_handle.UpdateUbahFotoProfilSeller(d); err != nil {
				return err
			}
		case models.MediaSellerBannerFoto.TableName(models.MediaSellerBannerFoto{}):
			if err := media_seller_handle.UpdateUbahFotoBannerSeller(d); err != nil {
				return err
			}
		case models.MediaEtalaseFoto.TableName(models.MediaEtalaseFoto{}):
			if err := media_seller_handle.UpdateUbahFotoEtalaseSeller(d); err != nil {
				return err
			}
		case models.MediaBarangIndukVideo.TableName(models.MediaBarangIndukVideo{}):
			if err := media_seller_handle.UpdateUbahBarangIndukVideo(d); err != nil {
				return err
			}
		case models.MediaKategoriBarangFoto.TableName(models.MediaKategoriBarangFoto{}):
			if err := media_seller_handle.UpdateUbahKategoriBarangFoto(d); err != nil {
				return err
			}
		case models.MediaDistributorDataDokumen.TableName(models.MediaDistributorDataDokumen{}):
			if err := media_seller_handle.UpdateTambahDistributorDataDokumen(d); err != nil {
				return err
			}
		case models.MediaDistributorDataNPWPFoto.TableName(models.MediaDistributorDataNPWPFoto{}):
			if err := media_seller_handle.UpdateTambahMediaDistributorDataNPWPFoto(d); err != nil {
				return err
			}
		case models.MediaDistributorDataNIBFoto.TableName(models.MediaDistributorDataNIBFoto{}):
			if err := media_seller_handle.UpdateTambahDistributorDataNIBFoto(d); err != nil {
				return err
			}
		case models.MediaDistributorDataSuratKerjasamaDokumen.TableName(models.MediaDistributorDataSuratKerjasamaDokumen{}):
			if err := media_seller_handle.UpdateTambahDistributorDataDokumen(d); err != nil {
				return err
			}
		case models.MediaBrandDataPerwakilanDokumen.TableName(models.MediaBrandDataPerwakilanDokumen{}):
			if err := media_seller_handle.UpdateTambahBrandDataPerwakilanDokumen(d); err != nil {
				return err
			}
		case models.MediaBrandDataPerwakilanDokumen.TableName(models.MediaBrandDataPerwakilanDokumen{}):
			if err := media_seller_handle.UpdateTambahBrandDataPerwakilanDokumen(d); err != nil {
				return err
			}
		case models.MediaBrandDataSertifikatFoto.TableName(models.MediaBrandDataSertifikatFoto{}):
			if err := media_seller_handle.UpdateTambahMediaBrandDataSertifikatFoto(d); err != nil {
				return err
			}
		case models.MediaBrandDataNIBFoto.TableName(models.MediaBrandDataNIBFoto{}):
			if err := media_seller_handle.UpdateTambahMediaBrandDataNIBFoto(d); err != nil {
				return err
			}
		case models.MediaBrandDataNPWPFoto.TableName(models.MediaBrandDataNPWPFoto{}):
			if err := media_seller_handle.UpdateTambahMediaBrandNPWPFoto(d); err != nil {
				return err
			}
		case models.MediaBrandDataLogoFoto.TableName(models.MediaBrandDataLogoFoto{}):
			if err := media_seller_handle.UpdateTambahMediaBrandDataLogoFoto(d); err != nil {
				return err
			}
		case models.MediaBrandDataSuratKerjasamaDokumen.TableName(models.MediaBrandDataSuratKerjasamaDokumen{}):
			if err := media_seller_handle.UpdateTambahBrandDataSuratKerjasamaDokumen(d); err != nil {
				return err
			}

		}

	case mb_cud_serializer.ConsumeDataProto:
		d = v.Parse()

		switch d.TableName {
		case models.AlamatGudang.TableName(models.AlamatGudang{}):
			if err := alamat_seller_handle.UpdateEditAlamatGudang(d); err != nil {
				return err
			}

		}
	}

	return nil
}
