package ekspedisi_cost

type CalculateDomesticCostResponse struct {
	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Status  string `json:"status"`
	} `json:"meta"`
	Data []struct {
		Name        string `json:"name"`
		Code        string `json:"code"`
		Service     string `json:"service"`
		Description string `json:"description"`
		Cost        int    `json:"cost"`
		ETD         string `json:"etd"`
	} `json:"data"`
}

type StarterDomesticCostReq struct {
	Origin      string
	Destination string
	Weight      string
	Courier     string
	Price       string
}
