package barang_seller_handle

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
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
	sot_threshold "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/threshold"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	notification_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/models"
	notification_request "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/request"
	notification_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/seeders"
)

func CreateMasukanBarangInduk(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "CreateMasukanBarangInduk"
	var Objek sot_models.BarangInduk
	var wg sync.WaitGroup
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BarangInduk = cass_models.BarangInduk{
		ID:               Objek.ID,
		SellerID:         Objek.SellerID,
		IdDiskon:         Objek.IdDiskon,
		NamaBarang:       Objek.NamaBarang,
		JenisBarang:      Objek.JenisBarang,
		Deskripsi:        Objek.Deskripsi,
		OriginalKategori: Objek.OriginalKategori,
		HargaKategoris:   Objek.HargaKategoris,
		CreatedAt:        Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.BarangInduk = se_models.BarangInduk{
		ID:               Objek.ID,
		SellerID:         Objek.SellerID,
		IdDiskon:         Objek.IdDiskon,
		NamaBarang:       Objek.NamaBarang,
		JenisBarang:      Objek.JenisBarang,
		Deskripsi:        Objek.Deskripsi,
		OriginalKategori: Objek.OriginalKategori,
		HargaKategoris:   Objek.HargaKategoris,
		CreatedAt:        Objek.CreatedAt,
	}

	if task_info, err := se_index.BarangIndukIndex.AddDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke dalam search engine dengan uid task %s", task_info.IndexUID)
	}

	var NamaSeller string = ""
	if err := read.WithContext(ctx).Model(&sot_models.Seller{}).Select("nama").Where(&sot_models.Seller{
		ID: Objek.SellerID,
	}).Limit(1).Take(&NamaSeller).Error; err != nil {
		fmt.Println("gagal mengambil nama seller")
	}

	var follower_seller_total int64 = 0
	if err := read.WithContext(ctx).Model(&sot_threshold.SellerThreshold{}).Select("follower").Where(&sot_threshold.SellerThreshold{
		IdSeller: int64(Objek.SellerID),
	}).Limit(1).Take(&follower_seller_total).Error; err != nil && err == gorm.ErrRecordNotFound {
		fmt.Println("Gagal menemukan threshold follower")
	}

	var IdsPengguna []int64 = make([]int64, 0, follower_seller_total)

	if err := read.WithContext(ctx).Model(&sot_models.Follower{}).Select("id_follower").Where(&sot_models.Follower{
		IdFollowed: int64(Objek.SellerID),
	}).Limit(int(follower_seller_total)).Take(&IdsPengguna).Error; err != nil {
		fmt.Println(err)
	}

	if len(IdsPengguna) == 0 || IdsPengguna[0] == 0 {
		return errors.New("tak ada data id pengguna yang didapatkan jadi broadcast notif tak jadi")
	} else {
		// ðŸ”” Notifikasi Pengguna: Broadcast ke semua followers kalau ada produk baru
		for _, id := range IdsPengguna {
			konteks, cancel := context.WithTimeout(context.Background(), time.Second*4)
			wg.Add(1)
			go func(idUser int64, ctx_t context.Context, batal context.CancelFunc) {
				defer wg.Done()
				defer batal()
				var NotifikasiPengguna = notification_models.NotificationPengguna{
					IDPengguna: idUser, // Mengirimkan array ID followers sekaligus (Bulk/Broadcast)
					Pengirim:   notification_seeders.Sistem,
					Judul:      "âœ¨ Toko Favoritmu Punya Produk Baru!",
					Pesan:      fmt.Sprintf("Jangan sampai kehabisan! Toko %s baru saja merilis '%s'. Yuk, cek sekarang!", NamaSeller, Objek.NamaBarang),
					Pop:        0.8, // Nilai priority sedikit lebih tinggi agar menarik perhatian
					Archive:    true,
					Inbox:      true,
					Activity:   false,
					CreatedAt:  time.Now().Format(time.RFC3339),
					ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339), // Expired 7 hari karena tipe promosi
					Data: struct {
						Metadata map[string]interface{} `json:"metadata"`
						Special  interface{}            `json:"special"`
					}{
						Metadata: map[string]interface{}{
							"seller_id":       Objek.SellerID,
							"barang_induk_id": Objek.ID,
						},
						Special: map[string]interface{}{
							"click_action": "OPEN_PRODUCT_DETAIL", // Langsung arahkan ke detail produk saat diklik
						},
					},
				}
				// Kirim ke API Notifikasi bagian Pengguna/Customer
				_ = notification_request.PostToNotification[notification_models.NotificationPengguna](
					ctx,
					NotifikasiPengguna,
					cache.HostRunningAPIInNotifikasi,
					cache.PortRunningAPIInNotifikasi,
					cache.PenggunaPathNotifikasiMasuk, // Pastikan path ini mengarah ke endpoint pengguna/customer
				)
			}(id, konteks, cancel)
		}
	}

	// ðŸ”” Notifikasi Seller: Konfirmasi produk baru berhasil didaftarkan
	if Objek.SellerID != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ“¦ Produk Berhasil Ditambahkan",
			Pesan:     fmt.Sprintf("Produk '%s' Anda telah berhasil didaftarkan ke sistem.", Objek.NamaBarang),
			Pop:       0.5,
			Archive:   true,
			Inbox:     false,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.SellerID, "barang_induk_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "OPEN_PRODUCT_DETAIL"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	wg.Wait()
	return nil
}

