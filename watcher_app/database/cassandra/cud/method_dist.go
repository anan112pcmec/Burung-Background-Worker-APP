package cass_cud

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

const timeout = 6

type PartitionKey map[string]bool

// id, tahun_update, bulan_update), event_time
var PartitionKeyHistorical PartitionKey = PartitionKey{
	"id":           true,
	"tahun_update": true,
	"bulan_update": true,
	"event_time":   true,
}

var PartitionKeySotReplica PartitionKey = PartitionKey{
	"id": true,
}

func InsertData(ctx context.Context, session *gocql.Session, tablename string, append_data ...map[string]interface{}) error {
	if len(append_data) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs []error

	for _, data := range append_data {
		if len(data) == 0 {
			continue
		}

		wg.Add(1)
		fctx, cancel := context.WithTimeout(ctx, time.Second*timeout)

		go func(konteks context.Context, ctxCancel context.CancelFunc, d map[string]interface{}) {
			defer wg.Done()
			defer ctxCancel()

			// Alokasi memori pas dengan ukuran data input
			totalCapacity := len(d)
			columns := make([]string, 0, totalCapacity)
			placeholders := make([]string, 0, totalCapacity)
			values := make([]interface{}, 0, totalCapacity)

			// Masukkan data murni apa adanya
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

			err := session.Query(queryStr, values...).
				Consistency(gocql.Quorum).
				ExecContext(konteks)

			if err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("failed to insert to %s: %w", tablename, err))
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

func UpdateData(ctx context.Context, session *gocql.Session, tablename string, id int64, update_data ...map[string]interface{}) error {
	if len(update_data) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs []error

	for _, data := range update_data {
		if len(data) == 0 {
			continue
		}

		wg.Add(1)
		fctx, cancel := context.WithTimeout(ctx, time.Second*timeout)

		go func(konteks context.Context, ctxCancel context.CancelFunc, d map[string]interface{}) {
			defer wg.Done()
			defer ctxCancel()

			// Alokasi memori pas + 1 untuk menampung 'id' di klausa WHERE
			setClauses := make([]string, 0, len(d))
			values := make([]interface{}, 0, len(d)+1)

			// Membangun SET klausa dari data asli murni
			for col, val := range d {
				if PartitionKeyHistorical[col] {
					continue
				}

				if PartitionKeySotReplica[col] {
					continue
				}
				setClauses = append(setClauses, fmt.Sprintf("%s = ?", col))
				values = append(values, val)
			}

			queryStr := fmt.Sprintf(
				"UPDATE %s SET %s WHERE id = ?",
				tablename,
				strings.Join(setClauses, ", "),
			)

			// id ditaruh paling akhir setelah nilai SET clauses
			values = append(values, id)

			err := session.Query(queryStr, values...).
				Consistency(gocql.Quorum).
				ExecContext(konteks)

			if err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("failed to update id %s: %w", id, err))
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

func DeleteData(ctx context.Context, session *gocql.Session, tablename string, id ...int64) error {
	if len(id) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs []error

	for _, i := range id {
		wg.Add(1)
		fctx, cancel := context.WithTimeout(ctx, time.Second*timeout)

		go func(konteks context.Context, ctxCancel context.CancelFunc, idnya int64) {
			defer wg.Done()
			defer ctxCancel()

			// PERBAIKAN: Gunakan placeholder '?' demi efisiensi query plan cache di Cassandra
			queryStr := fmt.Sprintf(
				"DELETE FROM %s WHERE id = ?",
				tablename,
			)

			err := session.Query(queryStr, idnya).
				Consistency(gocql.Quorum).
				ExecContext(konteks)

			if err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("failed to delete id %d from %s: %w", idnya, tablename, err))
				mu.Unlock()
			}
		}(fctx, cancel, i)
	}

	wg.Wait()

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}
