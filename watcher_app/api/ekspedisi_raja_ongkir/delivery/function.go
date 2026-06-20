package ekspedisi_delivery

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// lpn adalah The last 5 digits of the recipient's phone number are required for additional validation by certain couriers, including JNE

func LacakEkspedisi(apikey, awb, courier string, lpn int) (error, ResponseTrackingEkspedisi) {
	url := fmt.Sprintf(`https://rajaongkir.komerce.id/api/v1/track/waybill?awb=%s&courier=%s`, awb, courier)
	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return err, ResponseTrackingEkspedisi{}
	}
	req.Header.Add("key", apikey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err, ResponseTrackingEkspedisi{}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err, ResponseTrackingEkspedisi{}
	}

	var hasil ResponseTrackingEkspedisi
	if err := json.Unmarshal(body, &hasil); err != nil {
		return err, ResponseTrackingEkspedisi{}
	}

	fmt.Println(string(body))
	return nil, hasil
}
