package transaction_pengguna_handle

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/cache"
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

func DeleteCheckoutBarangUser(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteCheckoutBarangUser"
	var Objek sot_models.Keranjang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	var ObjekCass cass_models.Keranjang = cass_models.Keranjang{
		ID:            Objek.ID,
		IdPengguna:    Objek.IdPengguna,
		IdSeller:      Objek.IdSeller,
		IdBarangInduk: Objek.IdBarangInduk,
		IdKategori:    Objek.IdKategori,
		Jumlah:        Objek.Jumlah,
		Status:        Objek.Status,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateLockTransaksiVa(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateLockTransaksiVa"
	var Objek sot_models.Transaksi
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
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
		DeletedAt:           Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal menghapus data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
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
		UpdatedAt:           Objek.UpdatedAt,
		DeletedAt:           &Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.TransaksiIndex.AddDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("Berhasil memasukan data ke dalam search engine dengan antrean UID %s", task_info.IndexUID)
	}

	// Setup teks cuplikan buat notifikasi
	judulVA := "ðŸ’³ Tagihan Baru Menanti Pembayaran"
	pesanVA := fmt.Sprintf("Pesananmu dengan kode order %s telah dikunci. Yuk, segera lakukan pembayaran lewat Virtual Account sebelum batas waktunya habis biar pesanan lu langsung diproses!", Objek.KodeOrderSistem)

	// ==========================================
	// NOTIFIKASI 1: UNTUK PENGGUNA (PEMBELI) - PAKE POP-UP
	// ==========================================
	var NotificationLockVA notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: Objek.IdPengguna,
		Pengirim:   notification_seeders.Sistem,
		Judul:      judulVA,
		Pesan:      pesanVA,
		Activity:   true,
		Inbox:      false,
		Archive:    true,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
		Pop:        4.5, // Muncul di layar pembeli 4.5 detik
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"transaksi_id":      Objek.ID,
				"kode_order_sistem": Objek.KodeOrderSistem,
				"action_type":       "transaction_lock_va",
				"total_tagihan":     Objek.Total,
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_PAYMENT_INSTRUCTION_PAGE",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotificationLockVA, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk); err != nil {
		fmt.Printf("Gagal mengirim notifikasi lock VA ke pengguna %d: %v\n", Objek.IdPengguna, err)
	}

	// ==========================================
	// NOTIFIKASI 2: UNTUK SELLER - SILENT INBOX ONLY (POP: 0)
	// ==========================================
	judulSeller := "ðŸ“¦ Calon Pesanan Baru!"
	pesanSeller := fmt.Sprintf("Pembeli telah mengunci pesanan %s. Kami akan kabari lagi jika pembayaran Virtual Account mereka sudah lunas ya!", Objek.KodeOrderSistem)

	var NotificationSellerVA notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: int64(Objek.IdSeller), // Kirim ke seller terkait
		Pengirim:   notification_seeders.Sistem,
		Judul:      judulSeller,
		Pesan:      pesanSeller,
		Activity:   true,
		Inbox:      false,
		Archive:    true,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 3).Format(time.RFC3339), // Log info keep 3 hari aja cukup
		Pop:        0.0,                                              // ðŸ”¥ Request lu: Pop 0 biar FE tau ini silent, masuk inbox doang!
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"transaksi_id":      Objek.ID,
				"kode_order_sistem": Objek.KodeOrderSistem,
				"action_type":       "seller_incoming_lock_va",
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_SELLER_ORDER_DASHBOARD", // FE buka dashboard list order seller
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotificationSellerVA, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk); err != nil {
		fmt.Printf("Gagal mengirim silent notification lock VA ke seller %d: %v\n", Objek.IdSeller, err)
	}

	fmt.Printf("Berhasil memproses dual-notifikasi (Pop & Silent) Lock Transaksi VA untuk ID %d\n", Objek.ID)

	return nil
}

