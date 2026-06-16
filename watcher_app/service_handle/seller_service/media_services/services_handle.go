package media_seller_handle

import (
	"fmt"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func CreateUbahFotoProfilSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaSellerProfilFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateUbahFotoProfilSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaSellerProfilFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusFotoProfilSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaSellerProfilFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateUbahFotoBannerSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaSellerBannerFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateUbahFotoBannerSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaSellerBannerFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusFotoBannerSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaSellerBannerFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahkanFotoTokoFisikSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaSellerTokoFisikFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusFotoTokoFisikSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaSellerTokoFisikFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateUbahFotoEtalaseSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaEtalaseFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateUbahFotoEtalaseSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaEtalaseFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusFotoEtalaseSeller(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaEtalaseFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahkanMediaBarangIndukFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBarangIndukFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusMediaBarangIndukFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBarangIndukFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateUbahBarangIndukVideo(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBarangIndukVideo
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateUbahBarangIndukVideo(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBarangIndukVideo
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusBarangIndukVideo(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBarangIndukVideo
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateUbahKategoriBarangFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaKategoriBarangFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateUbahKategoriBarangFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaKategoriBarangFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusKategoriBarangFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaKategoriBarangFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahDistributorDataDokumen(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaDistributorDataDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateTambahDistributorDataDokumen(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaDistributorDataDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusMediaDistributorDataDokumen(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaDistributorDataDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahMediaDistributorDataNPWPFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaDistributorDataNPWPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateTambahMediaDistributorDataNPWPFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaDistributorDataNPWPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusMediaDistributorDataNPWPFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaDistributorDataNPWPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahDistributorDataNIBFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaDistributorDataNIBFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateTambahDistributorDataNIBFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaDistributorDataNIBFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusDistributorDataNIBFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaDistributorDataNIBFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahDistributorDataSuratKerjasamaDokumen(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaDistributorDataSuratKerjasamaDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateTambahDistributorDataSuratKerjasamaDokumen(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaDistributorDataSuratKerjasamaDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusDistributorDataSuratKerjasamaDokumen(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaDistributorDataSuratKerjasamaDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahBrandDataPerwakilanDokumen(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataPerwakilanDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateTambahBrandDataPerwakilanDokumen(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataPerwakilanDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusBrandDataPerwakilanDokumen(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataPerwakilanDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahMediaBrandDataSertifikatFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataSertifikatFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateTambahMediaBrandDataSertifikatFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataSertifikatFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusMediaBrandDataSertifikatFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataSertifikatFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahMediaBrandDataNIBFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataNIBFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateTambahMediaBrandDataNIBFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataNIBFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusMediaBrandDataNIBFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataNIBFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahMediaBrandNPWPFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataNPWPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateTambahMediaBrandNPWPFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataNPWPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusMediaBrandNPWPFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataNPWPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahMediaBrandDataLogoFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataLogoFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateTambahMediaBrandDataLogoFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataLogoFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusMediaBrandDataLogoFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataLogoFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahBrandDataSuratKerjasamaDokumen(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataSuratKerjasamaDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func UpdateTambahBrandDataSuratKerjasamaDokumen(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataSuratKerjasamaDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func DeleteHapusBrandDataSuratKerjasamaDokumen(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaBrandDataSuratKerjasamaDokumen
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahMediaTransaksiApprovedFoto(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaTransaksiApprovedFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}

func CreateTambahMediaTransaksiApprovedVideo(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.MediaTransaksiApprovedVideo
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data")
	} else {
		fmt.Println(Objek)
	}
	return nil
}
