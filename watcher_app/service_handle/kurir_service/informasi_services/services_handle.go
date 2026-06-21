package informasi_kurir_handle

import (
	"context"
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

	cass_cud "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/cud"
	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func CreateAjukanInformasiKendaraan(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services = "CreateAjukanInformasiKendaraan"
	var Objek sot_models.InformasiKendaraanKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.InformasiKendaraanKurir = cass_models.InformasiKendaraanKurir{
		ID:             Objek.ID,
		IDkurir:        Objek.IDkurir,
		JenisKendaraan: Objek.JenisKendaraan,
		NamaKendaraan:  Objek.NamaKendaraan,
		RodaKendaraan:  Objek.RodaKendaraan,
		STNK:           Objek.STNK,
		BPKB:           Objek.BPKB,
		NoRangka:       Objek.NoRangka,
		NoMesin:        Objek.NoMesin,
		Status:         Objek.Status,
		CreatedAt:      Objek.CreatedAt,
		UpdatedAt:      Objek.UpdatedAt,
		DeletedAt:      Objek.DeletedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		fmt.Println("Gagal memasukan data ke dalam sot replica async dalam services" + handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		fmt.Println("Gagal memasukan data ke dalam historical dalam services" + handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateEditInformasiKendaraan(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.InformasiKendaraanKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateAjukanInformasiKurir(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.InformasiKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateEditInformasiKurir(Data mb_cud_serializer.ParsedDataMessage) error {
	var Objek sot_models.InformasiKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
