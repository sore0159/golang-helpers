package group

import (
	"mule/mybad"
	sq "mule/mydb/sql"
)

var (
	Check = mybad.BuildCheck("package", "mydb/group")
)

// SQLStruct is just something for templating generics
type SQLStruct struct {
	INSERT bool
	DELETE bool
	UPDATE bool
}

func WherePK(pkCols []string) sq.Condition {
	wheres := make([]sq.P, len(pkCols))
	for i, col := range pkCols {
		wheres[i] = sq.P{col, nil}
	}
	return sq.AllEQ(
		wheres...,
	)
}

type GroupMaker interface {
	SQLGroup() SQLGrouper
}
