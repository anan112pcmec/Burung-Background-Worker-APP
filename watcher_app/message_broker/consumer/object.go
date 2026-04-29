package mb_cud_consumer

import (
	"context"
	"sync"

	"github.com/rabbitmq/amqp091-go"

	mb_cud_queue_provisioning "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/provisioning/cud_exchange/queue"
)

type Consumer struct {
	Ch          *amqp091.Channel
	QueueCreate *mb_cud_queue_provisioning.CreateQueue
	QueueUpdate *mb_cud_queue_provisioning.UpdateQueue
	QueueDelete *mb_cud_queue_provisioning.DeleteQueue
	Mu          sync.Mutex
}

func (c *Consumer) WatchPublish(ctx context.Context) error {

	// 🔒 QoS biar gak overconsume
	c.Mu.Lock()
	err := c.Ch.Qos(10, 0, false)
	c.Mu.Unlock()
	if err != nil {
		return err
	}

	// 🔥 consume dari masing-masing queue
	createConsume, err := c.Ch.Consume(
		c.QueueCreate.QueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	updateConsume, err := c.Ch.Consume(
		c.QueueUpdate.QueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	deleteConsume, err := c.Ch.Consume(
		c.QueueDelete.QueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// 🔥 jalanin handler terpisah
	go c.HandleCreate(ctx, createConsume)
	go c.HandleUpdate(ctx, updateConsume)
	go c.HandleDelete(ctx, deleteConsume)

	// 🔥 blocking sampai context selesai
	<-ctx.Done()

	// optional: cleanup
	c.Mu.Lock()
	_ = c.Ch.Close()
	c.Mu.Unlock()

	return nil
}
