package cass_models

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"gorm.io/gorm"
)

type EntitySocialMedia struct {
	ID         int64
	EntityId   int64
	Whatsapp   string
	Facebook   string
	TikTok     string
	Instagram  string
	Metadata   []byte
	EntityType string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	Pencatatan
}

func (EntitySocialMedia) TableNameHistorical() string {
	return "entity_social_media_historical"
}

func (e *EntitySocialMedia) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct EntitySocialMedia dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		entity_id bigint,
		whatsapp text,
		facebook text,
		tik_tok text,
		instagram text,
		metadata blob,
		entity_type text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, e.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", e.TableNameHistorical())
	return nil
}

func (e *EntitySocialMedia) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":           e.ID,
		"entity_id":    e.EntityId,
		"whatsapp":     e.Whatsapp,
		"facebook":     e.Facebook,
		"tik_tok":      e.TikTok,
		"instagram":    e.Instagram,
		"metadata":     e.Metadata,
		"entity_type":  e.EntityType,
		"created_at":   e.CreatedAt,
		"updated_at":   e.UpdatedAt,
		"deleted_at":   e.DeletedAt,
		"tahun_update": e.TahunUpdate,
		"bulan_update": e.BulanUpdate,
		"event_time":   e.EventTime,
	}
}

// DropTable disesuaikan menggunakan e.TableName() secara dinamis
func (e *EntitySocialMedia) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, e.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", e.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", e.TableNameHistorical())
	return nil
}

type Komentar struct {
	ID            int64
	IdBarangInduk int32
	Baranginduk   BarangInduk
	IdEntity      int64
	JenisEntity   string
	Komentar      string
	IsSeller      bool
	Dibalas       int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Pencatatan
}

func (Komentar) TableNameHistorical() string {
	return "komentar_historical"
}

func (k *Komentar) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct Komentar dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_barang_induk int,
		id_entity bigint,
		jenis_entity text,
		komentar text,
		is_seller boolean,
		dibalas bigint,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, k.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", k.TableNameHistorical())
	return nil
}

func (k *Komentar) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":              k.ID,
		"id_barang_induk": k.IdBarangInduk,
		"id_entity":       k.IdEntity,
		"jenis_entity":    k.JenisEntity,
		"komentar":        k.Komentar,
		"is_seller":       k.IsSeller,
		"dibalas":         k.Dibalas,
		"created_at":      k.CreatedAt,
		"updated_at":      k.UpdatedAt,
		"deleted_at":      k.DeletedAt,
		"tahun_update":    k.TahunUpdate,
		"bulan_update":    k.BulanUpdate,
		"event_time":      k.EventTime,
	}
}

// DropTable disesuaikan menggunakan k.TableName() secara dinamis
func (k *Komentar) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, k.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", k.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", k.TableNameHistorical())
	return nil
}

type KomentarChild struct {
	ID          int64
	IdKomentar  int64
	Komentar    Komentar
	IdEntity    int64
	JenisEntity string
	IsiKomentar string
	IsSeller    bool
	Mention     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Pencatatan
}

func (KomentarChild) TableNameHistorical() string {
	return "komentar_child_historical"
}

func (k *KomentarChild) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct KomentarChild dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_komentar bigint,
		id_entity bigint,
		jenis_entity text,
		isi_komentar text,
		is_seller boolean,
		mention text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, k.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", k.TableNameHistorical())
	return nil
}

func (k *KomentarChild) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":           k.ID,
		"id_komentar":  k.IdKomentar,
		"id_entity":    k.IdEntity,
		"jenis_entity": k.JenisEntity,
		"isi_komentar": k.IsiKomentar,
		"is_seller":    k.IsSeller,
		"mention":      k.Mention,
		"created_at":   k.CreatedAt,
		"updated_at":   k.UpdatedAt,
		"deleted_at":   k.DeletedAt,
		"tahun_update": k.TahunUpdate,
		"bulan_update": k.BulanUpdate,
		"event_time":   k.EventTime,
	}
}

// DropTable disesuaikan menggunakan k.TableName() secara dinamis
func (k *KomentarChild) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, k.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", k.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", k.TableNameHistorical())
	return nil
}

// ///////////////////////////////////////////////////////////////////////////////////////////
// ENGAGEMENT PENGGUNA
// ///////////////////////////////////////////////////////////////////////////////////////////

type Keranjang struct {
	ID             int64
	IdPengguna     int64
	Pengguna       Pengguna
	IdSeller       int32
	Seller         Seller
	IdBarangInduk  int32
	BarangInduk    BarangInduk
	IdKategori     int64
	Kategoribarang KategoriBarang
	Jumlah         int16
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	Pencatatan
}

