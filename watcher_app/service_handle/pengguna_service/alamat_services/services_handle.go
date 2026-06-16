package alamat_pengguna_handle

import (
	"fmt"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"

)

// data body yang diinput merupakan model relasi alamat pengguna
func CreateAlamatPub(Data mb_cud_serializer.ParsedDataMessage) error {
	var DataAlamat sot_models.AlamatPengguna
	if err := helper.DecodeJSONBody(Data, &DataAlamat); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", DataAlamat.IDPengguna)
	return nil
}

func UpdateAlamatPub(Data mb_cud_serializer.ParsedDataMessage) error {
	var DataAlamat sot_models.AlamatPengguna
	if err := helper.DecodeJSONBody(Data, &DataAlamat); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", DataAlamat.IDPengguna)
	return nil
}

func DeleteAlamatPub(Data mb_cud_serializer.ParsedDataMessage) error {
	var DataAlamat sot_models.AlamatPengguna
	if err := helper.DecodeJSONBody(Data, &DataAlamat); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", DataAlamat.IDPengguna)
	return nil
}
