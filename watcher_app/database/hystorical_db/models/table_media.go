package historical_models

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"gorm.io/gorm"
)

type MediaPenggunaProfilFoto struct {
	ID         int64
	IdPengguna int64
	Pengguna   Pengguna
	Key        string
	Format     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	Pencatatan
}

func (MediaPenggunaProfilFoto) TableName() string {
	return "media_pengguna_profil_foto_historical"
}

func (m *MediaPenggunaProfilFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaPenggunaProfilFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		key text,
		format text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaPenggunaProfilFoto) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":           m.ID,
		"id_pengguna":  m.IdPengguna,
		"key":          m.Key,
		"format":       m.Format,
		"created_at":   m.CreatedAt,
		"updated_at":   m.UpdatedAt,
		"deleted_at":   deletedAtInterface,
		"tahun_update": m.TahunUpdate,
		"bulan_update": m.BulanUpdate,
		"event_time":   m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaPenggunaProfilFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaPenggunaProfilFoto) PathName() string {
	return "/media-pengguna-profil-foto/"
}

type MediaSellerProfilFoto struct {
	ID        int64
	IdSeller  int64
	Seller    Seller
	Key       string
	Format    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Pencatatan
}

func (MediaSellerProfilFoto) PathName() string {
	return "/media_seller_profil_foto/"
}

func (m *MediaSellerProfilFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaSellerProfilFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller bigint,
		key text,
		format text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaSellerProfilFoto) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":           m.ID,
		"id_seller":    m.IdSeller,
		"key":          m.Key,
		"format":       m.Format,
		"created_at":   m.CreatedAt,
		"updated_at":   m.UpdatedAt,
		"deleted_at":   deletedAtInterface,
		"tahun_update": m.TahunUpdate,
		"bulan_update": m.BulanUpdate,
		"event_time":   m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaSellerProfilFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaSellerProfilFoto) TableName() string {
	return "media_seller_foto_profil_historical"
}

type MediaSellerBannerFoto struct {
	ID        int64
	IdSeller  int64
	Seller    Seller
	Key       string
	Format    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Pencatatan
}

func (MediaSellerBannerFoto) PathName() string {
	return "/media_seller_banner_foto/"
}

func (m *MediaSellerBannerFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaSellerBannerFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller bigint,
		key text,
		format text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaSellerBannerFoto) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":           m.ID,
		"id_seller":    m.IdSeller,
		"key":          m.Key,
		"format":       m.Format,
		"created_at":   m.CreatedAt,
		"updated_at":   m.UpdatedAt,
		"deleted_at":   deletedAtInterface,
		"tahun_update": m.TahunUpdate,
		"bulan_update": m.BulanUpdate,
		"event_time":   m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaSellerBannerFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaSellerBannerFoto) TableName() string {
	return "media_seller_banner_foto_historical"
}

type MediaSellerTokoFisikFoto struct {
	ID        int64
	IdSeller  int32
	Seller    Seller
	Key       string
	Format    string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
	Pencatatan
}

func (MediaSellerTokoFisikFoto) PathName() string {
	return "/media_seller_toko_fisik_foto/"
}

func (m *MediaSellerTokoFisikFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaSellerTokoFisikFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		key text,
		format text,
		created_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaSellerTokoFisikFoto) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":           m.ID,
		"id_seller":    m.IdSeller,
		"key":          m.Key,
		"format":       m.Format,
		"created_at":   m.CreatedAt,
		"deleted_at":   deletedAtInterface,
		"tahun_update": m.TahunUpdate,
		"bulan_update": m.BulanUpdate,
		"event_time":   m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaSellerTokoFisikFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaSellerTokoFisikFoto) TableName() string {
	return "media_seller_toko_fisik_foto_historical"
}

type MediaKurirProfilFoto struct {
	ID        int64
	IdKurir   int64
	Kurir     Kurir
	Key       string
	Format    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Pencatatan
}

func (MediaKurirProfilFoto) PathName() string {
	return "/media_kurir_profil_foto/"
}

