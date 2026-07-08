package cass_models

import (
	"context"
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

type TableName interface {
	TableNameHistorical() string
	TableNameSotReplica() string
}

type Method interface {
	CreateHistoricalTable(ctx context.Context, s *gocql.Session) error
	CreateSotReplicaTable(ctx context.Context, s *gocql.Session) error
	ParseToCUDType() map[string]interface{}
}

func DropTable(ctx context.Context, session *gocql.Session, tablename string) error {
	query := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, tablename)

	if err := session.Query(query).ExecContext(ctx); err != nil {
		return fmt.Errorf("gagal drop tabel %s: %w", tablename, err)
	}

	fmt.Printf("Berhasil drop tabel %s\n", tablename)
	return nil
}
