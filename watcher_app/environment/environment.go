package environment

import (
	"fmt"
	"log"
	"sync"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/meilisearch/meilisearch-go"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	mb_cud_consumer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/consumer"
	mb_cud_queue_provisioning "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/provisioning/cud_exchange/queue"
	mb_cud_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/seeders/cud_exchange"
)

const (
	ENVFILE = "env"
	YAML    = "yaml"
	JSON    = "json"
)

type Environment struct {
	DBMASTERHOST, DBMASTERUSER, DBMASTERPORT, DBMASTERPASS, DBMASTERNAME                           string
	DBREPLICAHOST, DBREPLICAUSER, DBREPLICAPORT, DBREPLICAPASS, DBREPLICANAME                      string
	RDSHOST, RDSPORT                                                                               string
	RDSAUTHENTICATION, RDSSESSION                                                                  int
	MEILIHOST, MEILIKEY, MEILIPORT                                                                 string
	RMQ_HOST, RMQ_USER, RMQ_PASS, EXCHANGE, RMQ_PORT                                               string
	RMQ_NOTIF_EXCHANGE                                                                             string
	CASS_HISTORICAL_KEYSPACE, CASS_HISTORICAL_USER, CASS_HISTORICAL_PASS, CASS_HISTORICAL_PORT     string
	CASS_SOT_REPLICA_KEYSPACE, CASS_SOT_REPLICA_USER, CASS_SOT_REPLICA_PASS, CASS_SOT_REPLICA_PORT string
}

type InternalDBReadWriteSystem struct {
	Write *gorm.DB
	Read  *gorm.DB
}