func (m *MediaKurirProfilFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaKurirProfilFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kurir bigint,
		key text,
		format text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaKurirProfilFoto) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":           m.ID,
		"id_kurir":     m.IdKurir,
		"key":          m.Key,
		"format":       m.Format,
		"created_at":   m.CreatedAt,
		"updated_at":   m.UpdatedAt,
		"deleted_at":   deletedAtInterface,
		"tahun_update": m.TahunUpdate,
		"bulan_update": m.BulanUpdate,
		"event_time":   m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaKurirProfilFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaKurirProfilFoto) TableName() string {
	return "media_kurir_profil_foto_historical"
}

type MediaEtalaseFoto struct {
	ID        int64
	IdEtalase int64
	Etalase   Etalase
	Key       string
	Format    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Pencatatan
}

func (MediaEtalaseFoto) PathName() string {
	return "/media_etalase_foto/"
}

func (m *MediaEtalaseFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaEtalaseFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_etalase bigint,
		key text,
		format text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaEtalaseFoto) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":           m.ID,
		"id_etalase":   m.IdEtalase,
		"key":          m.Key,
		"format":       m.Format,
		"created_at":   m.CreatedAt,
		"updated_at":   m.UpdatedAt,
		"deleted_at":   deletedAtInterface,
		"tahun_update": m.TahunUpdate,
		"bulan_update": m.BulanUpdate,
		"event_time":   m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaEtalaseFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaEtalaseFoto) TableName() string {
	return "media_etalase_foto_historical"
}

type MediaBarangIndukFoto struct {
	ID            int64
	IdBarangInduk int64
	BarangInduk   BarangInduk
	Key           string
	Format        string
	CreatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	Pencatatan
}

func (MediaBarangIndukFoto) PathName() string {
	return "/media_barang_induk_foto/"
}

func (m *MediaBarangIndukFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBarangIndukFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_barang_induk bigint,
		key text,
		format text,
		created_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaBarangIndukFoto) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":              m.ID,
		"id_barang_induk": m.IdBarangInduk,
		"key":             m.Key,
		"format":          m.Format,
		"created_at":      m.CreatedAt,
		"deleted_at":      deletedAtInterface,
		"tahun_update":    m.TahunUpdate,
		"bulan_update":    m.BulanUpdate,
		"event_time":      m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaBarangIndukFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaBarangIndukFoto) TableName() string {
	return "media_barang_induk_foto_historical"
}

type MediaBarangIndukVideo struct {
	ID            int64
	IdBarangInduk int64
	BarangInduk   BarangInduk
	Key           string
	Format        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	Pencatatan
}

func (MediaBarangIndukVideo) PathName() string {
	return "/media_barang_induk_video/"
}

func (m *MediaBarangIndukVideo) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBarangIndukVideo dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_barang_induk bigint,
		key text,
		format text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaBarangIndukVideo) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":              m.ID,
		"id_barang_induk": m.IdBarangInduk,
		"key":             m.Key,
		"format":          m.Format,
		"created_at":      m.CreatedAt,
		"updated_at":      m.UpdatedAt,
		"deleted_at":      deletedAtInterface,
		"tahun_update":    m.TahunUpdate,
		"bulan_update":    m.BulanUpdate,
		"event_time":      m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaBarangIndukVideo) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaBarangIndukVideo) TableName() string {
	return "media_barang_induk_video_historical"
}

type MediaKategoriBarangFoto struct {
	ID               int64
	IdKategoriBarang int64
	KategoriBarang   KategoriBarang
	IdBarangInduk    int64
	BarangInduk      BarangInduk
	Key              string
	Format           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	Pencatatan
}

func (MediaKategoriBarangFoto) PathName() string {
	return "/media_kategori_barang_foto/"
}

func (m *MediaKategoriBarangFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaKategoriBarangFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kategori_barang bigint,
		id_barang_induk bigint,
		key text,
		format text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaKategoriBarangFoto) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":                 m.ID,
		"id_kategori_barang": m.IdKategoriBarang,
		"id_barang_induk":    m.IdBarangInduk,
		"key":                m.Key,
		"format":             m.Format,
		"created_at":         m.CreatedAt,
		"updated_at":         m.UpdatedAt,
		"deleted_at":         deletedAtInterface,
		"tahun_update":       m.TahunUpdate,
		"bulan_update":       m.BulanUpdate,
		"event_time":         m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaKategoriBarangFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaKategoriBarangFoto) TableName() string {
	return "media_kategori_barang_foto_historical"
}

