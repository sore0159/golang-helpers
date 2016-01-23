package mydb

type SQLerNewer interface {
	New() SQLer
}

// Get is for automated group selecting; it requires a
// SQLerSource interface to handle all the slice population
// for it
//
// For GetOne just do
// err := db.QueryRow(query, queryArgs...).Scan(ptrs...)
// (test against sql.ErrNoRows if wanted)
func Get(db DBer, source SQLerNewer, query string, qArgs ...interface{}) (err error) {
	rows, err := db.Query(query, qArgs...)
	if my, bad := Check(err, "get query failure", "query", query); bad {
		return my
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if my, bad := Check(err, "scan rows columns failure", "query", query); bad {
		return my
	}
	for rows.Next() {
		next := source.New()
		ptrs, err := ColPtrs(next, cols)
		if my, bad := Check(err, "scanrows next ptrs failure", "cols", cols); bad {
			return my
		}
		err = rows.Scan(ptrs...)
		if my, bad := Check(err, "scan rowscan failure", "query", query, "cols", cols); bad {
			return my
		}
	}
	if my, bad := Check(rows.Err(), "scan rows final", "query", query); bad {
		return my
	}
	return nil
}
