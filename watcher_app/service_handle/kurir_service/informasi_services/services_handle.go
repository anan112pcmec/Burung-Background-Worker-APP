package informasi_kurir_handle

import (
	"context"
	"fmt"
	"strings"
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

func CreateAjukanInformasiKendaraan(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services = "CreateAjukanInformasiKendaraan"
	var Objek sot_models.InformasiKendaraanKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.InformasiKendaraanKurir = cass_models.InformasiKendaraanKurir{
		ID:             Objek.ID,
		IDkurir:        Objek.IDkurir,
		JenisKendaraan: Objek.JenisKendaraan,
		NamaKendaraan:  Objek.NamaKendaraan,
		RodaKendaraan:  Objek.RodaKendaraan,
		STNK:           Objek.STNK,
		BPKB:           Objek.BPKB,
		NoRangka:       Objek.NoRangka,
		NoMesin:        Objek.NoMesin,
		Status:         Objek.Status,
		CreatedAt:      Objek.CreatedAt,
		UpdatedAt:      Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		fmt.Println("Gagal memasukan data ke dalam sot replica async dalam services" + handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		fmt.Println("Gagal memasukan data ke dalam historical dalam services" + handle_services)
	}

	var NamaKurir string = ""
	if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where(&sot_models.Kurir{
		ID: Objek.IDkurir,
	}).Limit(1).Take(&NamaKurir).Error; err != nil {
		return err
	}

	if NamaKurir == "" {
		fmt.Println("Gagal mendapatkan nama kurir")
	}

	var JudulNotif string = strings.Trim(handle_services, "Create")

	var Notifikasi notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   Objek.ID,
		Pengirim:  notification_seeders.Sistem,
		Judul:     JudulNotif,
		Pesan:     fmt.Sprintf("Halo %s, Terimakasih telah mengisi informasi, kendaraanmu %s dengan no rangka %s akan secepatnya kami proses, supaya kamu bisa mulai narik", NamaKurir, Objek.NamaKendaraan, Objek.NoRangka),
		Pop:       0.8,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"informasi_kendaraan_id": Objek.ID,
				"action_type":            "create",
				"platform":               "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REDIRECT_TO_DASHBOARD",
			},
		},
	}

	if err := notification_request.PostToNotification[notification_models.NotificationKurir](ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		return err
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateEditInformasiKendaraan(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditInformasiKendaraan"
	var Objek sot_models.InformasiKendaraanKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.InformasiKendaraanKurir = cass_models.InformasiKendaraanKurir{
		ID:             Objek.ID,
		IDkurir:        Objek.IDkurir,
		JenisKendaraan: Objek.JenisKendaraan,
		NamaKendaraan:  Objek.NamaKendaraan,
		RodaKendaraan:  Objek.RodaKendaraan,
		STNK:           Objek.STNK,
		BPKB:           Objek.BPKB,
		NoRangka:       Objek.NoRangka,
		NoMesin:        Objek.NoMesin,
		Status:         Objek.Status,
		CreatedAt:      Objek.CreatedAt,
		UpdatedAt:      Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		fmt.Println("Gagal memasukan data ke dalam sot replica async dalam services " + handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		fmt.Println("Gagal memasukan data ke dalam historical dalam services " + handle_services)
	}

	// Ambil nama kurir biar personal
	var NamaKurir string = ""
	if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where(&sot_models.Kurir{
		ID: Objek.IDkurir,
	}).Limit(1).Take(&NamaKurir).Error; err != nil {
		return err
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// Copywriting dinamis berdasarkan status approval dari Admin
	var pesanUpdate string
	var judulUpdate string = "ðŸ”„ Update Informasi Kendaraan"
	statusDokumen := strings.ToUpper(Objek.Status)

	switch statusDokumen {
	case "APPROVED":
		judulUpdate = "âœ… Kendaraan Lu Berhasil Diverifikasi!"
		pesanUpdate = fmt.Sprintf("Mantap %s! Pengajuan kendaraan %s lu udah disetujui tim internal. Siap-siap dapet orderan gacor ya!", NamaKurir, Objek.NamaKendaraan)
	case "REJECTED":
		judulUpdate = "âŒ Dokumen Kendaraan Ditolak"
		pesanUpdate = fmt.Sprintf("Waduh %s, dokumen info kendaraan %s lu ditolak nih. Coba cek lagi kesesuaian nomor STNK/BPKB lu ya.", NamaKurir, Objek.NamaKendaraan)
	default: // PENDING / Perubahan data manual dari Kurir sendiri
		pesanUpdate = fmt.Sprintf("Halo %s, perubahan data untuk kendaraan %s telah disimpan dan sedang ditinjau ulang oleh tim kami.", NamaKurir, Objek.NamaKendaraan)
	}

	var NotifikasiUpdate notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   Objek.IDkurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     judulUpdate,
		Pesan:     pesanUpdate,
		Pop:       3.5,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 1, 0).Format(time.RFC3339), // Simpan 1 bulan jika disetujui/ditolak
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"informasi_kendaraan_id": Objek.ID,
				"action_type":            "update_kendaraan",
				"platform":               "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REDIRECT_TO_VEHICLE_DETAIL",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotifikasiUpdate, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		return err
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateAjukanInformasiKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateAjukanInformasiKurir"
	var Objek sot_models.InformasiKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.InformasiKurir = cass_models.InformasiKurir{
		ID:           Objek.ID,
		IDkurir:      Objek.IDkurir,
		TanggalLahir: Objek.TanggalLahir,
		Alasan:       Objek.Alasan,
		Ktp:          Objek.Ktp,
		InformasiSim: Objek.InformasiSim,
		Status:       Objek.Status,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		fmt.Println("Gagal memasukan data ke dalam sot replica async dalam services " + handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		fmt.Println("Gagal memasukan data ke dalam historical dalam services " + handle_services)
	}

	// Ambil nama kurir
	var NamaKurir string = ""
	if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where(&sot_models.Kurir{
		ID: Objek.IDkurir,
	}).Limit(1).Take(&NamaKurir).Error; err != nil {
		return err
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	var Notifikasi notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   Objek.IDkurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     "ðŸªª Pengajuan Data Profil Kurir",
		Pesan:     fmt.Sprintf("Halo %s, data KTP dan SIM lu berhasil kami terima. Tim legal kami bakal nge-validasi data lu secepatnya ya!", NamaKurir),
		Pop:       3.0,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, 7).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"informasi_kurir_id": Objek.ID,
				"action_type":        "create_informasi_diri",
				"platform":           "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REDIRECT_TO_VERIFICATION_STATUS",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, Notifikasi, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		return err
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateEditInformasiKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, read *gorm.DB, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateEditInformasiKurir"
	var Objek sot_models.InformasiKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.InformasiKurir = cass_models.InformasiKurir{
		ID:           Objek.ID,
		IDkurir:      Objek.IDkurir,
		TanggalLahir: Objek.TanggalLahir,
		Alasan:       Objek.Alasan,
		Ktp:          Objek.Ktp,
		InformasiSim: Objek.InformasiSim,
		Status:       Objek.Status,
		CreatedAt:    Objek.CreatedAt,
		UpdatedAt:    Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		fmt.Println("Gagal memasukan data ke dalam sot replica async dalam services " + handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		fmt.Println("Gagal memasukan data ke dalam historical dalam services " + handle_services)
	}

	// Ambil nama kurir
	var NamaKurir string = ""
	if err := read.WithContext(ctx).Model(&sot_models.Kurir{}).Select("nama").Where(&sot_models.Kurir{
		ID: Objek.IDkurir,
	}).Limit(1).Take(&NamaKurir).Error; err != nil {
		return err
	}
	if NamaKurir == "" {
		NamaKurir = "Kurir"
	}

	// Copywriting dinamis berdasarkan status verifikasi profil diri kurir
	var pesanDiri string
	var judulDiri string = "ðŸ†” Perubahan Profil Kurir"
	statusVerifikasi := strings.ToUpper(Objek.Status)

	switch statusVerifikasi {
	case "APPROVED":
		judulDiri = "ðŸŽ‰ Akun Kurir Lu Resmi Aktif!"
		pesanDiri = fmt.Sprintf("Selamat %s, dokumen identitas (KTP/SIM) lu udah lolos verifikasi sistem. Sekarang status lu resmi jadi kurir aktif. Yuk, gas pol cari orderan!", NamaKurir)
	case "REJECTED":
		judulDiri = "âš ï¸ Verifikasi Akun Tertunda"
		pesanDiri = fmt.Sprintf("Mohon maaf %s, data identitas yang lu kirim belum cocok dengan standar kami. Tolong upload ulang foto KTP/SIM dengan pencahayaan yang jelas ya.", NamaKurir)
	default:
		pesanDiri = fmt.Sprintf("Halo %s, pembaruan dokumen KTP/SIM berhasil disimpan dan masuk ke dalam antrean review admin.", NamaKurir)
	}

	var NotifikasiUpdate notification_models.NotificationKurir = notification_models.NotificationKurir{
		IDKurir:   Objek.IDkurir,
		Pengirim:  notification_seeders.Sistem,
		Judul:     judulDiri,
		Pesan:     pesanDiri,
		Pop:       4.5, // Pop-up agak lamaan dikit biar dibaca seksama kalo urusan status akun
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 1, 0).Format(time.RFC3339),
		Data: struct {
			Metadata map[string]interface{} `json:"metadata"`
			Special  interface{}            `json:"special"`
		}{
			Metadata: map[string]interface{}{
				"informasi_kurir_id": Objek.ID,
				"action_type":        "update_informasi_diri",
				"platform":           "kurir_mobile_app",
			},
			Special: map[string]interface{}{
				"click_action": "REDIRECT_TO_ACCOUNT_DASHBOARD",
			},
		},
	}

	if err := notification_request.PostToNotification(ctx, NotifikasiUpdate, cache.HostRunningAPIInNotifikasi, cache.PortRunningAPIInNotifikasi, cache.KurirPathNotifikasiMasuk); err != nil {
		return err
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
