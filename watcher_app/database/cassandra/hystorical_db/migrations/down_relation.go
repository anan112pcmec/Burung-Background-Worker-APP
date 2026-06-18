package historical_migrations

import (
	"context"
	"fmt"
	"sync"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"

)

func DownRelation(ctx context.Context, session *gocql.Session) []error {
	var errs []error
	var wg sync.WaitGroup
	var mu sync.RWMutex

	for _, model := range model_list {
		wg.Add(1)
		fctx, cancel := context.WithTimeout(ctx, time.Second*6)
		go func(konteks context.Context, ctxCancel context.CancelFunc, m interface{}) {
			defer wg.Done()
			defer ctxCancel()

			if historicalModel, ok := m.(cass_models.Method); ok {
				if err := historicalModel.DropTable(ctx, session); err != nil {
					mu.Lock()
					errs = append(errs, err)
					mu.Unlock()
				}
			} else {
				fmt.Printf("Objek %T tidak mengimplementasikan historical_models.Method\n", m)
			}
		}(fctx, cancel, model)
	}
	wg.Wait()

	return errs

}
