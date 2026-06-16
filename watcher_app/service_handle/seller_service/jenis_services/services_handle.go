package jenis_seller_handle

import (
	"fmt"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func CreateMasukanDataDistributor(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.DistributorData
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateEditDataDistributor(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.DistributorData
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusDataDistributor(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.DistributorData
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateMasukanDataBrand(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.BrandData
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateEditDataBrand(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.BrandData
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusDataBrand(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.BrandData
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}
