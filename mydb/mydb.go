package mydb

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"mule/mybad"
)

var Check = mybad.BuildCheck("package", "mydb")

// SQLer allows functions to use either db or tx
type DBer interface {
	Exec(string, ...interface{}) (sql.Result, error)
	QueryRow(string, ...interface{}) *sql.Row
	Query(string, ...interface{}) (*sql.Rows, error)
	Prepare(string) (*sql.Stmt, error)
}

// Maybe this shouldn't be here to keep this package SQL-flavor agnostic.
// Who knows?  Useful for me, for now.
func LoadDB(user, pass, dbName string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, pass, dbName))
	if my, bad := Check(err, "loaddb failure", "username", user, "dbname", dbName); bad {
		return nil, my
	}
	err = db.Ping()
	if my, bad := Check(err, "loaddb ping failure", "username", user, "dbname", dbName); bad {
		return nil, my
	}
	return db, nil
}
