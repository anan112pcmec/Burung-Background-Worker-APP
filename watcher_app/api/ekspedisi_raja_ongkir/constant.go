package ekspedisi

import (
	"math/rand"
	"time"
)

const (
	JNE     string = "jne"
	POS     string = "pos"
	TIKI    string = "tiki"
	JNT     string = "jnt"
	SICEPAT string = "sicepat"
)

var NamaProv = map[string]string{
	"banten":        "banten",
	"jawa_barat":    "jawa_barat",
	"jawa_tengah":   "jawa_tengah",
	"di_yogyakarta": "di_yogyakarta",
	"dki_jakarta":   "dki_jakarta",
	"jawa_timur":    "jawa_timur",
}

var JawaCitiesJNE = map[string]bool{
	// Banten
	"cilegon":           true,
	"pandeglang":        true,
	"lebak":             true,
	"serang":            true,
	"tangerang":         true,
	"tangerang selatan": true,

	// Jawa Barat
	"bandung":       true,
	"cimahi":        true,
	"sumedang":      true,
	"garut":         true,
	"bandung barat": true,
	"cianjur":       true,
	"bekasi":        true,
	"bogor":         true,
	"cirebon":       true,
	"indramayu":     true,
	"kuningan":      true,
	"majalengka":    true,
	"depok":         true,
	"karawang":      true,
	"purwakarta":    true,
	"subang":        true,
	"sukabumi":      true,
	"tasikmalaya":   true,
	"banjar":        true,
	"ciamis":        true,
	"pangandaran":   true,

	// Jawa Tengah
	"cilacap":      true,
	"magelang":     true,
	"kebumen":      true,
	"wonosobo":     true,
	"purworejo":    true,
	"temanggung":   true,
	"surakarta":    true,
	"boyolali":     true,
	"karanganyar":  true,
	"klaten":       true,
	"sragen":       true,
	"sukoharjo":    true,
	"wonogiri":     true,
	"semarang":     true,
	"jepara":       true,
	"kudus":        true,
	"pekalongan":   true,
	"batang":       true,
	"blora":        true,
	"demak":        true,
	"kendal":       true,
	"pati":         true,
	"pemalang":     true,
	"grobogan":     true,
	"rembang":      true,
	"salatiga":     true,
	"purbalingga":  true,
	"banjarnegara": true,
	"tegal":        true,
	"brebes":       true,
	"banyumas":     true,

	// DI Yogyakarta
	"yogyakarta":   true,
	"bantul":       true,
	"sleman":       true,
	"kulon progo":  true,
	"gunung kidul": true,

	// DKI Jakarta
	"jakarta barat":    true,
	"jakarta selatan":  true,
	"jakarta pusat":    true,
	"jakarta utara":    true,
	"jakarta timur":    true,
	"kepulauan seribu": true,

	// Jawa Timur
	"jember":      true,
	"banyuwangi":  true,
	"bondowoso":   true,
	"kediri":      true,
	"madiun":      true,
	"magetan":     true,
	"ngawi":       true,
	"pacitan":     true,
	"ponorogo":    true,
	"mojokerto":   true,
	"jombang":     true,
	"nganjuk":     true,
	"malang":      true,
	"blitar":      true,
	"batu":        true,
	"probolinggo": true,
	"lumajang":    true,
	"situbondo":   true,
	"pasuruan":    true,
	"bojonegoro":  true,
	"surabaya":    true,
	"gresik":      true,
	"lamongan":    true,
	"bangkalan":   true,
	"pamekasan":   true,
	"sampang":     true,
	"sidoarjo":    true,
	"sumenep":     true,
	"tuban":       true,
	"tulungagung": true,
	"trenggalek":  true,
}

