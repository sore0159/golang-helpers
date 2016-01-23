package mydb

import "fmt"

// Exec executes the query and checks if anything is affected
//
// If you don't need rows affected just do
// _, err := db.Exec(query, qArgs...)
func Exec(db SQLer, query string, qArgs ...interface{}) (err error) {
	res, err := db.Exec(query, qArgs...)
	if my, bad := Check(err, "exec failure", "query", query); bad {
		return my
	}
	aff, err := res.RowsAffected()
	if my, bad := Check(err, "rows affected failure", "query", query); bad {
		return my
	}
	if aff < 1 {
		my, _ := Check(fmt.Errorf("rows affected inadequate"), "exec rows affected failure", "query", query, "affected", aff)
		return my
	}
	return nil
}
