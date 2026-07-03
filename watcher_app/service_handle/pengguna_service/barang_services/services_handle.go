package barang_pengguna_handle

import (
	"context"
	"errors"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"

	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	se_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/search_engine/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/cache"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	notification_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/models"
	notification_request "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/request"
	notification_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/seeders"
)

func CreateLikesBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateLikesBarang"
	var Objek sot_models.BarangDisukai
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BarangDisukai = cass_models.BarangDisukai{
		ID:            Objek.ID,
		IdPengguna:    Objek.IdPengguna,
		IdBarangInduk: Objek.IdBarangInduk,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasuakan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.IdPengguna)
	return nil
}

func DeleteUnlikesBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteUnlikesBarang"
	var Objek sot_models.BarangDisukai
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BarangDisukai = cass_models.BarangDisukai{
		ID:            Objek.ID,
		IdPengguna:    Objek.IdPengguna,
		IdBarangInduk: Objek.IdBarangInduk,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedDataLikesBarang map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal menghapus data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedDataLikesBarang)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedDataLikesBarang); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.IdPengguna)
	return nil
}

func CreateMasukanKomentarBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateMasukanKomentarBarang"
	var Objek sot_models.Komentar
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Komentar = cass_models.Komentar{
		ID:            Objek.ID,
		IdBarangInduk: Objek.IdBarangInduk,
		IdEntity:      Objek.IdEntity,
		JenisEntity:   Objek.JenisEntity,
		Komentar:      Objek.Komentar,
		IsSeller:      Objek.IsSeller,
		Dibalas:       Objek.Dibalas,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	var idSeller int32 = 0
	if err := read.WithContext(ctx).Model(&sot_models.BarangInduk{}).Select("id_seller").Where(&sot_models.BarangInduk{
		ID: Objek.IdBarangInduk,
	}).Limit(1).Take(&idSeller).Error; err != nil {
		fmt.Println("Gagal mendapatkan data id seller")
	}

	if idSeller == 0 {
		return errors.New("gagal mendapatkan id seller")
	}

	// 1. Inisialisasi variabel untuk kebutuhan dinamis
	var targetUserID int64
	var targetPath string
	var judulNotif, pesanNotif string

	// Potong komentar biar gak kepanjangan kalau tampil di banner notifikasi
	cuplikanKomentar := Objek.Komentar
	if len(cuplikanKomentar) > 60 {
		cuplikanKomentar = cuplikanKomentar[:57] + "..."
	}

	// 2. Tentukan arah notifikasi berdasarkan siapa yang berkomentar
	if Objek.IsSeller {
		// Kasus: Seller menjawab pertanyaan pembeli
		targetUserID = Objek.IdEntity                        // Notif dikirim ke pembeli asli
		targetPath = cache.PenggunaPathNotifikasiMasuk // Path khusus pengguna/pembeli
		judulNotif = "ðŸ’¬ Pertanyaanmu Dijawab Seller!"
		pesanNotif = fmt.Sprintf("Toko baru saja membalas diskusimu: \"%s\". Cek jawabannya sekarang!", cuplikanKomentar)
	} else {
		// Kasus: Pembeli bertanya di produk seller
		targetUserID = int64(idSeller)                     // Notif dikirim ke seller
		targetPath = cache.SellerPathNotifikasiMasuk // Path khusus seller sesuai request lu
		judulNotif = "ðŸ›’ Ada Pertanyaan Baru Produk!"
		pesanNotif = fmt.Sprintf("Calon pembeli nanyain produk lu nih: \"%s\". Yuk bales biar makin laris!", cuplikanKomentar)
	}

	// 3. Susun struct notifikasi dengan metadata komplit buat FE deep-linking
	var NotificationKomentar notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: targetUserID,
		Pengirim:   notification_seeders.Sistem,
		Judul:      judulNotif,
		Pesan:      pesanNotif,
		Pop:        1,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 14).Format(time.RFC3339), // Notif diskusi produk keep 2 minggu aja cukup
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"barang_induk_id": Objek.IdBarangInduk,
				"komentar_id":     Objek.ID,
				"action_type":     "new_diskusi_produk",
				"is_seller_reply": Objek.IsSeller,
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_DISKUSI_PRODUK_PAGE", // FE langsung buka halaman diskusi produk terkait
			},
		},
	}

	// 4. Tembak ke API Notifikasi dengan path yang sudah dinamis
	if err := notification_request.PostToNotification(ctx, NotificationKomentar, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, targetPath); err != nil {
		return fmt.Errorf("gagal mengirim notifikasi komentar: %w", err)
	}

	return nil
}

func UpdateEditKomentarBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatedEditKomentarBarang"
	var Objek cass_models.Komentar
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Komentar = cass_models.Komentar{
		ID:            Objek.ID,
		IdBarangInduk: Objek.IdBarangInduk,
		IdEntity:      Objek.IdEntity,
		JenisEntity:   Objek.JenisEntity,
		Komentar:      Objek.Komentar,
		IsSeller:      Objek.IsSeller,
		Dibalas:       Objek.Dibalas,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
		DeletedAt:     Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}
	return nil
}

func DeleteHapusKomentarBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusKomentarBarang"
	var Objek sot_models.Komentar
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Komentar = cass_models.Komentar{
		ID:            Objek.ID,
		IdBarangInduk: Objek.IdBarangInduk,
		IdEntity:      Objek.IdEntity,
		JenisEntity:   Objek.JenisEntity,
		Komentar:      Objek.Komentar,
		IsSeller:      Objek.IsSeller,
		Dibalas:       Objek.Dibalas,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}
	return nil
}

func CreateMasukanChildKomentar(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateMasukanChildKomentar"
	var Objek sot_models.KomentarChild
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KomentarChild = cass_models.KomentarChild{
		ID:          Objek.ID,
		IdKomentar:  Objek.IdKomentar,
		IdEntity:    Objek.IdEntity,
		JenisEntity: Objek.JenisEntity,
		IsiKomentar: Objek.IsiKomentar,
		IsSeller:    Objek.IsSeller,
		Mention:     Objek.Mention,
		CreatedAt:   Objek.CreatedAt,
		UpdatedAt:   Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	// 1. Inisialisasi variabel target notifikasi
	var targetUserID int64
	var targetPath string
	var judulNotif, pesanNotif string

	// Potong isi komentar anak biar pas di push notification banner
	cuplikanChild := Objek.IsiKomentar
	if len(cuplikanChild) > 60 {
		cuplikanChild = cuplikanChild[:57] + "..."
	}

	// 2. Logika penentuan target berdasarkan fitur Mention atau IsSeller
	if Objek.Mention != "" {
		// Skenario A: Ada user/seller yang di-mention/tag khusus
		// NOTE: Di sini asumsinya Objek.IdEntity target mention bisa didapat atau diarahkan ke entitas pembeli/penjual.
		// Sementara kita pakai logika deteksi: jika yang mention adalah seller, target kemungkinan besar pengguna, dst.
		if Objek.IsSeller {
			targetUserID = Objek.IdEntity // Mengarah ke pembeli terkait thread
			targetPath = cache.PenggunaPathNotifikasiMasuk
		} else {
			targetUserID = Objek.IdEntity // Mengarah ke user lain / seller
			targetPath = cache.PenggunaPathNotifikasiMasuk
		}
		judulNotif = "ðŸ”” Kamu Disebut dalam Diskusi!"
		pesanNotif = fmt.Sprintf("@%s menyebutmu di komentar: \"%s\". Yuk join obrolannya!", Objek.Mention, cuplikanChild)

	} else if Objek.IsSeller {
		// Skenario B: Seller yang balas komentar biasa tanpa mention khusus
		targetUserID = Objek.IdEntity // Mengirim ke pembeli utama pemilik thread
		targetPath = cache.PenggunaPathNotifikasiMasuk
		judulNotif = "ðŸ’¬ Pertanyaanmu Dibalas Seller!"
		pesanNotif = fmt.Sprintf("Seller baru saja menanggapi diskusi: \"%s\"", cuplikanChild)

	} else {
		// Skenario C: Pembeli/User lain yang membalas thread di toko seller
		// Karena di model ini tidak bawa idSeller langsung dari database (seperti induknya),
		// idealnya lu ambil dulu idSeller via DB, tapi kalau mau cepat/sementara pakai target idEntity pembeli/seller lain:
		targetUserID = Objek.IdEntity
		targetPath = cache.SellerPathNotifikasiMasuk // Masuk ke dashboard/notif seller
		judulNotif = "ðŸ’¬ Ada Balasan Baru di Diskusi Produk!"
		pesanNotif = fmt.Sprintf("Ada tanggapan baru di lapak lu nih: \"%s\"", cuplikanChild)
	}

	// 3. Bangun struct Notifikasi
	var NotificationChild notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: targetUserID,
		Pengirim:   notification_seeders.Sistem,
		Judul:      judulNotif,
		Pesan:      pesanNotif,
		Pop:        1,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339), // Notif child comment simpan 7 hari aja biar gak numpuk
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"komentar_induk_id": Objek.IdKomentar,
				"child_id":          Objek.ID,
				"action_type":       "new_child_comment",
				"has_mention":       Objek.Mention != "",
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_THREAD_DISKUSI_PAGE", // FE langsung fokus ke sub-comment/thread terkait
			},
		},
	}

	// 4. Kirim Notifikasi secara dinamis
	if err := notification_request.PostToNotification(ctx, NotificationChild, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, targetPath); err != nil {
		return fmt.Errorf("gagal mengirim notifikasi child komentar: %w", err)
	}

	return nil
}

