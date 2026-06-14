package historical_models

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

)

type RekeningSeller struct {
	ID              int64
	IDSeller        int32
	Seller          Seller
	NamaBank        string
	NomorRekening   string
	PemilikRekening string
	IsDefault       bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
	Pencatatan
}

func (RekeningSeller) TableName() string {
	return "rekening_seller_historical"
}

func (r *RekeningSeller) CreateTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller int,
		nama_bank text,
		nomor_rekening text,
		pemilik_rekening text,
		is_default boolean,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, r.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", r.TableName())
	return nil
}

func (r *RekeningSeller) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":               r.ID,
		"id_seller":        r.IDSeller,
		"nama_bank":        r.NamaBank,
		"nomor_rekening":   r.NomorRekening,
		"pemilik_rekening": r.PemilikRekening,
		"is_default":       r.IsDefault,
		"created_at":       r.CreatedAt,
		"updated_at":       r.UpdatedAt,
		"deleted_at":       r.DeletedAt,
		"tahun_update":     r.TahunUpdate,
		"bulan_update":     r.BulanUpdate,
		"event_time":       r.EventTime,
	}
}

func (r *RekeningSeller) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, r.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", r.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", r.TableName())
	return nil
}

type RekeningKurir struct {
	ID              int64
	IdKurir         int64
	Kurir           Kurir
	NamaBank        string
	NomorRekening   string
	PemilikRekening string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
	Pencatatan
}

func (RekeningKurir) TableName() string {
	return "rekening_kurir_historical"
}

func (r *RekeningKurir) CreateTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kurir bigint,
		nama_bank text,
		nomor_rekening text,
		pemilik_rekening text,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, r.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", r.TableName())
	return nil
}

func (r *RekeningKurir) ParseToInsertType() map[string]interface{} {
	return map[string]interface{}{
		"id":               r.ID,
		"id_kurir":         r.IdKurir,
		"nama_bank":        r.NamaBank,
		"nomor_rekening":   r.NomorRekening,
		"pemilik_rekening": r.PemilikRekening,
		"created_at":       r.CreatedAt,
		"updated_at":       r.UpdatedAt,
		"deleted_at":       r.DeletedAt,
		"tahun_update":     r.TahunUpdate,
		"bulan_update":     r.BulanUpdate,
		"event_time":       r.EventTime,
	}
}

func (r *RekeningKurir) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, r.TableName())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", r.TableName(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", r.TableName())
	return nil
}
