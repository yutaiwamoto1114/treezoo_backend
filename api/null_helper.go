package api

import (
	"database/sql"
	"time"
)

// ヘルパー関数: Null値の適切な変換
func nilIfNullString(ns sql.NullString) interface{} {
	if ns.Valid {
		return ns.String
	}
	return nil
}

func nilIfNullBool(nb sql.NullBool) interface{} {
	if nb.Valid {
		return nb.Bool
	}
	return nil
}

func nilIfNullInt64(ni sql.NullInt64) interface{} {
	if ni.Valid {
		return ni.Int64
	}
	return nil
}

func nilIfNullTime(nt sql.NullTime) interface{} {
	if nt.Valid {
		return nt.Time.Format(time.RFC3339) // ISO8601 形式で返す
	}
	return nil
}
