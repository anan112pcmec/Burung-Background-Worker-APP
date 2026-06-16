package transaksi_seller_handle

import (
	"fmt"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"

)

func UpdateApproveOrderTransaksi(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.Transaksi
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateKirimOrderTransaksiEkspedisi(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.PengirimanEkspedisi
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateKirimOrderTransaksiBiasa(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.Pengiriman
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdatedUnApproveOrderTransaksi(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.Transaksi
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}
