package enums

import (
	"fmt"

	"gorm.io/gorm"
)

func UpEnumsEntity(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	enumMap := map[string][]string{
		"jenis_entity": {"pengguna", "seller", "kurir"},
		"status":       {"Online", "Offline"},
		"jenis_seller": {"Brands", "Distributors", "Personal"},
		/* Udh bikin enum */ "seller_dedication": {"Pakaian & Fashion", "Kosmetik & Kecantikan", "Elektronik & Gadget", "Buku & Media", "Makanan & Minuman", "Ibu & Bayi", "Mainan", "Olahraga & Outdoor", "Otomotif & Sparepart", "Rumah Tangga", "Alat Tulis", "Perhiasan & Aksesoris", "Produk Digital", "Bangunan & Perkakas", "Musik & Instrumen", "Film & Broadcasting", "Semua Barang"},

		/*Udh bikin enum*/ "jenis_layanan_kurir": {"Reguler", "Fast", "Instant"},
		"status_keranjang":                       {"Ready", "UnReady"},
		"status_perizinan":                       {"Pending", "Diizinkan", "Dilarang"},
		/*Udh bikin enum*/ "jenis_kendaraan_kurir": {"Motor", "Mobil", "Truk", "Pickup", "Lainnya", "Unknown"},
		"roda_kendaraan_kurir":                     {"2", "3", "4"},
		"status_kurir":                             {"Idle", "OnDelivery", "Off"},
		"status_jenis_seller":                      {"Pending", "Confirmed", "Declined"},
		"status_diskon_produk":                     {"Draft", "Aktif", "Selesai"},
		"status_barang_di_diskon":                  {"Waiting", "Applied"},
		"mode_bid_kurir":                           {"manual", "auto"},
		/* Udh bikin enum */ "nama_provinsi": {"banten", "jawa_barat", "jawa_tengah", "di_yogyakarta", "dki_jakarta", "jawa_timur"},
		/* Udh bikin enum */ "nama_kota": {
			"cilegon",
			"pandeglang",
			"lebak",
			"serang",
			"tangerang",
			"tangerang selatan",

			// Jawa Barat
			"bandung",
			"cimahi",
			"sumedang",
			"garut",
			"bandung barat",
			"cianjur",
			"bekasi",
			"bogor",
			"cirebon",
			"indramayu",
			"kuningan",
			"majalengka",
			"depok",
			"karawang",
			"purwakarta",
			"subang",
			"sukabumi",
			"tasikmalaya",
			"banjar",
			"ciamis",
			"pangandaran",

			// Jawa Tengah
			"cilacap",
			"magelang",
			"kebumen",
			"wonosobo",
			"purworejo",
			"temanggung",
			"surakarta",
			"boyolali",
			"karanganyar",
			"klaten",
			"sragen",
			"sukoharjo",
			"wonogiri",
			"semarang",
			"jepara",
			"kudus",
			"pekalongan",
			"batang",
			"blora",
			"demak",
			"kendal",
			"pati",
			"pemalang",
			"grobogan",
			"rembang",
			"salatiga",
			"purbalingga",
			"banjarnegara",
			"tegal",
			"brebes",
			"banyumas",

			// DI Yogyakarta
			"yogyakarta",
			"bantul",
			"sleman",
			"kulon progo",
			"gunung kidul",

			// DKI Jakarta
			"jakarta barat",
			"jakarta selatan",
			"jakarta pusat",
			"jakarta utara",
			"jakarta timur",
			"kepulauan seribu",

			// Jawa Timur
			"jember",
			"banyuwangi",
			"bondowoso",
			"kediri",
			"madiun",
			"magetan",
			"ngawi",
			"pacitan",
			"ponorogo",
			"mojokerto",
			"jombang",
			"nganjuk",
			"malang",
			"blitar",
			"batu",
			"probolinggo",
			"lumajang",
			"situbondo",
			"pasuruan",
			"bojonegoro",
			"surabaya",
			"gresik",
			"lamongan",
			"bangkalan",
			"pamekasan",
			"sampang",
			"sidoarjo",
			"sumenep",
			"tuban",
			"tulungagung",
			"trenggalek"},
		"status_bid_data":      {"Mengumpulkan", "Siap Antar"},
		"status_bid_scheduler": {"Wait", "Ambil", "Kirim"},
	}

	for enumName, values := range enumMap {
		var exists bool
		checkSQL := "SELECT EXISTS(SELECT 1 FROM pg_type WHERE typname = ?);"
		if err := tx.Raw(checkSQL, enumName).Scan(&exists).Error; err != nil {
			tx.Rollback()
			return err
		}

		if !exists {
			createSQL := fmt.Sprintf("CREATE TYPE %s AS ENUM (%s);", enumName, joinWithQuotes(values))
			if err := tx.Exec(createSQL).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

func UpBarangEnums(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	enumMap := map[string][]string{
		"status_varian": {"Ready", "Dipesan", "Diproses", "Terjual", "Down", "Pending"},
	}

	for enumName, values := range enumMap {
		// Cek apakah enum sudah ada
		var exists bool
		checkSQL := "SELECT EXISTS(SELECT 1 FROM pg_type WHERE typname = ?);"
		if err := tx.Raw(checkSQL, enumName).Scan(&exists).Error; err != nil {
			tx.Rollback()
			return err
		}

		if !exists {
			// Create type baru
			createSQL := fmt.Sprintf("CREATE TYPE %s AS ENUM (%s);", enumName, joinWithQuotes(values))
			if err := tx.Exec(createSQL).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

func UpEnumsTransaksi(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	enumMap := map[string][]string{
		"status_transaksi": {"Dibayar", "Diproses", "Waiting", "Dikirim", "Selesai", "Dibatalkan"},
		// Dibayar adalah status default sebuah transaksi sampai seller melakukan approval.
		// Setelah transaksi di-approve oleh seller, status akan berubah menjadi "Diproses".
		// Status akan berubah lagi menjadi "Waiting" setelah seller memutuskan untuk mengirim barang.
		// Status menjadi "Dikirim" ketika seller sudah menyerahkan barang ke kurir.
		// Status menjadi "Selesai" ketika pengguna telah menerima barang dan mengonfirmasi bahwa transaksi telah selesai.
		// "Dibatalkan" digunakan ketika pengguna atau seller membatalkan transaksi, baik karena kesepakatan maupun sepihak.
		// Pembatalan hanya bisa dilakukan selama status masih "Dibayar".

		"status_pengiriman": {"Waiting", "Picked Up", "Diperjalanan", "Sampai", "Trouble"},
		// "Packaging" muncul di tabel pengiriman dan terjadi ketika status transaksi adalah "Diproses".
		// Status berubah menjadi "Picked Up" ketika status transaksi berubah menjadi "Dikirim".
		// Status menjadi "Diperjalanan" ketika kurir memperbarui status sendiri selama proses pengantaran.
		// Dalam status ini, kurir juga mengupdate catatan serta koordinat (latitude, longitude)
		// di tabel "jejak_pengiriman" yang merupakan child dari tabel "pengiriman".
		// Status "Sampai" akan dikonfirmasi oleh kurir saat barang tiba di tujuan.
		// Segala hal tak terduga seperti barang tidak sesuai, masalah di jalan, atau kerusakan barang
		// akan masuk ke status "Trouble". Namun untuk saat ini, kita berasumsi semua berjalan lancar.

		"status_pengiriman_ekspedisi": {"Waiting", "Dikirim", "Sampai Agent", "Masuk Gateaway", "Sampai Agent Tujuan", "Dikirim Agent", "Sampai"},

		"status_paid_failed": {"Ditinjau", "Pending", "Batal", "Lanjut"},
		// "Ditinjau" berarti sistem sedang melakukan pemeriksaan terhadap transaksi gagal.
		// Kegagalan umumnya disebabkan oleh kesalahan foreign key, sehingga sistem akan melakukan self-healing data.
		// Setelah proses perbaikan (self-healing) selesai, status akan otomatis berubah menjadi "Pending".
		// Pada tahap "Pending", pengguna dapat memilih untuk melanjutkan atau membatalkan.
		// Secara default, jika tidak ada tindakan, status akan otomatis berubah menjadi "Batal" setelah 5 jam.
		// Jika pengguna memilih untuk melanjutkan, status berubah menjadi "Lanjut".
		// Dalam status "Lanjut", data akan dialihkan ke tabel transaksi dan pembayaran,
		// kemudian proses akan berlanjut seperti transaksi normal pada umumnya.
	}

	for enumName, values := range enumMap {
		var exists bool
		checkSQL := "SELECT EXISTS(SELECT 1 FROM pg_type WHERE typname = ?);"
		if err := tx.Raw(checkSQL, enumName).Scan(&exists).Error; err != nil {
			tx.Rollback()
			return err
		}

		if !exists {
			createSQL := fmt.Sprintf("CREATE TYPE %s AS ENUM (%s);", enumName, joinWithQuotes(values))
			if err := tx.Exec(createSQL).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

func joinWithQuotes(values []string) string {
	res := ""
	for i, v := range values {
		res += fmt.Sprintf("'%s'", v)
		if i != len(values)-1 {
			res += ","
		}
	}
	return res
}
