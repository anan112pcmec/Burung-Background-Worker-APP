package cass_models

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
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
}

func (EntitySocialMedia) TableNameHistorical() string {
	return "entity_social_media_historical"
}

func (e EntitySocialMedia) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		entity_id bigint,
		whatsapp text,
		facebook text,
		tiktok text,
		instagram text,
		metadata blob,
		entity_type text,
		created_at timestamp,
		updated_at timestamp,
		PRIMARY KEY (id)
	)`, e.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", e.TableNameSotReplica())
	return nil
}

func (e EntitySocialMedia) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		entity_id bigint,
		whatsapp text,
		facebook text,
		tiktok text,
		instagram text,
		metadata blob,
		entity_type text,
		created_at timestamp,
		updated_at timestamp,
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

func (e EntitySocialMedia) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":          e.ID,
		"entity_id":   e.EntityId,
		"whatsapp":    e.Whatsapp,
		"facebook":    e.Facebook,
		"tiktok":      e.TikTok,
		"instagram":   e.Instagram,
		"metadata":    e.Metadata,
		"entity_type": e.EntityType,
		"created_at":  e.CreatedAt,
		"updated_at":  e.UpdatedAt,
	}
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
	DeletedAt     time.Time
}

func (k Komentar) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, k.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", k.TableNameSotReplica())
	return nil
}

func (Komentar) TableNameHistorical() string {
	return "komentar_historical"
}

func (k Komentar) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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

func (k Komentar) ParseToCUDType() map[string]interface{} {
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
	}
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
	DeletedAt   time.Time
}

func (k KomentarChild) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_komentar bigint,
		id_entity bigint,
		jenis_entity text,
		komentar text,
		is_seller boolean,
		mention text,
		created_at timestamp,
		updated_at timestamp,
		PRIMARY KEY (id)
	)`, k.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", k.TableNameSotReplica())
	return nil
}

func (KomentarChild) TableNameHistorical() string {
	return "komentar_child_historical"
}

func (k KomentarChild) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_komentar bigint,
		id_entity bigint,
		jenis_entity text,
		komentar text,
		is_seller boolean,
		mention text,
		created_at timestamp,
		updated_at timestamp,
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

func (k KomentarChild) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":           k.ID,
		"id_komentar":  k.IdKomentar,
		"id_entity":    k.IdEntity,
		"jenis_entity": k.JenisEntity,
		"komentar":     k.IsiKomentar,
		"is_seller":    k.IsSeller,
		"mention":      k.Mention,
		"created_at":   k.CreatedAt,
		"updated_at":   k.UpdatedAt,
	}
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
}

func (k Keranjang) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_seller int,
		id_barang_induk int,
		id_kategori_barang bigint,
		jumlah smallint,
		status text,
		created_at timestamp,
		updated_at timestamp,
		PRIMARY KEY (id)
	)`, k.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", k.TableNameSotReplica())
	return nil
}

func (Keranjang) TableNameHistorical() string {
	return "keranjang_historical"
}

func (k Keranjang) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_seller int,
		id_barang_induk int,
		id_kategori_barang bigint,
		jumlah smallint,
		status text,
		created_at timestamp,
		updated_at timestamp,
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

func (k Keranjang) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                 k.ID,
		"id_pengguna":        k.IdPengguna,
		"id_seller":          k.IdSeller,
		"id_barang_induk":    k.IdBarangInduk,
		"id_kategori_barang": k.IdKategori,
		"jumlah":             k.Jumlah,
		"status":             k.Status,
		"created_at":         k.CreatedAt,
		"updated_at":         k.UpdatedAt,
	}
}

type BarangDisukai struct {
	ID            int64
	IdPengguna    int64
	Pengguna      Pengguna
	IdBarangInduk int32
	BarangInduk   BarangInduk
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (b BarangDisukai) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_barang_induk int,
		created_at timestamp,
		updated_at timestamp,
		PRIMARY KEY (id)
	)`, b.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", b.TableNameSotReplica())
	return nil
}

