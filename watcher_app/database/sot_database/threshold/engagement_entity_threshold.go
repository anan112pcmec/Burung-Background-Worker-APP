package sot_threshold

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	sot_threshold_seeders_nama "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/threshold/seeders/nama_threshold"
)

type BrandDataThreshold struct {
	ID          int64                `gorm:"primaryKey;autoIncrement" json:"id_brand_data_threshold"`
	IdBrandData int64                `gorm:"column:id_brand_data;index;not null" json:"id_brand_data"`
	BrandData   sot_models.BrandData `gorm:"foreignKey:IdBrandData;references:ID" json:"-"`

	MediaBrandDataLogoBrandFoto         int32 `gorm:"column:media_brand_data_logo_brand_foto;type:int4;default:0" json:"media_brand_data_logo_brand_foto"`
	MediaBrandDataNibFoto               int32 `gorm:"column:media_brand_data_nib_foto;type:int4;default:0" json:"media_brand_data_nib_foto"`
	MediaBrandDataNpwpFoto              int32 `gorm:"column:media_brand_data_npwp_foto;type:int4;default:0" json:"media_brand_data_npwp_foto"`
	MediaBrandDataPerwakilanDokumen     int32 `gorm:"column:media_brand_data_perwakilan_dokumen;type:int4;default:0" json:"media_brand_data_perwakilan_dokumen"`
	MediaBrandDataSertifikatFoto        int32 `gorm:"column:media_brand_data_sertifikat_foto;type:int4;default:0" json:"media_brand_data_sertifikat_foto"`
	MediaBrandDataSuratKerjasamaDokumen int32 `gorm:"column:media_brand_data_surat_kerjasama_dokumen;type:int4;default:0" json:"media_brand_data_surat_kerjasama_dokumen"`
}

func (BrandDataThreshold) TableName() string {
	return sot_threshold_seeders_nama.BrandDataThreshold
}

func (b *BrandDataThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&BrandDataThreshold{}).Select("id").Where(&BrandDataThreshold{
		IdBrandData: b.IdBrandData,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&BrandDataThreshold{
		IdBrandData: b.IdBrandData,
	}).Error
}

func (b *BrandDataThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&BrandDataThreshold{}).Where(&BrandDataThreshold{
		IdBrandData: b.IdBrandData,
	}).Updates(updates).Error
}

func (b *BrandDataThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&BrandDataThreshold{}).Where(&BrandDataThreshold{
		IdBrandData: b.IdBrandData,
	}).Updates(updates).Error
}

func (b *BrandDataThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&BrandDataThreshold{}).Where(&BrandDataThreshold{
		IdBrandData: b.IdBrandData,
	}).Updates(
		updates,
	).Error
}

func (b *BrandDataThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&BrandDataThreshold{}).Where(&BrandDataThreshold{
		IdBrandData: b.IdBrandData,
	}).Updates(
		updates,
	).Error
}

type DistributorDataThreshold struct {
	ID                int64                      `gorm:"primaryKey;autoIncrement" json:"id_distributor_data_threshold"`
	IdDistributorData int64                      `gorm:"column:id_distributor_data;index;not null" json:"id_distributor_data"`
	DistributorData   sot_models.DistributorData `gorm:"foreignKey:IdDistributorData;references:ID" json:"-"`

	MediaDistributorDataDokumen               int32 `gorm:"column:media_distributor_data_dokumen;type:int4;default:0" json:"media_distributor_data_dokumen"`
	MediaDistributorDataNibFoto               int32 `gorm:"column:media_distributor_data_nib_foto;type:int4;default:0" json:"media_distributor_data_nib_foto"`
	MediaDistributorDataNpwpFoto              int32 `gorm:"column:media_distributor_data_npwp_foto;type:int4;default:0" json:"media_distributor_data_npwp_foto"`
	MediaDistributorDataSuratKerjasamaDokumen int32 `gorm:"column:media_distributor_data_surat_kerjasama_dokumen;type:int4;default:0" json:"media_distributor_data_surat_kerjasama_dokumen"`
}

