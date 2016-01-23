package mydb

import (
	"fmt"
)

// Insert takes an Inserter and runs InsertQ, calling Scan based
// on the second value of InsertQ
// Insert returns a success bool and logs failures
func Insert(db SQLer, query string, cols []string, scanCols []string, items ...SQLColumner) (err error) {
	stmt, err := db.Prepare(query)
	if my, bad := Check(err, "insert prepare failure", "query", query); bad {
		return my
	}
	defer stmt.Close()
	for i, item := range items {
		args, err := ColVals(item, cols)
		if my, bad := Check(err, "insert colvals failure", "index", i, "query", query); bad {
			return my
		}
		if len(scanCols) > 0 {
			ptrs, err := ColPtrs(item, scanCols)
			if my, bad := Check(err, "insert colptrs failure", "index", i, "query", query); bad {
				return my
			}
			row := stmt.QueryRow(args...)
			if my, bad := Check(row.Scan(ptrs...), "insert rowscan failure", "index", i, "query", query); bad {
				return my
			}
		} else {
			res, err := stmt.Exec(args...)
			if my, bad := Check(err, "insert stmt exec failure", "index", i, "query", query); bad {
				return my
			}
			aff, err := res.RowsAffected()
			if my, bad := Check(err, "insert rows affected failure", "index", i, "query", query); bad {
				return my
			}
			if aff < 1 {
				my, _ := Check(fmt.Errorf("insert rows affected inadequate"), "exec rows affected failure", "index", i, "query", query, "affected", aff)
				return my
			}
		}
	}
	return nil
}