func (Keranjang) TableNameHistorical() string {
	return "keranjang_historical"
}

func (k *Keranjang) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct Keranjang dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_seller int,
		id_barang_induk int,
		id_kategori bigint,
		jumlah smallint,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, k.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", k.TableNameHistorical())
	return nil
}

func (k *Keranjang) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":              k.ID,
		"id_pengguna":     k.IdPengguna,
		"id_seller":       k.IdSeller,
		"id_barang_induk": k.IdBarangInduk,
		"id_kategori":     k.IdKategori,
		"jumlah":          k.Jumlah,
		"status":          k.Status,
		"created_at":      k.CreatedAt,
		"updated_at":      k.UpdatedAt,
		"deleted_at":      k.DeletedAt,
		"tahun_update":    k.TahunUpdate,
		"bulan_update":    k.BulanUpdate,
		"event_time":      k.EventTime,
	}
}

func (k *Keranjang) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, k.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", k.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", k.TableNameHistorical())
	return nil
}

type BarangDisukai struct {
	ID            int64
	IdPengguna    int64
	Pengguna      Pengguna
	IdBarangInduk int32
	BarangInduk   BarangInduk
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Pencatatan
}

func (BarangDisukai) TableNameHistorical() string {
	return "barang_disukai_historical"
}

func (b *BarangDisukai) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct BarangDisukai dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_barang_induk int,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, b.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", b.TableNameHistorical())
	return nil
}

func (b *BarangDisukai) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":              b.ID,
		"id_pengguna":     b.IdPengguna,
		"id_barang_induk": b.IdBarangInduk,
		"created_at":      b.CreatedAt,
		"updated_at":      b.UpdatedAt,
		"deleted_at":      b.DeletedAt,
		"tahun_update":    b.TahunUpdate,
		"bulan_update":    b.BulanUpdate,
		"event_time":      b.EventTime,
	}
}

// DropTable disesuaikan menggunakan b.TableName() secara dinamis
func (b *BarangDisukai) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, b.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", b.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", b.TableNameHistorical())
	return nil
}

type BarangWishlist struct {
	ID            int64
	IdPengguna    int64
	Pengguna      Pengguna
	IdBarangInduk int32
	BarangInduk   BarangInduk
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Pencatatan
}

func (b BarangWishlist) TableNameHistorical() string {
	return "barang_wishlist_historical"
}

func (b *BarangWishlist) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct BarangWishlist dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_barang_induk int,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, b.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", b.TableNameHistorical())
	return nil
}

func (b *BarangWishlist) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":              b.ID,
		"id_pengguna":     b.IdPengguna,
		"id_barang_induk": b.IdBarangInduk,
		"created_at":      b.CreatedAt,
		"updated_at":      b.UpdatedAt,
		"deleted_at":      b.DeletedAt,
		"tahun_update":    b.TahunUpdate,
		"bulan_update":    b.BulanUpdate,
		"event_time":      b.EventTime,
	}
}

// DropTable disesuaikan menggunakan b.TableName() secara dinamis
func (b *BarangWishlist) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, b.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", b.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", b.TableNameHistorical())
	return nil
}

type AlamatPengguna struct {
	ID              int64
	IDPengguna      int64
	Pengguna        Pengguna
	PanggilanAlamat string
	NomorTelephone  string
	NamaAlamat      string
	Provinsi        string
	Kota            string
	KodePos         string
	KodeNegara      string
	Deskripsi       string
	Longitude       float64
	Latitude        float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	Pencatatan
}

func (AlamatPengguna) TableNameHistorical() string {
	return "alamat_pengguna_historical"
}

func (a *AlamatPengguna) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct AlamatPengguna dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		panggilan_alamat text,
		nomor_telephone text,
		nama_alamat text,
		provinsi text,
		kota text,
		kode_pos text,
		kode_negara text,
		deskripsi text,
		longitude double,
		latitude double,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, a.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", a.TableNameHistorical())
	return nil
}

func (a *AlamatPengguna) ParseToInsertType() map[string]interface{} {
	var deletedAtInterface interface{} = nil
	if a.DeletedAt.Valid {
		deletedAtInterface = a.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":               a.ID,
		"id_pengguna":      a.IDPengguna,
		"panggilan_alamat": a.PanggilanAlamat,
		"nomor_telephone":  a.NomorTelephone,
		"nama_alamat":      a.NamaAlamat,
		"provinsi":         a.Provinsi,
		"kota":             a.Kota,
		"kode_pos":         a.KodePos,
		"kode_negara":      a.KodeNegara,
		"deskripsi":        a.Deskripsi,
		"longitude":        a.Longitude,
		"latitude":         a.Latitude,
		"created_at":       a.CreatedAt,
		"updated_at":       a.UpdatedAt,
		"deleted_at":       deletedAtInterface,
		"tahun_update":     a.TahunUpdate,
		"bulan_update":     a.BulanUpdate,
		"event_time":       a.EventTime,
	}
}

