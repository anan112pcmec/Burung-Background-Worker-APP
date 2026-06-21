package consume_kurir_dispatcher

import (
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	se_index_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database_index/se_indexarch_engine/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database_index/sot_database_index/models"
	mb_cud_se_indexrializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/se_indexrializer"
	alamat_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/kurir_se_indexrvice/alamat_se_indexrvices"
	media_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/kurir_se_indexrvice/media_se_indexrvices"
	pengiriman_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/kurir_se_indexrvice/pengiriman_se_indexrvices"
	rekening_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/se_indexrvice_handle/kurir_se_indexrvice/rekening_se_indexrvices"
)

func KurirDeletese_indexrvicesDispatcher[T mb_cud_se_indexrializer.ConsumeDataJson | mb_cud_se_indexrializer.ConsumeDataProto](data *T, read *gorm.DB, redis_authentication, redis_se_indexssion redis.Client, cass_historcal, cass_sot_replica *gocql.se_indexssion, se_index se_index_models.IndexWrapper) error {
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
	case_index sot_models.AlamatKurir{}.TableName():
		if err := alamat_kurir_handle.DeleteHapusAlamatKurir(d); err != nil {
			return err
		}
	case_index sot_models.MediaKurirProfilFoto{}.TableName():
		if err := media_kurir_handle.DeleteHapusKurirProfilFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaInformasiKendaraanKurirKendaraanFoto{}.TableName():
		if err := media_kurir_handle.DeleteHapusMediaInformasiKendaraanKurirKendaraanFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaInformasiKendaraanKurirBPKBFoto{}.TableName():
		if err := media_kurir_handle.DeleteHapusInformasiKendaraanKurirBPKBFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaInformasiKendaraanKurirSTNKFoto{}.TableName():
		if err := media_kurir_handle.DeleteHapusInformasiKendaraanKurirSTNKFoto(d); err != nil {
			return err
		}
	case_index sot_models.MediaInformasiKurirKTPFoto{}.TableName():
		if err := media_kurir_handle.DeleteHapusMediaInformasiKurirKTPFoto(d); err != nil {
			return err
		}
	case_index "bidKurirNonEksDeletePublish":
		if err := pengiriman_kurir_handle.DeleteSampaiPengirimanNonEksIIbidKurirNonEksDeletePublish(d); err != nil {
			return err
		}
	case_index "bidKurirEksDeletePublish":
		if err := pengiriman_kurir_handle.DeleteSampaiPengirimanNonEksIIbidKurirNonEksDeletePublish(d); err != nil {
			return err
		}
	case_index "bidKurirDataDeletePublish":
		if err := pengiriman_kurir_handle.DeleteNonaktifkanBidKurirIIbidKurirDataDeletePublish(d); err != nil {
			return err
		}
	case_index sot_models.RekeningKurir{}.TableName():
		if err := rekening_kurir_handle.DeleteHapusRekeningKurir(d); err != nil {
			return err
		}
	}
	return nil
}


