package consume_kurir_dispatcher

import (
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	se_index_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database_index/se_indexarch_engine/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database_index/sot_database_index/models"
	mb_cud_se_indexrializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/se_indexrializer"
	auth_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/auth"
	alamat_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/kurir_se_indexrvice/alamat_se_indexrvices"
	informasi_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/kurir_se_indexrvice/informasi_se_indexrvices"
	media_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/kurir_se_indexrvice/media_se_indexrvices"
	pengiriman_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/kurir_se_indexrvice/pengiriman_se_indexrvices"
	profiling_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/kurir_se_indexrvice/profiling_se_indexrvices"
	rekening_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/kurir_se_indexrvice/rekening_se_indexrvices"
	social_media_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/kurir_se_indexrvice/social_media_se_indexrvices"
)

func KurirUpdatese_indexrvicesDispatcher[T mb_cud_se_indexrializer.ConsumeDataJson | mb_cud_se_indexrializer.ConsumeDataProto](data *T, read *gorm.DB, redis_authentication, redis_se_indexssion redis.Client, cass_historcal, cass_sot_replica *gocql.se_indexssion, se_index se_index_models.IndexWrapper) error {
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
	case_index "KurirLogin":
		if err := auth_handle.UpdateKurirLogin(d); err != nil {
			return err
		}
	case_index sot_models.AlamatKurir{}.TableName():
		if err := alamat_kurir_handle.UpdatedEditAlamatKurir(d); err != nil {
			return err
		}
	case_index sot_models.InformasiKendaraanKurir{}.TableName():
		if err := informasi_kurir_handle.UpdateEditInformasiKendaraan(d); err != nil {
			return err
		}
	case_index sot_models.InformasiKurir{}.TableName():
		if err := informasi_kurir_handle.UpdateEditInformasiKurir(d); err != nil {
			return err
		}
	case_index sot_models.MediaKurirProfilFoto{}.TableName():
		if err := media_kurir_handle.UpdateUbahKurirProfilFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaInformasiKendaraanKurirKendaraanFoto{}.TableName():
		if err := media_kurir_handle.UpdateTambahMediaInformasiKendaraanKurirKendaraanFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaInformasiKendaraanKurirBPKBFoto{}.TableName():
		if err := media_kurir_handle.UpdateTambahInformasiKendaraanKurirBPKBFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaInformasiKendaraanKurirSTNKFoto{}.TableName():
		if err := media_kurir_handle.UpdateTambahInformasiKendaraanKurirSTNKFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaInformasiKurirKTPFoto{}.TableName():
		if err := media_kurir_handle.UpdateTambahMediaInformasiKurirKTPFoto(d); err != nil {
			return err
		}
	case_index "kurirUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateAktifkanBidKurir(d); err != nil {
			return err
		}
	case_index "bidKurirDataUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateUpdatePosisiBidKurir(d); err != nil {
			return err
		}
	case_index "pengirimanUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateAmbilPengirimanNonEksManualRegulerIIpengirimanUpdatedPublish(d); err != nil {
			return err
		}
	case_index "bidKurirDataAmbilPengirimanUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataAmbilPengirimanUpdatedPublish(d); err != nil {
			return err
		}
	case_index "bidKurirDataStatusUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataStatusUpdatedPublish(d); err != nil {
			return err
		}
	case_index "pengirimanEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateAmbilPengirimanEksManualRegulerIIpengirimanEksUpdatedPublish(d); err != nil {
			return err
		}
	case_index "bidKurirDataAmbilPengirimanEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataUpdatedPublish(d); err != nil {
			return err
		}
	case_index "bidKurirDataAmbilPengirimanEksStatusUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataAmbilPengirimanEksStatusUpdatedPublish(d); err != nil {
			return err
		}
	case_index "schedulerEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateLockSiapAntarBidKurirIIEksScheduler(d); err != nil {
			return err
		}
	case_index "schedulerNonEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateLockSiapAntarBidKurirIINonEksScheduler(d); err != nil {
			return err
		}
	case_index "bidKurirDataLockSiapAntarUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateLockSiapAntarBidKurirIIbidKurirDataLockSiapAntarUpdatedPublish(d); err != nil {
			return err
		}
	case_index "kurirLockSiapAntarUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateLockSiapAntarBidKurirIIkurirLockSiapAntarUpdatedPublish(d); err != nil {
			return err
		}
	case_index "schedulerPickedUpNonEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdatePickedUpPengirimanNonEksIIschedulerPickedUpNonEksUpdatedPublish(d); err != nil {
			return err
		}
	case_index "pengirimanPickedUpNonEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdatedPickedUpPengirimanNonEksIIpengirimanPickedUpNonEksUpdatedPublish(d); err != nil {
			return err
		}
	case_index "transaksiPickedUpNonEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdatedPickedUpPengirimanNonEksIItransaksiPickedUpNonEksUpdatedPublish(d); err != nil {
			return err
		}
	case_index "bidKurirPengirimanNonEksSchedulerUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateKirimPengirimanNonEksIIbidKurirPengirimanNonEksSchedulerUpdatedPublish(d); err != nil {
			return err
		}
	case_index "pengirimanPengirimanUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateKirimPengirimanNonEksIIpengirimanPengirimanUpdatedPublish(d); err != nil {
			return err
		}
	case_index "jejakpengirimanPengirimanUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateKirimPengirimanNonEksIIjejakpengirimanPengirimanUpdatedPublish(d); err != nil {
			return err
		}
	case_index "jejakPengirimanUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateUpdateInformasiPerjalananPengirimanNonEks(d); err != nil {
			return err
		}
	case_index "pengirimanSampaiUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanNonEksIIpengirimanSampaiUpdatedPublish(d); err != nil {
			return err
		}
	case_index "bidKurirDataSampaiUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanNonEksIIbidKurirDataSampaiUpdatedPublish(d); err != nil {
			return err
		}
	case_index "jejakPengirimanSampaiUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanNonEksIIjejakPengirimanSampaiUpdatedPublish(d); err != nil {
			return err
		}
	case_index "transaksiSampaiUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanNonEksIItransaksiSampaiUpdatedPublish(d); err != nil {
			return err
		}
	case_index "schedulerEksPickedUpUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdatePickedUpPengirimanEksIIschedulerEksPickedUpUpdatedPublish(d); err != nil {
			return err
		}
	case_index "pengirimanEksPickedUpUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdatePickedUpPengirimanEksIIpengirimanEksPickedUpUpdatedPublish(d); err != nil {
			return err
		}
	case_index "transaksiPickedUpUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdatePickedUpPengirimanEksIItransaksiPickedUpUpdatedPublish(d); err != nil {
			return err
		}
	case_index "schedulerPengirimanUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateKirimPengirimanEksIIschedulerPengirimanUpdatedPublish(d); err != nil {
			return err
		}
	case_index "pengirimanPengirimanEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateKirimPengirimanEksIIpengirimanPengirimanEksUpdatedPublish(d); err != nil {
			return err
		}
	case_index "jejakPengirimanPengirimanEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateKirimPengirimanNonEksIIjejakpengirimanPengirimanUpdatedPublish(d); err != nil {
			return err
		}
	case_index "jejakPengirimanEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateInformasiPerjalananPengirimanEks(d); err != nil {
			return err
		}
	case_index "pengirimanSampaiEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanEksIIpengirimanSampaiEksUpdatedPublish(d); err != nil {
			return err
		}
	case_index "bidKurirDataEksSampaiUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanEksIIbidKurirDataEksSampaiUpdatedPublish(d); err != nil {
			return err
		}
	case_index "jejakPengirimanEksSampaiUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanEksIIjejakPengirimanEksSampaiUpdatedPublish(d); err != nil {
			return err
		}
	case_index "transaksiSampaiEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanEksIItransaksiSampaiEksUpdatedPublish(d); err != nil {
			return err
		}
	case_index "kurirUpdatedSampaiEksPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanEksIIkurirUpdatedSampaiEksPublish(d); err != nil {
			return err
		}
	case_index "kurirNonaktifkanBidUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateNonaktifkanBidKurirIIkurirNonaktifkanBidUpdatedPublish(d); err != nil {
			return err
		}
	case_index "kurirDataPersonalProfilingUpdatedPublish":
		if err := profiling_kurir_handle.UpdatePersonalProfilingKurir(d); err != nil {
			return err
		}
	case_index "kurirDataGeneralProfilingUpdatedPublish":
		if err := profiling_kurir_handle.UpdateGeneralProfilingKurir(d); err != nil {
			return err
		}
	case_index sot_models.RekeningKurir{}.TableName():
		if err := rekening_kurir_handle.UpdateEditRekeningKurir(d); err != nil {
			return err
		}
	case_index sot_models.EntitySocialMedia{}.TableName():
		if err := social_media_kurir_handle.UpdateEngagementSocialMediaKurir(d); err != nil {
			return err
		}
	}
	return nil

}


