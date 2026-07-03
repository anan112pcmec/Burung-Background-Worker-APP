package payout_sistem_handle

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/cache"
	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	notification_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/models"
	notification_request "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/request"
	notification_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/seeders"
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

	// ðŸ”” Notifikasi Seller: Penarikan dana baru diajukan / diproses
	if Objek.IdSeller != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.IdSeller),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ’¸ Penarikan Dana Diproses",
			Pesan:     fmt.Sprintf("Permintaan pencairan dana sebesar Rp%d sedang diproses ke rekening %s.", Objek.Amount, Objek.BankCode),
			Pop:       0.5,
			Archive:   true,
			Inbox:     true,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.IdSeller, "payout_id": Objek.ID, "amount": Objek.Amount, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "OPEN_PAYOUT_DETAIL"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

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
		IdSeller:         Objek.IdSeller,
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

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memperbarui data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data riwayat ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil memperbarui data seller", Objek.ID)

	// ðŸ”” Notifikasi Seller: Perubahan status payout (Berhasil / Gagal)
	if Objek.IdSeller != 0 {
		judulNotif := "ðŸ’° Penarikan Dana Berhasil"
		pesanNotif := fmt.Sprintf("Dana sebesar Rp%d telah berhasil ditransfer ke rekening Anda.", Objek.Amount)

		if Objek.Status == "FAILED" || Objek.Status == "REJECTED" {
			judulNotif = "âš ï¸ Penarikan Dana Gagal"
			pesanNotif = fmt.Sprintf("Pencairan dana Rp%d gagal diproses. Alasan: %s", Objek.Amount, Objek.Reason)
		}

		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.IdSeller),
			Pengirim:  notification_seeders.Sistem,
			Judul:     judulNotif,
			Pesan:     pesanNotif,
			Pop:       0.5,
			Archive:   true,
			Inbox:     true,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.IdSeller, "payout_id": Objek.ID, "status": Objek.Status, "reason": Objek.Reason},
				Special:  map[string]interface{}{"click_action": "OPEN_PAYOUT_HISTORY"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

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
		IdKurir:          Objek.IdKurir,
		IdPengiriman:     Objek.IdPengiriman,
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

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memperbarui data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data riwayat ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil memperbarui data kurir", Objek.ID)

	// ðŸ”” Notifikasi Kurir: Info dompet / pencairan ongkir pengiriman selesai didistribusikan
	if Objek.IdKurir != 0 {
		judulNotif := "ðŸ’µ Pendapatan Masuk!"
		pesanNotif := fmt.Sprintf("Ongkir/Insentif sebesar Rp%d untuk pengiriman #%d telah ditransfer.", Objek.Amount, Objek.IdPengiriman)

		if Objek.Status == "FAILED" || Objek.Status == "REJECTED" {
			judulNotif = "âš ï¸ Payout Kurir Gagal"
			pesanNotif = fmt.Sprintf("Gagal mengirimkan Rp%d ke rekening kurir. Status: %s", Objek.Amount, Objek.Reason)
		}

		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.IdKurir,
			Pengirim:  notification_seeders.Sistem,
			Judul:     judulNotif,
			Pesan:     pesanNotif,
			Pop:       0.5,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"kurir_id": Objek.IdKurir, "pengiriman_id": Objek.IdPengiriman, "payout_id": Objek.ID, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "OPEN_COURIER_WALLET"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	return nil
}
