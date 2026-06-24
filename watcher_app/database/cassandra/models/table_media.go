package cass_models

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
}

func (MediaPenggunaProfilFoto) TableNameHistorical() string {
	return "media_pengguna_profil_foto_historical"
}

func (m *MediaPenggunaProfilFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaPenggunaProfilFoto) ParseToCUDType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":          m.ID,
		"id_pengguna": m.IdPengguna,
		"key":         m.Key,
		"format":      m.Format,
		"created_at":  m.CreatedAt,
		"updated_at":  m.UpdatedAt,
		"deleted_at":  deletedAtInterface,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaPenggunaProfilFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
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
}

func (MediaSellerProfilFoto) PathName() string {
	return "/media_seller_profil_foto/"
}

func (m *MediaSellerProfilFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaSellerProfilFoto) ParseToCUDType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":         m.ID,
		"id_seller":  m.IdSeller,
		"key":        m.Key,
		"format":     m.Format,
		"created_at": m.CreatedAt,
		"updated_at": m.UpdatedAt,
		"deleted_at": deletedAtInterface,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaSellerProfilFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaSellerProfilFoto) TableNameHistorical() string {
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
}

func (MediaSellerBannerFoto) PathName() string {
	return "/media_seller_banner_foto/"
}

func (m *MediaSellerBannerFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaSellerBannerFoto) ParseToCUDType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":         m.ID,
		"id_seller":  m.IdSeller,
		"key":        m.Key,
		"format":     m.Format,
		"created_at": m.CreatedAt,
		"updated_at": m.UpdatedAt,
		"deleted_at": deletedAtInterface,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaSellerBannerFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaSellerBannerFoto) TableNameHistorical() string {
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
}

func (MediaSellerTokoFisikFoto) PathName() string {
	return "/media_seller_toko_fisik_foto/"
}

func (m *MediaSellerTokoFisikFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaSellerTokoFisikFoto) ParseToCUDType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":         m.ID,
		"id_seller":  m.IdSeller,
		"key":        m.Key,
		"format":     m.Format,
		"created_at": m.CreatedAt,
		"deleted_at": deletedAtInterface,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaSellerTokoFisikFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaSellerTokoFisikFoto) TableNameHistorical() string {
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
}

func (MediaKurirProfilFoto) PathName() string {
	return "/media_kurir_profil_foto/"
}

func (m *MediaKurirProfilFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaKurirProfilFoto) ParseToCUDType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":         m.ID,
		"id_kurir":   m.IdKurir,
		"key":        m.Key,
		"format":     m.Format,
		"created_at": m.CreatedAt,
		"updated_at": m.UpdatedAt,
		"deleted_at": deletedAtInterface,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaKurirProfilFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaKurirProfilFoto) TableNameHistorical() string {
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
}

func (MediaEtalaseFoto) PathName() string {
	return "/media_etalase_foto/"
}

func (m *MediaEtalaseFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaEtalaseFoto) ParseToCUDType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if m.DeletedAt.Valid {
		deletedAtInterface = m.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":         m.ID,
		"id_etalase": m.IdEtalase,
		"key":        m.Key,
		"format":     m.Format,
		"created_at": m.CreatedAt,
		"updated_at": m.UpdatedAt,
		"deleted_at": deletedAtInterface,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaEtalaseFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaEtalaseFoto) TableNameHistorical() string {
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
}

func (MediaBarangIndukFoto) PathName() string {
	return "/media_barang_induk_foto/"
}

func (m *MediaBarangIndukFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaBarangIndukFoto) ParseToCUDType() map[string]interface{} {
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
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaBarangIndukFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaBarangIndukFoto) TableNameHistorical() string {
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
}

func (MediaBarangIndukVideo) PathName() string {
	return "/media_barang_induk_video/"
}

func (m *MediaBarangIndukVideo) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaBarangIndukVideo) ParseToCUDType() map[string]interface{} {
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
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaBarangIndukVideo) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaBarangIndukVideo) TableNameHistorical() string {
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
}

func (MediaKategoriBarangFoto) PathName() string {
	return "/media_kategori_barang_foto/"
}

func (m *MediaKategoriBarangFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaKategoriBarangFoto) ParseToCUDType() map[string]interface{} {
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
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaKategoriBarangFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaKategoriBarangFoto) TableNameHistorical() string {
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
}

