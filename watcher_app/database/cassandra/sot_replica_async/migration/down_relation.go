package sot_replica_migration

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
)

func DownMigration(ctx context.Context, s *gocql.Session) []error {
	var errs []error = []error{}

	for _, v := range model_list {
		fctx, cancel := context.WithTimeout(ctx, time.Second*10)

		if sotReplica, ok := v.(cass_models.Method); ok {
			if err := sotReplica.DropTable(fctx, s); err != nil {
				errs = append(errs, err)
			}
		} else {
			fmt.Printf("Objek %T tidak mengimplementasikan sot_replica.Method\n", v)
		}

		cancel()
	}

	return errs
}
