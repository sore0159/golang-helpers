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

// EQ is for functions that want to be vardic
// with alternating strings and []interface{}
//
// It's basically so you don't have to import
// mydb/sql yourself to get at sql.P
// if you just want to access this package
type EQ struct {
	Col string
	Val interface{}
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

func Convert2P(mine ...EQ) []sq.P {
	them := make([]sq.P, len(mine))
	for i, x := range mine {
		them[i].Col = x.Col
		them[i].Val = x.Val
	}
	return them
}

type GroupMaker interface {
	SQLGroup() SQLGrouper
}
