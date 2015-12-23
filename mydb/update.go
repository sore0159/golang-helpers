package mydb

import (
	"database/sql"
	"reflect"
)

type Updater interface {
	UpdateQ() string
	Commit()
}

type UpdateInserter interface {
	Updater
	Inserter
}

// UpdateList takes a list of items that satisfy Updater and
// returns a list of Updaters
// Probably won't use this much
func UpdateList(list interface{}) []Updater {
	listV := reflect.ValueOf(list)
	if listV.Type().Kind() != reflect.Slice {
		Log("bad type for mydb updatelist: not a slice:", listV)
		return nil
	}
	l := listV.Len()
	if l < 1 {
		return nil
	}
	r := make([]Updater, l)
	for i := 0; i < l; i++ {
		itemV := listV.Index(i)
		test, ok := itemV.Interface().(Updater)
		if ok {
			r[i] = test
		} else {
			Log("bad item", i, "for mydb updatelist: not an Updater", listV)
			return nil
		}
	}
	return r
}

// Update takes a list of Updaters and updates them all in one
// transaction, returning a success bool
// Update rollsback and logs errors on failure
// If UpdateQ gives "" the item will not be updated
func Update(db *sql.DB, list []Updater) (ok bool) {
	tx, err := db.Begin()
	if err != nil {
		Log("update transaction begin error:", err)
		return false
	}
	for _, x := range list {
		query := x.UpdateQ()
		if query == "" {
			continue
		}
		res, err := tx.Exec(query)
		if err != nil {
			Log("update exec failed:", err, "||", query)
			err = tx.Rollback()
			if err != nil {
				Log("update rollback error:", err)
			}
			return false
		}
		if aff, err := res.RowsAffected(); err != nil || aff < 1 {
			if err != nil {
				Log("update rowsaff error:", err)
				err = tx.Rollback()
				if err != nil {
					Log("update rollback error:", err)
				}
			} else {
				Log("update failed: 0 rows affected || ", query)
				err = tx.Rollback()
				if err != nil {
					Log("update rollback error:", err)
				}
			}
			return false
		}
	}
	err = tx.Commit()
	if err != nil {
		Log("update commit error:", err)
		return false
	}
	for _, x := range list {
		x.Commit()
	}
	return true
}

func Upsert(db SQLer, item UpdateInserter) (ok bool) {
	query := item.UpdateQ()
	if query == "" {
		return true
	}
	res, err := db.Exec(query)
	if err != nil {
		Log("upsert exec failed:", err, "||", query)
		return false
	}
	if aff, err := res.RowsAffected(); err != nil {
		Log("upsert rowsaff error:", err)
		return false
	} else if aff == 0 {
		return Insert(db, item)
	}
	return true
}
