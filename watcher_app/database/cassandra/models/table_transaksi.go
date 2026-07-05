package cass_models

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"gorm.io/gorm"
)

// =========================================================================
// PEMBAYARAN
// =========================================================================

type Pembayaran struct {
	ID              int64
	IdPengguna      int64
	Pengguna        Pengguna
	KodeTransaksiPG string
	KodeOrderSistem string
	Provider        string
	Total           int32
	PaymentType     string
	PaidAt          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

func (Pembayaran) TableNameHistorical() string {
	return "pembayaran_historical"
}

func (p Pembayaran) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		kode_transaksi_pg text,
		kode_order_sistem text,
		provider text,
		total int,
		payment_type text,
		paid_at text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, p.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", p.TableNameHistorical())
	return nil
}

func (p Pembayaran) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                p.ID,
		"id_pengguna":       p.IdPengguna,
		"kode_transaksi_pg": p.KodeTransaksiPG,
		"kode_order_sistem": p.KodeOrderSistem,
		"provider":          p.Provider,
		"total":             p.Total,
		"payment_type":      p.PaymentType,
		"paid_at":           p.PaidAt,
		"created_at":        p.CreatedAt,
		"updated_at":        p.UpdatedAt,
		"deleted_at":        p.DeletedAt,
	}
}

func (p Pembayaran) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, p.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", p.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", p.TableNameHistorical())
	return nil
}

func (p Pembayaran) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		kode_transaksi_pg text,
		kode_order_sistem text,
		provider text,
		total int,
		payment_type text,
		paid_at text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, p.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", p.TableNameSotReplica())
	return nil
}

// =========================================================================
// TRANSAKSI
// =========================================================================

type Transaksi struct {
	ID                  int64
	IdPengguna          int64
	Pengguna            Pengguna
	IdSeller            int32
	Seller              Seller
	IdBarangInduk       int64
	BarangInduk         BarangInduk
	IdKategoriBarang    int64
	KategoriBarang      KategoriBarang
	IdAlamatPengguna    int64
	AlamatPengguna      AlamatPengguna
	IdAlamatGudang      int64
	AlamatGudang        AlamatGudang
	IdAlamatEkspedisi   int64
	IdPembayaran        int64
	Pembayaran          Pembayaran
	KendaraanPengiriman string
	JenisPengiriman     string
	JarakTempuh         string
	BeratTotalKg        int16
	KodeOrderSistem     string
	KodeResiEkspedisi   *string
	Status              string
	DibatalkanOleh      *string
	Catatan             string
	KuantitasBarang     int32
	IsEkspedisi         bool
	SellerPaid          int64
	KurirPaid           int64
	EkspedisiPaid       int64
	Total               int64
	Reviewed            bool
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}

func (Transaksi) TableNameHistorical() string {
	return "transaksi_historical"
}

func (t Transaksi) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_seller int,
		id_barang_induk bigint,
		id_kategori_barang bigint,
		id_alamat_pengguna bigint,
		id_alamat_gudang bigint,
		id_alamat_ekspedisi bigint,
		id_pembayaran bigint,
		kendaraan_pengiriman text,
		jenis_pengiriman text,
		jarak_tempuh text,
		berat_total_kg smallint,
		kode_order_sistem text,
		kode_resi_ekspedisi text,
		status text,
		dibatalkan_oleh text,
		catatan text,
		kuantitas_barang int,
		is_ekspedisi boolean,
		seller_paid bigint,
		kurir_paid bigint,
		ekspedisi_paid bigint,
		total bigint,
		reviewed boolean,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, t.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", t.TableNameHistorical())
	return nil
}

func (t Transaksi) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                   t.ID,
		"id_pengguna":          t.IdPengguna,
		"id_seller":            t.IdSeller,
		"id_barang_induk":      t.IdBarangInduk,
		"id_kategori_barang":   t.IdKategoriBarang,
		"id_alamat_pengguna":   t.IdAlamatPengguna,
		"id_alamat_gudang":     t.IdAlamatGudang,
		"id_alamat_ekspedisi":  t.IdAlamatEkspedisi,
		"id_pembayaran":        t.IdPembayaran,
		"kendaraan_pengiriman": t.KendaraanPengiriman,
		"jenis_pengiriman":     t.JenisPengiriman,
		"jarak_tempuh":         t.JarakTempuh,
		"berat_total_kg":       t.BeratTotalKg,
		"kode_order_sistem":    t.KodeOrderSistem,
		"kode_resi_ekspedisi":  t.KodeResiEkspedisi,
		"status":               t.Status,
		"dibatalkan_oleh":      t.DibatalkanOleh,
		"catatan":              t.Catatan,
		"kuantitas_barang":     t.KuantitasBarang,
		"is_ekspedisi":         t.IsEkspedisi,
		"seller_paid":          t.SellerPaid,
		"kurir_paid":           t.KurirPaid,
		"ekspedisi_paid":       t.EkspedisiPaid,
		"total":                t.Total,
		"reviewed":             t.Reviewed,
		"created_at":           t.CreatedAt,
		"updated_at":           t.UpdatedAt,
		"deleted_at":           t.DeletedAt,
	}
}

