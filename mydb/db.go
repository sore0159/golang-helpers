package mydb

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var Log = log.Println

func SetLogger(f func(...interface{})) {
	Log = f
}

// SQLer in a type to allow for using either transactions
// or databases to perform operations
type SQLer interface {
	Exec(string, ...interface{}) (sql.Result, error)
	QueryRow(string, ...interface{}) *sql.Row
	Query(string, ...interface{}) (*sql.Rows, error)
}

// Maybe this shouldn't be here to keep this package SQL-flavor agnostic.
// Who knows?  Useful for me, for now.
func LoadDB(user, pass, dbName string) (*sql.DB, bool) {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, pass, dbName))
	if err != nil {
		Log("loaddb error:", err)
		return nil, false
	}
	err = db.Ping()
	if err != nil {
		Log("loaddb ping error:", err)
		return nil, false
	}
	return db, true
}