func (DistributorDataThreshold) TableName() string {
	return sot_threshold_seeders_nama.DistributorDataThreshold
}

func (d *DistributorDataThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&DistributorDataThreshold{}).Select("id").Where(&DistributorDataThreshold{
		IdDistributorData: d.IdDistributorData,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&DistributorDataThreshold{
		IdDistributorData: d.IdDistributorData,
	}).Error
}

func (d *DistributorDataThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&DistributorDataThreshold{}).Where(&DistributorDataThreshold{
		IdDistributorData: d.IdDistributorData,
	}).Updates(updates).Error
}

func (d *DistributorDataThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&DistributorDataThreshold{}).Where(&DistributorDataThreshold{
		IdDistributorData: d.IdDistributorData,
	}).Updates(updates).Error
}

func (d *DistributorDataThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&DistributorDataThreshold{}).Where(&DistributorDataThreshold{
		IdDistributorData: d.IdDistributorData,
	}).Updates(
		updates,
	).Error
}

func (d *DistributorDataThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&DistributorDataThreshold{}).Where(&DistributorDataThreshold{
		IdDistributorData: d.IdDistributorData,
	}).Updates(
		updates,
	).Error
}

// 13. EtalaseThreshold
type EtalaseThreshold struct {
	ID        int64              `gorm:"primaryKey;autoIncrement" json:"id_etalase_threshold"`
	IdEtalase int64              `gorm:"column:id_etalase;index;not null" json:"id_etalase"`
	Etalase   sot_models.Etalase `gorm:"foreignKey:IdEtalase;references:ID" json:"-"`

	BarangKeEtalase  int32 `gorm:"column:barang_ke_etalase;type:int4;default:0" json:"barang_ke_etalase"`
	MediaEtalaseFoto int32 `gorm:"column:media_etalase_foto;type:int4;default:0" json:"media_etalase_foto"`
}

func (EtalaseThreshold) TableName() string {
	return sot_threshold_seeders_nama.EtalaseThreshold
}

func (e *EtalaseThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&EtalaseThreshold{}).Select("id").Where(&EtalaseThreshold{
		IdEtalase: e.IdEtalase,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&EtalaseThreshold{
		IdEtalase: e.IdEtalase,
	}).Error
}

func (e *EtalaseThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&EtalaseThreshold{}).Where(&EtalaseThreshold{
		IdEtalase: e.IdEtalase,
	}).Updates(updates).Error
}

func (e *EtalaseThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&EtalaseThreshold{}).Where(&EtalaseThreshold{
		IdEtalase: e.IdEtalase,
	}).Updates(updates).Error
}

func (e *EtalaseThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&EtalaseThreshold{}).Where(&EtalaseThreshold{
		IdEtalase: e.IdEtalase,
	}).Updates(
		updates,
	).Error
}

func (e *EtalaseThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&EtalaseThreshold{}).Where(&EtalaseThreshold{
		IdEtalase: e.IdEtalase,
	}).Updates(
		updates,
	).Error
}

// 14. ReviewThreshold
type ReviewThreshold struct {
	ID       int64             `gorm:"primaryKey;autoIncrement" json:"id_review_threshold"`
	IdReview int64             `gorm:"column:id_review;index;not null" json:"id_review"`
	Review   sot_models.Review `gorm:"foreignKey:IdReview;references:ID" json:"-"`

	MediaReviewFoto  int32 `gorm:"column:media_review_foto;type:int4;default:0" json:"media_review_foto"`
	MediaReviewVideo int32 `gorm:"column:media_review_video;type:int4;default:0" json:"media_review_video"`
	ReviewLike       int32 `gorm:"column:review_like;type:int4;default:0" json:"review_like"`
	ReviewDislike    int32 `gorm:"column:review_dislike;type:int4;default:0" json:"review_dislike"`
}

func (ReviewThreshold) TableName() string {
	return sot_threshold_seeders_nama.ReviewThreshold
}

