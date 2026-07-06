package media_kurir_handle

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"gorm.io/gorm"

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

func UpdateUbahKurirProfilFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahKurirProfilFoto"
	var Objek sot_models.MediaKurirProfilFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaKurirProfilFoto = cass_models.MediaKurirProfilFoto{
		ID:        Objek.ID,
		IdKurir:   Objek.IdKurir,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
		UpdatedAt: Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
			return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
		}
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusKurirProfilFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusKurirProfilFoto"
	var Objek sot_models.MediaKurirProfilFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaKurirProfilFoto = cass_models.MediaKurirProfilFoto{
		ID:        Objek.ID,
		IdKurir:   Objek.IdKurir,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
		UpdatedAt: Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari  sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahMediaInformasiKendaraanKurirKendaraanFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaInformasiKendaraanKurirKendaraanFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirKendaraanFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirKendaraanFoto = cass_models.MediaInformasiKendaraanKurirKendaraanFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var IdKurir int64 = 0
	if err := read.WithContext(ctx).Model(&sot_models.InformasiKendaraanKurir{}).Select("id_kurir").Where(&sot_models.InformasiKendaraanKurir{
		ID: Objek.IdInformasiKendaraanKurir,
	}).Limit(1).Take(&IdKurir).Error; err != nil {
		return err
	}

	// Tarik nama kurir berdasarkan IdKurir yang barusan didapet
	var NamaKurir string = ""
	if IdKurir != 0 {
		if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", IdKurir).Limit(1).Take(&NamaKurir).Error; err != nil {
			fmt.Println("Gagal mengambil nama kurir:", err)
		}
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// ðŸ”” NOTIFIKASI CREATE FOTO (Muncul Pop-Up)
	var Notifikasi notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ“¸ Foto Kendaraan Berhasil Diunggah",
		Pesan:     fmt.Sprintf("Halo %s, foto kendaraan lu sukses di-upload. Berkas fisik ini bakal segera ditinjau oleh tim verifikator kami.", NamaKurir),
		Pop:       2.5, // Pop-up muncul 2.5 detik
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339), // Simpan 7 hari
		Activity:  true,
		Inbox:     false,
		Archive:   true,
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"media_foto_id":          Objek.ID,
				"informasi_kendaraan_id": Objek.IdInformasiKendaraanKurir,
				"action_type":            "create_foto_kendaraan",
				"platform":               "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REDIRECT_TO_VEHICLE_GALLERY",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		fmt.Println("Gagal mengirim notifikasi create foto kendaraan:", err)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateTambahMediaInformasiKendaraanKurirKendaraanFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateTambahMediaInformasiKendaraanKurirKendaraanFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirKendaraanFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirKendaraanFoto = cass_models.MediaInformasiKendaraanKurirKendaraanFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var IdKurir int64 = 0
	if err := read.WithContext(ctx).Model(&sot_models.InformasiKendaraanKurir{}).Select("id_kurir").Where(&sot_models.InformasiKendaraanKurir{
		ID: Objek.IdInformasiKendaraanKurir,
	}).Limit(1).Take(&IdKurir).Error; err != nil {
		return err
	}

	var NamaKurir string = ""
	if IdKurir != 0 {
		if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", IdKurir).Limit(1).Take(&NamaKurir).Error; err != nil {
			fmt.Println("Gagal mengambil nama kurir:", err)
		}
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// ðŸ”” NOTIFIKASI UPDATE FOTO (Silent Update, Langsung Masuk Inbox)
	var NotifikasiUpdate notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ”„ Foto Berkas Kendaraan Diperbarui",
		Pesan:     fmt.Sprintf("Halo %s, perubahan foto kendaraan lu berhasil disimpan ke database internal.", NamaKurir),
		Pop:       0, // Sesuai request: 0 biar ga usah muncul pop-up, langsung masuk inbox
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 5).Format(time.RFC3339), // Disimpan 5 hari di inbox
		Activity:  true,
		Inbox:     false,
		Archive:   true,
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"media_foto_id":          Objek.ID,
				"informasi_kendaraan_id": Objek.IdInformasiKendaraanKurir,
				"action_type":            "update_foto_kendaraan",
				"platform":               "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REFRESH_VEHICLE_GALLERY",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotifikasiUpdate, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		fmt.Println("Gagal mengirim notifikasi update foto kendaraan:", err)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusMediaInformasiKendaraanKurirKendaraanFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusMediaInformasiKendaraanKurirKendaraanFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirKendaraanFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirKendaraanFoto = cass_models.MediaInformasiKendaraanKurirKendaraanFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahInformasiKendaraanKurirBPKBFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahInformasiKendaraankurirBPKBFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirBPKBFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirBPKBFoto = cass_models.MediaInformasiKendaraanKurirBPKBFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Ambil IdKurir dari tabel InformasiKendaraanKurir
	var IdKurir int64 = 0
	if err := read.WithContext(ctx).Model(&sot_models.InformasiKendaraanKurir{}).Select("id_kurir").Where(&sot_models.InformasiKendaraanKurir{
		ID: Objek.IdInformasiKendaraanKurir,
	}).Limit(1).Take(&IdKurir).Error; err != nil {
		return err
	}

	// Tarik nama kurir biar personal
	var NamaKurir string = ""
	if IdKurir != 0 {
		if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", IdKurir).Limit(1).Take(&NamaKurir).Error; err != nil {
			fmt.Println("Gagal mengambil nama kurir:", err)
		}
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// ðŸ”” NOTIFIKASI CREATE BPKB FOTO (Muncul Pop-Up)
	var Notifikasi notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ“‘ Dokumen BPKB Sukses Diunggah",
		Pesan:     fmt.Sprintf("Halo %s, berkas foto BPKB kendaraan lu udah masuk ke sistem internal. Dokumen aman dan siap direview oleh tim verifikasi data.", NamaKurir),
		Pop:       3.0, // Alert dokumen penting kasih 3 detik
		Activity:  true,
		Inbox:     false,
		Archive:   true,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339), // Eksis 7 hari di inbox
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"media_bpkb_id":          Objek.ID,
				"informasi_kendaraan_id": Objek.IdInformasiKendaraanKurir,
				"action_type":            "create_foto_bpkb",
				"platform":               "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REDIRECT_TO_VEHICLE_DOCUMENTS",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		fmt.Println("Gagal mengirim notifikasi create foto BPKB:", err)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateTambahInformasiKendaraanKurirBPKBFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateTambahInformasiKendaraanKurirBPKBFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirBPKBFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirBPKBFoto = cass_models.MediaInformasiKendaraanKurirBPKBFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Ambil IdKurir dari tabel InformasiKendaraanKurir
	var IdKurir int64 = 0
	if err := read.WithContext(ctx).Model(&sot_models.InformasiKendaraanKurir{}).Select("id_kurir").Where(&sot_models.InformasiKendaraanKurir{
		ID: Objek.IdInformasiKendaraanKurir,
	}).Limit(1).Take(&IdKurir).Error; err != nil {
		return err
	}

	// Tarik nama kurir
	var NamaKurir string = ""
	if IdKurir != 0 {
		if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", IdKurir).Limit(1).Take(&NamaKurir).Error; err != nil {
			fmt.Println("Gagal mengambil nama kurir:", err)
		}
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// ðŸ”” NOTIFIKASI UPDATE BPKB FOTO (Silent Update, Langsung Masuk Inbox)
	var NotifikasiUpdate notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ”„ Berkas BPKB Diperbarui",
		Pesan:     fmt.Sprintf("Halo %s, lu baru saja memperbarui file foto BPKB kendaraan. Sistem akan memperbarui berkas antrean verifikasi lu.", NamaKurir),
		Pop:       0, // Langsung masuk inbox tanpa memunculkan pop-up di layar
		Activity:  true,
		Inbox:     false,
		Archive:   true,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 5).Format(time.RFC3339), // Disimpan 5 hari di inbox
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"media_bpkb_id":          Objek.ID,
				"informasi_kendaraan_id": Objek.IdInformasiKendaraanKurir,
				"action_type":            "update_foto_bpkb",
				"platform":               "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REFRESH_VEHICLE_DOCUMENTS",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotifikasiUpdate, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		fmt.Println("Gagal mengirim notifikasi update foto BPKB:", err)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