func CreateMentionChildKomentar(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateMentionChildKomentar"
	var Objek sot_models.KomentarChild
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KomentarChild = cass_models.KomentarChild{
		ID:          Objek.ID,
		IdKomentar:  Objek.IdKomentar,
		IdEntity:    Objek.IdEntity,
		JenisEntity: Objek.JenisEntity,
		IsiKomentar: Objek.IsiKomentar,
		IsSeller:    Objek.IsSeller,
		Mention:     Objek.Mention,
		CreatedAt:   Objek.CreatedAt,
		UpdatedAt:   Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	// 1. Potong isi komentar biar rapi di banner notifikasi
	cuplikanMention := Objek.IsiKomentar
	if len(cuplikanMention) > 60 {
		cuplikanMention = cuplikanMention[:57] + "..."
	}

	// 2. Tentukan jalur notifikasi (Default ke pengguna, kalau seller ngetag ya masuk ke path pengguna juga)
	var targetPath string = cache.PenggunaPathNotifikasiMasuk
	if Objek.IsSeller {
		// Kalau seller yang nge-mention, berarti yang dituju pasti pembeli/pengguna biasa
		targetPath = cache.PenggunaPathNotifikasiMasuk
	} else {
		// Kalau sesama pembeli atau pembeli nyenggol seller, bisa disesuaikan.
		// Di sini kita default-kan ke pengguna dulu, aman.
		targetPath = cache.PenggunaPathNotifikasiMasuk
	}

	// 3. Setup kalimat yang bikin user penasaran & langsung nge-klik
	judulNotif := "ðŸ”” Seseorang Menyebutmu!"
	pesanNotif := fmt.Sprintf("@%s nge-tag lu di diskusi produk: \"%s\". Cek obrolannya gih!", Objek.Mention, cuplikanMention)

	var NotificationMention notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: Objek.IdEntity, // ID target yang di-mention
		Pengirim:   notification_seeders.Sistem,
		Judul:      judulNotif,
		Pesan:      pesanNotif,
		Pop:        1,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339), // Simpan seminggu aja
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"komentar_induk_id": Objek.IdKomentar,
				"child_id":          Objek.ID,
				"action_type":       "mention_child_comment",
				"mentioned_user":    Objek.Mention,
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_MENTION_THREAD_PAGE", // FE langsung scroll ke komentar yang ada tag-nya
			},
		},
	}

	// 4. Kirim langsung ke service notifikasi, kelar!
	if err := notification_request.PostToNotification(ctx, NotificationMention, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, targetPath); err != nil {
		return fmt.Errorf("gagal mengirim notifikasi mention child: %w", err)
	}
	return nil
}

func UpdateEditChildKomentar(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditChildKomentar"
	var Objek sot_models.KomentarChild
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KomentarChild = cass_models.KomentarChild{
		ID:          Objek.ID,
		IdKomentar:  Objek.IdKomentar,
		IdEntity:    Objek.IdEntity,
		JenisEntity: Objek.JenisEntity,
		IsiKomentar: Objek.IsiKomentar,
		IsSeller:    Objek.IsSeller,
		Mention:     Objek.Mention,
		CreatedAt:   Objek.CreatedAt,
		UpdatedAt:   Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}
	return nil
}

