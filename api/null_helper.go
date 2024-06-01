package api

import (
	"database/sql"
	"time"
)

// ヘルパー関数: Null値の適切な変換

// for all nullable type
func nilIfNull[T any](n sql.Null[T]) interface{} {
	if n.Valid {
		return n.Value
	}
	return nil
}

// for null string
func nilIfNullString(ns sql.NullString) interface{} {
	if ns.Valid {
		return ns.String
	}
	return nil
}

// for null bool
func nilIfNullBool(nb sql.NullBool) interface{} {
	if nb.Valid {
		return nb.Bool
	}
	return nil
}

// for null int64
func nilIfNullInt64(ni sql.NullInt64) interface{} {
	if ni.Valid {
		return ni.Int64
	}
	return nil
}

// for null time
func nilIfNullTime(nt sql.NullTime) interface{} {
	if nt.Valid {
		return nt.Time.Format(time.RFC3339) // ISO8601 形式で返す
	}
	return nil
}