func (BarangDisukai) TableNameHistorical() string {
	return "barang_disukai_historical"
}

func (b BarangDisukai) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_barang_induk int,
		created_at timestamp,
		updated_at timestamp,
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

func (b BarangDisukai) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":              b.ID,
		"id_pengguna":     b.IdPengguna,
		"id_barang_induk": b.IdBarangInduk,
		"created_at":      b.CreatedAt,
		"updated_at":      b.UpdatedAt,
	}
}

type BarangWishlist struct {
	ID            int64
	IdPengguna    int64
	Pengguna      Pengguna
	IdBarangInduk int32
	BarangInduk   BarangInduk
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (b BarangWishlist) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_barang_induk int,
		created_at timestamp,
		updated_at timestamp,
		PRIMARY KEY (id)
	)`, b.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", b.TableNameSotReplica())
	return nil
}

func (b BarangWishlist) TableNameHistorical() string {
	return "barang_wishlist_historical"
}

func (b BarangWishlist) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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

func (b BarangWishlist) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":              b.ID,
		"id_pengguna":     b.IdPengguna,
		"id_barang_induk": b.IdBarangInduk,
		"created_at":      b.CreatedAt,
		"updated_at":      b.UpdatedAt,
	}
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
	DeletedAt       time.Time // REVISI: dari time.Time → gorm.DeletedAt (soft delete, sesuai sot_models)
}

func (a AlamatPengguna) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		panggilan_alamat text,
		nomor_telefon text,
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
		PRIMARY KEY (id)
	)`, a.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", a.TableNameSotReplica())
	return nil
}

func (AlamatPengguna) TableNameHistorical() string {
	return "alamat_pengguna_historical"
}

func (a AlamatPengguna) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		panggilan_alamat text,
		nomor_telefon text,
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

func (a AlamatPengguna) ParseToCUDType() map[string]interface{} {
	// REVISI: handle gorm.DeletedAt dengan benar

	return map[string]interface{}{
		"id":               a.ID,
		"id_pengguna":      a.IDPengguna,
		"panggilan_alamat": a.PanggilanAlamat,
		"nomor_telefon":    a.NomorTelephone,
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
		"deleted_at":       a.DeletedAt,
	}
}

type Wishlist struct {
	ID            int64
	IdPengguna    int64
	Pengguna      Pengguna
	IdBarangInduk int32
	BarangInduk   BarangInduk
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (w Wishlist) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_barang_induk int,
		created_at timestamp,
		updated_at timestamp,
		PRIMARY KEY (id)
	)`, w.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", w.TableNameSotReplica())
	return nil
}

func (Wishlist) TableNameHistorical() string {
	return "wishlist_historical"
}

func (w Wishlist) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_barang_induk int,
		created_at timestamp,
		updated_at timestamp,
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

func (w Wishlist) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":              w.ID,
		"id_pengguna":     w.IdPengguna,
		"id_barang_induk": w.IdBarangInduk,
		"created_at":      w.CreatedAt,
		"updated_at":      w.UpdatedAt,
	}
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
	DeletedAt     time.Time
}

func (r Review) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, r.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", r.TableNameSotReplica())
	return nil
}

func (Review) TableNameHistorical() string {
	return "review_historical"
}

func (r Review) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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

func (r Review) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":              r.ID,
		"id_pengguna":     r.IdPengguna,
		"id_barang_induk": r.IdBarangInduk,
		"rating":          r.Rating,
		"ulasan":          r.Ulasan,
		"created_at":      r.CreatedAt,
		"updated_at":      r.UpdatedAt,
		"deleted_at":      r.DeletedAt,
	}
}

type ReviewLike struct {
	ID         int64
	IdPengguna int64
	Pengguna   Pengguna
	IdReview   int64
	Review     Review
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (r ReviewLike) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_review bigint,
		created_at timestamp,
		updated_at timestamp,
		PRIMARY KEY (id)
	)`, r.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", r.TableNameSotReplica())
	return nil
}

