package historical_models

import (
	"context"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

type Method interface {
	TableName() string
	CreateTable(ctx context.Context, s *gocql.Session) error
	ParseToInsertType() map[string]interface{}
	DropTable(ctx context.Context, s *gocql.Session) error
}
