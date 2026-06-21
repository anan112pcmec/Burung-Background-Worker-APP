package auth_handle

import (
	"fmt"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func CreateValidatePenggunaRegistration(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.Pengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateValidateSellerRegistration(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.Seller
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateValidateKurirRegistration(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.Kurir
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdatePenggunaLogin(Data mb_cud_serializer.ParsedDataMessage) error {
	var Pengguna sot_models.Pengguna
	if err := helper.DecodeJSONBody(Data, &Pengguna); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Pengguna.ID)
	return nil
}

func UpdateSellerLogin(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.Seller
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateKurirLogin(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.Kurir
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