func (MediaDistributorDataDokumen) PathName() string {
	return "/media_distributor_data_dokumen/"
}

func (m *MediaDistributorDataDokumen) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaDistributorDataDokumen) ParseToCUDType() map[string]interface{} {
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
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaDistributorDataDokumen) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaDistributorDataDokumen) TableNameHistorical() string {
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
}

func (MediaDistributorDataNPWPFoto) PathName() string {
	return "/media_distributor_data_npwp_foto/"
}

func (m *MediaDistributorDataNPWPFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaDistributorDataNPWPFoto) ParseToCUDType() map[string]interface{} {
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
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaDistributorDataNPWPFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaDistributorDataNPWPFoto) TableNameHistorical() string {
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
}

func (MediaDistributorDataNIBFoto) PathName() string {
	return "/media_distributor_data_nib_foto/"
}

func (m *MediaDistributorDataNIBFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaDistributorDataNIBFoto) ParseToCUDType() map[string]interface{} {
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
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaDistributorDataNIBFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaDistributorDataNIBFoto) TableNameHistorical() string {
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
}

func (MediaDistributorDataSuratKerjasamaDokumen) PathName() string {
	return "/media_distributor_data_surat_kerjasama_dokumen/"
}

func (m *MediaDistributorDataSuratKerjasamaDokumen) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaDistributorDataSuratKerjasamaDokumen) ParseToCUDType() map[string]interface{} {
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
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaDistributorDataSuratKerjasamaDokumen) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaDistributorDataSuratKerjasamaDokumen) TableNameHistorical() string {
	return "media_distributor_data_surat_kerjasama_dokumen_historical"
}

type MediaBrandDataPerwakilanDokumen struct {
	ID          int64
	IdBrandData int64
	BrandData   BrandData
	Key         string
	Format      string
}

func (MediaBrandDataPerwakilanDokumen) PathName() string {
	return "/media_brand_data_perwakilan_dokumen/"
}

func (m *MediaBrandDataPerwakilanDokumen) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaBrandDataPerwakilanDokumen) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.ID,
		"id_brand_data": m.IdBrandData,
		"key":           m.Key,
		"format":        m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaBrandDataPerwakilanDokumen) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaBrandDataPerwakilanDokumen) TableNameHistorical() string {
	return "media_brand_data_perwakilan_dokumen_historical"
}

type MediaBrandDataSertifikatFoto struct {
	ID          int64
	IdBrandData int64
	BrandData   BrandData
	Key         string
	Format      string
}

func (MediaBrandDataSertifikatFoto) PathName() string {
	return "/media_brand_data_sertifikat_foto/"
}

func (m *MediaBrandDataSertifikatFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaBrandDataSertifikatFoto) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.ID,
		"id_brand_data": m.IdBrandData,
		"key":           m.Key,
		"format":        m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaBrandDataSertifikatFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaBrandDataSertifikatFoto) TableNameHistorical() string {
	return "media_brand_data_sertifikat_foto_historical"
}

type MediaBrandDataNIBFoto struct {
	ID          int64
	IdBrandData int64
	BrandData   BrandData
	Key         string
	Format      string
}

func (MediaBrandDataNIBFoto) PathName() string {
	return "/media_brand_data_nib_foto/"
}

func (m *MediaBrandDataNIBFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaBrandDataNIBFoto) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.ID,
		"id_brand_data": m.IdBrandData,
		"key":           m.Key,
		"format":        m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaBrandDataNIBFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaBrandDataNIBFoto) TableNameHistorical() string {
	return "media_brand_data_nib_foto_historical"
}

type MediaBrandDataNPWPFoto struct {
	ID          int64
	IdBrandData int64
	BrandData   BrandData
	Key         string
	Format      string
}

func (MediaBrandDataNPWPFoto) PathName() string {
	return "/media_brand_data_npwp_foto/"
}

func (m *MediaBrandDataNPWPFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaBrandDataNPWPFoto) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.ID,
		"id_brand_data": m.IdBrandData,
		"key":           m.Key,
		"format":        m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaBrandDataNPWPFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaBrandDataNPWPFoto) TableNameHistorical() string {
	return "media_brand_data_npwp_foto_historical"
}

