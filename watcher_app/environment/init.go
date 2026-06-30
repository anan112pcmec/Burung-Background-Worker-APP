package environment

import "os"

func init() {
	PenggunaPathNotifikasiMasuk = os.Getenv("PENGGUNAPATHNOTIFIKASIMASUK")
	SellerPathNotifikasiMasuk = os.Getenv("SELLERPATHNOTIFIKASIMASUK")
	KurirPathNotifikasiMasuk = os.Getenv("KURIRPATHNOTIFIKASIMASUK")
	PortRunningAPIInNotifikasi = os.Getenv("APIINNOTIFIKASIPORT")
	HostRunningAPIInNotifikasi = os.Getenv("APIINNOTIFIKASIHOST")

}