func (ReviewLike) TableNameHistorical() string {
	return "review_like_historical"
}

func (r ReviewLike) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_review bigint,
		created_at timestamp,
		updated_at timestamp,
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

func (r ReviewLike) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":          r.ID,
		"id_pengguna": r.IdPengguna,
		"id_review":   r.IdReview,
		"created_at":  r.CreatedAt,
		"updated_at":  r.UpdatedAt,
	}
}

type ReviewDislike struct {
	ID         int64
	IdPengguna int64
	Pengguna   Pengguna
	IdReview   int64
	Review     Review
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (r ReviewDislike) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_review bigint,
		created_at timestamp,
		updated_at timestamp,
		PRIMARY KEY (id)
	)`, r.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", r.TableNameSotReplica())
	return nil
}

func (ReviewDislike) TableNameHistorical() string {
	return "review_dislike_historical"
}

func (r ReviewDislike) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_review bigint,
		created_at timestamp,
		updated_at timestamp,
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

func (r ReviewDislike) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":          r.ID,
		"id_pengguna": r.IdPengguna,
		"id_review":   r.IdReview,
		"created_at":  r.CreatedAt,
		"updated_at":  r.UpdatedAt,
	}
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
	DeletedAt        time.Time
}

func (j Jenis_Seller) TableNameSotReplica() string {
	return "jenis_seller_sot_replica"
}

func (j Jenis_Seller) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		validation_status text,
		alasan_seller text,
		alasan_admin text,
		target_jenis text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, j.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", j.TableNameSotReplica())
	return nil
}

func (Jenis_Seller) TableNameHistorical() string {
	return "jenis_seller_validation_historical"
}

func (j Jenis_Seller) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		validation_status text,
		alasan_seller text,
		alasan_admin text,
		target_jenis text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, j.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", j.TableNameHistorical())
	return nil
}

func (j Jenis_Seller) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                j.ID,
		"id_seller":         j.IdSeller,
		"validation_status": j.ValidationStatus,
		"alasan_seller":     j.Alasan,
		"alasan_admin":      j.AlasanAdmin,
		"target_jenis":      j.TargetJenis,
		"created_at":        j.CreatedAt,
		"updated_at":        j.UpdatedAt,
		"deleted_at":        j.DeletedAt,
	}
}

type BatalTransaksi struct {
	ID             int64
	IdTransaksi    int64
	ITransaksi     Transaksi
	DibatalkanOleh string
	Alasan         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

func (j BatalTransaksi) TableNameSotReplica() string {
	return "batal_transaksi_sot_replica"
}

func (b BatalTransaksi) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_transaksi bigint,
		dibatalkan_oleh text,
		alasan text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, b.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", b.TableNameSotReplica())
	return nil
}

func (BatalTransaksi) TableNameHistorical() string {
	return "batal_transaksi_historical"
}

func (b BatalTransaksi) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, b.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", b.TableNameHistorical())
	return nil
}

func (b BatalTransaksi) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":              b.ID,
		"id_transaksi":    b.IdTransaksi,
		"dibatalkan_oleh": b.DibatalkanOleh,
		"alasan":          b.Alasan,
		"created_at":      b.CreatedAt,
		"updated_at":      b.UpdatedAt,
		"deleted_at":      b.DeletedAt,
	}
}

type Follower struct {
	ID         int64
	IdFollower int64
	Pengguna   Pengguna
	IdFollowed int64
	Seller     Seller
	CreatedAt  time.Time
	UpdatedAt  time.Time
	// DeletedAt  time.Time
}

func (j Follower) TableNameSotReplica() string {
	return "follower_sot_replica"
}

func (f Follower) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_follower bigint,
		id_followed bigint,
		created_at timestamp,
		updated_at timestamp,
		PRIMARY KEY (id)
	)`, f.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", f.TableNameSotReplica())
	return nil
}

