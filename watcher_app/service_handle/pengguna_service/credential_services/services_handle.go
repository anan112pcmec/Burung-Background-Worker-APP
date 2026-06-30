package credential_pengguna_handle

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	cache_db_function "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cache_db/function"
	cache_db_session "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cache_db/session"
	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	se_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/search_engine/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/environment"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	notification_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/models"
	notification_request "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/request"
	notification_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/seeders"
)

func UpdateValidateUbahPasswordPenggunaViaOtp(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services string = "UpdateValidateUbahPasswordPenggunaViaOtp"
	var Objek sot_models.Pengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Pengguna = cass_models.Pengguna{
		ID:             Objek.ID,
		Username:       Objek.Username,
		Nama:           Objek.Nama,
		Email:          Objek.Email,
		PasswordHash:   Objek.PasswordHash,
		PinHash:        Objek.PinHash,
		StatusPengguna: Objek.StatusPengguna,
		CreatedAt:      Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Pengguna = se_models.Pengguna{
		ID:             Objek.ID,
		Username:       Objek.Username,
		Nama:           Objek.Nama,
		Email:          Objek.Email,
		PasswordHash:   Objek.PasswordHash,
		PinHash:        Objek.PinHash,
		StatusPengguna: Objek.StatusPengguna,
		CreatedAt:      Objek.CreatedAt,
	}

	if task_info, err := se_index.PenggunaIndex.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke dalam search engine dengan UID %s", task_info.IndexUID)
	}

	if err := cache_db_function.UpdateSessionData[sot_models.Pengguna](ctx, *rds_session, cache_db_session.GetSessionKey[*sot_models.Pengguna](&Objek), Objek); err != nil {
		return fmt.Errorf("gagal memasukan mengubah data sesi pengguna")
	}

	// Setup notifikasi security alert (Tegas, formal, & informatif)
	judulSecurity := "⚠️ Keamanan Akun: Sandi Berhasil Diubah"
	pesanSecurity := "Kata sandi akun Anda telah berhasil diperbarui melalui verifikasi OTP. Jika Anda tidak merasa melakukan perubahan ini, segera hubungi Pusat Bantuan."

	var NotificationSecurity notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: Objek.ID, // Langsung kirim ke pengguna terkait
		Pengirim:   notification_seeders.Sistem,
		Judul:      judulSecurity,
		Pesan:      pesanSecurity,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 30).Format(time.RFC3339), // Simpan 30 hari di inbox karena info krusial
		Pop:        5.0,                                               // 5 detik, biar user sadar ada aktivitas keamanan di akunnya
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"user_id":     Objek.ID,
				"action_type": "security_password_changed",
				"auth_method": "otp",
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_SECURITY_SETTINGS_PAGE", // FE arahin ke halaman pengaturan keamanan/log aktivitas device
			},
		},
	}

	// Kirim ke path pengguna umum
	if err := notification_request.PostToNotification(ctx, NotificationSecurity, environment.HostRunningAPIInNotifikasi, environment.PortRunningAPIInNotifikasi, environment.PenggunaPathNotifikasiMasuk); err != nil {
		// Log aja kalau gagal kirim notif, jangan gagalin proses ganti password-nya
		fmt.Printf("Gagal mengirim notifikasi keamanan password untuk user %d: %v\n", Objek.ID, err)
	}

	return nil
}

func UpdateValidateUbahPasswordPenggunaViaPin(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services string = "UpdateValidateUbahPasswordPenggunaViaPin"
	var Objek sot_models.Pengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Pengguna = cass_models.Pengguna{
		ID:             Objek.ID,
		Username:       Objek.Username,
		Nama:           Objek.Nama,
		Email:          Objek.Email,
		PasswordHash:   Objek.PasswordHash,
		PinHash:        Objek.PinHash,
		StatusPengguna: Objek.StatusPengguna,
		CreatedAt:      Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Pengguna = se_models.Pengguna{
		ID:             Objek.ID,
		Username:       Objek.Username,
		Nama:           Objek.Nama,
		Email:          Objek.Email,
		PasswordHash:   Objek.PasswordHash,
		PinHash:        Objek.PinHash,
		StatusPengguna: Objek.StatusPengguna,
		CreatedAt:      Objek.CreatedAt,
	}

	if task_info, err := se_index.PenggunaIndex.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Println("berhasil memasukan data ke dalam search engine dengan UID %s", task_info.IndexUID)
	}

	if err := cache_db_function.UpdateSessionData[sot_models.Pengguna](ctx, *rds_session, cache_db_session.GetSessionKey[*sot_models.Pengguna](&Objek), Objek); err != nil {
		return fmt.Errorf("gagal memasukan mengubah data sesi pengguna")
	}

	// Notifikasi Keamanan: Ubah Password via PIN
	var NotificationPasswordViaPin notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: Objek.ID,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "⚠️ Keamanan Akun: Sandi Berhasil Diubah",
		Pesan:      "Kata sandi akun Anda telah berhasil diperbarui melalui verifikasi PIN keamanan. Jika Anda tidak merasa melakukan perubahan ini, segera amankan akun Anda.",
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 30).Format(time.RFC3339), // Log keamanan disimpan 30 hari
		Pop:        5.0,                                               // 5 detik biar terbaca jelas
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"user_id":     Objek.ID,
				"action_type": "security_password_changed",
				"auth_method": "pin",
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_SECURITY_SETTINGS_PAGE",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotificationPasswordViaPin, environment.HostRunningAPIInNotifikasi, environment.PortRunningAPIInNotifikasi, environment.PenggunaPathNotifikasiMasuk); err != nil {
		fmt.Printf("Gagal mengirim notifikasi keamanan password via PIN untuk user %d: %v\n", Objek.ID, err)
	}

	fmt.Printf("Berhasil memproses notifikasi keamanan ubah password via PIN untuk User ID %d\n", Objek.ID)

	return nil
}

