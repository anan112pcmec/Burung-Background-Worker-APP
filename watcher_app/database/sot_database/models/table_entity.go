package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Pengguna struct {
	ID             int64          `gorm:"primaryKey;autoIncrement" json:"id_user"`
	Username       string         `gorm:"column:username;type:varchar(100);not null;default:''" json:"username"`
	Nama           string         `gorm:"column:nama;type:text;not null;default:''" json:"nama"`
	Email          string         `gorm:"column:email;type:varchar(100);not null;uniqueIndex" json:"email"`
	PasswordHash   string         `gorm:"column:password_hash;type:varchar(250);not null;default:''" json:"password_hash"`
	PinHash        string         `gorm:"column:pin_hash;type:varchar(250);not null;default:''" json:"pin_hash"`
	StatusPengguna string         `gorm:"column:status;type:status;not null;default:'Offline'" json:"status"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Pengguna) TableName() string {
	return "pengguna"
}

type Seller struct {
	ID               int32          `gorm:"primaryKey;autoIncrement" json:"id_seller"`
	Username         string         `gorm:"column:username;type:varchar(100);notnull;default:''" json:"username"`
	Nama             string         `gorm:"column:nama;type:varchar(150);not null;default:''" json:"nama"`
	Email            string         `gorm:"column:email;type:varchar(150);not null;default:''" json:"email"`
	Jenis            string         `gorm:"column:jenis;type:jenis_seller;not null;default:'Personal'" json:"jenis"`
	SellerDedication string         `gorm:"column:seller_dedication;type:seller_dedication;not null;default:'Semua Barang'" json:"seller_dedication"`
	JamOperasional   string         `gorm:"column:jam_operasional;type:text;not null;default:''" json:"jam_operasional"`
	Punchline        string         `gorm:"column:punchline;type:text;not null;default:''" json:"punchline"`
	Password         string         `gorm:"column:password_hash;type:varchar(250);not null;default:''" json:"password_hash"`
	Deskripsi        string         `gorm:"column:deskripsi;type:text;not null;default:''" json:"deskripsi"`
	StatusSeller     string         `gorm:"column:status;type:status;not null;default:'Offline'" json:"status"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (s *Seller) Validating() error {
	if s.ID == 0 {
		return fmt.Errorf("id tidak valid")
	}
	if s.Username == "" {
		return fmt.Errorf("username tidak valid")
	}
	if s.Email == "" {
		return fmt.Errorf("email tidak valid")
	}
	return nil
}

func (Seller) TableName() string {
	return "seller"
}

type JenisLayananKurir string

type Kurir struct {
	ID            int64          `gorm:"primaryKey;autoIncrement" json:"id_kurir"`
	Nama          string         `gorm:"column:nama;type:varchar(150);not null;default:''" json:"nama"`
	Username      string         `gorm:"column:username;type:text;not null" json:"username"`
	Email         string         `gorm:"column:email;type:varchar(150);not null;default:''" json:"email"`
	Jenis         string         `gorm:"column:jenis;type:jenis_layanan_kurir;not null;default:'Reguler'" json:"jenis"`
	PasswordHash  string         `gorm:"column:password_hash;type:varchar(250);not null;default:''" json:"password_hash"`
	Deskripsi     string         `gorm:"column:deskripsi;type:text;not null;default:''" json:"deskripsi"`
	StatusKurir   string         `gorm:"column:status;type:status;not null;default:'Offline'" json:"status"`
	StatusBid     string         `gorm:"column:status_bid;type:status_kurir;not null;default:'Off'" json:"status_bid"`
	VerifiedKurir bool           `gorm:"column:verified;type:boolean;not null;default:false" json:"verified"`
	Rating        float32        `gorm:"column:rating;type:float;default:0" json:"rating"`
	TipeKendaraan string         `gorm:"column:jenis_kendaraan;type:jenis_kendaraan_kurir;default:'Unknown'" json:"tipe_kendaraan"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Kurir) TableName() string {
	return "kurir"
}
