package social_media_pengguna_handle

import (
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func CreateEngageTautkanSocialMediaPengguna(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek models.EntitySocialMedia
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	return nil
}

func UpdateEngageTautkanSocialMediaPengguna(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek models.EntitySocialMedia
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	return nil
}

func UpdateEngageHapusSocialMedia(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek models.EntitySocialMedia
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	return nil
}

func CreateFollowSeller(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek models.Follower
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	return nil
}

func DeleteUnfollowSeller(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek models.Follower
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	return nil
}
