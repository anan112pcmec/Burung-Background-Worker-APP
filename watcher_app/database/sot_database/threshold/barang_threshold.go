package sot_threshold

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	sot_threshold_seeders_nama "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/threshold/seeders/nama_threshold"
)

type BarangIndukThreshold struct {
	ID            int64                  `gorm:"primaryKey;autoIncrement" json:"id_barang_induk_threshold"`
	IdBarangInduk int64                  `gorm:"column:id_barang_induk;index;not null" json:"id_barang_induk"`
	BarangInduk   sot_models.BarangInduk `gorm:"foreignKey:IdBarangInduk;references:ID" json:"-"`

	BarangDiDiskon          int32 `gorm:"column:barang_di_diskon;type:int4;default:0" json:"barang_di_diskon"`
	BarangDisukai           int32 `gorm:"column:barang_disukai;type:int4;default:0" json:"barang_disukai"`
	BarangKeEtalase         int32 `gorm:"column:barang_ke_etalase;type:int4;default:0" json:"barang_ke_etalase"`
	KategoriBarang          int32 `gorm:"column:kategori_barang;type:int4;default:0" json:"kategori_barang"`
	Keranjang               int32 `gorm:"column:keranjang;type:int4;default:0" json:"keranjang"`
	Komentar                int32 `gorm:"column:komentar;type:int4;default:0" json:"komentar"`
	MediaBarangIndukFoto    int32 `gorm:"column:media_barang_induk_foto;type:int4;default:0" json:"media_barang_induk_foto"`
	MediaBarangIndukVideo   int32 `gorm:"column:media_barang_induk_video;type:int4;default:0" json:"media_barang_induk_video"`
	MediaKategoriBarangFoto int32 `gorm:"column:media_kategori_barang_foto;type:int4;default:0" json:"media_kategori_barang_foto"`
	Review                  int32 `gorm:"column:review;type:int4;default:0" json:"review"`
	Transaksi               int32 `gorm:"column:transaksi;type:int4;default:0" json:"transaksi"`
	VarianBarang            int32 `gorm:"column:varian_barang;type:int4;default:0" json:"varian_barang"`
	Wishlist                int32 `gorm:"column:wishlist;type:int4;default:0" json:"wishlist"`
}

func (BarangIndukThreshold) TableName() string {
	return sot_threshold_seeders_nama.BarangIndukThreshold
}

func (b *BarangIndukThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&BarangIndukThreshold{}).Select("id").Where(&BarangIndukThreshold{
		IdBarangInduk: b.IdBarangInduk,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&BarangIndukThreshold{
		IdBarangInduk: b.IdBarangInduk,
	}).Error
}

func (b *BarangIndukThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&BarangIndukThreshold{}).Where(&BarangIndukThreshold{
		ID: b.IdBarangInduk,
	}).Updates(updates).Error
}

func (b *BarangIndukThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&BarangIndukThreshold{}).Where(&BarangIndukThreshold{
		ID: b.IdBarangInduk,
	}).Updates(updates).Error
}

func (b *BarangIndukThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&BarangIndukThreshold{}).Where(&BarangIndukThreshold{
		ID: b.IdBarangInduk,
	}).Updates(
		updates,
	).Error
}

func (b *BarangIndukThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&BarangIndukThreshold{}).Where(&BarangIndukThreshold{
		ID: b.IdBarangInduk,
	}).Updates(
		updates,
	).Error
}

func (b *BarangIndukThreshold) GetKolomCount(ctx context.Context, db *gorm.DB, kolom string) (error, int) {
	var counter int = 0
	if err := db.WithContext(ctx).Model(&BarangIndukThreshold{}).Select(kolom).Where(&BarangIndukThreshold{
		IdBarangInduk: b.IdBarangInduk,
	}).Limit(1).Take(&counter).Error; err != nil {
		return fmt.Errorf("gagal mengambil count"), counter
	}

	return nil, counter
}

type KategoriBarangThreshold struct {
	ID               int64                     `gorm:"primaryKey;autoIncrement" json:"id_kategori_barang_threshold"`
	IdKategoriBarang int64                     `gorm:"column:id_kategori_barang;index;not null" json:"id_kategori_barang"`
	KategoriBarang   sot_models.KategoriBarang `gorm:"foreignKey:IdKategoriBarang;references:ID" json:"-"`

	// Counter
	BarangDiDiskon          int32 `gorm:"column:barang_di_diskon;type:int4;default:0" json:"barang_di_diskon"`
	Keranjang               int32 `gorm:"column:keranjang;type:int4;default:0" json:"keranjang"`
	MediaKategoriBarangFoto int32 `gorm:"column:media_kategori_barang_foto;type:int4;default:0" json:"media_kategori_barang_foto"`
	Transaksi               int32 `gorm:"column:transaksi;type:int4;default:0" json:"transaksi"`
	VarianBarang            int32 `gorm:"column:varian_barang;type:int4;default:0" json:"varian_barang"`
}

func (KategoriBarangThreshold) TableName() string {
	return sot_threshold_seeders_nama.KategoriBarangThreshold
}

func (k *KategoriBarangThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&KategoriBarangThreshold{}).Select("id").Where(&KategoriBarangThreshold{
		IdKategoriBarang: k.IdKategoriBarang,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&KategoriBarangThreshold{
		IdKategoriBarang: k.IdKategoriBarang,
	}).Error
}

func (k *KategoriBarangThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&KategoriBarangThreshold{}).Where(&KategoriBarangThreshold{
		IdKategoriBarang: k.IdKategoriBarang,
	}).Updates(updates).Error
}

func (k *KategoriBarangThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&KategoriBarangThreshold{}).Where(&KategoriBarangThreshold{
		IdKategoriBarang: k.IdKategoriBarang,
	}).Updates(updates).Error
}

func (k *KategoriBarangThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&KategoriBarangThreshold{}).Where(&KategoriBarangThreshold{
		IdKategoriBarang: k.IdKategoriBarang,
	}).Updates(
		updates,
	).Error
}

func (k *KategoriBarangThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&KategoriBarangThreshold{}).Where(&KategoriBarangThreshold{
		IdKategoriBarang: k.IdKategoriBarang,
	}).Updates(
		updates,
	).Error
}

func (k *KategoriBarangThreshold) GetKolomCount(ctx context.Context, db *gorm.DB, kolom string) (error, int) {
	var counter int = 0
	if err := db.WithContext(ctx).Model(&KategoriBarangThreshold{}).Select(kolom).Where(&KategoriBarangThreshold{
		IdKategoriBarang: k.IdKategoriBarang,
	}).Limit(1).Take(&counter).Error; err != nil {
		return fmt.Errorf("gagal mengambil count"), counter
	}

	return nil, counter
}
