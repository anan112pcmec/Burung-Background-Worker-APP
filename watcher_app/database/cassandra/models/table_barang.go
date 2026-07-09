package cass_models

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"gorm.io/gorm"
)

// REVISI: tambah interface BarangContract sesuai sot_models
type BarangContract interface {
	Validating() string
}

type BarangInduk struct {
	ID               int32
	SellerID         int32
	Seller           Seller
	IdDiskon         int64
	NamaBarang       string
	JenisBarang      string
	Deskripsi        string
	OriginalKategori int64
	HargaKategoris   int32
	CreatedAt        time.Time
	// REVISI: tambah UpdatedAt dan DeletedAt sesuai sot_models
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (b BarangInduk) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// REVISI: tambah updated_at dan deleted_at
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id int,
		id_seller int,
		id_diskon bigint,
		nama_barang text,
		jenis_barang text,
		deskripsi text,
		original_kategori bigint,
		harga_kategori_barang int,
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

func (BarangInduk) TableNameHistorical() string {
	return "barang_induk_historical"
}

// REVISI: tambah implementasi BarangContract interface
func (b BarangInduk) Validating() string {
	return b.TableNameHistorical()
}

func (b BarangInduk) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	// REVISI: tambah kolom updated_at dan deleted_at
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id int,
		id_seller int,
		id_diskon bigint,
		nama_barang text,
		jenis_barang text,
		deskripsi text,
		original_kategori bigint,
		harga_kategori_barang int,
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

func (b BarangInduk) ParseToCUDType() map[string]interface{} {
	// REVISI: tambah updated_at dan deleted_at
	var deletedAtInterface interface{} = nil
	if b.DeletedAt.Valid {
		deletedAtInterface = b.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":                    b.ID,
		"id_seller":             b.SellerID,
		"id_diskon":             b.IdDiskon,
		"nama_barang":           b.NamaBarang,
		"jenis_barang":          b.JenisBarang,
		"deskripsi":             b.Deskripsi,
		"original_kategori":     b.OriginalKategori,
		"harga_kategori_barang": b.HargaKategoris,
		"created_at":            b.CreatedAt,
		"updated_at":            b.UpdatedAt,
		"deleted_at":            deletedAtInterface,
	}
}

type KategoriBarang struct {
	ID             int64
	SellerID       int32
	IdBarangInduk  int32
	IDAlamat       int64
	IDRekening     int64
	Nama           string
	Deskripsi      string
	Warna          string
	Stok           int32
	Harga          int32
	PotonganDiskon int32
	BeratGram      int16
	DimensiPanjang int16
	DimensiLebar   int16
	Sku            string
	IsOriginal     bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

func (k KategoriBarang) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// REVISI: tambah updated_at dan deleted_at
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		id_barang_induk int,
		id_alamat_gudang bigint,
		id_rekening bigint,
		nama text,
		deskripsi text,
		warna text,
		stok int,
		harga int,
		diskon int,
		berat_gram smallint,
		dimensi_panjang_cm smallint,
		dimensi_lebar_cm smallint,
		sku text,
		is_original boolean,
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

func (KategoriBarang) TableNameHistorical() string {
	return "kategori_barang_historical"
}

// REVISI: tambah implementasi BarangContract interface
func (k KategoriBarang) Validating() string {
	return k.TableNameHistorical()
}

func (k KategoriBarang) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	// REVISI: tambah kolom updated_at dan deleted_at
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		id_barang_induk int,
		id_alamat_gudang bigint,
		id_rekening bigint,
		nama text,
		deskripsi text,
		warna text,
		stok int,
		harga int,
		diskon int,
		berat_gram smallint,
		dimensi_panjang_cm smallint,
		dimensi_lebar_cm smallint,
		sku text,
		is_original boolean,
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

func (k KategoriBarang) ParseToCUDType() map[string]interface{} {
	// REVISI: tambah updated_at dan deleted_at; gunakan IDRekening
	var deletedAtInterface interface{} = nil
	if k.DeletedAt.Valid {
		deletedAtInterface = k.DeletedAt.Time
	}

	return map[string]interface{}{
		"id":                 k.ID,
		"id_seller":          k.SellerID,
		"id_barang_induk":    k.IdBarangInduk,
		"id_alamat_gudang":   k.IDAlamat,
		"id_rekening":        k.IDRekening,
		"nama":               k.Nama,
		"deskripsi":          k.Deskripsi,
		"warna":              k.Warna,
		"stok":               k.Stok,
		"harga":              k.Harga,
		"diskon":             k.PotonganDiskon,
		"berat_gram":         k.BeratGram,
		"dimensi_panjang_cm": k.DimensiPanjang,
		"dimensi_lebar_cm":   k.DimensiLebar,
		"sku":                k.Sku,
		"is_original":        k.IsOriginal,
		"created_at":         k.CreatedAt,
		"updated_at":         k.UpdatedAt,
		"deleted_at":         deletedAtInterface,
	}
}

type VarianBarang struct {
	ID            int64
	IdBarangInduk int32
	BarangInduk   BarangInduk
	IdKategori    int64
	Kategori      KategoriBarang
	IdTransaksi   int64
	Sku           string
	Status        string
	HoldBy        int64
	HolderEntity  string
}

func (v VarianBarang) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_barang_induk int,
		id_kategori bigint,
		id_transaksi bigint,
		sku text,
		status text,
		hold_by bigint,
		holder_entity text,
		PRIMARY KEY (id)
	)`, v.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel sot_replica %s\n", v.TableNameSotReplica())
	return nil
}

func (VarianBarang) TableNameHistorical() string {
	return "varian_barang_historical"
}

// REVISI: tambah implementasi BarangContract interface
func (v VarianBarang) Validating() string {
	return v.TableNameHistorical()
}

func (v VarianBarang) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_barang_induk int,
		id_kategori bigint,
		id_transaksi bigint,
		sku text,
		status text,
		hold_by bigint,
		holder_entity text,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, v.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", v.TableNameHistorical())
	return nil
}

func (v VarianBarang) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":              v.ID,
		"id_barang_induk": v.IdBarangInduk,
		"id_kategori":     v.IdKategori,
		"id_transaksi":    v.IdTransaksi,
		"sku":             v.Sku,
		"status":          v.Status,
		"hold_by":         v.HoldBy,
		"holder_entity":   v.HolderEntity,
	}
}