var JawaCities = map[string]string{
	// Banten
	"cilegon":           "143",
	"pandeglang":        "144",
	"lebak":             "147",
	"serang":            "148",
	"tangerang":         "592",
	"tangerang selatan": "594",

	// Jawa Barat
	"bandung":       "55",
	"cimahi":        "56",
	"sumedang":      "57",
	"garut":         "59",
	"bandung barat": "60",
	"cianjur":       "62",
	"bekasi":        "63",
	"bogor":         "77",
	"cirebon":       "129",
	"indramayu":     "131",
	"kuningan":      "132",
	"majalengka":    "133",
	"depok":         "199",
	"karawang":      "329",
	"purwakarta":    "532",
	"subang":        "533",
	"sukabumi":      "538",
	"tasikmalaya":   "632",
	"banjar":        "633",
	"ciamis":        "634",
	"pangandaran":   "635",

	// Jawa Tengah
	"cilacap":      "149",
	"magelang":     "383",
	"kebumen":      "384",
	"wonosobo":     "385",
	"purworejo":    "386",
	"temanggung":   "387",
	"surakarta":    "539",
	"boyolali":     "540",
	"karanganyar":  "541",
	"klaten":       "542",
	"sragen":       "543",
	"sukoharjo":    "544",
	"wonogiri":     "545",
	"semarang":     "560",
	"jepara":       "561",
	"kudus":        "562",
	"pekalongan":   "563",
	"batang":       "564",
	"blora":        "565",
	"demak":        "567",
	"kendal":       "568",
	"pati":         "569",
	"pemalang":     "570",
	"grobogan":     "571",
	"rembang":      "572",
	"salatiga":     "573",
	"purbalingga":  "575",
	"banjarnegara": "576",
	"tegal":        "588",
	"brebes":       "589",
	"banyumas":     "591",

	// DI Yogyakarta
	"yogyakarta":   "259",
	"bantul":       "260",
	"sleman":       "261",
	"kulon progo":  "262",
	"gunung kidul": "263",

	// DKI Jakarta
	"jakarta barat":    "135",
	"jakarta selatan":  "136",
	"jakarta pusat":    "137",
	"jakarta utara":    "138",
	"jakarta timur":    "139",
	"kepulauan seribu": "141",

	// Jawa Timur
	"jember":      "256",
	"banyuwangi":  "257",
	"bondowoso":   "258",
	"kediri":      "289",
	"madiun":      "353",
	"magetan":     "355",
	"ngawi":       "357",
	"pacitan":     "359",
	"ponorogo":    "360",
	"mojokerto":   "388",
	"jombang":     "389",
	"nganjuk":     "390",
	"malang":      "391",
	"blitar":      "392",
	"batu":        "393",
	"probolinggo": "394",
	"lumajang":    "395",
	"situbondo":   "396",
	"pasuruan":    "531",
	"bojonegoro":  "566",
	"surabaya":    "577",
	"gresik":      "578",
	"lamongan":    "579",
	"bangkalan":   "580",
	"pamekasan":   "581",
	"sampang":     "582",
	"sidoarjo":    "583",
	"sumenep":     "584",
	"tuban":       "585",
	"tulungagung": "586",
	"trenggalek":  "587",
}

func RandomCityByProvince(province string) string {
	var cities []string

	switch province {
	case "banten":
		cities = []string{"cilegon", "pandeglang", "lebak", "serang", "tangerang", "tangerang selatan"}
	case "jawa_barat":
		cities = []string{"bandung", "cimahi", "sumedang", "garut", "bandung barat", "cianjur", "bekasi", "bogor", "cirebon", "indramayu", "kuningan", "majalengka", "depok", "karawang", "purwakarta", "subang", "sukabumi", "tasikmalaya", "banjar", "ciamis", "pangandaran"}
	case "jawa_tengah":
		cities = []string{"cilacap", "magelang", "kebumen", "wonosobo", "purworejo", "temanggung", "surakarta", "boyolali", "karanganyar", "klaten", "sragen", "sukoharjo", "wonogiri", "semarang", "jepara", "kudus", "pekalongan", "batang", "blora", "demak", "kendal", "pati", "pemalang", "grobogan", "rembang", "salatiga", "purbalingga", "banjarnegara", "tegal", "brebes", "banyumas"}
	case "di_yogyakarta":
		cities = []string{"yogyakarta", "bantul", "sleman", "kulon progo", "gunung kidul"}
	case "dki_jakarta":
		cities = []string{"jakarta barat", "jakarta selatan", "jakarta pusat", "jakarta utara", "jakarta timur", "kepulauan seribu"}
	case "jawa_timur":
		cities = []string{"jember", "banyuwangi", "bondowoso", "kediri", "madiun", "magetan", "ngawi", "pacitan", "ponorogo", "mojokerto", "jombang", "nganjuk", "malang", "blitar", "batu", "probolinggo", "lumajang", "situbondo", "pasuruan", "bojonegoro", "surabaya", "gresik", "lamongan", "bangkalan", "pamekasan", "sampang", "sidoarjo", "sumenep", "tuban", "tulungagung", "trenggalek"}
	default:
		return ""
	}

	rand.Seed(time.Now().UnixNano())
	return cities[rand.Intn(len(cities))]
}
