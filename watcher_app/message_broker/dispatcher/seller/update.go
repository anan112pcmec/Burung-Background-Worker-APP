package consume_seller_dispatcher

import (
	"context"
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	se_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/search_engine/models"
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

func SellerUpdateServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](ctx context.Context, data *T, read *gorm.DB, redis_authentication, redis_session *redis.Client, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	var d mb_cud_serializer.ParsedDataMessage
	switch v := any(data).(type) {
	case *mb_cud_serializer.ConsumeDataJson:
		d = v.Parse()
	case *mb_cud_serializer.ConsumeDataProto:
		d = v.Parse()
	default:
		return fmt.Errorf("unsupported data type")
	}

	switch d.TableName {
	case "SellerLogin":
		if err := auth_handle.UpdateSellerLogin(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.AlamatGudang{}.TableName(): // 1
		if err := alamat_seller_handle.UpdateEditAlamatGudang(d, ctx, cass_historcal, cass_sot_replica, se_index); err != nil {
			return err
		}
	case sot_models.BarangInduk{}.TableName(): // 2
		if err := barang_seller_handle.UpdateEditBarangInduk(d, ctx, cass_historcal, cass_sot_replica, se_index); err != nil {
			return err
		}
	case sot_models.KategoriBarang{}.TableName(): // 3
		if err := barang_seller_handle.UpdateEditKategoriBarang(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.Komentar{}.TableName(): // 4
		if err := barang_seller_handle.UpdateEditKomentarBarang(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "DownStokBarangInduk":
		if err := barang_seller_handle.UpdateDownStokBarangInduk(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "DownKategoriBarang":
		if err := barang_seller_handle.UpdateDownKategoriBarang(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.KomentarChild{}.TableName(): // 5
		if err := barang_seller_handle.UpdateEditChildKomentar(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "UbahHargaKategoriBarang":
		if err := barang_seller_handle.UpdateUbahHargaKategoriBarang(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "ValidateUbahPasswordSeller": // 6
		if err := credential_seller_handle.UpdateValidateUbahPasswordSeller(d, ctx, cass_historcal, cass_sot_replica, se_index, redis_session); err != nil {
			return err
		}
	case "EditRekeningSeller": // 7
		if err := credential_seller_handle.UpdateEditRekeningSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "SetDefaultRekeningSeller": // 8
		if err := credential_seller_handle.UpdateSetDefaultRekeningSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.DiskonProduk{}.TableName(): // 9
		if err := diskon_seller_handle.UpdateEditDiskonProduk(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.Etalase{}.TableName(): // 10
		if err := etalase_seller_handle.UpdateEditEtalaseSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.DistributorData{}.TableName(): // 11
		if err := jenis_seller_handle.UpdateEditDataDistributor(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.BrandData{}.TableName(): // 12
		if err := jenis_seller_handle.UpdateEditDataBrand(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaSellerProfilFoto{}.TableName(): // 13
		if err := media_seller_handle.UpdateUbahFotoProfilSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaSellerBannerFoto{}.TableName(): // 14
		if err := media_seller_handle.UpdateUbahFotoBannerSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaEtalaseFoto{}.TableName(): // 15
		if err := media_seller_handle.UpdateUbahFotoEtalaseSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBarangIndukVideo{}.TableName(): // 16
		if err := media_seller_handle.UpdateUbahBarangIndukVideo(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaKategoriBarangFoto{}.TableName(): // 17
		if err := media_seller_handle.UpdateUbahKategoriBarangFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataDokumen{}.TableName(): // 18
		if err := media_seller_handle.UpdateUbahDistributorDataDokumen(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataNPWPFoto{}.TableName(): // 19
		if err := media_seller_handle.UpdateUbahMediaDistributorDataNPWPFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataNIBFoto{}.TableName(): // 20
		if err := media_seller_handle.UpdateUbahDistributorDataNIBFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataSuratKerjasamaDokumen{}.TableName(): // 21
		if err := media_seller_handle.UpdateUbahDistributorDataDokumen(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataPerwakilanDokumen{}.TableName(): // 22
		if err := media_seller_handle.UpdateUbahBrandDataPerwakilanDokumen(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataPerwakilanDokumen{}.TableName(): // 23 (Duplikat bawaan asli)
		if err := media_seller_handle.UpdateUbahBrandDataPerwakilanDokumen(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataSertifikatFoto{}.TableName(): // 24
		if err := media_seller_handle.UpdateUbahMediaBrandDataSertifikatFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataNIBFoto{}.TableName(): // 25
		if err := media_seller_handle.UpdateUbahMediaBrandDataNIBFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataNPWPFoto{}.TableName(): // 26
		if err := media_seller_handle.UpdateUbahMediaBrandNPWPFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataLogoFoto{}.TableName(): // 27
		if err := media_seller_handle.UpdateUbahMediaBrandDataLogoFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataSuratKerjasamaDokumen{}.TableName(): // 28
		if err := media_seller_handle.UpdateUbahBrandDataSuratKerjasamaDokumen(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "UpdatePersonalSeller": // 29
		if err := profiling_seller_handle.UpdateUpdatePersonalSeller(d, ctx, cass_historcal, cass_sot_replica, se_index, redis_session); err != nil {
			return err
		}
	case "UpdateInfoGeneralPublic": // 30
		if err := profiling_seller_handle.UpdateUpdateInfoGeneralPublic(d, ctx, cass_historcal, cass_sot_replica, se_index, redis_session); err != nil {
			return err
		}
	case sot_models.EntitySocialMedia{}.TableName(): // 31
		if err := social_media_seller_handle.UpdateEngageSocialMediaSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "ApproveOrderTransaksi": // 32
		if err := transaksi_seller_handle.UpdateApproveOrderTransaksi(d, ctx, cass_historcal, cass_sot_replica, se_index); err != nil {
			return err
		}
	case "UnApproveOrderTransaksi": // 33
		if err := transaksi_seller_handle.UpdateApproveOrderTransaksi(d, ctx, cass_historcal, cass_sot_replica, se_index); err != nil {
			return err
		}
	}

	return nil
}
