package mydb

import (
	"database/sql"
	"reflect"
)

// Scanner lets us use Rowers on both sql.Rows and sql.Row
type Scanner interface {
	Scan(...interface{}) error
}

// Rower values know what values to give a scanner for a scan
type Rower interface {
	RowScan(Scanner) error
}

// Get runs the query and fills the list referenced by listPtr
// with data constructed by maker() and filled by calling RowScan
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
		Log("getall query error:", err, "\n||", query)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		next := maker()
		if err = next.RowScan(rows); err != nil {
			Log("getall scanning error:", err, "\n||", query)
			return false
		}
		listV.Set(reflect.Append(listV, reflect.ValueOf(next)))
	}
	if err = rows.Err(); err != nil {
		Log("getall rows final error:", err, "\n||", query)
		return false
	}
	return true
}

// GetOne is a simple QueryRow that logs errors and
// returns a successbool: null result is an error
func GetOne(db SQLer, query string, item Rower) (ok bool) {
	row := db.QueryRow(query)
	err := item.RowScan(row)
	if err != nil {
		Log("mydb GetOne scan error:", err, "\n||", query)
		return false
	}
	return true
}

// GetOneIf is a simple QueryRow that logs errors and
// returns a successbool; null result is false but
// but is not logged
func GetOneIf(db SQLer, query string, item Rower) (ok bool) {
	row := db.QueryRow(query)
	err := item.RowScan(row)
	if err != nil {
		if err != sql.ErrNoRows {
			Log("mydb GetOne scan error:", err, "\n||", query)
		}
		return false
	}
	return true
}
