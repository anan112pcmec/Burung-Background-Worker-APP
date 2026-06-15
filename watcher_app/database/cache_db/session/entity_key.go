package session_cache_db

import (
	"fmt"
	"strconv"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
)

func GetSessionKey[T *models.Pengguna | *models.Seller | *models.Kurir](entity T) (error, string) {
	switch v := any(entity).(type) {
	case models.Pengguna:
		return nil, fmt.Sprintf(`session_user_%s_%s_%s`, strconv.Itoa(int(v.ID)), v.Username, v.Email)
	case models.Seller:
		return nil, fmt.Sprintf(`session_seller_%s_%s_%s`, strconv.Itoa(int(v.ID)), v.Username, v.Email)
	case models.Kurir:
		return nil, fmt.Sprintf(`session_kurir_%s_%s_%s`, strconv.Itoa(int(v.ID)), v.Username, v.Email)
	default:
		return fmt.Errorf("Tak mendapatkan session"), ""
	}
}
