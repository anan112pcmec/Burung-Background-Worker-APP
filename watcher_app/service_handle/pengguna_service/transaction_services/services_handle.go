package transaction_pengguna_handle

import (
	"fmt"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"

)

func DeleteCheckoutBarangUser(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.Keranjang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	// bertujuan untuk menghapus keranjang bukan checkout berarti mendelete

	return nil
}

func CreateLockTransaksiVa(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.Transaksi
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateLockTransaksiWallet(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.Transaksi
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateLockTransaksiGerai(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.Transaksi
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}