func DeleteHapusChildKomentar(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusChildKomentar"
	var Objek sot_models.KomentarChild
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KomentarChild = cass_models.KomentarChild{
		ID:          Objek.ID,
		IdKomentar:  Objek.IdKomentar,
		IdEntity:    Objek.IdEntity,
		JenisEntity: Objek.JenisEntity,
		IsiKomentar: Objek.IsiKomentar,
		IsSeller:    Objek.IsSeller,
		Mention:     Objek.Mention,
		CreatedAt:   Objek.CreatedAt,
		UpdatedAt:   Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}
	return nil
}

func CreateTambahKeranjangBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahKeranjangBarang"
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

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	judulTambahKeranjang := "ðŸ›’ Berhasil Masuk Keranjang!"
	pesanTambahKeranjang := fmt.Sprintf("Mantap! Produk pilihan lu (ID Barang: %d) sebanyak %d pcs udah aman di keranjang. Yuk langsung checkout sebelum kehabisan!", Objek.IdBarangInduk, Objek.Jumlah)

	var NotificationCart notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: Objek.IdPengguna, // Dikirim balik ke pengguna sesuai request lu
		Pengirim:   notification_seeders.Sistem,
		Judul:      judulTambahKeranjang,
		Pesan:      pesanTambahKeranjang,
		Pop:        1.2,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339), // Simpan seminggu aja
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"keranjang_id":    Objek.ID,
				"barang_induk_id": Objek.IdBarangInduk,
				"action_type":     "add_to_cart_user",
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_MY_CART_PAGE", // FE langsung arahin user ke halaman keranjang belanja mereka
			},
		},
	}

	// Kirim menggunakan PenggunaPathNotifikasiMasuk
	if err := notification_request.PostToNotification(ctx, NotificationCart, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk); err != nil {
		return fmt.Errorf("gagal mengirim notifikasi tambah keranjang ke pengguna: %w", err)
	}
	return nil
}

func UpdateEditKeranjangBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditKeranjangBarang"
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

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}
	return nil
}

func DeleteHapusKeranjangBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusKeranjangBarang"
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

func CreateBerikanReviewBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateBerikanReviewBarang"
	var Objek sot_models.Review
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	var ObjekCass cass_models.Review = cass_models.Review{
		ID:            Objek.ID,
		IdPengguna:    Objek.IdPengguna,
		IdBarangInduk: Objek.IdBarangInduk,
		Rating:        Objek.Rating,
		Ulasan:        Objek.Ulasan,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	var idSeller int32 = 0
	if err := read.WithContext(ctx).Model(&sot_models.BarangInduk{}).Select("id_seller").Where(&sot_models.BarangInduk{
		ID: Objek.IdBarangInduk,
	}).Limit(1).Take(&idSeller).Error; err != nil {
		fmt.Println("Gagal mendapatkan data id seller")
	}

	if idSeller == 0 {
		return errors.New("gagal mendapatkan id seller")
	}

	// Potong ulasan biar pas di banner notifikasi
	cuplikanUlasan := Objek.Ulasan
	if len(cuplikanUlasan) > 50 {
		cuplikanUlasan = cuplikanUlasan[:47] + "..."
	}

	// Bikin representasi bintang (emoji) berdasarkan rating biar visualnya keren
	var emojiRating string
	for i := 0; i < int(Objek.Rating); i++ {
		emojiRating += "â­"
	}
	if emojiRating == "" {
		emojiRating = "â­"
	}

	// ==========================================
	// NOTIFIKASI 1: UNTUK PENGGUNA (PEMBELI)
	// ==========================================
	pesanPengguna := fmt.Sprintf("Makasih banyak udah kasih ulasan %s buat produk ini! Kontribusi lu berharga banget buat pembeli lain. ðŸ™", emojiRating)

	var NotificationToUser notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: Objek.IdPengguna,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸŽ‰ Ulasan Lu Berhasil Dikirim!",
		Pesan:      pesanPengguna,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 1, 0).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"review_id":       Objek.ID,
				"barang_induk_id": Objek.IdBarangInduk,
				"action_type":     "user_give_review",
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_PRODUCT_REVIEW_PAGE", // FE arahin ke list review produk itu
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotificationToUser, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk); err != nil {
		fmt.Printf("Gagal mengirim notifikasi review ke pengguna: %v\n", err)
		// Tetap lanjut, jangan di-return err biar transaksi utama gak rollback cuma gara-gara notif gagal
	}

	// ==========================================
	// NOTIFIKASI 2: UNTUK SELLER
	// ==========================================
	pesanSeller := fmt.Sprintf("Asyik! Toko lu dapet ulasan %s baru: \"%s\". Yuk cek kata pembeli dan bales ulasannya!", emojiRating, cuplikanUlasan)

	var NotificationToSeller notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: int64(idSeller),
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸ“ˆ Ada Ulasan Produk Baru!",
		Pesan:      pesanSeller,
		Pop:        4.3,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 1, 0).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"review_id":       Objek.ID,
				"barang_induk_id": Objek.IdBarangInduk,
				"action_type":     "seller_received_review",
				"rating":          Objek.Rating,
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_SELLER_REVIEW_DASHBOARD", // FE bawa seller ke dashboard ulasan toko
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotificationToSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk); err != nil {
		fmt.Printf("Gagal mengirim notifikasi review ke seller: %v\n", err)
	}
	return nil
}

