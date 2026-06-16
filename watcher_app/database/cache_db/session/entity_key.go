package session_cache_db

import (
	"fmt"
	"strconv"

	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"

)

func GetSessionKey[T *sot_models.Pengguna | *sot_models.Seller | *sot_models.Kurir](entity T) (error, string) {
	switch v := any(entity).(type) {
	case sot_models.Pengguna:
		return nil, fmt.Sprintf(`session_user_%s_%s_%s`, strconv.Itoa(int(v.ID)), v.Username, v.Email)
	case sot_models.Seller:
		return nil, fmt.Sprintf(`session_seller_%s_%s_%s`, strconv.Itoa(int(v.ID)), v.Username, v.Email)
	case sot_models.Kurir:
		return nil, fmt.Sprintf(`session_kurir_%s_%s_%s`, strconv.Itoa(int(v.ID)), v.Username, v.Email)
	default:
		return fmt.Errorf("Tak mendapatkan session"), ""
	}
}
