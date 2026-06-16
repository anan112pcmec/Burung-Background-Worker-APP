package credential_pengguna_handle

import (
	"fmt"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func UpdateValidateUbahPasswordPenggunaViaOtp(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.Pengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	return nil
}

func UpdateValidateUbahPasswordPenggunaViaPin(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.Pengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	return nil
}

func CreateMembuatSecretPinPengguna(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.Pengguna
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	return nil
}

func UpdateSecretPinPengguna(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.Komentar
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	return nil
}
