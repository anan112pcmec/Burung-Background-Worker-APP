package sot_replica_migration

import (
	"context"
	"fmt"
	"sync"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
)

func UpRelation(ctx context.Context, s *gocql.Session) []error {
	var errs []error = []error{}
	var wg sync.WaitGroup
	var rw sync.RWMutex
	for _, v := range model_list {
		fctx, cancel := context.WithTimeout(ctx, time.Second*6)
		wg.Add(1)
		go func(konteks context.Context, batal context.CancelFunc, model interface{}, sesi *gocql.Session) {
			defer wg.Done()
			defer batal()

			if sotReplica, tru := model.(cass_models.Method); tru {
				if err := sotReplica.CreateSotReplicaTable(konteks, sesi); err != nil {
					rw.Lock()
					errs = append(errs, err)
					rw.Unlock()
				}
			} else {
				fmt.Printf("Objek %T tidak mengimplementasikan sot_replica.Method\n", model)
			}

		}(fctx, cancel, v, s)
	}

	wg.Wait()
	return errs
}
