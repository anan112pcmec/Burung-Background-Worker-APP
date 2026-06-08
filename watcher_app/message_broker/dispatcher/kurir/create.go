package consume_kurir_dispatcher

import (
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	alamat_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/alamat_services"
	informasi_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/informasi_services"
	media_kurir_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/kurir_service/media_services"

)

func KurirCreateServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](data T) error {
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
		if err := alamat_kurir_handle.CreateMasukanAlamatKurir(d); err != nil {
			return err
		}
	case models.InformasiKendaraanKurir{}.TableName():
		if err := informasi_kurir_handle.CreateAjukanInformasiKendaraan(d); err != nil {
			return err
		}
	case models.InformasiKurir{}.TableName():
		if err := informasi_kurir_handle.CreateAjukanInformasiKurir(d); err != nil {
			return err
		}
	case models.MediaInformasiKendaraanKurirKendaraanFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahMediaInformasiKendaraanKurirKendaraanFoto(d); err != nil {
			return err
		}
	case models.MediaInformasiKendaraanKurirBPKBFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahInformasiKendaraanKurirBPKBFoto(d); err != nil {
			return err
		}
	case models.MediaInformasiKendaraanKurirSTNKFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahInformasiKendaraanKurirSTNKFoto(d); err != nil {
			return err
		}
	case models.MediaInformasiKurirKTPFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahMediaInformasiKurirKTPFoto(d); err != nil {
			return err
		}
	case models.MediaPengirimanPickedUpFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahMediaPengirimanPickedUpFoto(d); err != nil {
			return err
		}
	case models.MediaPengirimanSampaiFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahMediaPengirimanSampaiFoto(d); err != nil {
			return err
		}
	case models.MediaPengirimanEkspedisiPickedUpFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahMediaPengirimanEkspedisiPickedUpFoto(d); err != nil {
			return err
		}
	case models.MediaPengirimanEkspedisiSampaiAgentFoto{}.TableName():
		if err := media_kurir_handle.CreateTambahMediaPengirimanEkspedisiSampaiAgentFoto(d); err != nil {
			return err
		}

	}

	return nil
}
