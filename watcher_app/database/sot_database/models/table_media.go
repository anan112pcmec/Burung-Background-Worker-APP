package models

import (
	"time"

	"gorm.io/gorm"
)

type MediaPenggunaProfilFoto struct {
	ID         int64          `gorm:"primaryKey;autoIncrement" json:"id_media_pengguna_profil_foto"`
	IdPengguna int64          `gorm:"column:id_pengguna;not null" json:"id_pengguna"`
	Pengguna   Pengguna       `gorm:"foreignKey:IdPengguna;references:ID" json:"-"`
	Key        string         `gorm:"column:key;type:text;not null" json:"key"`
	Format     string         `gorm:"column:format;type:varchar(20);not null" json:"format"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (MediaPenggunaProfilFoto) TableName() string {
	return "media_pengguna_profil_foto"
}

func (MediaPenggunaProfilFoto) PathName() string {
	return "/media-pengguna-profil-foto/"
}

type MediaSellerProfilFoto struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id_media_seller_profil_foto"`
	IdSeller  int64          `gorm:"column:id_seller;not null" json:"id_seller"`
	Seller    Seller         `gorm:"foreignKey:IdSeller;references:ID" json:"-"`
	Key       string         `gorm:"column:key;type:text;not null" json:"key"`
	Format    string         `gorm:"column:format;type:varchar(20);not null" json:"format"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (MediaSellerProfilFoto) PathName() string {
	return "/media_seller_profil_foto/"
}

func (MediaSellerProfilFoto) TableName() string {
	return "media_seller_foto_profil"
}

type MediaSellerBannerFoto struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id_media_seller_banner_foto"`
	IdSeller  int64          `gorm:"column:id_seller;not null" json:"id_seller"`
	Seller    Seller         `gorm:"foreignKey:IdSeller;references:ID" json:"-"`
	Key       string         `gorm:"column:key;type:text;not null" json:"key"`
	Format    string         `gorm:"column:format;type:varchar(20);not null" json:"format"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (MediaSellerBannerFoto) PathName() string {
	return "/media_seller_banner_foto/"
}

func (MediaSellerBannerFoto) TableName() string {
	return "media_seller_banner_foto"
}

type MediaSellerTokoFisikFoto struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id_media_seller_toko_fisik_foto"`
	IdSeller  int32          `gorm:"column:id_seller;not null" json:"id_seller"`
	Seller    Seller         `gorm:"foreignKey:IdSeller;references:ID" json:"-"`
	Key       string         `gorm:"column:key;type:text;not null" json:"key"`
	Format    string         `gorm:"column:format;type:varchar(20);not null" json:"format"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (MediaSellerTokoFisikFoto) PathName() string {
	return "/media_seller_toko_fisik_foto/"
}

func (MediaSellerTokoFisikFoto) TableName() string {
	return "media_seller_toko_fisik_foto"
}

type MediaKurirProfilFoto struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id_media_kurir_profil_foto"`
	IdKurir   int64          `gorm:"column:id_kurir;not null" json:"id_kurir"`
	Kurir     Kurir          `gorm:"foreignKey:IdKurir;references:ID" json:"-"`
	Key       string         `gorm:"column:key;type:text;not null" json:"key"`
	Format    string         `gorm:"column:format;type:varchar(20);not null" json:"format"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (MediaKurirProfilFoto) PathName() string {
	return "/media_kurir_profil_foto/"
}

func (MediaKurirProfilFoto) TableName() string {
	return "media_kurir_profil_foto"
}

type MediaEtalaseFoto struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id_media_etalase_foto"`
	IdEtalase int64          `gorm:"column:id_etalase;not null" json:"id_etalase"`
	Etalase   Etalase        `gorm:"foreignKey:IdEtalase;references:ID" json:"-"`
	Key       string         `gorm:"column:key;type:text;not null" json:"key"`
	Format    string         `gorm:"column:format;type:varchar(20);not null" json:"format"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (MediaEtalaseFoto) PathName() string {
	return "/media_etalase_foto/"
}

func (MediaEtalaseFoto) TableName() string {
	return "media_etalase_foto"
}

type MediaBarangIndukFoto struct {
	ID            int64          `gorm:"primaryKey;autoIncrement" json:"id_media_barang_induk_foto"`
	IdBarangInduk int64          `gorm:"column:id_barang_induk;not null" json:"id_barang_induk"`
	BarangInduk   BarangInduk    `gorm:"foreignKey:IdBarangInduk;references:ID" json:"-"`
	Key           string         `gorm:"column:key;type:text;not null" json:"key"`
	Format        string         `gorm:"column:format;type:varchar(20);not null" json:"format"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (MediaBarangIndukFoto) PathName() string {
	return "/media_barang_induk_foto/"
}

func (MediaBarangIndukFoto) TableName() string {
	return "media_barang_induk_foto"
}

type MediaBarangIndukVideo struct {
	ID            int64          `gorm:"primaryKey;autoIncrement" json:"id_media_barang_induk_video"`
	IdBarangInduk int64          `gorm:"column:id_barang_induk;not null" json:"id_barang_induk"`
	BarangInduk   BarangInduk    `gorm:"foreignKey:IdBarangInduk;references:ID" json:"-"`
	Key           string         `gorm:"column:key;type:text;not null" json:"key"`
	Format        string         `gorm:"column:format;type:varchar(20);not null" json:"format"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (MediaBarangIndukVideo) PathName() string {
	return "/media_barang_induk_video/"
}

func (MediaBarangIndukVideo) TableName() string {
	return "media_barang_induk_video"
}

type MediaKategoriBarangFoto struct {
	ID               int64          `gorm:"primaryKey;autoIncrement" json:"id_media_kategori_barang_foto"`
	IdKategoriBarang int64          `gorm:"column:id_kategori_barang;not null" json:"id_kategori_barang"`
	KategoriBarang   KategoriBarang `gorm:"foreignKey:IdKategoriBarang;references:ID" json:"-"`
	IdBarangInduk    int64          `gorm:"column:id_barang_induk;not null" json:"id_barang_induk"`
	BarangInduk      BarangInduk    `gorm:"foreignKey:IdBarangInduk;references:ID" json:"-"`
	Key              string         `gorm:"column:key;type:text;not null" json:"key"`
	Format           string         `gorm:"column:format;type:varchar(20);not null" json:"format"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (MediaKategoriBarangFoto) PathName() string {
	return "/media_kategori_barang_foto/"
}

func (MediaKategoriBarangFoto) TableName() string {
	return "media_kategori_barang_foto"
}

type MediaDistributorDataDokumen struct {
	ID                int64           `gorm:"primaryKey;autoIncrement" json:"id_media_ditributor_data_dokumen_foto"`
	IdDistributorData int64           `gorm:"column:id_distributor_data;not null" json:"id_distributor_data"`
	DistributorData   DistributorData `gorm:"foreignKey:IdDistributorData;references:ID" json:"-"`
	Key               string          `gorm:"column:key;type:text;not null" json:"key"`
	Format            string          `gorm:"column:format;type:varchar(20);not null" json:"format"`
	CreatedAt         time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}

func (MediaDistributorDataDokumen) PathName() string {
	return "/media_distributor_data_dokumen/"
}

func (MediaDistributorDataDokumen) TableName() string {
	return "media_distributor_data_dokumen"
}

type MediaDistributorDataNPWPFoto struct {
	ID                int64           `gorm:"primaryKey;autoIncrement" json:"id_media_ditributor_data_npwp_foto"`
	IdDistributorData int64           `gorm:"column:id_distributor_data;not null" json:"id_distributor_data"`
	DistributorData   DistributorData `gorm:"foreignKey:IdDistributorData;references:ID" json:"-"`
	Key               string          `gorm:"column:key;type:text;not null" json:"key"`
	Format            string          `gorm:"column:format;type:varchar(20);not null" json:"format"`
	CreatedAt         time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}

func (MediaDistributorDataNPWPFoto) PathName() string {
	return "/media_distributor_data_npwp_foto/"
}

func (MediaDistributorDataNPWPFoto) TableName() string {
	return "media_distributor_data_npwp_foto"
}

type MediaDistributorDataNIBFoto struct {
	ID                int64           `gorm:"primaryKey;autoIncrement" json:"id_media_ditributor_data_nib_foto"`
	IdDistributorData int64           `gorm:"column:id_distributor_data;not null" json:"id_distributor_data"`
	DistributorData   DistributorData `gorm:"foreignKey:IdDistributorData;references:ID" json:"-"`
	Key               string          `gorm:"column:key;type:text;not null" json:"key"`
	Format            string          `gorm:"column:format;type:varchar(20);not null" json:"format"`
	CreatedAt         time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}

func (MediaDistributorDataNIBFoto) PathName() string {
	return "/media_distributor_data_nib_foto/"
}

func (MediaDistributorDataNIBFoto) TableName() string {
	return "media_distributor_data_nib_foto"
}

type MediaDistributorDataSuratKerjasamaDokumen struct {
	ID                int64           `gorm:"primaryKey;autoIncrement" json:"id_media_ditributor_data_surat_kerjasama_dokumen"`
	IdDistributorData int64           `gorm:"column:id_distributor_data;not null" json:"id_distributor_data"`
	DistributorData   DistributorData `gorm:"foreignKey:IdDistributorData;references:ID" json:"-"`
	Key               string          `gorm:"column:key;type:text;not null" json:"key"`
	Format            string          `gorm:"column:format;type:varchar(20);not null" json:"format"`
	CreatedAt         time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}

func (MediaDistributorDataSuratKerjasamaDokumen) PathName() string {
	return "/media_distributor_data_surat_kerjasama_dokumen/"
}

func (MediaDistributorDataSuratKerjasamaDokumen) TableName() string {
	return "media_distributor_data_surat_kerjasama_dokumen"
}

type MediaBrandDataPerwakilanDokumen struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id_media_brand_data_perwakilan_dokumen"`
	IdBrandData int64     `gorm:"column:id_brand_data" json:"id_brand_data"`
	BrandData   BrandData `gorm:"foreignKey:IdBrandData;references:ID" json:"-"`
	Key         string    `gorm:"column:key;type:text;not null" json:"key"`
	Format      string    `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaBrandDataPerwakilanDokumen) PathName() string {
	return "/media_brand_data_perwakilan_dokumen/"
}

func (MediaBrandDataPerwakilanDokumen) TableName() string {
	return "media_brand_data_perwakilan_dokumen"
}

type MediaBrandDataSertifikatFoto struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id_media_brand_data_sertifikat_foto"`
	IdBrandData int64     `gorm:"column:id_brand_data" json:"id_brand_data"`
	BrandData   BrandData `gorm:"foreignKey:IdBrandData;references:ID" json:"-"`
	Key         string    `gorm:"column:key;type:text;not null" json:"key"`
	Format      string    `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaBrandDataSertifikatFoto) PathName() string {
	return "/media_brand_data_sertifikat_foto/"
}

func (MediaBrandDataSertifikatFoto) TableName() string {
	return "media_brand_data_sertifikat_foto"
}

type MediaBrandDataNIBFoto struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id_media_brand_data_nib_foto"`
	IdBrandData int64     `gorm:"column:id_brand_data" json:"id_brand_data"`
	BrandData   BrandData `gorm:"foreignKey:IdBrandData;references:ID" json:"-"`
	Key         string    `gorm:"column:key;type:text;not null" json:"key"`
	Format      string    `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaBrandDataNIBFoto) PathName() string {
	return "/media_brand_data_nib_foto/"
}

func (MediaBrandDataNIBFoto) TableName() string {
	return "media_brand_data_nib_foto"
}

type MediaBrandDataNPWPFoto struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id_media_brand_data_npwp_foto"`
	IdBrandData int64     `gorm:"column:id_brand_data" json:"id_brand_data"`
	BrandData   BrandData `gorm:"foreignKey:IdBrandData;references:ID" json:"-"`
	Key         string    `gorm:"column:key;type:text;not null" json:"key"`
	Format      string    `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaBrandDataNPWPFoto) PathName() string {
	return "/media_brand_data_npwp_foto/"
}

func (MediaBrandDataNPWPFoto) TableName() string {
	return "media_brand_data_npwp_foto"
}

type MediaBrandDataLogoFoto struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id_media_brand_data_logo_foto"`
	IdBrandData int64     `gorm:"column:id_brand_data" json:"id_brand_data"`
	BrandData   BrandData `gorm:"foreignKey:IdBrandData;references:ID" json:"-"`
	Key         string    `gorm:"column:key;type:text;not null" json:"key"`
	Format      string    `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaBrandDataLogoFoto) PathName() string {
	return "/media_brand_data_logo_foto/"
}

func (MediaBrandDataLogoFoto) TableName() string {
	return "media_brand_data_logo_foto"
}

type MediaBrandDataSuratKerjasamaDokumen struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id_media_brand_data_surat_kerjasama_dokumen"`
	IdBrandData int64     `gorm:"column:id_brand_data" json:"id_brand_data"`
	BrandData   BrandData `gorm:"foreignKey:IdBrandData;references:ID" json:"-"`
	Key         string    `gorm:"column:key;type:text;not null" json:"key"`
	Format      string    `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaBrandDataSuratKerjasamaDokumen) PathName() string {
	return "/media_brand_data_surat_kerjasama_dokumen/"
}

func (MediaBrandDataSuratKerjasamaDokumen) TableName() string {
	return "media_brand_data_surat_kerjasama_dokumen"
}

type MediaInformasiKendaraanKurirKendaraanFoto struct {
	ID                        int64                   `gorm:"primaryKey;autoIncrement" json:"id_media_informasi_kendaraan_kurir_kendaraan_foto"`
	IdInformasiKendaraanKurir int64                   `gorm:"column:id_informasi_kendaraan_kurir;not null" json:"id_informasi_kendaraan_kurir"`
	InformasiKendaraanKurir   InformasiKendaraanKurir `gorm:"foreignKey:IdInformasiKendaraanKurir;references:ID" json:"-"`
	Key                       string                  `gorm:"column:key;type:text;not null" json:"key"`
	Format                    string                  `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaInformasiKendaraanKurirKendaraanFoto) PathName() string {
	return "/media_informasi_kendaraan_kurir_kendaraan_foto/"
}

func (MediaInformasiKendaraanKurirKendaraanFoto) TableName() string {
	return "media_informasi_kendaraan_kurir_kendaraan_foto"
}

type MediaInformasiKendaraanKurirBPKBFoto struct {
	ID                        int64                   `gorm:"primaryKey;autoIncrement" json:"id_media_informasi_kendaraan_kurir_bpkb_foto"`
	IdInformasiKendaraanKurir int64                   `gorm:"column:id_informasi_kendaraan_kurir;not null" json:"id_informasi_kendaraan_kurir"`
	InformasiKendaraanKurir   InformasiKendaraanKurir `gorm:"foreignKey:IdInformasiKendaraanKurir;references:ID" json:"-"`
	Key                       string                  `gorm:"column:key;type:text;not null" json:"key"`
	Format                    string                  `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaInformasiKendaraanKurirBPKBFoto) PathName() string {
	return "/media_informasi_kendaraan_kurir_bpkb_foto/"
}

func (MediaInformasiKendaraanKurirBPKBFoto) TableName() string {
	return "media_informasi_kendaraan_kurir_bpkb_foto"
}

type MediaInformasiKendaraanKurirSTNKFoto struct {
	ID                        int64                   `gorm:"primaryKey;autoIncrement" json:"id_media_informasi_kendaraan_kurir_stnk_foto"`
	IdInformasiKendaraanKurir int64                   `gorm:"column:id_informasi_kendaraan_kurir;not null" json:"id_informasi_kendaraan_kurir"`
	InformasiKendaraanKurir   InformasiKendaraanKurir `gorm:"foreignKey:IdInformasiKendaraanKurir;references:ID" json:"-"`
	Key                       string                  `gorm:"column:key;type:text;not null" json:"key"`
	Format                    string                  `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaInformasiKendaraanKurirSTNKFoto) PathName() string {
	return "/media_informasi_kendaraan_kurir_stnk_foto/"
}

func (MediaInformasiKendaraanKurirSTNKFoto) TableName() string {
	return "media_informasi_kendaraan_kurir_stnk_foto"
}

type MediaInformasiKurirKTPFoto struct {
	ID               int64          `gorm:"primaryKey;autoIncrement" json:"id_media_informasi_kurir_ktp_foto"`
	IdInformasiKurir int64          `gorm:"column:id_informasi_kurir;not null" json:"id_informasi_kurir"`
	InformasiKurir   InformasiKurir `gorm:"foreignKey:IdInformasiKurir;references:ID" json:"-"`
	Key              string         `gorm:"column:key;type:text;not null" json:"key"`
	Format           string         `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaInformasiKurirKTPFoto) PathName() string {
	return "/media_informasi_kurir_ktp_foto/"
}

func (MediaInformasiKurirKTPFoto) TableName() string {
	return "media_informasi_kurir_ktp_foto"
}

type MediaReviewFoto struct {
	ID       int64  `gorm:"primaryKey;autoIncrement" json:"id_media_review_foto"`
	IdReview int64  `gorm:"column:id_review;not null" json:"id_review"`
	Review   Review `gorm:"foreignKey:IdReview;references:ID" json:"-"`
	Key      string `gorm:"column:key;type:text;not null" json:"key"`
	Format   string `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaReviewFoto) PathName() string {
	return "/media_review_foto/"
}

func (MediaReviewFoto) TableName() string {
	return "media_review_foto"
}

type MediaReviewVideo struct {
	ID       int64  `gorm:"primaryKey;autoIncrement" json:"id_media_review_video"`
	IdReview int64  `gorm:"column:id_review;not null" json:"id_review"`
	Review   Review `gorm:"foreignKey:IdReview;references:ID" json:"-"`
	Key      string `gorm:"column:key;type:text;not null" json:"key"`
	Format   string `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaReviewVideo) PathName() string {
	return "/media_review_video/"
}

func (MediaReviewVideo) TableName() string {
	return "media_review_video"
}

type MediaTransaksiApprovedFoto struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id_media_transaksi_approved_foto"`
	IdTransaksi int64     `gorm:"column:id_transaksi;not null" json:"id_transaksi"`
	Transaksi   Transaksi `gorm:"foreignKey:IdTransaksi;references:ID" json:"-"`
	Key         string    `gorm:"column:key;type:text;not null" json:"key"`
	Format      string    `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaTransaksiApprovedFoto) PathName() string {
	return "/media_transaksi_approved_foto/"
}

func (MediaTransaksiApprovedFoto) TableName() string {
	return "media_transaksi_approved_foto"
}

type MediaTransaksiApprovedVideo struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id_media_transaksi_approved_video"`
	IdTransaksi int64     `gorm:"column:id_transaksi;not null" json:"id_transaksi"`
	Transaksi   Transaksi `gorm:"foreignKey:IdTransaksi;references:ID" json:"-"`
	Key         string    `gorm:"column:key;type:text;not null" json:"key"`
	Format      string    `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaTransaksiApprovedVideo) PathName() string {
	return "/media_transaksi_approved_video/"
}

func (MediaTransaksiApprovedVideo) TableName() string {
	return "media_transaksi_approved_video"
}

type MediaPengirimanPickedUpFoto struct {
	ID           int64      `gorm:"primaryKey;autoIncrement" json:"id_media_pengiriman_picked_up_foto"`
	IdPengiriman int64      `gorm:"column:id_pengiriman;not null" json:"id_pengiriman"`
	Pengiriman   Pengiriman `gorm:"foreignKey:IdPengiriman;references:ID" json:"-"`
	Key          string     `gorm:"column:key;type:text;not null" json:"key"`
	Format       string     `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaPengirimanPickedUpFoto) PathName() string {
	return "/media_pengiriman_picked_up_foto/"
}

func (MediaPengirimanPickedUpFoto) TableName() string {
	return "media_pengiriman_picked_up_foto"
}

type MediaPengirimanSampaiFoto struct {
	ID           int64      `gorm:"primaryKey;autoIncrement" json:"id_media_pengiriman_sampai_foto"`
	IdPengiriman int64      `gorm:"column:id_pengiriman;not null" json:"id_pengiriman"`
	Pengiriman   Pengiriman `gorm:"foreignKey:IdPengiriman;references:ID" json:"-"`
	Key          string     `gorm:"column:key;type:text;not null" json:"key"`
	Format       string     `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaPengirimanSampaiFoto) PathName() string {
	return "/media_pengiriman_sampai_foto/"

}

func (MediaPengirimanSampaiFoto) TableName() string {
	return "media_pengiriman_sampai_foto"
}

type MediaPengirimanEkspedisiPickedUpFoto struct {
	ID                    int64               `gorm:"primaryKey;autoIncrement" json:"id_media_pengiriman_ekspedisi_approved_foto"`
	IdPengirimanEkspedisi int64               `gorm:"column:id_pengiriman_ekspedisi;not null" json:"id_pengiriman_ekspedisi"`
	PengirimanEkspedisi   PengirimanEkspedisi `gorm:"foreignKey:IdPengirimanEkspedisi;references:ID" json:"-"`
	Key                   string              `gorm:"column:key;type:text;not null" json:"key"`
	Format                string              `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaPengirimanEkspedisiPickedUpFoto) PathName() string {
	return "/media_pengiriman_ekspedisi_picked_up_foto/"
}

func (MediaPengirimanEkspedisiPickedUpFoto) TableName() string {
	return "media_pengiriman_ekspedisi_picked_up_foto"
}

type MediaPengirimanEkspedisiSampaiAgentFoto struct {
	ID                    int64               `gorm:"primaryKey;autoIncrement" json:"id_media_pengiriman_ekspedisi_sampai_agent_foto"`
	IdPengirimanEkspedisi int64               `gorm:"column:id_pengiriman_ekspedisi;not null" json:"id_pengiriman_ekspedisi"`
	PengirimanEkspedisi   PengirimanEkspedisi `gorm:"foreignKey:IdPengirimanEkspedisi;references:ID" json:"-"`
	Key                   string              `gorm:"column:key;type:text;not null" json:"key"`
	Format                string              `gorm:"column:format;type:varchar(20);not null" json:"format"`
}

func (MediaPengirimanEkspedisiSampaiAgentFoto) PathName() string {
	return "/media_pengiriman_ekspedisi_sampai_agent_foto/"
}

func (MediaPengirimanEkspedisiSampaiAgentFoto) TableName() string {
	return "media_pengiriman_ekspedisi_sampai_agent_foto"
}