type MediaBrandDataLogoFoto struct {
	ID          int64
	IdBrandData int64
	BrandData   BrandData
	Key         string
	Format      string
}

func (MediaBrandDataLogoFoto) PathName() string {
	return "/media_brand_data_logo_foto/"
}

func (m *MediaBrandDataLogoFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaBrandDataLogoFoto) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.ID,
		"id_brand_data": m.IdBrandData,
		"key":           m.Key,
		"format":        m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaBrandDataLogoFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaBrandDataLogoFoto) TableNameHistorical() string {
	return "media_brand_data_logo_foto_historical"
}

type MediaBrandDataSuratKerjasamaDokumen struct {
	ID          int64
	IdBrandData int64
	BrandData   BrandData
	Key         string
	Format      string
}

func (MediaBrandDataSuratKerjasamaDokumen) PathName() string {
	return "/media_brand_data_surat_kerjasama_dokumen/"
}

func (m *MediaBrandDataSuratKerjasamaDokumen) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaBrandDataSuratKerjasamaDokumen) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":     m.ID,
		"key":    m.Key,
		"format": m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaBrandDataSuratKerjasamaDokumen) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaBrandDataSuratKerjasamaDokumen) TableNameHistorical() string {
	return "media_brand_data_surat_kerjasama_dokumen_historical"
}

type MediaInformasiKendaraanKurirKendaraanFoto struct {
	ID                        int64
	IdInformasiKendaraanKurir int64
	InformasiKendaraanKurir   InformasiKendaraanKurir
	Key                       string
	Format                    string
}

func (MediaInformasiKendaraanKurirKendaraanFoto) PathName() string {
	return "/media_informasi_kendaraan_kurir_kendaraan_foto/"
}

func (m *MediaInformasiKendaraanKurirKendaraanFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaInformasiKendaraanKurirKendaraanFoto) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                           m.ID,
		"id_informasi_kendaraan_kurir": m.IdInformasiKendaraanKurir,
		"key":                          m.Key,
		"format":                       m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaInformasiKendaraanKurirKendaraanFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaInformasiKendaraanKurirKendaraanFoto) TableNameHistorical() string {
	return "media_informasi_kendaraan_kurir_kendaraan_foto_historical"
}

type MediaInformasiKendaraanKurirBPKBFoto struct {
	ID                        int64
	IdInformasiKendaraanKurir int64
	InformasiKendaraanKurir   InformasiKendaraanKurir
	Key                       string
	Format                    string
}

func (MediaInformasiKendaraanKurirBPKBFoto) PathName() string {
	return "/media_informasi_kendaraan_kurir_bpkb_foto/"

}

func (m *MediaInformasiKendaraanKurirBPKBFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaInformasiKendaraanKurirBPKBFoto) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                           m.ID,
		"id_informasi_kendaraan_kurir": m.IdInformasiKendaraanKurir,
		"key":                          m.Key,
		"format":                       m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaInformasiKendaraanKurirBPKBFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaInformasiKendaraanKurirBPKBFoto) TableNameHistorical() string {
	return "media_informasi_kendaraan_kurir_bpkb_foto_historical"
}

type MediaInformasiKendaraanKurirSTNKFoto struct {
	ID                        int64
	IdInformasiKendaraanKurir int64
	InformasiKendaraanKurir   InformasiKendaraanKurir
	Key                       string
	Format                    string
}

func (MediaInformasiKendaraanKurirSTNKFoto) PathName() string {
	return "/media_informasi_kendaraan_kurir_stnk_foto/"
}

func (m *MediaInformasiKendaraanKurirSTNKFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaInformasiKendaraanKurirSTNKFoto) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                           m.ID,
		"id_informasi_kendaraan_kurir": m.IdInformasiKendaraanKurir,
		"key":                          m.Key,
		"format":                       m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaInformasiKendaraanKurirSTNKFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaInformasiKendaraanKurirSTNKFoto) TableNameHistorical() string {
	return "media_informasi_kendaraan_kurir_stnk_foto_historical"
}

type MediaInformasiKurirKTPFoto struct {
	ID               int64
	IdInformasiKurir int64
	InformasiKurir   InformasiKurir
	Key              string
	Format           string
}

