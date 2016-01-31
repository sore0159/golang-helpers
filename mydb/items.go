package mydb

func UpdateItem(db DBer, table string, set, conditions []interface{}) error {
	setCols, setArgs, err := C(set).Split()
	if my, bad := Check(err, "update item failure on set splic", "set", set, "conditions", conditions); bad {
		return my
	}
	condCols, condArgs, err := C(conditions).Split()
	if my, bad := Check(err, "update item failure on conditions split", "set", set, "conditions", conditions); bad {
		return my
	}
	query := UpdateQ(table, setCols, condCols)
	args := append(setArgs, condArgs...)
	err = ExecCheck(db.Exec(query, args...))
	if my, bad := Check(err, "update item failure", "table", table, "query", query, "args", args); bad {
		return my
	}
	return nil
}

func DropItems(db DBer, table string, conditions []interface{}) error {
	query, args, err := DeleteQA(table, conditions)
	if my, bad := Check(err, "dropitem failure", "table", table, "conditions", conditions); bad {
		return my
	}
	return ExecCheck(db.Exec(query, args...))
}

func DropItemsIf(db DBer, table string, conditions []interface{}) error {
	query, args, err := DeleteQA(table, conditions)
	if my, bad := Check(err, "dropitem failure", "table", table, "conditions", conditions); bad {
		return my
	}
	_, err = db.Exec(query, args...)
	return err
}