func (r *ReviewThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&ReviewThreshold{}).Select("id").Where(&ReviewThreshold{
		IdReview: r.IdReview,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&ReviewThreshold{
		IdReview: r.IdReview,
	}).Error
}

func (r *ReviewThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&ReviewThreshold{}).Where(&ReviewThreshold{
		IdReview: r.IdReview,
	}).Updates(updates).Error
}

func (r *ReviewThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&ReviewThreshold{}).Where(&ReviewThreshold{
		IdReview: r.IdReview,
	}).Updates(updates).Error
}

func (r *ReviewThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&ReviewThreshold{}).Where(&ReviewThreshold{
		IdReview: r.IdReview,
	}).Updates(
		updates,
	).Error
}

func (r *ReviewThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&ReviewThreshold{}).Where(&ReviewThreshold{
		IdReview: r.IdReview,
	}).Updates(
		updates,
	).Error
}

// 15. InformasiKendaraanKurirThreshold
type InformasiKendaraanKurirThreshold struct {
	ID                        int64                              `gorm:"primaryKey;autoIncrement" json:"id_informasi_kendaraan_kurir_threshold"`
	IdInformasiKendaraanKurir int64                              `gorm:"column:id_informasi_kendaraan_kurir;index;not null" json:"id_informasi_kendaraan_kurir"`
	InformasiKendaraanKurir   sot_models.InformasiKendaraanKurir `gorm:"foreignKey:IdInformasiKendaraanKurir;references:ID" json:"-"`

	MediaBpkbFoto      int32 `gorm:"column:media_informasi_kendaraan_kurir_bpkb_foto;type:int4;default:0" json:"media_informasi_kendaraan_kurir_bpkb_foto"`
	MediaKendaraanFoto int32 `gorm:"column:media_informasi_kendaraan_kurir_kendaraan_foto;type:int4;default:0" json:"media_informasi_kendaraan_kurir_kendaraan_foto"`
	MediaStnkFoto      int32 `gorm:"column:media_informasi_kendaraan_kurir_stnk_foto;type:int4;default:0" json:"media_informasi_kendaraan_kurir_stnk_foto"`
}

func (InformasiKendaraanKurirThreshold) TableName() string {
	return sot_threshold_seeders_nama.InformasiKendaraanKurirThreshold
}

func (i *InformasiKendaraanKurirThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&InformasiKendaraanKurirThreshold{}).Select("id").Where(&InformasiKendaraanKurirThreshold{
		IdInformasiKendaraanKurir: i.IdInformasiKendaraanKurir,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&InformasiKendaraanKurirThreshold{
		IdInformasiKendaraanKurir: i.IdInformasiKendaraanKurir,
	}).Error
}

func (i *InformasiKendaraanKurirThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&InformasiKendaraanKurirThreshold{}).Where(&InformasiKendaraanKurirThreshold{
		IdInformasiKendaraanKurir: i.IdInformasiKendaraanKurir,
	}).Updates(updates).Error
}

func (i *InformasiKendaraanKurirThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&InformasiKendaraanKurirThreshold{}).Where(&InformasiKendaraanKurirThreshold{
		IdInformasiKendaraanKurir: i.IdInformasiKendaraanKurir,
	}).Updates(updates).Error
}

func (i *InformasiKendaraanKurirThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&InformasiKendaraanKurirThreshold{}).Where(&InformasiKendaraanKurirThreshold{
		IdInformasiKendaraanKurir: i.IdInformasiKendaraanKurir,
	}).Updates(
		updates,
	).Error
}

func (i *InformasiKendaraanKurirThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&InformasiKendaraanKurirThreshold{}).Where(&InformasiKendaraanKurirThreshold{
		IdInformasiKendaraanKurir: i.IdInformasiKendaraanKurir,
	}).Updates(
		updates,
	).Error
}

