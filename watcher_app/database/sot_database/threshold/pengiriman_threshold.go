package sot_threshold

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	sot_threshold_seeders_nama "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/threshold/seeders/nama_threshold"
)

type PengirimanEkspedisiThreshold struct {
	ID                    int64                          `gorm:"primaryKey;autoIncrement" json:"id_pengiriman_ekspedisi_threshold"`
	IdPengirimanEkspedisi int64                          `gorm:"column:id_pengiriman_ekspedisi;index;not null" json:"id_pengiriman_ekspedisi"`
	PengirimanEkspedisi   sot_models.PengirimanEkspedisi `gorm:"foreignKey:IdPengirimanEkspedisi;references:ID" json:"-"`

	BidKurirEksScheduler                    int32 `gorm:"column:bid_kurir_eks_scheduler;type:int4;default:0" json:"bid_kurir_eks_scheduler"`
	MediaPengirimanEkspedisiPickedUpFoto    int32 `gorm:"column:media_pengiriman_ekspedisi_picked_up_foto;type:int4;default:0" json:"media_pengiriman_ekspedisi_picked_up_foto"`
	MediaPengirimanEkspedisiSampaiAgentFoto int32 `gorm:"column:media_pengiriman_ekspedisi_sampai_agent_foto;type:int4;default:0" json:"media_pengiriman_ekspedisi_sampai_agent_foto"`
}

func (PengirimanEkspedisiThreshold) TableName() string {
	return sot_threshold_seeders_nama.PengirimanEkspedisiThreshold
}

func (p *PengirimanEkspedisiThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&PengirimanEkspedisiThreshold{}).Select("id").Where(&PengirimanEkspedisiThreshold{
		IdPengirimanEkspedisi: p.IdPengirimanEkspedisi,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&PengirimanEkspedisiThreshold{
		IdPengirimanEkspedisi: p.IdPengirimanEkspedisi,
	}).Error
}

func (p *PengirimanEkspedisiThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&PengirimanEkspedisiThreshold{}).Where(&PengirimanEkspedisiThreshold{
		IdPengirimanEkspedisi: p.IdPengirimanEkspedisi,
	}).Updates(updates).Error
}

func (p *PengirimanEkspedisiThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&PengirimanEkspedisiThreshold{}).Where(&PengirimanEkspedisiThreshold{
		IdPengirimanEkspedisi: p.IdPengirimanEkspedisi,
	}).Updates(updates).Error
}

func (p *PengirimanEkspedisiThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&PengirimanEkspedisiThreshold{}).Where(&PengirimanEkspedisiThreshold{
		IdPengirimanEkspedisi: p.IdPengirimanEkspedisi,
	}).Updates(
		updates,
	).Error
}

func (p *PengirimanEkspedisiThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&PengirimanEkspedisiThreshold{}).Where(&PengirimanEkspedisiThreshold{
		IdPengirimanEkspedisi: p.IdPengirimanEkspedisi,
	}).Updates(
		updates,
	).Error
}

type PengirimanNonEkspedisiThreshold struct {
	ID           int64                 `gorm:"primaryKey;autoIncrement" json:"id_pengiriman_non_ekspedisi_threshold"`
	IdPengiriman int64                 `gorm:"column:id_pengiriman;index;not null" json:"id_pengiriman"`
	Pengiriman   sot_models.Pengiriman `gorm:"foreignKey:IdPengiriman;references:ID" json:"-"`

	BidKurirNonEksScheduler     int32 `gorm:"column:bid_kurir_non_eks_scheduler;type:int4;default:0" json:"bid_kurir_non_eks_scheduler"`
	MediaPengirimanPickedUpFoto int32 `gorm:"column:media_pengiriman_picked_up_foto;type:int4;default:0" json:"media_pengiriman_picked_up_foto"`
	MediaPengirimanSampaiFoto   int32 `gorm:"column:media_pengiriman_sampai_foto;type:int4;default:0" json:"media_pengiriman_sampai_foto"`
}

func (PengirimanNonEkspedisiThreshold) TableName() string {
	return sot_threshold_seeders_nama.PengirimanNonEkspedisiThreshold
}

func (p *PengirimanNonEkspedisiThreshold) Inisialisasi(ctx context.Context, db *gorm.DB) error {
	var id_data_threshold int64 = 0
	if err := db.WithContext(ctx).Model(&PengirimanNonEkspedisiThreshold{}).Select("id").Where(&PengirimanNonEkspedisiThreshold{
		IdPengiriman: p.IdPengiriman,
	}).Limit(1).Scan(&id_data_threshold).Error; err != nil {
		return err
	}

	if id_data_threshold != 0 {
		return fmt.Errorf("gagal sudah memiliki threshold")
	}

	return db.WithContext(ctx).Create(&PengirimanNonEkspedisiThreshold{
		IdPengiriman: p.IdPengiriman,
	}).Error
}

func (p *PengirimanNonEkspedisiThreshold) Increment(ctx context.Context, db *gorm.DB, koloms ...string) error {

	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s + ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&PengirimanNonEkspedisiThreshold{}).Where(&PengirimanNonEkspedisiThreshold{
		IdPengiriman: p.IdPengiriman,
	}).Updates(updates).Error
}

func (p *PengirimanNonEkspedisiThreshold) Decrement(ctx context.Context, db *gorm.DB, koloms ...string) error {
	if len(koloms) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(koloms))
	for _, kolom := range koloms {
		updates[kolom] = gorm.Expr(fmt.Sprintf("%s - ?", kolom), 1)
	}

	return db.WithContext(ctx).Model(&PengirimanNonEkspedisiThreshold{}).Where(&PengirimanNonEkspedisiThreshold{
		IdPengiriman: p.IdPengiriman,
	}).Updates(updates).Error
}

func (p *PengirimanNonEkspedisiThreshold) CustomIncrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s + ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&PengirimanNonEkspedisiThreshold{}).Where(&PengirimanNonEkspedisiThreshold{
		IdPengiriman: p.IdPengiriman,
	}).Updates(
		updates,
	).Error
}

func (p *PengirimanNonEkspedisiThreshold) CustomDecrement(ctx context.Context, db *gorm.DB, kj []CustomCounter) error {
	if len(kj) == 0 {
		return fmt.Errorf("gagal data kosong")
	}

	updates := make(map[string]interface{}, len(kj))
	for _, kolom := range kj {
		updates[kolom.FieldName] = gorm.Expr(fmt.Sprintf("%s - ?", kolom.FieldName), kolom.Count)
	}

	return db.WithContext(ctx).Model(&PengirimanNonEkspedisiThreshold{}).Where(&PengirimanNonEkspedisiThreshold{
		IdPengiriman: p.IdPengiriman,
	}).Updates(
		updates,
	).Error
}
