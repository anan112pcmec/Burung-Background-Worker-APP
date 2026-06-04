package consume_kurir_dispatcher

import (
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	alamat_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/alamat_services"
	media_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/media_services"
)

func KurirDeleteServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](data T) error {
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
	case models.AlamatKurir{}.TableName():
		if err := alamat_kurir_handle.DeleteHapusAlamatKurir(d); err != nil {
			return err
		}
	case models.MediaKurirProfilFoto{}.TableName():
		if err := media_kurir_handle.DeleteHapusKurirProfilFoto(d); err != nil {
			return err
		}
	case models.MediaInformasiKendaraanKurirKendaraanFoto{}.TableName():
		if err := media_kurir_handle.DeleteHapusMediaInformasiKendaraanKurirKendaraanFoto(d); err != nil {
			return err
		}
	case models.MediaInformasiKendaraanKurirBPKBFoto{}.TableName():
		if err := media_kurir_handle.DeleteHapusInformasiKendaraanKurirBPKBFoto(d); err != nil {
			return err
		}
	case models.MediaInformasiKendaraanKurirSTNKFoto{}.TableName():
		if err := media_kurir_handle.DeleteHapusInformasiKendaraanKurirSTNKFoto(d); err != nil {
			return err
		}
	case models.MediaInformasiKurirKTPFoto{}.TableName():
		if err := media_kurir_handle.DeleteHapusMediaInformasiKurirKTPFoto(d); err != nil {
			return err
		}

	}
	return nil
}