// 16. BidKurirDataThreshold
type BidKurirDataThreshold struct {
	ID             int64                   `gorm:"primaryKey;autoIncrement" json:"id_bid_kurir_data_threshold"`
	IdBidKurirData int64                   `gorm:"column:id_bid_kurir_data;index;not null" json:"id_bid_kurir_data"`
	BidKurirData   sot_models.BidKurirData `gorm:"foreignKey:IdBidKurirData;references:ID" json:"-"`

	BidKurirEksScheduler    int32 `gorm:"column:bid_kurir_eks_scheduler;type:int4;default:0" json:"bid_kurir_eks_scheduler"`
	BidKurirNonEksScheduler int32 `gorm:"column:bid_kurir_non_eks_scheduler;type:int4;default:0" json:"bid_kurir_non_eks_scheduler"`
}

func (BidKurirDataThreshold) TableName() string {
	return sot_threshold_seeders_nama.BidKurirDataThreshold
}

func (b *BidKurirDataThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&BidKurirDataThreshold{}).Select("id").Where(&BidKurirDataThreshold{
		IdBidKurirData: b.IdBidKurirData,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&BidKurirDataThreshold{
		IdBidKurirData: b.IdBidKurirData,
	}).Error
}

func (b *BidKurirDataThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&BidKurirDataThreshold{}).Where(&BidKurirDataThreshold{
		IdBidKurirData: b.IdBidKurirData,
	}).Updates(updates).Error
}

func (b *BidKurirDataThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&BidKurirDataThreshold{}).Where(&BidKurirDataThreshold{
		IdBidKurirData: b.IdBidKurirData,
	}).Updates(updates).Error
}

func (b *BidKurirDataThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&BidKurirDataThreshold{}).Where(&BidKurirDataThreshold{
		IdBidKurirData: b.IdBidKurirData,
	}).Updates(
		updates,
	).Error
}

func (b *BidKurirDataThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&BidKurirDataThreshold{}).Where(&BidKurirDataThreshold{
		IdBidKurirData: b.IdBidKurirData,
	}).Updates(
		updates,
	).Error
}

// 17. DiskonProdukThreshold
type DiskonProdukThreshold struct {
	ID             int64                   `gorm:"primaryKey;autoIncrement" json:"id_diskon_produk_threshold"`
	IdDiskonProduk int64                   `gorm:"column:id_diskon_produk;index;not null" json:"id_diskon_produk"`
	DiskonProduk   sot_models.DiskonProduk `gorm:"foreignKey:IdDiskonProduk;references:ID" json:"-"`

	BarangDiDiskon int32 `gorm:"column:barang_di_diskon;type:int4;default:0" json:"barang_di_diskon"`
}

func (DiskonProdukThreshold) TableName() string {
	return sot_threshold_seeders_nama.DiskonProdukThreshold
}

func (d *DiskonProdukThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&DiskonProdukThreshold{}).Select("id").Where(&DiskonProdukThreshold{
		IdDiskonProduk: d.IdDiskonProduk,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&DiskonProdukThreshold{
		IdDiskonProduk: d.IdDiskonProduk,
	}).Error
}

func (d *DiskonProdukThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&DiskonProdukThreshold{}).Where(&DiskonProdukThreshold{
		IdDiskonProduk: d.IdDiskonProduk,
	}).Updates(updates).Error
}

func (d *DiskonProdukThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&DiskonProdukThreshold{}).Where(&DiskonProdukThreshold{
		IdDiskonProduk: d.IdDiskonProduk,
	}).Updates(updates).Error
}

func (d *DiskonProdukThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&DiskonProdukThreshold{}).Where(&DiskonProdukThreshold{
		IdDiskonProduk: d.IdDiskonProduk,
	}).Updates(
		updates,
	).Error
}

func (d *DiskonProdukThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&DiskonProdukThreshold{}).Where(&DiskonProdukThreshold{
		IdDiskonProduk: d.IdDiskonProduk,
	}).Updates(
		updates,
	).Error
}

