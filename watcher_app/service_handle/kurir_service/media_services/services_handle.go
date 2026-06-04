package media_kurir_handle

import (
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"

)

func UpdateUbahKurirProfilFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaKurirProfilFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusKurirProfilFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaKurirProfilFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahMediaInformasiKendaraanKurirKendaraanFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaInformasiKendaraanKurirKendaraanFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateTambahMediaInformasiKendaraanKurirKendaraanFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaInformasiKendaraanKurirKendaraanFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusMediaInformasiKendaraanKurirKendaraanFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaInformasiKendaraanKurirKendaraanFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahInformasiKendaraanKurirBPKBFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaInformasiKendaraanKurirBPKBFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateTambahInformasiKendaraanKurirBPKBFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaInformasiKendaraanKurirBPKBFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusInformasiKendaraanKurirBPKBFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaInformasiKendaraanKurirBPKBFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahInformasiKendaraanKurirSTNKFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaInformasiKendaraanKurirSTNKFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateTambahInformasiKendaraanKurirSTNKFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaInformasiKendaraanKurirSTNKFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusInformasiKendaraanKurirSTNKFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaInformasiKendaraanKurirSTNKFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahMediaInformasiKurirKTPFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaInformasiKurirKTPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateTambahMediaInformasiKurirKTPFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaInformasiKurirKTPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusMediaInformasiKurirKTPFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaInformasiKurirKTPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahMediaPengirimanPickedUpFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaPengirimanPickedUpFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahMediaPengirimanSampaiFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaPengirimanSampaiFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahMediaPengirimanEkspedisiPickedUpFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaPengirimanEkspedisiPickedUpFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahMediaPengirimanEkspedisiSampaiAgentFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek models.MediaPengirimanEkspedisiSampaiAgentFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
