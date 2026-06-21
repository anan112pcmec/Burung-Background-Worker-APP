package historical_format

import "time"

type Pencatatan struct {
	TahunUpdate int
	BulanUpdate int
	EventTime   time.Time
}

func Sekarang() Pencatatan {
	now := time.Now()

	return Pencatatan{
		TahunUpdate: now.Year(),
		BulanUpdate: int(now.Month()),
		EventTime:   now,
	}
}

func PencatatanCombine(dataPencatatan Pencatatan, data map[string]interface{}) {
	data["tahun_update"] = dataPencatatan.TahunUpdate
	data["bulan_update"] = dataPencatatan.BulanUpdate
	data["event_time"] = dataPencatatan.EventTime

}