func UpdateEditBarangInduk(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "UpdateEditBarangInduk"
	var Objek sot_models.BarangInduk
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BarangInduk = cass_models.BarangInduk{
		ID:               Objek.ID,
		SellerID:         Objek.SellerID,
		IdDiskon:         Objek.IdDiskon,
		NamaBarang:       Objek.NamaBarang,
		JenisBarang:      Objek.JenisBarang,
		Deskripsi:        Objek.Deskripsi,
		OriginalKategori: Objek.OriginalKategori,
		HargaKategoris:   Objek.HargaKategoris,
		CreatedAt:        Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var ObjekSearchEngine se_models.BarangInduk = se_models.BarangInduk{
		ID:               Objek.ID,
		SellerID:         Objek.SellerID,
		IdDiskon:         Objek.IdDiskon,
		NamaBarang:       Objek.NamaBarang,
		JenisBarang:      Objek.JenisBarang,
		Deskripsi:        Objek.Deskripsi,
		OriginalKategori: Objek.OriginalKategori,
		HargaKategoris:   Objek.HargaKategoris,
		CreatedAt:        Objek.CreatedAt,
	}

	if task_info, err := se_index.BarangIndukIndex.UpdateDocumentsWithContext(ctx, &ObjekSearchEngine, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke dalam search engine dengan uid task %s", task_info.IndexUID)
	}

	// ðŸ”” Silent Update Seller: Sinkronisasi data detail produk induk di background app
	if Objek.SellerID != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Data Produk Diperbarui",
			Pesan:     fmt.Sprintf("Informasi produk '%s' berhasil disinkronisasi.", Objek.NamaBarang),
			Pop:       0, // Silent update
			Archive:   true,
			Inbox:     false,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.SellerID, "barang_induk_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_PRODUCT"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func DeleteHapusBarangInduk(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session, se_index se_models.IndexWrapper) error {
	const handle_services string = "DeleteHapusBarangInduk"
	var Objek sot_models.BarangInduk
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.BarangInduk = cass_models.BarangInduk{
		ID:               Objek.ID,
		SellerID:         Objek.SellerID,
		IdDiskon:         Objek.IdDiskon,
		NamaBarang:       Objek.NamaBarang,
		JenisBarang:      Objek.JenisBarang,
		Deskripsi:        Objek.Deskripsi,
		OriginalKategori: Objek.OriginalKategori,
		HargaKategoris:   Objek.HargaKategoris,
		CreatedAt:        Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID)); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	idStr := strconv.FormatInt(int64(Objek.ID), 10)

	if task_info, err := se_index.BarangIndukIndex.DeleteDocumentWithContext(ctx, idStr, &meilisearch.DocumentOptions{
		PrimaryKey: meilisearch.StringPtr("id"),
	}); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam search engine %s dalam %s", err, handle_services)
	} else {
		fmt.Printf("berhasil memasukan data ke dalam search engine dengan uid task %s", task_info.IndexUID)
	}

	// ðŸ”” Silent Update Seller: Hapus produk dari list inventory UI lokal seller
	if Objek.SellerID != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ—‘ï¸ Produk Dihapus",
			Pesan:     fmt.Sprintf("Produk '%s' telah dihapus dari etalase.", Objek.NamaBarang),
			Pop:       0, // Silent update
			Archive:   true,
			Inbox:     false,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.SellerID, "barang_induk_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_REMOVE_PRODUCT"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func CreateMasukanKategoriBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateMasukanKategoriBarang"
	var Objek sot_models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KategoriBarang = cass_models.KategoriBarang{
		ID:             Objek.ID,
		SellerID:       int32(Objek.SellerID),
		IdBarangInduk:  int32(Objek.IdBarangInduk),
		IDAlamat:       Objek.IDAlamat,
		IDRekening:     Objek.IDRekening,
		Nama:           Objek.Nama,
		Deskripsi:      Objek.Deskripsi,
		Warna:          Objek.Warna,
		Stok:           int32(Objek.Stok),
		Harga:          int32(Objek.Harga),
		PotonganDiskon: int32(Objek.PotonganDiskon),
		BeratGram:      Objek.BeratGram,
		DimensiPanjang: Objek.DimensiPanjang,
		DimensiLebar:   Objek.DimensiLebar,
		Sku:            Objek.Sku,
		IsOriginal:     Objek.IsOriginal,
		CreatedAt:      Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Notifikasi Seller: Penambahan varian/kategori barang baru
	if Objek.SellerID != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "âœ¨ Varian Baru Ditambahkan",
			Pesan:     fmt.Sprintf("Varian baru '%s' dengan stok %d berhasil ditambahkan.", Objek.Nama, Objek.Stok),
			Pop:       0.5,
			Archive:   true,
			Inbox:     true,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.SellerID, "kategori_barang_id": Objek.ID, "barang_induk_id": Objek.IdBarangInduk},
				Special:  map[string]interface{}{"click_action": "OPEN_PRODUCT_VARIANT"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func UpdateEditKategoriBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditKategoriBarang"
	var Objek sot_models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KategoriBarang = cass_models.KategoriBarang{
		ID:             Objek.ID,
		SellerID:       int32(Objek.SellerID),
		IdBarangInduk:  int32(Objek.IdBarangInduk),
		IDAlamat:       Objek.IDAlamat,
		IDRekening:     Objek.IDRekening,
		Nama:           Objek.Nama,
		Deskripsi:      Objek.Deskripsi,
		Warna:          Objek.Warna,
		Stok:           int32(Objek.Stok),
		Harga:          int32(Objek.Harga),
		PotonganDiskon: int32(Objek.PotonganDiskon),
		BeratGram:      Objek.BeratGram,
		DimensiPanjang: Objek.DimensiPanjang,
		DimensiLebar:   Objek.DimensiLebar,
		Sku:            Objek.Sku,
		IsOriginal:     Objek.IsOriginal,
		CreatedAt:      Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Seller: Perubahan detail varian (stok, sku, deskripsi)
	if Objek.SellerID != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Varian Produk Diperbarui",
			Pesan:     fmt.Sprintf("Informasi varian '%s' telah disinkronisasi.", Objek.Nama),
			Pop:       0, // Silent update
			Archive:   true,
			Inbox:     false,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.SellerID, "kategori_barang_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_VARIANT"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func UpdateUbahHargaKategoriBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahHargaKategoriBarang"
	var Objek sot_models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KategoriBarang = cass_models.KategoriBarang{
		ID:             Objek.ID,
		SellerID:       int32(Objek.SellerID),
		IdBarangInduk:  int32(Objek.IdBarangInduk),
		IDAlamat:       Objek.IDAlamat,
		IDRekening:     Objek.IDRekening,
		Nama:           Objek.Nama,
		Deskripsi:      Objek.Deskripsi,
		Warna:          Objek.Warna,
		Stok:           int32(Objek.Stok),
		Harga:          int32(Objek.Harga),
		PotonganDiskon: int32(Objek.PotonganDiskon),
		BeratGram:      Objek.BeratGram,
		DimensiPanjang: Objek.DimensiPanjang,
		DimensiLebar:   Objek.DimensiLebar,
		Sku:            Objek.Sku,
		IsOriginal:     Objek.IsOriginal,
		CreatedAt:      Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Seller: Perubahan harga varian tertentu
	if Objek.SellerID != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ’° Harga Varian Berubah",
			Pesan:     fmt.Sprintf("Perubahan harga varian '%s' menjadi Rp%d berhasil diterapkan.", Objek.Nama, Objek.Harga),
			Pop:       0, // Silent update
			Archive:   true,
			Inbox:     false,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.SellerID, "kategori_barang_id": Objek.ID, "harga_baru": Objek.Harga},
				Special:  map[string]interface{}{"click_action": "SILENT_REFRESH_PRICE"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func DeleteHapusBarangKategori(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusBarangKategori"
	var Objek sot_models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KategoriBarang = cass_models.KategoriBarang{
		ID:             Objek.ID,
		SellerID:       int32(Objek.SellerID),
		IdBarangInduk:  int32(Objek.IdBarangInduk),
		IDAlamat:       Objek.IDAlamat,
		IDRekening:     Objek.IDRekening,
		Nama:           Objek.Nama,
		Deskripsi:      Objek.Deskripsi,
		Warna:          Objek.Warna,
		Stok:           int32(Objek.Stok),
		Harga:          int32(Objek.Harga),
		PotonganDiskon: int32(Objek.PotonganDiskon),
		BeratGram:      Objek.BeratGram,
		DimensiPanjang: Objek.DimensiPanjang,
		DimensiLebar:   Objek.DimensiLebar,
		Sku:            Objek.Sku,
		IsOriginal:     Objek.IsOriginal,
		CreatedAt:      Objek.CreatedAt,
	}
	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID)); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Seller: Hapus varian dari list inventory UI lokal seller
	if Objek.SellerID != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ—‘ï¸ Varian Produk Dihapus",
			Pesan:     fmt.Sprintf("Varian '%s' telah dihapus dari sistem.", Objek.Nama),
			Pop:       0, // Silent update
			Archive:   true,
			Inbox:     false,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.SellerID, "kategori_barang_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_REMOVE_VARIANT"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}
func UpdateDownStokBarangInduk(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditKategoriBarang"
	var Objek sot_models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KategoriBarang = cass_models.KategoriBarang{
		ID:             Objek.ID,
		SellerID:       int32(Objek.SellerID),
		IdBarangInduk:  int32(Objek.IdBarangInduk),
		IDAlamat:       Objek.IDAlamat,
		IDRekening:     Objek.IDRekening,
		Nama:           Objek.Nama,
		Deskripsi:      Objek.Deskripsi,
		Warna:          Objek.Warna,
		Stok:           int32(Objek.Stok),
		Harga:          int32(Objek.Harga),
		PotonganDiskon: int32(Objek.PotonganDiskon),
		BeratGram:      Objek.BeratGram,
		DimensiPanjang: Objek.DimensiPanjang,
		DimensiLebar:   Objek.DimensiLebar,
		Sku:            Objek.Sku,
		IsOriginal:     Objek.IsOriginal,
		CreatedAt:      Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Seller: Update Stok
	if Objek.SellerID != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ”„ Stok Varian Diperbarui",
			Pesan:     fmt.Sprintf("Stok untuk varian '%s' telah diperbarui.", Objek.Nama),
			Pop:       0,
			Archive:   true,
			Inbox:     false,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.SellerID, "kategori_barang_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_UPDATE_STOK"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}

func UpdateDownKategoriBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditKategoriBarang"
	var Objek sot_models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.KategoriBarang = cass_models.KategoriBarang{
		ID:             Objek.ID,
		SellerID:       int32(Objek.SellerID),
		IdBarangInduk:  int32(Objek.IdBarangInduk),
		IDAlamat:       Objek.IDAlamat,
		IDRekening:     Objek.IDRekening,
		Nama:           Objek.Nama,
		Deskripsi:      Objek.Deskripsi,
		Warna:          Objek.Warna,
		Stok:           int32(Objek.Stok),
		Harga:          int32(Objek.Harga),
		PotonganDiskon: int32(Objek.PotonganDiskon),
		BeratGram:      Objek.BeratGram,
		DimensiPanjang: Objek.DimensiPanjang,
		DimensiLebar:   Objek.DimensiLebar,
		Sku:            Objek.Sku,
		IsOriginal:     Objek.IsOriginal,
		CreatedAt:      Objek.CreatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ”” Silent Update Seller: Update Kategori/Detail Barang
	if Objek.SellerID != 0 {
		var Notifikasi = notification_models.NotificationSeller{
			IDSeller:  int64(Objek.SellerID),
			Pengirim:  notification_seeders.Sistem,
			Judul:     "ðŸ“ Data Varian Diubah",
			Pesan:     fmt.Sprintf("Informasi varian '%s' telah diperbarui.", Objek.Nama),
			Pop:       0,
			Archive:   true,
			Inbox:     false,
			Activity:  true,
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiredAt: time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"seller_id": Objek.SellerID, "kategori_barang_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "SILENT_UPDATE_VARIANT"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationSeller](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)
	}

	return nil
}
func CreateMasukanKomentarBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
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
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func UpdateEditKomentarBarang(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditKomentarBarang"
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

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
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

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID)); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}

func CreateMasukanChildKomentar(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
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
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var idPengguna int64 = 0
	if err := read.WithContext(ctx).Model(&sot_models.Komentar{}).Select("id_entity").Where(&sot_models.Komentar{
		ID: Objek.IdKomentar,
	}).Limit(1).Take(&idPengguna).Error; err != nil {
		fmt.Println("gagal mendapatkan id pengguna")
	}

	if idPengguna == 0 {
		return errors.New("Gagal mendapatkan id pengguna")
	} else {
		// ðŸ”” Push Notifikasi Balasan Komentar
		var Notifikasi = notification_models.NotificationPengguna{
			IDPengguna: idPengguna,
			Pengirim:   notification_seeders.Sistem,
			Judul:      "ðŸ’¬ Komentar Anda Dibalas",
			Pesan:      fmt.Sprintf("Seseorang membalas komentar Anda: \"%s\"", Objek.IsiKomentar),
			Pop:        1, // Munculkan pop-up / push notif aktif
			Archive:    false,
			Inbox:      true,
			Activity:   true,
			CreatedAt:  time.Now().Format(time.RFC3339),
			ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339), // Expired 7 hari
			Data: struct {
				Metadata map[string]interface{} `json:"metadata"`
				Special  interface{}            `json:"special"`
			}{
				Metadata: map[string]interface{}{"komentar_parent_id": Objek.IdKomentar, "komentar_child_id": Objek.ID},
				Special:  map[string]interface{}{"click_action": "OPEN_COMMENT_REPLY"},
			},
		}
		_ = notification_request.PostToNotification[notification_models.NotificationPengguna](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)
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

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
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

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), int64(ObjekCass.ID)); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	// Sesuai prinsip append-only, penghapusan tetap mencatat state baru ke historical db
	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	return nil
}
