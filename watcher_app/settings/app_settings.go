package settings

import "time"

const (
	// =========================================================================
	// 1. TIMEOUTS & DEADLINES (System-wide)
	// =========================================================================
	// Batasan waktu maksimum untuk operasi I/O dan Context di level aplikasi
	TimeoutContext        = 6 * time.Second
	TimeoutDatabaseQuery  = 5 * time.Second
	TimeoutCacheOperation = 2 * time.Second

	// HTTP Server Timeouts (Mencegah Slowloris attack & koneksi menggantung)
	TimeoutServerRead  = 10 * time.Second // Waktu maksimal baca request body
	TimeoutServerWrite = 15 * time.Second // Waktu maksimal tulis response
	TimeoutServerIdle  = 60 * time.Second // Waktu jaga koneksi keep-alive

	// =========================================================================
	// 2. DATABASE & POOLING TUNING (GORM / sql.DB)
	// =========================================================================
	DBMaxOpenConns    = 25               // Maksimal koneksi simultan ke DB
	DBMaxIdleConns    = 5                // Koneksi idle yang tetap dijaga tetap hidup
	DBConnMaxLifetime = 15 * time.Minute // Umur maksimal satu koneksi sebelum di-recycle

	// =========================================================================
	// 3. HTTP & API LIMITATIONS (Payload & Security)
	// =========================================================================
	// Batasan ukuran request body (misal untuk upload foto/file lewat multipart form)
	MaxMultipartMemory = 10 << 20 // 10 MB (dalam bytes)
	MaxRequestBodySize = 2 << 20  // 2 MB untuk request JSON biasa

	// =========================================================================
	// 4. RATE LIMITING PARAMETERS
	// =========================================================================
	// Jika middleware rate limiter kamu di-uncomment nanti
	RateLimitRequestsPerMin = 60
	RateLimitBurst          = 10

	// =========================================================================
	// 5. RETRY POLICIES (Resilience)
	// =========================================================================
	// Batasan mencoba ulang saat koneksi pihak ketiga (MinIO/RabbitMQ) sempat terputus
	MaxRetryAttempts = 3
	RetryBackoffTime = 500 * time.Millisecond
)
