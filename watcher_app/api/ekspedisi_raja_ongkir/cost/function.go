package ekspedisi_cost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const RajaOngkirDomesticCostEndpoint = "https://rajaongkir.komerce.id/api/v1/calculate/domestic-cost"

func (s *StarterDomesticCostReq) DomesticCostReq(ctx context.Context) CalculateDomesticCostResponse {
	form := url.Values{}
	form.Set("origin", s.Origin)
	form.Set("destination", s.Destination)
	form.Set("weight", s.Weight)
	form.Set("courier", s.Courier)
	form.Set("price", s.Price)

	fmt.Println("origin:", s.Origin, "destination:", s.Destination, "weight:", s.Weight, "Courier:", s.Courier, "price:", s.Price)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, RajaOngkirDomesticCostEndpoint, strings.NewReader(form.Encode()))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("key", "peogMXfye4692639b41ad550nv6xANvu")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var result CalculateDomesticCostResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s\nMessage: %s\n\n", resp.Status, result.Meta.Message)
	for _, service := range result.Data {
		fmt.Printf("Courier: %s | Service: %s | Cost: %d | ETD: %s | Description: %s\n",
			service.Code, service.Service, service.Cost, service.ETD, service.Description)
	}

	return result
}
