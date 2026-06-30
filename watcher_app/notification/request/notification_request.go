package notification_request

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	notification_contract "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/contract"
	notification_error "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/error"
	notification_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/models"
)

func PostToNotification[T notification_models.NotificationPengguna | notification_models.NotificationSeller | notification_models.NotificationKurir](ctx context.Context, data T, host, port, path string) error {
	marshallData, err := notification_contract.NotificationManager[T]{}.MarshallNotification(data)
	if err != nil {
		return notification_error.ErrorDataTidakCocok
	}

	bodyReader := bytes.NewBuffer(marshallData)

	req := &http.Request{
		Method: http.MethodPost,
		URL: &url.URL{
			Scheme: "http",                           // Jangan lupa skemanya (http/https) biar gak error pas dikirim
			Host:   fmt.Sprintf("%s:%s", host, port), // Tambahkan port service notification-mu jika ada
			Path:   path,
		},
		Header: make(http.Header),
		Body:   io.NopCloser(bodyReader), // Membungkus bytes buffer jadi io.ReadCloser
	}

	// 2. Bagus urusannya kalau ditambahin Header Content-Type JSON
	req.Header.Set("Content-Type", "application/json")

	// 3. Selalu biasakan pakai Context bawaan minimal context.Background()
	// supaya request-nya bisa di-trace atau di-cancel jika timeout
	req = req.WithContext(ctx)

	// 4. Kirim request menggunakan http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return notification_error.ErrorGagalKirim // Kembalikan error jika koneksi/server tujuan gagal dihubungi
	}
	defer resp.Body.Close() // Wajib di-close biar gak memory leak!

	// 5. Cek status code dari server notification
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return notification_error.ErrorGagalKirim // Sesuaikan dengan struct error-mu
	}

	return nil
}
