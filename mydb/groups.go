package mydb

type GroupMaker interface {
	SQLGroup() SQLGrouper
}

type SQLGrouper interface {
	SQLTable() string
	PKCols() []string

	DeleteList() []SQLer

	SelectCols() []string
	New() SQLer

	UpdateCols() []string
	UpdateList() []SQLer

	InsertCols() []string
	InsertScanCols() []string
	InsertList() []SQLer
}

type DeleteGrouper interface {
	PKCols() []string
	SQLTable() string
	DeleteList() []SQLer
}

type SelectGrouper interface {
	SQLTable() string
	SelectCols() []string
	New() SQLer
}
type UpdateGrouper interface {
	SQLTable() string
	UpdateCols() []string
	PKCols() []string
	UpdateList() []SQLer
}
type InsertGrouper interface {
	SQLTable() string
	InsertCols() []string
	InsertScanCols() []string
	InsertList() []SQLer
}

func UpdateGroup(db DBer, group UpdateGrouper) error {
	sqlers := group.UpdateList()
	if len(sqlers) == 0 {
		return nil
	}
	table := group.SQLTable()
	cols := group.UpdateCols()
	condCols := group.PKCols()
	query := UpdateQ(table, cols, condCols)
	allCols := append(cols, condCols...)
	return Update(db, true, query, allCols, sqlers...)
}

func GetGroup(db DBer, group SelectGrouper, conditions []interface{}) error {
	table := group.SQLTable()
	cols := group.SelectCols()
	query, args, err := SelectQA(table, cols, conditions)
	if my, bad := Check(err, "get interface failure", "table", table, "conditions", conditions, "cols", cols); bad {
		return my
	}
	err = Get(db, group, query, args...)
	if my, bad := Check(err, "get interface failure", "query", query, "args", args); bad {
		return my
	}
	return nil
}
func MakeGroup(db DBer, group InsertGrouper) error {
	sqlers := group.InsertList()
	if len(sqlers) == 0 {
		return nil
	}
	table := group.SQLTable()
	cols := group.InsertCols()
	scanCols := group.InsertScanCols()
	query := InsertQ(table, cols, scanCols)
	return Insert(db, query, cols, scanCols, sqlers...)
}

func DropGroup(db DBer, group DeleteGrouper) error {
	table := group.SQLTable()
	pkCols := group.PKCols()
	return Delete(db, table, pkCols, group.DeleteList()...)
}
