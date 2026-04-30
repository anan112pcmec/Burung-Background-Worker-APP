package nama_kota

const (
	// Banten
	Cilegon          string = "cilegon"
	Pandeglang       string = "pandeglang"
	Lebak            string = "lebak"
	Serang           string = "serang"
	Tangerang        string = "tangerang"
	TangerangSelatan string = "tangerang selatan"

	// Jawa Barat
	Bandung      string = "bandung"
	Cimahi       string = "cimahi"
	Sumedang     string = "sumedang"
	Garut        string = "garut"
	BandungBarat string = "bandung barat"
	Cianjur      string = "cianjur"
	Bekasi       string = "bekasi"
	Bogor        string = "bogor"
	Cirebon      string = "cirebon"
	Indramayu    string = "indramayu"
	Kuningan     string = "kuningan"
	Majalengka   string = "majalengka"
	Depok        string = "depok"
	Karawang     string = "karawang"
	Purwakarta   string = "purwakarta"
	Subang       string = "subang"
	Sukabumi     string = "sukabumi"
	Tasikmalaya  string = "tasikmalaya"
	Banjar       string = "banjar"
	Ciamis       string = "ciamis"
	Pangandaran  string = "pangandaran"

	// Jawa Tengah
	Cilacap      string = "cilacap"
	Magelang     string = "magelang"
	Kebumen      string = "kebumen"
	Wonosobo     string = "wonosobo"
	Purworejo    string = "purworejo"
	Temanggung   string = "temanggung"
	Surakarta    string = "surakarta"
	Boyolali     string = "boyolali"
	Karanganyar  string = "karanganyar"
	Klaten       string = "klaten"
	Sragen       string = "sragen"
	Sukoharjo    string = "sukoharjo"
	Wonogiri     string = "wonogiri"
	Semarang     string = "semarang"
	Jepara       string = "jepara"
	Kudus        string = "kudus"
	Pekalongan   string = "pekalongan"
	Batang       string = "batang"
	Blora        string = "blora"
	Demak        string = "demak"
	Kendal       string = "kendal"
	Pati         string = "pati"
	Pemalang     string = "pemalang"
	Grobogan     string = "grobogan"
	Rembang      string = "rembang"
	Salatiga     string = "salatiga"
	Purbalingga  string = "purbalingga"
	Banjarnegara string = "banjarnegara"
	Tegal        string = "tegal"
	Brebes       string = "brebes"
	Banyumas     string = "banyumas"

	// DI Yogyakarta
	Yogyakarta  string = "yogyakarta"
	Bantul      string = "bantul"
	Sleman      string = "sleman"
	KulonProgo  string = "kulon progo"
	GunungKidul string = "gunung kidul"

	// DKI Jakarta
	JakartaBarat    string = "jakarta barat"
	JakartaSelatan  string = "jakarta selatan"
	JakartaPusat    string = "jakarta pusat"
	JakartaUtara    string = "jakarta utara"
	JakartaTimur    string = "jakarta timur"
	KepulauanSeribu string = "kepulauan seribu"

	// Jawa Timur
	Jember      string = "jember"
	Banyuwangi  string = "banyuwangi"
	Bondowoso   string = "bondowoso"
	Kediri      string = "kediri"
	Madiun      string = "madiun"
	Magetan     string = "magetan"
	Ngawi       string = "ngawi"
	Pacitan     string = "pacitan"
	Ponorogo    string = "ponorogo"
	Mojokerto   string = "mojokerto"
	Jombang     string = "jombang"
	Nganjuk     string = "nganjuk"
	Malang      string = "malang"
	Blitar      string = "blitar"
	Batu        string = "batu"
	Probolinggo string = "probolinggo"
	Lumajang    string = "lumajang"
	Situbondo   string = "situbondo"
	Pasuruan    string = "pasuruan"
	Bojonegoro  string = "bojonegoro"
	Surabaya    string = "surabaya"
	Gresik      string = "gresik"
	Lamongan    string = "lamongan"
	Bangkalan   string = "bangkalan"
	Pamekasan   string = "pamekasan"
	Sampang     string = "sampang"
	Sidoarjo    string = "sidoarjo"
	Sumenep     string = "sumenep"
	Tuban       string = "tuban"
	Tulungagung string = "tulungagung"
	Trenggalek  string = "trenggalek"
)