type MediaDistributorDataDokumen struct {
	ID                int64
	IdDistributorData int64
	DistributorData   DistributorData
	Key               string
	Format            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
	Pencatatan
}

func (MediaDistributorDataDokumen) PathName() string {
	return "/media_distributor_data_dokumen/"
}

func (m *MediaDistributorDataDokumen) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaDistributorDataDokumen dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_distributor_data bigint,
		key text,
		format text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaDistributorDataDokumen) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":                  m.ID,
		"id_distributor_data": m.IdDistributorData,
		"key":                 m.Key,
		"format":              m.Format,
		"created_at":          m.CreatedAt,
		"updated_at":          m.UpdatedAt,
		"deleted_at":          deletedAtInterface,
		"tahun_update":        m.TahunUpdate,
		"bulan_update":        m.BulanUpdate,
		"event_time":          m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaDistributorDataDokumen) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaDistributorDataDokumen) TableName() string {
	return "media_distributor_data_dokumen_historical"
}

type MediaDistributorDataNPWPFoto struct {
	ID                int64
	IdDistributorData int64
	DistributorData   DistributorData
	Key               string
	Format            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
	Pencatatan
}

func (MediaDistributorDataNPWPFoto) PathName() string {
	return "/media_distributor_data_npwp_foto/"
}

func (m *MediaDistributorDataNPWPFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaDistributorDataNPWPFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_distributor_data bigint,
		key text,
		format text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaDistributorDataNPWPFoto) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":                  m.ID,
		"id_distributor_data": m.IdDistributorData,
		"key":                 m.Key,
		"format":              m.Format,
		"created_at":          m.CreatedAt,
		"updated_at":          m.UpdatedAt,
		"deleted_at":          deletedAtInterface,
		"tahun_update":        m.TahunUpdate,
		"bulan_update":        m.BulanUpdate,
		"event_time":          m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaDistributorDataNPWPFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaDistributorDataNPWPFoto) TableName() string {
	return "media_distributor_data_npwp_foto_historical"
}

type MediaDistributorDataNIBFoto struct {
	ID                int64
	IdDistributorData int64
	DistributorData   DistributorData
	Key               string
	Format            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
	Pencatatan
}

func (MediaDistributorDataNIBFoto) PathName() string {
	return "/media_distributor_data_nib_foto/"
}

func (m *MediaDistributorDataNIBFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaDistributorDataNIBFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_distributor_data bigint,
		key text,
		format text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaDistributorDataNIBFoto) ParseToInsertType() map[string]interface{} {
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":                  m.ID,
		"id_distributor_data": m.IdDistributorData,
		"key":                 m.Key,
		"format":              m.Format,
		"created_at":          m.CreatedAt,
		"updated_at":          m.UpdatedAt,
		"deleted_at":          deletedAtInterface,
		"tahun_update":        m.TahunUpdate,
		"bulan_update":        m.BulanUpdate,
		"event_time":          m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaDistributorDataNIBFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaDistributorDataNIBFoto) TableName() string {
	return "media_distributor_data_nib_foto_historical"
}

type MediaDistributorDataSuratKerjasamaDokumen struct {
	ID                int64
	IdDistributorData int64
	DistributorData   DistributorData
	Key               string
	Format            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
	Pencatatan
}

func (MediaDistributorDataSuratKerjasamaDokumen) PathName() string {
	return "/media_distributor_data_surat_kerjasama_dokumen/"
}

func (m *MediaDistributorDataSuratKerjasamaDokumen) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaDistributorDataSuratKerjasamaDokumen dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_distributor_data bigint,
		key text,
		format text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaDistributorDataSuratKerjasamaDokumen) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":                  m.ID,
		"id_distributor_data": m.IdDistributorData,
		"key":                 m.Key,
		"format":              m.Format,
		"created_at":          m.CreatedAt,
		"updated_at":          m.UpdatedAt,
		"deleted_at":          deletedAtInterface,
		"tahun_update":        m.TahunUpdate,
		"bulan_update":        m.BulanUpdate,
		"event_time":          m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaDistributorDataSuratKerjasamaDokumen) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaDistributorDataSuratKerjasamaDokumen) TableName() string {
	return "media_distributor_data_surat_kerjasama_dokumen_historical"
}

