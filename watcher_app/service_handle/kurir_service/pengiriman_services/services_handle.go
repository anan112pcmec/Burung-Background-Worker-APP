package pengiriman_kurir_handle

import (
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func CreateAktifkanBidKurir(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAktifkanBidKurir(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateUpdatePosisiBidKurir(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateAmbilPengirimanNonEksManualRegulerIIBidKurirNonEksSchedulerCreatePublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAmbilPengirimanNonEksManualRegulerIIpengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataAmbilPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataStatusUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateAmbilPengirimanEksManualRegulerIIbidKurirEksSchedulerCreatePublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAmbilPengirimanEksManualRegulerIIpengirimanEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataAmbilPengirimanEksStatusUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateLockSiapAntarBidKurirIIEksScheduler(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateLockSiapAntarBidKurirIINonEksScheduler(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateLockSiapAntarBidKurirIIbidKurirDataLockSiapAntarUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateLockSiapAntarBidKurirIIkurirLockSiapAntarUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreatePickedUpPengirimanNonEksIIjejakPengirimanCreatePublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.JejakPengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanNonEksIIschedulerPickedUpNonEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatedPickedUpPengirimanNonEksIIpengirimanPickedUpNonEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatedPickedUpPengirimanNonEksIItransaksiPickedUpNonEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanNonEksIIbidKurirPengirimanNonEksSchedulerUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanNonEksIIpengirimanPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanNonEksIIjejakpengirimanPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.JejakPengiriman
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateUpdateInformasiPerjalananPengirimanNonEks(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.JejakPengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteSampaiPengirimanNonEksIIbidKurirNonEksDeletePublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanNonEksIIpengirimanSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanNonEksIIbidKurirDataSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanNonEksIIjejakPengirimanSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.JejakPengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanNonEksIItransaksiSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateSampaiPengirimanNonEksIIpayOutSellerCreatePublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.PayOutSeller

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateSampaiPengirimanNonEksIIpayOutKurirCreatePublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.PayOutKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreatePickedUpPengirimanEksIIjejakPengirimanEksCreatePublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.JejakPengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanEksIIschedulerEksPickedUpUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanEksIIpengirimanEksPickedUpUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.PengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanEksIItransaksiPickedUpUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanEksIIschedulerPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanEksIIpengirimanEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.PengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanEksIIpengirimanPengirimanEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.JejakPengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateInformasiPerjalananPengirimanEks(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.JejakPengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteSampaipengirimanEksIIbidKurirEksDeletePublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIpengirimanSampaiEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.PengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIbidKurirDataEksSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIjejakPengirimanEksSampaiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.JejakPengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIItransaksiSampaiEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateSampaiPengirimanEksIIpayOutKurirEksCreatePublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.PayOutKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIkurirUpdatedSampaiEksPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteNonaktifkanBidKurirIIbidKurirDataDeletePublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateNonaktifkanBidKurirIIkurirNonaktifkanBidUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
