package auth_handle

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/cache"
	cache_db_function "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cache_db/function"
	cache_db_session "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cache_db/session"
	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	se_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/search_engine/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	notification_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/models"
	notification_request "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/request"
	notification_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/seeders"
)

func CreateValidatePenggunaRegistration(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateValidatePenggunaRegistration"
	var Objek sot_models.Pengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
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
		UpdatedAt:      Objek.UpdatedAt,
		DeletedAt:      Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica %s dalam %s", err, handle_services)
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
		StatusPengguna: Objek.StatusPengguna,
		CreatedAt:      Objek.CreatedAt,
	}

	if _, err := se_index.PenggunaIndex.AddDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke search engine %s dalam %s", err, handle_services)
	}

	// ðŸ”” Notifikasi Welcome Pengguna Baru
	if Objek.ID != 0 {
		var Notif = notification_models.NotificationPengguna{
			IDPengguna: Objek.ID,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸŽ‰ Registrasi Berhasil!",
			Pesan:      fmt.Sprintf("Halo %s, selamat datang di platform kami!", Objek.Nama),
			Pop:        1,
			Activity:   true,
			Inbox:      false,
			Archive:    true,
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		}
		_ = notification_request.PostToNotification(ctx, Notif, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil meregistrasikan pengguna", Objek.ID)
	return nil
}

func CreateValidateSellerRegistration(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateValidateSellerRegistration"
	var Objek sot_models.Seller
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	}

	var ObjekCass cass_models.Seller = cass_models.Seller{
		ID:               int(Objek.ID),
		Username:         Objek.Username,
		Nama:             Objek.Nama,
		Email:            Objek.Email,
		Jenis:            Objek.Jenis,
		SellerDedication: Objek.SellerDedication,
		JamOperasional:   Objek.JamOperasional,
		Punchline:        Objek.Punchline,
		Password:         Objek.Password,
		Deskripsi:        Objek.Deskripsi,
		StatusSeller:     Objek.StatusSeller,
		CreatedAt:        Objek.CreatedAt,
		UpdatedAt:        Objek.UpdatedAt,
		DeletedAt:        Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Notifikasi Seller Baru
	if Objek.ID != 0 {
		var Notif = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.ID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸª Toko Anda Berhasil Terdaftar",
			Pesan:     fmt.Sprintf("Selamat, Toko %s sekarang sudah aktif. Yuk mulai upload produk pertamamu!", Objek.Nama),
			Pop:       1,
			Activity:  true,
			Inbox:     false,
			Archive:   true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notif, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil meregistrasikan seller", Objek.ID)
	return nil
}

func CreateValidateKurirRegistration(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateValidateKurirRegistration"
	var Objek sot_models.Kurir
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	}

	var ObjekCass cass_models.Kurir = cass_models.Kurir{
		ID:            Objek.ID,
		Nama:          Objek.Nama,
		Username:      Objek.Username,
		Email:         Objek.Email,
		Jenis:         Objek.Jenis,
		PasswordHash:  Objek.PasswordHash,
		Deskripsi:     Objek.Deskripsi,
		StatusKurir:   Objek.StatusKurir,
		StatusBid:     Objek.StatusBid,
		VerifiedKurir: Objek.VerifiedKurir,
		Rating:        Objek.Rating,
		TipeKendaraan: Objek.TipeKendaraan,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
		DeletedAt:     Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Notifikasi Kurir Baru
	if Objek.ID != 0 {
		var Notif = notification_models.NotificationKurir{
			IDKurir:   Objek.ID,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ›µ Akun Kurir Aktif",
			Pesan:     "Pendaftaran kurir disetujui. Siap-siap terima orderan!",
			Pop:       1,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		}
		_ = notification_request.PostToNotification(ctx, Notif, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil meregistrasikan kurir", Objek.ID)
	return nil
}

// ==========================================
// LOGIN LIFECYCLE (UPDATE SESSION & STATE)
// ==========================================

func UpdatePenggunaLogin(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, rds_session *redis.Client) error {
	const handle_services string = "UpdatePenggunaLogin"
	var Pengguna sot_models.Pengguna
	if err := helper.DecodeJSONBody(Data, &Pengguna); err != nil {
		return fmt.Errorf("gagal mengolah data")
	}

	var ObjekCass cass_models.Pengguna = cass_models.Pengguna{
		ID:             Pengguna.ID,
		Username:       Pengguna.Username,
		Nama:           Pengguna.Nama,
		Email:          Pengguna.Email,
		PasswordHash:   Pengguna.PasswordHash,
		PinHash:        Pengguna.PinHash,
		StatusPengguna: Pengguna.StatusPengguna,
		CreatedAt:      Pengguna.CreatedAt,
		UpdatedAt:      Pengguna.UpdatedAt,
		DeletedAt:      Pengguna.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate login ke sot replica %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal mencatat login ke historical %s dalam %s", err, handle_services)
	}

	// ðŸ”„ Refresh Session Data di Redis Cache
	if err := cache_db_function.UpdateSessionData[sot_models.Pengguna](ctx, *rds_session, cache_db_session.GetSessionKey[*sot_models.Pengguna](&Pengguna), Pengguna); err != nil {
		return fmt.Errorf("gagal mereset sesi pengguna di redis")
	}

	// ðŸ”” Silent Update untuk mencocokkan Session State lokal device
	if Pengguna.ID != 0 {
		var Notif = notification_models.NotificationPengguna{
			IDPengguna: Pengguna.ID,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "Halo Login Berhasil",
			Pesan:      "Sesi login Anda telah diperbarui.",
			Pop:        0,
			Activity:   true,
			Inbox:      false,
			Archive:    true,
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"pengguna_id": Pengguna.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_SESSION"},
			},
		}
		if err := notification_request.PostToNotification(ctx, Notif, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk); err != nil {
			fmt.Println("[TRACE ERROR NOTIFICATION POST]", err)
		} else {
			fmt.Println("[TRACE SUCESS NOTIFICATION POST]", err)
		}
	}

	fmt.Println("Berhasil memperbarui data login pengguna", Pengguna.ID)
	return nil
}

func UpdateSellerLogin(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateSellerLogin"
	var Objek sot_models.Seller
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	}

	var ObjekCass cass_models.Seller = cass_models.Seller{
		ID:               int(Objek.ID),
		Username:         Objek.Username,
		Nama:             Objek.Nama,
		Email:            Objek.Email,
		Jenis:            Objek.Jenis,
		SellerDedication: Objek.SellerDedication,
		JamOperasional:   Objek.JamOperasional,
		Punchline:        Objek.Punchline,
		Password:         Objek.Password,
		Deskripsi:        Objek.Deskripsi,
		StatusSeller:     Objek.StatusSeller,
		CreatedAt:        Objek.CreatedAt,
		UpdatedAt:        Objek.UpdatedAt,
		DeletedAt:        Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate login seller ke sot replica %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal mencatat login seller ke historical %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Sync Dashboard Seller setelah login berhasil
	if Objek.ID != 0 {
		var Notif = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.ID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ” Seller Login",
			Pesan:     "Sinkronisasi dashboard penjual.",
			Pop:       0,
			Activity:  true,
			Inbox:     false,
			Archive:   true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notif, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil memperbarui data login seller", Objek.ID)
	return nil
}

func UpdateKurirLogin(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateKurirLogin"
	var Objek sot_models.Kurir
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	}

	var ObjekCass cass_models.Kurir = cass_models.Kurir{
		ID:            Objek.ID,
		Nama:          Objek.Nama,
		Username:      Objek.Username,
		Email:         Objek.Email,
		Jenis:         Objek.Jenis,
		PasswordHash:  Objek.PasswordHash,
		Deskripsi:     Objek.Deskripsi,
		StatusKurir:   Objek.StatusKurir,
		StatusBid:     Objek.StatusBid,
		VerifiedKurir: Objek.VerifiedKurir,
		Rating:        Objek.Rating,
		TipeKendaraan: Objek.TipeKendaraan,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
		DeletedAt:     Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate login kurir ke sot replica %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal mencatat login kurir ke historical %s dalam %s", err, handle_services)
	}

	if Objek.ID != 0 {
		var Notif = notification_models.NotificationKurir{
			IDKurir:   Objek.ID,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ” Kurir Login",
			Pesan:     "Mengaktifkan pelacakan background.",
			Pop:       0,
			Activity:  true,
			Inbox:     false,
			Archive:   true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
		}
		_ = notification_request.PostToNotification(ctx, Notif, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil memperbarui data login kurir", Objek.ID)
	return nil
}
