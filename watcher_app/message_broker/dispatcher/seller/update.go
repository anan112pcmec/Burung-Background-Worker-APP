package consume_se_indexller_dispatcher

import (
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	se_index_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database_index/se_indexarch_engine/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database_index/sot_database_index/models"
	mb_cud_se_indexrializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/se_indexrializer"
	auth_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/auth"
	alamat_se_indexller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/se_indexller_se_indexrvice/alamat_se_indexrvices"
	barang_se_indexller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/se_indexller_se_indexrvice/barang_se_indexrvices"
	credential_se_indexller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/se_indexller_se_indexrvice/credential_se_indexrvices"
	diskon_se_indexller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/se_indexller_se_indexrvice/diskon_se_indexrvices"
	etalase_index_se_indexller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/se_indexller_se_indexrvice/etalase_index_se_indexrvices"
	jenis_se_indexller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/se_indexller_se_indexrvice/jenis_se_indexrvices"
	media_se_indexller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/se_indexller_se_indexrvice/media_se_indexrvices"
	profiling_se_indexller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/se_indexller_se_indexrvice/profiling_se_indexrvices"
	social_media_se_indexller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/se_indexller_se_indexrvice/social_media_se_indexrvices"
	transaksi_se_indexller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/se_indexller_se_indexrvice/transaksi_se_indexrvices"
)