type MediaBrandDataPerwakilanDokumen struct {
	ID          int64
	IdBrandData int64
	BrandData   BrandData
	Key         string
	Format      string
	Pencatatan
}

func (MediaBrandDataPerwakilanDokumen) PathName() string {
	return "/media_brand_data_perwakilan_dokumen/"
}

func (m *MediaBrandDataPerwakilanDokumen) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBrandDataPerwakilanDokumen dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_brand_data bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaBrandDataPerwakilanDokumen) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.ID,
		"id_brand_data": m.IdBrandData,
		"key":           m.Key,
		"format":        m.Format,
		"tahun_update":  m.TahunUpdate,
		"bulan_update":  m.BulanUpdate,
		"event_time":    m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaBrandDataPerwakilanDokumen) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaBrandDataPerwakilanDokumen) TableName() string {
	return "media_brand_data_perwakilan_dokumen_historical"
}

type MediaBrandDataSertifikatFoto struct {
	ID          int64
	IdBrandData int64
	BrandData   BrandData
	Key         string
	Format      string
	Pencatatan
}

func (MediaBrandDataSertifikatFoto) PathName() string {
	return "/media_brand_data_sertifikat_foto/"
}

func (m *MediaBrandDataSertifikatFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBrandDataSertifikatFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_brand_data bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaBrandDataSertifikatFoto) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.ID,
		"id_brand_data": m.IdBrandData,
		"key":           m.Key,
		"format":        m.Format,
		"tahun_update":  m.TahunUpdate,
		"bulan_update":  m.BulanUpdate,
		"event_time":    m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaBrandDataSertifikatFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaBrandDataSertifikatFoto) TableName() string {
	return "media_brand_data_sertifikat_foto_historical"
}

type MediaBrandDataNIBFoto struct {
	ID          int64
	IdBrandData int64
	BrandData   BrandData
	Key         string
	Format      string
	Pencatatan
}

func (MediaBrandDataNIBFoto) PathName() string {
	return "/media_brand_data_nib_foto/"
}

func (m *MediaBrandDataNIBFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBrandDataNIBFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_brand_data bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaBrandDataNIBFoto) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.ID,
		"id_brand_data": m.IdBrandData,
		"key":           m.Key,
		"format":        m.Format,
		"tahun_update":  m.TahunUpdate,
		"bulan_update":  m.BulanUpdate,
		"event_time":    m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaBrandDataNIBFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaBrandDataNIBFoto) TableName() string {
	return "media_brand_data_nib_foto_historical"
}

type MediaBrandDataNPWPFoto struct {
	ID          int64
	IdBrandData int64
	BrandData   BrandData
	Key         string
	Format      string
	Pencatatan
}

func (MediaBrandDataNPWPFoto) PathName() string {
	return "/media_brand_data_npwp_foto/"
}

func (m *MediaBrandDataNPWPFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBrandDataNPWPFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_brand_data bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaBrandDataNPWPFoto) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.ID,
		"id_brand_data": m.IdBrandData,
		"key":           m.Key,
		"format":        m.Format,
		"tahun_update":  m.TahunUpdate,
		"bulan_update":  m.BulanUpdate,
		"event_time":    m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaBrandDataNPWPFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaBrandDataNPWPFoto) TableName() string {
	return "media_brand_data_npwp_foto_historical"
}

type MediaBrandDataLogoFoto struct {
	ID          int64
	IdBrandData int64
	BrandData   BrandData
	Key         string
	Format      string
	Pencatatan
}

func (MediaBrandDataLogoFoto) PathName() string {
	return "/media_brand_data_logo_foto/"
}

func (m *MediaBrandDataLogoFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBrandDataLogoFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_brand_data bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaBrandDataLogoFoto) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.ID,
		"id_brand_data": m.IdBrandData,
		"key":           m.Key,
		"format":        m.Format,
		"tahun_update":  m.TahunUpdate,
		"bulan_update":  m.BulanUpdate,
		"event_time":    m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaBrandDataLogoFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaBrandDataLogoFoto) TableName() string {
	return "media_brand_data_logo_foto_historical"
}

