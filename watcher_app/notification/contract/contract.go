package notification_contract

import (
	"encoding/json"
	"fmt"

	notification_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/notification/models"
)

type ContractNotification[T notification_models.NotificationPengguna | notification_models.NotificationSeller | notification_models.NotificationKurir] interface {
	MarshallNotification(data T) ([]byte, error)
	ParseNotification(data interface{}) (error, T)
}

type NotificationManager[T notification_models.NotificationPengguna | notification_models.NotificationSeller | notification_models.NotificationKurir] struct{}

func (n NotificationManager[T]) MarshallNotification(data T) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshall notification: %w", err)
	}
	return bytes, nil
}

// 2. Implementasi ParseNotification
// Mengubah data (bisa berupa []byte atau string) kembali menjadi struct T
func (n NotificationManager[T]) ParseNotification(data interface{}) (error, T) {
	var result T

	var byteData []byte
	switch v := data.(type) {
	case []byte:
		byteData = v
	case string:
		byteData = []byte(v)
	default:
		return fmt.Errorf("unsupported data type for parsing"), result
	}

	err := json.Unmarshal(byteData, &result)
	if err != nil {
		return fmt.Errorf("failed to parse notification: %w", err), result
	}

	return nil, result
}
