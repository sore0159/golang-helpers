package sql

import (
	"fmt"
)

type dropper struct {
	Table string
	Where Condition
}

func (d *dropper) WHERE(cond Condition) *dropper {
	d.Where = cond
	return d
}

func (d *dropper) Compile() (query string, args []interface{}) {
	query = fmt.Sprintf("DELETE FROM %s", d.Table)
	if d.Where != nil {
		var str string
		str, args = d.Where.SQL(args)
		query = fmt.Sprintf("%s WHERE %s", query, str)
	}
	return query, args
}
func (d *dropper) Args() []interface{} {
	if d.Where == nil {
		return nil
	}
	return d.Where.Args(nil)
}
