package group

import (
	"mule/mydb/db"
	sq "mule/mydb/sql"
)

func Update(d db.DBer, group UpdateGrouper) error {
	sqlers := group.UpdateList()
	if len(sqlers) == 0 {
		return nil
	}
	table := group.SQLTable()
	cols := group.UpdateCols()
	pkCols := group.PKCols()
	builder := sq.UPDATE(table).COLS(cols...).WHERE(WherePK(pkCols))
	query, _ := builder.Compile()
	qArgsList := make([][]interface{}, len(sqlers))
	for i, item := range sqlers {
		list, err := ColVals(item, cols)
		if my, bad := Check(err, "group update fail on ColVals for item", "index", i, "item", item, "cols", cols); bad {
			return my
		}
		qArgsList[i] = list
	}
	return db.Exec(d, true, query, qArgsList...)
}

// Select is for the main use case of a series of columns that
// all must match:
// Example:
//		err := group.Select(d, group, group.EQ{"gid", 1}, group.EQ{"planet", 2})
func Select(d db.DBer, group SelectGrouper, conditions ...EQ) error {
	where := sq.AllEQ(Convert2P(conditions...)...)
	return SelectWhere(d, group, where)
}

func SelectWhere(d db.DBer, group SelectGrouper, where sq.Condition) error {
	table := group.SQLTable()
	sCols := group.SelectCols()
	builder := sq.SELECT(table).COLS(sCols...)
	if where != nil {
		builder = builder.WHERE(where)
	}
	query, qArgs := builder.Compile()
	var makeErr error
	maker := func() []interface{} {
		item := group.New()
		list, err := ColPtrs(item, sCols)
		if err != nil {
			makeErr = err
		}
		return list
	}
	err := db.Query(d, query, qArgs, maker)
	if my, bad := Check(makeErr, "select failure on colPtrs in make function", "sCols", sCols); bad {
		return my
	}
	if my, bad := Check(err, "select failure on queryscan", "query", query, "qArgs", qArgs); bad {
		return my
	}
	return nil
}

func Insert(d db.DBer, group InsertGrouper) error {
	sqlers := group.InsertList()
	if len(sqlers) == 0 {
		return nil
	}
	table := group.SQLTable()
	iCols := group.InsertCols()
	sCols := group.InsertScanCols()
	builder := sq.INSERT(table).COLS(iCols...)
	willScan := len(sCols) > 0
	if willScan {
		builder = builder.RETURN(sCols...)
	}
	query, _ := builder.Compile()
	qArgsList := make([][]interface{}, len(sqlers))
	var scanArgsList [][]interface{}
	if willScan {
		scanArgsList = make([][]interface{}, len(sqlers))
	}
	var err error
	for i, item := range sqlers {
		qArgsList[i], err = ColVals(item, iCols)
		if my, bad := Check(err, "insert failure on ColVals qArgs", "iCols", iCols, "index", i); bad {
			return my
		}
		if willScan {
			scanArgsList[i], err = ColPtrs(item, sCols)
			if my, bad := Check(err, "insert failure on ColPtrs sArgs", "sCols", sCols, "index", i); bad {
				return my
			}
		}
	}
	if willScan {
		return db.QueryRow(d, query, qArgsList, scanArgsList)
	} else {
		return db.Exec(d, true, query, qArgsList...)
	}
}

func Delete(d db.DBer, group DeleteGrouper) error {
	sqlers := group.DeleteList()
	if len(sqlers) == 0 {
		return nil
	}
	table := group.SQLTable()
	pkCols := group.PKCols()
	query, _ := sq.DELETE(table).WHERE(WherePK(pkCols)).Compile()
	argList := make([][]interface{}, len(sqlers))
	var err error
	for i, item := range sqlers {
		argList[i], err = ColVals(item, pkCols)
		if my, bad := Check(err, "delete failure on Colvals pkcols", "pkCols", pkCols, "index", i); bad {
			return my
		}
	}
	return db.Exec(d, true, query, argList...)
}

func DeleteWhere(d db.DBer, table string, where sq.Condition) error {
	query, args := sq.DELETE(table).WHERE(where).Compile()
	return db.Exec(d, false, query, args)
}
