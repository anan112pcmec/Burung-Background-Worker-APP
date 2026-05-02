package wishlist_pengguna_handle

import (
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func CreateTambahBarangKeWishlist(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek models.Wishlist
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	return nil
}

func DeleteHapusBarangDariWishlist(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek models.Wishlist
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	return nil
}
