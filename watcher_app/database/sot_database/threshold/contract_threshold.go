package sot_threshold

import (
	"context"

	"gorm.io/gorm"
)

type CustomCounter struct {
	FieldName string
	Count     int
}

type ThresholdContract interface {
	Inisialisasi(int64, context.Context, *gorm.DB) error
	Increment(int64, context.Context, gorm.DB, ...string) error
	Decrement(int64, context.Context, gorm.DB, ...string) error
	CustomIncrement(int64, context.Context, gorm.DB, []CustomCounter) error
	CustomDecrement(int64, context.Context, gorm.DB, []CustomCounter) error
}

func InisialisasiThreshold(t ThresholdContract, id_fk int64, ctx context.Context, db *gorm.DB) error {
	return t.Inisialisasi(id_fk, ctx, db)
}
