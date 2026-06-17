package cass_models

import (
	"context"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

)

type TableName interface {
	TableNameHistorical() string
	TableNameSotReplica() string
}

type Method interface {
	CreateTable(ctx context.Context, s *gocql.Session) error
	UpdateTable(ctx context.Context, s *gocql.Session) error
	ParseToCUDType() map[string]interface{}
	DropTable(ctx context.Context, s *gocql.Session) error
}