func (t Transaksi) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, t.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", t.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", t.TableNameHistorical())
	return nil
}

func (t Transaksi) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_seller int,
		id_barang_induk bigint,
		id_kategori_barang bigint,
		id_alamat_pengguna bigint,
		id_alamat_gudang bigint,
		id_alamat_ekspedisi bigint,
		id_pembayaran bigint,
		kendaraan_pengiriman text,
		jenis_pengiriman text,
		jarak_tempuh text,
		berat_total_kg smallint,
		kode_order_sistem text,
		kode_resi_ekspedisi text,
		status text,
		dibatalkan_oleh text,
		catatan text,
		kuantitas_barang int,
		is_ekspedisi boolean,
		seller_paid bigint,
		kurir_paid bigint,
		ekspedisi_paid bigint,
		total bigint,
		reviewed boolean,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, t.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", t.TableNameSotReplica())
	return nil
}

// =========================================================================
// TRANSAKSI FAILED
// =========================================================================

type TransaksiFailed struct {
	ID                  int64
	IdPengguna          int64
	IdSeller            int32
	IdBarangInduk       int32
	IdKategoriBarang    int64
	IdAlamatPengguna    int64
	IdAlamatGudang      int64
	IdAlamatEkspedisi   int64
	IdPembayaran        int64
	KendaraanPengiriman string
	JenisPengiriman     string
	JarakTempuh         string
	BeratTotalKg        int16
	KodeOrderSistem     string
	KodeResiEkspedisi   *string
	Status              string
	DibatalkanOleh      *string
	Catatan             string
	KuantitasBarang     int32
	IsEkspedisi         bool
	SellerPaid          int64
	KurirPaid           int64
	EkspedisiPaid       int64
	Total               int64
	Reviewed            bool
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}

func (TransaksiFailed) TableNameHistorical() string {
	return "transaksi_failed_historical"
}

func (t TransaksiFailed) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_seller int,
		id_barang_induk int,
		id_kategori_barang bigint,
		id_alamat_pengguna bigint,
		id_alamat_gudang bigint,
		id_alamat_ekspedisi bigint,
		id_pembayaran bigint,
		kendaraan_pengiriman text,
		jenis_pengiriman text,
		jarak_tempuh text,
		berat_total_kg smallint,
		kode_order_sistem text,
		kode_resi_ekspedisi text,
		status text,
		dibatalkan_oleh text,
		catatan text,
		kuantitas_barang int,
		is_ekspedisi boolean,
		seller_paid bigint,
		kurir_paid bigint,
		ekspedisi_paid bigint,
		total bigint,
		reviewed boolean,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, t.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", t.TableNameHistorical())
	return nil
}

func (t TransaksiFailed) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                   t.ID,
		"id_pengguna":          t.IdPengguna,
		"id_seller":            t.IdSeller,
		"id_barang_induk":      t.IdBarangInduk,
		"id_kategori_barang":   t.IdKategoriBarang,
		"id_alamat_pengguna":   t.IdAlamatPengguna,
		"id_alamat_gudang":     t.IdAlamatGudang,
		"id_alamat_ekspedisi":  t.IdAlamatEkspedisi,
		"id_pembayaran":        t.IdPembayaran,
		"kendaraan_pengiriman": t.KendaraanPengiriman,
		"jenis_pengiriman":     t.JenisPengiriman,
		"jarak_tempuh":         t.JarakTempuh,
		"berat_total_kg":       t.BeratTotalKg,
		"kode_order_sistem":    t.KodeOrderSistem,
		"kode_resi_ekspedisi":  t.KodeResiEkspedisi,
		"status":               t.Status,
		"dibatalkan_oleh":      t.DibatalkanOleh,
		"catatan":              t.Catatan,
		"kuantitas_barang":     t.KuantitasBarang,
		"is_ekspedisi":         t.IsEkspedisi,
		"seller_paid":          t.SellerPaid,
		"kurir_paid":           t.KurirPaid,
		"ekspedisi_paid":       t.EkspedisiPaid,
		"total":                t.Total,
		"reviewed":             t.Reviewed,
		"created_at":           t.CreatedAt,
		"updated_at":           t.UpdatedAt,
		"deleted_at":           t.DeletedAt,
	}
}

func (t TransaksiFailed) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, t.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", t.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", t.TableNameHistorical())
	return nil
}

func (t TransaksiFailed) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_pengguna bigint,
		id_seller int,
		id_barang_induk int,
		id_kategori_barang bigint,
		id_alamat_pengguna bigint,
		id_alamat_gudang bigint,
		id_alamat_ekspedisi bigint,
		id_pembayaran bigint,
		kendaraan_pengiriman text,
		jenis_pengiriman text,
		jarak_tempuh text,
		berat_total_kg smallint,
		kode_order_sistem text,
		kode_resi_ekspedisi text,
		status text,
		dibatalkan_oleh text,
		catatan text,
		kuantitas_barang int,
		is_ekspedisi boolean,
		seller_paid bigint,
		kurir_paid bigint,
		ekspedisi_paid bigint,
		total bigint,
		reviewed boolean,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, t.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", t.TableNameSotReplica())
	return nil
}