func (a *AlamatPengguna) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, a.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", a.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", a.TableNameHistorical())
	return nil
}

type Wishlist struct {
	ID            int64
	IdPengguna    int64
	Pengguna      Pengguna
	IdBarangInduk int32
	BarangInduk   BarangInduk
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Pencatatan
}

func (Wishlist) TableNameHistorical() string {
	return "wishlist_historical"
}

func (w *Wishlist) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct Wishlist dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_barang_induk int,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, w.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", w.TableNameHistorical())
	return nil
}

func (w *Wishlist) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":              w.ID,
		"id_pengguna":     w.IdPengguna,
		"id_barang_induk": w.IdBarangInduk,
		"created_at":      w.CreatedAt,
		"updated_at":      w.UpdatedAt,
		"deleted_at":      w.DeletedAt,
		"tahun_update":    w.TahunUpdate,
		"bulan_update":    w.BulanUpdate,
		"event_time":      w.EventTime,
	}
}

// DropTable disesuaikan menggunakan w.TableName() secara dinamis
func (w *Wishlist) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, w.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", w.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", w.TableNameHistorical())
	return nil
}

type Review struct {
	ID            int64
	IdPengguna    int64
	Pengguna      Pengguna
	IdBarangInduk int32
	BarangInduk   BarangInduk
	Rating        float32
	Ulasan        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Pencatatan
}

func (Review) TableNameHistorical() string {
	return "review_historical"
}

func (r *Review) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct Review dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_barang_induk int,
		rating float,
		ulasan text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, r.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", r.TableNameHistorical())
	return nil
}

func (r *Review) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":              r.ID,
		"id_pengguna":     r.IdPengguna,
		"id_barang_induk": r.IdBarangInduk,
		"rating":          r.Rating,
		"ulasan":          r.Ulasan,
		"created_at":      r.CreatedAt,
		"updated_at":      r.UpdatedAt,
		"deleted_at":      r.DeletedAt,
		"tahun_update":    r.TahunUpdate,
		"bulan_update":    r.BulanUpdate,
		"event_time":      r.EventTime,
	}
}

// DropTable disesuaikan menggunakan r.TableName() secara dinamis
func (r *Review) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, r.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", r.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", r.TableNameHistorical())
	return nil
}

type ReviewLike struct {
	ID         int64
	IdPengguna int64
	Pengguna   Pengguna
	IdReview   int64
	Review     Review
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	Pencatatan
}

func (ReviewLike) TableNameHistorical() string {
	return "review_like_historical"
}

func (r *ReviewLike) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct ReviewLike dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_review bigint,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, r.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", r.TableNameHistorical())
	return nil
}

func (r *ReviewLike) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":           r.ID,
		"id_pengguna":  r.IdPengguna,
		"id_review":    r.IdReview,
		"created_at":   r.CreatedAt,
		"updated_at":   r.UpdatedAt,
		"deleted_at":   r.DeletedAt,
		"tahun_update": r.TahunUpdate,
		"bulan_update": r.BulanUpdate,
		"event_time":   r.EventTime,
	}
}

// DropTable disesuaikan menggunakan r.TableName() secara dinamis
func (r *ReviewLike) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, r.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", r.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", r.TableNameHistorical())
	return nil
}

type ReviewDislike struct {
	ID         int64
	IdPengguna int64
	Pengguna   Pengguna
	IdReview   int64
	Review     Review
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	Pencatatan
}

func (ReviewDislike) TableNameHistorical() string {
	return "review_dislike_historical"
}

func (r *ReviewDislike) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct ReviewDislike dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_review bigint,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, r.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", r.TableNameHistorical())
	return nil
}

func (r *ReviewDislike) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":           r.ID,
		"id_pengguna":  r.IdPengguna,
		"id_review":    r.IdReview,
		"created_at":   r.CreatedAt,
		"updated_at":   r.UpdatedAt,
		"deleted_at":   r.DeletedAt,
		"tahun_update": r.TahunUpdate,
		"bulan_update": r.BulanUpdate,
		"event_time":   r.EventTime,
	}
}

// DropTable disesuaikan menggunakan r.TableName() secara dinamis
func (r *ReviewDislike) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, r.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", r.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", r.TableNameHistorical())
	return nil
}

