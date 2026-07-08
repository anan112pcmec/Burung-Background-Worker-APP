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
	// 1. Marshalling data
	marshallData, err := notification_contract.NotificationManager[T]{}.MarshallNotification(data)
	if err != nil {
		// TRACE: Log error saat marshalling gagal
		fmt.Printf("[TRACE][ERROR] Gagal marshalling data: %v\n", err)
		return notification_error.ErrorDataTidakCocok
	}

	bodyReader := bytes.NewBuffer(marshallData)

	// Buat full URL untuk bahan tracing & request
	fullURL := fmt.Sprintf("http://%s:%s%s", host, port, path)

	// TRACE: Log sebelum request dikirim (Sangat berguna untuk debugging payload)
	fmt.Printf("[TRACE][START] Mengirim notifikasi ke %s | Payload Size: %d bytes\n", fullURL, len(marshallData))

	req := &http.Request{
		Method: http.MethodPost,
		URL: &url.URL{
			Scheme: "http",
			Host:   fmt.Sprintf("%s:%s", host, port),
			Path:   path,
		},
		Header: make(http.Header),
		Body:   io.NopCloser(bodyReader),
	}

	req.Header.Set("Content-Type", "application/json")

	// Pastikan ctx tidak nil sebelum dipasang, jika nil gunakan context.Background()
	if ctx == nil {
		ctx = context.Background()
	}
	req = req.WithContext(ctx)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// TRACE: Log error koneksi / HTTP client
		fmt.Printf("[TRACE][ERROR] Gagal mengirim HTTP request ke %s. Error: %v\n", fullURL, err)
		return notification_error.ErrorGagalKirim
	}
	defer resp.Body.Close()

	// TRACE: Log status code yang diterima dari server
	fmt.Printf("[TRACE][RESPONSE] Menerima respon dari %s | Status: %s (%d)\n", fullURL, resp.Status, resp.StatusCode)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		// TRACE: Log jika status code bukan 200/201
		fmt.Printf("[TRACE][WARN] Server merespon dengan status tidak sukses: %d\n", resp.StatusCode)
		return notification_error.ErrorGagalKirim
	}

	// TRACE: Log sukses akhir
	fmt.Printf("[TRACE][SUCCESS] Notifikasi berhasil dikirim ke %s\n", fullURL)
	return nil
}
