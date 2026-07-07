package mb_cud_consumer

import (
	"context"
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	se_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/search_engine/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	consume_kurir_dispatcher "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/dispatcher/kurir"
	consume_pengguna_dispatcher "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/dispatcher/pengguna"
	consume_seller_dispatcher "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/dispatcher/seller"
	mb_cud_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/seeders/cud_exchange"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func (c *Consumer) HandleCreate(ctx context.Context, msgs <-chan amqp091.Delivery, read *gorm.DB, redis_authentication, redis_session *redis.Client, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-msgs:
			if !ok {
				return
			}

			fmt.Printf("[CREATE] %s\n", string(msg.Body))

			// 1. Coba hajar pakai JSON dulu
			if dataJson, ok := helper.ParseDataMessage[mb_cud_serializer.ConsumeDataJson](msg); ok {
				fmt.Println("JSON detected")
				fmt.Println("[DATAJSON TRACE: ]", dataJson)

				switch dataJson.Role {
				case mb_cud_seeders.Pengguna:
					if err := consume_pengguna_dispatcher.PenggunaCreateServicesDispatcher(ctx, &dataJson, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				case mb_cud_seeders.Seller:
					if err := consume_seller_dispatcher.SellerCreateServicesDispatcher(ctx, &dataJson, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				case mb_cud_seeders.Kurir:
					if err := consume_kurir_dispatcher.KurirCreateServicesDispatcher(ctx, &dataJson, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				default:
					fmt.Println("role tidak diketahui")
				}

				_ = msg.Ack(false)
				continue
			}

			// 2. Kalau JSON gagal, fallback ke PROTO
			if dataProto, ok := helper.ParseDataMessage[mb_cud_serializer.ConsumeDataProto](msg); ok {
				fmt.Println("PROTO detected")

				switch dataProto.Role {
				case mb_cud_seeders.Pengguna:
					if err := consume_pengguna_dispatcher.PenggunaCreateServicesDispatcher(ctx, &dataProto, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				case mb_cud_seeders.Seller:
					if err := consume_seller_dispatcher.SellerCreateServicesDispatcher(ctx, &dataProto, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				case mb_cud_seeders.Kurir:
					if err := consume_kurir_dispatcher.KurirCreateServicesDispatcher(ctx, &dataProto, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				default:
					fmt.Println("role tidak diketahui")
				}

				_ = msg.Ack(false)
				continue
			}

			// 3. Gagal parsing dua-duanya
			fmt.Println("Unknown type after parsing")
			_ = msg.Nack(false, true)
		}
	}
}

func (c *Consumer) HandleUpdate(ctx context.Context, msgs <-chan amqp091.Delivery, read *gorm.DB, redis_authentication, redis_session *redis.Client, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-msgs:
			if !ok {
				return
			}

			fmt.Printf("[UPDATE] %s\n", string(msg.Body))

			// 1. Coba JSON dulu
			if dataJson, ok := helper.ParseDataMessage[mb_cud_serializer.ConsumeDataJson](msg); ok {
				fmt.Println("JSON detected")

				switch dataJson.Role {
				case mb_cud_seeders.Pengguna:
					if err := consume_pengguna_dispatcher.PenggunaUpdateServicesDispatcher(ctx, &dataJson, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				case mb_cud_seeders.Seller:
					if err := consume_seller_dispatcher.SellerUpdateServicesDispatcher(ctx, &dataJson, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				case mb_cud_seeders.Kurir:
					if err := consume_kurir_dispatcher.KurirUpdateServicesDispatcher(ctx, &dataJson, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				default:
					fmt.Println("role tidak diketahui")
				}

				_ = msg.Ack(false)
				continue
			}

			// 2. Fallback ke PROTO
			if dataProto, ok := helper.ParseDataMessage[mb_cud_serializer.ConsumeDataProto](msg); ok {
				fmt.Println("PROTO detected")

				switch dataProto.Role {
				case mb_cud_seeders.Pengguna:
					if err := consume_pengguna_dispatcher.PenggunaUpdateServicesDispatcher(ctx, &dataProto, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				case mb_cud_seeders.Seller:
					if err := consume_seller_dispatcher.SellerUpdateServicesDispatcher(ctx, &dataProto, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				case mb_cud_seeders.Kurir:
					if err := consume_kurir_dispatcher.KurirUpdateServicesDispatcher(ctx, &dataProto, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				default:
					fmt.Println("role tidak diketahui")
				}

				_ = msg.Ack(false)
				continue
			}

			// 3. Gagal parsing dua-duanya
			fmt.Println("Unknown type after parsing")
			_ = msg.Nack(false, true)
		}
	}
}

func (c *Consumer) HandleDelete(ctx context.Context, msgs <-chan amqp091.Delivery, read *gorm.DB, redis_authentication, redis_session *redis.Client, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-msgs:
			if !ok {
				return
			}

			fmt.Printf("[DELETE] %s\n", string(msg.Body))

			// 1. Coba JSON dulu
			if dataJson, ok := helper.ParseDataMessage[mb_cud_serializer.ConsumeDataJson](msg); ok {
				fmt.Println("JSON detected")

				switch dataJson.Role {
				case mb_cud_seeders.Pengguna:
					if err := consume_pengguna_dispatcher.PenggunaDeleteServicesDispatcher(ctx, &dataJson, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				case mb_cud_seeders.Seller:
					if err := consume_seller_dispatcher.SellerDeleteServicesDispatcher(ctx, &dataJson, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				case mb_cud_seeders.Kurir:
					if err := consume_kurir_dispatcher.KurirDeleteServicesDispatcher(ctx, &dataJson, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				default:
					fmt.Println("role tidak diketahui")
				}

				_ = msg.Ack(false)
				continue
			}

			// 2. Fallback ke PROTO
			if dataProto, ok := helper.ParseDataMessage[mb_cud_serializer.ConsumeDataProto](msg); ok {
				fmt.Println("PROTO detected")

				switch dataProto.Role {
				case mb_cud_seeders.Pengguna:
					if err := consume_pengguna_dispatcher.PenggunaDeleteServicesDispatcher(ctx, &dataProto, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				case mb_cud_seeders.Seller:
					if err := consume_seller_dispatcher.SellerDeleteServicesDispatcher(ctx, &dataProto, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				case mb_cud_seeders.Kurir:
					if err := consume_kurir_dispatcher.KurirDeleteServicesDispatcher(ctx, &dataProto, read, redis_authentication, redis_session, cass_historcal, cass_sot_replica, se_index); err != nil {
						fmt.Println(err)
					}
				default:
					fmt.Println("role tidak diketahui")
				}

				_ = msg.Ack(false)
				continue
			}

			// 3. Gagal parsing dua-duanya
			fmt.Println("Unknown type after parsing")
			_ = msg.Nack(false, true)
		}
	}
}
