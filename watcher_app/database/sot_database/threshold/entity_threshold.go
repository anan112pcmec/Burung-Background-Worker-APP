package sot_threshold

import (
	"context"
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	sot_threshold_seeders_nama "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/threshold/seeders/nama_threshold"
	"gorm.io/gorm"
)

type PenggunaThreshold struct {
	ID         int64           `gorm:"primaryKey;autoIncrement" json:"id_pengguna_threshold"`
	IdPengguna int64           `gorm:"column:id_pengguna;not null" json:"id_pengguna"`
	Pengguna   models.Pengguna `gorm:"foreignKey:IdPengguna;references:ID" json:"-"`

	// Counters
	AlamatPengguna      int32 `gorm:"column:alamat_pengguna;type:int4;default:0" json:"alamat_pengguna"`
	BarangDisukai       int32 `gorm:"column:barang_disukai;type:int4;default:0" json:"barang_disukai"`
	Following           int32 `gorm:"column:following;type:int4;default:0" json:"following"`
	Keranjang           int32 `gorm:"column:keranjang;type:int4;default:0" json:"keranjang"`
	MediaPenggunaProfil int32 `gorm:"column:media_pengguna_profil_foto;type:int4;default:0" json:"media_pengguna_profil_foto"`
	Pembayaran          int32 `gorm:"column:pembayaran;type:int4;default:0" json:"pembayaran"`
	Review              int32 `gorm:"column:review;type:int4;default:0" json:"review"`
	ReviewDislike       int32 `gorm:"column:review_dislike;type:int4;default:0" json:"review_dislike"`
	Transaksi           int32 `gorm:"column:transaksi;type:int4;default:0" json:"transaksi"`
	Wishlist            int32 `gorm:"column:wishlist;type:int4;default:0" json:"wishlist"`
}

func (PenggunaThreshold) TableName() string {
	return sot_threshold_seeders_nama.PenggunaThreshold
}

func (p *PenggunaThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&PenggunaThreshold{}).Select("id").Where(&PenggunaThreshold{
		IdPengguna: p.IdPengguna,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&PenggunaThreshold{
		IdPengguna: p.IdPengguna,
	}).Error
}

func (p *PenggunaThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&PenggunaThreshold{}).Where(&PenggunaThreshold{
		IdPengguna: p.IdPengguna,
	}).Updates(updates).Error
}

func (p *PenggunaThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&PenggunaThreshold{}).Where(&PenggunaThreshold{
		IdPengguna: p.IdPengguna,
	}).Updates(updates).Error
}

func (p *PenggunaThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&PenggunaThreshold{}).Where(&PenggunaThreshold{
		IdPengguna: p.IdPengguna,
	}).Updates(
		updates,
	).Error
}

func (p *PenggunaThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&PenggunaThreshold{}).Where(&PenggunaThreshold{
		IdPengguna: p.IdPengguna,
	}).Updates(
		updates,
	).Error
}

type SellerThreshold struct {
	ID       int64         `gorm:"primaryKey;autoIncrement" json:"id_seller_threshold"`
	IdSeller int64         `gorm:"column:id_seller;not null" json:"id_seller"`
	Seller   models.Seller `gorm:"foreignKey:IdSeller;references:ID" json:"-"`

	// Counters
	AlamatGudang              int32 `gorm:"column:alamat_gudang;type:int4;default:0" json:"alamat_gudang"`
	BarangDiDiskon            int32 `gorm:"column:barang_di_diskon;type:int4;default:0" json:"barang_di_diskon"`
	BarangInduk               int32 `gorm:"column:barang_induk;type:int4;default:0" json:"barang_induk"`
	BrandData                 int32 `gorm:"column:brand_data;type:int4;default:0" json:"brand_data"`
	DiskonProduk              int32 `gorm:"column:diskon_produk;type:int4;default:0" json:"diskon_produk"`
	DistributorData           int32 `gorm:"column:distributor_data;type:int4;default:0" json:"distributor_data"`
	Etalase                   int32 `gorm:"column:etalase;type:int4;default:0" json:"etalase"`
	Follower                  int32 `gorm:"column:follower;type:int4;default:0" json:"follower"`
	JenisSellerValidation     int32 `gorm:"column:jenis_seller_validation;type:int4;default:0" json:"jenis_seller_validation"`
	Rekening                  int32 `gorm:"column:rekening;type:int4;default:0" json:"rekening"`
	KategoriBarang            int32 `gorm:"column:kategori_barang;type:int4;default:0" json:"kategori_barang"`
	Keranjang                 int32 `gorm:"column:keranjang;type:int4;default:0" json:"keranjang"`
	MediaSellerBannerFoto     int32 `gorm:"column:media_seller_banner_foto;type:int4;default:0" json:"media_seller_banner_foto"`
	MediaSellerFotoProfil     int32 `gorm:"column:media_seller_foto_profil;type:int4;default:0" json:"media_seller_foto_profil"`
	MediaSellerTokoFisikFoto  int32 `gorm:"column:media_seller_toko_fisik_foto;type:int4;default:0" json:"media_seller_toko_fisik_foto"`
	PayoutSeller              int32 `gorm:"column:payout_seller;type:int4;default:0" json:"payout_seller"`
	Pengiriman                int32 `gorm:"column:pengiriman;type:int4;default:0" json:"pengiriman"`
	PengirimanEkspedisi       int32 `gorm:"column:pengiriman_ekspedisi;type:int4;default:0" json:"pengiriman_ekspedisi"`
	ThresholdOrderSellers     int32 `gorm:"column:threshold_order_sellers;type:int4;default:0" json:"threshold_order_sellers"`
	ThresholdTransaksiSellers int32 `gorm:"column:threshold_transaksi_sellers;type:int4;default:0" json:"threshold_transaksi_sellers"`
	Transaksi                 int32 `gorm:"column:transaksi;type:int4;default:0" json:"transaksi"`
}

