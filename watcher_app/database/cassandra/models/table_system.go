package cass_models

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"gorm.io/gorm"
)

// =========================================================================
// ALAMAT EKSPEDISI
// =========================================================================

type AlamatEkspedisi struct {
	ID              int64
	Kota            string
	NamaAlamat      string
	Lokasi          string
	Longitude       float64
	Latitude        float64
	PengirimanCount int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

func (AlamatEkspedisi) TableNameHistorical() string {
	return "alamat_ekspedisi_historical"
}

func (AlamatEkspedisi) TableNameSotReplica() string {
	return "alamat_ekspedisi"
}

func (a AlamatEkspedisi) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		kota text,
		nama_alamat text,
		lokasi text,
		longitude double,
		latitude double,
		pengiriman_count bigint,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, a.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}
	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", a.TableNameHistorical())
	return nil
}

func (a AlamatEkspedisi) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":               a.ID,
		"kota":             a.Kota,
		"nama_alamat":      a.NamaAlamat,
		"lokasi":           a.Lokasi,
		"longitude":        a.Longitude,
		"latitude":         a.Latitude,
		"pengiriman_count": a.PengirimanCount,
		"created_at":       a.CreatedAt,
		"updated_at":       a.UpdatedAt,
		"deleted_at":       a.DeletedAt,
	}
}

func (a AlamatEkspedisi) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		kota text,
		nama_alamat text,
		lokasi text,
		longitude double,
		latitude double,
		pengiriman_count bigint,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, a.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}
	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", a.TableNameSotReplica())
	return nil
}

// =========================================================================
// REKENING SISTEM
// =========================================================================

