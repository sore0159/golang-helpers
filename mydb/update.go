package mydb

import "fmt"

// Update is for mass updating and requires a
// qArg maker function for its iterations
//
// For UpdateOne just do
// err = db.Exec(query, args..)
func Update(db SQLer, must bool, query string, cols []string, items ...SQLColumner) error {
	stmt, err := db.Prepare(query)
	if my, bad := Check(err, "update prepare failure", "query", query); bad {
		return my
	}
	defer stmt.Close()
	for i, item := range items {
		args, err := ColVals(item, cols)
		if my, bad := Check(err, "update colvals failure", "index", i, "query", query, "cols", cols); bad {
			return my
		}
		res, err := stmt.Exec(args...)
		if my, bad := Check(err, "update exec failure", "index", i, "query", query); bad {
			return my
		}
		if must {
			aff, err := res.RowsAffected()
			if my, bad := Check(err, "update rows affected failure", "index", i, "query", query); bad {
				return my
			}
			if aff != 1 {
				my, _ := Check(fmt.Errorf("inadequate rows affected", aff), "update rows affected failure", "index", i, "query", query, "affected", aff)
				return my
			}
		}
	}
	return nil
}