// 18. KomentarThreshold
type KomentarThreshold struct {
	ID         int64               `gorm:"primaryKey;autoIncrement" json:"id_komentar_threshold"`
	IdKomentar int64               `gorm:"column:id_komentar;index;not null" json:"id_komentar"`
	Komentar   sot_models.Komentar `gorm:"foreignKey:IdKomentar;references:ID" json:"-"`

	KomentarChild int32 `gorm:"column:komentar_child;type:int4;default:0" json:"komentar_child"`
}

func (KomentarThreshold) TableName() string {
	return sot_threshold_seeders_nama.KomentarThreshold
}

func (k *KomentarThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&KomentarThreshold{}).Select("id").Where(&KomentarThreshold{
		IdKomentar: k.IdKomentar,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&KomentarThreshold{
		IdKomentar: k.IdKomentar,
	}).Error
}

func (k *KomentarThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&KomentarThreshold{}).Where(&KomentarThreshold{
		IdKomentar: k.IdKomentar,
	}).Updates(updates).Error
}

func (k *KomentarThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&KomentarThreshold{}).Where(&KomentarThreshold{
		IdKomentar: k.IdKomentar,
	}).Updates(updates).Error
}

func (k *KomentarThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&KomentarThreshold{}).Where(&KomentarThreshold{
		IdKomentar: k.IdKomentar,
	}).Updates(
		updates,
	).Error
}

func (k *KomentarThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&KomentarThreshold{}).Where(&KomentarThreshold{
		IdKomentar: k.IdKomentar,
	}).Updates(
		updates,
	).Error
}

// 19. InformasiKurirThreshold
type InformasiKurirThreshold struct {
	ID               int64                     `gorm:"primaryKey;autoIncrement" json:"id_informasi_kurir_threshold"`
	IdInformasiKurir int64                     `gorm:"column:id_informasi_kurir;index;not null" json:"id_informasi_kurir"`
	InformasiKurir   sot_models.InformasiKurir `gorm:"foreignKey:IdInformasiKurir;references:ID" json:"-"`

	MediaKtpFoto int32 `gorm:"column:media_informasi_kurir_ktp_foto;type:int4;default:0" json:"media_informasi_kurir_ktp_foto"`
}

func (InformasiKurirThreshold) TableName() string {
	return sot_threshold_seeders_nama.InformasiKurirThreshold
}

func (i *InformasiKurirThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&InformasiKurirThreshold{}).Select("id").Where(&InformasiKurirThreshold{
		IdInformasiKurir: i.IdInformasiKurir,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&InformasiKurirThreshold{
		IdInformasiKurir: i.IdInformasiKurir,
	}).Error
}

func (i *InformasiKurirThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&InformasiKurirThreshold{}).Where(&InformasiKurirThreshold{
		IdInformasiKurir: i.IdInformasiKurir,
	}).Updates(updates).Error
}

func (i *InformasiKurirThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&InformasiKurirThreshold{}).Where(&InformasiKurirThreshold{
		IdInformasiKurir: i.IdInformasiKurir,
	}).Updates(updates).Error
}

func (i *InformasiKurirThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&InformasiKurirThreshold{}).Where(&InformasiKurirThreshold{
		IdInformasiKurir: i.IdInformasiKurir,
	}).Updates(
		updates,
	).Error
}

func (i *InformasiKurirThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&InformasiKurirThreshold{}).Where(&InformasiKurirThreshold{
		IdInformasiKurir: i.IdInformasiKurir,
	}).Updates(
		updates,
	).Error
}

// 20. PembayaranThreshold

// 21. RekeningSellerThreshold
type RekeningSellerThreshold struct {
	ID               int64                     `gorm:"primaryKey;autoIncrement" json:"id_rekening_seller_threshold"`
	IdRekeningSeller int64                     `gorm:"column:id_rekening_seller;index;not null" json:"id_rekening_seller"`
	RekeningSeller   sot_models.RekeningSeller `gorm:"foreignKey:IdRekeningSeller;references:ID" json:"-"`

	KategoriBarang int32 `gorm:"column:kategori_barang;type:int4;default:0" json:"kategori_barang"`
}