func (Follower) TableNameHistorical() string {
	return "follower_historical"
}

func (f Follower) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_follower bigint,
		id_followed bigint,
		created_at timestamp,
		updated_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, f.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", f.TableNameHistorical())
	return nil
}

func (f Follower) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":          f.ID,
		"id_follower": f.IdFollower,
		"id_followed": f.IdFollowed,
		"created_at":  f.CreatedAt,
		"updated_at":  f.UpdatedAt,
	}
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
	DeletedAt       time.Time
}

func (j AlamatGudang) TableNameSotReplica() string {
	return "alamat_gudang_sot_replica"
}

func (a AlamatGudang) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		panggilan_alamat text,
		nomor_telefon text,
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
		PRIMARY KEY (id)
	)`, a.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", a.TableNameSotReplica())
	return nil
}

func (AlamatGudang) TableNameHistorical() string {
	return "alamat_gudang_historical"
}

func (a AlamatGudang) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		panggilan_alamat text,
		nomor_telefon text,
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

func (a AlamatGudang) ParseToCUDType() map[string]interface{} {

	return map[string]interface{}{
		"id":               a.ID,
		"id_seller":        a.IDSeller,
		"panggilan_alamat": a.PanggilanAlamat,
		"nomor_telefon":    a.NomorTelephone,
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
		"deleted_at":       a.DeletedAt,
	}
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
	// REVISI: tambah CreatedAt, UpdatedAt, DeletedAt sesuai sot_models
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (j DistributorData) TableNameSotReplica() string {
	return "distributor_data_sot_replica"
}

func (d DistributorData) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// REVISI: tambah created_at, updated_at, deleted_at sesuai sot_models
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
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, d.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", d.TableNameSotReplica())
	return nil
}

func (DistributorData) TableNameHistorical() string {
	return "distributor_data_historical_data"
}

func (d DistributorData) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	// REVISI: tambah kolom created_at, updated_at, deleted_at
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
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, d.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", d.TableNameHistorical())
	return nil
}

func (d DistributorData) ParseToCUDType() map[string]interface{} {
	// REVISI: tambah created_at, updated_at, deleted_at

	return map[string]interface{}{
		"id":                           d.ID,
		"seller_id":                    d.SellerId,
		"nama_perusahaan":              d.NamaPerusahaan,
		"nib":                          d.NIB,
		"npwp":                         d.NPWP,
		"dokumen_izin_distributor_url": d.DokumenIzinDistributorUrl,
		"alasan":                       d.Alasan,
		"status":                       d.Status,
		"created_at":                   d.CreatedAt,
		"updated_at":                   d.UpdatedAt,
		"deleted_at":                   d.DeletedAt,
	}
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
	// REVISI: tambah CreatedAt, UpdatedAt, DeletedAt sesuai sot_models
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// REVISI: CreateSotReplicaTable untuk BrandData sebelumnya tidak ada, sekarang ditambahkan
func (b BrandData) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		seller_id int,
		nama_perusahaan text,
		negara_asal text,
		lembaga_pendaftaran text,
		nomor_pendaftaran_merek text,
		sertifikat_merek_url text,
		dokumen_perwakilan_url text,
		nib text,
		npwp text,
		alasan text,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, b.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", b.TableNameSotReplica())
	return nil
}

func (BrandData) TableNameHistorical() string {
	return "brand_data_historical"
}

func (BrandData) TableNameSotReplica() string {
	return "brand_data_sot_replica"
}

// REVISI: CreateHistoricalTable untuk BrandData sebelumnya tidak ada, sekarang ditambahkan
func (b BrandData) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		seller_id int,
		nama_perusahaan text,
		negara_asal text,
		lembaga_pendaftaran text,
		nomor_pendaftaran_merek text,
		sertifikat_merek_url text,
		dokumen_perwakilan_url text,
		nib text,
		npwp text,
		alasan text,
		status text,
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

func (b BrandData) ParseToCUDType() map[string]interface{} {
	// REVISI: tambah created_at, updated_at, deleted_at

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
		"created_at":              b.CreatedAt,
		"updated_at":              b.UpdatedAt,
		"deleted_at":              b.DeletedAt,
	}
}

type Etalase struct {
	ID           int64
	SellerID     int64
	Seller       Seller
	Nama         string
	Deskripsi    string
	JumlahBarang int32
	// REVISI: tambah CreatedAt, UpdatedAt, DeletedAt sesuai sot_models
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (e Etalase) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		nama text,
		deskripsi text,
		jumlah_barang int,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, e.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", e.TableNameSotReplica())
	return nil
}

func (Etalase) TableNameHistorical() string {
	return "etalase_historical"
}

func (Etalase) TableNameSotReplica() string {
	return "etalase_sot_replica"
}

// REVISI: CreateHistoricalTable untuk Etalase sebelumnya salah (pakai skema BrandData)
func (e Etalase) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		nama text,
		deskripsi text,
		jumlah_barang int,
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

func (e Etalase) ParseToCUDType() map[string]interface{} {
	// REVISI: tambah created_at, updated_at, deleted_at; hapus kolom yang tidak relevan

	return map[string]interface{}{
		"id":            e.ID,
		"id_seller":     e.SellerID,
		"nama":          e.Nama,
		"deskripsi":     e.Deskripsi,
		"jumlah_barang": e.JumlahBarang,
		"created_at":    e.CreatedAt,
		"updated_at":    e.UpdatedAt,
		"deleted_at":    e.DeletedAt,
	}
}

type BarangKeEtalase struct {
	ID            int64
	IdEtalase     int64
	Etalase       Etalase
	IdBarangInduk int64
	BarangInduk   BarangInduk
}

func (j BarangKeEtalase) TableNameSotReplica() string {
	return "barang_ke_etalase_sot_replica"
}

func (b BarangKeEtalase) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_etalase bigint,
		id_barang_induk bigint,
		PRIMARY KEY (id)
	)`, b.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", b.TableNameSotReplica())
	return nil
}

