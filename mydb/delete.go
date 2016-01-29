package mydb

import "fmt"

func Delete(db DBer, table string, pkCols []string, items ...SQLer) error {
	itemLen := len(items)
	if itemLen == 0 {
		return nil
	}
	conds := make([]interface{}, 0, len(pkCols)*2)
	for _, col := range pkCols {
		conds = append(conds, col, nil)
	}
	query, _, err := DeleteQA(table, conds)
	if my, bad := Check(err, "delete querygen failure", "table", table, "conds", conds); bad {
		return my
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
