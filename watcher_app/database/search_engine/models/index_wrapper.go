package se_models

import "github.com/meilisearch/meilisearch-go"

type IndexWrapper struct {
	BarangIndukIndex, SellerIndex, TransaksiIndex, AlamatPenggunaIndex, PenggunaIndex, AlamatEkspedisiIndex, AlamatKurir, AlamatGudang, KurirIndex, FollowerSeller meilisearch.IndexManager
}
