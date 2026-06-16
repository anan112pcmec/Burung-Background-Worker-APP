package sot_threshold

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	sot_threshold_seeders_nama "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/threshold/seeders/nama_threshold"
)

type TransaksiThreshold struct {
	ID          int64                `gorm:"primaryKey;autoIncrement" json:"id_transaksi_threshold"`
	IdTransaksi int64                `gorm:"column:id_transaksi;index;not null" json:"id_transaksi"`
	Transaksi   sot_models.Transaksi `gorm:"foreignKey:IdTransaksi;references:ID" json:"-"`

	MediaTransaksiApprovedFoto  int32 `gorm:"column:media_transaksi_approved_foto;type:int4;default:0" json:"media_transaksi_approved_foto"`
	MediaTransaksiApprovedVideo int32 `gorm:"column:media_transaksi_approved_video;type:int4;default:0" json:"media_transaksi_approved_video"`
	Pengiriman                  int32 `gorm:"column:pengiriman;type:int4;default:0" json:"pengiriman"`
	PengirimanEkspedisi         int32 `gorm:"column:pengiriman_ekspedisi;type:int4;default:0" json:"pengiriman_ekspedisi"`
}

func (TransaksiThreshold) TableName() string {
	return sot_threshold_seeders_nama.TransaksiThreshold
}

func (t TransaksiThreshold) Inisialisasi(id_fk int64, ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&TransaksiThreshold{}).Select("id").Where(&TransaksiThreshold{
		IdTransaksi: id_fk,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&TransaksiThreshold{
		IdTransaksi: id_fk,
	}).Error
}

func (t *TransaksiThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&TransaksiThreshold{}).Where(&TransaksiThreshold{
		IdTransaksi: t.IdTransaksi,
	}).Updates(updates).Error
}

func (t *TransaksiThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&TransaksiThreshold{}).Where(&TransaksiThreshold{
		IdTransaksi: t.IdTransaksi,
	}).Updates(updates).Error
}

func (t *TransaksiThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&TransaksiThreshold{}).Where(&TransaksiThreshold{
		IdTransaksi: t.IdTransaksi,
	}).Updates(
		updates,
	).Error
}

func (t *TransaksiThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&TransaksiThreshold{}).Where(&TransaksiThreshold{
		IdTransaksi: t.IdTransaksi,
	}).Updates(
		updates,
	).Error
}

type ThresholdOrderSeller struct {
	ID         int64             `gorm:"primaryKey;autoIncrement" json:"id_threshold_order_seller"`
	IdSeller   int32             `gorm:"column:id_seller;not null" json:"id_seller_threshold_order_seller"`
	Seller     sot_models.Seller `gorm:"foreignKey:IdSeller;references:ID" json:"-"`
	Dibatalkan int64             `gorm:"column:dibatalkan;type:int8;not null;default:0" json:"dibatalkan_threshold_order_seller"`
	Dibayar    int64             `gorm:"column:dibayar;type:int8;not null;default:0" json:"dibayar_threshold_order_seller"`
	Diproses   int64             `gorm:"column:diproses;type:int8;not null;default:0" json:"diproses_threshold_order_seller"`
	Waiting    int64             `gorm:"column:waiting;type:int8;not null;default:0" json:"waiting_threshold_order_seller"`
	Dikirim    int64             `gorm:"column:dikirim;type:int8;not null;default:0" json:"dikirim_threshold_order_seller"`
	Sampai     int64             `gorm:"column:sampai;type:int8;not null;default:0" json:"sampai_threshold_order_seller"`
	Total      int64             `gorm:"column:total;type:int8;not null;default:0" json:"total_threshold_order_seller"`
}

type PembayaranThreshold struct {
	ID           int64                 `gorm:"primaryKey;autoIncrement" json:"id_pembayaran_threshold"`
	IdPembayaran int64                 `gorm:"column:id_pembayaran;index;not null" json:"id_pembayaran"`
	Pembayaran   sot_models.Pembayaran `gorm:"foreignKey:IdPembayaran;references:ID" json:"-"`

	Transaksi int32 `gorm:"column:transaksi;type:int4;default:0" json:"transaksi"`
}

func (PembayaranThreshold) TableName() string {
	return sot_threshold_seeders_nama.PembayaranThreshold
}

func (p PembayaranThreshold) Inisialisasi(id_fk int64, ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&PembayaranThreshold{}).Select("id").Where(&PembayaranThreshold{
		IdPembayaran: id_fk,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&PembayaranThreshold{
		IdPembayaran: id_fk,
	}).Error
}

func (p *PembayaranThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&PembayaranThreshold{}).Where(&PembayaranThreshold{
		IdPembayaran: p.IdPembayaran,
	}).Updates(updates).Error
}

func (p *PembayaranThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&PembayaranThreshold{}).Where(&PembayaranThreshold{
		IdPembayaran: p.IdPembayaran,
	}).Updates(updates).Error
}

func (p *PembayaranThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&PembayaranThreshold{}).Where(&PembayaranThreshold{
		IdPembayaran: p.IdPembayaran,
	}).Updates(
		updates,
	).Error
}

func (p *PembayaranThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&PembayaranThreshold{}).Where(&PembayaranThreshold{
		IdPembayaran: p.IdPembayaran,
	}).Updates(
		updates,
	).Error
}

type TransaksiFailedThreshold struct {
	ID                int64                      `gorm:"primaryKey;autoIncrement" json:"id_transaksi_failed_threshold"`
	IdTransaksiFailed int64                      `gorm:"column:id_transaksi_failed;index;not null" json:"id_transaksi_failed"`
	TransaksiFailed   sot_models.TransaksiFailed `gorm:"foreignKey:IdTransaksiFailed;references:ID" json:"-"`
}
