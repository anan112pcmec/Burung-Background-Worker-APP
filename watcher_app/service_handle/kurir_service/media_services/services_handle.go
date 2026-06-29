package media_kurir_handle

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

func UpdateUbahKurirProfilFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateUbahKurirProfilFoto"
	var Objek sot_models.MediaKurirProfilFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaKurirProfilFoto = cass_models.MediaKurirProfilFoto{
		ID:        Objek.ID,
		IdKurir:   Objek.IdKurir,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
		UpdatedAt: Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		if err := cass_cud.InsertData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), parsedData); err != nil {
			return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
		}
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusKurirProfilFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusKurirProfilFoto"
	var Objek sot_models.MediaKurirProfilFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaKurirProfilFoto = cass_models.MediaKurirProfilFoto{
		ID:        Objek.ID,
		IdKurir:   Objek.IdKurir,
		Key:       Objek.Key,
		Format:    Objek.Format,
		CreatedAt: Objek.CreatedAt,
		UpdatedAt: Objek.UpdatedAt,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID); err != nil {
		return fmt.Errorf("gagal menghapus data dari  sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahMediaInformasiKendaraanKurirKendaraanFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaInformasiKendaraanKurirKendaraanFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirKendaraanFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirKendaraanFoto = cass_models.MediaInformasiKendaraanKurirKendaraanFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
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

func UpdateTambahMediaInformasiKendaraanKurirKendaraanFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateTambahMediaInformasiKendaraanKurirKendaraanFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirKendaraanFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirKendaraanFoto = cass_models.MediaInformasiKendaraanKurirKendaraanFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusMediaInformasiKendaraanKurirKendaraanFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusMediaInformasiKendaraanKurirKendaraanFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirKendaraanFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirKendaraanFoto = cass_models.MediaInformasiKendaraanKurirKendaraanFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahInformasiKendaraanKurirBPKBFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahInformasiKendaraankurirBPKBFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirBPKBFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirBPKBFoto = cass_models.MediaInformasiKendaraanKurirBPKBFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
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

func UpdateTambahInformasiKendaraanKurirBPKBFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateTambahInformasiKendaraanKurirBPKBFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirBPKBFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirBPKBFoto = cass_models.MediaInformasiKendaraanKurirBPKBFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusInformasiKendaraanKurirBPKBFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusInformasiKendaraanKurirBPKBFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirBPKBFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirBPKBFoto = cass_models.MediaInformasiKendaraanKurirBPKBFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID); err != nil {
		return fmt.Errorf("gagal menghapus data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahInformasiKendaraanKurirSTNKFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahInformasiKendaraanKurirSTNKFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirSTNKFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirSTNKFoto = cass_models.MediaInformasiKendaraanKurirSTNKFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
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

func UpdateTambahInformasiKendaraanKurirSTNKFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateTambahInformasiKendaraanKurirSTNKFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirSTNKFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirSTNKFoto = cass_models.MediaInformasiKendaraanKurirSTNKFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusInformasiKendaraanKurirSTNKFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusInformasiKendaraanKurirSTNKFoto"
	var Objek sot_models.MediaInformasiKendaraanKurirSTNKFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKendaraanKurirSTNKFoto = cass_models.MediaInformasiKendaraanKurirSTNKFoto{
		ID:                        Objek.ID,
		IdInformasiKendaraanKurir: Objek.IdInformasiKendaraanKurir,
		Key:                       Objek.Key,
		Format:                    Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahMediaInformasiKurirKTPFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaInformasiKurirKTPFoto"

	var Objek sot_models.MediaInformasiKurirKTPFoto

	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKurirKTPFoto = cass_models.MediaInformasiKurirKTPFoto{
		ID:               Objek.ID,
		IdInformasiKurir: Objek.IdInformasiKurir,
		Key:              Objek.Key,
		Format:           Objek.Format,
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

func UpdateTambahMediaInformasiKurirKTPFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "UpdateTambahMediaInformasiKurirKTPFoto"
	var Objek sot_models.MediaInformasiKurirKTPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKurirKTPFoto = cass_models.MediaInformasiKurirKTPFoto{
		ID:               Objek.ID,
		IdInformasiKurir: Objek.IdInformasiKurir,
		Key:              Objek.Key,
		Format:           Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.UpdateData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID, parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func DeleteHapusMediaInformasiKurirKTPFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "DeleteHapusMediaInformasiKurirKTPFoto"
	var Objek sot_models.MediaInformasiKurirKTPFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaInformasiKurirKTPFoto = cass_models.MediaInformasiKurirKTPFoto{
		ID:               Objek.ID,
		IdInformasiKurir: Objek.IdInformasiKurir,
		Key:              Objek.Key,
		Format:           Objek.Format,
	}

	var parsedData map[string]interface{} = ObjekCass.ParseToCUDType()

	if err := cass_cud.DeleteData(ctx, cass_sot_replica, ObjekCass.TableNameSotReplica(), Objek.ID); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam sot replica async %s dalam %s", err, handle_services)
	}

	historical_format.PencatatanCombine(historical_format.Sekarang(), parsedData)

	if err := cass_cud.InsertData(ctx, cass_historcal, ObjekCass.TableNameHistorical(), parsedData); err != nil {
		return fmt.Errorf("gagal memasukan data ke dalam historical db %s dalam %s", err, handle_services)
	}

	fmt.Println("Berhasil mendapatkan data", Objek.ID)
	return nil
}

func CreateTambahMediaPengirimanPickedUpFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaPengirimanPickedUpFoto"
	var Objek sot_models.MediaPengirimanPickedUpFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaPengirimanPickedUpFoto = cass_models.MediaPengirimanPickedUpFoto{
		ID:           Objek.ID,
		IdPengiriman: Objek.IdPengiriman,
		Key:          Objek.Key,
		Format:       Objek.Format,
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

func CreateTambahMediaPengirimanSampaiFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaPengirimanSampaiFoto"
	var Objek sot_models.MediaPengirimanSampaiFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaPengirimanSampaiFoto = cass_models.MediaPengirimanSampaiFoto{
		ID:           Objek.ID,
		IdPengiriman: Objek.IdPengiriman,
		Key:          Objek.Key,
		Format:       Objek.Format,
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

func CreateTambahMediaPengirimanEkspedisiPickedUpFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaPengirimanEkspedisiPickedUpFoto"
	var Objek sot_models.MediaPengirimanEkspedisiPickedUpFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaPengirimanEkspedisiPickedUpFoto = cass_models.MediaPengirimanEkspedisiPickedUpFoto{
		ID:                    Objek.ID,
		IdPengirimanEkspedisi: Objek.IdPengirimanEkspedisi,
		Key:                   Objek.Key,
		Format:                Objek.Format,
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

func CreateTambahMediaPengirimanEkspedisiSampaiAgentFoto(Data mb_cud_serializer.ParsedDataMessage, ctx context.Context, cass_historcal, cass_sot_replica *gocql.Session) error {
	const handle_services string = "CreateTambahMediaPengirimanEkspedisiSampaiAgentFoto"
	var Objek sot_models.MediaPengirimanEkspedisiSampaiAgentFoto
	if err := helper.DecodeJSONBody(Data, &Objek); err != nil {
		return fmt.Errorf("gagal mengolah data alamat")
	}

	var ObjekCass cass_models.MediaPengirimanEkspedisiSampaiAgentFoto = cass_models.MediaPengirimanEkspedisiSampaiAgentFoto{
		ID:                    Objek.ID,
		IdPengirimanEkspedisi: Objek.IdPengirimanEkspedisi,
		Key:                   Objek.Key,
		Format:                Objek.Format,
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