// ///////////////////////////////////////////////////////////////////////////////////////////
// ENGAGEMENT SELLER
// ///////////////////////////////////////////////////////////////////////////////////////////

type Jenis_Seller struct {
	ID               int64
	IdSeller         int32
	Seller           Seller
	ValidationStatus string
	Alasan           string
	AlasanAdmin      string
	TargetJenis      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
	Pencatatan
}

func (Jenis_Seller) TableName() string {
	return "jenis_seller_validation_historical"
}

func (j *Jenis_Seller) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct Jenis_Seller dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		validation_status text,
		alasan text,
		alasan_admin text,
		target_jenis text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, j.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", j.TableName())
	return nil
}

func (j *Jenis_Seller) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":                j.ID,
		"id_seller":         j.IdSeller,
		"validation_status": j.ValidationStatus,
		"alasan":            j.Alasan,
		"alasan_admin":      j.AlasanAdmin,
		"target_jenis":      j.TargetJenis,
		"created_at":        j.CreatedAt,
		"updated_at":        j.UpdatedAt,
		"deleted_at":        j.DeletedAt,
		"tahun_update":      j.TahunUpdate,
		"bulan_update":      j.BulanUpdate,
		"event_time":        j.EventTime,
	}
}

// DropTable disesuaikan menggunakan j.TableName() secara dinamis
func (j *Jenis_Seller) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, j.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", j.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", j.TableName())
	return nil
}

type BatalTransaksi struct {
	ID             int64
	IdTransaksi    int64
	ITransaksi     Transaksi
	DibatalkanOleh string
	Alasan         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	Pencatatan
}

func (BatalTransaksi) TableName() string {
	return "batal_transaksi_historical"
}

func (b *BatalTransaksi) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct BatalTransaksi dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_transaksi bigint,
		dibatalkan_oleh text,
		alasan text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", b.TableName())
	return nil
}

func (b *BatalTransaksi) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":              b.ID,
		"id_transaksi":    b.IdTransaksi,
		"dibatalkan_oleh": b.DibatalkanOleh,
		"alasan":          b.Alasan,
		"created_at":      b.CreatedAt,
		"updated_at":      b.UpdatedAt,
		"deleted_at":      b.DeletedAt,
		"tahun_update":    b.TahunUpdate,
		"bulan_update":    b.BulanUpdate,
		"event_time":      b.EventTime,
	}
}

// DropTable disesuaikan menggunakan b.TableName() secara dinamis
func (b *BatalTransaksi) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", b.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", b.TableName())
	return nil
}

type Follower struct {
	ID         int64
	IdFollower int64
	Pengguna   Pengguna
	IdFollowed int64
	Seller     Seller
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	Pencatatan
}

func (Follower) TableName() string {
	return "follower_historical"
}

func (f *Follower) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct Follower dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_follower bigint,
		id_followed bigint,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, f.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", f.TableName())
	return nil
}

func (f *Follower) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":           f.ID,
		"id_follower":  f.IdFollower,
		"id_followed":  f.IdFollowed,
		"created_at":   f.CreatedAt,
		"updated_at":   f.UpdatedAt,
		"deleted_at":   f.DeletedAt,
		"tahun_update": f.TahunUpdate,
		"bulan_update": f.BulanUpdate,
		"event_time":   f.EventTime,
	}
}

// DropTable disesuaikan menggunakan f.TableName() secara dinamis
func (f *Follower) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, f.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", f.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", f.TableName())
	return nil
}

type AlamatGudang struct {
	ID              int64
	IDSeller        int32
	Seller          Seller
	PanggilanAlamat string
	NomorTelephone  string
	NamaAlamat      string
	Provinsi        string
	Kota            string
	KodePos         string
	KodeNegara      string
	Deskripsi       string
	Longitude       float64
	Latitude        float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	Pencatatan
}

func (AlamatGudang) TableName() string {
	return "alamat_gudang_historical"
}

func (a *AlamatGudang) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct AlamatGudang dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		panggilan_alamat text,
		nomor_telephone text,
		nama_alamat text,
		provinsi text,
		kota text,
		kode_pos text,
		kode_negara text,
		deskripsi text,
		longitude double,
		latitude double,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, a.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", a.TableName())
	return nil
}