func (BarangKeEtalase) TableNameHistorical() string {
	return "barang_ke_etalase_historical"
}

func (b BarangKeEtalase) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_etalase bigint,
		id_barang_induk bigint,
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

func (b BarangKeEtalase) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":              b.ID,
		"id_etalase":      b.IdEtalase,
		"id_barang_induk": b.IdBarangInduk,
	}
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
	DeletedAt     time.Time
}

func (j DiskonProduk) TableNameSotReplica() string {
	return "diskon_produk_sot_replica"
}

func (d DiskonProduk) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		nama text,
		deskripsi text,
		diskon_persen double,
		berlaku_mulai timestamp,
		berlaku_sampai timestamp,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, d.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", d.TableNameSotReplica())
	return nil
}

func (DiskonProduk) TableNameHistorical() string {
	return "diskon_produk_historical"
}

func (d DiskonProduk) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
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
	)`, d.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", d.TableNameHistorical())
	return nil
}

func (d DiskonProduk) ParseToCUDType() map[string]interface{} {

	return map[string]interface{}{
		"id":             d.ID,
		"id_seller":      d.SellerId,
		"nama":           d.Nama,
		"deskripsi":      d.Deskripsi,
		"diskon_persen":  d.DiskonPersen,
		"berlaku_mulai":  d.BerlakuMulai,
		"berlaku_sampai": d.BerlakuSampai,
		"status":         d.Status,
		"created_at":     d.CreatedAt,
		"updated_at":     d.UpdatedAt,
		"deleted_at":     d.DeletedAt,
	}
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
	// REVISI: hapus DeletedAt (hard delete sesuai sot_models)
}

func (j BarangDiDiskon) TableNameSotReplica() string {
	return "barang_di_diskon_sot_replica"
}

func (b BarangDiDiskon) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// REVISI: hapus deleted_at (hard delete sesuai sot_models)
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		id_diskon bigint,
		id_barang_induk int,
		id_kategori_barang bigint,
		status text,
		created_at timestamp,
		updated_at timestamp,
		PRIMARY KEY (id)
	)`, b.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", b.TableNameSotReplica())
	return nil
}

