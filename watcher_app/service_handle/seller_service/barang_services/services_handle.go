package barang_seller_handle

import (
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func CreateMasukanBarangInduk(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BarangInduk
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateEditBarangInduk(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.BarangInduk
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusBarangInduk(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.AlamatGudang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateEditKategoriBarang(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateMasukanKategoriBarang(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusBarangKategori(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.KategoriBarang
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateMasukanKomentarBarang(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Komentar
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateEditKomentarBarang(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Komentar
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusKomentarBarang(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.Komentar
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateMasukanChildKomentar(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.KomentarChild
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateEditChildKomentar(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.KomentarChild
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusChildKomentar(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.KomentarChild
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}