func (a *AlamatGudang) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if a.DeletedAt.Valid {
		deletedAtInterface = a.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":               a.ID,
		"id_seller":        a.IDSeller,
		"panggilan_alamat": a.PanggilanAlamat,
		"nomor_telephone":  a.NomorTelephone,
		"nama_alamat":      a.NamaAlamat,
		"provinsi":         a.Provinsi,
		"kota":             a.Kota,
		"kode_pos":         a.KodePos,
		"kode_negara":      a.KodeNegara,
		"deskripsi":        a.Deskripsi,
		"longitude":        a.Longitude,
		"latitude":         a.Latitude,
		"created_at":       a.CreatedAt,
		"updated_at":       a.UpdatedAt,
		"deleted_at":       deletedAtInterface,
		"tahun_update":     a.TahunUpdate,
		"bulan_update":     a.BulanUpdate,
		"event_time":       a.EventTime,
	}
}

// DropTable disesuaikan menggunakan a.TableName() secara dinamis
func (a *AlamatGudang) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, a.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", a.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", a.TableName())
	return nil
}

type DistributorData struct {
	ID                        int64
	SellerId                  int32
	Seller                    Seller
	NamaPerusahaan            string
	NIB                       string
	NPWP                      string
	DokumenIzinDistributorUrl string
	Alasan                    string
	Status                    string
	Pencatatan
}

func (DistributorData) TableName() string {
	return "distributor_data_historical_data"
}

func (d *DistributorData) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct DistributorData dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		seller_id int,
		nama_perusahaan text,
		nib text,
		npwp text,
		dokumen_izin_distributor_url text,
		alasan text,
		status text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, d.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", d.TableName())
	return nil
}

func (d *DistributorData) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":                           d.ID,
		"seller_id":                    d.SellerId,
		"nama_perusahaan":              d.NamaPerusahaan,
		"nib":                          d.NIB,  // Tetap di-map ke nib (snake_case)
		"npwp":                         d.NPWP, // Tetap di-map ke npwp (snake_case)
		"dokumen_izin_distributor_url": d.DokumenIzinDistributorUrl,
		"alasan":                       d.Alasan,
		"status":                       d.Status,
		"tahun_update":                 d.TahunUpdate,
		"bulan_update":                 d.BulanUpdate,
		"event_time":                   d.EventTime,
	}
}

// DropTable disesuaikan menggunakan d.TableName() secara dinamis
func (d *DistributorData) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, d.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", d.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", d.TableName())
	return nil
}

type BrandData struct {
	ID                    int64
	SellerId              int32
	Seller                Seller
	NamaPerusahaan        string
	NegaraAsal            string
	LembagaPendaftaran    string
	NomorPendaftaranMerek string
	SertifikatMerekUrl    string
	DokumenPerwakilanUrl  string
	NIB                   string
	NPWP                  string
	Alasan                string
	Status                string
	Pencatatan
}

func (BrandData) TableName() string {
	return "brand_data_historical"
}

type Etalase struct {
	ID           int64
	SellerID     int64  `gorm:"column:id_seller;not null" json:"id_seller_etalase"`
	Seller       Seller `gorm:"foreignKey:SellerID;references:ID" json:"-"`
	Nama         string `gorm:"column:nama;type:varchar(100);not null" json:"nama_etalase"`
	Deskripsi    string `gorm:"column:deskripsi;type:text" json:"deskripsi_etalase"`
	JumlahBarang int32  `gorm:"column:jumlah_barang;not null;default:0" json:"jumlah_barang"`
}

func (Etalase) TableName() string {
	return "etalase_historical"
}

func (b *BrandData) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct BrandData dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		seller_id int,
		nama_perusahaan text,
		negara asal text,
		lembaga_pendaftaran text,
		nomor_pendaftaran_merek text,
		sertifikat_merek_url text,
		dokumen_perwakilan_url text,
		nib text,
		npwp text,
		alasan text,
		status text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", b.TableName())
	return nil
}

func (b *BrandData) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":                      b.ID,
		"seller_id":               b.SellerId,
		"nama_perusahaan":         b.NamaPerusahaan,
		"negara_asal":             b.NegaraAsal,
		"lembaga_pendaftaran":     b.LembagaPendaftaran,
		"nomor_pendaftaran_merek": b.NomorPendaftaranMerek,
		"sertifikat_merek_url":    b.SertifikatMerekUrl,
		"dokumen_perwakilan_url":  b.DokumenPerwakilanUrl,
		"nib":                     b.NIB,
		"npwp":                    b.NPWP,
		"alasan":                  b.Alasan,
		"status":                  b.Status,
		"tahun_update":            b.TahunUpdate,
		"bulan_update":            b.BulanUpdate,
		"event_time":              b.EventTime,
	}
}

// DropTable disesuaikan menggunakan b.TableName() secara dinamis
func (b *BrandData) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", b.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", b.TableName())
	return nil
}

type BarangKeEtalase struct {
	ID            int64
	IdEtalase     int64
	Etalase       Etalase
	IdBarangInduk int64
	BarangInduk   BarangInduk
	Pencatatan
}

