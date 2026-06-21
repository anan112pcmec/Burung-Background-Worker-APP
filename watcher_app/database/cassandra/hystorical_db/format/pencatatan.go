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