func (RekeningSellerThreshold) TableName() string {
	return sot_threshold_seeders_nama.RekeningSellerThreshold
}

func (r *RekeningSellerThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&RekeningSellerThreshold{}).Select("id").Where(&RekeningSellerThreshold{
		IdRekeningSeller: r.IdRekeningSeller,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&RekeningSellerThreshold{
		IdRekeningSeller: r.IdRekeningSeller,
	}).Error
}

func (r *RekeningSellerThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&RekeningSellerThreshold{}).Where(&RekeningSellerThreshold{
		IdRekeningSeller: r.IdRekeningSeller,
	}).Updates(updates).Error
}

func (r *RekeningSellerThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&RekeningSellerThreshold{}).Where(&RekeningSellerThreshold{
		IdRekeningSeller: r.IdRekeningSeller,
	}).Updates(updates).Error
}

func (r *RekeningSellerThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&RekeningSellerThreshold{}).Where(&RekeningSellerThreshold{
		IdRekeningSeller: r.IdRekeningSeller,
	}).Updates(
		updates,
	).Error
}

func (r *RekeningSellerThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&RekeningSellerThreshold{}).Where(&RekeningSellerThreshold{
		IdRekeningSeller: r.IdRekeningSeller,
	}).Updates(
		updates,
	).Error
}

type AlamatGudangThreshold struct {
	ID             int64                   `gorm:"primaryKey;autoIncrement" json:"id_alamat_gudang_threshold"`
	IdAlamatGudang int64                   `gorm:"column:id_alamat_gudang;index;not null" json:"id_alamat_gudang"`
	AlamatGudang   sot_models.AlamatGudang `gorm:"foreignKey:IdAlamatGudang;references:ID" json:"-"`

	KategoriBarang      int32 `gorm:"column:kategori_barang;type:int4;default:0" json:"kategori_barang"`
	Pengiriman          int32 `gorm:"column:pengiriman;type:int4;default:0" json:"pengiriman"`
	PengirimanEkspedisi int32 `gorm:"column:pengiriman_ekspedisi;type:int4;default:0" json:"pengiriman_ekspedisi"`
	Transaksi           int32 `gorm:"column:transaksi;type:int4;default:0" json:"transaksi"`
}

func (AlamatGudangThreshold) TableName() string {
	return sot_threshold_seeders_nama.AlamatGudangThreshold
}

func (a *AlamatGudangThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&AlamatGudangThreshold{}).Select("id").Where(&AlamatGudangThreshold{
		IdAlamatGudang: a.IdAlamatGudang,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&AlamatGudangThreshold{
		IdAlamatGudang: a.IdAlamatGudang,
	}).Error
}

func (a *AlamatGudangThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&AlamatGudangThreshold{}).Where(&AlamatGudangThreshold{
		IdAlamatGudang: a.IdAlamatGudang,
	}).Updates(updates).Error
}

func (a *AlamatGudangThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&AlamatGudangThreshold{}).Where(&AlamatGudangThreshold{
		IdAlamatGudang: a.IdAlamatGudang,
	}).Updates(updates).Error
}

func (a *AlamatGudangThreshold) CustomIncrement(id_fk int64, ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&AlamatGudangThreshold{}).Where(&AlamatGudangThreshold{
		ID: id_fk,
	}).Updates(
		updates,
	).Error
}

func (a AlamatGudangThreshold) CustomDecrement(id_fk int64, ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&AlamatGudangThreshold{}).Where(&AlamatGudangThreshold{
		ID: id_fk,
	}).Updates(
		updates,
	).Error
}

type AlamatPenggunaThreshold struct {
	ID               int64                     `gorm:"primaryKey;autoIncrement" json:"id_alamat_pengguna_threshold"`
	IdAlamatPengguna int64                     `gorm:"column:id_alamat_pengguna;index;not null" json:"id_alamat_pengguna"`
	AlamatPengguna   sot_models.AlamatPengguna `gorm:"foreignKey:IdAlamatPengguna;references:ID" json:"-"`

	Pengiriman int32 `gorm:"column:pengiriman;type:int4;default:0" json:"pengiriman"`
	Transaksi  int32 `gorm:"column:transaksi;type:int4;default:0" json:"transaksi"`
}

