package mb_cud_consumer

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	consume_kurir_dispatcher "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/dispatcher/kurir"
	consume_pengguna_dispatcher "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/dispatcher/pengguna"
	consume_seller_dispatcher "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/dispatcher/seller"
	mb_cud_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/seeders/cud_exchange"
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
			data, status = helper.ParseDataMessage[mb_cud_serializer.ConsumeDataJson](msg)

			c.Mu.Lock()

			if !status {
				// fallback ke Proto
				data, status = helper.ParseDataMessage[mb_cud_serializer.ConsumeDataProto](msg)
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

				switch v.Role {
				case mb_cud_seeders.Pengguna:

					if err := consume_pengguna_dispatcher.PenggunaCreateServicesDispatcher[mb_cud_serializer.ConsumeDataJson](&v); err != nil {
						fmt.Println(err)
					}

				case mb_cud_seeders.Seller:

					if err := consume_seller_dispatcher.SellerCreateServicesDispatcher[mb_cud_serializer.ConsumeDataJson](v); err != nil {
						fmt.Println(err)
					}

				case mb_cud_seeders.Kurir:
					if err := consume_kurir_dispatcher.KurirCreateServicesDispatcher[mb_cud_serializer.ConsumeDataJson](v); err != nil {
						fmt.Println(err)
					}
				default:
					fmt.Println("role tidak diketahui")
				}

			case mb_cud_serializer.ConsumeDataProto:
				// PROTO flow
				fmt.Println("PROTO detected")
				fmt.Printf("Raw bytes: %v\n", v.Data)

				switch v.Role {
				case mb_cud_seeders.Pengguna:

					if err := consume_pengguna_dispatcher.PenggunaCreateServicesDispatcher[mb_cud_serializer.ConsumeDataProto](&v); err != nil {
						fmt.Println(err)
					}

				case mb_cud_seeders.Seller:

					if err := consume_seller_dispatcher.SellerCreateServicesDispatcher[mb_cud_serializer.ConsumeDataProto](v); err != nil {
						fmt.Println(err)
					}

				case mb_cud_seeders.Kurir:
					if err := consume_kurir_dispatcher.KurirCreateServicesDispatcher[mb_cud_serializer.ConsumeDataProto](v); err != nil {
						fmt.Println(err)
					}
				default:
					fmt.Println("role tidak diketahui")
				}

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
			data, status = helper.ParseDataMessage[mb_cud_serializer.ConsumeDataJson](msg)

			c.Mu.Lock()

			if !status {
				// fallback ke Proto
				data, status = helper.ParseDataMessage[mb_cud_serializer.ConsumeDataProto](msg)
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

				switch v.Role {
				case mb_cud_seeders.Pengguna:

					if err := consume_pengguna_dispatcher.PenggunaCreateServicesDispatcher[mb_cud_serializer.ConsumeDataJson](&v); err != nil {
						fmt.Println(err)
					}

				case mb_cud_seeders.Seller:

					if err := consume_seller_dispatcher.SellerCreateServicesDispatcher[mb_cud_serializer.ConsumeDataJson](v); err != nil {
						fmt.Println(err)
					}

				case mb_cud_seeders.Kurir:
					if err := consume_kurir_dispatcher.KurirCreateServicesDispatcher[mb_cud_serializer.ConsumeDataJson](v); err != nil {
						fmt.Println(err)
					}
				default:
					fmt.Println("role tidak diketahui")
				}

			case mb_cud_serializer.ConsumeDataProto:
				// PROTO flow
				fmt.Println("PROTO detected")
				fmt.Printf("Raw bytes: %v\n", v.Data)

				switch v.Role {
				case mb_cud_seeders.Pengguna:

					if err := consume_pengguna_dispatcher.PenggunaUpdateServicesDispatcher[mb_cud_serializer.ConsumeDataProto](&v); err != nil {
						fmt.Println(err)
					}

				case mb_cud_seeders.Seller:

					if err := consume_seller_dispatcher.SellerUpdateServicesDispatcher[mb_cud_serializer.ConsumeDataProto](v); err != nil {
						fmt.Println(err)
					}

				case mb_cud_seeders.Kurir:
					if err := consume_kurir_dispatcher.KurirUpdateServicesDispatcher[mb_cud_serializer.ConsumeDataProto](v); err != nil {
						fmt.Println(err)
					}
				default:
					fmt.Println("role tidak diketahui")
				}

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
			data, status = helper.ParseDataMessage[mb_cud_serializer.ConsumeDataJson](msg)

			c.Mu.Lock()

			if !status {
				// fallback ke Proto
				data, status = helper.ParseDataMessage[mb_cud_serializer.ConsumeDataProto](msg)
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

				switch v.Role {
				case mb_cud_seeders.Pengguna:

					if err := consume_pengguna_dispatcher.PenggunaCreateServicesDispatcher[mb_cud_serializer.ConsumeDataJson](&v); err != nil {
						fmt.Println(err)
					}

				case mb_cud_seeders.Seller:

					if err := consume_seller_dispatcher.SellerCreateServicesDispatcher[mb_cud_serializer.ConsumeDataJson](v); err != nil {
						fmt.Println(err)
					}

				case mb_cud_seeders.Kurir:
					if err := consume_kurir_dispatcher.KurirCreateServicesDispatcher[mb_cud_serializer.ConsumeDataJson](v); err != nil {
						fmt.Println(err)
					}
				default:
					fmt.Println("role tidak diketahui")
				}

			case mb_cud_serializer.ConsumeDataProto:
				// PROTO flow
				fmt.Println("PROTO detected")
				fmt.Printf("Raw bytes: %v\n", v.Data)

				switch v.Role {
				case mb_cud_seeders.Pengguna:

					if err := consume_pengguna_dispatcher.PenggunaCreateServicesDispatcher[mb_cud_serializer.ConsumeDataProto](&v); err != nil {
						fmt.Println(err)
					}

				case mb_cud_seeders.Seller:

					if err := consume_seller_dispatcher.SellerCreateServicesDispatcher[mb_cud_serializer.ConsumeDataProto](v); err != nil {
						fmt.Println(err)
					}

				case mb_cud_seeders.Kurir:
					if err := consume_kurir_dispatcher.KurirCreateServicesDispatcher[mb_cud_serializer.ConsumeDataProto](v); err != nil {
						fmt.Println(err)
					}
				default:
					fmt.Println("role tidak diketahui")
				}

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
