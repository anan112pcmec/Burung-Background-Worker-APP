package rekening_kurir_handle

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

func CreateMasukanRekeningKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateMasukanRekeningKurir"
	var Objek sot_models.RekeningKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.RekeningKurir = cass_models.RekeningKurir{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		NamaBank:        Objek.NamaBank,
		NomorRekening:   Objek.NomorRekening,
		PemilikRekening: Objek.PemilikRekening,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func UpdateEditRekeningKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdatedEditRekeningKurir"
	var Objek sot_models.RekeningKurir
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.RekeningKurir = cass_models.RekeningKurir{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		NamaBank:        Objek.NamaBank,
		NomorRekening:   Objek.NomorRekening,
		PemilikRekening: Objek.PemilikRekening,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusRekeningKurir(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusRekeningKurir"
	var Objek sot_models.RekeningKurir

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.RekeningKurir = cass_models.RekeningKurir{
		ID:              Objek.ID,
		IdKurir:         Objek.IdKurir,
		NamaBank:        Objek.NamaBank,
		NomorRekening:   Objek.NomorRekening,
		PemilikRekening: Objek.PemilikRekening,
		CreatedAt:       Objek.CreatedAt,
		UpdatedAt:       Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), ObjekCass.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}