type RekeningSistem struct {
	ID              int64
	NamaBank        string
	NomorRekening   string
	PemilikRekening string
	PusatKota       string
	CurrentActive   bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

func (RekeningSistem) TableNameHistorical() string {
	return "rekening_sistem_historical"
}

func (RekeningSistem) TableNameSotReplica() string {
	return "rekening_sistem"
}

func (r RekeningSistem) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		nama_bank text,
		nomor_rekening text,
		pemilik_rekening text,
		pusat_kota text,
		current_active boolean,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		tahun_update int,
		bulan_update int,
		event_time timestamp,
		PRIMARY KEY ((id, tahun_update, bulan_update), event_time)
	)`, r.TableNameHistorical())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}
	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", r.TableNameHistorical())
	return nil
}

func (r RekeningSistem) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":               r.ID,
		"nama_bank":        r.NamaBank,
		"nomor_rekening":   r.NomorRekening,
		"pemilik_rekening": r.PemilikRekening,
		"pusat_kota":       r.PusatKota,
		"current_active":   r.CurrentActive,
		"created_at":       r.CreatedAt,
		"updated_at":       r.UpdatedAt,
		"deleted_at":       r.DeletedAt,
	}
}

func (r RekeningSistem) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		nama_bank text,
		nomor_rekening text,
		pemilik_rekening text,
		pusat_kota text,
		current_active boolean,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, r.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}
	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", r.TableNameSotReplica())
	return nil
}

// =========================================================================
// PAYOUT SISTEM
// =========================================================================

type PayOutSistem struct {
	ID               int64
	IdDisburstment   int64
	IdTransaksi      int64
	Transaksi        Transaksi
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
	DeletedAt        gorm.DeletedAt
}

func (PayOutSistem) TableNameHistorical() string {
	return "payout_sistem_historical"
}

func (PayOutSistem) TableNameSotReplica() string {
	return "payout_sistem"
}

func (p PayOutSistem) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_disburstment bigint,
		id_transaksi bigint,
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
		deleted_at timestamp,
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

func (p PayOutSistem) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                 p.ID,
		"id_disburstment":    p.IdDisburstment,
		"id_transaksi":       p.IdTransaksi,
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
		"deleted_at":         p.DeletedAt,
	}
}

func (p PayOutSistem) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		id_disburstment bigint,
		id_transaksi bigint,
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
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, p.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}
	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", p.TableNameSotReplica())
	return nil
}

// =========================================================================
// KEBIJAKAN SISTEM
// =========================================================================

type KebijakanSistem struct {
	ID                            int64
	DitetapkanOleh                string
	IDAdmin                       int64
	NamaAdmin                     string
	KomisiSistemPerTransaksi      float32
	LimitMembuatDiskonPersonal    int32
	LimitMembuatDiskonDistributor int32
	LimitMembuatDiskonBrand       int32
	MaxJarakKmReguler             int
	MaxJarakKmExpress             int
	MaxJarakKmInstant             int
	EstimasiHariReguler           int
	EstimasiHariExpress           int
	EstimasiHariInstant           int
	TarifPengirimanRegulerPerKm   int
	TarifPengirimanExpressPerKm   int
	TarifPengirimanInstantPerKm   int
	MaksimalBidKurirReguler       int
	MaksimalBidKurirExpress       int
	MaksimalBidKurirInstant       int
	JarakMasukEkspedisi           int
	TarifPerkg                    int
	Catatan                       string
	BerlakuMulai                  time.Time
	BerlakuSampai                 time.Time
	StatusActive                  bool
	CreatedAt                     time.Time
	UpdatedAt                     time.Time
	DeletedAt                     gorm.DeletedAt
}

func (KebijakanSistem) TableNameHistorical() string {
	return "kebijakan_sistem_historical"
}

func (KebijakanSistem) TableNameSotReplica() string {
	return "kebijakan_sistem"
}

func (k KebijakanSistem) CreateHistoricalTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		ditetapkan_oleh text,
		id_admin bigint,
		nama_admin text,
		komisi_sistem_per_transaksi float,
		limit_membuat_diskon_personal int,
		limit_membuat_diskon_distributor int,
		limit_membuat_diskon_brand int,
		max_jarak_km_reguler int,
		max_jarak_km_express int,
		max_jarak_km_instant int,
		estimasi_hari_reguler int,
		estimasi_hari_express int,
		estimasi_hari_instant int,
		tarif_pengiriman_reguler_per_km int,
		tarif_pengiriman_express_per_km int,
		tarif_pengiriman_instant_per_km int,
		maksimal_bid_kurir_reguler int,
		maksimal_bid_kurir_express int,
		maksimal_bid_kurir_instant int,
		jarak_masuk_ekspedisi int,
		tarif_perkg int,
		catatan text,
		berlaku_mulai timestamp,
		berlaku_sampai timestamp,
		status_active boolean,
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

func (k KebijakanSistem) ParseToCUDType() map[string]interface{} {
	return map[string]interface{}{
		"id":                               k.ID,
		"ditetapkan_oleh":                  k.DitetapkanOleh,
		"id_admin":                         k.IDAdmin,
		"nama_admin":                       k.NamaAdmin,
		"komisi_sistem_per_transaksi":      k.KomisiSistemPerTransaksi,
		"limit_membuat_diskon_personal":    k.LimitMembuatDiskonPersonal,
		"limit_membuat_diskon_distributor": k.LimitMembuatDiskonDistributor,
		"limit_membuat_diskon_brand":       k.LimitMembuatDiskonBrand,
		"max_jarak_km_reguler":             k.MaxJarakKmReguler,
		"max_jarak_km_express":             k.MaxJarakKmExpress,
		"max_jarak_km_instant":             k.MaxJarakKmInstant,
		"estimasi_hari_reguler":            k.EstimasiHariReguler,
		"estimasi_hari_express":            k.EstimasiHariExpress,
		"estimasi_hari_instant":            k.EstimasiHariInstant,
		"tarif_pengiriman_reguler_per_km":  k.TarifPengirimanRegulerPerKm,
		"tarif_pengiriman_express_per_km":  k.TarifPengirimanExpressPerKm,
		"tarif_pengiriman_instant_per_km":  k.TarifPengirimanInstantPerKm,
		"maksimal_bid_kurir_reguler":       k.MaksimalBidKurirReguler,
		"maksimal_bid_kurir_express":       k.MaksimalBidKurirExpress,
		"maksimal_bid_kurir_instant":       k.MaksimalBidKurirInstant,
		"jarak_masuk_ekspedisi":            k.JarakMasukEkspedisi,
		"tarif_perkg":                      k.TarifPerkg,
		"catatan":                          k.Catatan,
		"berlaku_mulai":                    k.BerlakuMulai,
		"berlaku_sampai":                   k.BerlakuSampai,
		"status_active":                    k.StatusActive,
		"created_at":                       k.CreatedAt,
		"updated_at":                       k.UpdatedAt,
		"deleted_at":                       k.DeletedAt,
	}
}

func (k KebijakanSistem) CreateSotReplicaTable(ctx context.Context, session *gocql.Session) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		ditetapkan_oleh text,
		id_admin bigint,
		nama_admin text,
		komisi_sistem_per_transaksi float,
		limit_membuat_diskon_personal int,
		limit_membuat_diskon_distributor int,
		limit_membuat_diskon_brand int,
		max_jarak_km_reguler int,
		max_jarak_km_express int,
		max_jarak_km_instant int,
		estimasi_hari_reguler int,
		estimasi_hari_express int,
		estimasi_hari_instant int,
		tarif_pengiriman_reguler_per_km int,
		tarif_pengiriman_express_per_km int,
		tarif_pengiriman_instant_per_km int,
		maksimal_bid_kurir_reguler int,
		maksimal_bid_kurir_express int,
		maksimal_bid_kurir_instant int,
		jarak_masuk_ekspedisi int,
		tarif_perkg int,
		catatan text,
		berlaku_mulai timestamp,
		berlaku_sampai timestamp,
		status_active boolean,
		created_at timestamp,
		updated_at timestamp,
		deleted_at timestamp,
		PRIMARY KEY (id)
	)`, k.TableNameSotReplica())

	if err := session.Query(query).ExecContext(ctx); err != nil {
		fmt.Println("Gagal eksekusi query:", err)
		return err
	}
	fmt.Printf("Berhasil Eksekusi query membuat tabel %s\n", k.TableNameSotReplica())
	return nil
}