func CreateMembuatSecretPinPengguna(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services string = "CreateMembuatSecretPinPengguna"
	var Objek sot_models.Pengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Pengguna = cass_models.Pengguna{
		ID:             Objek.ID,
		Username:       Objek.Username,
		Nama:           Objek.Nama,
		Email:          Objek.Email,
		PasswordHash:   Objek.PasswordHash,
		PinHash:        Objek.PinHash,
		StatusPengguna: Objek.StatusPengguna,
		CreatedAt:      Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Pengguna = se_models.Pengguna{
		ID:             Objek.ID,
		Username:       Objek.Username,
		Nama:           Objek.Nama,
		Email:          Objek.Email,
		PasswordHash:   Objek.PasswordHash,
		PinHash:        Objek.PinHash,
		StatusPengguna: Objek.StatusPengguna,
		CreatedAt:      Objek.CreatedAt,
	}

	if task_info, err := se_index.PenggunaIndex.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke dalam search engine dengan UID %s", task_info.IndexUID)
	}

	if err := cache_db_function.UpdateSessionData[sot_models.Pengguna](ctx, *rds_session, cache_db_session.GetSessionKey[*sot_models.Pengguna](&Objek), Objek); err != nil {
		return fmt.Errorf("gagal memasukan mengubah data sesi pengguna")
	}

	// Notifikasi Keamanan: Pembuatan PIN Baru
	var NotificationCreatePin notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: Objek.ID,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "🔒 PIN Keamanan Berhasil Dibuat",
		Pesan:      "PIN keamanan akun Anda telah berhasil didaftarkan. Gunakan PIN ini untuk menjaga keamanan transaksi dan perubahan data penting Anda.",
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 14).Format(time.RFC3339), // Simpan log 14 hari
		Pop:        4.0,                                               // 4 detik dirasa cukup untuk info registrasi sukses
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"user_id":     Objek.ID,
				"action_type": "security_pin_created",
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_ACCOUNT_DASHBOARD", // Arahkan ke dashboard akun utama setelah sukses
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotificationCreatePin, environment.HostRunningAPIInNotifikasi, environment.PortRunningAPIInNotifikasi, environment.PenggunaPathNotifikasiMasuk); err != nil {
		fmt.Printf("Gagal mengirim notifikasi pembuatan PIN untuk user %d: %v\n", Objek.ID, err)
	}

	fmt.Printf("Berhasil memproses notifikasi pembuatan PIN baru untuk User ID %d\n", Objek.ID)

	return nil
}

func UpdateSecretPinPengguna(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services string = "UpdateSecretPinPengguna"
	var Objek sot_models.Pengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Pengguna = cass_models.Pengguna{
		ID:             Objek.ID,
		Username:       Objek.Username,
		Nama:           Objek.Nama,
		Email:          Objek.Email,
		PasswordHash:   Objek.PasswordHash,
		PinHash:        Objek.PinHash,
		StatusPengguna: Objek.StatusPengguna,
		CreatedAt:      Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Pengguna = se_models.Pengguna{
		ID:             Objek.ID,
		Username:       Objek.Username,
		Nama:           Objek.Nama,
		Email:          Objek.Email,
		PasswordHash:   Objek.PasswordHash,
		PinHash:        Objek.PinHash,
		StatusPengguna: Objek.StatusPengguna,
		CreatedAt:      Objek.CreatedAt,
	}

	if task_info, err := se_index.PenggunaIndex.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Println("berhasil memasukan data ke dalam search engine dengan UID %s", task_info.IndexUID)
	}

	if err := cache_db_function.UpdateSessionData[sot_models.Pengguna](ctx, *rds_session, cache_db_session.GetSessionKey[*sot_models.Pengguna](&Objek), Objek); err != nil {
		return fmt.Errorf("gagal memasukan mengubah data sesi pengguna")
	}

	return nil
}