func (BarangDiDiskon) TableNameHistorical() string {
	return "barang_di_diskon_historical"
}

func (b BarangDiDiskon) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	// REVISI: hapus kolom deleted_at (hard delete)
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		id_diskon bigint,
		id_barang_induk int,
		id_kategori_barang bigint,
		status text,
		created_at timestamp,
		updated_at timestamp,
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

func (b BarangDiDiskon) ParseToCUDType() map[string]interface{} {
	// REVISI: hapus deleted_at (hard delete)
	return map[string]interface{}{
		"id":                 b.ID,
		"id_seller":          b.SellerId,
		"id_diskon":          b.IdDiskon,
		"id_barang_induk":    b.IdBarangInduk,
		"id_kategori_barang": b.IdKategoriBarang,
		"status":             b.Status,
		"created_at":         b.CreatedAt,
		"updated_at":         b.UpdatedAt,
	}
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
	DeletedAt    time.Time
}

func (j InformasiKurir) TableNameSotReplica() string {
	return "informasi_kurir_sot_replica"
}

func (i InformasiKurir) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kurir bigint,
		tanggal_lahir text,
		alasan text,
		informasi_ktp boolean,
		informasi_sim boolean,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, i.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", i.TableNameSotReplica())
	return nil
}

func (InformasiKurir) TableNameHistorical() string {
	return "informasi_kurir_historical"
}

func (i InformasiKurir) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kurir bigint,
		tanggal_lahir text,
		alasan text,
		informasi_ktp boolean,
		informasi_sim boolean,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, i.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", i.TableNameHistorical())
	return nil
}

func (i InformasiKurir) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":            i.ID,
		"id_kurir":      i.IDkurir,
		"tanggal_lahir": i.TanggalLahir,
		"alasan":        i.Alasan,
		"informasi_ktp": i.Ktp,
		"informasi_sim": i.InformasiSim,
		"status":        i.Status,
		"created_at":    i.CreatedAt,
		"updated_at":    i.UpdatedAt,
		"deleted_at":    i.DeletedAt,
	}
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
	DeletedAt      time.Time
}

func (j InformasiKendaraanKurir) TableNameSotReplica() string {
	return "informasi_kendaraan_kurir_sot_replica"
}

func (i InformasiKendaraanKurir) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kurir bigint,
		jenis_kendaraan text,
		nama_kendaraan text,
		roda_kendaraan text,
		informasi_stnk boolean,
		informasi_bpkb boolean,
		nomor_rangka text,
		nomor_mesin text,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, i.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", i.TableNameSotReplica())
	return nil
}

func (InformasiKendaraanKurir) TableNameHistorical() string {
	return "informasi_kendaraan_kurir_historical"
}

func (i InformasiKendaraanKurir) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kurir bigint,
		jenis_kendaraan text,
		nama_kendaraan text,
		roda_kendaraan text,
		informasi_stnk boolean,
		informasi_bpkb boolean,
		nomor_rangka text,
		nomor_mesin text,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, i.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", i.TableNameHistorical())
	return nil
}

func (i InformasiKendaraanKurir) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":              i.ID,
		"id_kurir":        i.IDkurir,
		"jenis_kendaraan": i.JenisKendaraan,
		"nama_kendaraan":  i.NamaKendaraan,
		"roda_kendaraan":  i.RodaKendaraan,
		"informasi_stnk":  i.STNK,
		"informasi_bpkb":  i.BPKB,
		"nomor_rangka":    i.NoRangka,
		"nomor_mesin":     i.NoMesin,
		"status":          i.Status,
		"created_at":      i.CreatedAt,
		"updated_at":      i.UpdatedAt,
		"deleted_at":      i.DeletedAt,
	}
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
}

