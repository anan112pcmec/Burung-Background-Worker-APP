package mb_cud_serializer

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
)

func ParseDataMessage[T ConsumeDataJson | ConsumeDataProto](Data amqp091.Delivery) (T, bool) {
	var datanya T

	if err := helper.DecodeDeliveryBody(Data, &datanya); err != nil {
		fmt.Println("Jenis Data Tidak Diketahui")

		return datanya, false
	}

	return datanya, true
}