func (SellerThreshold) TableName() string {
	return sot_threshold_seeders_nama.SellerThreshold
}

func (s *SellerThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&SellerThreshold{}).Select("id").Where(&SellerThreshold{
		IdSeller: s.IdSeller,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&SellerThreshold{
		IdSeller: s.IdSeller,
	}).Error
}

func (s *SellerThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&SellerThreshold{}).Where(&SellerThreshold{
		IdSeller: s.IdSeller,
	}).Updates(updates).Error
}

func (s *SellerThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&SellerThreshold{}).Where(&SellerThreshold{
		IdSeller: s.IdSeller,
	}).Updates(updates).Error
}

func (s *SellerThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&SellerThreshold{}).Where(&SellerThreshold{
		IdSeller: s.IdSeller,
	}).Updates(
		updates,
	).Error
}

func (s *SellerThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&SellerThreshold{}).Where(&SellerThreshold{
		IdSeller: s.IdSeller,
	}).Updates(
		updates,
	).Error
}

type KurirThreshold struct {
	ID      int64        `gorm:"primaryKey;autoIncrement" json:"id_kurir_threshold"`
	IdKurir int64        `gorm:"column:id_kurir;not null" json:"id_kurir"`
	Kurir   models.Kurir `gorm:"foreignKey:IdKurir;references:ID" json:"-"`

	// Counters
	AlamatKurir             int32 `gorm:"column:alamat_kurir;type:int4;default:0" json:"alamat_kurir"`
	BidKurirData            int32 `gorm:"column:bid_kurir_data;type:int4;default:0" json:"bid_kurir_data"`
	BidKurirEksScheduler    int32 `gorm:"column:bid_kurir_eks_scheduler;type:int4;default:0" json:"bid_kurir_eks_scheduler"`
	BidKurirNonEksScheduler int32 `gorm:"column:bid_kurir_non_eks_scheduler;type:int4;default:0" json:"bid_kurir_non_eks_scheduler"`
	InformasiKendaraanKurir int32 `gorm:"column:informasi_kendaraan_kurir;type:int4;default:0" json:"informasi_kendaraan_kurir"`
	InformasiKurir          int32 `gorm:"column:informasi_kurir;type:int4;default:0" json:"informasi_kurir"`
	MediaKurirProfilFoto    int32 `gorm:"column:media_kurir_profil_foto;type:int4;default:0" json:"media_kurir_profil_foto"`
	PayoutKurir             int32 `gorm:"column:payout_kurir;type:int4;default:0" json:"payout_kurir"`
	RekeningKurir           int32 `gorm:"column:rekening_kurir;type:int4;default:0" json:"rekening_kurir"`
}

func (KurirThreshold) TableName() string {
	return sot_threshold_seeders_nama.KurirThreshold
}

func (k *KurirThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&KurirThreshold{}).Select("id").Where(&KurirThreshold{
		IdKurir: k.IdKurir,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&KurirThreshold{
		IdKurir: k.IdKurir,
	}).Error
}

func (k *KurirThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&KurirThreshold{}).Where(&KurirThreshold{
		IdKurir: k.IdKurir,
	}).Updates(updates).Error
}

func (k *KurirThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&KurirThreshold{}).Where(&KurirThreshold{
		IdKurir: k.IdKurir,
	}).Updates(updates).Error
}

func (k *KurirThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&KurirThreshold{}).Where(&KurirThreshold{
		IdKurir: k.IdKurir,
	}).Updates(
		updates,
	).Error
}

func (k *KurirThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&KurirThreshold{}).Where(&KurirThreshold{
		IdKurir: k.IdKurir,
	}).Updates(
		updates,
	).Error
}
