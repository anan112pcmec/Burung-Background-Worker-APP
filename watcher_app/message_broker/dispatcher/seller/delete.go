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
	alamat_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/alamat_services"
	barang_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/barang_services"
	credential_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/credential_services"
	diskon_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/diskon_services"
	etalase_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/etalase_services"
	jenis_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/jenis_services"
	media_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/media_services"
)

func SellerDeleteServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](ctx context.Context, data *T, read *gorm.DB, redis_authentication, redis_session *redis.Client, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
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
		if err := alamat_seller_handle.DeleteHapusAlamatGudang(d, ctx, cass_historcal, cass_sot_replica, se_index); err != nil {
			return err
		}
	case sot_models.BarangInduk{}.TableName(): // 2
		if err := barang_seller_handle.DeleteHapusBarangInduk(d, ctx, cass_historcal, cass_sot_replica, se_index); err != nil {
			return err
		}
	case sot_models.KategoriBarang{}.TableName(): // 3
		if err := barang_seller_handle.DeleteHapusBarangKategori(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.Komentar{}.TableName(): // 4
		if err := barang_seller_handle.DeleteHapusKomentarBarang(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.KomentarChild{}.TableName(): // 5
		if err := barang_seller_handle.DeleteHapusChildKomentar(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.RekeningSeller{}.TableName(): // 6
		if err := credential_seller_handle.DeleteHapusRekeningSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.DiskonProduk{}.TableName(): // 7
		if err := diskon_seller_handle.DeleteHapusDiskonProduk(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.BarangDiDiskon{}.TableName(): // 8
		if err := diskon_seller_handle.DeleteHapusDiskonPadaBarang(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.Etalase{}.TableName(): // 9
		if err := etalase_seller_handle.DeleteHapusEtalaseSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.BarangKeEtalase{}.TableName(): // 10
		if err := etalase_seller_handle.DeleteHapusBarangDariEtalase(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.DistributorData{}.TableName(): // 11
		if err := jenis_seller_handle.DeleteHapusDataDistributor(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.BrandData{}.TableName(): // 12
		if err := jenis_seller_handle.DeleteHapusDataBrand(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaSellerProfilFoto{}.TableName(): // 13
		if err := media_seller_handle.DeleteHapusFotoProfilSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaSellerBannerFoto{}.TableName(): // 14
		if err := media_seller_handle.DeleteHapusFotoBannerSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaSellerBannerFoto{}.TableName(): // 15 (Duplikat bawaan asli)
		if err := media_seller_handle.DeleteHapusFotoBannerSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaEtalaseFoto{}.TableName(): // 16
		if err := media_seller_handle.DeleteHapusFotoEtalaseSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBarangIndukFoto{}.TableName(): // 17
		if err := media_seller_handle.DeleteHapusMediaBarangIndukFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBarangIndukVideo{}.TableName(): // 18
		if err := media_seller_handle.DeleteHapusBarangIndukVideo(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaKategoriBarangFoto{}.TableName(): // 19
		if err := media_seller_handle.DeleteHapusKategoriBarangFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataDokumen{}.TableName(): // 20
		if err := media_seller_handle.DeleteHapusMediaDistributorDataDokumen(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataNPWPFoto{}.TableName(): // 21
		if err := media_seller_handle.DeleteHapusMediaDistributorDataNPWPFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataNIBFoto{}.TableName(): // 22
		if err := media_seller_handle.DeleteHapusDistributorDataNIBFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataSuratKerjasamaDokumen{}.TableName(): // 23
		if err := media_seller_handle.DeleteHapusMediaDistributorDataDokumen(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataPerwakilanDokumen{}.TableName(): // 24
		if err := media_seller_handle.DeleteHapusBrandDataPerwakilanDokumen(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataPerwakilanDokumen{}.TableName(): // 25 (Duplikat bawaan asli)
		if err := media_seller_handle.DeleteHapusBrandDataPerwakilanDokumen(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataSertifikatFoto{}.TableName(): // 26
		if err := media_seller_handle.DeleteHapusMediaBrandDataSertifikatFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataNIBFoto{}.TableName(): // 27
		if err := media_seller_handle.DeleteHapusMediaBrandDataNIBFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataNPWPFoto{}.TableName(): // 28
		if err := media_seller_handle.DeleteHapusMediaBrandNPWPFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataLogoFoto{}.TableName(): // 29
		if err := media_seller_handle.DeleteHapusMediaBrandDataLogoFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataSuratKerjasamaDokumen{}.TableName(): // 30
		if err := media_seller_handle.DeleteHapusBrandDataSuratKerjasamaDokumen(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	}

	return nil
}
