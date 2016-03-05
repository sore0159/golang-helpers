package db

import (
	"database/sql"
	"fmt"
)

// Examples:
//		Exec(db, true,
//			"UPDATE TABLE chairs SET age = $1 WHERE chid = $2",
//			[]interface{}{1, 1}, []interface{}{3, 2},
//		)
//		Exec(db, false,
//			"DELETE FROM chairs WHERE legs IN ($1, $2, $3)",
//			[]interface{}{1,2,3},
//		)
func Exec(db DBer, must bool, query string, qArgsList ...[]interface{}) error {
	itemLen := len(qArgsList)
	if itemLen == 0 {
		return nil
	}
	stmt, err := db.Prepare(query)
	if my, bad := Check(err, "delete prepare failure", "query", query); bad {
		return my
	}
	defer stmt.Close()
	if must {
		for i, qArgs := range qArgsList {
			err := ExecCheck(stmt.Exec(qArgs...))
			if my, bad := Check(err, "delete exec failure on Must check", "index", i, "query", query, "args", qArgs); bad {
				return my
			}
		}
	} else {
		for i, qArgs := range qArgsList {
			_, err := stmt.Exec(qArgs...)
			if my, bad := Check(err, "delete exec failure", "index", i, "query", query, "args", qArgs); bad {
				return my
			}
		}
	}
	return nil
}

// For wrapping db.Exec(....) to check if rows aff > 0
func ExecCheck(res sql.Result, err error) error {
	if my, bad := Check(err, "exec failure"); bad {
		return my
	}
	aff, err := res.RowsAffected()
	if my, bad := Check(err, "result rows affected method call failure"); bad {
		return my
	}
	if aff < 1 {
		my, _ := Check(fmt.Errorf("rows affected inadequate"), "exec rows affected failure", "affected", aff)
		return my
	}
	return nil
}