type MediaBrandDataSuratKerjasamaDokumen struct {
	ID        int64
	BrandData BrandData
	Key       string
	Format    string
	Pencatatan
}

func (MediaBrandDataSuratKerjasamaDokumen) PathName() string {
	return "/media_brand_data_surat_kerjasama_dokumen/"
}

func (m *MediaBrandDataSuratKerjasamaDokumen) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBrandDataSuratKerjasamaDokumen dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaBrandDataSuratKerjasamaDokumen) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":           m.ID,
		"key":          m.Key,
		"format":       m.Format,
		"tahun_update": m.TahunUpdate,
		"bulan_update": m.BulanUpdate,
		"event_time":   m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaBrandDataSuratKerjasamaDokumen) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaBrandDataSuratKerjasamaDokumen) TableName() string {
	return "media_brand_data_surat_kerjasama_dokumen_historical"
}

type MediaInformasiKendaraanKurirKendaraanFoto struct {
	ID                        int64
	IdInformasiKendaraanKurir int64
	InformasiKendaraanKurir   InformasiKendaraanKurir
	Key                       string
	Format                    string
	Pencatatan
}

func (MediaInformasiKendaraanKurirKendaraanFoto) PathName() string {
	return "/media_informasi_kendaraan_kurir_kendaraan_foto/"
}

func (m *MediaInformasiKendaraanKurirKendaraanFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaInformasiKendaraanKurirKendaraanFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_informasi_kendaraan_kurir bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaInformasiKendaraanKurirKendaraanFoto) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":                           m.ID,
		"id_informasi_kendaraan_kurir": m.IdInformasiKendaraanKurir,
		"key":                          m.Key,
		"format":                       m.Format,
		"tahun_update":                 m.TahunUpdate,
		"bulan_update":                 m.BulanUpdate,
		"event_time":                   m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaInformasiKendaraanKurirKendaraanFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaInformasiKendaraanKurirKendaraanFoto) TableName() string {
	return "media_informasi_kendaraan_kurir_kendaraan_foto_historical"
}

type MediaInformasiKendaraanKurirBPKBFoto struct {
	ID                        int64
	IdInformasiKendaraanKurir int64
	InformasiKendaraanKurir   InformasiKendaraanKurir
	Key                       string
	Format                    string
	Pencatatan
}

func (MediaInformasiKendaraanKurirBPKBFoto) PathName() string {
	return "/media_informasi_kendaraan_kurir_bpkb_foto/"

}

func (m *MediaInformasiKendaraanKurirBPKBFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaInformasiKendaraanKurirBPKBFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_informasi_kendaraan_kurir bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaInformasiKendaraanKurirBPKBFoto) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":                           m.ID,
		"id_informasi_kendaraan_kurir": m.IdInformasiKendaraanKurir,
		"key":                          m.Key,
		"format":                       m.Format,
		"tahun_update":                 m.TahunUpdate,
		"bulan_update":                 m.BulanUpdate,
		"event_time":                   m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaInformasiKendaraanKurirBPKBFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaInformasiKendaraanKurirBPKBFoto) TableName() string {
	return "media_informasi_kendaraan_kurir_bpkb_foto_historical"
}

type MediaInformasiKendaraanKurirSTNKFoto struct {
	ID                        int64
	IdInformasiKendaraanKurir int64
	InformasiKendaraanKurir   InformasiKendaraanKurir
	Key                       string
	Format                    string
	Pencatatan
}

func (MediaInformasiKendaraanKurirSTNKFoto) PathName() string {
	return "/media_informasi_kendaraan_kurir_stnk_foto/"
}

func (m *MediaInformasiKendaraanKurirSTNKFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaInformasiKendaraanKurirSTNKFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_informasi_kendaraan_kurir bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaInformasiKendaraanKurirSTNKFoto) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":                           m.ID,
		"id_informasi_kendaraan_kurir": m.IdInformasiKendaraanKurir,
		"key":                          m.Key,
		"format":                       m.Format,
		"tahun_update":                 m.TahunUpdate,
		"bulan_update":                 m.BulanUpdate,
		"event_time":                   m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaInformasiKendaraanKurirSTNKFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaInformasiKendaraanKurirSTNKFoto) TableName() string {
	return "media_informasi_kendaraan_kurir_stnk_foto_historical"
}

