package cache_db_function

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	session_cache_db "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cache_db/session"
	sot_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/sot_database/models"
	"github.com/anan112pcmec/Burung-backend-2/watcher_app/helper"
)

func UpdateSessionData[T sot_models.Pengguna | sot_models.Seller | sot_models.Kurir](ctx context.Context, rds_session redis.Client, key_session string, data T) error {
	fctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	res, err := rds_session.HGetAll(fctx, key_session).Result()
	if err != nil {
		return err
	}

	var fields map[string]any
	var keyz string

	switch v := any(data).(type) {
	case sot_models.Pengguna:
		fields = helper.StructToJSONMap(v)
		if err, key := session_cache_db.GetSessionKey(&v); err != nil {
			return err
		} else {
			keyz = key
		}
	case sot_models.Seller:
		fields = helper.StructToJSONMap(v)
		if err, key := session_cache_db.GetSessionKey(&v); err != nil {
			return err
		} else {
			keyz = key
		}
	case sot_models.Kurir:
		fields = helper.StructToJSONMap(v)
		if err, key := session_cache_db.GetSessionKey(&v); err != nil {
			return err
		} else {
			keyz = key
		}
	}

	if keyz == key_session {
		for k, newVal := range fields {
			if oldVal, exists := res[k]; exists && oldVal == fmt.Sprintf(`%v`, newVal) {
				delete(fields, k)
			}
		}
	}

	if len(fields) == 0 {
		return nil
	}

	if err := rds_session.HSet(fctx, keyz, fields).Err(); err != nil {
		return err
	}
	return nil
}