func (BarangKeEtalase) TableName() string {
	return "barang_ke_etalase_historical"
}

func (b *BarangKeEtalase) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct BarangKeEtalase dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_etalase bigint,
		id_barang_induk bigint,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", b.TableName())
	return nil
}

func (b *BarangKeEtalase) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":              b.ID,
		"id_etalase":      b.IdEtalase,
		"id_barang_induk": b.IdBarangInduk,
		"tahun_update":    b.TahunUpdate,
		"bulan_update":    b.BulanUpdate,
		"event_time":      b.EventTime,
	}
}

// DropTable disesuaikan menggunakan b.TableName() secara dinamis
func (b *BarangKeEtalase) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", b.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", b.TableName())
	return nil
}

type DiskonProduk struct {
	ID            int64
	SellerId      int32
	Seller        Seller
	Nama          string
	Deskripsi     string
	DiskonPersen  float64
	BerlakuMulai  time.Time
	BerlakuSampai time.Time
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	Pencatatan
}

func (DiskonProduk) TableName() string {
	return "diskon_produk_historical"
}

func (d *DiskonProduk) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct DiskonProduk dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		seller_id int,
		nama text,
		deskripsi text,
		diskon_persen double,
		berlaku_mulai timestamp,
		berlaku_sampai timestamp,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, d.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", d.TableName())
	return nil
}

func (d *DiskonProduk) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if d.DeletedAt.Valid {
		deletedAtInterface = d.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":             d.ID,
		"seller_id":      d.SellerId,
		"nama":           d.Nama,
		"deskripsi":      d.Deskripsi,
		"diskon_persen":  d.DiskonPersen,
		"berlaku_mulai":  d.BerlakuMulai,
		"berlaku_sampai": d.BerlakuSampai,
		"status":         d.Status,
		"created_at":     d.CreatedAt,
		"updated_at":     d.UpdatedAt,
		"deleted_at":     deletedAtInterface,
		"tahun_update":   d.TahunUpdate,
		"bulan_update":   d.BulanUpdate,
		"event_time":     d.EventTime,
	}
}

// DropTable disesuaikan menggunakan d.TableName() secara dinamis
func (d *DiskonProduk) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, d.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", d.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", d.TableName())
	return nil
}

type BarangDiDiskon struct {
	ID               int64
	SellerId         int32
	Seller           Seller
	IdDiskon         int64
	DiskonProduk     DiskonProduk
	IdBarangInduk    int32
	BarangInduk      BarangInduk
	IdKategoriBarang int64
	KategoriBarang   KategoriBarang
	Status           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	Pencatatan
}

func (BarangDiDiskon) TableName() string {
	return "barang_di_diskon_historical"
}

func (b *BarangDiDiskon) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct BarangDiDiskon dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		seller_id int,
		id_diskon bigint,
		id_barang_induk int,
		id_kategori_barang bigint,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", b.TableName())
	return nil
}

func (b *BarangDiDiskon) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if b.DeletedAt.Valid {
		deletedAtInterface = b.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":                 b.ID,
		"seller_id":          b.SellerId,
		"id_diskon":          b.IdDiskon,
		"id_barang_induk":    b.IdBarangInduk,
		"id_kategori_barang": b.IdKategoriBarang,
		"status":             b.Status,
		"created_at":         b.CreatedAt,
		"updated_at":         b.UpdatedAt,
		"deleted_at":         deletedAtInterface,
		"tahun_update":       b.TahunUpdate,
		"bulan_update":       b.BulanUpdate,
		"event_time":         b.EventTime,
	}
}

// DropTable disesuaikan menggunakan b.TableName() secara dinamis
func (b *BarangDiDiskon) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", b.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", b.TableName())
	return nil
}

// ///////////////////////////////////////////////////////////////////////////////////////////
// ENGAGEMENT KURIR
// ///////////////////////////////////////////////////////////////////////////////////////////

type InformasiKurir struct {
	ID           int64
	IDkurir      int64
	Kurir        Kurir
	TanggalLahir string
	Alasan       string
	Ktp          bool
	InformasiSim bool
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Pencatatan
}

func (InformasiKurir) TableName() string {
	return "informasi_kurir_historical"
}

func (i *InformasiKurir) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct InformasiKurir dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kurir bigint,
		tanggal_lahir text,
		alasan text,
		ktp boolean,
		informasi_sim boolean,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, i.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", i.TableName())
	return nil
}

