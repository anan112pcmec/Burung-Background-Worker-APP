package social_media_seller_handle

import (
	"fmt"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func CreateEngageSocialMediaSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.EntitySocialMedia
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdatedngageSocialMediaSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.EntitySocialMedia
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}