func CreateLockTransaksiWallet(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateLockTransaksiWallet"
	var Objek sot_models.Transaksi
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
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
		DeletedAt:           Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal menghapus data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
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
		UpdatedAt:           Objek.UpdatedAt,
		DeletedAt:           &Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.TransaksiIndex.AddDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("Berhasil memasukan data ke dalam search engine dengan antrean UID %s", task_info.IndexUID)
	}

	// Notif Pembeli (Pop-up aktif)
	var NotificationWalletUser notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: Objek.IdPengguna,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸ‘› Konfirmasi Pembayaran Dompet",
		Pesan:      fmt.Sprintf("Pesanan %s menunggu konfirmasi saldo digitalmu. Yuk selesaikan pembayaran sekarang agar pesanan langsung diproses!", Objek.KodeOrderSistem),
		Activity:   true,
		Inbox:      false,
		Archive:    true,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
		Pop:        3.0, // 3 detik cukup buat trigger bayar instan
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"transaksi_id":      Objek.ID,
				"kode_order_sistem": Objek.KodeOrderSistem,
				"action_type":       "transaction_lock_wallet",
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_WALLET_PAYMENT_PAGE", // FE arahin ke pin wallet auth
			},
		},
	}
	if err := notification_request.PostToNotification(ctx, NotificationWalletUser, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk); err != nil {
		fmt.Printf("Gagal mengirim notifikasi wallet ke pengguna %d: %v\n", Objek.IdPengguna, err)
	}

	// Notif Seller (Silent - Masuk Inbox)
	var NotificationWalletSeller notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: int64(Objek.IdSeller),
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸ“¦ Calon Pesanan Baru (Wallet)",
		Pesan:      fmt.Sprintf("Pembeli sedang menyelesaikan pembayaran transaksi %s menggunakan dompet digital.", Objek.KodeOrderSistem),
		Activity:   true,
		Inbox:      false,
		Archive:    true,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
		Pop:        0.0, // Silent inbox
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"transaksi_id":      Objek.ID,
				"kode_order_sistem": Objek.KodeOrderSistem,
				"action_type":       "seller_incoming_lock_wallet",
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_SELLER_ORDER_DASHBOARD",
			},
		},
	}
	if err := notification_request.PostToNotification(ctx, NotificationWalletSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk); err != nil {
		fmt.Printf("Gagal mengirim silent notification wallet ke seller %d: %v\n", Objek.IdSeller, err)
	}

	fmt.Printf("Berhasil memproses dual-notifikasi Transaksi Wallet untuk ID %d\n", Objek.ID)

	return nil
}

func CreateLockTransaksiGerai(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateLockTransaksiGerai"
	var Objek sot_models.Transaksi
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
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
		DeletedAt:           Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal menghapus data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
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
		UpdatedAt:           Objek.UpdatedAt,
		DeletedAt:           &Objek.DeletedAt.Time,
	}

	if task_info, err := se_index.TransaksiIndex.AddDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("Berhasil memasukan data ke dalam search engine dengan antrean UID %s", task_info.IndexUID)
	}

	var NotificationGeraiUser notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: Objek.IdPengguna,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸª Kode Bayar Gerai Ritel Keluar",
		Pesan:      fmt.Sprintf("Transaksi %s berhasil dibuat. Tunjukkan kode pembayaran di gerai ritel terdekat pilihanmu sebelum batas waktu habis ya!", Objek.KodeOrderSistem),
		CreatedAt:  time.Now().Format(time.RFC3339),
		Activity:   true,
		Inbox:      false,
		Archive:    true,
		ExpiredAt:  time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
		Pop:        5.0, // 5 detik, biar user sadar kode bayarnya udah siap
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"transaksi_id":      Objek.ID,
				"kode_order_sistem": Objek.KodeOrderSistem,
				"action_type":       "transaction_lock_gerai",
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_GERAI_PAYMENT_CODE_PAGE", // FE tampilin barcode/kode bayar
			},
		},
	}
	if err := notification_request.PostToNotification(ctx, NotificationGeraiUser, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk); err != nil {
		fmt.Printf("Gagal mengirim notifikasi gerai ke pengguna %d: %v\n", Objek.IdPengguna, err)
	}

	// Notif Seller (Silent - Masuk Inbox)
	var NotificationGeraiSeller notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: int64(Objek.IdSeller),
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸ“¦ Calon Pesanan Baru (Gerai)",
		Pesan:      fmt.Sprintf("Pembeli telah mengambil kode bayar gerai untuk pesanan %s. Menunggu pembayaran lunas.", Objek.KodeOrderSistem),
		CreatedAt:  time.Now().Format(time.RFC3339),
		Activity:   true,
		Inbox:      false,
		Archive:    true,
		ExpiredAt:  time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
		Pop:        0.0, // Silent inbox
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"transaksi_id":      Objek.ID,
				"kode_order_sistem": Objek.KodeOrderSistem,
				"action_type":       "seller_incoming_lock_gerai",
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_SELLER_ORDER_DASHBOARD",
			},
		},
	}
	if err := notification_request.PostToNotification(ctx, NotificationGeraiSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk); err != nil {
		fmt.Printf("Gagal mengirim silent notification gerai ke seller %d: %v\n", Objek.IdSeller, err)
	}

	fmt.Printf("Berhasil memproses dual-notifikasi Transaksi Gerai untuk ID %d\n", Objek.ID)

	return nil
}
