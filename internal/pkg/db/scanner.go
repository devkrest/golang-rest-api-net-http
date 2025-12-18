package db

import (
	"database/sql"
	"fmt"
	"reflect"
)

// scanPlan pre-calculates the mapping between SQL columns and struct fields
type scanPlan struct {
	fieldIndices []int
}

// buildPlan creates a scanPlan for a given type and set of SQL columns
func buildPlan(t reflect.Type, columns []string) *scanPlan {
	plan := &scanPlan{
		fieldIndices: make([]int, len(columns)),
	}

	// Default to -1 (skip column)
	for i := range plan.fieldIndices {
		plan.fieldIndices[i] = -1
	}

	// Create a temporary map of db tags to field indices
	tagToIndex := make(map[string]int)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("db")
		if tag != "" {
			tagToIndex[tag] = i
		}
	}

	// Map columns to indices
	for i, col := range columns {
		if index, ok := tagToIndex[col]; ok {
			plan.fieldIndices[i] = index
		}
	}

	return plan
}

// Scan scans multiple rows into a slice of structs or a single struct
func Scan(rows *sql.Rows, dst any) error {
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("dst must be a pointer")
	}

	elem := v.Elem()
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// Handle slice: []*Struct or []Struct
	if elem.Kind() == reflect.Slice {
		return scanSlice(rows, elem, columns)
	}

	// Handle single struct: *Struct
	if elem.Kind() == reflect.Struct {
		if !rows.Next() {
			return sql.ErrNoRows
		}
		plan := buildPlan(elem.Type(), columns)
		return scanItem(rows, elem, plan)
	}

	return fmt.Errorf("unsupported dst type: %s", elem.Kind())
}

func scanSlice(rows *sql.Rows, slice reflect.Value, columns []string) error {
	itemType := slice.Type().Elem()
	isPtr := itemType.Kind() == reflect.Ptr
	if isPtr {
		itemType = itemType.Elem()
	}

	plan := buildPlan(itemType, columns)

	for rows.Next() {
		item := reflect.New(itemType).Elem()
		if err := scanItem(rows, item, plan); err != nil {
			return err
		}

		if isPtr {
			slice.Set(reflect.Append(slice, item.Addr()))
		} else {
			slice.Set(reflect.Append(slice, item))
		}
	}
	return rows.Err()
}

func scanItem(rows *sql.Rows, v reflect.Value, plan *scanPlan) error {
	pointers := make([]any, len(plan.fieldIndices))
	
	for i, fieldIdx := range plan.fieldIndices {
		if fieldIdx != -1 {
			pointers[i] = v.Field(fieldIdx).Addr().Interface()
		} else {
			var skip any
			pointers[i] = &skip
		}
	}

	return rows.Scan(pointers...)
}
