package mb_cud_queue_provisioning

import (
	"fmt"
	"log"

	mb_cud_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/seeders/cud_exchange"
	"github.com/rabbitmq/amqp091-go"
)

type DeleteQueue struct {
	ExchangeName string
	QueueName    string
	QueueBind    string
	Durable      bool
	AutoDelete   bool
	Internal     bool
	NoWait       bool
	Exclusive    bool
	Payload      interface{}
}

func (DeleteQueue) BindingName() string {
	return mb_cud_seeders.Delete
}

func (d *DeleteQueue) ProvisioningQueues(ch *amqp091.Channel) error {
	// Declare queue
	_, err := ch.QueueDeclare(
		d.QueueName,
		d.Durable,
		d.AutoDelete,
		d.Exclusive,
		false, // no-wait = false supaya error muncul kalau gagal
		nil,
	)
	if err != nil {
		if amqpErr, ok := err.(*amqp091.Error); ok {
			if amqpErr.Code == 406 { // precondition_failed
				log.Printf("⚠️ Queue %s sudah ada, lewati declare", d.QueueName)
			} else {
				return fmt.Errorf("declare queue %s failed: %w", d.QueueName, err)
			}
		} else {
			return fmt.Errorf("declare queue %s failed: %w", d.QueueName, err)
		}
	}

	// Bind queue to exchange
	if err := ch.QueueBind(
		d.QueueName,
		d.BindingName(),
		d.ExchangeName,
		false, // no-wait = false
		nil,
	); err != nil {
		if amqpErr, ok := err.(*amqp091.Error); ok {
			if amqpErr.Code == 406 { // precondition_failed
				log.Printf("⚠️ Queue %s sudah bound ke exchange %s", d.QueueName, d.ExchangeName)
			} else {
				return fmt.Errorf("bind queue %s failed: %w", d.QueueName, err)
			}
		} else {
			return fmt.Errorf("bind queue %s failed: %w", d.QueueName, err)
		}
	}

	return nil
}
