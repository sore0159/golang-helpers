package mydb

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"mule/mylog"
)

// Maybe detatch this package from mylog at some point
var Log = mylog.Err

// SQLer in a type to allow for using either transactions
// or databases to perform operations
type SQLer interface {
	Exec(string, ...interface{}) (sql.Result, error)
	QueryRow(string, ...interface{}) *sql.Row
	Query(string, ...interface{}) (*sql.Rows, error)
}

// Maybe this shouldn't be here to keep this package SQL-flavor agnostic.
// Who knows?  Useful for me, for now.
func LoadDB(user, pass, dbName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, pass, dbName))
	if err != nil {
		return nil, Log("loaddb error:", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, Log("loaddb ping error:", err)
	}
	return db, nil
}
