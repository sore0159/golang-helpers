package mydb

import (
	"database/sql"
	"reflect"
)

type Rower interface {
	RowsScan(*sql.Rows) error
}

// Get runs the query and fills the list referenced by listPtr
// with data constructed by maker() and filled by calling RowsScan
// Get returns a success bool and logs failures
func Get(db SQLer, query string, listPtr interface{}, maker func() Rower) (ok bool) {
	listPtrV := reflect.ValueOf(listPtr)
	if listPtrV.Type().Kind() != reflect.Ptr {
		Log("bad listPtr for mydb get: not a pointer")
		return false
	}
	listV := listPtrV.Elem()
	rows, err := db.Query(query)
	if err != nil {
		Log("getall query error:", err, "\n", query)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		next := maker()
		if err = next.RowsScan(rows); err != nil {
			Log("getall scanning error:", err, "\n", query)
			return false
		}
		listV.Set(reflect.Append(listV, reflect.ValueOf(next)))
	}
	if err = rows.Err(); err != nil {
		Log("getall rows final error:", err)
		return false
	}
	return true
}