func (e *Environment) RunConnectionEnvironment() (
	db *InternalDBReadWriteSystem,
	redis_authentication *redis.Client,
	redis_session *redis.Client,
	search_engine meilisearch.ServiceManager,
	cud_consumer *mb_cud_consumer.Consumer,
	cass_historical_session *gocql.Session,
	cass_sot_replica_session *gocql.Session,
) {
	dsnWrite := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		e.DBMASTERHOST, e.DBMASTERUSER, e.DBMASTERPASS, e.DBMASTERNAME, e.DBMASTERPORT,
	)

	log.Println("🔍 Mencoba koneksi ke PostgreSQL...")
	log.Println("🔗 DSN:", dsnWrite)

	var err error
	db.Write, err = gorm.Open(postgres.Open(dsnWrite), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn), // pakai level Warn agar log tidak terlalu ramai
	})
	if err != nil {
		log.Fatalf("❌ Gagal konek ke PostgreSQL: %v", err)
	}

	// Coba koneksi langsung
	sqlDBWrite, err := db.Write.DB()
	if err != nil {
		log.Fatalf("❌ Gagal mendapatkan *sql.DB dari GORM: %v", err)
	}

	// Coba ping database untuk memastikan koneksi aktif
	if err := sqlDBWrite.Ping(); err != nil {
		log.Fatalf("❌ Gagal ping ke PostgreSQL: %v", err)
	}

	// Atur pool koneksi
	sqlDBWrite.SetMaxOpenConns(100)
	sqlDBWrite.SetMaxIdleConns(50)
	sqlDBWrite.SetConnMaxLifetime(time.Hour)

	var currentDBWrite string
	if err := db.Write.Raw("SELECT current_database();").Scan(&currentDBWrite).Error; err != nil {
		log.Printf("⚠️ Tidak bisa membaca nama database: %v", err)
	} else {
		log.Println("✅ Berhasil terkoneksi ke database:", currentDBWrite)
	}

	///

	dsnRead := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		e.DBREPLICAHOST, e.DBREPLICAUSER, e.DBREPLICAPASS, e.DBREPLICANAME, e.DBREPLICAPORT,
	)

	db.Read, err = gorm.Open(postgres.Open(dsnRead), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn), // pakai level Warn agar log tidak terlalu ramai
	})
	if err != nil {
		log.Fatalf("❌ Gagal konek ke PostgreSQL: %v", err)
	}

	// Coba koneksi langsung
	sqlDBRead, err := db.Read.DB()
	if err != nil {
		log.Fatalf("❌ Gagal mendapatkan *sql.DB dari GORM: %v", err)
	}

	// Coba ping database untuk memastikan koneksi aktif
	if err := sqlDBRead.Ping(); err != nil {
		log.Fatalf("❌ Gagal ping ke PostgreSQL: %v", err)
	}

	// Atur pool koneksi
	sqlDBRead.SetMaxOpenConns(100)
	sqlDBRead.SetMaxIdleConns(50)
	sqlDBRead.SetConnMaxLifetime(time.Hour)

	var currentDBRead string
	if err := db.Read.Raw("SELECT current_database();").Scan(&currentDBRead).Error; err != nil {
		log.Printf("⚠️ Tidak bisa membaca nama database: %v", err)
	} else {
		log.Println("✅ Berhasil terkoneksi ke database:", currentDBRead)
	}

	redis_authentication = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", e.RDSHOST, e.RDSPORT),
		Password: "",
		DB:       e.RDSAUTHENTICATION,
	})

	redis_session = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", e.RDSHOST, e.RDSPORT),
		Password: "",
		DB:       e.RDSSESSION,
	})

	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", e.RMQ_USER, e.RMQ_PASS, e.RMQ_HOST, e.RMQ_PORT)
	notification, _ := amqp091.Dial(connStr)
	cud_ch, err := notification.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	cud_consumer = &mb_cud_consumer.Consumer{
		Ch: cud_ch,
		QueueCreate: &mb_cud_queue_provisioning.CreateQueue{
			ExchangeName: mb_cud_seeders.ExchangeName,
			QueueName:    mb_cud_seeders.Create,
			QueueBind:    mb_cud_queue_provisioning.CreateQueue{}.BindingName(),
			Durable:      true,
			AutoDelete:   false,
			Internal:     false,
			NoWait:       false,
			Exclusive:    false,
		},
		QueueUpdate: &mb_cud_queue_provisioning.UpdateQueue{
			ExchangeName: mb_cud_seeders.ExchangeName,
			QueueName:    mb_cud_seeders.Update,
			QueueBind:    mb_cud_queue_provisioning.UpdateQueue{}.BindingName(),
			Durable:      true,
			AutoDelete:   false,
			Internal:     false,
			NoWait:       false,
			Exclusive:    false,
		},
		QueueDelete: &mb_cud_queue_provisioning.DeleteQueue{
			ExchangeName: mb_cud_seeders.ExchangeName,
			QueueName:    mb_cud_seeders.Delete,
			QueueBind:    mb_cud_queue_provisioning.DeleteQueue{}.BindingName(),
			Durable:      true,
			AutoDelete:   false,
			Internal:     false,
			NoWait:       false,
			Exclusive:    false,
		},
		Mu: sync.Mutex{},
	}

	search_engine = meilisearch.New(fmt.Sprintf("http://%s:%s", e.MEILIHOST, e.MEILIPORT), meilisearch.WithAPIKey(e.MEILIKEY))

	ch := gocql.NewCluster(fmt.Sprintf("127.0.0.1:%s", e.CASS_HISTORICAL_PORT))
	ch.Keyspace = e.CASS_HISTORICAL_KEYSPACE
	ch.ReconnectionPolicy = &gocql.ExponentialReconnectionPolicy{
		MaxRetries:      8,                // 9 total percobaan (0s, 1s, 2s, 4s, 8s, 16s, 30s, 30s, 30s)
		InitialInterval: 1 * time.Second,  // Dimulai pada 1 detik
		MaxInterval:     30 * time.Second, // Membatasi pertumbuhan eksponensial hingga 30 detik
	}
	ch.Authenticator = gocql.PasswordAuthenticator{
		Username: e.CASS_HISTORICAL_USER,
		Password: e.CASS_HISTORICAL_PASS,
	}

	cass_historical_session, err = ch.CreateSession()
	if err != nil {
		log.Fatal("gagal membuat session dengan cassandra", err)
	} else {
		fmt.Println("berhasil terhubung ke cassandra")
	}

	csr := gocql.NewCluster(fmt.Sprintf("127.0.0.1:%s", e.CASS_SOT_REPLICA_PORT))
	csr.Keyspace = e.CASS_SOT_REPLICA_KEYSPACE
	csr.ReconnectionPolicy = &gocql.ExponentialReconnectionPolicy{
		MaxRetries:      8,                // 9 total percobaan (0s, 1s, 2s, 4s, 8s, 16s, 30s, 30s, 30s)
		InitialInterval: 1 * time.Second,  // Dimulai pada 1 detik
		MaxInterval:     30 * time.Second, // Membatasi pertumbuhan eksponensial hingga 30 detik
	}
	csr.Authenticator = gocql.PasswordAuthenticator{
		Username: e.CASS_SOT_REPLICA_USER,
		Password: e.CASS_SOT_REPLICA_PASS,
	}

	cass_sot_replica_session, err = csr.CreateSession()
	if err != nil {
		log.Fatal("gagal membuat session dengan cassandra", err)
	} else {
		fmt.Println("berhasil terhubung ke cassandra")
	}

	return
}
