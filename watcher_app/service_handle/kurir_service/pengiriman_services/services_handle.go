package pengiriman_kurir_handle

import (
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func UpdateAktifkanBidKurir(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirData

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

func UpdateAmbilPengirimanNonEksManualRegulerIIPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateAmbilPengirimanNonEksManualRegulerIIbidKurirDataUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
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

func UpdateAmbilPengirimanEksManualRegulerIIbidKurirDataStatusUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
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

func UpdateLockSiapAntarBidKurirIIbidKurirDataUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateLockSiapAntarBidKurirIIkurirUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
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

func UpdatePickedUpPengirimanNonEksIIdataSchedulerUpdated(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatedPickedUpPengirimanNonEksIIpengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatedPickedUpPengirimanNonEksIItransaksiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanNonEksIIbidKurirSchedulerUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirNonEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanNonEksIIpengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanNonEksIIjejakPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
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

func UpdateSampaiPengirimanNonEksIIpengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Pengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanNonEksIIbidKurirDataUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanNonEksIIjejakPengirimanUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.JejakPengiriman

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanNonEksIItransaksiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
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

func UpdatePickedUpPengirimanEksIIschedulerUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirEksScheduler

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanEksIIpengirimanEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.PengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePickedUpPengirimanEksIItransaksiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKirimPengirimanEksIIschedulerUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
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

func UpdateKirimPengirimanEksIIjejakPengirimanEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.JejakPengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateInformasiPerjalananPengiimanEks(Data mb_cud_serializer.ParsedDataMessage) error {
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

func UpdateSampaiPengirimanEksIIpengirimanEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.PengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIbidKurirDataUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BidKurirData

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIjejakPengirimanEksUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.JejakPengirimanEkspedisi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIItransaksiUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Transaksi

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIpayOutKurirCreatePublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.PayOutKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateSampaiPengirimanEksIIkurirUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
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

func UpdateNonaktifkanBidKurirIIkurirUpdatedPublish(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Kurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
