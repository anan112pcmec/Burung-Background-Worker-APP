package historical_models

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

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
	Pencatatan
}

func (BarangInduk) TableName() string {
	return "barang_induk_historical"
}

func (b *BarangInduk) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct BarangInduk dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id int,
		seller_id int,
		id_diskon bigint,
		nama_barang text,
		jenis_barang text,
		deskripsi text,
		original_kategori bigint,
		harga_kategoris int,
		created_at timestamp,
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

func (b *BarangInduk) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":                b.ID,
		"seller_id":         b.SellerID,
		"id_diskon":         b.IdDiskon,
		"nama_barang":       b.NamaBarang,
		"jenis_barang":      b.JenisBarang,
		"deskripsi":         b.Deskripsi,
		"original_kategori": b.OriginalKategori,
		"harga_kategoris":   b.HargaKategoris,
		"created_at":        b.CreatedAt,
		"tahun_update":      b.TahunUpdate,
		"bulan_update":      b.BulanUpdate,
		"event_time":        b.EventTime,
	}
}

// DropTable disesuaikan menggunakan b.TableName() secara dinamis
func (b *BarangInduk) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, b.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", b.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", b.TableName())
	return nil
}

type KategoriBarang struct {
	ID             int64
	SellerID       int
	IdBarangInduk  int
	IDAlamat       int64
	IdRekening     int64
	Nama           string
	Deskripsi      string
	Warna          string
	Stok           int
	Harga          int
	PotonganDiskon int
	BeratGram      int16
	DimensiPanjang int16
	DimensiLebar   int16
	Sku            string
	IsOriginal     bool
	CreatedAt      time.Time
	Pencatatan
}

func (KategoriBarang) TableName() string {
	return "kategori_barang_historical"
}

func (k *KategoriBarang) CreateTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		seller_id int,
		id_barang_induk int,
		id_alamat bigint,
		id_rekening bigint,
		nama text,
		deskripsi text,
		warna text,
		stok int,
		harga int,
		potongan_diskon int,
		berat_gram smallint,
		dimensi_panjang smallint,
		dimensi_lebar smallint,
		sku text,
		is_original boolean,
		created_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, k.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", k.TableName())
	return nil
}

func (k *KategoriBarang) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":              k.ID,
		"seller_id":       k.SellerID,
		"id_barang_induk": k.IdBarangInduk,
		"id_alamat":       k.IDAlamat,
		"id_rekening":     k.IdRekening,
		"nama":            k.Nama,
		"deskripsi":       k.Deskripsi,
		"warna":           k.Warna,
		"stok":            k.Stok,
		"harga":           k.Harga,
		"potongan_diskon": k.PotonganDiskon,
		"berat_gram":      k.BeratGram,
		"dimensi_panjang": k.DimensiPanjang,
		"dimensi_lebar":   k.DimensiLebar,
		"sku":             k.Sku,
		"is_original":     k.IsOriginal,
		"created_at":      k.CreatedAt,
		"tahun_update":    k.TahunUpdate,
		"bulan_update":    k.BulanUpdate,
		"event_time":      k.EventTime,
	}
}

// DropTable disesuaikan menggunakan k.TableName() secara dinamis
func (k *KategoriBarang) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, k.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", k.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", k.TableName())
	return nil
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
	Pencatatan
}

func (VarianBarang) TableName() string {
	return "varian_barang_historical"
}

func (v *VarianBarang) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct VarianBarang dan Pencatatan
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
	)`, v.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", v.TableName())
	return nil
}

func (v *VarianBarang) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":              v.ID,
		"id_barang_induk": v.IdBarangInduk,
		"id_kategori":     v.IdKategori,
		"id_transaksi":    v.IdTransaksi,
		"sku":             v.Sku,
		"status":          v.Status,
		"hold_by":         v.HoldBy,
		"holder_entity":   v.HolderEntity,
		"tahun_update":    v.TahunUpdate,
		"bulan_update":    v.BulanUpdate,
		"event_time":      v.EventTime,
	}
}

// DropTable disesuaikan menggunakan v.TableName() secara dinamis
func (v *VarianBarang) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, v.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", v.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", v.TableName())
	return nil
}