var KotaJawaMap = map[string]string{
	// Banten
	Cilegon:          Cilegon,
	Pandeglang:       Pandeglang,
	Lebak:            Lebak,
	Serang:           Serang,
	Tangerang:        Tangerang,
	TangerangSelatan: TangerangSelatan,

	// Jawa Barat
	Bandung:      Bandung,
	Cimahi:       Cimahi,
	Sumedang:     Sumedang,
	Garut:        Garut,
	BandungBarat: BandungBarat,
	Cianjur:      Cianjur,
	Bekasi:       Bekasi,
	Bogor:        Bogor,
	Cirebon:      Cirebon,
	Indramayu:    Indramayu,
	Kuningan:     Kuningan,
	Majalengka:   Majalengka,
	Depok:        Depok,
	Karawang:     Karawang,
	Purwakarta:   Purwakarta,
	Subang:       Subang,
	Sukabumi:     Sukabumi,
	Tasikmalaya:  Tasikmalaya,
	Banjar:       Banjar,
	Ciamis:       Ciamis,
	Pangandaran:  Pangandaran,

	// Jawa Tengah
	Cilacap:      Cilacap,
	Magelang:     Magelang,
	Kebumen:      Kebumen,
	Wonosobo:     Wonosobo,
	Purworejo:    Purworejo,
	Temanggung:   Temanggung,
	Surakarta:    Surakarta,
	Boyolali:     Boyolali,
	Karanganyar:  Karanganyar,
	Klaten:       Klaten,
	Sragen:       Sragen,
	Sukoharjo:    Sukoharjo,
	Wonogiri:     Wonogiri,
	Semarang:     Semarang,
	Jepara:       Jepara,
	Kudus:        Kudus,
	Pekalongan:   Pekalongan,
	Batang:       Batang,
	Blora:        Blora,
	Demak:        Demak,
	Kendal:       Kendal,
	Pati:         Pati,
	Pemalang:     Pemalang,
	Grobogan:     Grobogan,
	Rembang:      Rembang,
	Salatiga:     Salatiga,
	Purbalingga:  Purbalingga,
	Banjarnegara: Banjarnegara,
	Tegal:        Tegal,
	Brebes:       Brebes,
	Banyumas:     Banyumas,

	// DI Yogyakarta
	Yogyakarta:  Yogyakarta,
	Bantul:      Bantul,
	Sleman:      Sleman,
	KulonProgo:  KulonProgo,
	GunungKidul: GunungKidul,

	// DKI Jakarta
	JakartaBarat:    JakartaBarat,
	JakartaSelatan:  JakartaSelatan,
	JakartaPusat:    JakartaPusat,
	JakartaUtara:    JakartaUtara,
	JakartaTimur:    JakartaTimur,
	KepulauanSeribu: KepulauanSeribu,

	// Jawa Timur
	Jember:      Jember,
	Banyuwangi:  Banyuwangi,
	Bondowoso:   Bondowoso,
	Kediri:      Kediri,
	Madiun:      Madiun,
	Magetan:     Magetan,
	Ngawi:       Ngawi,
	Pacitan:     Pacitan,
	Ponorogo:    Ponorogo,
	Mojokerto:   Mojokerto,
	Jombang:     Jombang,
	Nganjuk:     Nganjuk,
	Malang:      Malang,
	Blitar:      Blitar,
	Batu:        Batu,
	Probolinggo: Probolinggo,
	Lumajang:    Lumajang,
	Situbondo:   Situbondo,
	Pasuruan:    Pasuruan,
	Bojonegoro:  Bojonegoro,
	Surabaya:    Surabaya,
	Gresik:      Gresik,
	Lamongan:    Lamongan,
	Bangkalan:   Bangkalan,
	Pamekasan:   Pamekasan,
	Sampang:     Sampang,
	Sidoarjo:    Sidoarjo,
	Sumenep:     Sumenep,
	Tuban:       Tuban,
	Tulungagung: Tulungagung,
	Trenggalek:  Trenggalek,
}
