package mydb

import (
	"database/sql"
	"fmt"
)

// For wrapping db.Exec(....) to check if rows aff > 0
func ExecCheck(res sql.Result, err error) error {
	if my, bad := Check(err, "exec failure"); bad {
		return my
	}
	aff, err := res.RowsAffected()
	if my, bad := Check(err, "rows affected failure"); bad {
		return my
	}
	if aff < 1 {
		my, _ := Check(fmt.Errorf("rows affected inadequate"), "exec rows affected failure", "affected", aff)
		return my
	}
	return nil
}
