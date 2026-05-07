package etalase_seller_handle

import (
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func CreateTambahEtalaseSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Etalase
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateEditEtalaseSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Etalase
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusEtalaseSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Etalase
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahkanBarangKeEtalase(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BarangKeEtalase
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusBarangDariEtalase(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BarangKeEtalase
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}
