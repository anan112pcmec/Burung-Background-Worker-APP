package cass_models

import "strings"

// TableNameSotReplica implementations delegate to TableNameHistorical
// and replace the trailing "_historical" with "_sot_replica".

func (r BarangInduk) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r KategoriBarang) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r VarianBarang) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}

func (r Pengiriman) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r JejakPengiriman) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r PengirimanEkspedisi) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r JejakPengirimanEkspedisi) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}

func (r EntitySocialMedia) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r Komentar) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r KomentarChild) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r Keranjang) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r BarangDisukai) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r BarangWishlist) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r AlamatPengguna) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r Wishlist) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r Review) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r ReviewLike) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r ReviewDislike) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}

func (r MediaPenggunaProfilFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaSellerProfilFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaSellerBannerFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaSellerTokoFisikFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaKurirProfilFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaEtalaseFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaBarangIndukFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaBarangIndukVideo) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaKategoriBarangFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}

func (r MediaDistributorDataDokumen) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaDistributorDataNPWPFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaDistributorDataNIBFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaDistributorDataSuratKerjasamaDokumen) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}

func (r MediaBrandDataPerwakilanDokumen) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaBrandDataSertifikatFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaBrandDataNIBFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaBrandDataNPWPFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaBrandDataLogoFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaBrandDataSuratKerjasamaDokumen) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}

func (r MediaInformasiKendaraanKurirKendaraanFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaInformasiKendaraanKurirBPKBFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaInformasiKendaraanKurirSTNKFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaInformasiKurirKTPFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}

func (r MediaReviewFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaReviewVideo) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaTransaksiApprovedFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaTransaksiApprovedVideo) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}

func (r MediaPengirimanPickedUpFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaPengirimanSampaiFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaPengirimanEkspedisiPickedUpFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r MediaPengirimanEkspedisiSampaiAgentFoto) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}

func (r PayOutKurir) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r PayOutSeller) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}

func (r RekeningSeller) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r RekeningKurir) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}

func (r Pembayaran) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r Transaksi) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
func (r TransaksiFailed) TableNameSotReplica() string {
	return strings.Replace(r.TableNameHistorical(), "_historical", "_sot_replica", 1)
}