type MediaInformasiKurirKTPFoto struct {
	ID               int64
	IdInformasiKurir int64
	InformasiKurir   InformasiKurir
	Key              string
	Format           string
	Pencatatan
}

func (MediaInformasiKurirKTPFoto) PathName() string {
	return "/media_informasi_kurir_ktp_foto/"
}

func (m *MediaInformasiKurirKTPFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaInformasiKurirKTPFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_informasi_kurir bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaInformasiKurirKTPFoto) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":                 m.ID,
		"id_informasi_kurir": m.IdInformasiKurir,
		"key":                m.Key,
		"format":             m.Format,
		"tahun_update":       m.TahunUpdate,
		"bulan_update":       m.BulanUpdate,
		"event_time":         m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaInformasiKurirKTPFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaInformasiKurirKTPFoto) TableName() string {
	return "media_informasi_kurir_ktp_foto_historical"
}

type MediaReviewFoto struct {
	ID       int64
	IdReview int64
	Review   Review
	Key      string
	Format   string
	Pencatatan
}

func (MediaReviewFoto) PathName() string {
	return "/media_review_foto/"
}

func (m *MediaReviewFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaReviewFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_review bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaReviewFoto) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":           m.ID,
		"id_review":    m.IdReview,
		"key":          m.Key,
		"format":       m.Format,
		"tahun_update": m.TahunUpdate,
		"bulan_update": m.BulanUpdate,
		"event_time":   m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaReviewFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaReviewFoto) TableName() string {
	return "media_review_foto_historical"
}

type MediaReviewVideo struct {
	ID       int64
	IdReview int64
	Review   Review
	Key      string
	Format   string
	Pencatatan
}

func (MediaReviewVideo) PathName() string {
	return "/media_review_video/"
}

func (m *MediaReviewVideo) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaReviewVideo dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_review bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaReviewVideo) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":           m.ID,
		"id_review":    m.IdReview,
		"key":          m.Key,
		"format":       m.Format,
		"tahun_update": m.TahunUpdate,
		"bulan_update": m.BulanUpdate,
		"event_time":   m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaReviewVideo) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaReviewVideo) TableName() string {
	return "media_review_video_historical"
}

type MediaTransaksiApprovedFoto struct {
	ID          int64
	IdTransaksi int64
	Transaksi   Transaksi
	Key         string
	Format      string
	Pencatatan
}

func (MediaTransaksiApprovedFoto) PathName() string {
	return "/media_transaksi_approved_foto/"
}

func (m *MediaTransaksiApprovedFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaTransaksiApprovedFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_transaksi bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaTransaksiApprovedFoto) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":           m.ID,
		"id_transaksi": m.IdTransaksi,
		"key":          m.Key,
		"format":       m.Format,
		"tahun_update": m.TahunUpdate,
		"bulan_update": m.BulanUpdate,
		"event_time":   m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaTransaksiApprovedFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaTransaksiApprovedFoto) TableName() string {
	return "media_transaksi_approved_foto_historical"
}

type MediaTransaksiApprovedVideo struct {
	ID          int64
	IdTransaksi int64
	Transaksi   Transaksi
	Key         string
	Format      string
	Pencatatan
}

func (MediaTransaksiApprovedVideo) PathName() string {
	return "/media_transaksi_approved_video/"
}

func (m *MediaTransaksiApprovedVideo) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaTransaksiApprovedVideo dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_transaksi bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaTransaksiApprovedVideo) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":           m.ID,
		"id_transaksi": m.IdTransaksi,
		"key":          m.Key,
		"format":       m.Format,
		"tahun_update": m.TahunUpdate,
		"bulan_update": m.BulanUpdate,
		"event_time":   m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaTransaksiApprovedVideo) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaTransaksiApprovedVideo) TableName() string {
	return "media_transaksi_approved_video_historical"
}