func (MediaInformasiKurirKTPFoto) PathName() string {
	return "/media_informasi_kurir_ktp_foto/"
}

func (m *MediaInformasiKurirKTPFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaInformasiKurirKTPFoto) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                 m.ID,
		"id_informasi_kurir": m.IdInformasiKurir,
		"key":                m.Key,
		"format":             m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaInformasiKurirKTPFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaInformasiKurirKTPFoto) TableNameHistorical() string {
	return "media_informasi_kurir_ktp_foto_historical"
}

type MediaReviewFoto struct {
	ID       int64
	IdReview int64
	Review   Review
	Key      string
	Format   string
}

func (MediaReviewFoto) PathName() string {
	return "/media_review_foto/"
}

func (m *MediaReviewFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaReviewFoto) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":        m.ID,
		"id_review": m.IdReview,
		"key":       m.Key,
		"format":    m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaReviewFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaReviewFoto) TableNameHistorical() string {
	return "media_review_foto_historical"
}

type MediaReviewVideo struct {
	ID       int64
	IdReview int64
	Review   Review
	Key      string
	Format   string
}

func (MediaReviewVideo) PathName() string {
	return "/media_review_video/"
}

func (m *MediaReviewVideo) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaReviewVideo) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":        m.ID,
		"id_review": m.IdReview,
		"key":       m.Key,
		"format":    m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaReviewVideo) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaReviewVideo) TableNameHistorical() string {
	return "media_review_video_historical"
}

type MediaTransaksiApprovedFoto struct {
	ID          int64
	IdTransaksi int64
	Transaksi   Transaksi
	Key         string
	Format      string
}

func (MediaTransaksiApprovedFoto) PathName() string {
	return "/media_transaksi_approved_foto/"
}

func (m *MediaTransaksiApprovedFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaTransaksiApprovedFoto) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":           m.ID,
		"id_transaksi": m.IdTransaksi,
		"key":          m.Key,
		"format":       m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaTransaksiApprovedFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaTransaksiApprovedFoto) TableNameHistorical() string {
	return "media_transaksi_approved_foto_historical"
}

type MediaTransaksiApprovedVideo struct {
	ID          int64
	IdTransaksi int64
	Transaksi   Transaksi
	Key         string
	Format      string
}

func (MediaTransaksiApprovedVideo) PathName() string {
	return "/media_transaksi_approved_video/"
}

func (m *MediaTransaksiApprovedVideo) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaTransaksiApprovedVideo) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":           m.ID,
		"id_transaksi": m.IdTransaksi,
		"key":          m.Key,
		"format":       m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaTransaksiApprovedVideo) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaTransaksiApprovedVideo) TableNameHistorical() string {
	return "media_transaksi_approved_video_historical"
}

type MediaPengirimanPickedUpFoto struct {
	ID           int64
	IdPengiriman int64
	Pengiriman   Pengiriman
	Key          string
	Format       string
}

func (MediaPengirimanPickedUpFoto) PathName() string {
	return "/media_pengiriman_picked_up_foto/"
}

func (m *MediaPengirimanPickedUpFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaPengirimanPickedUpFoto) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.ID,
		"id_pengiriman": m.IdPengiriman,
		"key":           m.Key,
		"format":        m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaPengirimanPickedUpFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaPengirimanPickedUpFoto) TableNameHistorical() string {
	return "media_pengiriman_picked_up_foto_historical"
}

type MediaPengirimanSampaiFoto struct {
	ID           int64
	IdPengiriman int64
	Pengiriman   Pengiriman
	Key          string
	Format       string
}

func (MediaPengirimanSampaiFoto) PathName() string {
	return "/media_pengiriman_sampai_foto/"

}

func (m *MediaPengirimanSampaiFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaPengirimanSampaiFoto) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.ID,
		"id_pengiriman": m.IdPengiriman,
		"key":           m.Key,
		"format":        m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaPengirimanSampaiFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaPengirimanSampaiFoto) TableNameHistorical() string {
	return "media_pengiriman_sampai_foto_historical"
}

type MediaPengirimanEkspedisiPickedUpFoto struct {
	ID                    int64
	IdPengirimanEkspedisi int64
	PengirimanEkspedisi   PengirimanEkspedisi
	Key                   string
	Format                string
}

