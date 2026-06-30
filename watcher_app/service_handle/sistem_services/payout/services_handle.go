package payout_sistem_handle

import (
	"context"
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func CreatePayOutSistem(data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreatePayoutSistem"
	var Objek sot_models.PayOutSistem

	if err := helper.DecodeJSONBody(data, &Objek); err != nil {
		return err
	}

	var ObjekCass cass_models.PayOutSistem = cass_models.PayOutSistem{
		ID:               Objek.ID,
		IdDisburstment:   Objek.IdDisburstment,
		IdTransaksi:      Objek.IdTransaksi,
		UserId:           Objek.UserId,
		Amount:           Objek.Amount,
		Status:           Objek.Status,
		Reason:           Objek.Reason,
		Timestamp:        Objek.Timestamp,
		BankCode:         Objek.BankCode,
		AccountNumber:    Objek.AccountNumber,
		RecipientName:    Objek.RecipientName,
		SenderBank:       Objek.SenderBank,
		Remark:           Objek.Remark,
		Receipt:          Objek.Receipt,
		TimeServed:       Objek.TimeServed,
		BundleId:         Objek.BundleId,
		CompanyId:        Objek.CompanyId,
		RecipientCity:    Objek.RecipientCity,
		CreatedFrom:      Objek.CreatedFrom,
		Direction:        Objek.Direction,
		Sender:           Objek.Sender,
		Fee:              Objek.Fee,
		BeneficiaryEmail: Objek.BeneficiaryEmail,
		IdempotencyKey:   Objek.IdempotencyKey,
		IsVirtualAccount: Objek.IsVirtualAccount,
		DeletedAt:        Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePayOutSistem(data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatePayoutSistem"
	var Objek sot_models.PayOutSistem

	if err := helper.DecodeJSONBody(data, &Objek); err != nil {
		return err
	}

	var ObjekCass cass_models.PayOutSistem = cass_models.PayOutSistem{
		ID:               Objek.ID,
		IdDisburstment:   Objek.IdDisburstment,
		IdTransaksi:      Objek.IdTransaksi,
		UserId:           Objek.UserId,
		Amount:           Objek.Amount,
		Status:           Objek.Status,
		Reason:           Objek.Reason,
		Timestamp:        Objek.Timestamp,
		BankCode:         Objek.BankCode,
		AccountNumber:    Objek.AccountNumber,
		RecipientName:    Objek.RecipientName,
		SenderBank:       Objek.SenderBank,
		Remark:           Objek.Remark,
		Receipt:          Objek.Receipt,
		TimeServed:       Objek.TimeServed,
		BundleId:         Objek.BundleId,
		CompanyId:        Objek.CompanyId,
		RecipientCity:    Objek.RecipientCity,
		CreatedFrom:      Objek.CreatedFrom,
		Direction:        Objek.Direction,
		Sender:           Objek.Sender,
		Fee:              Objek.Fee,
		BeneficiaryEmail: Objek.BeneficiaryEmail,
		IdempotencyKey:   Objek.IdempotencyKey,
		IsVirtualAccount: Objek.IsVirtualAccount,
		DeletedAt:        Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreatePayoutSeller(data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreatePayoutSeller"
	var Objek sot_models.PayOutSeller

	if err := helper.DecodeJSONBody(data, &Objek); err != nil {
		return err
	}

	var ObjekCass cass_models.PayOutSeller = cass_models.PayOutSeller{

		IdSeller:         Objek.IdSeller,
		ID:               Objek.ID,
		IdDisbursment:    Objek.IdDisbursment,
		IdTransaksi:      Objek.IdTransaksi,
		UserId:           Objek.UserId,
		Amount:           Objek.Amount,
		Status:           Objek.Status,
		Reason:           Objek.Reason,
		Timestamp:        Objek.Timestamp,
		BankCode:         Objek.BankCode,
		AccountNumber:    Objek.AccountNumber,
		RecipientName:    Objek.RecipientName,
		SenderBank:       Objek.SenderBank,
		Remark:           Objek.Remark,
		Receipt:          Objek.Receipt,
		TimeServed:       Objek.TimeServed,
		BundleId:         Objek.BundleId,
		CompanyId:        Objek.CompanyId,
		RecipientCity:    Objek.RecipientCity,
		CreatedFrom:      Objek.CreatedFrom,
		Direction:        Objek.Direction,
		Sender:           Objek.Sender,
		Fee:              Objek.Fee,
		BeneficiaryEmail: Objek.BeneficiaryEmail,
		IdempotencyKey:   Objek.IdempotencyKey,
		IsVirtualAccount: Objek.IsVirtualAccount,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)

	return nil
}

func UpdatePayoutSeller(data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatePayoutSeller"
	var Objek sot_models.PayOutSeller

	if err := helper.DecodeJSONBody(data, &Objek); err != nil {
		return err
	}

	var ObjekCass cass_models.PayOutSeller = cass_models.PayOutSeller{
		ID:               Objek.ID,
		IdSeller:         Objek.IdSeller, // 🔵 Dipastikan ikut ter-mapping
		IdDisbursment:    Objek.IdDisbursment,
		IdTransaksi:      Objek.IdTransaksi,
		UserId:           Objek.UserId,
		Amount:           Objek.Amount,
		Status:           Objek.Status,
		Reason:           Objek.Reason,
		Timestamp:        Objek.Timestamp,
		BankCode:         Objek.BankCode,
		AccountNumber:    Objek.AccountNumber,
		RecipientName:    Objek.RecipientName,
		SenderBank:       Objek.SenderBank,
		Remark:           Objek.Remark,
		Receipt:          Objek.Receipt,
		TimeServed:       Objek.TimeServed,
		BundleId:         Objek.BundleId,
		CompanyId:        Objek.CompanyId,
		RecipientCity:    Objek.RecipientCity,
		CreatedFrom:      Objek.CreatedFrom,
		Direction:        Objek.Direction,
		Sender:           Objek.Sender,
		Fee:              Objek.Fee,
		BeneficiaryEmail: Objek.BeneficiaryEmail,
		IdempotencyKey:   Objek.IdempotencyKey,
		IsVirtualAccount: Objek.IsVirtualAccount,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	// 🔵 Menggunakan UpdateData untuk tabel SOT Replica
	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memperbarui data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	// 🔵 Tetap menggunakan InsertData untuk Historical karena sifatnya append-only log (mencatat riwayat perubahan baru)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data riwayat ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil memperbarui data seller", Objek.ID)
	return nil
}

func UpdatePayoutKurir(data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatePayoutKurir"
	var Objek sot_models.PayOutKurir

	if err := helper.DecodeJSONBody(data, &Objek); err != nil {
		return err
	}

	var ObjekCass cass_models.PayOutKurir = cass_models.PayOutKurir{
		ID:               Objek.ID,
		IdKurir:          Objek.IdKurir,      // 🔵 Spesifik Kurir
		IdPengiriman:     Objek.IdPengiriman, // 🔵 Spesifik Kurir
		IdDisbursment:    Objek.IdDisbursment,
		UserId:           Objek.UserId,
		Amount:           Objek.Amount,
		Status:           Objek.Status,
		Reason:           Objek.Reason,
		Timestamp:        Objek.Timestamp,
		BankCode:         Objek.BankCode,
		AccountNumber:    Objek.AccountNumber,
		RecipientName:    Objek.RecipientName,
		SenderBank:       Objek.SenderBank,
		Remark:           Objek.Remark,
		Receipt:          Objek.Receipt,
		TimeServed:       Objek.TimeServed,
		BundleId:         Objek.BundleId,
		CompanyId:        Objek.CompanyId,
		RecipientCity:    Objek.RecipientCity,
		CreatedFrom:      Objek.CreatedFrom,
		Direction:        Objek.Direction,
		Sender:           Objek.Sender,
		Fee:              Objek.Fee,
		BeneficiaryEmail: Objek.BeneficiaryEmail,
		IdempotencyKey:   Objek.IdempotencyKey,
		IsVirtualAccount: Objek.IsVirtualAccount,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	// 🔵 Eksekusi update pada table Replica
	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memperbarui data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	// 🔵 Catat perubahan (termasuk status jika dihapus) ke tabel Historical log
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data riwayat ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil memperbarui data kurir", Objek.ID)
	return nil
}
