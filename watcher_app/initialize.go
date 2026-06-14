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
	"gorm.io/gorm"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/config"
	mb_cud_consumer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/consumer"
)

type Connection struct {
	DB            *gorm.DB
	RDSENTITY     *redis.Client
	RDSBARANG     *redis.Client
	RDSENGAGEMENT *redis.Client
	SE            meilisearch.ServiceManager
	CUD_CONSUMER  *mb_cud_consumer.Consumer
	HDB           *gocql.Session
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
	rdsentity, _ := strconv.Atoi(Getenvi("RDSENTITY", "0"))
	rdsbarang, _ := strconv.Atoi(Getenvi("RDSBARANG", "0"))
	rdsengagement, _ := strconv.Atoi(Getenvi("RDSENGAGEMENT", "0")) // ✅ typo fix

	env := config.Environment{
		DBHOST:             Getenvi("DBHOST", "NIL"),
		DBUSER:             Getenvi("DBUSER", "NIL"),
		DBPASS:             Getenvi("DBPASS", "NIL"),
		DBNAME:             Getenvi("DBNAME", "NIL"),
		DBPORT:             Getenvi("DBPORT", "NIL"),
		RDSHOST:            Getenvi("RDSHOST", "NIL"),
		RDSPORT:            Getenvi("RDSPORT", "NIL"),
		RDSENTITYDB:        rdsentity,
		RDSBARANGDB:        rdsbarang,
		RDSENGAGEMENTDB:    rdsengagement,
		MEILIHOST:          Getenvi("MEILIHOST", "NIL"),
		MEILIPORT:          Getenvi("MEILIPORT", "NIL"),
		MEILIKEY:           Getenvi("MEILIKEY", "NIL"),
		RMQ_HOST:           Getenvi("RMQ_HOST", "NIL"),
		RMQ_USER:           Getenvi("RMQ_USER", "NIL"),
		RMQ_PASS:           Getenvi("RMQ_PASS", "NIL"),
		RMQ_PORT:           Getenvi("RMQ_PORT", "NIL"),
		RMQ_NOTIF_EXCHANGE: Getenvi("RMQ_NOTIF_EXCHANGE", "NIL"),
		CASS_KEYSPACE:      Getenvi("CASS_KEYSPACE", "NIL"),
		CASS_USER:          Getenvi("CASS_USER", "NIL"),
		CASS_PASS:          Getenvi("CASS_PASS", "NIL"),
		CASS_PORT:          Getenvi("CASS_PORT", "NIL"),
	}

	// init connection
	conn.DB, conn.RDSENTITY, conn.RDSBARANG, conn.RDSENGAGEMENT, conn.SE, conn.CUD_CONSUMER, conn.HDB = env.RunConnectionEnvironment()

	var wg sync.WaitGroup

	// 🟢 start consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		conn.CUD_CONSUMER.WatchPublish(ctx, conn.DB, *conn.RDSBARANG, *conn.RDSENGAGEMENT, conn.SE)
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
