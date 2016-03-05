package db

func Query(db DBer, query string, qArgs []interface{}, scanArgMaker func() []interface{}) (err error) {
	rows, err := db.Query(query, qArgs...)
	if my, bad := Check(err, "queryscan rows creation failure", "query", query); bad {
		return my
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(scanArgMaker()...)
		if my, bad := Check(err, "queryscan rowscan failure", "query", query); bad {
			return my
		}
	}
	if my, bad := Check(rows.Err(), "queryscan rows final", "query", query); bad {
		return my
	}
	return nil
}
