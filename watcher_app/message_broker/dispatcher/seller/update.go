package consume_seller_dispatcher

import (
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	alamat_seller_handle "github.com/anan112pcmec/Burung-backend-2/watcher_app/service_handle/seller_service/alamat_services"

)

func SellerUpdateServicesDispatcher[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](data T) error {
	var d mb_cud_serializer.ParsedDataMessage
	switch v := any(data).(type) {
	case mb_cud_serializer.ConsumeDataJson:
		d = v.Parse()

		switch d.TableName {
		case models.AlamatGudang.TableName(models.AlamatGudang{}):
			if err := alamat_seller_handle.UpdateEditAlamatGudang(d); err != nil {
				return err
			}

		}

	case mb_cud_serializer.ConsumeDataProto:
		d = v.Parse()

		switch d.TableName {
		case models.AlamatGudang.TableName(models.AlamatGudang{}):
			if err := alamat_seller_handle.UpdateEditAlamatGudang(d); err != nil {
				return err
			}

		}
	}

	return nil
}
