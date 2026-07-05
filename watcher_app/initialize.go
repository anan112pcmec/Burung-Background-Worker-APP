package watcher_app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/joho/godotenv"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"

	historical_migrations "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/migrations"
	sot_replica_migration "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/sot_replica_async/migration"
	se_initialize "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/search_engine/initialize"
	se_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/search_engine/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/environment"
	mb_cud_consumer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/consumer"
)

type Connection struct {
	db                       environment.InternalDBReadWriteSystem
	redis_authentication     *redis.Client
	redis_session            *redis.Client
	search_engine            meilisearch.ServiceManager
	cud_consumer             *mb_cud_consumer.Consumer
	cass_historical_session  *gocql.Session
	cass_sot_replica_session *gocql.Session
}

func Getenvi(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var conn Connection

	if err := godotenv.Load(); err != nil {
		log.Fatalf("❌ Tidak ada file .env")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// parsing env
	rdsauthentication, _ := strconv.Atoi(Getenvi("RDSAUTHENTICATION", "0"))
	rdssession, _ := strconv.Atoi(Getenvi("RDSSESSION", "0"))

	env := environment.Environment{
		DBMASTERHOST:              Getenvi("DBMASTERHOST", "NIL"),
		DBMASTERUSER:              Getenvi("DBMASTERUSER", "NIL"),
		DBMASTERPORT:              Getenvi("DBMASTERPORT", "NIL"),
		DBMASTERPASS:              Getenvi("DBMASTERPASS", "NIL"),
		DBMASTERNAME:              Getenvi("DBMASTERNAME", "NIL"),
		DBREPLICAHOST:             Getenvi("DBREPLICAHOST", "NIL"),
		DBREPLICAUSER:             Getenvi("DBREPLICAUSER", "NIL"),
		DBREPLICAPASS:             Getenvi("DBREPLICAPASS", "NIL"),
		DBREPLICAPORT:             Getenvi("DBREPLICAPORT", "NIL"),
		DBREPLICANAME:             Getenvi("DBREPLICANAME", "NIL"),
		RDSHOST:                   Getenvi("RDSHOST", "NIL"),
		RDSPORT:                   Getenvi("RDSPORT", "NIL"),
		RDSAUTHENTICATION:         rdsauthentication,
		RDSSESSION:                rdssession,
		MEILIHOST:                 Getenvi("MEILIHOST", "NIL"),
		MEILIPORT:                 Getenvi("MEILIPORT", "NIL"),
		MEILIKEY:                  Getenvi("MEILIKEY", "NIL"),
		RMQ_HOST:                  Getenvi("RMQ_HOST", "NIL"),
		RMQ_USER:                  Getenvi("RMQ_USER", "NIL"),
		RMQ_PASS:                  Getenvi("RMQ_PASS", "NIL"),
		RMQ_PORT:                  Getenvi("RMQ_PORT", "NIL"),
		RMQ_NOTIF_EXCHANGE:        Getenvi("RMQ_NOTIF_EXCHANGE", "NIL"),
		CASS_HISTORICAL_KEYSPACE:  Getenvi("CASS_HISTORICAL_KEYSPACE", "NIL"),
		CASS_HISTORICAL_USER:      Getenvi("CASS_HISTORICAL_USER", "NIL"),
		CASS_HISTORICAL_PASS:      Getenvi("CASS_HISTORICAL_PASS", "NIL"),
		CASS_HISTORICAL_PORT:      Getenvi("CASS_HISTORICAL_PORT", "NIL"),
		CASS_SOT_REPLICA_KEYSPACE: Getenvi("CASS_SOT_REPLICA_KEYSPACE", "NIL"),
		CASS_SOT_REPLICA_USER:     Getenvi("CASS_SOT_REPLICA_USER", "NIL"),
		CASS_SOT_REPLICA_PASS:     Getenvi("CASS_SOT_REPLICA_PASS", "NIL"),
		CASS_SOT_REPLICA_PORT:     Getenvi("CASS_SOT_REPLICA_PORT", "NIL"),
	}

	// init connection
	conn.db, conn.redis_authentication, conn.redis_session, conn.search_engine, conn.cud_consumer, conn.cass_historical_session, conn.cass_sot_replica_session = env.RunConnectionEnvironment()

	if err := historical_migrations.DownRelation(ctx, conn.cass_historical_session); len(err) > 0 {
		fmt.Println("gagal")
	}
	if err := historical_migrations.UpRelation(ctx, conn.cass_historical_session); len(err) > 0 {
		for _, e := range err {
			fmt.Println(e)
		}
	}

	if err := sot_replica_migration.DownMigration(ctx, conn.cass_sot_replica_session); len(err) > 0 {
		fmt.Println("gagal")
	}

	if err := sot_replica_migration.UpRelation(ctx, conn.cass_sot_replica_session); len(err) > 0 {
		for _, e := range err {
			fmt.Println(e)
		}
	}

	var searchEngineIndex se_models.IndexWrapper = se_initialize.InitIndex(ctx, conn.search_engine)

	var wg sync.WaitGroup

	// 🟢 start consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		conn.cud_consumer.WatchPublish(ctx, conn.db.Write, conn.redis_authentication, conn.redis_session, conn.cass_historical_session, conn.cass_sot_replica_session, searchEngineIndex)
	}()

	fmt.Println("Watcher berjalan... tekan CTRL+C untuk exit")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	fmt.Println("Shutting down...")

	cancel() // stop semua goroutine via context

	wg.Wait() // tunggu semua selesai

	fmt.Println("Shutdown selesai")
}
