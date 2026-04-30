package models

import (
	"time"

	"gorm.io/gorm"
)

type EntitySocialMedia struct {
	ID         int64      `gorm:"primaryKey;autoIncrement" json:"id_social_media"`
	EntityId   int64      `gorm:"column:entity_id;type:int8;not null" json:"entity_id_social_media"`
	Whatsapp   string     `gorm:"column:whatsapp;type:varchar(20)" json:"whatsapp_social_media"`
	Facebook   string     `gorm:"column:facebook;type:text" json:"facebook_social_media"`
	TikTok     string     `gorm:"column:tiktok;type:text" json:"tiktok_social_media"`
	Instagram  string     `gorm:"column:instagram;type:text" json:"instagram_social_media"`
	Metadata   []byte     `gorm:"column:metadata;type:bytea" json:"metadata_social_media"`
	EntityType string     `gorm:"column:entity_type;type:varchar(20);not null" json:"entity_type_social_media"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (EntitySocialMedia) TableName() string {
	return "entity_social_media"
}

type Komentar struct {
	ID            int64       `gorm:"primaryKey;autoIncrement" json:"id_komentar"`
	IdBarangInduk int32       `gorm:"column:id_barang_induk;not null" json:"id_barang_induk"`
	Baranginduk   BarangInduk `gorm:"foreignKey:IdBarangInduk;references:ID" json:"-"`
	IdEntity      int64       `gorm:"column:id_entity;not null" json:"id_entity"`
	JenisEntity   string      `gorm:"column:jenis_entity;type:varchar(50);not null" json:"jenis_entity"`
	Komentar      string      `gorm:"column:komentar;type:text;not null" json:"isi_komentar"`
	IsSeller      bool        `gorm:"column:is_seller;not null;default:false" json:"is_seller"`
	Dibalas       int64       `gorm:"column:dibalas;type:int8;default:0;" json:"dibalas_komentar"`
	CreatedAt     time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     *time.Time  `gorm:"index" json:"deleted_at,omitempty"`

	// Relasi one-to-many
	KomentarChild []KomentarChild `gorm:"foreignKey:IdKomentar;constraint:OnDelete:CASCADE;" json:"komentar_child,omitempty"`
}

func (Komentar) TableName() string {
	return "komentar"
}

type KomentarChild struct {
	ID          int64      `gorm:"primaryKey;autoIncrement" json:"id_komentar_child"`
	IdKomentar  int64      `gorm:"column:id_komentar;not null" json:"id_komentar"`
	Komentar    Komentar   `gorm:"foreignKey:IdKomentar;references:ID;constraint:OnDelete:CASCADE;" json:"-"`
	IdEntity    int64      `gorm:"column:id_entity;not null" json:"id_entity"`
	JenisEntity string     `gorm:"column:jenis_entity;type:varchar(50);not null" json:"jenis_entity"`
	IsiKomentar string     `gorm:"column:komentar;type:text;not null" json:"isi_komentar"`
	IsSeller    bool       `gorm:"column:is_seller;not null;default:false" json:"is_seller_komentar"`
	Mention     string     `gorm:"column:mention;type:text;not null" json:"mention_komentar"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (KomentarChild) TableName() string {
	return "komentar_child"
}

// ///////////////////////////////////////////////////////////////////////////////////////////
// ENGAGEMENT PENGGUNA
// ///////////////////////////////////////////////////////////////////////////////////////////

type Keranjang struct {
	ID             int64          `gorm:"primaryKey;autoIncrement" json:"id_keranjang"`
	IdPengguna     int64          `gorm:"column:id_pengguna;not null" json:"id_pengguna_keranjang"`
	Pengguna       Pengguna       `gorm:"foreignKey:IdPengguna;references:ID"`
	IdSeller       int32          `gorm:"column:id_seller;not null" json:"id_seller_barang_induk_keranjang"`
	Seller         Seller         `gorm:"foreignKey:IdSeller;references:ID"`
	IdBarangInduk  int32          `gorm:"column:id_barang_induk;not null" json:"id_barang_induk_keranjang"`
	BarangInduk    BarangInduk    `gorm:"foreignKey:IdBarangInduk;references:ID"`
	IdKategori     int64          `gorm:"id_kategori_barang;not null" json:"id_kategori_barang_keranjang"`
	Kategoribarang KategoriBarang `gorm:"foreignKey:IdKategori;references:ID"`
	Jumlah         int16          `gorm:"column:jumlah;type:int2;not null" json:"jumlah_keranjang"`
	Status         string         `gorm:"column:status;type:status_keranjang;not null" json:"status_keranjang"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      *time.Time     `gorm:"index" json:"deleted_at,omitempty"`
}

func (Keranjang) TableName() string {
	return "keranjang"
}

type BarangDisukai struct {
	ID            int64       `gorm:"primaryKey;autoIncrement" json:"id_barang_disukai"`
	IdPengguna    int64       `gorm:"column:id_pengguna;not null" json:"id_pengguna_barang_disukai"`
	Pengguna      Pengguna    `gorm:"foreignKey:IdPengguna;references:ID"`
	IdBarangInduk int32       `gorm:"column:id_barang_induk;not null" json:"id_barang_induk_disukai"`
	BarangInduk   BarangInduk `gorm:"foreignKey:IdBarangInduk;references:ID"`
	CreatedAt     time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     *time.Time  `gorm:"index" json:"deleted_at,omitempty"`
}

func (BarangDisukai) TableName() string {
	return "barang_disukai"
}

type BarangWishlist struct {
	ID            int64       `gorm:"primaryKey;autoIncrement" json:"id_barang_wishlist"`
	IdPengguna    int64       `gorm:"column:id_pengguna;not null" json:"id_pengguna_barang_wishlist"`
	Pengguna      Pengguna    `gorm:"foreignKey:IdPengguna;references:ID"`
	IdBarangInduk int32       `gorm:"column:id_barang_induk;not null" json:"id_barang_induk_barang_wishlist"`
	BarangInduk   BarangInduk `gorm:"foreignKey:IdBarangInduk;references:ID"`
	CreatedAt     time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     *time.Time  `gorm:"index" json:"deleted_at,omitempty"`
}

type AlamatPengguna struct {
	ID              int64          `gorm:"primaryKey;autoIncrement" json:"id_alamat_user"`
	IDPengguna      int64          `gorm:"column:id_pengguna;not null" json:"id_pengguna_alamat_user"`
	Pengguna        Pengguna       `gorm:"foreignKey:IDPengguna;references:ID" json:"-"`
	PanggilanAlamat string         `gorm:"column:panggilan_alamat;type:varchar(250);not null" json:"panggilan_alamat_user"`
	NomorTelephone  string         `gorm:"column:nomor_telefon;type:varchar(20);not null" json:"nomor_telfon_alamat_user"`
	NamaAlamat      string         `gorm:"column:nama_alamat;type:text;not null" json:"nama_alamat_user"`
	Provinsi        string         `gorm:"column:provinsi;type:nama_provinsi;not null" json:"provinsi_alamat_user"`
	Kota            string         `gorm:"column:kota;type:nama_kota;not null" json:"kota_alamat_user"`
	KodePos         string         `gorm:"column:kode_pos;type:varchar(40);not null" json:"kode_pos_alamat_user"`
	KodeNegara      string         `gorm:"column:kode_negara;default:'IDN';not null" json:"kode_negara_alamat_user"`
	Deskripsi       string         `gorm:"column:deskripsi;type:text;" json:"deskripsi_alamat_user"`
	Longitude       float64        `gorm:"column:longitude;type:decimal(11,8);" json:"longitude_alamat_user"`
	Latitude        float64        `gorm:"column:latitude;type:decimal(11,8);" json:"latitude_alamat_user"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (AlamatPengguna) TableName() string {
	return "alamat_pengguna"
}

type Wishlist struct {
	ID            int64       `gorm:"primaryKey;autoIncrement" json:"id_wishlist"`
	IdPengguna    int64       `gorm:"column:id_pengguna;not null" json:"id_pengguna_wishlist"`
	Pengguna      Pengguna    `gorm:"foreignKey:IdPengguna;references:ID" json:"-"`
	IdBarangInduk int32       `gorm:"column:id_barang_induk;not null" json:"id_barang_induk_wishlist"`
	BarangInduk   BarangInduk `gorm:"foreignKey:IdBarangInduk;references:ID" json:"-"`
	CreatedAt     time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     *time.Time  `gorm:"index" json:"deleted_at,omitempty"`
}

func (Wishlist) TableName() string {
	return "wishlist"
}

type Review struct {
	ID            int64       `gorm:"primaryKey;autoIncrement" json:"id_review"`
	IdPengguna    int64       `gorm:"column:id_pengguna;not null" json:"id_pengguna_review"`
	Pengguna      Pengguna    `gorm:"foreignKey:IdPengguna;references:ID" json:"-"`
	IdBarangInduk int32       `gorm:"column:id_barang_induk;not null" json:"id_barang_induk_review"`
	BarangInduk   BarangInduk `gorm:"foreignKey:IdBarangInduk;references:ID" json:"-"`
	Rating        float32     `gorm:"column:rating;type:numeric(2,1);not null" json:"rating_review"`
	Ulasan        string      `gorm:"column:ulasan;type:text" json:"ulasan_review"`
	CreatedAt     time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     *time.Time  `gorm:"index" json:"deleted_at,omitempty"`
}

func (Review) TableName() string {
	return "review"
}

type ReviewLike struct {
	ID         int64      `gorm:"primaryKey;autoIncrement" json:"id_like_review"`
	IdPengguna int64      `gorm:"column:id_pengguna;not null" json:"id_pengguna_like_wishlist"`
	Pengguna   Pengguna   `gorm:"foreignKey:IdPengguna;references:ID" json:"-"`
	IdReview   int64      `gorm:"column:id_review;not null" json:"id_review_like_wishlist"`
	Review     Review     `gorm:"foreignKey:IdReview;references:ID" json:"-"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (ReviewLike) TableName() string {
	return "review_like"
}

type ReviewDislike struct {
	ID         int64      `gorm:"primaryKey;autoIncrement" json:"id_dislike_review"`
	IdPengguna int64      `gorm:"column:id_pengguna;not null" json:"id_pengguna_dislike_wishlist"`
	Pengguna   Pengguna   `gorm:"foreignKey:IdPengguna;references:ID" json:"-"`
	IdReview   int64      `gorm:"column:id_review;not null" json:"id_review_dislike_wishlist"`
	Review     Review     `gorm:"foreignKey:IdReview;references:ID" json:"-"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (ReviewDislike) TableName() string {
	return "review_dislike"
}

// ///////////////////////////////////////////////////////////////////////////////////////////
// ENGAGEMENT SELLER
// ///////////////////////////////////////////////////////////////////////////////////////////

type Jenis_Seller struct {
	ID               int64      `gorm:"primaryKey;autoIncrement" json:"id_jenis_seller"`
	IdSeller         int32      `gorm:"column:id_seller;not null" json:"id_seller_jenis_seller"`
	Seller           Seller     `gorm:"foreignKey:IdSeller;references:ID" json:"-"`
	ValidationStatus string     `gorm:"column:validation_status; not null; default:'Pending'" json:"validation_status_jenis_seller"`
	Alasan           string     `gorm:"alasan_seller;type:text" json:"alasan_seller_jenis_seller"`
	AlasanAdmin      string     `gorm:"alasan_admin;type:text" json:"alasan_admin_jenis_seller"`
	TargetJenis      string     `gorm:"column:target_jenis;type:jenis_seller" json:"target_jenis_seller"`
	CreatedAt        time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (Jenis_Seller) TableName() string {
	return "jenis_seller_validation"
}

type BatalTransaksi struct {
	ID             int64      `gorm:"primaryKey;autoIncrement" json:"id_batal_transaksi"`
	IdTransaksi    int64      `gorm:"column:id_transaksi;not null" json:"id_transaksi_batal_transaksi"`
	ITransaksi     Transaksi  `gorm:"foreignKey:IdTransaksi;references:ID" json:"-"`
	DibatalkanOleh string     `gorm:"column:dibatalkan_oleh;type:varchar(20);not null" json:"transaksi_dibatalkan_oleh"`
	Alasan         string     `gorm:"column:alasan;type:text;not null" json:"alasan_batal_transaksi"`
	CreatedAt      time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (BatalTransaksi) TableName() string {
	return "batal_transaksi"
}

type Follower struct {
	ID         int64      `gorm:"primaryKey;autoIncrement" json:"id_follow"`
	IdFollower int64      `gorm:"column:id_follower;not null" json:"id_follower"`
	Pengguna   Pengguna   `gorm:"foreignKey:IdFollower;references:ID"` // user yang follow
	IdFollowed int64      `gorm:"column:id_followed;not null" json:"id_followed"`
	Seller     Seller     `gorm:"foreignKey:IdFollowed;references:ID"` // seller yang di-follow
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (Follower) TableName() string {
	return "follower"
}

type AlamatGudang struct {
	ID              int64          `gorm:"primaryKey;autoIncrement" json:"id_alamat_gudang"`
	IDSeller        int32          `gorm:"column:id_seller;not null" json:"id_seller_alamat_gudang"`
	Seller          Seller         `gorm:"foreignKey:IDSeller;references:ID" json:"-"`
	PanggilanAlamat string         `gorm:"column:panggilan_alamat;type:varchar(250);not null" json:"panggilan_alamat_gudang"`
	NomorTelephone  string         `gorm:"column:nomor_telefon;type:varchar(20);not null" json:"nomor_telfon_alamat_gudang"`
	NamaAlamat      string         `gorm:"column:nama_alamat;type:text;not null" json:"nama_alamat_gudang"`
	Provinsi        string         `gorm:"column:provinsi;type:nama_provinsi;not null" json:"provinsi_alamat_gudang"`
	Kota            string         `gorm:"column:kota;type:nama_kota;not null" json:"kota_alamat_gudang"`
	KodePos         string         `gorm:"column:kode_pos;type:varchar(40);not null" json:"kode_pos_alamat_gudang"`
	KodeNegara      string         `gorm:"column:kode_negara;default:'IDN';not null" json:"kode_negara_alamat_gudang"`
	Deskripsi       string         `gorm:"column:deskripsi;type:text;" json:"deskripsi_alamat_gudang"`
	Longitude       float64        `gorm:"column:longitude;type:decimal(11,8);" json:"longitude_alamat_gudang"`
	Latitude        float64        `gorm:"column:latitude;type:decimal(11,8);" json:"latitude_alamat_gudang"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (AlamatGudang) TableName() string {
	return "alamat_gudang"
}

type DistributorData struct {
	ID                        int64  `gorm:"primaryKey;autoIncrement" json:"id_distributor_data"`
	SellerId                  int32  `gorm:"column:seller_id;not null" json:"id_seller_distributor_data"`
	Seller                    Seller `gorm:"foreignKey:SellerId;references:ID" json:"-"`
	NamaPerusahaan            string `gorm:"column:nama_perusahaan;type:text;not null" json:"nama_perusahaan_distributor_data"`
	NIB                       string `gorm:"column:nib;type:varchar(20);not null" json:"nib_distributor_data"`
	NPWP                      string `gorm:"column:npwp;type:varchar(30);not null" json:"npwp_distributor_data"`
	DokumenIzinDistributorUrl string `gorm:"column:dokumen_izin_distributor;type:text;not null" json:"dokumen_izin_distributor_data"`
	Alasan                    string `gorm:"column:alasan;type:text" json:"alasan_distributor_data"`
	Status                    string `gorm:"column:status;type:status_jenis_seller;not null" json:"status_distributor_data"`
}

func (DistributorData) TableName() string {
	return "distributor_data"
}

type BrandData struct {
	ID                    int64  `gorm:"primaryKey;autoIncrement" json:"id_brand_data"`
	SellerId              int32  `gorm:"column:seller_id;not null" json:"id_seller_brand_data"`
	Seller                Seller `gorm:"foreignKey:SellerId;references:ID" json:"-"`
	NamaPerusahaan        string `gorm:"column:nama_perusahaan;type:text;not null" json:"nama_perusahaan_brand_data"`
	NegaraAsal            string `gorm:"column:negara_asal;type:varchar(57);not null;default:'indonesia'" json:"negara_asal_brand_data"`
	LembagaPendaftaran    string `gorm:"column:lembaga_pendaftaran;type:text;not null" json:"lembaga_pendaftaran_brand_data"`
	NomorPendaftaranMerek string `gorm:"column:nomor_pendaftaran_merek;type:text;not null" json:"nomor_pendaftaran_merek_brand_data"`
	SertifikatMerekUrl    string `gorm:"column:sertifikat_merek_url;type:text;not null" json:"sertifikat_merek_url_brand_data"`
	DokumenPerwakilanUrl  string `gorm:"column:dokumen_perwakilan_url;type:text;not null" json:"dokumen_perwakilan_url_brand_data"`
	NIB                   string `gorm:"column:nib;type:varchar(20);not null" json:"nib_brand_data"`
	NPWP                  string `gorm:"column:npwp;type:varchar(30);not null" json:"npwp_brand_data"`
	Alasan                string `gorm:"column:alasan;type:text" json:"alasan_brand_data"`
	Status                string `gorm:"column:status;type:status_jenis_seller;not null" json:"status_brand_data"`
}

func (BrandData) TableName() string {
	return "brand_data"
}

type Etalase struct {
	ID           int64  `gorm:"primaryKey;autoIncrement" json:"id_etalase"`
	SellerID     int64  `gorm:"column:id_seller;not null" json:"id_seller_etalase"`
	Seller       Seller `gorm:"foreignKey:SellerID;references:ID" json:"-"`
	Nama         string `gorm:"column:nama;type:varchar(100);not null" json:"nama_etalase"`
	Deskripsi    string `gorm:"column:deskripsi;type:text" json:"deskripsi_etalase"`
	JumlahBarang int32  `gorm:"column:jumlah_barang;not null;default:0" json:"jumlah_barang"`
}

func (Etalase) TableName() string {
	return "etalase"
}

type BarangKeEtalase struct {
	ID            int64       `gorm:"primaryKey;autoIncrement" json:"id_barang_ke_etalase"`
	IdEtalase     int64       `gorm:"column:id_etalase;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"id_etalase_barang_ke_etalase"`
	Etalase       Etalase     `gorm:"foreignKey:IdEtalase;references:ID" json:"-"`
	IdBarangInduk int64       `gorm:"column:id_barang_induk;not null" json:"id_barang_induk_barang_ke_etalase"`
	BarangInduk   BarangInduk `gorm:"foreignKey:IdBarangInduk;references:ID" json:"-"`
}

func (BarangKeEtalase) TableName() string {
	return "barang_ke_etalase"
}

type DiskonProduk struct {
	ID            int64          `gorm:"primaryKey;autoIncrement" json:"id_diskon_produk"`
	SellerId      int32          `gorm:"column:id_seller;not null" json:"id_seller_diskon_produk"`
	Seller        Seller         `gorm:"foreignKey:SellerId;references:ID" json:"-"`
	Nama          string         `gorm:"column:nama;type:varchar(100);not null" json:"nama_diskon_produk"`
	Deskripsi     string         `gorm:"column:deskripsi;type:text" json:"deskripsi_diskon_produk"`
	DiskonPersen  float64        `gorm:"column:diskon_persen;type:decimal(5,2);not null" json:"diskon_persen_diskon_produk"`
	BerlakuMulai  time.Time      `gorm:"column:berlaku_mulai;type:date;not null" json:"berlaku_mulai_diskon_produk"`
	BerlakuSampai time.Time      `gorm:"column:berlaku_sampai;type:date;not null" json:"berlaku_sampai_diskon_produk"`
	Status        string         `gorm:"column:status;type:varchar(20);default:'Draft';not null" json:"status_diskon_produk"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (DiskonProduk) TableName() string {
	return "diskon_produk"
}

type BarangDiDiskon struct {
	ID               int64          `gorm:"primaryKey;autoIncrement" json:"id_barang_di_diskon"`
	SellerId         int32          `gorm:"column:id_seller;not null" json:"id_seller_barang_di_diskon"`
	Seller           Seller         `gorm:"foreignKey:SellerId;references:ID" json:"-"`
	IdDiskon         int64          `gorm:"column:id_diskon;not null;constraint:OnDelete:CASCADE" json:"id_diskon_barang_di_diskon"`
	DiskonProduk     DiskonProduk   `gorm:"foreignKey:IdDiskon;references:ID" json:"-"`
	IdBarangInduk    int32          `gorm:"column:id_barang_induk;not null" json:"id_barang_induk_barang_di_diskon"`
	BarangInduk      BarangInduk    `gorm:"foreignKey:IdBarangInduk;references:ID" json:"-"`
	IdKategoriBarang int64          `gorm:"column:id_kategori_barang;not null" json:"id_kategori_barang_barang_di_diskon"`
	KategoriBarang   KategoriBarang `gorm:"foreignKey:IdKategoriBarang;references:ID" json:"-"`
	Status           string         `gorm:"column:status;type:status_barang_di_diskon;default:'Waiting';not null" json:"status_barang_di_diskon"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (BarangDiDiskon) TableName() string {
	return "barang_di_diskon"
}

// ///////////////////////////////////////////////////////////////////////////////////////////
// ENGAGEMENT KURIR
// ///////////////////////////////////////////////////////////////////////////////////////////

type InformasiKurir struct {
	ID           int64      `gorm:"primaryKey;autoIncrement" json:"id_informasi_kurir"`
	IDkurir      int64      `gorm:"column:id_kurir;not null" json:"id_kurir_informasi_kurir"`
	Kurir        Kurir      `gorm:"foreignKey:IDkurir;references:ID" json:"-"`
	TanggalLahir string     `gorm:"column:tanggal_lahir;type:text" json:"tanggal_lahir_informasi_kurir" `
	Alasan       string     `gorm:"column:alasan;type:text" json:"alasan_informasi_kurir"`
	Ktp          bool       `gorm:"column:informasi_ktp;type:boolean;not null;default:false" json:"informasi_ktp_informasi_kurir"`
	InformasiSim bool       `gorm:"column:informasi_sim;type:boolean;not null;default:false" json:"informasi_sim_informasi_kurir"`
	Status       string     `gorm:"column:status;type:status_perizinan;not null; default:'Pending'" json:"status_informasi_kurir"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
	DeletedAt    *time.Time `gorm:"index"`
}

func (InformasiKurir) TableName() string {
	return "informasi_kurir"
}

type InformasiKendaraanKurir struct {
	ID             int64      `gorm:"primaryKey;autoIncrement" json:"id_informasi_kendaraan_kurir"`
	IDkurir        int64      `gorm:"column:id_kurir;not null" json:"id_kurir_kendaraan_kurir"`
	Kurir          Kurir      `gorm:"foreignKey:IDkurir;references:ID" json:"-"`
	JenisKendaraan string     `gorm:"column:jenis_kendaraan;type:jenis_kendaraan_kurir;not null; default:'Motor'" json:"jenis_kendaraan_informasi_kendaraan_kurir"`
	NamaKendaraan  string     `gorm:"column:nama_kendaraan;type:text;not null" json:"nama_kendaraan_informasi_kendaraan_kurir"`
	RodaKendaraan  string     `gorm:"column:roda_kendaraan;type:roda_kendaraan_kurir;not null" json:"roda_kendaraan_informasi_kendaraan_kurir"`
	STNK           bool       `gorm:"column:informasi_stnk;type:boolean;not null; default:false" json:"informasi_stnk_informasi_kendaraan_kurir"`
	BPKB           bool       `gorm:"column:informasi_bpkb;type:boolean;not null; default:false" json:"informasi_bpkb_informasi_kendaraan_kurir"`
	NoRangka       string     `gorm:"column:nomor_rangka;type:text" json:"nomor_rangka_informasi_kendaraan_kurir"`
	NoMesin        string     `gorm:"column:nomor_mesin;type:text;" json:"nomor_mesin_informasi_kendaraan_kurir"`
	Status         string     `gorm:"column:status;type:status_perizinan;not null;default:'Pending'" json:"status_informasi_kendaraan_kurir"`
	CreatedAt      time.Time  `gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime"`
	DeletedAt      *time.Time `gorm:"index"`
}

func (InformasiKendaraanKurir) TableName() string {
	return "informasi_kendaraan_kurir"
}

type AlamatKurir struct {
	ID              int64      `gorm:"primaryKey;autoIncrement" json:"id_alamat_kurir"`
	IdKurir         int64      `gorm:"column:id_kurir;not null" json:"id_kurir_alamat_kurir"`
	Kurir           Kurir      `gorm:"foreignKey:IdKurir;references:ID" json:"-"`
	PanggilanAlamat string     `gorm:"column:panggilan_alamat;type:text;not null" json:"panggilan_alamat_kurir"`
	NomorTelephone  string     `gorm:"column:nomor_telefon;type:varchar(25);not null" json:"nomor_telefon_alamat_kurir"`
	NamaAlamat      string     `gorm:"column:nama_alamat;type:text" json:"nama_alamat_kurir"`
	Provinsi        string     `gorm:"column:provinsi;type:nama_provinsi;not null" json:"provinsi_alamat_kurir"`
	Kota            string     `gorm:"column:kota;type:nama_kota;not null" json:"kota_alamat_kurir"`
	KodeNegara      string     `gorm:"column:kode_negara;type:varchar(10)" json:"kode_negara_alamat_kurir"`
	KodePos         string     `gorm:"column:kode_pos;type:varchar(10)" json:"kode_pos_alamat_kurir"`
	Deskripsi       string     `gorm:"column:deskripsi;type:text" json:"deskripsi_alamat_kurir"`
	Longitude       float64    `gorm:"column:longitude;type:decimal(11,8)" json:"longitude_alamat_kurir"`
	Latitude        float64    `gorm:"column:latitude;type:decimal(11,8)" json:"latitude_alamat_kurir"`
	CreatedAt       time.Time  `gorm:"autoCreateTime"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime"`
	DeletedAt       *time.Time `gorm:"index"`
}

func (AlamatKurir) TableName() string {
	return "alamat_kurir"
}

type BidKurirData struct {
	ID              int64          `gorm:"primaryKey;autoIncrement" json:"id_bid_kurir_data"`
	IdKurir         int64          `gorm:"column:id_kurir;not null" json:"id_kurir_bid_kurir_data"`
	Kurir           Kurir          `gorm:"foreignKey:IdKurir;references:ID" json:"-"`
	JenisPengiriman string         `gorm:"column:jenis_pengiriman;type:jenis_layanan_kurir;not null" json:"jenis_pengiriman_bid_kurir_data"`
	Mode            string         `gorm:"column:mode;type:mode_bid_kurir;not null" json:"mode_bid_kurir_data"`
	Provinsi        string         `gorm:"column:provinsi;type:nama_provinsi;not null" json:"provinsi_bid_kurir_data"`
	Kota            string         `gorm:"column:kota;type:nama_kota;not null" json:"kota_bid_kurir_data"`
	IsEkspedisi     bool           `gorm:"column:is_ekspedisi;type:boolean;not null;default:false" json:"is_ekspedisi_bid_kurir_data"`
	Alamat          string         `gorm:"column:alamat;type:text" json:"alamat_bid_kurir_data"`
	Longitude       float64        `gorm:"column:longitude;type:numeric(11,8);not null" json:"longitude_bid_kurir_data"`
	Latitude        float64        `gorm:"column:latitude;type:numeric(11,8);not null" json:"latitude_bid_kurir_data"`
	MaxKg           int16          `gorm:"column:max_kg;type:smallint;not null" json:"max_kg_bid_kurir_data"`
	SlotTersisa     int32          `gorm:"column:slot_tersisa;type:int;not null;default:0" json:"slot_tersisa_bid_kurir_data"`
	Dimulai         time.Time      `gorm:"column:dimulai;type:time;not null" json:"dimulai_bid_kurir_data"`
	Selesai         *time.Time     `gorm:"column:selesai;type:time" json:"selesai_bid_kurir_data"`
	JenisKendaraan  string         `gorm:"column:jenis_kendaraan;type:jenis_kendaraan_kurir;not null" json:"jenis_kendaraan_bid_kurir_data"`
	Status          string         `gorm:"column:status;type:status_bid_data;not null;default:'Mengumpulkan'" json:"status_bid_data"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (BidKurirData) TableName() string {
	return "bid_kurir_data"
}

type BidKurirNonEksScheduler struct {
	ID           int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	IdBid        int64          `gorm:"column:id_bid;not null" json:"id_bid"`
	BidKurirData BidKurirData   `gorm:"foreignKey:IdBid;references:ID" json:"-"`
	IdKurir      int64          `gorm:"column:id_kurir;not null" json:"id_kurir"`
	Kurir        Kurir          `gorm:"foreignKey:IdKurir;references:ID" json:"-"`
	Urutan       int8           `gorm:"column:urutan;type:int;not null"`
	IdPengiriman int64          `gorm:"column:id_pengiriman;not null" json:"id_pengiriman"`
	Pengiriman   Pengiriman     `gorm:"foreignKey:IdPengiriman;references:ID" json:"-"`
	Status       string         `gorm:"column:status;type:status_bid_scheduler;not null;default:Wait" json:"status"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (BidKurirNonEksScheduler) TableName() string {
	return "bid_kurir_non_eks_scheduler"
}

type BidKurirEksScheduler struct {
	ID                  int64               `gorm:"primaryKey;autoIncrement" json:"id"`
	IdBid               int64               `gorm:"column:id_bid;not null" json:"id_bid"`
	BidKurirData        BidKurirData        `gorm:"foreignKey:IdBid;references:ID" json:"-"`
	IdKurir             int64               `gorm:"column:id_kurir;not null" json:"id_kurir"`
	Kurir               Kurir               `gorm:"foreignKey:IdKurir;references:ID" json:"-"`
	Urutan              int8                `gorm:"column:urutan;type:int;not null"`
	IdPengirimanEks     int64               `gorm:"column:id_pengiriman_ekspedisi;not null" json:"id_pengiriman_ekspedisi"`
	PengirimanEkspedisi PengirimanEkspedisi `gorm:"foreignKey:IdPengirimanEks;references:ID" json:"-"`
	Status              string              `gorm:"column:status;type:status_bid_scheduler;not null;default:Wait" json:"status"`
	CreatedAt           time.Time           `gorm:"autoCreateTime"`
	UpdatedAt           time.Time           `gorm:"autoUpdateTime"`
	DeletedAt           gorm.DeletedAt      `gorm:"index"`
}

func (BidKurirEksScheduler) TableName() string {
	return "bid_kurir_eks_scheduler"
}