func (AlamatPenggunaThreshold) TableName() string {
	return sot_threshold_seeders_nama.AlamatPenggunaThreshold
}

func (a *AlamatPenggunaThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&AlamatPenggunaThreshold{}).Select("id").Where(&AlamatPenggunaThreshold{
		IdAlamatPengguna: a.IdAlamatPengguna,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&AlamatPenggunaThreshold{
		IdAlamatPengguna: a.IdAlamatPengguna,
	}).Error
}

func (a *AlamatPenggunaThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&AlamatPenggunaThreshold{}).Where(&AlamatPenggunaThreshold{
		IdAlamatPengguna: a.IdAlamatPengguna,
	}).Updates(updates).Error
}

func (a *AlamatPenggunaThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&AlamatPenggunaThreshold{}).Where(&AlamatPenggunaThreshold{
		IdAlamatPengguna: a.IdAlamatPengguna,
	}).Updates(updates).Error
}

func (a *AlamatPenggunaThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&AlamatPenggunaThreshold{}).Where(&AlamatPenggunaThreshold{
		IdAlamatPengguna: a.IdAlamatPengguna,
	}).Updates(
		updates,
	).Error
}

func (a *AlamatPenggunaThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&AlamatPenggunaThreshold{}).Where(&AlamatPenggunaThreshold{
		IdAlamatPengguna: a.IdAlamatPengguna,
	}).Updates(
		updates,
	).Error
}

// 22. AlamatEkspedisiThreshold
type AlamatEkspedisiThreshold struct {
	ID                int64                      `gorm:"primaryKey;autoIncrement" json:"id_alamat_ekspedisi_threshold"`
	IdAlamatEkspedisi int64                      `gorm:"column:id_alamat_ekspedisi;index;not null" json:"id_alamat_ekspedisi"`
	AlamatEkspedisi   sot_models.AlamatEkspedisi `gorm:"foreignKey:IdAlamatEkspedisi;references:ID" json:"-"`

	PengirimanEkspedisi int32 `gorm:"column:pengiriman_ekspedisi;type:int4;default:0" json:"pengiriman_ekspedisi"`
}

func (AlamatEkspedisiThreshold) TableName() string {
	return sot_threshold_seeders_nama.AlamatEkspedisiThreshold
}

func (a *AlamatEkspedisiThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&AlamatEkspedisiThreshold{}).Select("id").Where(&AlamatEkspedisiThreshold{
		IdAlamatEkspedisi: a.IdAlamatEkspedisi,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&AlamatEkspedisiThreshold{
		IdAlamatEkspedisi: a.IdAlamatEkspedisi,
	}).Error
}

func (a *AlamatEkspedisiThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&AlamatEkspedisiThreshold{}).Where(&AlamatEkspedisiThreshold{
		IdAlamatEkspedisi: a.IdAlamatEkspedisi,
	}).Updates(updates).Error
}

func (a *AlamatEkspedisiThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&AlamatEkspedisiThreshold{}).Where(&AlamatEkspedisiThreshold{
		IdAlamatEkspedisi: a.IdAlamatEkspedisi,
	}).Updates(updates).Error
}

func (a *AlamatEkspedisiThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&AlamatEkspedisiThreshold{}).Where(&AlamatEkspedisiThreshold{
		IdAlamatEkspedisi: a.IdAlamatEkspedisi,
	}).Updates(
		updates,
	).Error
}

func (a AlamatEkspedisiThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&AlamatEkspedisiThreshold{}).Where(&AlamatEkspedisiThreshold{
		IdAlamatEkspedisi: a.IdAlamatEkspedisi,
	}).Updates(
		updates,
	).Error
}
