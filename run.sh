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

# 13. Run backend Go application
echo ""
print_success "Semua checks passed! Starting backend..."
echo -e "${GREEN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo ""

# Disable auto-restart for cleaner exit
print_info "Menjalankan backend Go application..."
echo ""

go run main.go
EXIT_CODE=$?

# 14. Exit message
echo ""
if [ $EXIT_CODE -eq 0 ]; then
    print_success "Backend berhenti dengan normal"
else
    print_error "Backend berhenti dengan exit code: $EXIT_CODE"
fi

print_info "Docker containers masih berjalan"
print_info "Gunakan 'docker compose down' untuk menghentikan containers"
print_info "Gunakan 'docker compose logs' untuk melihat logs"

exit $EXIT_CODE