package models

import (
	"time"

	"gorm.io/gorm"
)

type AlamatEkspedisi struct {
	ID              int64          `gorm:"primaryKey;autoIncrement" json:"id_alamat_ekspedisi"`
	Kota            string         `gorm:"index;column:kota;type:nama_kota" json:"kota"`
	NamaAlamat      string         `gorm:"column:nama_alamat;type:text" json:"nama"`
	Lokasi          string         `gorm:"column:lokasi;type:text;not null" json:"lokasi"`
	Longitude       float64        `gorm:"column:longitude;type:numeric(11,8);not null" json:"longitude"`
	Latitude        float64        `gorm:"column:latitude;type:numeric(11,8);not null" json:"latitude"`
	PengirimanCount int64          `gorm:"column:pengiriman_count;type:int8;not null;default:0" json:"pengiriman_count" `
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (AlamatEkspedisi) TableName() string {
	return "alamat_ekspedisi"
}