func (j AlamatKurir) TableNameSotReplica() string {
	return "alamat_kurir_sot_replica"
}

func (a AlamatKurir) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kurir bigint,
		panggilan_alamat text,
		nomor_telefon text,
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
		PRIMARY KEY (id)
	)`, a.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", a.TableNameSotReplica())
	return nil
}

func (AlamatKurir) TableNameHistorical() string {
	return "alamat_kurir_historical"
}

func (a AlamatKurir) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kurir bigint,
		panggilan_alamat text,
		nomor_telefon text,
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

func (a AlamatKurir) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":               a.ID,
		"id_kurir":         a.IdKurir,
		"panggilan_alamat": a.PanggilanAlamat,
		"nomor_telefon":    a.NomorTelephone,
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
	}
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
	Selesai         time.Time
	JenisKendaraan  string
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

func (j BidKurirData) TableNameSotReplica() string {
	return "bid_kurir_data_sot_replica"
}

func (b BidKurirData) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, b.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", b.TableNameSotReplica())
	return nil
}

func (BidKurirData) TableNameHistorical() string {
	return "bid_kurir_data_historical"
}

func (b BidKurirData) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, b.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", b.TableNameHistorical())
	return nil
}

func (b BidKurirData) ParseToCUDType() map[string]interface{} {

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
		"deleted_at":       b.DeletedAt,
	}
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
	DeletedAt    time.Time
}

func (j BidKurirNonEksScheduler) TableNameSotReplica() string {
	return "bid_kurir_non_eks_scheduler_sot_replica"
}

func (b BidKurirNonEksScheduler) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
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
		PRIMARY KEY (id)
	)`, b.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", b.TableNameSotReplica())
	return nil
}

func (BidKurirNonEksScheduler) TableNameHistorical() string {
	return "bid_kurir_non_eks_scheduler_historical"
}

func (b BidKurirNonEksScheduler) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
	)`, b.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", b.TableNameHistorical())
	return nil
}

func (b BidKurirNonEksScheduler) ParseToCUDType() map[string]interface{} {

	return map[string]interface{}{
		"id":            b.ID,
		"id_bid":        b.IdBid,
		"id_kurir":      b.IdKurir,
		"urutan":        b.Urutan,
		"id_pengiriman": b.IdPengiriman,
		"status":        b.Status,
		"created_at":    b.CreatedAt,
		"updated_at":    b.UpdatedAt,
		"deleted_at":    b.DeletedAt,
	}
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
	DeletedAt           time.Time
}

func (j BidKurirEksScheduler) TableNameSotReplica() string {
	return "bid_kurir_eks_scheduler_sot_replica"
}

func (b BidKurirEksScheduler) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_bid bigint,
		id_kurir bigint,
		urutan tinyint,
		id_pengiriman_ekspedisi bigint,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, b.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", b.TableNameSotReplica())
	return nil
}

func (BidKurirEksScheduler) TableNameHistorical() string {
	return "bid_kurir_eks_scheduler_historical"
}

func (b BidKurirEksScheduler) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_bid bigint,
		id_kurir bigint,
		urutan tinyint,
		id_pengiriman_ekspedisi bigint,
		status text,
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

func (b BidKurirEksScheduler) ParseToCUDType() map[string]interface{} {

	return map[string]interface{}{
		"id":                      b.ID,
		"id_bid":                  b.IdBid,
		"id_kurir":                b.IdKurir,
		"urutan":                  b.Urutan,
		"id_pengiriman_ekspedisi": b.IdPengirimanEks,
		"status":                  b.Status,
		"created_at":              b.CreatedAt,
		"updated_at":              b.UpdatedAt,
		"deleted_at":              b.DeletedAt,
	}
}

// ///////////////////////////////////////////////////////////////////////////////////////////
// SOT REPLICA TABLES
// ///////////////////////////////////////////////////////////////////////////////////////////
