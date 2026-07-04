package wishlist_pengguna_handle

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"gorm.io/gorm"

	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/cache"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
	notification_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/models"
	notification_request "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/request"
	notification_seeders "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/seeders"
)

func CreateTambahBarangKeWishlist(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahBarangKeWishlist"
	var Objek sot_models.Wishlist
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Wishlist = cass_models.Wishlist{
		ID:            Objek.ID,
		IdPengguna:    Objek.IdPengguna,
		IdBarangInduk: Objek.IdBarangInduk,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica sync %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historica db %s dalam %s", err, handle_services)
	}

	// Setup notifikasi sukses tambah ke wishlist
	var NotificationAddWishlist notification_models.NotificationPengguna = notification_models.NotificationPengguna{
		IDPengguna: Objek.IdPengguna, // Kirim balik ke pembeli sebagai konfirmasi
		Pengirim:   notification_seeders.Sistem,
		Judul:      "â¤ï¸ Berhasil Disimpan ke Wishlist",
		Pesan:      "Produk incaranmu sukses masuk daftar favorit. Kami akan kabari kalau ada diskon atau promo menarik untuk produk ini ya!",
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339), // Keep seminggu di tab notif
		Pop:        2.0,                                              // Cukup 2 detik, biar gak ganggu user yang lagi asyik nyari barang lain
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"wishlist_id":     Objek.ID,
				"barang_induk_id": Objek.IdBarangInduk,
				"action_type":     "add_to_wishlist",
			},
			Special: map[string]interface{}{
				"click_action": "OPEN_MY_WISHLIST_PAGE", // FE arahin ke halaman wishlist si pembeli
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotificationAddWishlist, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk); err != nil {
		fmt.Printf("Gagal mengirim notifikasi tambah wishlist ke pengguna %d: %v\n", Objek.IdPengguna, err)
	}

	fmt.Printf("Berhasil memproses notifikasi tambah wishlist untuk ID %d\n", Objek.ID)
	return nil
}

func DeleteHapusBarangDariWishlist(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, Read *gorm.DB, cass_historical, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusBarangDariWishlist"
	var Objek sot_models.Wishlist
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	var ObjekCass cass_models.Wishlist = cass_models.Wishlist{
		ID:            Objek.ID,
		IdPengguna:    Objek.IdPengguna,
		IdBarangInduk: Objek.IdBarangInduk,
		CreatedAt:     Objek.CreatedAt,
		UpdatedAt:     Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica sync %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historical, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historica db %s dalam %s", err, handle_services)
	}

	return nil
}


