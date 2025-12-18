package db

import (
	"context"
	"database/sql"
)

// Exec executes a query without returning any rows (for Update, Delete)
func Exec(ctx context.Context, query string, args ...any) (int64, error) {
	result, err := DB.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// Update executes an update query and returns the number of rows affected
func Update(ctx context.Context, query string, args ...any) (int64, error) {
	return Exec(ctx, query, args...)
}

// Delete executes a delete query and returns the number of rows affected
func Delete(ctx context.Context, query string, args ...any) (int64, error) {
	return Exec(ctx, query, args...)
}

// Insert executes an insert query and returns the LastInsertId
func Insert(ctx context.Context, query string, args ...any) (int64, error) {
	result, err := DB.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// FindAll executes a query and scans all rows into the dst slice
func FindAll(ctx context.Context, query string, dst any, args ...any) error {
	rows, err := DB.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return Scan(rows, dst)
}

// FindOne executes a query and scans the first row into the dst struct
func FindOne(ctx context.Context, query string, dst any, args ...any) error {
	rows, err := DB.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return Scan(rows, dst)
}

// FindAllTx is FindAll but uses an existing transaction
func FindAllTx(ctx context.Context, tx *sql.Tx, query string, dst any, args ...any) error {
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return Scan(rows, dst)
}

// FindOneTx is FindOne but uses an existing transaction
func FindOneTx(ctx context.Context, tx *sql.Tx, query string, dst any, args ...any) error {
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return Scan(rows, dst)
}

// InsertTx is Insert but uses an existing transaction
func InsertTx(ctx context.Context, tx *sql.Tx, query string, args ...any) (int64, error) {
	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
