package models

import "time"

type RekeningSeller struct {
	ID              int64      `gorm:"primaryKey;autoIncrement" json:"id_rekening_seller"`
	IDSeller        int32      `gorm:"column:id_seller;not null;index" json:"id_seller"`
	Seller          Seller     `gorm:"foreignKey:IDSeller;references:ID" json:"-"`
	NamaBank        string     `gorm:"column:nama_bank;type:varchar(50);not null" json:"nama_bank_rekening_seller"`
	NomorRekening   string     `gorm:"column:nomor_rekening;type:varchar(50);not null" json:"nomor_rekening_seller"`
	PemilikRekening string     `gorm:"column:pemilik_rekening;type:varchar(100);not null" json:"pemilik_rekening_seller"`
	IsDefault       bool       `gorm:"column:is_default;default:false" json:"is_default_rekening_seller"`
	CreatedAt       time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (RekeningSeller) TableName() string {
	return "rekening_seller"
}

type RekeningKurir struct {
	ID              int64      `gorm:"primaryKey;autoIncrement" json:"id_rekening_kurir"`
	IdKurir         int64      `gorm:"column:id_kurir;not null" json:"id_kurir_rekening_kurir"`
	Kurir           Kurir      `gorm:"foreignKey:IdKurir;references:ID" json:"-"`
	NamaBank        string     `gorm:"column:nama_bank;type:varchar(50);not null" json:"nama_bank_rekening_kurir"`
	NomorRekening   string     `gorm:"column:nomor_rekening;type:varchar(50);not null" json:"nomor_rekening_kurir"`
	PemilikRekening string     `gorm:"column:pemilik_rekening;type:varchar(30);not null" json:"pemilik_rekening_kurir"`
	CreatedAt       time.Time  `gorm:"autoCreateTime"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime"`
	DeletedAt       *time.Time `gorm:"index"`
}

func (RekeningKurir) TableName() string {
	return "rekening_kurir"
}
