package se_initialize

import (
	"context"

	"github.com/meilisearch/meilisearch-go"

	se_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/search_engine/models"

)

func InitIndex(ctx context.Context, c meilisearch.ServiceManager) se_models.IndexWrapper {

	return se_models.IndexWrapper{
		BarangIndukIndex:     c.Index(se_models.BarangInduk{}.IndexName()),
		SellerIndex:          c.Index(se_models.Seller{}.IndexName()),
		TransaksiIndex:       c.Index(se_models.Transaksi{}.IndexName()),
		PenggunaIndex:        c.Index(se_models.Pengguna{}.IndexName()),
		AlamatEkspedisiIndex: c.Index(se_models.AlamatEkspedisi{}.IndexName()),
	}
}
