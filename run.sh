#!/bin/bash

# ====================================================
# 🚀 BACKEND STARTUP SCRIPT (Windows Git Bash Compatible)
# ====================================================

# Disable strict error exit for better control
set +e

# Colors for better readability
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
MAX_WAIT_TIME=60
COMPOSE_FILE="docker-compose.yml"
HEALTH_CHECK_INTERVAL=2

# ====================================================
# FUNCTION: Cleanup - matikan semua container saat exit
# ====================================================
CLEANUP_DONE=0
cleanup() {
    if [ "$CLEANUP_DONE" -eq 1 ]; then
        return
    fi
    CLEANUP_DONE=1

    echo ""
    print_warning "Menerima sinyal berhenti, mematikan semua container..."
    docker compose down
    if [ $? -eq 0 ]; then
        print_success "Semua container berhasil dimatikan"
    else
        print_error "Gagal mematikan sebagian/semua container"
    fi
}

trap cleanup EXIT
trap cleanup INT TERM

# ====================================================
# FUNCTION: Print colored messages
# ====================================================
print_error() {
    echo -e "${RED}❌ ERROR: $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

# ====================================================
# FUNCTION: Check if command exists
# ====================================================
command_exists() {
    command -v "$1" >/dev/null 2>&1
    return $?
}

# ====================================================
# FUNCTION: Wait for container to be healthy
# ====================================================
wait_for_containers() {
    local max_wait=$1
    local elapsed=0
    
    print_info "Menunggu containers menjadi healthy..."
    
    # Check if jq is available
    if ! command_exists jq; then
        print_warning "jq tidak terinstall, skip health check"
        sleep 5
        return 0
    fi
    
    while [ $elapsed -lt $max_wait ]; do
        local unhealthy=$(docker compose ps --format json 2>/dev/null | jq -r 'select(.Health != "healthy" and .State == "running") | .Name' 2>/dev/null)
        
        if [ -z "$unhealthy" ]; then
            print_success "Semua containers sudah healthy!"
            return 0
        fi
        
        echo -ne "\r⏳ Menunggu... ${elapsed}s/${max_wait}s"
        sleep $HEALTH_CHECK_INTERVAL
        elapsed=$((elapsed + HEALTH_CHECK_INTERVAL))
    done
    
    echo ""
    print_warning "Timeout menunggu containers. Melanjutkan..."
    return 0
}

# ====================================================
# MAIN SCRIPT
# ====================================================

echo -e "${BLUE}"
echo "==================================="
echo "   🚀 BACKEND STARTUP SCRIPT"
echo "==================================="
echo -e "${NC}"

# 1. Check if Docker is installed
print_info "Memeriksa instalasi Docker..."
if ! command_exists docker; then
    print_error "Docker tidak terinstall!"
    print_info "Install Docker Desktop dari: https://www.docker.com/products/docker-desktop"
    exit 1
fi
print_success "Docker terinstall"

# 2. Check if Docker is running
print_info "Memeriksa Docker daemon..."
if ! docker info > /dev/null 2>&1; then
    print_error "Docker tidak berjalan! Nyalakan Docker Desktop terlebih dahulu."
    exit 1
fi
print_success "Docker daemon berjalan"

# 3. Check if Docker Compose is available
print_info "Memeriksa Docker Compose..."
if ! docker compose version >/dev/null 2>&1; then
    print_error "Docker Compose tidak tersedia!"
    print_info "Pastikan Docker Desktop versi terbaru sudah terinstall"
    exit 1
fi
print_success "Docker Compose tersedia"

# 4. Check if docker-compose.yml exists
print_info "Memeriksa file $COMPOSE_FILE..."
if [ ! -f "$COMPOSE_FILE" ]; then
    print_error "File $COMPOSE_FILE tidak ditemukan!"
    print_info "Pastikan Anda berada di direktori yang benar"
    exit 1
fi
print_success "File $COMPOSE_FILE ditemukan"

# 5. Check if Go is installed
print_info "Memeriksa instalasi Go..."
if ! command_exists go; then
    print_error "Go tidak terinstall!"
    print_info "Install Go dari: https://go.dev/dl/"
    exit 1
fi
GO_VERSION=$(go version 2>/dev/null)
print_success "Go terinstall: $GO_VERSION"

# 6. Check if main.go exists
print_info "Memeriksa file main.go..."
if [ ! -f "main.go" ]; then
    print_error "File main.go tidak ditemukan!"
    print_info "Pastikan Anda berada di direktori project yang benar"
    exit 1
fi
print_success "File main.go ditemukan"

# 7. Check if containers are already running
print_info "Memeriksa status containers..."
RUNNING_CONTAINERS=$(docker compose ps --services --filter "status=running" 2>/dev/null)
if [ -n "$RUNNING_CONTAINERS" ]; then
    print_warning "Containers sudah berjalan!"
    echo "$RUNNING_CONTAINERS"
    echo ""
    read -p "Restart containers? (y/N): " -n 1 -r
    echo ""
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        print_info "Menghentikan containers lama..."
        docker compose down
        if [ $? -eq 0 ]; then
            print_success "Containers berhasil dihentikan"
        else
            print_error "Gagal menghentikan containers"
            exit 1
        fi
    else
        print_info "Menggunakan containers yang sudah berjalan"
    fi
fi

# 8. Start Docker Compose
echo ""
print_info "Menjalankan docker compose..."
docker compose up -d
if [ $? -ne 0 ]; then
    print_error "Gagal menjalankan docker compose!"
    print_info "Coba jalankan: docker compose logs"
    exit 1
fi
print_success "Docker compose berhasil dijalankan"

# 9. Wait for containers to be ready
echo ""
wait_for_containers $MAX_WAIT_TIME

# 10. Show container status
echo ""
print_info "Status containers:"
docker compose ps

# 11. Check for any failed containers
echo ""
print_info "Memeriksa containers yang gagal..."
FAILED_CONTAINERS=$(docker compose ps --format json 2>/dev/null | jq -r 'select(.State != "running") | .Name' 2>/dev/null)
if [ -n "$FAILED_CONTAINERS" ]; then
    print_warning "Beberapa containers gagal start:"
    echo "$FAILED_CONTAINERS"
    print_info "Periksa logs dengan: docker compose logs [container_name]"
else
    print_success "Semua containers berjalan dengan baik"
fi

# ==========================================
# 12. Cek user Cassandra historical
# ==========================================
print_info "Memeriksa user Cassandra historical..."
docker exec -i cassandra_historical_n1 cqlsh -u burung -p burung_secure123 -e "SELECT release_version FROM system.local;" >/dev/null 2>&1

if [ $? -eq 0 ]; then
    print_success "User 'burung' sudah bisa login ke Cassandra Historical, melewati konfigurasi."
else
    print_warning "User 'burung' belum bisa login. Mengaktifkan PasswordAuthenticator..."

    docker exec -i cassandra_historical_n1 bash -c "
        sed -i 's/^\(# *\)\?authenticator:.*/authenticator: PasswordAuthenticator/' /etc/cassandra/cassandra.yaml
        sed -i 's/^\(# *\)\?authorizer:.*/authorizer: CassandraAuthorizer/' /etc/cassandra/cassandra.yaml
    "

    print_info "Restart container cassandra_historical_n1..."
    docker restart cassandra_historical_n1

    print_info "Menunggu Cassandra Historical siap kembali (port 9042)..."
    until docker exec -i cassandra_historical_n1 cqlsh -u cassandra -p cassandra -e "SHOW HOST;" >/dev/null 2>&1; do
        sleep 3
    done

    # jeda ekstra: default superuser butuh waktu sampai system_auth siap
    sleep 10

    print_info "Membuat user 'burung' di Cassandra Historical..."
    CREATE_OK=0
    for i in 1 2 3 4 5; do
        docker exec -i cassandra_historical_n1 cqlsh -u cassandra -p cassandra -e "CREATE ROLE IF NOT EXISTS burung WITH PASSWORD = 'burung_secure123' AND LOGIN = true;" && { CREATE_OK=1; break; }
        sleep 5
    done

    if [ "$CREATE_OK" -ne 1 ]; then
        print_warning "Gagal membuat user 'burung' di Cassandra Historical (CREATE ROLE gagal terus)"
    else
        # ==========================================
        # VERIFIKASI: coba login pake akun burung yang baru dibuat
        # ==========================================
        print_info "Verifikasi login akun 'burung' di Cassandra Historical..."
        LOGIN_OK=0
        for i in 1 2 3 4 5; do
            docker exec -i cassandra_historical_n1 cqlsh -u burung -p burung_secure123 -e "SELECT release_version FROM system.local;" >/dev/null 2>&1 && { LOGIN_OK=1; break; }
            sleep 3
        done

        if [ "$LOGIN_OK" -eq 1 ]; then
            print_success "User 'burung' berhasil dibuat DAN bisa login ke Cassandra Historical"
        else
            print_error "User 'burung' dibuat tapi TETAP GAGAL login. Diagnostic:"
            echo "--- authenticator/authorizer aktif ---"
            docker exec -i cassandra_historical_n1 grep -E "^authenticator|^authorizer" /etc/cassandra/cassandra.yaml
            echo "--- cek role burung ada di system_auth ---"
            docker exec -i cassandra_historical_n1 cqlsh -u cassandra -p cassandra -e "LIST ROLES;"
            echo "--- 30 baris terakhir log cassandra ---"
            docker logs --tail 30 cassandra_historical_n1
        fi
    fi
fi

print_info "Memeriksa keyspace 'historical_db' di Cassandra SOT..."
docker exec -i cassandra_sot_n1 cqlsh -u cassandra -p cassandra -e "CREATE KEYSPACE IF NOT EXISTS historical_db WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};"

if [ $? -eq 0 ]; then
    print_success "Keyspace 'shistorical_db' siap"
    docker exec -i cassandra_sot_n1 cqlsh -u cassandra -p cassandra -e "GRANT ALL PERMISSIONS ON KEYSPACE historical_db TO burung;"
    print_success "Permission 'burung' ke 'historical_db' diberikan"
else
    print_error "Gagal membuat/memastikan keyspace 'sot_replica_db'"
fi

print_info "Memeriksa user Cassandra SOT..."
docker exec -i cassandra_sot_n1 cqlsh 127.0.0.1 9042 -u burung -p burung_secure123 -e "SELECT release_version FROM system.local;" >/dev/null 2>&1

if [ $? -eq 0 ]; then
    print_success "User 'burung' sudah bisa login ke Cassandra SOT, melewati konfigurasi."
else
    print_warning "User 'burung' belum bisa login. Mengaktifkan PasswordAuthenticator..."

    docker exec -i cassandra_sot_n1 bash -c "
        sed -i 's/^\(# *\)\?authenticator:.*/authenticator: PasswordAuthenticator/' /etc/cassandra/cassandra.yaml
        sed -i 's/^\(# *\)\?authorizer:.*/authorizer: CassandraAuthorizer/' /etc/cassandra/cassandra.yaml
    "

    print_info "Restart container cassandra_sot_n1..."
    docker restart cassandra_sot_n1

    print_info "Menunggu Cassandra SOT siap kembali (Host Port: 9043)..."
    until docker exec -i cassandra_sot_n1 cqlsh 127.0.0.1 9042 -u cassandra -p cassandra -e "SHOW HOST;" >/dev/null 2>&1; do
        sleep 3
    done

    sleep 10

    print_info "Membuat user 'burung' di Cassandra SOT..."
    CREATE_OK=0
    for i in 1 2 3 4 5; do
        docker exec -i cassandra_sot_n1 cqlsh -u cassandra -p cassandra -e "CREATE ROLE IF NOT EXISTS burung WITH PASSWORD = 'burung_secure123' AND LOGIN = true;" && { CREATE_OK=1; break; }
        sleep 5
    done

    if [ "$CREATE_OK" -ne 1 ]; then
        print_warning "Gagal membuat user 'burung' di Cassandra SOT (CREATE ROLE gagal terus)"
    else
        print_info "Verifikasi login akun 'burung' di Cassandra SOT..."
        LOGIN_OK=0
        for i in 1 2 3 4 5; do
            docker exec -i cassandra_sot_n1 cqlsh 127.0.0.1 9042 -u burung -p burung_secure123 -e "SELECT release_version FROM system.local;" >/dev/null 2>&1 && { LOGIN_OK=1; break; }
            sleep 3
        done

        if [ "$LOGIN_OK" -eq 1 ]; then
            print_success "User 'burung' berhasil dibuat DAN bisa login ke Cassandra SOT"
        else
            print_error "User 'burung' dibuat tapi TETAP GAGAL login. Diagnostic:"
            echo "--- authenticator/authorizer aktif ---"
            docker exec -i cassandra_sot_n1 grep -E "^authenticator|^authorizer" /etc/cassandra/cassandra.yaml
            echo "--- cek role burung ada di system_auth ---"
            docker exec -i cassandra_sot_n1 cqlsh -u cassandra -p cassandra -e "LIST ROLES;"
            echo "--- 30 baris terakhir log cassandra ---"
            docker logs --tail 30 cassandra_sot_n1
        fi
    fi
fi

print_info "Memeriksa keyspace 'sot_replica_db' di Cassandra SOT..."
docker exec -i cassandra_sot_n1 cqlsh -u cassandra -p cassandra -e "CREATE KEYSPACE IF NOT EXISTS sot_replica_db WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};"

if [ $? -eq 0 ]; then
    print_success "Keyspace 'sot_replica_db' siap"
    docker exec -i cassandra_sot_n1 cqlsh -u cassandra -p cassandra -e "GRANT ALL PERMISSIONS ON KEYSPACE sot_replica_db TO burung;"
    print_success "Permission 'burung' ke 'sot_replica_db' diberikan"
else
    print_error "Gagal membuat/memastikan keyspace 'sot_replica_db'"
fi

 docker exec -i cassandra_historical_n1 cqlsh -u cassandra -p cassandra -e "CREATE KEYSPACE IF NOT EXISTS historical_db WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};"
 sleep 2
 docker exec -i cassandra_sot_n1 cqlsh -u cassandra -p cassandra -e "CREATE KEYSPACE IF NOT EXISTS sot_replica_db WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};"
sleep 30

# 13. Run backend Go application
echo ""
print_success "Semua checks passed! Starting backend..."
echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo ""

# Disable auto-restart for cleaner exit
print_info "Menjalankan backend Go application..."
echo ""

go run main.go --rebootcass=true
EXIT_CODE=$?

# 14. Exit message
# 14. Exit message
echo ""
if [ $EXIT_CODE -eq 0 ]; then
    print_success "Backend berhenti dengan normal"
else
    print_error "Backend berhenti dengan exit code: $EXIT_CODE"
fi

exit $EXIT_CODE