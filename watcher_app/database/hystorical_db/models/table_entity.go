package historical_models

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

type Pengguna struct {
	ID             int64
	Username       string
	Nama           string
	Email          string
	PasswordHash   string
	PinHash        string
	StatusPengguna string
	CreatedAt      time.Time
	Pencatatan
}

func (p Pengguna) TableName() string {
	return "pengguna_historical"
}

func (p *Pengguna) CreateTable(ctx context.Context, s *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct Pengguna dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		username text,
		nama text,
		email text,
		password_hash text,
		pin_hash text,
		status_pengguna text,
		created_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, p.TableName())

	if err := s.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Println("Berhasil Eksekusi query membuat tabel pengguna_by_periode")
	return nil
}

func (p *Pengguna) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":              p.ID,
		"username":        p.Username,
		"nama":            p.Nama,
		"email":           p.Email,
		"password_hash":   p.PasswordHash,
		"pin_hash":        p.PinHash,
		"status_pengguna": p.StatusPengguna,
		"created_at":      p.CreatedAt,
		"tahun_update":    p.TahunUpdate,
		"bulan_update":    p.BulanUpdate,
		"event_time":      p.EventTime,
	}
}

// DropTable disesuaikan menggunakan p.TableName() secara dinamis
func (p *Pengguna) DropTable(ctx context.Context, s *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, p.TableName())

	if err := s.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", p.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", p.TableName())
	return nil
}

type Seller struct {
	ID               int
	Username         string
	Nama             string
	Email            string
	Jenis            string
	SellerDedication string
	JamOperasional   string
	Punchline        string
	Password         string
	Deskrisi         string
	StatusSeller     string
	CreatedAt        time.Time
	Pencatatan
}

func (s Seller) TableName() string {
	return "seller_historical"
}

func (s *Seller) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct Seller dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		username text,
		nama text,
		email text,
		jenis text,
		seller_dedication text,
		jam_operasional text,
		punchline text,
		password text,
		deskrisi text,
		status_seller text,
		created_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, s.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", s.TableName())
	return nil
}

func (s *Seller) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":                s.ID,
		"username":          s.Username,
		"nama":              s.Nama,
		"email":             s.Email,
		"jenis":             s.Jenis,
		"seller_dedication": s.SellerDedication,
		"jam_operasional":   s.JamOperasional,
		"punchline":         s.Punchline,
		"password":          s.Password,
		"deskrisi":          s.Deskrisi,
		"status_seller":     s.StatusSeller,
		"created_at":        s.CreatedAt,
		"tahun_update":      s.TahunUpdate,
		"bulan_update":      s.BulanUpdate,
		"event_time":        s.EventTime,
	}
}

func (s *Seller) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, s.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", s.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", s.TableName())
	return nil
}

type Kurir struct {
	ID            int64
	Nama          string
	Username      string
	Email         string
	Jenis         string
	PasswordHash  string
	Deskripsi     string
	StatusKurir   string
	StatusBid     string
	VerifiedKurir bool
	Rating        float32
	TipeKendaraan string
	CreatedAt     time.Time
	Pencatatan
}

func (k Kurir) TableName() string {
	return "kurir_historical"
}

func (k *Kurir) CreateTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct Kurir dan Pencatatan
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		nama text,
		username text,
		email text,
		jenis text,
		password_hash text,
		deskripsi text,
		status_kurir text,
		status_bid text,
		verified_kurir boolean,
		rating float,
		tipe_kendaraan text,
		created_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, k.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", k.TableName())
	return nil
}

func (k *Kurir) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":             k.ID,
		"nama":           k.Nama,
		"username":       k.Username,
		"email":          k.Email,
		"jenis":          k.Jenis,
		"password_hash":  k.PasswordHash,
		"deskripsi":      k.Deskripsi,
		"status_kurir":   k.StatusKurir,
		"status_bid":     k.StatusBid,
		"verified_kurir": k.VerifiedKurir,
		"rating":         k.Rating,
		"tipe_kendaraan": k.TipeKendaraan,
		"created_at":     k.CreatedAt,
		"tahun_update":   k.TahunUpdate,
		"bulan_update":   k.BulanUpdate,
		"event_time":     k.EventTime,
	}
}

// DropTable disesuaikan menggunakan k.TableName() secara dinamis
func (k *Kurir) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, k.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", k.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", k.TableName())
	return nil
}
