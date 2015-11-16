package mydb

// Exec just executes the query and returns a success bool (logging failures)
// Affecting 0 rows is considered a failure
func Exec(db SQLer, query string) (ok bool) {
	res, err := db.Exec(query)
	if err != nil {
		Log("query exec err:", err, "||", query)
		return false
	}
	if aff, err := res.RowsAffected(); err != nil || aff < 1 {
		if err != nil {
			Log("insert exec rowsaff err:", err, "||", query)
		} else {
			Log("insert exec rowsaff < 1: query", query)
		}
		return false
	}
	return true
}

// Exec just executes the query and returns a success bool (logging failures)
// Affecting 0 rows is considered a failure
func ExecIf(db SQLer, query string) (ok bool) {
	_, err := db.Exec(query)
	if err != nil {
		Log("query ExecIf err:", err, "||", query)
		return false
	}
	return true
}