func (MediaPengirimanEkspedisiPickedUpFoto) PathName() string {
	return "/media_pengiriman_ekspedisi_picked_up_foto/"
}

func (m *MediaPengirimanEkspedisiPickedUpFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaPengirimanEkspedisiPickedUpFoto) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                      m.ID,
		"id_pengiriman_ekspedisi": m.IdPengirimanEkspedisi,
		"key":                     m.Key,
		"format":                  m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaPengirimanEkspedisiPickedUpFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaPengirimanEkspedisiPickedUpFoto) TableNameHistorical() string {
	return "media_pengiriman_ekspedisi_picked_up_foto_historical"
}

type MediaPengirimanEkspedisiSampaiAgentFoto struct {
	ID                    int64
	IdPengirimanEkspedisi int64
	PengirimanEkspedisi   PengirimanEkspedisi
	Key                   string
	Format                string
}

func (MediaPengirimanEkspedisiSampaiAgentFoto) PathName() string {
	return "/media_pengiriman_ekspedisi_sampai_agent_foto/"
}

func (m *MediaPengirimanEkspedisiSampaiAgentFoto) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", m.TableNameHistorical())
	return nil
}

func (m *MediaPengirimanEkspedisiSampaiAgentFoto) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                      m.ID,
		"id_pengiriman_ekspedisi": m.IdPengirimanEkspedisi,
		"key":                     m.Key,
		"format":                  m.Format,
	}
}

// DropTable disesuaikan menggunakan m.TableNameHistorical() secara dinamis
func (m *MediaPengirimanEkspedisiSampaiAgentFoto) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, m.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", m.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", m.TableNameHistorical())
	return nil
}

func (MediaPengirimanEkspedisiSampaiAgentFoto) TableNameHistorical() string {
	return "media_pengiriman_ekspedisi_sampai_agent_foto_historical"
}

func (m *MediaPenggunaProfilFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaSellerProfilFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaSellerBannerFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaSellerTokoFisikFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaSellerTokoFisikFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		key text,
		format text,
		created_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaKurirProfilFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaEtalaseFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaBarangIndukFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBarangIndukFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_barang_induk bigint,
		key text,
		format text,
		created_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaBarangIndukVideo) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaKategoriBarangFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaDistributorDataDokumen) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaDistributorDataNPWPFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaDistributorDataNIBFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaDistributorDataSuratKerjasamaDokumen) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaBrandDataPerwakilanDokumen) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBrandDataPerwakilanDokumen dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_brand_data bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaBrandDataSertifikatFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBrandDataSertifikatFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_brand_data bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaBrandDataNIBFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBrandDataNIBFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_brand_data bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaBrandDataNPWPFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBrandDataNPWPFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_brand_data bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaBrandDataLogoFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBrandDataLogoFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_brand_data bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaBrandDataSuratKerjasamaDokumen) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaBrandDataSuratKerjasamaDokumen dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaInformasiKendaraanKurirKendaraanFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaInformasiKendaraanKurirKendaraanFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_informasi_kendaraan_kurir bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaInformasiKendaraanKurirBPKBFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaInformasiKendaraanKurirBPKBFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_informasi_kendaraan_kurir bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaInformasiKendaraanKurirSTNKFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaInformasiKendaraanKurirSTNKFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_informasi_kendaraan_kurir bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaInformasiKurirKTPFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaInformasiKurirKTPFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_informasi_kurir bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaReviewFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaReviewFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_review bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaReviewVideo) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaReviewVideo dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_review bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaTransaksiApprovedFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaTransaksiApprovedFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_transaksi bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaTransaksiApprovedVideo) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaTransaksiApprovedVideo dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_transaksi bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaPengirimanPickedUpFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaPengirimanPickedUpFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengiriman bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaPengirimanSampaiFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaPengirimanSampaiFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengiriman bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaPengirimanEkspedisiPickedUpFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaPengirimanEkspedisiPickedUpFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengiriman_ekspedisi bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}

func (m *MediaPengirimanEkspedisiSampaiAgentFoto) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct MediaPengirimanEkspedisiSampaiAgentFoto dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengiriman_ekspedisi bigint,
		key text,
		format text,
		PRIMARY KEY (id)
	)`, m.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica\n", m.TableNameSotReplica())
	return nil
}