func DeleteHapusInformasiKendaraanKurirBPKBFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusInformasiKendaraanKurirBPKBFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirBPKBFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirBPKBFoto = cass_models.MediaInformasiKendaraanKurirBPKBFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID); err != nil {
		return fmt.Errorf("gagal menghapus data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahInformasiKendaraanKurirSTNKFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahInformasiKendaraanKurirSTNKFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirSTNKFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirSTNKFoto = cass_models.MediaInformasiKendaraanKurirSTNKFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Tarik IdKurir dari tabel InformasiKendaraanKurir
	var IdKurir int64 = 0
	if err := read.WithContext(ctx).Model(&sot_models.InformasiKendaraanKurir{}).Select("id_kurir").Where(&sot_models.InformasiKendaraanKurir{
		ID: Objek.IdInformasiKendaraanKurir,
	}).Limit(1).Take(&IdKurir).Error; err != nil {
		return err
	}

	// Tarik nama kurir biar personal
	var NamaKurir string = ""
	if IdKurir != 0 {
		if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", IdKurir).Limit(1).Take(&NamaKurir).Error; err != nil {
			fmt.Println("Gagal mengambil nama kurir:", err)
		}
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// ðŸ”” NOTIFIKASI CREATE STNK FOTO (Muncul Pop-Up)
	var Notifikasi notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ“„ Berkas STNK Berhasil Diunggah",
		Pesan:     fmt.Sprintf("Halo %s, berkas foto STNK kendaraan lu udah aman masuk ke sistem. Tim verifikator bakal segera mengecek validitasnya ya!", NamaKurir),
		Pop:       3.0, // Pop-up muncul 3 detik
		Activity:  true,
		Inbox:     false,
		Archive:   true,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339), // Eksis 7 hari di inbox
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"media_stnk_id":          Objek.ID,
				"informasi_kendaraan_id": Objek.IdInformasiKendaraanKurir,
				"action_type":            "create_foto_stnk",
				"platform":               "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REDIRECT_TO_VEHICLE_DOCUMENTS",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		fmt.Println("Gagal mengirim notifikasi create foto STNK:", err)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateTambahInformasiKendaraanKurirSTNKFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateTambahInformasiKendaraanKurirSTNKFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirSTNKFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirSTNKFoto = cass_models.MediaInformasiKendaraanKurirSTNKFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Tarik IdKurir dari tabel InformasiKendaraanKurir
	var IdKurir int64 = 0
	if err := read.WithContext(ctx).Model(&sot_models.InformasiKendaraanKurir{}).Select("id_kurir").Where(&sot_models.InformasiKendaraanKurir{
		ID: Objek.IdInformasiKendaraanKurir,
	}).Limit(1).Take(&IdKurir).Error; err != nil {
		return err
	}

	// Tarik nama kurir
	var NamaKurir string = ""
	if IdKurir != 0 {
		if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", IdKurir).Limit(1).Take(&NamaKurir).Error; err != nil {
			fmt.Println("Gagal mengambil nama kurir:", err)
		}
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// ðŸ”” NOTIFIKASI UPDATE STNK FOTO (Silent Update, Langsung Masuk Inbox)
	var NotifikasiUpdate notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ”„ Foto STNK Diperbarui",
		Pesan:     fmt.Sprintf("Halo %s, file foto STNK kendaraan lu berhasil diperbarui ke sistem internal.", NamaKurir),
		Pop:       0, // Sesuai request: 0 biar silent, gak ganggu layar, langsung ngendap di inbox
		Activity:  true,
		Inbox:     false,
		Archive:   true,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 5).Format(time.RFC3339), // Simpan 5 hari di inbox
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"media_stnk_id":          Objek.ID,
				"informasi_kendaraan_id": Objek.IdInformasiKendaraanKurir,
				"action_type":            "update_foto_stnk",
				"platform":               "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REFRESH_VEHICLE_DOCUMENTS",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotifikasiUpdate, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		fmt.Println("Gagal mengirim notifikasi update foto STNK:", err)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusInformasiKendaraanKurirSTNKFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusInformasiKendaraanKurirSTNKFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirSTNKFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirSTNKFoto = cass_models.MediaInformasiKendaraanKurirSTNKFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahMediaInformasiKurirKTPFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaInformasiKurirKTPFoto"
	var Objek sot_models.MediaInformasiKurirKTPFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKurirKTPFoto = cass_models.MediaInformasiKurirKTPFoto{
		ID:               Objek.ID,
		IdInformasiKurir: Objek.IdInformasiKurir,
		Key:              Objek.Key,
		Format:           Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Tarik IdKurir dari tabel InformasiKurir
	var IdKurir int64 = 0
	if err := read.WithContext(ctx).Model(&sot_models.InformasiKurir{}).Select("id_kurir").Where(&sot_models.InformasiKurir{
		ID: Objek.IdInformasiKurir,
	}).Limit(1).Take(&IdKurir).Error; err != nil {
		return err
	}

	// Tarik nama kurir biar personal
	var NamaKurir string = ""
	if IdKurir != 0 {
		if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", IdKurir).Limit(1).Take(&NamaKurir).Error; err != nil {
			fmt.Println("Gagal mengambil nama kurir:", err)
		}
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// ðŸ”” NOTIFIKASI CREATE FOTO KTP (Muncul Pop-Up)
	var Notifikasi notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸªª Foto KTP Sukses Terunggah",
		Pesan:     fmt.Sprintf("Halo %s, foto KTP lu udah aman tersimpan di sistem. Berkas ini akan langsung diproses untuk kebutuhan verifikasi akun lu.", NamaKurir),
		Pop:       3.0, // Pop-up muncul selama 3 detik
		Activity:  true,
		Inbox:     false,
		Archive:   true,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339), // Simpan 7 hari
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"media_ktp_id":       Objek.ID,
				"informasi_kurir_id": Objek.IdInformasiKurir,
				"action_type":        "create_foto_ktp",
				"platform":           "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REDIRECT_TO_VERIFICATION_STATUS",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		fmt.Println("Gagal mengirim notifikasi create foto KTP:", err)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateTambahMediaInformasiKurirKTPFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateTambahMediaInformasiKurirKTPFoto"
	var Objek sot_models.MediaInformasiKurirKTPFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKurirKTPFoto = cass_models.MediaInformasiKurirKTPFoto{
		ID:               Objek.ID,
		IdInformasiKurir: Objek.IdInformasiKurir,
		Key:              Objek.Key,
		Format:           Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// ðŸ•µï¸â€â™‚ï¸ Tarik IdKurir dari tabel InformasiKurir
	var IdKurir int64 = 0
	if err := read.WithContext(ctx).Model(&sot_models.InformasiKurir{}).Select("id_kurir").Where(&sot_models.InformasiKurir{
		ID: Objek.IdInformasiKurir,
	}).Limit(1).Take(&IdKurir).Error; err != nil {
		return err
	}

	// Tarik nama kurir
	var NamaKurir string = ""
	if IdKurir != 0 {
		if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where("id = ?", IdKurir).Limit(1).Take(&NamaKurir).Error; err != nil {
			fmt.Println("Gagal mengambil nama kurir:", err)
		}
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// ðŸ”” NOTIFIKASI UPDATE FOTO KTP (Silent Update, Langsung Masuk Inbox)
	var NotifikasiUpdate notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   IdKurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ”„ Foto KTP Diperbarui",
		Pesan:     fmt.Sprintf("Halo %s, perubahan berkas foto KTP lu berhasil disimpan ke dalam sistem.", NamaKurir),
		Pop:       0, // Silent update, langsung masuk inbox tanpa ganggu screen kurir
		Activity:  true,
		Inbox:     false,
		Archive:   true,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 5).Format(time.RFC3339), // Simpan 5 hari di inbox
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"media_ktp_id":       Objek.ID,
				"informasi_kurir_id": Objek.IdInformasiKurir,
				"action_type":        "update_foto_ktp",
				"platform":           "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REFRESH_ACCOUNT_DOCUMENTS",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotifikasiUpdate, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		fmt.Println("Gagal mengirim notifikasi update foto KTP:", err)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
func DeleteHapusMediaInformasiKurirKTPFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusMediaInformasiKurirKTPFoto"
	var Objek sot_models.MediaInformasiKurirKTPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKurirKTPFoto = cass_models.MediaInformasiKurirKTPFoto{
		ID:               Objek.ID,
		IdInformasiKurir: Objek.IdInformasiKurir,
		Key:              Objek.Key,
		Format:           Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahMediaPengirimanPickedUpFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaPengirimanPickedUpFoto"
	var Objek sot_models.MediaPengirimanPickedUpFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaPengirimanPickedUpFoto = cass_models.MediaPengirimanPickedUpFoto{
		ID:           Objek.ID,
		IdPengiriman: Objek.IdPengiriman,
		Key:          Objek.Key,
		Format:       Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var (
		IdSeller   int64 = 0
		IdPengguna int64 = 0
	)

	var pengiriman sot_models.Pengiriman
	if err := read.WithContext(ctx).Model(&sot_models.Pengiriman{}).Select("id_seller, id_alamat_pengguna").Where(&sot_models.Pengiriman{
		ID: Objek.IdPengiriman,
	}).Limit(1).Take(&pengiriman).Error; err != nil {
		return err
	}

	IdSeller = pengiriman.IdSeller
	if err := read.WithContext(ctx).Model(&sot_models.AlamatPengguna{}).Select("id_pengguna").Where(&sot_models.AlamatPengguna{
		ID: pengiriman.IdAlamatPengguna,
	}).Limit(1).Take(&IdPengguna).Error; err != nil {
		return err
	}

	// ðŸ”” NOTIFIKASI PENGGUNA (PICKED UP)
	var NotifPengguna = notification_models.NotificationPengguna{
		IDPengguna: IdPengguna,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸ“¦ Paketmu Mulai Jalan!",
		Pesan:      "Hore! Paketmu sudah di-pickup oleh kurir dari toko seller dan sedang dalam perjalanan menuju lokasimu.",
		Pop:        3.0,
		Activity:   true,
		Inbox:      false,
		Archive:    true,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": Objek.IdPengiriman, "status": "picked_up"},
			Special:  map[string]interface{}{"click_action": "TRACK_DELIVERY"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)

	// ðŸ”” NOTIFIKASI SELLER (PICKED UP)
	var NotifSeller = notification_models.NotificationSeller{
		IDSeller:  IdSeller,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸšš Paket Berhasil Diserahkan",
		Pesan:     "Mantap! Kurir sudah melakukan pickup berkas paket pesanan pembeli dari tokomu.",
		Pop:       3.0,
		Activity:  true,
		Inbox:     false,
		Archive:   true,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": Objek.IdPengiriman, "status": "picked_up"},
			Special:  map[string]interface{}{"click_action": "MANAGE_ORDER"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// ==========================================
// 2. MEDIA PENGIRIMAN SAMPAI FOTO (KURIR INTERNAL)
// ==========================================
func CreateTambahMediaPengirimanSampaiFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaPengirimanSampaiFoto"
	var Objek sot_models.MediaPengirimanSampaiFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaPengirimanSampaiFoto = cass_models.MediaPengirimanSampaiFoto{
		ID:           Objek.ID,
		IdPengiriman: Objek.IdPengiriman,
		Key:          Objek.Key,
		Format:       Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	var (
		IdSeller   int64 = 0
		IdPengguna int64 = 0
	)

	var pengiriman sot_models.Pengiriman
	if err := read.WithContext(ctx).Model(&sot_models.Pengiriman{}).Select("id_seller, id_alamat_pengguna").Where(&sot_models.Pengiriman{
		ID: Objek.IdPengiriman,
	}).Limit(1).Take(&pengiriman).Error; err != nil {
		return err
	}

	IdSeller = pengiriman.IdSeller
	if err := read.WithContext(ctx).Model(&sot_models.AlamatPengguna{}).Select("id_pengguna").Where(&sot_models.AlamatPengguna{
		ID: pengiriman.IdAlamatPengguna,
	}).Limit(1).Take(&IdPengguna).Error; err != nil {
		return err
	}

	// ðŸ”” NOTIFIKASI PENGGUNA (ARRIVED)
	var NotifPengguna = notification_models.NotificationPengguna{
		IDPengguna: IdPengguna,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸŽ‰ Paketmu Sudah Sampai!",
		Pesan:      "Kurir telah menyerahkan paket di lokasi tujuan. Silakan cek bukti foto penyerahan dan pastikan kondisi barang aman ya!",
		Pop:        3.0,
		Activity:   true,
		Inbox:      false,
		Archive:    true,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": Objek.IdPengiriman, "status": "arrived"},
			Special:  map[string]interface{}{"click_action": "VIEW_PROOF_OF_DELIVERY"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)

	// ðŸ”” NOTIFIKASI SELLER (ARRIVED)
	var NotifSeller = notification_models.NotificationSeller{
		IDSeller:  IdSeller,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ Pesanan Selesai Diantar",
		Pesan:     "Paket kirimanmu telah sukses diserahkan ke tangan pembeli oleh pihak kurir.",
		Pop:       3.0,
		Activity:  true,
		Inbox:     false,
		Archive:   true,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": Objek.IdPengiriman, "status": "arrived"},
			Special:  map[string]interface{}{"click_action": "MANAGE_ORDER"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// ==========================================
// 3. MEDIA PENGIRIMAN EKSPEDISI PICKED UP FOTO
// ==========================================
func CreateTambahMediaPengirimanEkspedisiPickedUpFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaPengirimanEkspedisiPickedUpFoto"
	var Objek sot_models.MediaPengirimanEkspedisiPickedUpFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaPengirimanEkspedisiPickedUpFoto = cass_models.MediaPengirimanEkspedisiPickedUpFoto{
		ID:                    Objek.ID,
		IdPengirimanEkspedisi: Objek.IdPengirimanEkspedisi,
		Key:                   Objek.Key,
		Format:                Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// Tracing IdPengiriman dari tabel pengiriman_ekspedisis
	var IdPengiriman int64 = 0
	if err := read.WithContext(ctx).Table("pengiriman_ekspedisis").Select("id_pengiriman").Where("id = ?", Objek.IdPengirimanEkspedisi).Limit(1).Scan(&IdPengiriman).Error; err != nil {
		return err
	}

	var (
		IdSeller   int64 = 0
		IdPengguna int64 = 0
	)

	var pengiriman sot_models.Pengiriman
	if err := read.WithContext(ctx).Model(&sot_models.Pengiriman{}).Select("id_seller, id_alamat_pengguna").Where("id = ?", IdPengiriman).Limit(1).Take(&pengiriman).Error; err != nil {
		return err
	}

	IdSeller = pengiriman.IdSeller
	if err := read.WithContext(ctx).Model(&sot_models.AlamatPengguna{}).Select("id_pengguna").Where("id = ?", pengiriman.IdAlamatPengguna).Limit(1).Take(&IdPengguna).Error; err != nil {
		return err
	}

	// ðŸ”” NOTIFIKASI PENGGUNA (EKSPEDISI PICKED UP)
	var NotifPengguna = notification_models.NotificationPengguna{
		IDPengguna: IdPengguna,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸ“¦ Paket Diserahkan ke Ekspedisi",
		Pesan:      "Paket pesananmu kini sudah di-pickup oleh armada logistik ekspedisi rekanan dan segera bergerak menuju kota tujuan.",
		Pop:        3.0,
		Activity:   true,
		Inbox:      false,
		Archive:    true,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": IdPengiriman, "status": "expedition_picked_up"},
			Special:  map[string]interface{}{"click_action": "TRACK_EXPEDITION"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)

	// ðŸ”” NOTIFIKASI SELLER (EKSPEDISI PICKED UP)
	var NotifSeller = notification_models.NotificationSeller{
		IDSeller:  IdSeller,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸšš Serah Terima Ekspedisi Berhasil",
		Pesan:     "Bukti serah terima unit paket pesanan pembeli ke kurir ekspedisi sudah terdata valid di sistem.",
		Pop:       3.0,
		Activity:  true,
		Inbox:     false,
		Archive:   true,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": IdPengiriman, "status": "expedition_picked_up"},
			Special:  map[string]interface{}{"click_action": "MANAGE_ORDER"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

// ==========================================
// 4. MEDIA PENGIRIMAN EKSPEDISI SAMPAI AGENT FOTO
// ==========================================
func CreateTambahMediaPengirimanEkspedisiSampaiAgentFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaPengirimanEkspedisiSampaiAgentFoto"
	var Objek sot_models.MediaPengirimanEkspedisiSampaiAgentFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaPengirimanEkspedisiSampaiAgentFoto = cass_models.MediaPengirimanEkspedisiSampaiAgentFoto{
		ID:                    Objek.ID,
		IdPengirimanEkspedisi: Objek.IdPengirimanEkspedisi,
		Key:                   Objek.Key,
		Format:                Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	// Tracing IdPengiriman dari tabel pengiriman_ekspedisis
	var IdPengiriman int64 = 0
	if err := read.WithContext(ctx).Table("pengiriman_ekspedisis").Select("id_pengiriman").Where("id = ?", Objek.IdPengirimanEkspedisi).Limit(1).Scan(&IdPengiriman).Error; err != nil {
		return err
	}

	var (
		IdSeller   int64 = 0
		IdPengguna int64 = 0
	)

	var pengiriman sot_models.Pengiriman
	if err := read.WithContext(ctx).Model(&sot_models.Pengiriman{}).Select("id_seller, id_alamat_pengguna").Where("id = ?", IdPengiriman).Limit(1).Take(&pengiriman).Error; err != nil {
		return err
	}

	IdSeller = pengiriman.IdSeller
	if err := read.WithContext(ctx).Model(&sot_models.AlamatPengguna{}).Select("id_pengguna").Where("id = ?", pengiriman.IdAlamatPengguna).Limit(1).Take(&IdPengguna).Error; err != nil {
		return err
	}

	// ðŸ”” NOTIFIKASI PENGGUNA (ARRIVED AT AGENT)
	var NotifPengguna = notification_models.NotificationPengguna{
		IDPengguna: IdPengguna,
		Pengirim:   notification_seeders.Sistem,
		Judul:      "ðŸ¢ Paket Tiba di Agen Transit",
		Pesan:      "Paket kirimanmu sudah mendarat di gudang/agen ekspedisi terdekat dari lokasimu dan segera dijadwalkan untuk pengantaran kurir lokal.",
		Pop:        3.0,
		Activity:   true,
		Inbox:      false,
		Archive:    true,
		CreatedAt:  time.Now().Format(time.RFC3339),
		ExpiredAt:  time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": IdPengiriman, "status": "arrived_at_agent"},
			Special:  map[string]interface{}{"click_action": "TRACK_EXPEDITION"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifPengguna, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.PenggunaPathNotifikasiMasuk)

	// ðŸ”” NOTIFIKASI SELLER (ARRIVED AT AGENT)
	var NotifSeller = notification_models.NotificationSeller{
		IDSeller:  IdSeller,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸ“ Paket Memasuki Hub Tujuan",
		Pesan:     "Paket pesanan pembeli yang kamu kirim via ekspedisi terpantau sudah sampai di gudang agen transit kota tujuan.",
		Pop:       3.0,
		Activity:  true,
		Inbox:     false,
		Archive:   true,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{"id_pengiriman": IdPengiriman, "status": "arrived_at_agent"},
			Special:  map[string]interface{}{"click_action": "MANAGE_ORDER"},
		},
	}
	_ = notification_request.PostToNotification(ctx, NotifSeller, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.SellerPathNotifikasiMasuk)

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
