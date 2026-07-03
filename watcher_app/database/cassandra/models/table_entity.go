package cass_models

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
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

func (p Pengguna) TableNameHistorical() string {
	return "pengguna_historical"
}

func (p *Pengguna) CreateHistoricalTable(ctx context.Context, s *gocql.Session) error {
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
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, p.TableNameHistorical())

	if err := s.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Println("Berhasil Eksekusi query membuat tabel pengguna_by_periode")
	return nil
}

func (p Pengguna) TableNameSotReplica() string {
	return "pengguna_sot_replica"
}

func (p *Pengguna) CreateSotReplicaTable(ctx context.Context, s *gocql.Session) error {
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
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, p.TableNameSotReplica())

	if err := s.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Println("Berhasil Eksekusi query membuat tabel sot_replica")
	return nil
}

func (p *Pengguna) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":              p.ID,
		"username":        p.Username,
		"nama":            p.Nama,
		"email":           p.Email,
		"password_hash":   p.PasswordHash,
		"pin_hash":        p.PinHash,
		"status_pengguna": p.StatusPengguna,
		"created_at":      p.CreatedAt,
		"updated_at":      p.UpdatedAt,
		"deleted_at":      p.DeletedAt,
	}
}

// DropTable disesuaikan menggunakan p.TableName() secara dinamis
func (p *Pengguna) DropTable(ctx context.Context, s *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, p.TableNameHistorical())

	if err := s.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", p.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", p.TableNameHistorical())
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
	Deskripsi        string
	StatusSeller     string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

func (s Seller) TableNameHistorical() string {
	return "seller_historical"
}

func (s *Seller) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, s.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", s.TableNameHistorical())
	return nil
}

func (s Seller) TableNameSotReplica() string {
	return "seller_sot_replica"
}

func (s *Seller) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct Pengguna dan Pencatatan
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
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, s.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Println("Berhasil Eksekusi query membuat tabel sot_replica")
	return nil
}

func (s *Seller) ParseToCUDType() map[string]interface{} {
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
		"deskrispi":         s.Deskripsi,
		"status_seller":     s.StatusSeller,
		"created_at":        s.CreatedAt,
		"updated_at":        s.UpdatedAt,
		"deleted_at":        s.DeletedAt,
	}
}

func (s *Seller) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, s.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", s.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", s.TableNameHistorical())
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
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

func (k Kurir) TableNameHistorical() string {
	return "kurir_historical"
}

func (k *Kurir) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
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
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, k.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", k.TableNameHistorical())
	return nil
}

func (k Kurir) TableNameSotReplica() string {
	return "kurir_sot_replica"
}

func (k *Kurir) CreateSotReplicaTable(ctx context.Context, s *gocql.Session) error {
	// Query CREATE TABLE disesuaikan dengan field di struct Pengguna dan Pencatatan
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
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, k.TableNameSotReplica())

	if err := s.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Println("Berhasil Eksekusi query membuat tabel sot_replica")
	return nil
}

func (k *Kurir) ParseToCUDType() map[string]interface{} {
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
		"updated_at":     k.UpdatedAt,
		"deleted_at":     k.DeletedAt,
	}
}

// DropTable disesuaikan menggunakan k.TableName() secara dinamis
func (k *Kurir) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, k.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", k.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", k.TableNameHistorical())
	return nil
}
