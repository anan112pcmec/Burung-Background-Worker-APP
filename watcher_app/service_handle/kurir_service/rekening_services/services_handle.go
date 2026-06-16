package rekening_kurir_handle

import (
	"fmt"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"

)

func CreateMasukanRekeningKurir(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.RekeningKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateEditRekeningKurir(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.RekeningKurir
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusRekeningKurir(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.RekeningKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
