package mb_cud_exchange_provisioning

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	mb_cud_seeders "github.com/anan112pcmec/Burung-backend-1/app/message_broker/seeders/cud_exchange"
)

func ProvisionExchangeCUD(ch *amqp.Channel) error {

	if err := ch.ExchangeDeclare(
		mb_cud_seeders.ExchangeName, // name
		mb_cud_seeders.ExchangeType, // type
		true,                        // durable
		false,                       // auto-deleted
		false,                       // internal
		false,                       // no-wait
		nil,                         // arguments
	); err != nil {
		return fmt.Errorf("failed to declare exchange %s: %w", mb_cud_seeders.ExchangeName, err)
	}

	return nil
}
