package alamat_kurir_handle

import (
	"fmt"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func CreateMasukanAlamatKurir(Data mb_cud_serializer.ParsedDataMessage) error {
	var DataAlamatKurir sot_models.AlamatKurir
	if err := helper.DecodeJSONBody(Data, &DataAlamatKurir); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", DataAlamatKurir.IdKurir)
	return nil
}

func UpdatedEditAlamatKurir(Data mb_cud_serializer.ParsedDataMessage) error {
	var DataAlamatKurir sot_models.AlamatKurir
	if err := helper.DecodeJSONBody(Data, &DataAlamatKurir); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", DataAlamatKurir.IdKurir)
	return nil
}

func DeleteHapusAlamatKurir(Data mb_cud_serializer.ParsedDataMessage) error {
	var DataAlamatKurir sot_models.AlamatKurir
	if err := helper.DecodeJSONBody(Data, &DataAlamatKurir); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", DataAlamatKurir.IdKurir)
	return nil
}
