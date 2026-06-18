package cass_cud

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"

	historical_format "github.com/anan112pcmec/Burung-backend-2/watcher_app/database/cassandra/hystorical_db/format"
)

const timeout = 6

func InsertData(ctx context.Context, session *gocql.Session, tablename string, append_data *[]map[string]interface{}) error {
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

			pencatatan := historical_format.Sekarang()
			d["tahun_update"] = pencatatan.TahunUpdate
			d["bulan_update"] = pencatatan.BulanUpdate
			d["event_time"] = pencatatan.EventTime

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

func UpdateData(ctx context.Context, session *gocql.Session, tablename string, update_data *[]map[string]interface{}, conditions map[string]interface{}) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs []error

	for _, data := range *update_data {
		if len(data) == 0 {
			continue
		}

		wg.Add(1)
		fctx, cancel := context.WithTimeout(ctx, time.Second*timeout)

		go func(konteks context.Context, ctxCancel context.CancelFunc, d map[string]interface{}) {
			defer wg.Done()
			defer ctxCancel()

			// Pastikan metadata waktu tetap diperbarui saat update jika diperlukan
			pencatatan := historical_format.Sekarang()
			d["tahun_update"] = pencatatan.TahunUpdate
			d["bulan_update"] = pencatatan.BulanUpdate
			d["event_time"] = pencatatan.EventTime

			setClauses := make([]string, 0, len(d))
			whereClauses := make([]string, 0, len(conditions))
			values := make([]interface{}, 0, len(d)+len(conditions))

			// 1. Membangun SET klausa
			for col, val := range d {
				setClauses = append(setClauses, fmt.Sprintf("%s = ?", col))
				values = append(values, val)
			}

			// 2. Membangun WHERE klausa (Primary Keys)
			for col, val := range conditions {
				whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", col))
				values = append(values, val)
			}

			queryStr := fmt.Sprintf(
				"UPDATE %s SET %s WHERE %s",
				tablename,
				strings.Join(setClauses, ", "),
				strings.Join(whereClauses, " AND "),
			)

			fmt.Println(queryStr)

			err := session.Query(queryStr, values...).
				Consistency(gocql.Quorum).
				ExecContext(konteks)

			if err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("failed to update: %w", err))
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

func DeleteData(ctx context.Context, session *gocql.Session, tablename string, delete_conditions *[]map[string]interface{}) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs []error

	for _, cond := range *delete_conditions {
		if len(cond) == 0 {
			continue
		}

		wg.Add(1)
		fctx, cancel := context.WithTimeout(ctx, time.Second*timeout)

		go func(konteks context.Context, ctxCancel context.CancelFunc, c map[string]interface{}) {
			defer wg.Done()
			defer ctxCancel()

			whereClauses := make([]string, 0, len(c))
			values := make([]interface{}, 0, len(c))

			// Membangun WHERE klausa menggunakan Primary Key
			for col, val := range c {
				whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", col))
				values = append(values, val)
			}

			queryStr := fmt.Sprintf(
				"DELETE FROM %s WHERE %s",
				tablename,
				strings.Join(whereClauses, " AND "),
			)

			fmt.Println(queryStr)

			err := session.Query(queryStr, values...).
				Consistency(gocql.Quorum).
				ExecContext(konteks)

			if err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("failed to delete: %w", err))
				mu.Unlock()
			}
		}(fctx, cancel, cond)
	}

	wg.Wait()

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}
