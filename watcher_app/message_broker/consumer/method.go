package mb_cud_consumer

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"

	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func (c *Consumer) HandleCreate(ctx context.Context, msgs <-chan amqp091.Delivery) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-msgs:
			if !ok {
				return
			}

			fmt.Printf("[CREATE] %s\n", string(msg.Body))

			// 🔥 logic khusus CREATE di sini
			var data interface{}
			var status bool

			// coba JSON dulu
			data, status = mb_cud_serializer.ParseDataMessage[mb_cud_serializer.ConsumeDataJson](msg)

			c.Mu.Lock()

			if !status {
				// fallback ke Proto
				data, status = mb_cud_serializer.ParseDataMessage[mb_cud_serializer.ConsumeDataProto](msg)
				if !status {
					_ = msg.Nack(false, true)
					c.Mu.Unlock()
					continue
				}
			}
			switch v := data.(type) {

			case mb_cud_serializer.ConsumeDataJson:
				// JSON flow
				fmt.Println("JSON detected")
				fmt.Printf("Data: %+v\n", v)

				// contoh akses
				// v.Payload
				// v.TableName

			case mb_cud_serializer.ConsumeDataProto:
				// PROTO flow
				fmt.Println("PROTO detected")
				fmt.Printf("Raw bytes: %v\n", v.Payload)

				// di sini biasanya decode protobuf lagi
				// proto.Unmarshal(v.Payload, &yourStruct)

			default:
				fmt.Println("Unknown type after parsing")
				_ = msg.Nack(false, true)
				c.Mu.Unlock()
				continue
			}

			// kalau sukses
			_ = msg.Ack(false)
			c.Mu.Unlock()
		}
	}
}

func (c *Consumer) HandleUpdate(ctx context.Context, msgs <-chan amqp091.Delivery) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-msgs:
			if !ok {
				return
			}

			fmt.Printf("[UPDATE] %s\n", string(msg.Body))

			// 🔥 logic khusus UPDATE
			var data interface{}
			var status bool

			// coba JSON dulu
			data, status = mb_cud_serializer.ParseDataMessage[mb_cud_serializer.ConsumeDataJson](msg)

			c.Mu.Lock()

			if !status {
				// fallback ke Proto
				data, status = mb_cud_serializer.ParseDataMessage[mb_cud_serializer.ConsumeDataProto](msg)
				if !status {
					_ = msg.Nack(false, true)
					c.Mu.Unlock()
					continue
				}
			}
			switch v := data.(type) {

			case mb_cud_serializer.ConsumeDataJson:
				// JSON flow
				fmt.Println("JSON detected")
				fmt.Printf("Data: %+v\n", v)

				// contoh akses
				// v.Payload
				// v.TableName

			case mb_cud_serializer.ConsumeDataProto:
				// PROTO flow
				fmt.Println("PROTO detected")
				fmt.Printf("Raw bytes: %v\n", v.Payload)

				// di sini biasanya decode protobuf lagi
				// proto.Unmarshal(v.Payload, &yourStruct)

			default:
				fmt.Println("Unknown type after parsing")
				_ = msg.Nack(false, true)
				c.Mu.Unlock()
				continue
			}

			// kalau sukses
			_ = msg.Ack(false)
			c.Mu.Unlock()
		}
	}
}

func (c *Consumer) HandleDelete(ctx context.Context, msgs <-chan amqp091.Delivery) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-msgs:
			if !ok {
				return
			}

			fmt.Printf("[DELETE] %s\n", string(msg.Body))

			var data interface{}
			var status bool
			data, status = mb_cud_serializer.ParseDataMessage[mb_cud_serializer.ConsumeDataJson](msg)

			c.Mu.Lock()

			if !status {
				// fallback ke Proto
				data, status = mb_cud_serializer.ParseDataMessage[mb_cud_serializer.ConsumeDataProto](msg)
				if !status {
					_ = msg.Nack(false, true)
					c.Mu.Unlock()
					continue
				}
			}
			switch v := data.(type) {

			case mb_cud_serializer.ConsumeDataJson:
				// JSON flow
				fmt.Println("JSON detected")
				fmt.Printf("Data: %+v\n", v)

				// contoh akses
				// v.Payload
				// v.TableName

			case mb_cud_serializer.ConsumeDataProto:
				// PROTO flow
				fmt.Println("PROTO detected")
				fmt.Printf("Raw bytes: %v\n", v.Payload)

				// di sini biasanya decode protobuf lagi
				// proto.Unmarshal(v.Payload, &yourStruct)

			default:
				fmt.Println("Unknown type after parsing")
				_ = msg.Nack(false, true)
				c.Mu.Unlock()
				continue
			}

			// kalau sukses
			_ = msg.Ack(false)

			c.Mu.Unlock()
		}
	}
}
