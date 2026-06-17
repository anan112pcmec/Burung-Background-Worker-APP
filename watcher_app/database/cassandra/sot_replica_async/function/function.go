package sot_replica_function

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

const timeout = 6

func InsertReplicaData(ctx context.Context, session *gocql.Session, tablename string, append_data *[]map[string]interface{}) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs []error

	for _, data := range *append_data {
		if len(data) == 0 {
			continue
		}

		wg.Add(1)
		fctx, cancel := context.WithTimeout(ctx, time.Second*timeout)

		go func(konteks context.Context, ctxCancel context.CancelFunc, d map[string]interface{}) {
			defer wg.Done()
			defer ctxCancel()

			total := len(d)

			columns := make([]string, 0, total)
			placeholders := make([]string, 0, total)
			values := make([]interface{}, 0, total)

			for col, val := range d {
				columns = append(columns, col)
				placeholders = append(placeholders, "?")
				values = append(values, val)
			}

			queryStr := fmt.Sprintf(
				"INSERT INTO %s (%s) VALUES (%s)",
				tablename,
				strings.Join(columns, ", "),
				strings.Join(placeholders, ", "),
			)

			fmt.Println(queryStr)

			err := session.Query(queryStr, values...).
				Consistency(gocql.Quorum).
				ExecContext(konteks)

			if err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("failed to insert: %w", err))
				mu.Unlock()
			}
		}(fctx, cancel, data)
	}

	wg.Wait()

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}

func UpdateData() {}

func DeleteData() {}