type MediaPengirimanPickedUpFoto struct {
	ID           int64
	IdPengiriman int64
	Pengiriman   Pengiriman
	Key          string
	Format       string
	Pencatatan
}

func (MediaPengirimanPickedUpFoto) PathName() string {
	return "/media_pengiriman_picked_up_foto/"
}

func (m *MediaPengirimanPickedUpFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaPengirimanPickedUpFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengiriman bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaPengirimanPickedUpFoto) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.ID,
		"id_pengiriman": m.IdPengiriman,
		"key":           m.Key,
		"format":        m.Format,
		"tahun_update":  m.TahunUpdate,
		"bulan_update":  m.BulanUpdate,
		"event_time":    m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaPengirimanPickedUpFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaPengirimanPickedUpFoto) TableName() string {
	return "media_pengiriman_picked_up_foto_historical"
}

type MediaPengirimanSampaiFoto struct {
	ID           int64
	IdPengiriman int64
	Pengiriman   Pengiriman
	Key          string
	Format       string
	Pencatatan
}

func (MediaPengirimanSampaiFoto) PathName() string {
	return "/media_pengiriman_sampai_foto/"

}

func (m *MediaPengirimanSampaiFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaPengirimanSampaiFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengiriman bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaPengirimanSampaiFoto) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.ID,
		"id_pengiriman": m.IdPengiriman,
		"key":           m.Key,
		"format":        m.Format,
		"tahun_update":  m.TahunUpdate,
		"bulan_update":  m.BulanUpdate,
		"event_time":    m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaPengirimanSampaiFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaPengirimanSampaiFoto) TableName() string {
	return "media_pengiriman_sampai_foto_historical"
}

type MediaPengirimanEkspedisiPickedUpFoto struct {
	ID                    int64
	IdPengirimanEkspedisi int64
	PengirimanEkspedisi   PengirimanEkspedisi
	Key                   string
	Format                string
	Pencatatan
}

func (MediaPengirimanEkspedisiPickedUpFoto) PathName() string {
	return "/media_pengiriman_ekspedisi_picked_up_foto/"
}

func (m *MediaPengirimanEkspedisiPickedUpFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaPengirimanEkspedisiPickedUpFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengiriman_ekspedisi bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaPengirimanEkspedisiPickedUpFoto) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":                      m.ID,
		"id_pengiriman_ekspedisi": m.IdPengirimanEkspedisi,
		"key":                     m.Key,
		"format":                  m.Format,
		"tahun_update":            m.TahunUpdate,
		"bulan_update":            m.BulanUpdate,
		"event_time":              m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaPengirimanEkspedisiPickedUpFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaPengirimanEkspedisiPickedUpFoto) TableName() string {
	return "media_pengiriman_ekspedisi_picked_up_foto_historical"
}

type MediaPengirimanEkspedisiSampaiAgentFoto struct {
	ID                    int64
	IdPengirimanEkspedisi int64
	PengirimanEkspedisi   PengirimanEkspedisi
	Key                   string
	Format                string
	Pencatatan
}

func (MediaPengirimanEkspedisiSampaiAgentFoto) PathName() string {
	return "/media_pengiriman_ekspedisi_sampai_agent_foto/"
}

func (m *MediaPengirimanEkspedisiSampaiAgentFoto) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaPengirimanEkspedisiSampaiAgentFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengiriman_ekspedisi bigint,
		key text,
		format text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableName())
	return nil
}

func (m *MediaPengirimanEkspedisiSampaiAgentFoto) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":                      m.ID,
		"id_pengiriman_ekspedisi": m.IdPengirimanEkspedisi,
		"key":                     m.Key,
		"format":                  m.Format,
		"tahun_update":            m.TahunUpdate,
		"bulan_update":            m.BulanUpdate,
		"event_time":              m.EventTime,
	}
}

// DropTable disesuaikan menggunakan m.TableName() secara dinamis
func (m *MediaPengirimanEkspedisiSampaiAgentFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableName())
	return nil
}

func (MediaPengirimanEkspedisiSampaiAgentFoto) TableName() string {
	return "media_pengiriman_ekspedisi_sampai_agent_foto_historical"
}
