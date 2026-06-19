package consume_seller_dispatcher

import (
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	auth_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/auth"
	alamat_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/alamat_services"
	barang_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/barang_services"
	credential_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/credential_services"
	diskon_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/diskon_services"
	etalase_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/etalase_services"
	jenis_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/jenis_services"
	media_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/media_services"
	profiling_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/profiling_services"
	social_media_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/social_media_services"
	transaksi_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/transaksi_services"
)

func SellerUpdateServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](data *T, read *gorm.DB, redis_authentication, redis_session redis.Client, cass_historcal, cass_sot_replica *gocql.Session, se meilisearch.ServiceManager) error {
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
	case "SellerLogin":
		if err := auth_handle.UpdateSellerLogin(d); err != nil {
			return err
		}
	case sot_models.AlamatGudang{}.TableName(): // 1
		if err := alamat_seller_handle.UpdateEditAlamatGudang(d); err != nil {
			return err
		}
	case sot_models.BarangInduk{}.TableName(): // 2
		if err := barang_seller_handle.UpdateEditBarangInduk(d); err != nil {
			return err
		}
	case sot_models.KategoriBarang{}.TableName(): // 3
		if err := barang_seller_handle.UpdateEditKategoriBarang(d); err != nil {
			return err
		}
	case sot_models.Komentar{}.TableName(): // 4
		if err := barang_seller_handle.UpdateEditKomentarBarang(d); err != nil {
			return err
		}
	case sot_models.KomentarChild{}.TableName(): // 5
		if err := barang_seller_handle.UpdateEditChildKomentar(d); err != nil {
			return err
		}
	case "ValidateUbahPasswordSeller": // 6
		if err := credential_seller_handle.UpdateValidateUbahPasswordSeller(d); err != nil {
			return err
		}
	case "EditRekeningSeller": // 7
		if err := credential_seller_handle.UpdateEditRekeningSeller(d); err != nil {
			return err
		}
	case "SetDefaultRekeningSeller": // 8
		if err := credential_seller_handle.UpdateSetDefaultRekeningSeller(d); err != nil {
			return err
		}
	case sot_models.DiskonProduk{}.TableName(): // 9
		if err := diskon_seller_handle.UpdateEditDiskonProduk(d); err != nil {
			return err
		}
	case sot_models.Etalase{}.TableName(): // 10
		if err := etalase_seller_handle.UpdateEditEtalaseSeller(d); err != nil {
			return err
		}
	case sot_models.DistributorData{}.TableName(): // 11
		if err := jenis_seller_handle.UpdateEditDataDistributor(d); err != nil {
			return err
		}
	case sot_models.BrandData{}.TableName(): // 12
		if err := jenis_seller_handle.UpdateEditDataBrand(d); err != nil {
			return err
		}
	case sot_models.MediaSellerProfilFoto{}.TableName(): // 13
		if err := media_seller_handle.UpdateUbahFotoProfilSeller(d); err != nil {
			return err
		}
	case sot_models.MediaSellerBannerFoto{}.TableName(): // 14
		if err := media_seller_handle.UpdateUbahFotoBannerSeller(d); err != nil {
			return err
		}
	case sot_models.MediaEtalaseFoto{}.TableName(): // 15
		if err := media_seller_handle.UpdateUbahFotoEtalaseSeller(d); err != nil {
			return err
		}
	case sot_models.MediaBarangIndukVideo{}.TableName(): // 16
		if err := media_seller_handle.UpdateUbahBarangIndukVideo(d); err != nil {
			return err
		}
	case sot_models.MediaKategoriBarangFoto{}.TableName(): // 17
		if err := media_seller_handle.UpdateUbahKategoriBarangFoto(d); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataDokumen{}.TableName(): // 18
		if err := media_seller_handle.UpdateTambahDistributorDataDokumen(d); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataNPWPFoto{}.TableName(): // 19
		if err := media_seller_handle.UpdateTambahMediaDistributorDataNPWPFoto(d); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataNIBFoto{}.TableName(): // 20
		if err := media_seller_handle.UpdateTambahDistributorDataNIBFoto(d); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataSuratKerjasamaDokumen{}.TableName(): // 21
		if err := media_seller_handle.UpdateTambahDistributorDataDokumen(d); err != nil {
			return err
		}
	case sot_models.MediaBrandDataPerwakilanDokumen{}.TableName(): // 22
		if err := media_seller_handle.UpdateTambahBrandDataPerwakilanDokumen(d); err != nil {
			return err
		}
	case sot_models.MediaBrandDataPerwakilanDokumen{}.TableName(): // 23 (Duplikat bawaan asli)
		if err := media_seller_handle.UpdateTambahBrandDataPerwakilanDokumen(d); err != nil {
			return err
		}
	case sot_models.MediaBrandDataSertifikatFoto{}.TableName(): // 24
		if err := media_seller_handle.UpdateTambahMediaBrandDataSertifikatFoto(d); err != nil {
			return err
		}
	case sot_models.MediaBrandDataNIBFoto{}.TableName(): // 25
		if err := media_seller_handle.UpdateTambahMediaBrandDataNIBFoto(d); err != nil {
			return err
		}
	case sot_models.MediaBrandDataNPWPFoto{}.TableName(): // 26
		if err := media_seller_handle.UpdateTambahMediaBrandNPWPFoto(d); err != nil {
			return err
		}
	case sot_models.MediaBrandDataLogoFoto{}.TableName(): // 27
		if err := media_seller_handle.UpdateTambahMediaBrandDataLogoFoto(d); err != nil {
			return err
		}
	case sot_models.MediaBrandDataSuratKerjasamaDokumen{}.TableName(): // 28
		if err := media_seller_handle.UpdateTambahBrandDataSuratKerjasamaDokumen(d); err != nil {
			return err
		}
	case "UpdatePersonalSeller": // 29
		if err := profiling_seller_handle.UpdateUpdatePersonalSeller(d); err != nil {
			return err
		}
	case "UpdateInfoGeneralPublic": // 30
		if err := profiling_seller_handle.UpdateUpdateInfoGeneralPublic(d); err != nil {
			return err
		}
	case sot_models.EntitySocialMedia{}.TableName(): // 31
		if err := social_media_seller_handle.UpdatedngageSocialMediaSeller(d); err != nil {
			return err
		}
	case "ApproveOrderTransaksi": // 32
		if err := transaksi_seller_handle.UpdateApproveOrderTransaksi(d); err != nil {
			return err
		}
	case "UnApproveOrderTransaksi": // 33
		if err := transaksi_seller_handle.UpdatedUnApproveOrderTransaksi(d); err != nil {
			return err
		}
	}

	return nil
}
