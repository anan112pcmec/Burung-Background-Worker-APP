package environment

import (
	"fmt"
	"log"
	"net"
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

type connObserver struct {
	name string
}

func (o connObserver) ObserveConnect(oc gocql.ObservedConnect) {
	if oc.Err != nil {
		fmt.Printf("🔴 [%s] gagal connect ke host %v: %v (durasi: %v)\n",
			o.name, oc.Host.ConnectAddress(), oc.Err, oc.End.Sub(oc.Start))
	} else {
		fmt.Printf("🟢 [%s] berhasil connect ke host %v (durasi: %v)\n",
			o.name, oc.Host.ConnectAddress(), oc.End.Sub(oc.Start))
	}
}

type staticAddressTranslator struct {
	ip   net.IP
	port int
}

func (t staticAddressTranslator) Translate(addr net.IP, port int) (net.IP, int) {
	return t.ip, t.port
}

func connectWithRetry(cluster *gocql.ClusterConfig, name string, maxAttempts int) (*gocql.Session, error) {
	var session *gocql.Session
	var err error

	fmt.Printf("=== [%s] cluster config ===\n", name)
	fmt.Printf("Hosts: %v\n", cluster.Hosts)
	fmt.Printf("Port: %d\n", cluster.Port)
	fmt.Printf("Keyspace: %s\n", cluster.Keyspace)
	fmt.Printf("ProtoVersion: %d\n", cluster.ProtoVersion)
	fmt.Printf("Consistency: %v\n", cluster.Consistency)
	if auth, ok := cluster.Authenticator.(gocql.PasswordAuthenticator); ok {
		fmt.Printf("Auth Username: %q\n", auth.Username)
		fmt.Printf("Auth Password len: %q\n", auth.Password)
	} else {
		fmt.Printf("Authenticator: %T (bukan PasswordAuthenticator!)\n", cluster.Authenticator)
	}
	fmt.Printf("ConnectTimeout: %v\n", cluster.ConnectTimeout)
	fmt.Printf("Timeout: %v\n", cluster.Timeout)
	fmt.Println("===========================")

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		session, err = cluster.CreateSession()
		if err == nil {
			fmt.Printf("✅ berhasil terhubung ke %s (percobaan ke-%d)\n", name, attempt)
			return session, nil
		}

		fmt.Printf("❌ gagal connect ke %s (percobaan %d/%d)\n", name, attempt, maxAttempts)
		fmt.Printf("   error type: %T\n", err)
		fmt.Printf("   error detail: %+v\n", err)

		if attempt < maxAttempts {
			time.Sleep(3 * time.Second)
		}
	}

	return nil, err
}

func (e *Environment) RunConnectionEnvironment() (
	db InternalDBReadWriteSystem,
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

	connStr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/internal_system_burung",
		e.RMQ_USER,
		e.RMQ_PASS,
		e.RMQ_HOST,
		e.RMQ_PORT,
	)
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

	ch := gocql.NewCluster("127.0.0.1")
	ch.Port = 9042
	ch.DisableInitialHostLookup = true
	ch.ProtoVersion = 4
	ch.Keyspace = e.CASS_HISTORICAL_KEYSPACE
	ch.ConnectObserver = connObserver{name: "Cassandra Historical"}
	ch.AddressTranslator = staticAddressTranslator{ip: net.ParseIP("127.0.0.1"), port: 9042}
	ch.ReconnectionPolicy = &gocql.ExponentialReconnectionPolicy{
		MaxRetries:      8,
		InitialInterval: 1 * time.Second,
		MaxInterval:     30 * time.Second,
	}
	ch.Authenticator = gocql.PasswordAuthenticator{
		Username: e.CASS_HISTORICAL_USER,
		Password: e.CASS_HISTORICAL_PASS,
	}

	cass_historical_session, err = connectWithRetry(ch, "Cassandra Historical", 10)
	if err != nil {
		log.Fatal("gagal membuat session dengan cassandra historical: ", err)
	}

	csr := gocql.NewCluster("127.0.0.1")
	csr.Port = 9043
	csr.DisableInitialHostLookup = true
	csr.ProtoVersion = 4
	csr.Keyspace = e.CASS_SOT_REPLICA_KEYSPACE
	csr.ConnectObserver = connObserver{name: "Cassandra SOT"}
	csr.AddressTranslator = staticAddressTranslator{ip: net.ParseIP("127.0.0.1"), port: 9043}
	csr.ReconnectionPolicy = &gocql.ExponentialReconnectionPolicy{
		MaxRetries:      8,
		InitialInterval: 1 * time.Second,
		MaxInterval:     30 * time.Second,
	}
	csr.Authenticator = gocql.PasswordAuthenticator{
		Username: e.CASS_SOT_REPLICA_USER,
		Password: e.CASS_SOT_REPLICA_PASS,
	}

	cass_sot_replica_session, err = connectWithRetry(csr, "Cassandra SOT Replica", 10)
	if err != nil {
		log.Fatal("gagal membuat session dengan cassandra sot replica: ", err)
	}

	return
}
