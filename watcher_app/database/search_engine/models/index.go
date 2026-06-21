package se_models

import (
	"time"
)

type Seller struct {
	ID               int32      `json:"id"`
	Username         string     `json:"username"`
	Nama             string     `json:"nama"`
	Email            string     `json:"email"`
	Jenis            string     `json:"jenis"`
	SellerDedication string     `json:"seller_dedication"`
	JamOperasional   string     `json:"jam_operasional"`
	Punchline        string     `json:"punchline"`
	Password         string     `json:"password_hash"`
	Deskripsi        string     `json:"deskripsi"`
	StatusSeller     string     `json:"status"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

func (s Seller) IndexName() string {
	return "seller_se"
}

type BarangInduk struct {
	ID               int32      `json:"id"`
	SellerID         int32      `json:"id_seller"`
	IdDiskon         int64      `json:"id_diskon"`
	NamaBarang       string     `json:"nama_barang"`
	JenisBarang      string     `json:"jenis_barang,omitempty"`
	Deskripsi        string     `json:"deskripsi,omitempty"`
	OriginalKategori int64      `json:"original_kategori,omitempty"`
	HargaKategoris   int32      `json:"harga_kategori_barang"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

func (b BarangInduk) IndexName() string {
	return "barang_induk_se"
}

type Transaksi struct {
	ID                  int64       `json:"id"`
	IdPengguna          int64       `json:"id_pengguna"`
	IdSeller            int32       `json:"id_seller"`
	Seller              Seller      `json:"-"`
	IdBarangInduk       int64       `json:"id_barang_induk"`
	BarangInduk         BarangInduk `json:"-"`
	IdKategoriBarang    int64       `json:"id_kategori_barang"`
	IdAlamatPengguna    int64       `json:"id_alamat_pengguna"`
	IdAlamatGudang      int64       `json:"id_alamat_gudang"`
	IdAlamatEkspedisi   int64       `json:"id_alamat_ekspedisi"`
	IdPembayaran        int64       `json:"id_pembayaran"`
	KendaraanPengiriman string      `json:"kendaraan_pengiriman"`
	JenisPengiriman     string      `json:"jenis_pengiriman"`
	JarakTempuh         string      `json:"jarak_tempuh"`
	BeratTotalKg        int16       `json:"berat_total_kg"`
	KodeOrderSistem     string      `json:"kode_order_sistem"`
	KodeResiEkspedisi   *string     `json:"kode_resi_ekspedisi,omitempty"`
	Status              string      `json:"status"`
	DibatalkanOleh      *string     `json:"dibatalkan_oleh,omitempty"`
	Catatan             string      `json:"catatan,omitempty"`
	KuantitasBarang     int32       `json:"kuantitas_barang"`
	IsEkspedisi         bool        `json:"is_ekspedisi"`
	SellerPaid          int64       `json:"seller_paid"`
	KurirPaid           int64       `json:"kurir_paid"`
	EkspedisiPaid       int64       `json:"ekspedisi_paid"`
	Total               int64       `json:"total"`
	Reviewed            bool        `json:"reviewed"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at"`
	DeletedAt           *time.Time  `json:"deleted_at,omitempty"`
}

func (t Transaksi) IndexName() string {
	return "transaksi_se"
}

type AlamatPengguna struct {
	ID              int64      `json:"id"` // Tetap dipertahankan sesuai nama primary key/kolom id-nya
	IDPengguna      int64      `json:"id_pengguna"`
	PanggilanAlamat string     `json:"panggilan_alamat"`
	NomorTelephone  string     `json:"nomor_telefon"`
	NamaAlamat      string     `json:"nama_alamat"`
	Provinsi        string     `json:"provinsi"`
	Kota            string     `json:"kota"`
	KodePos         string     `json:"kode_pos"`
	KodeNegara      string     `json:"kode_negara"`
	Deskripsi       string     `json:"deskripsi,omitempty"`
	Longitude       float64    `json:"longitude"`
	Latitude        float64    `json:"latitude"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

func (t AlamatPengguna) IndexName() string {
	return "alamat_pengguna_se"
}

type AlamatKurir struct {
	ID              int64      `json:"id"`
	IdKurir         int64      `json:"id_kurir_alamat_kurir"`
	PanggilanAlamat string     `json:"panggilan_alamat_kurir"`
	NomorTelephone  string     `json:"nomor_telefon_alamat_kurir"`
	NamaAlamat      string     `json:"nama_alamat_kurir"`
	Provinsi        string     `json:"provinsi_alamat_kurir"`
	Kota            string     `json:"kota_alamat_kurir"`
	KodeNegara      string     `json:"kode_negara_alamat_kurir"`
	KodePos         string     `json:"kode_pos_alamat_kurir"`
	Deskripsi       string     `json:"deskripsi_alamat_kurir"`
	Longitude       float64    `json:"longitude_alamat_kurir"`
	Latitude        float64    `json:"latitude_alamat_kurir"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

func (a AlamatKurir) IndexName() string {
	return "alamat_kurir_se"
}

type AlamatGudang struct {
	ID              int64      `json:"id"`
	IDSeller        int32      `json:"id_seller_alamat_gudang"`
	Seller          Seller     `json:"-"`
	PanggilanAlamat string     `json:"panggilan_alamat_gudang"`
	NomorTelephone  string     `json:"nomor_telfon_alamat_gudang"`
	NamaAlamat      string     `json:"nama_alamat_gudang"`
	Provinsi        string     `json:"provinsi_alamat_gudang"`
	Kota            string     `json:"kota_alamat_gudang"`
	KodePos         string     `json:"kode_pos_alamat_gudang"`
	KodeNegara      string     `json:"kode_negara_alamat_gudang"`
	Deskripsi       string     `json:"deskripsi_alamat_gudang"`
	Longitude       float64    `json:"longitude_alamat_gudang"`
	Latitude        float64    `json:"latitude_alamat_gudang"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

func (a AlamatGudang) IndexName() string {
	return "alamat_gudang_se"
}

type Pengguna struct {
	ID             int64      `json:"id"`
	Username       string     `json:"username"`
	Nama           string     `json:"nama"`
	Email          string     `json:"email"`
	PasswordHash   string     `json:"password_hash"`
	PinHash        string     `json:"pin_hash"`
	StatusPengguna string     `json:"status"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
}

func (p Pengguna) IndexName() string {
	return "pengguna_se"
}

type Kurir struct {
	ID            int64      `json:"id_kurir"`
	Nama          string     `json:"nama"`
	Username      string     `json:"username"`
	Email         string     `json:"email"`
	Jenis         string     `json:"jenis"`
	PasswordHash  string     `json:"password_hash"`
	Deskripsi     string     `json:"deskripsi"`
	StatusKurir   string     `json:"status"`
	StatusBid     string     `json:"status_bid"`
	VerifiedKurir bool       `json:"verified"`
	Rating        float32    `json:"rating"`
	TipeKendaraan string     `json:"tipe_kendaraan"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}

func (k Kurir) IndexName() string {
	return "kurir_se"
}

type AlamatEkspedisi struct {
	ID              int64      `json:"id"`
	Kota            string     `json:"kota"`
	NamaAlamat      string     `json:"nama_alamat"`
	Lokasi          string     `json:"lokasi"`
	Longitude       float64    `json:"longitude"`
	Latitude        float64    `json:"latitude"`
	PengirimanCount int64      `json:"pengiriman_count"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

func (a AlamatEkspedisi) IndexName() string {
	return "alamat_ekspedisi_se"
}
