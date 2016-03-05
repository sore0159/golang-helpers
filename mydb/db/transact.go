package db

import (
	"database/sql"
)

func Transact(db *sql.DB, f func(DBer) error) (err error) {
	tx, err := db.Begin()
	if my, bad := Check(err, "transaction begin failure"); bad {
		return my
	}
	err = f(tx)
	if my, bad := Check(err, "transact execute failure"); bad {
		if err2 := tx.Rollback(); err2 != nil {
			my.AddContext("rollback failure", err2)
		}
		return my
	} else {
		if my, bad := Check(tx.Commit(), "transact commit failure"); bad {
			return my
		}
		return nil
	}
}
