package pengiriman_kurir_handle

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

func CreateAktifkanBidKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateAktifkanBidKurir"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         *Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Ambil nama kurir berdasarkan IdKurir dari model
	var NamaKurir string = ""
	if Objek.IdKurir != 0 {
		if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", Objek.IdKurir).Limit(1).Take(&NamaKurir).Error; err != nil {
			fmt.Println("Gagal mengambil nama kurir:", err)
		}
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// ðŸ”” NOTIFIKASI CREATE AKTIFKAN BID KURIR (Muncul Pop-Up)
	var Notifikasi notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   Objek.IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸŸ¢ Status Bidding Aktif!",
		Pesan:     fmt.Sprintf("Halo %s, sesi bidding lu di area %s telah diaktifkan. Bersiaplah menerima alokasi orderan masuk!", NamaKurir, Objek.Kota),
		Pop:       3.0, // Muncul pop-up selama 3 detik biar kurir dapet konfirmasi mantap
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 3).Format(time.RFC3339), // Eksis 3 hari di inbox
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"bid_id":           Objek.ID,
				"action_type":      "activate_bid",
				"jenis_pengiriman": Objek.JenisPengiriman,
				"mode":             Objek.Mode,
				"platform":         "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REDIRECT_TO_DASHBOARD_LIVE",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		fmt.Println("Gagal mengirim notifikasi aktifkan bid kurir:", err)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAktifkanBidKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services string = "UpdateAktifkanBidKurir"
	var Objek sot_models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
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
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Kurir = se_models.Kurir{
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
		DeletedAt:     &Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.KurirIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke search engine dengan info: %s ", task_info.IndexUID)
	}

	if err := cache_db_function.UpdateSessionData(ctx, *rds_session, cache_db_session.GetSessionKey(&Objek), Objek); err != nil {
		return fmt.Errorf("gagal mengupdate session data %s dalam %s", err, handle_services)
	}

	// 🔔 Silent Update Kurir: Sinkronisasi data aktivasi bid ke perangkat lokal
	if Objek.ID != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.ID,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "🔄 Fitur Bid Diaktifkan",
			Pesan:     "Fitur bid Anda telah aktif. Sistem sedang menyinkronkan data ke perangkat.",
			Pop:       0, // Background sync tanpa mengganggu kurir
			Activity:  true,
			Archive:   false,
			Inbox:     false,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"kurir_id": Objek.ID, "sync_type": "BID_ACTIVATION"},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_BID_STATUS"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}
	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateUbahBidKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahBidKurir"
	var Objek sot_models.BidKurirData // Diselaraskan pakai BidKurirData sesuai pasangannya

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         *Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateUpdatePosisiBidKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatePosisiBidKurir"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         *Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data posisi ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data posisi ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Ambil nama kurir berdasarkan IdKurir dari model
	var NamaKurir string = ""
	if Objek.IdKurir != 0 {
		if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", Objek.IdKurir).Limit(1).Take(&NamaKurir).Error; err != nil {
			fmt.Println("Gagal mengambil nama kurir:", err)
		}
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// ðŸ”” NOTIFIKASI UPDATE POSISI/BID KURIR (Silent Update, Masuk Inbox)
	var NotifikasiUpdate notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   Objek.IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ“ Area Bidding Berhasil Diperbarui",
		Pesan:     fmt.Sprintf("Halo %s, koordinat posisi dan data bid aktif lu di area %s, %s telah diperbarui.", NamaKurir, Objek.Kota, Objek.Provinsi),
		Pop:       0, // Tetap 0 biar silent, ngebantu kenyamanan berkendara di jalan
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 3).Format(time.RFC3339), // Cukup simpan 3 hari karena data posisi dinamis
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"bid_id":       Objek.ID,
				"action_type":  "update_position_bid",
				"latitude":     Objek.Latitude,
				"longitude":    Objek.Longitude,
				"slot_tersisa": Objek.SlotTersisa,
				"platform":     "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REFRESH_BIDDING_MAP",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotifikasiUpdate, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		fmt.Println("Gagal mengirim notifikasi update posisi/bid kurir:", err)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateAmbilPengirimanNonEksManualRegulerIIBidKurirNonEksSchedulerCreatePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateAmbilPengirimanNonEksManualRegulerIIBidKurirNonEksSchedulerCreatePublish"
	var Objek sot_models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirNonEksScheduler = cass_models.BidKurirNonEksScheduler{
		ID:           Objek.ID,
		IdBid:        Objek.IdBid,
		IdKurir:      Objek.IdKurir,
		Urutan:       Objek.Urutan,
		IdPengiriman: Objek.IdPengiriman,
		Status:       Objek.Status,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var (
		IdPengguna int64 = 0
		IdSeller   int64 = 0
	)

	var pengiriman sot_models.Pengiriman
	if err := read.WithContext(ctx).Model(&sot_models.Pengiriman{}).Select("id_seller, id_alamat_pengguna").Where("id = ?", Objek.IdPengiriman).Limit(1).Take(&pengiriman).Error; err != nil {
		return err
	}

	IdSeller = pengiriman.IdSeller
	if err := read.WithContext(ctx).Model(&sot_models.AlamatPengguna{}).Select("id_pengguna").Where("id = ?", pengiriman.IdAlamatPengguna).Limit(1).Take(&IdPengguna).Error; err != nil {
		return err
	}

	// ðŸ”” Notifikasi Pengguna (Pencarian Kurir)
	var NotifPengguna = notification_models.NotificationPengguna{
		IDPengguna: IdPengguna,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸ” Mencari Kurir Terdekat",
		Pesan:      "Sistem kami sedang menjadwalkan alokasi kurir reguler terbaik untuk menjemput paket pesananmu.",
		Pop:        3.0,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": Objek.IdPengiriman, "scheduler_id": Objek.ID},
			Special:  map[string]interface{}{"click_action": "TRACK_DELIVERY_STATUS"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)

	// ðŸ”” Notifikasi Seller (Pencarian Kurir)
	var NotifSeller = notification_models.NotificationSeller{
		IDSeller:  IdSeller,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ¤– Sistem Mencari Kurir",
		Pesan:     "Alokasi otomatis sedang berjalan, kurir terdekat sedang diarahkan menuju lokasi toko Anda.",
		Pop:       3.0,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": Objek.IdPengiriman, "scheduler_id": Objek.ID},
			Special:  map[string]interface{}{"click_action": "MANAGE_ORDER"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// ===================================================================================================
// 2. DATA PENGIRIMAN UPDATED (NOTIFIKASI PENGGUNA, SELLER, & KURIR YANG DI-ASSIGN)
// ===================================================================================================
func UpdateAmbilPengirimanNonEksManualRegulerIIpengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateAmbilPengirimanNonEksManualRegulerIIpengirimanUpdatedPublish"
	var Objek sot_models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.Pengiriman = cass_models.Pengiriman{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatPengguna:  Objek.IdAlamatPengguna,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var IdPengguna int64 = 0
	if err := read.WithContext(ctx).Model(&sot_models.AlamatPengguna{}).Select("id_pengguna").Where("id = ?", Objek.IdAlamatPengguna).Limit(1).Take(&IdPengguna).Error; err != nil {
		fmt.Println("Gagal mengambil id pengguna:", err)
	}

	var NamaKurir string = ""
	if *Objek.IdKurir != 0 {
		if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", Objek.IdKurir).Limit(1).Take(&NamaKurir).Error; err != nil {
			fmt.Println("Gagal mengambil nama kurir:", err)
		}
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir Internal"
	}

	// ðŸ”” Notifikasi Pengguna (Kurir Berhasil Didapatkan)
	var NotifPengguna = notification_models.NotificationPengguna{
		IDPengguna: IdPengguna,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸ›µ Kurir Telah Ditemukan!",
		Pesan:      fmt.Sprintf("Pesananmu beres di-pickup oleh %s dan siap diantar langsung ke rumahmu.", NamaKurir),
		Pop:        3.0,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": Objek.ID, "id_kurir": Objek.IdKurir, "status": Objek.Status},
			Special:  map[string]interface{}{"click_action": "TRACK_DELIVERY_LIVE"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)

	// ðŸ”” Notifikasi Seller (Kurir Menuju Lokasi)
	var NotifSeller = notification_models.NotificationSeller{
		IDSeller:  Objek.IdSeller,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ¤ Kurir Siap Meluncur",
		Pesan:     fmt.Sprintf("Kurir %s telah dikonfirmasi mengamankan orderan ini. Harap siapkan paket Anda.", NamaKurir),
		Pop:       3.0,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": Objek.ID, "id_kurir": Objek.IdKurir},
			Special:  map[string]interface{}{"click_action": "MANAGE_ORDER"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// ===================================================================================================
// 3. UPDATE DATA KAPASITAS/SLOT BID KURIR (SILENT UPDATE - POP: 0)
// ===================================================================================================
func UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataAmbilPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataAmbilPengirimanUpdatedPublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         *Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var NamaKurir string = ""
	if Objek.IdKurir != 0 {
		_ = read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", Objek.IdKurir).Limit(1).Take(&NamaKurir).Error
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// ðŸ”” Silent Notif Kurir (Update Slot Sisa Muatan)
	var NotifKurir = notification_models.NotificationKurir{
		IDKurir:   Objek.IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ”„ Kapasitas Bid Diperbarui",
		Pesan:     fmt.Sprintf("Halo %s, slot kapasitas bid lu otomatis berkurang. Sisa slot penampungan: %d.", NamaKurir, Objek.SlotTersisa),
		Pop:       0, // Silent update biar gak menutupi layar navigasi kurir
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 2).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"bid_id": Objek.ID, "slot_tersisa": Objek.SlotTersisa},
			Special:  map[string]interface{}{"click_action": "REFRESH_DASHBOARD_CAPACITY"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// ===================================================================================================
// 4. UPDATE DATA STATUS BID KURIR (SILENT UPDATE - POP: 0)
// ===================================================================================================
func UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataStatusUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataStatusUpdatedPublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         *Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var NamaKurir string = ""
	if Objek.IdKurir != 0 {
		_ = read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", Objek.IdKurir).Limit(1).Take(&NamaKurir).Error
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// ðŸ”” Silent Notif Kurir (Update Status Kerja)
	var NotifKurir = notification_models.NotificationKurir{
		IDKurir:   Objek.IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ”„ Sinkronisasi Status Bidding",
		Pesan:     fmt.Sprintf("Halo %s, status alokasi bidding lu di sistem saat ini bernilai [%s].", NamaKurir, Objek.Status),
		Pop:       0, // Tetap silent agar background sync berjalan mulus tanpa mengganggu kurir
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 2).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"bid_id": Objek.ID, "status": Objek.Status},
			Special:  map[string]interface{}{"click_action": "REFRESH_STATUS_CORE"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateAmbilPengirimanEksManualRegulerIIbidKurirEksSchedulerCreatePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateAmbilPengirimanEksManualRegulerIIbidKurirEksSchedulerCreatePublish"
	var Objek sot_models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirEksScheduler = cass_models.BidKurirEksScheduler{
		ID:              Objek.ID,
		IdBid:           Objek.IdBid,
		IdKurir:         Objek.IdKurir,
		Urutan:          Objek.Urutan,
		IdPengirimanEks: Objek.IdPengirimanEks,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// Tracing IdPengiriman lewat tabel pengiriman_ekspedisis
	var IdPengiriman int64 = 0
	if err := read.WithContext(ctx).Table("pengiriman_ekspedisis").Select("id_pengiriman").Where("id = ?", Objek.IdPengirimanEks).Limit(1).Scan(&IdPengiriman).Error; err != nil {
		return err
	}

	var (
		IdPengguna int64 = 0
		IdSeller   int64 = 0
	)

	var pengiriman sot_models.Pengiriman
	if err := read.WithContext(ctx).Model(&sot_models.Pengiriman{}).Select("id_seller, id_alamat_pengguna").Where("id = ?", IdPengiriman).Limit(1).Take(&pengiriman).Error; err != nil {
		return err
	}

	IdSeller = pengiriman.IdSeller
	if err := read.WithContext(ctx).Model(&sot_models.AlamatPengguna{}).Select("id_pengguna").Where("id = ?", pengiriman.IdAlamatPengguna).Limit(1).Take(&IdPengguna).Error; err != nil {
		return err
	}

	// ðŸ”” Notifikasi Pengguna
	var NotifPengguna = notification_models.NotificationPengguna{
		IDPengguna: IdPengguna,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸ“¦ Mempersiapkan Kurir Ekspedisi",
		Pesan:      "Sistem sedang mencocokkan jadwal penjemputan paketmu dengan armada kurir ekspedisi mitra.",
		Pop:        3.0,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": IdPengiriman, "scheduler_eks_id": Objek.ID},
			Special:  map[string]interface{}{"click_action": "TRACK_DELIVERY_STATUS"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)

	// ðŸ”” Notifikasi Seller
	var NotifSeller = notification_models.NotificationSeller{
		IDSeller:  IdSeller,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ¤– Penjadwalan Ekspedisi Berjalan",
		Pesan:     "Kurir logistik pihak ketiga sedang di-assign secara otomatis untuk menjemput produk dari gudang/toko Anda.",
		Pop:       3.0,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": IdPengiriman, "scheduler_eks_id": Objek.ID},
			Special:  map[string]interface{}{"click_action": "MANAGE_ORDER"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// ===================================================================================================
// 2. PENGIRIMAN EKSPEDISI UPDATED (NOTIFIKASI PENGGUNA, SELLER, & KURIR EKSPEDISI)
// ===================================================================================================
func UpdateAmbilPengirimanEksManualRegulerIIpengirimanEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateAmbilPengirimanEksManualRegulerIIpengirimanEksUpdatedPublish"
	var Objek sot_models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.Pengiriman = cass_models.Pengiriman{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatPengguna:  Objek.IdAlamatPengguna,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var IdPengguna int64 = 0
	if err := read.WithContext(ctx).Model(&sot_models.AlamatPengguna{}).Select("id_pengguna").Where("id = ?", Objek.IdAlamatPengguna).Limit(1).Take(&IdPengguna).Error; err != nil {
		fmt.Println("Gagal mengambil id pengguna:", err)
	}

	var NamaKurir string = ""
	if *Objek.IdKurir != 0 {
		_ = read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", Objek.IdKurir).Limit(1).Take(&NamaKurir).Error
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir Vendor Ekspedisi"
	}

	// ðŸ”” Notifikasi Pengguna
	var NotifPengguna = notification_models.NotificationPengguna{
		IDPengguna: IdPengguna,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸšš Kurir Ekspedisi Dikonfirmasi!",
		Pesan:      fmt.Sprintf("Paketmu dikonfirmasi akan dibawa oleh %s. Status manifest pengiriman siap berjalan.", NamaKurir),
		Pop:        3.0,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": Objek.ID, "id_kurir": Objek.IdKurir, "status": Objek.Status},
			Special:  map[string]interface{}{"click_action": "TRACK_EXPEDITION_LIVE"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)

	// ðŸ”” Notifikasi Seller
	var NotifSeller = notification_models.NotificationSeller{
		IDSeller:  Objek.IdSeller,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ“¦ Kurir Mitra Ekspedisi Ditugaskan",
		Pesan:     fmt.Sprintf("Armada dari %s telah ditunjuk untuk menangani rute pengiriman ini.", NamaKurir),
		Pop:       3.0,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": Objek.ID, "id_kurir": Objek.IdKurir},
			Special:  map[string]interface{}{"click_action": "MANAGE_ORDER"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)

	// ðŸ”” Notifikasi Kurir Ekspedisi
	if *Objek.IdKurir != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   *Objek.IdKurir,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ’¼ Orderan Ekspedisi Baru Masuk",
			Pesan:     fmt.Sprintf("Halo %s, satu paket reguler via logistik ekspedisi resmi ditugaskan ke manifest penjemputan lu.", NamaKurir),
			Pop:       3.0,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman": Objek.ID, "jarak": Objek.JarakTempuh, "ongkir_paid": Objek.KurirPaid},
				Special:  map[string]interface{}{"click_action": "REDIRECT_TO_EXPEDITION_MANIFEST"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// ===================================================================================================
// 3. SINKRONISASI SLOT BID KURIR EKSPEDISI (SILENT UPDATE - POP: 0)
// ===================================================================================================
func UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataUpdatedPublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         *Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var NamaKurir string = ""
	if Objek.IdKurir != 0 {
		_ = read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", Objek.IdKurir).Limit(1).Take(&NamaKurir).Error
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir Ekspedisi"
	}

	// ðŸ”” Silent Notif Kurir
	var NotifKurir = notification_models.NotificationKurir{
		IDKurir:   Objek.IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ”„ Slot Bid Ekspedisi Terkalkulasi",
		Pesan:     fmt.Sprintf("Halo %s, slot muatan armada ekspedisi lu disesuaikan. Sisa slot: %d.", NamaKurir, Objek.SlotTersisa),
		Pop:       0, // Tetap silent agar background sync aman
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 2).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"bid_id": Objek.ID, "slot_tersisa": Objek.SlotTersisa, "is_ekspedisi": Objek.IsEkspedisi},
			Special:  map[string]interface{}{"click_action": "REFRESH_CAPACITY_DASHBOARD"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// ===================================================================================================
// 4. SINKRONISASI STATUS BID EKSPEDISI (SILENT UPDATE - POP: 0)
// ===================================================================================================
func UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataAmbilPengirimanEksStatusUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataAmbilPengirimanEksStatusUpdatedPublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         *Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var NamaKurir string = ""
	if Objek.IdKurir != 0 {
		_ = read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", Objek.IdKurir).Limit(1).Take(&NamaKurir).Error
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// ðŸ”” Silent Notif Kurir
	var NotifKurir = notification_models.NotificationKurir{
		IDKurir:   Objek.IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ”„ Sync Status Bid Ekspedisi",
		Pesan:     fmt.Sprintf("Halo %s, status manifest bid logistik lu saat ini ter-update [%s].", NamaKurir, Objek.Status),
		Pop:       0,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 2).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"bid_id": Objek.ID, "status": Objek.Status},
			Special:  map[string]interface{}{"click_action": "REFRESH_SYSTEM_CORE"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// ===================================================================================================
// 5. LOCK SCHEDULER SIAP ANTAR (NOTIFIKASI AKTIF KE KURIR)
// ===================================================================================================
func UpdateLockSiapAntarBidKurirIIEksScheduler(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateLockSiapAntarBidKurirIIEksScheduler"
	var Objek sot_models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirEksScheduler = cass_models.BidKurirEksScheduler{
		ID:              Objek.ID,
		IdBid:           Objek.IdBid,
		IdKurir:         Objek.IdKurir,
		Urutan:          Objek.Urutan,
		IdPengirimanEks: Objek.IdPengirimanEks,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var NamaKurir string = ""
	if Objek.IdKurir != 0 {
		_ = read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", Objek.IdKurir).Limit(1).Take(&NamaKurir).Error
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir Ekspedisi"
	}

	// ðŸ”” Notifikasi Kurir (Lock Siap Antar - Wajib Pop Up)
	if Objek.IdKurir != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.IdKurir,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”’ Jadwal Terkunci: Siap Antar!",
			Pesan:     fmt.Sprintf("Halo %s, manifes pengiriman eks #%d telah dikunci. Status: SIAP ANTAR. Silakan menuju lokasi jemput.", NamaKurir, Objek.IdPengirimanEks),
			Pop:       3.0, // Munculkan pop-up agar kurir sigap berkendara ke lokasi pickup/hub
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"scheduler_id": Objek.ID, "id_pengiriman_eks": Objek.IdPengirimanEks, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "OPEN_NAVIGATION_TO_PICKUP"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateLockSiapAntarBidKurirIINonEksScheduler(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateLockSiapAntarBidKurirIINonEksScheduler"
	var Objek sot_models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirNonEksScheduler = cass_models.BidKurirNonEksScheduler{
		ID:           Objek.ID,
		IdBid:        Objek.IdBid,
		IdKurir:      Objek.IdKurir,
		Urutan:       Objek.Urutan,
		IdPengiriman: Objek.IdPengiriman,
		Status:       Objek.Status,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var NamaKurir string = ""
	if Objek.IdKurir != 0 {
		_ = read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", Objek.IdKurir).Limit(1).Take(&NamaKurir).Error
	}
	if NamaKurir == "" {
		NamaKurir = "Rekan Kurir"
	}

	// ðŸ”” Notifikasi Kurir (Lock Siap Antar - Pop-Up Wajib Nyala)
	if Objek.IdKurir != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.IdKurir,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”’ Jadwal Terkunci: Gas Pickup!",
			Pesan:     fmt.Sprintf("Halo %s, alokasi orderan reguler #%d sudah dikunci ke manifest-mu. Status: SIAP ANTAR. Segera merapat ke lokasi seller!", NamaKurir, Objek.IdPengiriman),
			Pop:       3.0,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"scheduler_id": Objek.ID, "id_pengiriman": Objek.IdPengiriman, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "OPEN_ROUTING_MAP_TO_SELLER"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// ===================================================================================================
// 2. SINKRONISASI SLOT BID PASCA LOCK (SILENT UPDATE - POP: 0)
// ===================================================================================================
func UpdateLockSiapAntarBidKurirIIbidKurirDataLockSiapAntarUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateLockSiapAntarBidKurirIIbidKurirDataLockSiapAntarUpdatedPublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         *Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var NamaKurir string = ""
	if Objek.IdKurir != 0 {
		_ = read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", Objek.IdKurir).Limit(1).Take(&NamaKurir).Error
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// ðŸ”” Silent Notif Kurir
	var NotifKurir = notification_models.NotificationKurir{
		IDKurir:   Objek.IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ”„ Kapasitas Slot Terkunci",
		Pesan:     fmt.Sprintf("Halo %s, slot muatan otomatis dikunci akibat alokasi aktif. Sisa slot: %d.", NamaKurir, Objek.SlotTersisa),
		Pop:       0,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"bid_id": Objek.ID, "slot_tersisa": Objek.SlotTersisa},
			Special:  map[string]interface{}{"click_action": "REFRESH_CAPACITY"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// ===================================================================================================
// 3. UPDATE DATA PROFIL KURIR & ENGINE SYNC (SILENT UPDATE - POP: 0)
// ===================================================================================================
func UpdateLockSiapAntarBidKurirIIkurirLockSiapAntarUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services = "UpdateLockSiapAntarBidKurirIIKurirLockSiapAntarUpdatedPublish"
	var Objek sot_models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
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
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}
	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Kurir = se_models.Kurir{
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
	}

	if task_info, err := se_index.KurirIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke search engine dengan info: %s ", task_info.IndexUID)
	}

	if err := cache_db_function.UpdateSessionData(ctx, *rds_session, cache_db_session.GetSessionKey(&Objek), Objek); err != nil {
		return fmt.Errorf("gagal memperbarui data di cache %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Notif Kurir (Status Kurir Berubah Jadi Mengantar Paket di Core Engine)
	var NotifKurir = notification_models.NotificationKurir{
		IDKurir:   Objek.ID,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ”„ Status Akun Sinkron",
		Pesan:     fmt.Sprintf("Profil dan status kerja lu [%s] sukses diselaraskan ke server utama.", Objek.StatusKurir),
		Pop:       0,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"kurir_id": Objek.ID, "status_kurir": Objek.StatusKurir, "status_bid": Objek.StatusBid},
			Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_SESSION"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// ===================================================================================================
// 4. JEJAK PENGIRIMAN / PICKED UP (NOTIFIKASI AKTIF KE PENGGUNA & SELLER)
// ===================================================================================================
func CreatePickedUpPengirimanNonEksIIjejakPengirimanCreatePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreatePickedUpPengirimanNonEksIIjejakPengirimanCreatePublish"
	var Objek sot_models.JejakPengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengiriman = cass_models.JejakPengiriman{
		ID:           Objek.ID,
		IdPengiriman: Objek.IdPengiriman,
		Lokasi:       Objek.Lokasi,
		Keterangan:   Objek.Keterangan,
		Latitude:     Objek.Latitude,
		Longtitude:   Objek.Longtitude,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Lakukan tracing data Pengiriman untuk mengambil Seller ID & Pengguna ID
	var (
		IdPengguna int64 = 0
		IdSeller   int64 = 0
	)

	var pengiriman sot_models.Pengiriman
	if err := read.WithContext(ctx).Model(&sot_models.Pengiriman{}).Select("id_seller, id_alamat_pengguna").Where("id = ?", Objek.IdPengiriman).Limit(1).Take(&pengiriman).Error; err == nil {
		IdSeller = pengiriman.IdSeller
		_ = read.WithContext(ctx).Model(&sot_models.AlamatPengguna{}).Select("id_pengguna").Where("id = ?", pengiriman.IdAlamatPengguna).Limit(1).Take(&IdPengguna).Error
	}

	// ðŸ”” Notifikasi Pengguna (Paket telah di-pickup / dibawa lari kurir)
	if IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸš€ Paketmu Sedang Diantar!",
			Pesan:      fmt.Sprintf("Hore! Paketmu telah berhasil di-pickup dari toko penjuall. Lokasi terkini: %s (%s). Pantau pergerakannya yuk!", Objek.Lokasi, Objek.Keterangan),
			Pop:        3.0, // Alert pop-up agar pembeli senang memantau peta
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman": Objek.IdPengiriman, "jejak_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "OPEN_LIVE_TRACKING_PAGE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	// ðŸ”” Notifikasi Seller (Serah terima produk selesai)
	if IdSeller != 0 {
		var NotifSeller = notification_models.NotificationSeller{
			IDSeller:  IdSeller,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ“¦ Paket Sukses Diserahkan",
			Pesan:     fmt.Sprintf("Barang dengan ID Pengiriman #%d telah dibawa oleh kurir dari toko Anda untuk pengantaran rute reguler.", Objek.IdPengiriman),
			Pop:       3.0,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman": Objek.IdPengiriman},
				Special:  map[string]interface{}{"click_action": "MONAGE_ORDER_DELIVERING"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanNonEksIIschedulerPickedUpNonEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatePickedUpPengirimanNonEksIIschedulerPickedUpNonEksUpdatedPublish"
	var Objek sot_models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirNonEksScheduler = cass_models.BidKurirNonEksScheduler{
		ID:           Objek.ID,
		IdBid:        Objek.IdBid,
		IdKurir:      Objek.IdKurir,
		Urutan:       Objek.Urutan,
		IdPengiriman: Objek.IdPengiriman,
		Status:       Objek.Status,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Scheduler Kurir
	if Objek.IdKurir != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.IdKurir,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Sync Scheduler Picked Up",
			Pesan:     fmt.Sprintf("Manifest scheduler #%d diperbarui ke status %s.", Objek.ID, Objek.Status),
			Pop:       0,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"scheduler_id": Objek.ID, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanNonEksIIpengirimanPickedUpNonEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatePickedUpPengirimanNonEksIIpengirimanPickedUpNonEksUpdatedPublish"
	var Objek sot_models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.Pengiriman = cass_models.Pengiriman{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatPengguna:  Objek.IdAlamatPengguna,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var IdPengguna int64 = 0
	_ = read.WithContext(ctx).Model(&sot_models.AlamatPengguna{}).Select("id_pengguna").Where("id = ?", Objek.IdAlamatPengguna).Limit(1).Take(&IdPengguna).Error

	// ðŸ”” Notifikasi Pengguna
	if IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸ“¦ Pengiriman Di-update: Picked Up",
			Pesan:      fmt.Sprintf("Status logistik pengiriman #%d saat ini: %s.", Objek.ID, Objek.Status),
			Pop:        3.0,
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman": Objek.ID, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "TRACK_DELIVERY"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanNonEksIItransaksiPickedUpNonEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdatePickedUpPengirimanNonEksIItransaksiPickedUpNonEksUpdatedPublish"
	var Objek sot_models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.Transaksi = cass_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
		UpdatedAt:           Objek.UpdatedAt,
		DeletedAt:           Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Transaksi = se_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
	}

	if task_info, err := se_index.TransaksiIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data transaksi ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data transaksi ke search engine dengan info: %s ", task_info.IndexUID)
	}

	// ðŸ”” Notifikasi Transaksi Pembeli
	var NotifPengguna = notification_models.NotificationPengguna{
		IDPengguna: Objek.IdPengguna,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸ›’ Pesananmu Telah Di-pickup!",
		Pesan:      fmt.Sprintf("Transaksi #%s sudah dikonfirmasi bawa oleh kurir. Bersiap menerima kedatangan paketmu ya!", Objek.KodeOrderSistem),
		Pop:        3.0,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_transaksi": Objek.ID, "kode_order": Objek.KodeOrderSistem, "status": Objek.Status},
			Special:  map[string]interface{}{"click_action": "OPEN_TRANSACTION_DETAIL"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)

	// ðŸ”” Notifikasi Transaksi Seller
	var NotifSeller = notification_models.NotificationSeller{
		IDSeller:  int64(Objek.IdSeller),
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ’¸ Pesanan Telah Berpindah Tangan",
		Pesan:     fmt.Sprintf("Order #%s sukses dipickup. Status transaksi berganti menjadi %s.", Objek.KodeOrderSistem, Objek.Status),
		Pop:       3.0,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_transaksi": Objek.ID, "kode_order": Objek.KodeOrderSistem},
			Special:  map[string]interface{}{"click_action": "MONAGE_ORDER_VIEW"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// ===================================================================================================
// FASE 2: ON THE WAY / DELIVERING DEVELOPMENTS
// ===================================================================================================

func UpdateKirimPengirimanNonEksIIbidKurirPengirimanNonEksSchedulerUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateKirimPengirimanNonEksIIbidKurirPengirimanNonEksSchedulerUpdatedPublish"
	var Objek sot_models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirNonEksScheduler = cass_models.BidKurirNonEksScheduler{
		ID:           Objek.ID,
		IdBid:        Objek.IdBid,
		IdKurir:      Objek.IdKurir,
		Urutan:       Objek.Urutan,
		IdPengiriman: Objek.IdPengiriman,
		Status:       Objek.Status,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanNonEksIIpengirimanPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateKirimPengirimanNonEksIIpengirimanPengirimanUpdatedPublish"
	var Objek sot_models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.Pengiriman = cass_models.Pengiriman{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatPengguna:  Objek.IdAlamatPengguna,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var IdPengguna int64 = 0
	_ = read.WithContext(ctx).Model(&sot_models.AlamatPengguna{}).Select("id_pengguna").Where("id = ?", Objek.IdAlamatPengguna).Limit(1).Take(&IdPengguna).Error

	// ðŸ”” Notifikasi Pengguna (Kurir Sedang OTW Menuju Lokasimu)
	if IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸ›µ Kurir Sedang Menuju Rumahmu!",
			Pesan:      fmt.Sprintf("Paket dengan manifest #%d saat ini berstatus [%s]. Kurir internal kami sedang bergegas menuju alamat pengantaran Anda.", Objek.ID, Objek.Status),
			Pop:        3.0, // Alert aktif agar pembeli stand by di lokasi
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman": Objek.ID, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "TRACK_LIVE_MAP_COURIER"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanNonEksIIjejakpengirimanPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateKirimPengirimanNonEksIIjejakpengirimanPengirimanUpdatedPublish"
	var Objek sot_models.JejakPengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengiriman = cass_models.JejakPengiriman{
		ID:           Objek.ID,
		IdPengiriman: Objek.IdPengiriman,
		Lokasi:       Objek.Lokasi,
		Keterangan:   Objek.Keterangan,
		Latitude:     Objek.Latitude,
		Longtitude:   Objek.Longtitude,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateUpdateInformasiPerjalananPengirimanNonEks(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateInformasiPerjalananPengirimanNonEks"
	var Objek sot_models.JejakPengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengiriman = cass_models.JejakPengiriman{
		ID:           Objek.ID,
		IdPengiriman: Objek.IdPengiriman,
		Lokasi:       Objek.Lokasi,
		Keterangan:   Objek.Keterangan,
		Latitude:     Objek.Latitude,
		Longtitude:   Objek.Longtitude,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var IdPengguna int64 = 0
	var pengiriman sot_models.Pengiriman
	if err := read.WithContext(ctx).Model(&sot_models.Pengiriman{}).Select("id_alamat_pengguna").Where("id = ?", Objek.IdPengiriman).Limit(1).Take(&pengiriman).Error; err == nil {
		_ = read.WithContext(ctx).Model(&sot_models.AlamatPengguna{}).Select("id_pengguna").Where("id = ?", pengiriman.IdAlamatPengguna).Limit(1).Take(&IdPengguna).Error
	}

	// ðŸ”” Notifikasi Pengguna (Update Transit Perjalanan)
	if IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸ“ Paketmu Sedang Transit",
			Pesan:      fmt.Sprintf("Update terbaru nih! Paketmu sekarang berada di: %s (%s). Selangkah lebih dekat ke tempatmu!", Objek.Lokasi, Objek.Keterangan),
			Pop:        0, // Ga perlu pop up mengganggu, cukup muncul di tray notification
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman": Objek.IdPengiriman, "jejak_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "TRACK_DELIVERY"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteSampaiPengirimanNonEksIIbidKurirNonEksDeletePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteSampaiPengirimanNonEksIIbidKurirNonEksDeletePublish"
	var Objek sot_models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirNonEksScheduler = cass_models.BidKurirNonEksScheduler{
		ID:           Objek.ID,
		IdBid:        Objek.IdBid,
		IdKurir:      Objek.IdKurir,
		Urutan:       Objek.Urutan,
		IdPengiriman: Objek.IdPengiriman,
		Status:       Objek.Status,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan log hapus ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Scheduler Kurir (Tugas dibersihkan dari manifest kerja aktif kurir)
	if Objek.IdKurir != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.IdKurir,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Manifest Cleared",
			Pesan:     fmt.Sprintf("Tugas pengiriman #%d telah diselesaikan dan dihapus dari manifest aktif Anda.", Objek.IdPengiriman),
			Pop:       0,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"scheduler_id": Objek.ID, "status": "CLEARED"},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanNonEksIIpengirimanSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateSampaiPengirimanNonEksIIpengirimanSampaiUpdatedPublish"
	var Objek sot_models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.Pengiriman = cass_models.Pengiriman{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatPengguna:  Objek.IdAlamatPengguna,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Cari ID Pengguna lewat Alamat Pengguna untuk target notifikasi pembeli
	var IdPengguna int64 = 0
	_ = read.WithContext(ctx).Model(&sot_models.AlamatPengguna{}).Select("id_pengguna").Where("id = ?", Objek.IdAlamatPengguna).Limit(1).Take(&IdPengguna).Error

	// ðŸ”” 1. Notifikasi ke Pembeli (Bujuk cek depan rumah & pencet tombol selesai)
	if IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸŽ‰ Paketmu Sudah Sampai Tujuan!",
			Pesan:      fmt.Sprintf("Hore! Pengiriman #%d statusnya telah [%s]. Buruan cek depan rumah! Kalau isinya aman, yuk klik 'Selesai' dan beri rating kurir serta tokomu ya!", Objek.ID, Objek.Status),
			Pop:        5.0, // Suara & getar kencang karena barang yang ditunggu sudah mendarat
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman": Objek.ID, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "OPEN_ORDER_DETAIL_PAGE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	// ðŸ”” 2. Notifikasi ke Penjual / Seller (Kasih tau kalau pesanan beres & siap cair)
	if Objek.IdSeller != 0 {
		var NotifSeller = notification_models.NotificationSeller{
			IDSeller:  Objek.IdSeller,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ’° Pembeli Telah Menerima Paket!",
			Pesan:     fmt.Sprintf("Mantap! Pesanan dengan ID Pengiriman #%d telah sukses diantarkan kurir ke pembeli. Saldo Anda akan segera diteruskan setelah pembeli melakukan konfirmasi.", Objek.ID),
			Pop:       3.0,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman": Objek.ID},
				Special:  map[string]interface{}{"click_action": "MONAGE_ORDER_DELIVERED"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanNonEksIIbidKurirDataSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateSampaiPengirimanNonEksIIbidKurirDataSampaiUpdatedPublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         *Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Slot/Status Bid Kurir (Biar aplikasi kurir auto-refresh kuota)
	if Objek.IdKurir != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.IdKurir,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Kuota Bid Diperbarui",
			Pesan:     fmt.Sprintf("Slot kerja berhasil disesuaikan. Sisa slot aktif Anda saat ini: %d.", Objek.SlotTersisa),
			Pop:       0, // Silent refresh di background
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"bid_data_id": Objek.ID, "slot_tersisa": Objek.SlotTersisa},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_CAPACITY"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanNonEksIIjejakPengirimanSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateSampaiPengirimanNonEksIIjejakPengirimanSampaiUpdatedPublish"
	var Objek sot_models.JejakPengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengiriman = cass_models.JejakPengiriman{
		ID:           Objek.ID,
		IdPengiriman: Objek.IdPengiriman,
		Lokasi:       Objek.Lokasi,
		Keterangan:   Objek.Keterangan,
		Latitude:     Objek.Latitude,
		Longtitude:   Objek.Longtitude,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
		DeletedAt:    Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Tracing data Pengiriman untuk ambil IdPengguna & IdSeller (MULTI-ENTITY)
	var (
		IdPengguna int64 = 0
		IdSeller   int64 = 0
	)
	var pengiriman sot_models.Pengiriman
	if err := read.WithContext(ctx).Model(&sot_models.Pengiriman{}).Select("id_seller, id_alamat_pengguna").Where("id = ?", Objek.IdPengiriman).Limit(1).Take(&pengiriman).Error; err == nil {
		IdSeller = pengiriman.IdSeller
		_ = read.WithContext(ctx).Model(&sot_models.AlamatPengguna{}).Select("id_pengguna").Where("id = ?", pengiriman.IdAlamatPengguna).Limit(1).Take(&IdPengguna).Error
	}

	// ðŸ”” 1. Notifikasi ke Pembeli (Jejak Akhir Penerimaan + Ajakan Review)
	if IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸ“¦ Selesai! Paketmu Telah Diterima",
			Pesan:      fmt.Sprintf("Yeay, tracking log mencatat pesananmu sudah sampai di %s (%s). Suka dengan pelayanannya? Yuk, kasih review terbaikmu sekarang!", Objek.Lokasi, Objek.Keterangan),
			Pop:        3.0,
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman": Objek.IdPengiriman, "jejak_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "LEAVE_A_REVIEW"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	// ðŸ”” 2. Notifikasi ke Penjual / Seller (Log Validasi Akhir Pengiriman Beres)
	if IdSeller != 0 {
		var NotifSeller = notification_models.NotificationSeller{
			IDSeller:  IdSeller,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "âœ… Log Pengiriman Selesai",
			Pesan:     fmt.Sprintf("Sistem mencatat riwayat perjalanan akhir untuk pengiriman #%d telah terverifikasi sukses di lokasi: %s.", Objek.IdPengiriman, Objek.Lokasi),
			Pop:       0, // Log sekunder, tidak perlu getar berlebih
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman": Objek.IdPengiriman},
				Special:  map[string]interface{}{"click_action": "MONAGE_ORDER_DELIVERED"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// INFOO: INI CONTOH YANG ADA NOTIFIKASI
func UpdateSampaiPengirimanNonEksIItransaksiSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdateSampaiPengirimanNonEksIItransaksiSampaiUpdatedPublish"
	var Objek sot_models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	// 1. Mapping ke Cassandra Model
	var ObjekCass cass_models.Transaksi = cass_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
		UpdatedAt:           Objek.UpdatedAt,
		DeletedAt:           Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	// 2. Update ke Cassandra Sot Replica
	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	// 3. Catat ke Cassandra Historical DB
	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 4. Sinkronisasi ke Search Engine (Meilisearch)
	var ObjekSearchEngine se_models.Transaksi = se_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
	}

	if task_info, err := se_index.TransaksiIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data transaksi ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data transaksi ke search engine dengan info: %s ", task_info.IndexUID)
	}

	// ðŸ”” 1. Notifikasi ke Pembeli (Transaksi Selesai Resmi)
	if Objek.IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: Objek.IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸ›ï¸ Transaksi Selesai!",
			Pesan:      fmt.Sprintf("Pesananmu dengan kode %s telah selesai diproses. Terima kasih sudah berbelanja! Ditunggu orderan berikutnya, ya.", Objek.KodeOrderSistem),
			Pop:        3.0,
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_transaksi": Objek.ID, "kode_order": Objek.KodeOrderSistem},
				Special:  map[string]interface{}{"click_action": "OPEN_TRANSACTION_DETAIL"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	// ðŸ”” 2. Notifikasi ke Penjual (Transaksi Selesai, Siap Terima Dana)
	if Objek.IdSeller != 0 {
		var NotifSeller = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.IdSeller),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ“ˆ Penjualan Sukses Selesai",
			Pesan:     fmt.Sprintf("Selamat! Transaksi %s senilai Rp %d dari pembeli telah selesai sepenuhnya. Data keuangan Anda sedang disinkronkan oleh sistem.", Objek.KodeOrderSistem, Objek.Total),
			Pop:       3.0,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_transaksi": Objek.ID, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "MONAGE_ORDER_COMPLETED"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateSampaiPengirimanNonEksIIpayOutSellerCreatePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateSampaiPengirimanNonEksIIpayOutSellerCreatePublish"
	var Objek sot_models.PayOutSeller

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.PayOutSeller = cass_models.PayOutSeller{
		ID:               Objek.ID,
		IdSeller:         Objek.IdSeller,
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

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Notifikasi ke Penjual (Duit Hasil Jualan Cair Ke Rekening!)
	if Objek.IdSeller != 0 {
		var NotifSeller = notification_models.NotificationSeller{
			IDSeller:  Objek.IdSeller,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ’° Dana Penjualan Telah Ditransfer!",
			Pesan:     fmt.Sprintf("Hore! Dana sebesar Rp %d berhasil dicairkan ke rekening %s (%s) Anda dengan status [%s]. Silakan cek mutasi rekening Anda.", Objek.Amount, Objek.BankCode, Objek.AccountNumber, Objek.Status),
			Pop:       5.0, // Urusan duit wajib pop-up paling kenceng biar seller puas
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"payout_id": Objek.ID, "amount": Objek.Amount, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "OPEN_FINANCIAL_LOG"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateSampaiPengirimanNonEksIIpayOutKurirCreatePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateSampaiPengirimanNonEksIIpayOutKurirCreatePublish"
	var Objek sot_models.PayOutKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.PayOutKurir = cass_models.PayOutKurir{
		ID:               Objek.ID,
		IdKurir:          Objek.IdKurir,
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

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Notifikasi ke Kurir (Gaji / Insentif Rute Cair)
	if Objek.IdKurir != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.IdKurir,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ’¸ Ongkir/Insentif Pengantaran Cair!",
			Pesan:     fmt.Sprintf("Kerja bagus! Pendapatan sebesar Rp %d hasil pengantaran sukses ditransfer ke rekening bank %s Anda. Status: [%s]. Cek dompet digitalmu sekarang!", Objek.Amount, Objek.BankCode, Objek.Status),
			Pop:       5.0, // Kurir paling senang dapat notif duit masuk, kasih pop-up maksimal!
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"payout_id": Objek.ID, "amount": Objek.Amount, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "OPEN_WALLET_BALANCE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
func CreatePickedUpPengirimanEksIIjejakPengirimanEksCreatePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreatePickedUpPengirimanEksIIjejakPengirimanEksCreatePublish"
	var Objek sot_models.JejakPengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengirimanEkspedisi = cass_models.JejakPengirimanEkspedisi{
		ID:                    Objek.ID,
		IdPengirimanEkspedisi: Objek.IdPengirimanEkspedisi,
		Lokasi:                Objek.Lokasi,
		Keterangan:            Objek.Keterangan,
		Latitude:              Objek.Latitude,
		Longitude:             Objek.Longitude,
		CreatedAt:             Objek.CreatedAt,
		UpdatedAt:             Objek.UpdatedAt,
		DeletedAt:             Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Ambil IdTransaksi dari Pengiriman Ekspedisi untuk dapetin info Pembeli
	var IdPengguna int64 = 0
	var pe sot_models.PengirimanEkspedisi
	if err := read.WithContext(ctx).Model(&sot_models.PengirimanEkspedisi{}).Select("id_transaksi").Where("id = ?", Objek.IdPengirimanEkspedisi).Limit(1).Take(&pe).Error; err == nil {
		_ = read.WithContext(ctx).Model(&sot_models.Transaksi{}).Select("id_pengguna").Where("id = ?", pe.IdTransaksi).Limit(1).Take(&IdPengguna).Error
	}

	// ðŸ”” Notifikasi Pembeli: Jejak pertama logistik eksternal tercatat
	if IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸšš Paketmu Mulas Bergerak!",
			Pesan:      fmt.Sprintf("Status logistik eksternal baru: paketmu saat ini berada di %s (%s).", Objek.Lokasi, Objek.Keterangan),
			Pop:        1.0,
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman_eks": Objek.IdPengirimanEkspedisi},
				Special:  map[string]interface{}{"click_action": "OPEN_TRACKING_PAGE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanEksIIschedulerEksPickedUpUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatePickedUpPengirimanEksIIschedulerEksPickedUpUpdatedPublish"
	var Objek sot_models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirEksScheduler = cass_models.BidKurirEksScheduler{
		ID:              Objek.ID,
		IdBid:           Objek.IdBid,
		IdKurir:         Objek.IdKurir,
		Urutan:          Objek.Urutan,
		IdPengirimanEks: Objek.IdPengirimanEks,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Kurir Ekspedisi (Update status antrean rute pick-up kurir internal depo)
	if Objek.IdKurir != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.IdKurir,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Status Rute Diperbarui",
			Pesan:     fmt.Sprintf("Penjadwalan manifest #%d diubah menjadi [%s].", Objek.IdPengirimanEks, Objek.Status),
			Pop:       0,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"scheduler_id": Objek.ID, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanEksIIpengirimanEksPickedUpUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatePickedUpPengirimanEksIIpengirimanEksPickedUpUpdatedPublish"
	var Objek sot_models.PengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.PengirimanEkspedisi = cass_models.PengirimanEkspedisi{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatEkspedisi: Objek.IdAlamatEkspedisi,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Ambil IdPengguna via Transaksi
	var IdPengguna int64 = 0
	_ = read.WithContext(ctx).Model(&sot_models.Transaksi{}).Select("id_pengguna").Where("id = ?", Objek.IdTransaksi).Limit(1).Take(&IdPengguna).Error

	// ðŸ”” Notifikasi ke Penjual: Barang sukses diserahkan ke Ekspedisi
	if Objek.IdSeller != 0 {
		var NotifSeller = notification_models.NotificationSeller{
			IDSeller:  Objek.IdSeller,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ“¦ Serah Terima Ekspedisi Berhasil",
			Pesan:     fmt.Sprintf("Paket dari pengiriman #%d sukses di-pickup oleh pihak mitra ekspedisi. Tanggung jawab pengiriman kini beralih ke kurir logistik eksternal.", Objek.ID),
			Pop:       2.0,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman_eks": Objek.ID},
				Special:  map[string]interface{}{"click_action": "MONAGE_SHIPPING_STATUS"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	// ðŸ”” Notifikasi ke Pembeli: Status pengiriman total berubah jadi PICKED_UP
	if IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸ“¦ Paket Telah Di-pickup Kurir",
			Pesan:      fmt.Sprintf("Asik! Paketmu untuk transaksi #%d sudah diterima oleh pihak ekspedisi rekanan dan siap diberangkatkan ke kota tujuan.", Objek.IdTransaksi),
			Pop:        3.0,
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_transaksi": Objek.IdTransaksi},
				Special:  map[string]interface{}{"click_action": "OPEN_ORDER_TRACKING"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanEksIItransaksiPickedUpUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdatePickedUpPengirimanEksIItransaksiPickedUpUpdatedPublish"
	var Objek sot_models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.Transaksi = cass_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
		UpdatedAt:           Objek.UpdatedAt,
		DeletedAt:           Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Transaksi = se_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
	}

	if task_info, err := se_index.TransaksiIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data transaksi ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data transaksi ke search engine dengan info: %s ", task_info.IndexUID)
	}

	// ðŸ”” Notifikasi Pembeli: Transaksi resmi berganti status utama di dashboard mereka
	if Objek.IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: Objek.IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸ’¼ Pesanan Sedang Dikirim",
			Pesan:      fmt.Sprintf("Transaksi #%s sudah diproses kurir ekspedisi. Pantau terus halaman tracking untuk update real-time logistiknya.", Objek.KodeOrderSistem),
			Pop:        2.0,
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_transaksi": Objek.ID, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "OPEN_TRANSACTION_DETAIL"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanEksIIschedulerPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateKirimPengirimanEksIIschedulerPengirimanUpdatedPublish"
	var Objek sot_models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirEksScheduler = cass_models.BidKurirEksScheduler{
		ID:              Objek.ID,
		IdBid:           Objek.IdBid,
		IdKurir:         Objek.IdKurir,
		Urutan:          Objek.Urutan,
		IdPengirimanEks: Objek.IdPengirimanEks,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Kurir: Status transit antar depo berubah
	if Objek.IdKurir != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.IdKurir,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Transit Log Updated",
			Pesan:     fmt.Sprintf("Jadwal pengiriman ekspedisi #%d status updated to [%s].", Objek.IdPengirimanEks, Objek.Status),
			Pop:       0,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"scheduler_id": Objek.ID, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_ROUTING"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanEksIIpengirimanEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateKirimPengirimanEksIIpengirimanEksUpdatedPublish"
	var Objek sot_models.PengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.PengirimanEkspedisi = cass_models.PengirimanEkspedisi{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatEkspedisi: Objek.IdAlamatEkspedisi,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Ambil IdPengguna lewat Transaksi
	var IdPengguna int64 = 0
	_ = read.WithContext(ctx).Model(&sot_models.Transaksi{}).Select("id_pengguna").Where("id = ?", Objek.IdTransaksi).Limit(1).Take(&IdPengguna).Error

	// ðŸ”” Notifikasi Pembeli: Konfirmasi kalau paket bener-bener on the way antar hub logistik
	if IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸš€ Paketmu Sedang Dalam Perjalanan!",
			Pesan:      fmt.Sprintf("Pengiriman eksternal #%d diperbarui menjadi [%s]. Armada kurir sedang mengarah ke wilayah transit terdekat.", Objek.ID, Objek.Status),
			Pop:        2.0,
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman_eks": Objek.ID, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "OPEN_TRACKING_PAGE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanEksIIpengirimanPengirimanEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateKirimPengirimanEksIIpengirimanPengirimanEksUpdatedPublish"
	var Objek sot_models.JejakPengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengirimanEkspedisi = cass_models.JejakPengirimanEkspedisi{
		ID:                    Objek.ID,
		IdPengirimanEkspedisi: Objek.IdPengirimanEkspedisi,
		Lokasi:                Objek.Lokasi,
		Keterangan:            Objek.Keterangan,
		Latitude:              Objek.Latitude,
		Longitude:             Objek.Longitude,
		CreatedAt:             Objek.CreatedAt,
		UpdatedAt:             Objek.UpdatedAt,
		DeletedAt:             Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Ambil IdPengguna lewat PengirimanEkspedisi -> Transaksi
	var IdPengguna int64 = 0
	var pe sot_models.PengirimanEkspedisi
	if err := read.WithContext(ctx).Model(&sot_models.PengirimanEkspedisi{}).Select("id_transaksi").Where("id = ?", Objek.IdPengirimanEkspedisi).Limit(1).Take(&pe).Error; err == nil {
		_ = read.WithContext(ctx).Model(&sot_models.Transaksi{}).Select("id_pengguna").Where("id = ?", pe.IdTransaksi).Limit(1).Take(&IdPengguna).Error
	}

	// ðŸ”” Notifikasi Pembeli: Setiap log detail dari manifest pihak ketiga masuk
	if IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸ“ Update Lokasi Paket",
			Pesan:      fmt.Sprintf("Paket terdeteksi tiba di %s: %s.", Objek.Lokasi, Objek.Keterangan),
			Pop:        1.0, // Log perjalanan berkala cukup getar halus
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"jejak_id": Objek.ID, "id_pengiriman_eks": Objek.IdPengirimanEkspedisi},
				Special:  map[string]interface{}{"click_action": "OPEN_TRACKING_PAGE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
func UpdateInformasiPerjalananPengirimanEks(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateInformasiPerjalananPengirimanEks"
	var Objek sot_models.JejakPengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengirimanEkspedisi = cass_models.JejakPengirimanEkspedisi{
		ID:                    Objek.ID,
		IdPengirimanEkspedisi: Objek.IdPengirimanEkspedisi,
		Lokasi:                Objek.Lokasi,
		Keterangan:            Objek.Keterangan,
		Latitude:              Objek.Latitude,
		Longitude:             Objek.Longitude,
		CreatedAt:             Objek.CreatedAt,
		UpdatedAt:             Objek.UpdatedAt,
		DeletedAt:             Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Ambil IdTransaksi dari Pengiriman Ekspedisi untuk melacak Pembeli
	var IdPengguna int64 = 0
	var pe sot_models.PengirimanEkspedisi
	if err := read.WithContext(ctx).Model(&sot_models.PengirimanEkspedisi{}).Select("id_transaksi").Where("id = ?", Objek.IdPengirimanEkspedisi).Limit(1).Take(&pe).Error; err == nil {
		_ = read.WithContext(ctx).Model(&sot_models.Transaksi{}).Select("id_pengguna").Where("id = ?", pe.IdTransaksi).Limit(1).Take(&IdPengguna).Error
	}

	// ðŸ”” Notifikasi ke Pembeli: Log informasi rute diperbarui
	if IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸ“ Pembaruan Informasi Perjalanan",
			Pesan:      fmt.Sprintf("Ada pembaruan manifes kurir: Paket Anda terdeteksi di %s (%s).", Objek.Lokasi, Objek.Keterangan),
			Pop:        1.0,
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman_eks": Objek.IdPengirimanEkspedisi},
				Special:  map[string]interface{}{"click_action": "OPEN_TRACKING_PAGE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteSampaipengirimanEksIIbidKurirEksDeletePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteSampaipengirimanEksIIbidKurirEksDeletePublish"
	var Objek sot_models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirEksScheduler = cass_models.BidKurirEksScheduler{
		ID:              Objek.ID,
		IdBid:           Objek.IdBid,
		IdKurir:         Objek.IdKurir,
		Urutan:          Objek.Urutan,
		IdPengirimanEks: Objek.IdPengirimanEks,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan log hapus ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Kurir: Bersihkan jadwal manifest dari antrean lokal aplikasi kurir
	if Objek.IdKurir != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.IdKurir,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ—‘ï¸ Penjadwalan Selesai",
			Pesan:     fmt.Sprintf("Manifest rute #%d telah ditutup dan dihapus dari daftar tugas aktif Anda.", Objek.IdPengirimanEks),
			Pop:       0, // Pembersihan antrean tidak perlu mengganggu layar HP kurir
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"scheduler_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_REMOVE_QUEUE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIpengirimanSampaiEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateSampaiPengirimanEksIIpengirimanSampaiEksUpdatedPublish"
	var Objek sot_models.PengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.PengirimanEkspedisi = cass_models.PengirimanEkspedisi{
		ID:                Objek.ID,
		IdTransaksi:       Objek.IdTransaksi,
		IdSeller:          Objek.IdSeller,
		IdAlamatGudang:    Objek.IdAlamatGudang,
		IdAlamatEkspedisi: Objek.IdAlamatEkspedisi,
		IdKurir:           Objek.IdKurir,
		BeratBarang:       Objek.BeratBarang,
		KendaraanRequired: Objek.KendaraanRequired,
		JenisPengiriman:   Objek.JenisPengiriman,
		JarakTempuh:       Objek.JarakTempuh,
		KurirPaid:         Objek.KurirPaid,
		Status:            Objek.Status,
		CreatedAt:         Objek.CreatedAt,
		UpdatedAt:         Objek.UpdatedAt,
		DeletedAt:         Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Ambil IdPengguna via Transaksi
	var IdPengguna int64 = 0
	_ = read.WithContext(ctx).Model(&sot_models.Transaksi{}).Select("id_pengguna").Where("id = ?", Objek.IdTransaksi).Limit(1).Take(&IdPengguna).Error

	// ðŸ”” 1. Notifikasi ke Pembeli: Paket mendarat dengan selamat
	if IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸŽ‰ Paketmu Sudah Sampai!",
			Pesan:      fmt.Sprintf("Hore! Pengiriman #%d untuk transaksi Anda telah berhasil dikirimkan ke alamat tujuan. Silakan periksa paket Anda.", Objek.ID),
			Pop:        4.0, // Penting banget agar user langsung ngecek pagar rumah
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_transaksi": Objek.IdTransaksi, "id_pengiriman_eks": Objek.ID},
				Special:  map[string]interface{}{"click_action": "OPEN_TRANSACTION_DETAIL"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	// ðŸ”” 2. Notifikasi ke Penjual: Konfirmasi barang sampai tujuan
	if Objek.IdSeller != 0 {
		var NotifSeller = notification_models.NotificationSeller{
			IDSeller:  Objek.IdSeller,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ Pengiriman Selesai",
			Pesan:     fmt.Sprintf("Paket dari pesanan #%d telah sukses diserahkan ke pembeli oleh mitra ekspedisi. Menunggu konfirmasi akhir dari pembeli.", Objek.IdTransaksi),
			Pop:       2.0,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_transaksi": Objek.IdTransaksi},
				Special:  map[string]interface{}{"click_action": "MONAGE_ORDER_COMPLETED"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIbidKurirDataEksSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateSampaiPengirimanEksIIbidKurirDataEksSampaiUpdatedPublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         *Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Kurir: Sinkronisasi sisa slot muatan truk/motor logistik depo eksternal
	if Objek.IdKurir != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.IdKurir,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Slot Kuota Diperbarui",
			Pesan:     fmt.Sprintf("Sisa slot muatan kurir diperbarui menjadi %d kg untuk wilayah %s.", Objek.SlotTersisa, Objek.Kota),
			Pop:       0, // Sync background data aplikasi kurir
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"bid_kurir_data_id": Objek.ID, "slot_tersisa": Objek.SlotTersisa},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_CAPACITY"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIjejakPengirimanEksSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateSampaiPengirimanEksIIjejakPengirimanEksSampaiUpdatedPublish"
	var Objek sot_models.JejakPengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.JejakPengirimanEkspedisi = cass_models.JejakPengirimanEkspedisi{
		ID:                    Objek.ID,
		IdPengirimanEkspedisi: Objek.IdPengirimanEkspedisi,
		Lokasi:                Objek.Lokasi,
		Keterangan:            Objek.Keterangan,
		Latitude:              Objek.Latitude,
		Longitude:             Objek.Longitude,
		CreatedAt:             Objek.CreatedAt,
		UpdatedAt:             Objek.UpdatedAt,
		DeletedAt:             Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Ambil IdTransaksi dari Pengiriman Ekspedisi untuk melacak Pembeli
	var IdPengguna int64 = 0
	var pe sot_models.PengirimanEkspedisi
	if err := read.WithContext(ctx).Model(&sot_models.PengirimanEkspedisi{}).Select("id_transaksi").Where("id = ?", Objek.IdPengirimanEkspedisi).Limit(1).Take(&pe).Error; err == nil {
		_ = read.WithContext(ctx).Model(&sot_models.Transaksi{}).Select("id_pengguna").Where("id = ?", pe.IdTransaksi).Limit(1).Take(&IdPengguna).Error
	}

	// ðŸ”” Notifikasi Pembeli: Manifes manifes logistik menandakan barang telah diterima di titik akhir
	if IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸ“ Logistik Ekspedisi Selesai",
			Pesan:      fmt.Sprintf("Titik akhir terdata: Paket sudah tiba di %s (%s).", Objek.Lokasi, Objek.Keterangan),
			Pop:        1.0,
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_pengiriman_eks": Objek.IdPengirimanEkspedisi},
				Special:  map[string]interface{}{"click_action": "OPEN_TRACKING_PAGE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIItransaksiSampaiEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdateSampaiPengirimanEksIItransaksiSampaiEksUpdatedPublish"
	var Objek sot_models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.Transaksi = cass_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
		UpdatedAt:           Objek.UpdatedAt,
		DeletedAt:           Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Transaksi = se_models.Transaksi{
		ID:                  Objek.ID,
		IdPengguna:          Objek.IdPengguna,
		IdSeller:            Objek.IdSeller,
		IdBarangInduk:       Objek.IdBarangInduk,
		IdKategoriBarang:    Objek.IdKategoriBarang,
		IdAlamatPengguna:    Objek.IdAlamatPengguna,
		IdAlamatGudang:      Objek.IdAlamatGudang,
		IdAlamatEkspedisi:   Objek.IdAlamatEkspedisi,
		IdPembayaran:        Objek.IdPembayaran,
		KendaraanPengiriman: Objek.KendaraanPengiriman,
		JenisPengiriman:     Objek.JenisPengiriman,
		JarakTempuh:         Objek.JarakTempuh,
		BeratTotalKg:        Objek.BeratTotalKg,
		KodeOrderSistem:     Objek.KodeOrderSistem,
		KodeResiEkspedisi:   Objek.KodeResiEkspedisi,
		Status:              Objek.Status,
		DibatalkanOleh:      Objek.DibatalkanOleh,
		Catatan:             Objek.Catatan,
		KuantitasBarang:     Objek.KuantitasBarang,
		IsEkspedisi:         Objek.IsEkspedisi,
		SellerPaid:          Objek.SellerPaid,
		KurirPaid:           Objek.KurirPaid,
		EkspedisiPaid:       Objek.EkspedisiPaid,
		Total:               Objek.Total,
		Reviewed:            Objek.Reviewed,
		CreatedAt:           Objek.CreatedAt,
	}

	if task_info, err := se_index.TransaksiIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data transaksi ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data transaksi ke search engine dengan info: %s ", task_info.IndexUID)
	}

	// ðŸ”” Notifikasi Pembeli: Transaksi utama selesai, dorong pembeli untuk klik selesaikan / review produk
	if Objek.IdPengguna != 0 {
		var NotifPengguna = notification_models.NotificationPengguna{
			IDPengguna: Objek.IdPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸ“¦ Paket Diterima! Yuk Konfirmasi",
			Pesan:      fmt.Sprintf("Pesanan #%s dinyatakan sampai tujuan oleh sistem ekspedisi eksternal. Jangan lupa konfirmasi terima barang dan beri ulasan ya!", Objek.KodeOrderSistem),
			Pop:        3.0,
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 5).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"id_transaksi": Objek.ID},
				Special:  map[string]interface{}{"click_action": "OPEN_REVIEW_PAGE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateSampaiPengirimanEksIIpayOutKurirEksCreatePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateSampaiPengirimanEksIIpayOutKurirEksCreatePublish"
	var Objek sot_models.PayOutKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.PayOutKurir = cass_models.PayOutKurir{
		ID:               Objek.ID,
		IdKurir:          Objek.IdKurir,
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

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Notifikasi Finansial Kurir: Pengiriman sukses berbuah saldo/ongkir cair ke rekening
	if Objek.IdKurir != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.IdKurir,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ’° Dana Pengiriman Telah Cair!",
			Pesan:     fmt.Sprintf("Selamat! Saldo komisi sebesar Rp %v sukses ditransfer ke rekening %s Anda.", Objek.Amount, Objek.BankCode),
			Pop:       4.0, // Push notif keuangan penting banget dikasih pop tinggi
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"payout_id": Objek.ID, "amount": Objek.Amount, "status": Objek.Status},
				Special:  map[string]interface{}{"click_action": "OPEN_WALLET_HISTORY"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIkurirUpdatedSampaiEksPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services string = "UpdateSampaiPengirimanEksIIkurirUpdatedSampaiEksPublish"
	var Objek sot_models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
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
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.Kurir = se_models.Kurir{
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
	}

	if task_info, err := se_index.KurirIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data kurir ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data kurir ke search engine dengan info: %s ", task_info.IndexUID)
	}

	if err := cache_db_function.UpdateSessionData(ctx, *rds_session, cache_db_session.GetSessionKey(&Objek), Objek); err != nil {
		return fmt.Errorf("gagal memperbarui data kurir di cache %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Kurir: Sinkronisasi pembaruan profil / performa rating kurir secara berkala
	if Objek.ID != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.ID,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Sinkronisasi Profil Berhasil",
			Pesan:     "Data statistik performa dan rating kurir Anda berhasil diperbarui.",
			Pop:       0,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"kurir_id": Objek.ID, "rating_terkini": Objek.Rating},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_PROFILE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteNonaktifkanBidKurirIIbidKurirDataDeletePublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteNonaktifkanBidKurirIIbidKurirDataDeletePublish"
	var Objek sot_models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	var ObjekCass cass_models.BidKurirData = cass_models.BidKurirData{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		JenisPengiriman: Objek.JenisPengiriman,
		Mode:            Objek.Mode,
		Provinsi:        Objek.Provinsi,
		Kota:            Objek.Kota,
		IsEkspedisi:     Objek.IsEkspedisi,
		Alamat:          Objek.Alamat,
		Longitude:       Objek.Longitude,
		Latitude:        Objek.Latitude,
		MaxKg:           Objek.MaxKg,
		SlotTersisa:     Objek.SlotTersisa,
		Dimulai:         Objek.Dimulai,
		Selesai:         *Objek.Selesai,
		JenisKendaraan:  Objek.JenisKendaraan,
		Status:          Objek.Status,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
		DeletedAt:       Objek.DeletedAt.Time,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan log hapus ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Kurir: Hapus slot penawaran (Bid) logistik dari perangkat lokal karena masa berlaku/penugasan habis
	if Objek.IdKurir != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.IdKurir,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ—‘ï¸ Alokasi Penjadwalan Ditutup",
			Pesan:     "Slot operasional pengiriman ekspedisi untuk periode ini resmi dinonaktifkan.",
			Activity:  true,
			Inbox:     false,
			Archive:   true,
			Pop:       0, // Silent hapus di aplikasi klien
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"bid_kurir_data_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_REMOVE_SLOT"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateNonaktifkanBidKurirIIkurirNonaktifkanBidUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper, rds_session *redis.Client) error {
	const handle_services string = "UpdateNonaktifkanBidKurirIIkurirNonaktifkanBidUpdatedPublish"
	var Objek sot_models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data dalam %s", handle_services)
	}

	// 1. Update ke Cassandra Sot Replica & Historical
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
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal mengupdate data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// 2. Sinkronisasi ke Search Engine (Meilisearch)
	var ObjekSearchEngine se_models.Kurir = se_models.Kurir{
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
	}

	if task_info, err := se_index.KurirIndex.UpdateDocuments(&ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data kurir ke search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data kurir ke search engine dengan info: %s ", task_info.IndexUID)
	}

	// 3. Update Session Data di Cache Redis
	if err := cache_db_function.UpdateSessionData(ctx, *rds_session, cache_db_session.GetSessionKey(&Objek), Objek); err != nil {
		return fmt.Errorf("gagal memperbarui data kurir di cache %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Kurir: Sinkronisasi penonaktifan status fitur pencarian kerja (Bid Status)
	if Objek.ID != 0 {
		var NotifKurir = notification_models.NotificationKurir{
			IDKurir:   Objek.ID,
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ“´ Sesi Bidding Ditutup",
			Pesan:     "Status pencarian penawaran (bid) Anda saat ini telah dinonaktifkan oleh sistem.",
			Pop:       0, // Jalur background sync, merubah state UI aplikasi kurir tanpa pop up
			Activity:  true,
			Inbox:     false,
			Archive:   true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"kurir_id": Objek.ID, "status_bid": Objek.StatusBid},
				Special:  map[string]interface{}{"click_action": "SILENT_DISABLE_BID_MODE"},
			},
		}
		_ = notification_request.PostToNotification(ctx, NotifKurir, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
