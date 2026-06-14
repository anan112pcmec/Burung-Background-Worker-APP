package historical_models

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"gorm.io/gorm"
)

type Pengiriman struct {
	ID                int64
	IdTransaksi       int64
	Transaksi         Transaksi
	IdSeller          int64
	Seller            Seller
	IdAlamatGudang    int64
	AlamatGudang      AlamatGudang
	IdAlamatPengguna  int64
	AlamatPengguna    AlamatPengguna
	IdKurir           *int64
	BeratBarang       int16
	KendaraanRequired string
	JenisPengiriman   string
	JarakTempuh       string
	KurirPaid         int64
	Status            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
	Pencatatan
}

func (Pengiriman) TableName() string {
	return "pengiriman_historical"
}

func (p *Pengiriman) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan seluruh field di struct Pengiriman dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_transaksi bigint,
		id_seller bigint,
		id_alamat_gudang bigint,
		id_alamat_pengguna bigint,
		id_kurir bigint,
		berat_barang smallint,
		kendaraan_required text,
		jenis_pengiriman text,
		jarak_tempuh text,
		kurir_paid bigint,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, p.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", p.TableName())
	return nil
}

func (p *Pengiriman) ParseToInsertType() map[string]interface{} {

	return map[string]interface{}{
		"id":                 p.ID,
		"id_transaksi":       p.IdTransaksi,
		"id_seller":          p.IdSeller,
		"id_alamat_gudang":   p.IdAlamatGudang,
		"id_alamat_pengguna": p.IdAlamatPengguna,
		"id_kurir":           p.IdKurir,
		"berat_barang":       p.BeratBarang,
		"kendaraan_required": p.KendaraanRequired,
		"jenis_pengiriman":   p.JenisPengiriman,
		"jarak_tempuh":       p.JarakTempuh,
		"kurir_paid":         p.KurirPaid,
		"status":             p.Status,
		"created_at":         p.CreatedAt,
		"updated_at":         p.UpdatedAt,
		"deleted_at":         p.DeletedAt,
		"tahun_update":       p.TahunUpdate,
		"bulan_update":       p.BulanUpdate,
		"event_time":         p.EventTime,
	}
}

// DropTable disesuaikan menggunakan p.TableName() secara dinamis
func (p *Pengiriman) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, p.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", p.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", p.TableName())
	return nil
}

type JejakPengiriman struct {
	ID           int64
	IdPengiriman int64
	Pengiriman   Pengiriman
	Lokasi       string
	Keterangan   string
	Latitude     float64
	Longtitude   float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Pencatatan
}

func (JejakPengiriman) TableName() string {
	return "jejak_pengiriman_historical"
}

func (j *JejakPengiriman) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan seluruh field di struct JejakPengiriman dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengiriman bigint,
		lokasi text,
		keterangan text,
		latitude double,
		longtitude double,
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

func (j *JejakPengiriman) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":            j.ID,
		"id_pengiriman": j.IdPengiriman,
		"lokasi":        j.Lokasi,
		"keterangan":    j.Keterangan,
		"latitude":      j.Latitude,
		"longtitude":    j.Longtitude,
		"created_at":    j.CreatedAt,
		"updated_at":    j.UpdatedAt,
		"deleted_at":    j.DeletedAt,
		"tahun_update":  j.TahunUpdate,
		"bulan_update":  j.BulanUpdate,
		"event_time":    j.EventTime,
	}
}

func (j *JejakPengiriman) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, j.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", j.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", j.TableName())
	return nil
}

type PengirimanEkspedisi struct {
	ID                int64
	IdTransaksi       int64
	Transaksi         Transaksi
	IdSeller          int64
	Seller            Seller
	IdAlamatGudang    int64
	AlamatGudang      AlamatGudang
	IdAlamatEkspedisi int64
	IdKurir           *int64
	BeratBarang       int16
	KendaraanRequired string
	JenisPengiriman   string
	JarakTempuh       string
	KurirPaid         int64
	Status            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
	Pencatatan
}

func (PengirimanEkspedisi) TableName() string {
	return "pengiriman_ekspedisi_historical"
}

func (p *PengirimanEkspedisi) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan seluruh field di struct PengirimanEkspedisi dan komponen pencatatan historis
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_transaksi bigint,
		id_seller bigint,
		id_alamat_gudang bigint,
		id_alamat_ekspedisi bigint,
		id_kurir bigint,
		berat_barang smallint,
		kendaraan_required text,
		jenis_pengiriman text,
		jarak_tempuh text,
		kurir_paid bigint,
		status text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, p.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", p.TableName())
	return nil
}

func (p *PengirimanEkspedisi) ParseToInsertType(tahunUpdate int, bulanUpdate int, eventTime time.Time) map[string]interface{} {

	return map[string]interface{}{
		"id":                  p.ID,
		"id_transaksi":        p.IdTransaksi,
		"id_seller":           p.IdSeller,
		"id_alamat_gudang":    p.IdAlamatGudang,
		"id_alamat_ekspedisi": p.IdAlamatEkspedisi,
		"id_kurir":            p.IdKurir,
		"berat_barang":        p.BeratBarang,
		"kendaraan_required":  p.KendaraanRequired,
		"jenis_pengiriman":    p.JenisPengiriman,
		"jarak_tempuh":        p.JarakTempuh,
		"kurir_paid":          p.KurirPaid,
		"status":              p.Status,
		"created_at":          p.CreatedAt,
		"updated_at":          p.UpdatedAt,
		"deleted_at":          p.DeletedAt,
		"tahun_update":        p.TahunUpdate, // Dilewatkan via argumen fungsi karena tidak ada embedded Pencatatan
		"bulan_update":        p.BulanUpdate, // Dilewatkan via argumen fungsi karena tidak ada embedded Pencatatan
		"event_time":          p.EventTime,   // Dilewatkan via argumen fungsi karena tidak ada embedded Pencatatan
	}
}

// DropTable disesuaikan menggunakan p.TableName() secara dinamis
func (p *PengirimanEkspedisi) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, p.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", p.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", p.TableName())
	return nil
}

type JejakPengirimanEkspedisi struct {
	ID                    int64
	IdPengirimanEkspedisi int64
	Pengiriman            PengirimanEkspedisi
	Lokasi                string
	Keterangan            string
	Latitude              float64
	Longitude             float64
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt
	Pencatatan
}

func (JejakPengirimanEkspedisi) TableName() string {
	return "jejak_pengiriman_ekspedisi_historical"
}

func (j *JejakPengirimanEkspedisi) CreateTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengiriman_ekspedisi bigint,
		lokasi text,
		keterangan text,
		latitude double,
		longitude double,
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

func (j *JejakPengirimanEkspedisi) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":                      j.ID,
		"id_pengiriman_ekspedisi": j.IdPengirimanEkspedisi,
		"lokasi":                  j.Lokasi,
		"keterangan":              j.Keterangan,
		"latitude":                j.Latitude,
		"longitude":               j.Longitude,
		"created_at":              j.CreatedAt,
		"updated_at":              j.UpdatedAt,
		"deleted_at":              j.DeletedAt,
		"tahun_update":            j.TahunUpdate,
		"bulan_update":            j.BulanUpdate,
		"event_time":              j.EventTime,
	}
}

func (j *JejakPengirimanEkspedisi) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, j.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", j.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", j.TableName())
	return nil
}
