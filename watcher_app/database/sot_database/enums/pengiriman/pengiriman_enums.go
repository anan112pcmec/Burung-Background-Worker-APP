package pengiriman_enums

// Enums jenis pengiriman
var (
	Reguler = "Reguler"
	Fast    = "Fast"
	Instant = "Instant"
)

// Enums untuk pengiriman non ekspedisi

const Waiting string = "Waiting"
const PickedUp string = "Picked Up"
const Diperjalanan string = "Diperjalanan"
const Sampai string = "Sampai"
const Trouble = "Trouble"

// Enums untuk pengiriman ekspedisi

const WaitingEkspedisi string = "Waiting"
const DikirimEkspedisi string = "Dikirim"
const SampaiAgentEkspedisi string = "Sampai Agent"
const SampaiAgentTujuanEkspedisi string = "Sampai Agent Tujuan"
const DikirimAgentEkspedisi string = "Dikirim Agent"
const SampaiEkspedisi string = "Sampai"