func (i *InformasiKurir) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":            i.ID,
		"id_kurir":      i.IDkurir,
		"tanggal_lahir": i.TanggalLahir,
		"alasan":        i.Alasan,
		"ktp":           i.Ktp,
		"informasi_sim": i.InformasiSim,
		"status":        i.Status,
		"created_at":    i.CreatedAt,
		"updated_at":    i.UpdatedAt,
		"deleted_at":    i.DeletedAt,
		"tahun_update":  i.TahunUpdate,
		"bulan_update":  i.BulanUpdate,
		"event_time":    i.EventTime,
	}
}

// DropTable disesuaikan menggunakan i.TableName() secara dinamis
func (i *InformasiKurir) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, i.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", i.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", i.TableName())
	return nil
}

type InformasiKendaraanKurir struct {
	ID             int64
	IDkurir        int64
	Kurir          Kurir
	JenisKendaraan string
	NamaKendaraan  string
	RodaKendaraan  string
	STNK           bool
	BPKB           bool
	NoRangka       string
	NoMesin        string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	Pencatatan
}

func (InformasiKendaraanKurir) TableName() string {
	return "informasi_kendaraan_kurir_historical"
}

func (i *InformasiKendaraanKurir) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct InformasiKendaraanKurir dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kurir bigint,
		jenis_kendaraan text,
		nama_kendaraan text,
		roda_kendaraan text,
		stnk boolean,
		bpkb boolean,
		no_rangka text,
		no_mesin text,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, i.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", i.TableName())
	return nil
}

func (i *InformasiKendaraanKurir) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":              i.ID,
		"id_kurir":        i.IDkurir,
		"jenis_kendaraan": i.JenisKendaraan,
		"nama_kendaraan":  i.NamaKendaraan,
		"roda_kendaraan":  i.RodaKendaraan,
		"stnk":            i.STNK,
		"bpkb":            i.BPKB,
		"no_rangka":       i.NoRangka,
		"no_mesin":        i.NoMesin,
		"status":          i.Status,
		"created_at":      i.CreatedAt,
		"updated_at":      i.UpdatedAt,
		"deleted_at":      i.DeletedAt,
		"tahun_update":    i.TahunUpdate,
		"bulan_update":    i.BulanUpdate,
		"event_time":      i.EventTime,
	}
}

func (i *InformasiKendaraanKurir) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, i.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", i.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", i.TableName())
	return nil
}

type AlamatKurir struct {
	ID              int64
	IdKurir         int64
	Kurir           Kurir
	PanggilanAlamat string
	NomorTelephone  string
	NamaAlamat      string
	Provinsi        string
	Kota            string
	KodeNegara      string
	KodePos         string
	Deskripsi       string
	Longitude       float64
	Latitude        float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
	Pencatatan
}

func (AlamatKurir) TableName() string {
	return "alamat_kurir_historical"
}

func (a *AlamatKurir) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct AlamatKurir dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kurir bigint,
		panggilan_alamat text,
		nomor_telephone text,
		nama_alamat text,
		provinsi text,
		kota text,
		kode_negara text,
		kode_pos text,
		deskripsi text,
		longitude double,
		latitude double,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, a.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", a.TableName())
	return nil
}

func (a *AlamatKurir) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":               a.ID,
		"id_kurir":         a.IdKurir,
		"panggilan_alamat": a.PanggilanAlamat,
		"nomor_telephone":  a.NomorTelephone,
		"nama_alamat":      a.NamaAlamat,
		"provinsi":         a.Provinsi,
		"kota":             a.Kota,
		"kode_negara":      a.KodeNegara,
		"kode_pos":         a.KodePos,
		"deskripsi":        a.Deskripsi,
		"longitude":        a.Longitude,
		"latitude":         a.Latitude,
		"created_at":       a.CreatedAt,
		"updated_at":       a.UpdatedAt,
		"deleted_at":       a.DeletedAt,
		"tahun_update":     a.TahunUpdate,
		"bulan_update":     a.BulanUpdate,
		"event_time":       a.EventTime,
	}
}

// DropTable disesuaikan menggunakan a.TableName() secara dinamis
func (a *AlamatKurir) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, a.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", a.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", a.TableName())
	return nil
}

type BidKurirData struct {
	ID              int64
	IdKurir         int64
	Kurir           Kurir
	JenisPengiriman string
	Mode            string
	Provinsi        string
	Kota            string
	IsEkspedisi     bool
	Alamat          string
	Longitude       float64
	Latitude        float64
	MaxKg           int16
	SlotTersisa     int32
	Dimulai         time.Time
	Selesai         *time.Time
	JenisKendaraan  string
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	Pencatatan
}

func (BidKurirData) TableName() string {
	return "bid_kurir_data_historical"
}

