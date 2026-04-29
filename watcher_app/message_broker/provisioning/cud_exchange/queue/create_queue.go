package mb_cud_queue_provisioning

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"

	mb_cud_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/seeders/cud_exchange"
)

type CreateQueue struct {
	ExchangeName string
	QueueName    string
	QueueBind    string
	Durable      bool
	AutoDelete   bool
	Internal     bool
	NoWait       bool
	Exclusive    bool
}

func (CreateQueue) BindingName() string {
	return mb_cud_seeders.Create
}

func (c *CreateQueue) ProvisioningQueues(ch *amqp091.Channel) error {
	// Declare queue
	_, err := ch.QueueDeclare(
		c.QueueName,
		c.Durable,
		c.AutoDelete,
		c.Exclusive,
		false, // no-wait = false supaya error muncul kalau gagal
		nil,
	)
	if err != nil {
		if amqpErr, ok := err.(*amqp091.Error); ok {
			if amqpErr.Code == 406 { // precondition_failed
				log.Printf("⚠️ Queue %s sudah ada, lewati declare", c.QueueName)
			} else {
				return fmt.Errorf("declare queue %s failed: %w", c.QueueName, err)
			}
		} else {
			return fmt.Errorf("declare queue %s failed: %w", c.QueueName, err)
		}
	}

	// Bind queue to exchange
	if err := ch.QueueBind(
		c.QueueName,
		c.BindingName(),
		c.ExchangeName,
		false, // no-wait = false
		nil,
	); err != nil {
		if amqpErr, ok := err.(*amqp091.Error); ok {
			if amqpErr.Code == 406 { // precondition_failed
				log.Printf("⚠️ Queue %s sudah bound ke exchange %s", c.QueueName, c.ExchangeName)
			} else {
				return fmt.Errorf("bind queue %s failed: %w", c.QueueName, err)
			}
		} else {
			return fmt.Errorf("bind queue %s failed: %w", c.QueueName, err)
		}
	}

	return nil
}
