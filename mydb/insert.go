package mydb

import (
	"database/sql"
)

// Inserter values have a default insert query and
// scan return behavior
type Inserter interface {
	InsertQ() (query string, scan bool)
	InsertScan(*sql.Row) error
}

// Insert takes an Inserter and runs InsertQ, calling Scan based
// on the second value of InsertQ
// Insert returns a success bool and logs failures
func Insert(db SQLer, item Inserter) (ok bool) {
	query, scan := item.InsertQ()
	if scan {
		row := db.QueryRow(query)
		if err := item.InsertScan(row); err != nil {
			Log("insert scan error:", err, "||", query)
			return false
		}
		return true
	} else {
		res, err := db.Exec(query)
		if err != nil {
			Log("insert exec error:", err, "\n||", query)
			return false
		}
		if aff, err := res.RowsAffected(); err != nil || aff < 1 {
			if err != nil {
				Log("insert exec rowsaff err:", err)
			} else {
				Log("insert exec rowsaff < 1: query", query)
			}
			return false
		}
	}
	return true
}
