# Burung Background Worker App

Background worker service untuk memproses perubahan data SOT secara asinkron — mendistribusikan event ke seluruh storage engine dan meneruskan notifikasi internal.

## Stack

<div align="center">
  <table>
    <tr>
      <td align="center" width="90">
        <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_LightBlue.png" height="40" alt="go logo" />
        <br/><sub><b>Go</b></sub>
      </td>
      <td align="center" width="90">
        <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/postgresql/postgresql-original.svg" height="40" alt="postgresql logo" />
        <br/><sub><b>PostgreSQL</b></sub>
      </td>
      <td align="center" width="90">
        <img src="https://cdn.simpleicons.org/redis/FF4438" height="40" alt="redis logo" />
        <br/><sub><b>Redis</b></sub>
      </td>
      <td align="center" width="90">
        <img src="https://cdn.simpleicons.org/rabbitmq/FF6600" height="40" alt="rabbitmq logo" />
        <br/><sub><b>RabbitMQ</b></sub>
      </td>
      <td align="center" width="90">
        <img src="https://cdn.simpleicons.org/apachecassandra/1287B1" height="40" alt="cassandra logo" />
        <br/><sub><b>Cassandra</b></sub>
      </td>
      <td align="center" width="90">
        <img src="https://cdn.simpleicons.org/meilisearch/FF5CAA" height="40" alt="meilisearch logo" />
        <br/><sub><b>Meilisearch</b></sub>
      </td>
      <td align="center" width="90">
        <img src="https://cdn.simpleicons.org/clickhouse/FFCC01" height="40" alt="clickhouse logo" />
        <br/><sub><b>ClickHouse</b></sub>
      </td>
    </tr>
  </table>
</div>
## Architecture

<img src="https://github.com/user-attachments/assets/0f2d3565-e6a5-48b2-a37c-15f7e62a84c8" />

| Layer | Komponen | Peran |
|---|---|---|
| Message broker | RabbitMQ | Consume SOT change event |
| Worker | burung-background-worker-app | Orchestrate distribusi & forward |
| Historical DB | Cassandra `historical_db` | Audit trail & riwayat perubahan |
| Async Replica | Cassandra `sot_replica_async` | Async replica dari SOT |
| Search engine | Meilisearch | Full-text index untuk query |
| Analytical DB | ClickHouse | OLAP & reporting |
| Session cache | Redis | Sinkronisasi cache session pengguna |
| Notification | Burung-Internal-Notificationing-App | Forward perubahan via HTTP |

## Flow

```
RabbitMQ (SOT change event)
        │
        ▼ consume
burung-background-worker-app
        │
        ├──────────────────────────────────────┐
        │                                      │
        ├── Cassandra historical_db            │
        │       └── Cassandra sot_replica_async│
        │                                      │
        ├── Meilisearch (search index)         │
        │                                      │
        ├── ClickHouse (analytical DB)         │
        │                                      │
        └── Redis (user session cache sync) ───┘
                        │
                        ▼ HTTP POST · PATCH · PUT · DELETE
        Burung-Internal-Notificationing-App
                        │
                        ▼
                     Selesai ✓
```

## Getting Started

```bash
git clone https://github.com/<your-org>/burung-background-worker-app.git
cd burung-background-worker-app

cp .env.example .env

go run ./cmd/main.go
```

## Configuration

### Environment Variables

```dotenv
# PostgreSQL Master (write)
DBMASTERHOST=localhost
DBMASTERUSER=postgres
DBMASTERPORT=5432
DBMASTERPASS=your_password
DBMASTERNAME=your_db_name

# PostgreSQL Replica (read)
DBREPLICAHOST=localhost
DBREPLICAUSER=postgres
DBREPLICAPORT=5432
DBREPLICAPASS=your_password
DBREPLICANAME=your_db_name

# Redis
RDSHOST=localhost
RDSPORT=6379
RDSAUTHENTICATION=1        # Redis DB index untuk auth
RDSSESSION=2               # Redis DB index untuk session

# Meilisearch
MEILIHOST=localhost
MEILIPORT=7700
MEILIKEY=your_meili_master_key

# RabbitMQ
RMQ_HOST=localhost
RMQ_USER=your_rmq_user
RMQ_PASS=your_rmq_pass
RMQ_PORT=5672
RMQ_NOTIF_EXCHANGE=notification_burung

# SMTP (Gmail)
CONFIG_SMTP_HOST=smtp.gmail.com
CONFIG_SMTP_PORT=587
CONFIG_SENDER_NAME=App Name <your_email@gmail.com>
CONFIG_AUTH_EMAIL=your_email@gmail.com
CONFIG_AUTH_PASSWORD=your_gmail_app_password   # Google App Password

# Cassandra — Historical DB
CASS_HISTORICAL_SPACEKEY=your_keyspace
CASS_HISTORICAL_USER=cassandra
CASS_HISTORICAL_PASS=your_password
CASS_HISTORICAL_PORT=9042

# Cassandra — SOT Replica Async
CASS_SOT_REPLICA_SPACEKEY=your_keyspace
CASS_SOT_REPLICA_USER=cassandra
CASS_SOT_REPLICA_PASS=your_password
CASS_SOT_REPLICA_PORT=9042
```

> **Jangan commit file `.env` ke repository.** Pastikan sudah ada di `.gitignore`.

```gitignore
.env
```