func (b *BidKurirData) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct BidKurirData dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kurir bigint,
		jenis_pengiriman text,
		mode text,
		provinsi text,
		kota text,
		is_ekspedisi boolean,
		alamat text,
		longitude double,
		latitude double,
		max_kg smallint,
		slot_tersisa int,
		dimulai timestamp,
		selesai timestamp,
		jenis_kendaraan text,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", b.TableName())
	return nil
}

func (b *BidKurirData) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if b.DeletedAt.Valid {
		deletedAtInterface = b.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":               b.ID,
		"id_kurir":         b.IdKurir,
		"jenis_pengiriman": b.JenisPengiriman,
		"mode":             b.Mode,
		"provinsi":         b.Provinsi,
		"kota":             b.Kota,
		"is_ekspedisi":     b.IsEkspedisi,
		"alamat":           b.Alamat,
		"longitude":        b.Longitude,
		"latitude":         b.Latitude,
		"max_kg":           b.MaxKg,
		"slot_tersisa":     b.SlotTersisa,
		"dimulai":          b.Dimulai,
		"selesai":          b.Selesai,
		"jenis_kendaraan":  b.JenisKendaraan,
		"status":           b.Status,
		"created_at":       b.CreatedAt,
		"updated_at":       b.UpdatedAt,
		"deleted_at":       deletedAtInterface,
		"tahun_update":     b.TahunUpdate,
		"bulan_update":     b.BulanUpdate,
		"event_time":       b.EventTime,
	}
}

// DropTable disesuaikan menggunakan b.TableName() secara dinamis
func (b *BidKurirData) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", b.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", b.TableName())
	return nil
}

type BidKurirNonEksScheduler struct {
	ID           int64
	IdBid        int64
	BidKurirData BidKurirData
	IdKurir      int64
	Kurir        Kurir
	Urutan       int8
	IdPengiriman int64
	Pengiriman   Pengiriman
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Pencatatan
}

func (BidKurirNonEksScheduler) TableName() string {
	return "bid_kurir_non_eks_scheduler_historical"
}

func (b *BidKurirNonEksScheduler) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct BidKurirNonEksScheduler dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_bid bigint,
		id_kurir bigint,
		urutan tinyint,
		id_pengiriman bigint,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", b.TableName())
	return nil
}

func (b *BidKurirNonEksScheduler) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if b.DeletedAt.Valid {
		deletedAtInterface = b.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":            b.ID,
		"id_bid":        b.IdBid,
		"id_kurir":      b.IdKurir,
		"urutan":        b.Urutan,
		"id_pengiriman": b.IdPengiriman,
		"status":        b.Status,
		"created_at":    b.CreatedAt,
		"updated_at":    b.UpdatedAt,
		"deleted_at":    deletedAtInterface,
		"tahun_update":  b.TahunUpdate,
		"bulan_update":  b.BulanUpdate,
		"event_time":    b.EventTime,
	}
}

// DropTable disesuaikan menggunakan b.TableName() secara dinamis
func (b *BidKurirNonEksScheduler) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", b.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", b.TableName())
	return nil
}

type BidKurirEksScheduler struct {
	ID                  int64
	IdBid               int64
	BidKurirData        BidKurirData
	IdKurir             int64
	Kurir               Kurir
	Urutan              int8
	IdPengirimanEks     int64
	PengirimanEkspedisi PengirimanEkspedisi
	Status              string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
	Pencatatan
}

func (BidKurirEksScheduler) TableName() string {
	return "bid_kurir_eks_scheduler_historical"
}

func (b *BidKurirEksScheduler) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct BidKurirEksScheduler dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_bid bigint,
		id_kurir bigint,
		urutan tinyint,
		id_pengiriman_eks bigint,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", b.TableName())
	return nil
}

func (b *BidKurirEksScheduler) ParseToInsertType() map[string]interface{} {
	// Memeriksa apakah DeletedAt di GORM valid sebelum di-insert
	var deletedAtInterface interface{} = nil
	if b.DeletedAt.Valid {
		deletedAtInterface = b.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":                b.ID,
		"id_bid":            b.IdBid,
		"id_kurir":          b.IdKurir,
		"urutan":            b.Urutan,
		"id_pengiriman_eks": b.IdPengirimanEks,
		"status":            b.Status,
		"created_at":        b.CreatedAt,
		"updated_at":        b.UpdatedAt,
		"deleted_at":        deletedAtInterface,
		"tahun_update":      b.TahunUpdate,
		"bulan_update":      b.BulanUpdate,
		"event_time":        b.EventTime,
	}
}

// DropTable disesuaikan menggunakan b.TableName() secara dinamis
func (b *BidKurirEksScheduler) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", b.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", b.TableName())
	return nil
}
