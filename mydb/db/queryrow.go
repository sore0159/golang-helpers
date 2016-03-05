package db

func QueryRow(db DBer, query string, qArgsList [][]interface{}, scanArgsList [][]interface{}) error {
	stmt, err := db.Prepare(query)
	if my, bad := Check(err, "MultiQuery prepare failure", "query", query); bad {
		return my
	}
	defer stmt.Close()
	for i, qArgs := range qArgsList {
		scanArgs := scanArgsList[i]
		row := stmt.QueryRow(qArgs...)
		if my, bad := Check(row.Scan(scanArgs...), "MultiQuery rowscan failure", "index", i, "query", query); bad {
			return my
		}
	}
	return nil
}

func SingleQueryRow(db DBer, query string, qArgs []interface{}, scanArgs []interface{}) (found bool, err error) {
	row := db.QueryRow(query, qArgs...)
	err = row.Scan(scanArgs...)
	if err == ErrNoRows {
		return false, nil
	}
	if my, bad := Check(err, "queryrowscan failure on scan", "query", query, "qArgs", qArgs, "scanArgs", scanArgs); bad {
		return false, my
	}
	return true, nil
}
