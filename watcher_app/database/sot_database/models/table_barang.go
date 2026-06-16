package sot_models

import (
	"time"

	"gorm.io/gorm"
)

type BarangContract interface {
	Validating() string
}

type BarangInduk struct {
	ID               int32          `gorm:"primaryKey;autoIncrement" json:"id_barang_induk"`
	SellerID         int32          `gorm:"index;column:id_seller;not null" json:"id_seller"`
	Seller           Seller         `gorm:"foreignKey:SellerID;references:ID" json:"-"`
	IdDiskon         int64          `gorm:"column:id_diskon"`
	NamaBarang       string         `gorm:"column:nama_barang;type:varchar(200);not null" json:"nama"`
	JenisBarang      string         `gorm:"index;column:jenis_barang;type:seller_dedication;not null;default:'Semua Barang'" json:"jenis,omitempty"`
	Deskripsi        string         `gorm:"column:deskripsi;type:text" json:"deskripsi,omitempty"`
	OriginalKategori int64          `gorm:"column:original_kategori;type:int8" json:"original_kategori,omitempty"`
	HargaKategoris   int32          `gorm:"column:harga_kategori_barang" json:"harga_kategori_barang"`
	CreatedAt        time.Time      `gorm:"autoCreateTime"  json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"  json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" db:"deleted_at" json:"deleted_at,omitempty"`
}

func (BarangInduk) TableName() string {
	return "barang_induk"
}

type KategoriBarang struct {
	ID             int64          `gorm:"primaryKey;autoIncrement" json:"id_kategori_barang"`
	SellerID       int32          `gorm:"column:id_seller;not null" json:"id_seller_kategori_barang"`
	Seller         Seller         `gorm:"foreignKey:SellerID;references:ID" json:"-"`
	IdBarangInduk  int32          `gorm:"column:id_barang_induk;not null" json:"id_barang_induk_kategori"`
	BarangInduk    BarangInduk    `gorm:"foreignKey:IdBarangInduk;references:ID;constraint:OnDelete:CASCADE;" json:"-"`
	IDAlamat       int64          `gorm:"column:id_alamat_gudang;type:int8" json:"id_alamat_gudang_kategori_barang"`
	AlamatGudang   AlamatGudang   `gorm:"foreignKey:IDAlamat;references:ID" json:"-"`
	IDRekening     int64          `gorm:"column:id_rekening;type:int8" json:"id_rekening_kategori_barang"`
	RekeningSeller RekeningSeller `gorm:"foreignKey:IDRekening;references:ID" json:"-"`
	Nama           string         `gorm:"column:nama;type:varchar(120);not null" json:"nama_kategori_barang"`
	Deskripsi      string         `gorm:"column:deskripsi;type:text" json:"deskripsi_kategori_barang"`
	Warna          string         `gorm:"column:warna;type:varchar(50)" json:"warna_kategori_barang"`
	Stok           int32          `gorm:"column:stok;type:int4;not null" json:"stok_kategori_barang"`
	Harga          int32          `gorm:"column:harga;type:int4;not null" json:"harga_kategori_barang"`
	PotonganDiskon int32          `gorm:"column:diskon;type:int4;default:0" json:"-"`
	BeratGram      int16          `gorm:"column:berat_gram;type:int2" json:"berat_gram_kategori_barang"`
	DimensiPanjang int16          `gorm:"column:dimensi_panjang_cm;type:int2" json:"dimensi_panjang_cm_kategori_barang"`
	DimensiLebar   int16          `gorm:"column:dimensi_lebar_cm;type:int2" json:"dimensi_tinggi_cm_kategori_barang"`
	Sku            string         `gorm:"column:sku;type:varchar(20);not null" json:"sku_kategori"`
	IsOriginal     bool           `gorm:"column:is_original;type:bool;not null;default:false" json:"is_original_kategori_barang"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (KategoriBarang) TableName() string {
	return "kategori_barang"
}

type VarianBarang struct {
	ID            int64          `gorm:"primaryKey;autoIncrement" json:"id_varian_barang"`
	IdBarangInduk int32          `gorm:"column:id_barang_induk;not null" json:"id_barang_induk_varian_barang"`
	BarangInduk   BarangInduk    `gorm:"foreignKey:IdBarangInduk;references:ID" json:"-"`
	IdKategori    int64          `gorm:"column:id_kategori;not null" json:"id_kategori_varian_barang"`
	Kategori      KategoriBarang `gorm:"foreignKey:IdKategori;references:ID;constraint:OnDelete:CASCADE;" json:"-"`
	IdTransaksi   int64          `gorm:"column:id_transaksi;type:int8" json:"id_transksi_varian_barang,omitempty"`
	Sku           string         `gorm:"column:sku;type:varchar(100);not null" json:"Sku_varian_barang,omitempty"`
	Status        string         `gorm:"column:status;type:status_varian;not null;default:'Ready'" json:"status_varian_barang,omitempty"`
	HoldBy        int64          `gorm:"column:hold_by;type:int8;default:0" json:"hold_by_varian_barang"`
	HolderEntity  string         `gorm:"column:holder_entity;type:varchar(30)" json:"holder_entity_varian_barang"`
}

func (VarianBarang) TableName() string {
	return "varian_barang"
}
