package historical_models

import "time"

type Pencatatan struct {
	TahunUpdate int
	BulanUpdate int
	EventTime   time.Time
}

func (p Pencatatan) Sekarang() *Pencatatan {
	now := time.Now()

	return &Pencatatan{
		TahunUpdate: now.Year(),
		BulanUpdate: int(now.Month()),
		EventTime:   now,
	}
}
