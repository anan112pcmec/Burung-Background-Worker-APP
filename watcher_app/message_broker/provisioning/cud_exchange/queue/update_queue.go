package mb_cud_queue_provisioning

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"

	mb_cud_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/seeders/cud_exchange"
)

type UpdateQueue struct {
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

func (UpdateQueue) BindingName() string {
	return mb_cud_seeders.Update
}

func (u *UpdateQueue) ProvisioningQueues(ch *amqp091.Channel) error {
	// Declare queue
	_, err := ch.QueueDeclare(
		u.QueueName,
		u.Durable,    // durable
		u.AutoDelete, // auto-delete
		u.Exclusive,  // exclusive
		false,        // no-wait harus false supaya error muncul
		nil,
	)
	if err != nil {
		// Jika queue sudah ada tapi dengan properti sama, ignore error
		if amqpErr, ok := err.(*amqp091.Error); ok {
			if amqpErr.Code == 406 { // precondition_failed
				log.Printf("⚠️ Queue %s sudah ada, lewati declare", u.QueueName)
			} else {
				return fmt.Errorf("declare queue %s failed: %w", u.QueueName, err)
			}
		} else {
			return fmt.Errorf("declare queue %s failed: %w", u.QueueName, err)
		}
	}

	// Bind queue to exchange
	if err := ch.QueueBind(
		u.QueueName,
		u.BindingName(),
		u.ExchangeName,
		false, // no-wait = false supaya error muncul kalau gagal
		nil,
	); err != nil {
		// Sama, kalau sudah bound dengan routing key yang sama, aman untuk ignore
		if amqpErr, ok := err.(*amqp091.Error); ok {
			if amqpErr.Code == 406 { // precondition_failed
				log.Printf("⚠️ Queue %s sudah bound ke exchange %s", u.QueueName, u.ExchangeName)
			} else {
				return fmt.Errorf("bind queue %s failed: %w", u.QueueName, err)
			}
		} else {
			return fmt.Errorf("bind queue %s failed: %w", u.QueueName, err)
		}
	}

	return nil
}
