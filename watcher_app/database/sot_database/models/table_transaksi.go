package models

import (
	"time"

	"gorm.io/gorm"
)

type Pembayaran struct {
	ID              int64          `gorm:"primaryKey;autoIncrement" json:"id_pembayaran"`
	IdPengguna      int64          `gorm:"column:id_pengguna;not null" json:"id_pengguna"`
	Pengguna        Pengguna       `gorm:"foreignKey:IdPengguna;references:ID" json:"-"`
	KodeTransaksiPG string         `gorm:"column:kode_transaksi_pg;not null" json:"kode_transaksi_pg_pembayaran"`
	KodeOrderSistem string         `gorm:"column:kode_order_sistem;type:varchar(250);unique;not null" json:"kode_order_sistem_pembayaran"`
	Provider        string         `gorm:"column:provider;type:text;not null;default:''" json:"provider_pembayaran"`
	Total           int32          `gorm:"column:total;type:int4;not null;default:0" json:"total_pembayaran"`
	PaymentType     string         `gorm:"column:payment_type;type:varchar(120);not null" json:"payment_type_pembayaran"`
	PaidAt          string         `gorm:"column:paid_at;type:text;not null;default:''" json:"paid_at_pembayaran"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Pembayaran) TableName() string {
	return "pembayaran"
}

type Transaksi struct {
	ID                  int64          `gorm:"primaryKey;autoIncrement" json:"id_transaksi"`
	IdPengguna          int64          `gorm:"index;column:id_pengguna;not null" json:"id_pengguna"`
	Pengguna            Pengguna       `gorm:"foreignKey:IdPengguna;references:ID" json:"-"`
	IdSeller            int32          `gorm:"column:id_seller;not null" json:"id_seller"`
	Seller              Seller         `gorm:"foreignKey:IdSeller;references:ID" json:"-"`
	IdBarangInduk       int64          `gorm:"index;column:id_barang_induk;not null" json:"id_barang_induk"`
	BarangInduk         BarangInduk    `gorm:"foreignKey:IdBarangInduk;references:ID" json:"-"`
	IdKategoriBarang    int64          `gorm:"index;column:id_kategori_barang;not null" json:"id_kategori_barang"`
	KategoriBarang      KategoriBarang `gorm:"foreignKey:IdKategoriBarang;references:ID" json:"-"`
	IdAlamatPengguna    int64          `gorm:"index;column:id_alamat_pengguna;not null" json:"id_alamat_pengguna"`
	AlamatPengguna      AlamatPengguna `gorm:"foreignKey:IdAlamatPengguna;references:ID" json:"-"`
	IdAlamatGudang      int64          `gorm:"index;column:id_alamat_gudang;type:int8;not null" json:"id_alamat_gudang"`
	AlamatGudang        AlamatGudang   `gorm:"foreignKey:IdAlamatGudang;references:ID" json:"-"`
	IdAlamatEkspedisi   int64          `gorm:"column:id_alamat_ekspedisi;type:int8;not null" json:"id_alamat_ekspedisi"`
	IdPembayaran        int64          `gorm:"column:id_pembayaran;not null" json:"id_pembayaran"`
	Pembayaran          Pembayaran     `gorm:"foreignKey:IdPembayaran;references:ID" json:"-"`
	KendaraanPengiriman string         `gorm:"column:kendaraan_pengiriman;type:jenis_kendaraan_kurir;default:'Motor';not null" json:"kendaraan_pengiriman"`
	JenisPengiriman     string         `gorm:"column:jenis_pengiriman;type:jenis_layanan_kurir;not null" json:"jenis_pengiriman"`
	JarakTempuh         string         `gorm:"column:jarak_tempuh;not null" json:"jarak_tempuh"`
	BeratTotalKg        int16          `gorm:"column:berat_total_kg;type:int2;not null" json:"berat_total_kg_pengiriman"`
	KodeOrderSistem     string         `gorm:"column:kode_order_sistem;type:varchar(100);not null" json:"kode_order_sistem"`
	KodeResiEkspedisi   *string        `gorm:"column:kode_resi_ekspedisi;type:varchar(100)" json:"kode_resi_ekspedisi"`
	Status              string         `gorm:"index;column:status;type:status_transaksi;default:'Dibayar';not null" json:"status"`
	DibatalkanOleh      *string        `gorm:"column:dibatalkan_oleh;type:jenis_entity" json:"dibatalkan_oleh"`
	Catatan             string         `gorm:"column:catatan;type:text" json:"catatan"`
	KuantitasBarang     int32          `gorm:"column:kuantitas_barang;type:int4;not null" json:"kuantitas_barang"`
	IsEkspedisi         bool           `gorm:"column:is_ekspedisi;not null;default:false" json:"is_ekspedisi"`
	SellerPaid          int64          `gorm:"column:seller_paid;type:int8;not null" json:"seller_paid"`
	KurirPaid           int64          `gorm:"column:kurir_paid;type:int8;not null" json:"kurir_paid"`
	EkspedisiPaid       int64          `gorm:"column:ekspedisi_paid;type:int8;not null" json:"ekspedisi_paid"`
	Total               int64          `gorm:"column:total;type:int8;not null" json:"total"`
	Reviewed            bool           `gorm:"column:reviewed;type:bool;not null;default:false" json:"reviewed"`
	CreatedAt           time.Time      `gorm:"autoCreateTime" db:"created_at" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime" db:"updated_at" json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" db:"deleted_at" json:"deleted_at,omitempty"`
}

func (Transaksi) TableName() string {
	return "transaksi"
}

type TransaksiFailed struct {
	ID                  int64          `gorm:"primaryKey;autoIncrement" json:"id_transaksi"`
	IdPengguna          int64          `gorm:"column:id_pengguna;not null" json:"id_pengguna"`
	IdSeller            int32          `gorm:"column:id_seller;not null" json:"id_seller"`
	IdBarangInduk       int32          `gorm:"column:id_barang_induk;not null" json:"id_barang_induk"`
	IdKategoriBarang    int64          `gorm:"column:id_kategori_barang;not null" json:"id_kategori_barang"`
	IdAlamatPengguna    int64          `gorm:"column:id_alamat_pengguna;not null" json:"id_alamat_pengguna"`
	IdAlamatGudang      int64          `gorm:"column:id_alamat_gudang;type:int8;not null" json:"id_alamat_gudang"`
	IdAlamatEkspedisi   int64          `gorm:"column:id_alamat_ekspedisi;type:int8;not null" json:"id_alamat_ekspedisi"`
	IdPembayaran        int64          `gorm:"column:id_pembayaran;not null" json:"id_pembayaran"`
	KendaraanPengiriman string         `gorm:"column:kendaraan_pengiriman;type:jenis_kendaraan_kurir;default:'Motor';not null" json:"kendaraan_pengiriman"`
	JenisPengiriman     string         `gorm:"column:jenis_pengiriman;type:jenis_layanan_kurir;not null" json:"jenis_pengiriman"`
	JarakTempuh         string         `gorm:"column:jarak_tempuh;not null" json:"jarak_tempuh"`
	BeratTotalKg        int16          `gorm:"column:berat_total_kg;type:int2;not null" json:"berat_total_kg_pengiriman"`
	KodeOrderSistem     string         `gorm:"column:kode_order_sistem;type:varchar(100);not null" json:"kode_order_sistem"`
	KodeResiEkspedisi   *string        `gorm:"column:kode_resi_ekspedisi;type:varchar(100)" json:"kode_resi_ekspedisi"`
	Status              string         `gorm:"column:status;type:status_transaksi;default:'Dibayar';not null" json:"status"`
	DibatalkanOleh      *string        `gorm:"column:dibatalkan_oleh;type:jenis_entity" json:"dibatalkan_oleh"`
	Catatan             string         `gorm:"column:catatan;type:text" json:"catatan"`
	KuantitasBarang     int32          `gorm:"column:kuantitas_barang;type:int4;not null" json:"kuantitas"`
	IsEkspedisi         bool           `gorm:"column:is_ekspedisi;not null;default:false" json:"is_ekspedisi"`
	SellerPaid          int64          `gorm:"column:seller_paid;type:int8;not null" json:"seller_paid"`
	KurirPaid           int64          `gorm:"column:kurir_paid;type:int8;not null" json:"kurir_paid"`
	EkspedisiPaid       int64          `gorm:"column:ekspedisi_paid;type:int8;not null" json:"ekspedisi_paid"`
	Total               int64          `gorm:"column:total;type:int8;not null" json:"total_transaksi"`
	Reviewed            bool           `gorm:"column:reviewed;type:bool;not null;default:false" json:"reviewed"`
	CreatedAt           time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (TransaksiFailed) TableName() string {
	return "transaksi_failed"
}
