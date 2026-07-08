package historical_migrations

import (
	"context"
	"fmt"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

	cass_models "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/models"
)

func DownRelation(ctx context.Context, s *gocql.Session) []error {
	var errs []error = []error{}

	for _, model := range model_list {
		fctx, cancel := context.WithTimeout(ctx, time.Second*10)

		if historicalModel, ok := model.(cass_models.TableName); ok {
			if err := cass_models.DropTable(fctx, s, historicalModel.TableNameHistorical()); err != nil {
				errs = append(errs, err)
			}
		} else {
			fmt.Printf("Objek %T tidak mengimplementasikan historical_models.Method\n", model)
		}

		cancel()
	}

	return errs
}