func UpdateBerikanReviewBarangIIUpdateTransaksi(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdateberikanReviewBarangiiUpdateTransaksi"
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

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
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

	if task_info, err := se_index.TransaksiIndex.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("Berhasil memasukan data ke dalam search engine dengan antrean UID %s", task_info.IndexUID)
	}

	return nil
}

func CreateLikeReviewBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdateLikeReviewBarang"
	var Objek sot_models.ReviewLike
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.ReviewLike = cass_models.ReviewLike{
		ID:         Objek.ID,
		IdPengguna: Objek.IdPengguna,
		IdReview:   Objek.IdReview,
		CreatedAt:  Objek.CreatedAt,
		UpdatedAt:  Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	var idPenggunaPenulisUlasan int64 = 0
	if err := read.WithContext(ctx).Model(&sot_models.Review{}).Select("id_pengguna").Where(&sot_models.Review{
		ID: Objek.IdReview,
	}).Limit(1).Take(&idPenggunaPenulisUlasan).Error; err != nil {
		return err
	}

	if idPenggunaPenulisUlasan == 0 {
		return errors.New("gagal mendapatkan id pengguna penulis ulasan untuk notifikasi")
	}

	// 1. Setup kalimat notifikasi penambah dopamin user wkwk
	judulLike := "â¤ï¸ Ulasanmu Bermanfaat!"
	pesanLike := "Seseorang baru saja menyukai ulasan produk yang kamu tulis. Terima kasih ya sudah membantu pembeli lain!"

	var NotificationLikeReview notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: idPenggunaPenulisUlasan, // Dikirim ke orang yang nulis ulasan asli!
		Pengirim:   notification_seeders.Sistem,
		Judul:      judulLike,
		Pesan:      pesanLike,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339), // Simpan seminggu aja di tab notif
		Pop:        2.3,                                              // 3.5 detik, pas buat baca pesan apresiasi ringan
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"like_id":     Objek.ID,
				"review_id":   Objek.IdReview,
				"action_type": "review_get_liked",
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_MY_REVIEW_DETAIL", // FE bisa arahin user ke ulasan milik mereka sendiri
			},
		},
	}

	// 2. Kirim ke path pengguna umum
	if err := notification_request.PostToNotification(ctx, NotificationLikeReview, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk); err != nil {
		fmt.Printf("Gagal mengirim notifikasi apresiasi like review: %v\n", err)
	}

	return nil
}

func DeleteUnlikeReviewBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "DeleteUnlikeReviewBarang"
	var Objek sot_models.ReviewLike
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.ReviewLike = cass_models.ReviewLike{
		ID:         Objek.ID,
		IdPengguna: Objek.IdPengguna,
		IdReview:   Objek.IdReview,
		CreatedAt:  Objek.CreatedAt,
		UpdatedAt:  Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}
	return nil
}

