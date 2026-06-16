package barang_pengguna_handle

import (
	"fmt"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func CreateLikesBarang(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.BarangDisukai
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.IdPengguna)
	return nil
}

func DeleteUnlikesBarang(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.BarangDisukai
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.IdPengguna)
	return nil
}

func CreateMasukanKomentarBarang(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.Komentar
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}

	return nil
}

func UpdateEditKomentarBarang(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek int64
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusKomentarBarang(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.Komentar
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateMasukanChildKomentar(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.KomentarChild
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateMentionChildKomentar(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.KomentarChild
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateEditChildKomentar(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek int64
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusChildKomentar(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.KomentarChild
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahKeranjangBarang(Data mb_cud_serializer.ParsedDataMessage) error {

	var Objek sot_models.Keranjang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateEditKeranjangBarang(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek int64
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusKeranjangBarang(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.Keranjang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateBerikanReviewBarang(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.Review
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateLikeReviewBarang(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek int64
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateUnlikeReviewBarang(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek int64
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}