func se_indexllerUpdatese_indexrvicesDispatcher[T mb_cud_se_indexrializer.ConsumeDataJson | mb_cud_se_indexrializer.ConsumeDataProto](data *T, read *gorm.DB, redis_authentication, redis_se_indexssion redis.Client, cass_historcal, cass_sot_replica *gocql.se_indexssion, se_index se_index_models.IndexWrapper) error {
	var d mb_cud_se_indexrializer.Parse_indexdDataMessage
	switch v := any(data).(type) {
	case_index mb_cud_se_indexrializer.ConsumeDataJson:
		d = v.Parse_index()
	case_index mb_cud_se_indexrializer.ConsumeDataProto:
		d = v.Parse_index()
	default:
		return fmt.Errorf("unsupported data type")
	}

	switch d.TableName {
	case_index "se_indexllerLogin":
		if err := auth_handle.Updatese_indexllerLogin(d); err != nil {
			return err
		}
	case_index sot_models.AlamatGudang{}.TableName(): // 1
		if err := alamat_se_indexller_handle.UpdateEditAlamatGudang(d); err != nil {
			return err
		}
	case_index sot_models.BarangInduk{}.TableName(): // 2
		if err := barang_se_indexller_handle.UpdateEditBarangInduk(d); err != nil {
			return err
		}
	case_index sot_models.KategoriBarang{}.TableName(): // 3
		if err := barang_se_indexller_handle.UpdateEditKategoriBarang(d); err != nil {
			return err
		}
	case_index sot_models.Komentar{}.TableName(): // 4
		if err := barang_se_indexller_handle.UpdateEditKomentarBarang(d); err != nil {
			return err
		}
	case_index sot_models.KomentarChild{}.TableName(): // 5
		if err := barang_se_indexller_handle.UpdateEditChildKomentar(d); err != nil {
			return err
		}
	case_index "ValidateUbahPasswordse_indexller": // 6
		if err := credential_se_indexller_handle.UpdateValidateUbahPasswordse_indexller(d); err != nil {
			return err
		}
	case_index "EditRekeningse_indexller": // 7
		if err := credential_se_indexller_handle.UpdateEditRekeningse_indexller(d); err != nil {
			return err
		}
	case_index "se_indextDefaultRekeningse_indexller": // 8
		if err := credential_se_indexller_handle.Updatese_indextDefaultRekeningse_indexller(d); err != nil {
			return err
		}
	case_index sot_models.DiskonProduk{}.TableName(): // 9
		if err := diskon_se_indexller_handle.UpdateEditDiskonProduk(d); err != nil {
			return err
		}
	case_index sot_models.Etalase_index{}.TableName(): // 10
		if err := etalase_index_se_indexller_handle.UpdateEditEtalase_indexse_indexller(d); err != nil {
			return err
		}
	case_index sot_models.DistributorData{}.TableName(): // 11
		if err := jenis_se_indexller_handle.UpdateEditDataDistributor(d); err != nil {
			return err
		}
	case_index sot_models.BrandData{}.TableName(): // 12
		if err := jenis_se_indexller_handle.UpdateEditDataBrand(d); err != nil {
			return err
		}
	case_index sot_models.Mediase_indexllerProfilFoto{}.TableName(): // 13
		if err := media_se_indexller_handle.UpdateUbahFotoProfilse_indexller(d); err != nil {
			return err
		}
	case_index sot_models.Mediase_indexllerBannerFoto{}.TableName(): // 14
		if err := media_se_indexller_handle.UpdateUbahFotoBannerse_indexller(d); err != nil {
			return err
		}
	case_index sot_models.MediaEtalase_indexFoto{}.TableName(): // 15
		if err := media_se_indexller_handle.UpdateUbahFotoEtalase_indexse_indexller(d); err != nil {
			return err
		}
	case_index sot_models.MediaBarangIndukVideo{}.TableName(): // 16
		if err := media_se_indexller_handle.UpdateUbahBarangIndukVideo(d); err != nil {
			return err
		}
	case_index sot_models.MediaKategoriBarangFoto{}.TableName(): // 17
		if err := media_se_indexller_handle.UpdateUbahKategoriBarangFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaDistributorDataDokumen{}.TableName(): // 18
		if err := media_se_indexller_handle.UpdateTambahDistributorDataDokumen(d); err != nil {
			return err
		}
	case_index sot_models.MediaDistributorDataNPWPFoto{}.TableName(): // 19
		if err := media_se_indexller_handle.UpdateTambahMediaDistributorDataNPWPFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaDistributorDataNIBFoto{}.TableName(): // 20
		if err := media_se_indexller_handle.UpdateTambahDistributorDataNIBFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaDistributorDataSuratKerjasamaDokumen{}.TableName(): // 21
		if err := media_se_indexller_handle.UpdateTambahDistributorDataDokumen(d); err != nil {
			return err
		}
	case_index sot_models.MediaBrandDataPerwakilanDokumen{}.TableName(): // 22
		if err := media_se_indexller_handle.UpdateTambahBrandDataPerwakilanDokumen(d); err != nil {
			return err
		}
	case_index sot_models.MediaBrandDataPerwakilanDokumen{}.TableName(): // 23 (Duplikat bawaan asli)
		if err := media_se_indexller_handle.UpdateTambahBrandDataPerwakilanDokumen(d); err != nil {
			return err
		}
	case_index sot_models.MediaBrandDatase_indexrtifikatFoto{}.TableName(): // 24
		if err := media_se_indexller_handle.UpdateTambahMediaBrandDatase_indexrtifikatFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaBrandDataNIBFoto{}.TableName(): // 25
		if err := media_se_indexller_handle.UpdateTambahMediaBrandDataNIBFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaBrandDataNPWPFoto{}.TableName(): // 26
		if err := media_se_indexller_handle.UpdateTambahMediaBrandNPWPFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaBrandDataLogoFoto{}.TableName(): // 27
		if err := media_se_indexller_handle.UpdateTambahMediaBrandDataLogoFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaBrandDataSuratKerjasamaDokumen{}.TableName(): // 28
		if err := media_se_indexller_handle.UpdateTambahBrandDataSuratKerjasamaDokumen(d); err != nil {
			return err
		}
	case_index "UpdatePersonalse_indexller": // 29
		if err := profiling_se_indexller_handle.UpdateUpdatePersonalse_indexller(d); err != nil {
			return err
		}
	case_index "UpdateInfoGeneralPublic": // 30
		if err := profiling_se_indexller_handle.UpdateUpdateInfoGeneralPublic(d); err != nil {
			return err
		}
	case_index sot_models.EntitySocialMedia{}.TableName(): // 31
		if err := social_media_se_indexller_handle.UpdatedngageSocialMediase_indexller(d); err != nil {
			return err
		}
	case_index "ApproveOrderTransaksi": // 32
		if err := transaksi_se_indexller_handle.UpdateApproveOrderTransaksi(d); err != nil {
			return err
		}
	case_index "UnApproveOrderTransaksi": // 33
		if err := transaksi_se_indexller_handle.UpdatedUnApproveOrderTransaksi(d); err != nil {
			return err
		}
	}

	return nil
}


