package mydb

import (
	"fmt"
	sq "mule/mydb/sql"
)

func Delete(db DBer, table string, pkCols []string, items ...SQLer) error {
	itemLen := len(items)
	if itemLen == 0 {
		return nil
	}
	var query string
	if len(pkCols) == 1 {
		query, _ = sq.DELETE(table).WHERE(sq.EQ(pkCols[0], nil)).Compile()
	} else {
		conds := make([]sq.Condition, len(pkCols))
		for i, pk := range pkCols {
			conds[i] = sq.EQ(pk, nil)
		}
		query, _ = sq.DELETE(table).WHERE(sq.AND(conds...)).Compile()
	}
	stmt, err := db.Prepare(query)
	if my, bad := Check(err, "delete prepare failure", "query", query); bad {
		return my
	}
	defer stmt.Close()
	for i, item := range items {
		args, err := ColVals(item, pkCols)
		if my, bad := Check(err, "delete colvals failure", "index", i, "query", query, "pkCols", pkCols); bad {
			return my
		}
		res, err := stmt.Exec(args...)
		if my, bad := Check(err, "delete exec failure", "index", i, "query", query, "args", args); bad {
			return my
		}
		aff, err := res.RowsAffected()
		if my, bad := Check(err, "delete rows affected failure", "index", i, "query", query, "args", args); bad {
			return my
		}
		if aff != 1 {
			my, _ := Check(fmt.Errorf("inadequate rows affected %v", aff), "delete rows affected failure", "index", i, "query", query, "args", args, "affected", aff)
			return my
		}
	}
	return nil
}
