package cass_models

import (
	"context"
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

type PayOutKurir struct {
	ID               int64
	IdKurir          int64
	Kurir            Kurir
	IdPengiriman     int64 // 🔵 Ditambahkan sesuai sot_models
	Pengiriman       Pengiriman
	IdDisbursment    int64
	UserId           int
	Amount           int
	Status           string
	Reason           string
	Timestamp        string
	BankCode         string
	AccountNumber    string
	RecipientName    string
	SenderBank       string
	Remark           string
	Receipt          string
	TimeServed       string
	BundleId         int64
	CompanyId        int64
	RecipientCity    int
	CreatedFrom      string
	Direction        string
	Sender           string
	Fee              int
	BeneficiaryEmail string
	IdempotencyKey   string
	IsVirtualAccount bool
}

func (PayOutKurir) TableNameHistorical() string {
	return "pay_out_kurir_historical"
}

func (PayOutKurir) TableNameSotReplica() string {
	return "pay_out_kurir"
}

func (p *PayOutKurir) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kurir bigint,
		id_pengiriman bigint,
		id_disbursment bigint,
		user_id int,
		amount int,
		status text,
		reason text,
		timestamp text,
		bank_code text,
		account_number text,
		recipient_name text,
		sender_bank text,
		remark text,
		receipt text,
		time_served text,
		bundle_id bigint,
		company_id bigint,
		recipient_city int,
		created_from text,
		direction text,
		sender text,
		fee int,
		beneficiary_email text,
		idempotency_key text,
		is_virtual_account boolean,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, p.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", p.TableNameHistorical())
	return nil
}

func (p *PayOutKurir) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                 p.ID,
		"id_kurir":           p.IdKurir,
		"id_pengiriman":      p.IdPengiriman,
		"id_disbursment":     p.IdDisbursment,
		"user_id":            p.UserId,
		"amount":             p.Amount,
		"status":             p.Status,
		"reason":             p.Reason,
		"timestamp":          p.Timestamp,
		"bank_code":          p.BankCode,
		"account_number":     p.AccountNumber,
		"recipient_name":     p.RecipientName,
		"sender_bank":        p.SenderBank,
		"remark":             p.Remark,
		"receipt":            p.Receipt,
		"time_served":        p.TimeServed,
		"bundle_id":          p.BundleId,
		"company_id":         p.CompanyId,
		"recipient_city":     p.RecipientCity,
		"created_from":       p.CreatedFrom,
		"direction":          p.Direction,
		"sender":             p.Sender,
		"fee":                p.Fee,
		"beneficiary_email":  p.BeneficiaryEmail,
		"idempotency_key":    p.IdempotencyKey,
		"is_virtual_account": p.IsVirtualAccount,
	}
}

func (p *PayOutKurir) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, p.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", p.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", p.TableNameHistorical())
	return nil
}

func (p *PayOutKurir) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_kurir bigint,
		id_pengiriman bigint,
		id_disbursment bigint,
		user_id int,
		amount int,
		status text,
		reason text,
		timestamp text,
		bank_code text,
		account_number text,
		recipient_name text,
		sender_bank text,
		remark text,
		receipt text,
		time_served text,
		bundle_id bigint,
		company_id bigint,
		recipient_city int,
		created_from text,
		direction text,
		sender text,
		fee int,
		beneficiary_email text,
		idempotency_key text,
		is_virtual_account boolean,
		PRIMARY KEY (id)
	)`, p.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", p.TableNameSotReplica())
	return nil
}

type PayOutSeller struct {
	ID               int64
	IdSeller         int64
	Seller           Seller
	IdTransaksi      int64 // 🔵 Ditambahkan sesuai sot_models
	Transaksi        Transaksi
	IdDisbursment    int64
	UserId           int
	Amount           int
	Status           string
	Reason           string
	Timestamp        string
	BankCode         string
	AccountNumber    string
	RecipientName    string
	SenderBank       string
	Remark           string
	Receipt          string
	TimeServed       string
	BundleId         int64
	CompanyId        int64
	RecipientCity    int
	CreatedFrom      string
	Direction        string
	Sender           string
	Fee              int
	BeneficiaryEmail string
	IdempotencyKey   string
	IsVirtualAccount bool
}

func (PayOutSeller) TableNameHistorical() string {
	return "payout_seller_historical"
}

func (PayOutSeller) TableNameSotReplica() string {
	return "payout_seller"
}

func (p *PayOutSeller) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller bigint,
		id_transaksi bigint,
		id_disbursment bigint,
		user_id int,
		amount int,
		status text,
		reason text,
		timestamp text,
		bank_code text,
		account_number text,
		recipient_name text,
		sender_bank text,
		remark text,
		receipt text,
		time_served text,
		bundle_id bigint,
		company_id bigint,
		recipient_city int,
		created_from text,
		direction text,
		sender text,
		fee int,
		beneficiary_email text,
		idempotency_key text,
		is_virtual_account boolean,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, p.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", p.TableNameHistorical())
	return nil
}

func (p *PayOutSeller) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                 p.ID,
		"id_seller":          p.IdSeller,
		"id_transaksi":       p.IdTransaksi,
		"id_disbursment":     p.IdDisbursment,
		"user_id":            p.UserId,
		"amount":             p.Amount,
		"status":             p.Status,
		"reason":             p.Reason,
		"timestamp":          p.Timestamp,
		"bank_code":          p.BankCode,
		"account_number":     p.AccountNumber,
		"recipient_name":     p.RecipientName,
		"sender_bank":        p.SenderBank,
		"remark":             p.Remark,
		"receipt":            p.Receipt,
		"time_served":        p.TimeServed,
		"bundle_id":          p.BundleId,
		"company_id":         p.CompanyId,
		"recipient_city":     p.RecipientCity,
		"created_from":       p.CreatedFrom,
		"direction":          p.Direction,
		"sender":             p.Sender,
		"fee":                p.Fee,
		"beneficiary_email":  p.BeneficiaryEmail,
		"idempotency_key":    p.IdempotencyKey,
		"is_virtual_account": p.IsVirtualAccount,
	}
}

func (p *PayOutSeller) DropTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, p.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", p.TableNameHistorical(), err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", p.TableNameHistorical())
	return nil
}

func (p *PayOutSeller) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_seller bigint,
		id_transaksi bigint,
		id_disbursment bigint,
		user_id int,
		amount int,
		status text,
		reason text,
		timestamp text,
		bank_code text,
		account_number text,
		recipient_name text,
		sender_bank text,
		remark text,
		receipt text,
		time_served text,
		bundle_id bigint,
		company_id bigint,
		recipient_city int,
		created_from text,
		direction text,
		sender text,
		fee int,
		beneficiary_email text,
		idempotency_key text,
		is_virtual_account boolean,
		PRIMARY KEY (id)
	)`, p.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}

	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", p.TableNameSotReplica())
	return nil
}
