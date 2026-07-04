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
	social_media_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/social_media_services"
	transaksi_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/transaksi_services"
)

func SellerCreateServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](ctx context.Context, data *T, read *gorm.DB, redis_authentication, redis_session *redis.Client, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
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
	case sot_models.Seller{}.TableName():
		if err := auth_handle.CreateValidateSellerRegistration(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.AlamatGudang{}.TableName():
		if err := alamat_seller_handle.CreateTambahAlamatGudang(d, ctx, cass_historcal, cass_sot_replica, se_index); err != nil {
			return err
		}
	case sot_models.BarangInduk{}.TableName():
		if err := barang_seller_handle.CreateMasukanBarangInduk(d, ctx, read, cass_historcal, cass_sot_replica, se_index); err != nil {
			return err
		}
	case sot_models.KategoriBarang{}.TableName():
		if err := barang_seller_handle.CreateMasukanKategoriBarang(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.Komentar{}.TableName():
		if err := barang_seller_handle.CreateMasukanKomentarBarang(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.KomentarChild{}.TableName():
		if err := barang_seller_handle.CreateMasukanChildKomentar(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.RekeningSeller{}.TableName():
		if err := credential_seller_handle.CreateTambahRekeningSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.DiskonProduk{}.TableName():
		if err := diskon_seller_handle.CreateTambahDiskonProduk(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.BarangDiDiskon{}.TableName():
		if err := diskon_seller_handle.CreateTetapkanDiskonPadaBarang(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.Etalase{}.TableName():
		if err := etalase_seller_handle.CreateTambahEtalaseSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.BarangKeEtalase{}.TableName():
		if err := etalase_seller_handle.CreateTambahkanBarangKeEtalase(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.DistributorData{}.TableName():
		if err := jenis_seller_handle.CreateMasukanDataDistributor(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.BrandData{}.TableName():
		if err := jenis_seller_handle.CreateMasukanDataBrand(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaSellerProfilFoto{}.TableName():
		if err := media_seller_handle.CreateTambahFotoProfilSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaSellerBannerFoto{}.TableName():
		if err := media_seller_handle.CreateTambahFotoBannerSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaSellerTokoFisikFoto{}.TableName():
		if err := media_seller_handle.CreateTambahkanFotoTokoFisikSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaEtalaseFoto{}.TableName():
		if err := media_seller_handle.CreateTambahFotoEtalaseSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBarangIndukFoto{}.TableName():
		if err := media_seller_handle.CreateTambahkanMediaBarangIndukFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBarangIndukVideo{}.TableName():
		if err := media_seller_handle.CreateTambahBarangIndukVideo(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaKategoriBarangFoto{}.TableName():
		if err := media_seller_handle.CreateTambahKategoriBarangFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataDokumen{}.TableName():
		if err := media_seller_handle.CreateTambahDistributorDataDokumen(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataNPWPFoto{}.TableName():
		if err := media_seller_handle.CreateTambahMediaDistributorDataNPWPFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataNIBFoto{}.TableName():
		if err := media_seller_handle.CreateTambahDistributorDataNIBFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaDistributorDataSuratKerjasamaDokumen{}.TableName():
		if err := media_seller_handle.CreateTambahDistributorDataDokumen(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataPerwakilanDokumen{}.TableName():
		if err := media_seller_handle.CreateTambahBrandDataPerwakilanDokumen(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataSertifikatFoto{}.TableName():
		if err := media_seller_handle.CreateTambahMediaBrandDataSertifikatFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataNIBFoto{}.TableName():
		if err := media_seller_handle.CreateTambahMediaBrandDataNIBFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataNPWPFoto{}.TableName():
		if err := media_seller_handle.CreateTambahMediaBrandNPWPFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataLogoFoto{}.TableName():
		if err := media_seller_handle.CreateTambahMediaBrandDataLogoFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaBrandDataSuratKerjasamaDokumen{}.TableName():
		if err := media_seller_handle.CreateTambahBrandDataSuratKerjasamaDokumen(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaTransaksiApprovedFoto{}.TableName():
		if err := media_seller_handle.CreateTambahMediaTransaksiApprovedFoto(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaTransaksiApprovedVideo{}.TableName():
		if err := media_seller_handle.CreateTambahMediaTransaksiApprovedVideo(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.EntitySocialMedia{}.TableName():
		if err := social_media_seller_handle.CreateEngageSocialMediaSeller(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.PengirimanEkspedisi{}.TableName():
		if err := transaksi_seller_handle.CreateKirimOrderTransaksiEkspedisi(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.Pengiriman{}.TableName():
		if err := transaksi_seller_handle.CreateKirimOrderTransaksiBiasa(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	}

	return nil
}
