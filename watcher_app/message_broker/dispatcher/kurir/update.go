package consume_kurir_dispatcher

import (
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	auth_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/auth"
	alamat_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/alamat_services"
	informasi_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/informasi_services"
	media_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/media_services"
	pengiriman_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/pengiriman_services"
	profiling_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/profiling_services"
	rekening_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/rekening_services"
	social_media_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/social_media_services"
)

func KurirUpdateServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](data *T, read *gorm.DB, redis_authentication, redis_session redis.Client, cass_historcal, cass_sot_replica *gocql.Session, se meilisearch.ServiceManager) error {
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
	case "KurirLogin":
		if err := auth_handle.UpdateKurirLogin(d); err != nil {
			return err
		}
	case sot_models.AlamatKurir{}.TableName():
		if err := alamat_kurir_handle.UpdatedEditAlamatKurir(d); err != nil {
			return err
		}
	case sot_models.InformasiKendaraanKurir{}.TableName():
		if err := informasi_kurir_handle.UpdateEditInformasiKendaraan(d); err != nil {
			return err
		}
	case sot_models.InformasiKurir{}.TableName():
		if err := informasi_kurir_handle.UpdateEditInformasiKurir(d); err != nil {
			return err
		}
	case sot_models.MediaKurirProfilFoto{}.TableName():
		if err := media_kurir_handle.UpdateUbahKurirProfilFoto(d); err != nil {
			return err
		}
	case sot_models.MediaInformasiKendaraanKurirKendaraanFoto{}.TableName():
		if err := media_kurir_handle.UpdateTambahMediaInformasiKendaraanKurirKendaraanFoto(d); err != nil {
			return err
		}
	case sot_models.MediaInformasiKendaraanKurirBPKBFoto{}.TableName():
		if err := media_kurir_handle.UpdateTambahInformasiKendaraanKurirBPKBFoto(d); err != nil {
			return err
		}
	case sot_models.MediaInformasiKendaraanKurirSTNKFoto{}.TableName():
		if err := media_kurir_handle.UpdateTambahInformasiKendaraanKurirSTNKFoto(d); err != nil {
			return err
		}
	case sot_models.MediaInformasiKurirKTPFoto{}.TableName():
		if err := media_kurir_handle.UpdateTambahMediaInformasiKurirKTPFoto(d); err != nil {
			return err
		}
	case "kurirUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateAktifkanBidKurir(d); err != nil {
			return err
		}
	case "bidKurirDataUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateUpdatePosisiBidKurir(d); err != nil {
			return err
		}
	case "pengirimanUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateAmbilPengirimanNonEksManualRegulerIIpengirimanUpdatedPublish(d); err != nil {
			return err
		}
	case "bidKurirDataAmbilPengirimanUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataAmbilPengirimanUpdatedPublish(d); err != nil {
			return err
		}
	case "bidKurirDataStatusUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataStatusUpdatedPublish(d); err != nil {
			return err
		}
	case "pengirimanEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateAmbilPengirimanEksManualRegulerIIpengirimanEksUpdatedPublish(d); err != nil {
			return err
		}
	case "bidKurirDataAmbilPengirimanEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataUpdatedPublish(d); err != nil {
			return err
		}
	case "bidKurirDataAmbilPengirimanEksStatusUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataAmbilPengirimanEksStatusUpdatedPublish(d); err != nil {
			return err
		}
	case "schedulerEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateLockSiapAntarBidKurirIIEksScheduler(d); err != nil {
			return err
		}
	case "schedulerNonEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateLockSiapAntarBidKurirIINonEksScheduler(d); err != nil {
			return err
		}
	case "bidKurirDataLockSiapAntarUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateLockSiapAntarBidKurirIIbidKurirDataLockSiapAntarUpdatedPublish(d); err != nil {
			return err
		}
	case "kurirLockSiapAntarUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateLockSiapAntarBidKurirIIkurirLockSiapAntarUpdatedPublish(d); err != nil {
			return err
		}
	case "schedulerPickedUpNonEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdatePickedUpPengirimanNonEksIIschedulerPickedUpNonEksUpdatedPublish(d); err != nil {
			return err
		}
	case "pengirimanPickedUpNonEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdatedPickedUpPengirimanNonEksIIpengirimanPickedUpNonEksUpdatedPublish(d); err != nil {
			return err
		}
	case "transaksiPickedUpNonEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdatedPickedUpPengirimanNonEksIItransaksiPickedUpNonEksUpdatedPublish(d); err != nil {
			return err
		}
	case "bidKurirPengirimanNonEksSchedulerUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateKirimPengirimanNonEksIIbidKurirPengirimanNonEksSchedulerUpdatedPublish(d); err != nil {
			return err
		}
	case "pengirimanPengirimanUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateKirimPengirimanNonEksIIpengirimanPengirimanUpdatedPublish(d); err != nil {
			return err
		}
	case "jejakpengirimanPengirimanUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateKirimPengirimanNonEksIIjejakpengirimanPengirimanUpdatedPublish(d); err != nil {
			return err
		}
	case "jejakPengirimanUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateUpdateInformasiPerjalananPengirimanNonEks(d); err != nil {
			return err
		}
	case "pengirimanSampaiUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanNonEksIIpengirimanSampaiUpdatedPublish(d); err != nil {
			return err
		}
	case "bidKurirDataSampaiUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanNonEksIIbidKurirDataSampaiUpdatedPublish(d); err != nil {
			return err
		}
	case "jejakPengirimanSampaiUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanNonEksIIjejakPengirimanSampaiUpdatedPublish(d); err != nil {
			return err
		}
	case "transaksiSampaiUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanNonEksIItransaksiSampaiUpdatedPublish(d); err != nil {
			return err
		}
	case "schedulerEksPickedUpUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdatePickedUpPengirimanEksIIschedulerEksPickedUpUpdatedPublish(d); err != nil {
			return err
		}
	case "pengirimanEksPickedUpUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdatePickedUpPengirimanEksIIpengirimanEksPickedUpUpdatedPublish(d); err != nil {
			return err
		}
	case "transaksiPickedUpUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdatePickedUpPengirimanEksIItransaksiPickedUpUpdatedPublish(d); err != nil {
			return err
		}
	case "schedulerPengirimanUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateKirimPengirimanEksIIschedulerPengirimanUpdatedPublish(d); err != nil {
			return err
		}
	case "pengirimanPengirimanEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateKirimPengirimanEksIIpengirimanPengirimanEksUpdatedPublish(d); err != nil {
			return err
		}
	case "jejakPengirimanPengirimanEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateKirimPengirimanNonEksIIjejakpengirimanPengirimanUpdatedPublish(d); err != nil {
			return err
		}
	case "jejakPengirimanEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateInformasiPerjalananPengirimanEks(d); err != nil {
			return err
		}
	case "pengirimanSampaiEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanEksIIpengirimanSampaiEksUpdatedPublish(d); err != nil {
			return err
		}
	case "bidKurirDataEksSampaiUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanEksIIbidKurirDataEksSampaiUpdatedPublish(d); err != nil {
			return err
		}
	case "jejakPengirimanEksSampaiUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanEksIIjejakPengirimanEksSampaiUpdatedPublish(d); err != nil {
			return err
		}
	case "transaksiSampaiEksUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanEksIItransaksiSampaiEksUpdatedPublish(d); err != nil {
			return err
		}
	case "kurirUpdatedSampaiEksPublish":
		if err := pengiriman_kurir_handle.UpdateSampaiPengirimanEksIIkurirUpdatedSampaiEksPublish(d); err != nil {
			return err
		}
	case "kurirNonaktifkanBidUpdatedPublish":
		if err := pengiriman_kurir_handle.UpdateNonaktifkanBidKurirIIkurirNonaktifkanBidUpdatedPublish(d); err != nil {
			return err
		}
	case "kurirDataPersonalProfilingUpdatedPublish":
		if err := profiling_kurir_handle.UpdatePersonalProfilingKurir(d); err != nil {
			return err
		}
	case "kurirDataGeneralProfilingUpdatedPublish":
		if err := profiling_kurir_handle.UpdateGeneralProfilingKurir(d); err != nil {
			return err
		}
	case sot_models.RekeningKurir{}.TableName():
		if err := rekening_kurir_handle.UpdateEditRekeningKurir(d); err != nil {
			return err
		}
	case sot_models.EntitySocialMedia{}.TableName():
		if err := social_media_kurir_handle.UpdateEngagementSocialMediaKurir(d); err != nil {
			return err
		}
	}
	return nil

}
