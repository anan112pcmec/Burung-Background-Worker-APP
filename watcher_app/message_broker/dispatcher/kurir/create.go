package consume_kurir_dispatcher

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
	alamat_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/alamat_services"
	informasi_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/informasi_services"
	media_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/media_services"
	pengiriman_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/pengiriman_services"
	rekening_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/rekening_services"
	social_media_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/social_media_services"
)

func KurirCreateServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](ctx context.Context, data *T, read *gorm.DB, redis_authentication, redis_session *redis.Client, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
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
	case sot_models.Kurir{}.TableName():
		if err := auth_handle.CreateValidateKurirRegistration(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.AlamatKurir{}.TableName():
		if err := alamat_kurir_handle.CreateMasukanAlamatKurir(d, ctx, read, cass_historcal, cass_sot_replica, se_index); err != nil {
			return err
		}
	case sot_models.InformasiKendaraanKurir{}.TableName():
		if err := informasi_kurir_handle.CreateAjukanInformasiKendaraan(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.InformasiKurir{}.TableName():
		if err := informasi_kurir_handle.CreateAjukanInformasiKurir(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaInformasiKendaraanKurirKendaraanFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahMediaInformasiKendaraanKurirKendaraanFoto(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaInformasiKendaraanKurirBPKBFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahInformasiKendaraanKurirBPKBFoto(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaInformasiKendaraanKurirSTNKFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahInformasiKendaraanKurirSTNKFoto(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaInformasiKurirKTPFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahMediaInformasiKurirKTPFoto(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaPengirimanPickedUpFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahMediaPengirimanPickedUpFoto(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaPengirimanSampaiFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahMediaPengirimanSampaiFoto(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaPengirimanEkspedisiPickedUpFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahMediaPengirimanEkspedisiPickedUpFoto(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.MediaPengirimanEkspedisiSampaiAgentFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahMediaPengirimanEkspedisiSampaiAgentFoto(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "bidKurirDataCreatePublish":
		if err := pengiriman_kurir_handle.CreateAktifkanBidKurir(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "bidKurirNonEksSchedulerCreatePublish":
		if err := pengiriman_kurir_handle.CreateAmbilPengirimanNonEksManualRegulerIIBidKurirNonEksSchedulerCreatePublish(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "bidKurirEksSchedulerCreatePublish":
		if err := pengiriman_kurir_handle.CreateAmbilPengirimanEksManualRegulerIIbidKurirEksSchedulerCreatePublish(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "payOutSellerCreatePublish":
		if err := pengiriman_kurir_handle.CreateSampaiPengirimanNonEksIIpayOutSellerCreatePublish(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "payOutKurirCreatePublish":
		if err := pengiriman_kurir_handle.CreateSampaiPengirimanNonEksIIpayOutKurirCreatePublish(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "jejakPengirimanEksCreatePublish":
		if err := pengiriman_kurir_handle.CreatePickedUpPengirimanEksIIjejakPengirimanEksCreatePublish(d, ctx, read, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case "payOutKurirEksCreatePublish":
		if err := pengiriman_kurir_handle.CreateSampaiPengirimanEksIIpayOutKurirEksCreatePublish(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.RekeningKurir{}.TableName():
		if err := rekening_kurir_handle.CreateMasukanRekeningKurir(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	case sot_models.EntitySocialMedia{}.TableName():
		if err := social_media_kurir_handle.CreateEngagementSocialMediaKurir(d, ctx, cass_historcal, cass_sot_replica); err != nil {
			return err
		}
	}

	return nil
}
