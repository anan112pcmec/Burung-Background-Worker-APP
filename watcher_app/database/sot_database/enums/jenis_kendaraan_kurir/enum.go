package jenis_kendaraan_kurir

const (
	Motor   = "Motor"
	Mobil   = "Mobil"
	Truk    = "Truk"
	Pickup  = "Pickup"
	Unknown = "Unknown"
)

var JenisKendaraan map[string]string = map[string]string{
	Motor:   Motor,
	Mobil:   Mobil,
	Truk:    Truk,
	Pickup:  Pickup,
	Unknown: Unknown,
}
