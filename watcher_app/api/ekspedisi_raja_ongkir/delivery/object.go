package ekspedisi_delivery

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
}

type Summary struct {
	CourierCode   string `json:"courier_code"`
	CourierName   string `json:"courier_name"`
	WaybillNumber string `json:"waybill_number"`
	ServiceCode   string `json:"service_code"`
	WaybillDate   string `json:"waybill_date"`
	ShipperName   string `json:"shipper_name"`
	ReceiverName  string `json:"receiver_name"`
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	Status        string `json:"status"`
}

type Details struct {
	WaybillNumber    string `json:"waybill_number"`
	WaybillDate      string `json:"waybill_date"`
	WaybillTime      string `json:"waybill_time"`
	Weight           string `json:"weight"` // Diubah ke string karena template asli menggunakan "{{ package.weight }}"
	Origin           string `json:"origin"`
	Destination      string `json:"destination"`
	ShipperName      string `json:"shipper_name"`
	ShipperAddress1  string `json:"shipper_address1"`
	ShipperAddress2  string `json:"shipper_address2"`
	ShipperAddress3  string `json:"shipper_address3"`
	ShipperCity      string `json:"shipper_city"`
	ReceiverName     string `json:"receiver_name"`
	ReceiverAddress1 string `json:"receiver_address1"`
	ReceiverAddress2 string `json:"receiver_address2"`
	ReceiverAddress3 string `json:"receiver_address3"`
	ReceiverCity     string `json:"receiver_city"`
}

type DeliveryStatus struct {
	Status      string `json:"status"`
	PodReceiver string `json:"pod_receiver"`
	PodDate     string `json:"pod_date"`
	PodTime     string `json:"pod_time"`
}

type Manifest struct {
	ManifestCode        string `json:"manifest_code"`
	ManifestDescription string `json:"manifest_description"`
	ManifestDate        string `json:"manifest_date"`
	ManifestTime        string `json:"manifest_time"`
	CityName            string `json:"city_name"`
}

type Data struct {
	Delivered      string         `json:"delivered"`
	Summary        Summary        `json:"summary"`
	Details        Details        `json:"details"`
	DeliveryStatus DeliveryStatus `json:"delivery_status"`
	Manifest       []Manifest     `json:"manifest"`
}
type ResponseTrackingEkspedisi struct {
	Meta `json:"meta"`
	Data `json:"data"`
}
