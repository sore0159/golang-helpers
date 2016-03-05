package sql

import (
	"fmt"
	"strings"
)

type updater struct {
	Table   string
	Changes []Change
	Where   Condition
	Return  []string
}

type Change struct {
	Col string
	Val interface{}
}

func (c Change) Args(args []interface{}) []interface{} {
	return append(args, c.Val)
}
func (c Change) SQL(args []interface{}) (string, []interface{}) {
	args = append(args, c.Val)
	return fmt.Sprintf("%s = $%d", c.Col, len(args)), args
}

func (u *updater) COLS(cols ...string) *updater {
	c := make([]Change, len(cols))
	for i, n := range cols {
		c[i].Col = n
	}
	u.Changes = c
	return u
}
func (u *updater) VALS(vals ...interface{}) *updater {
	for i, v := range vals {
		u.Changes[i].Val = v
	}
	return u
}
func (u *updater) WHERE(cond Condition) *updater {
	u.Where = cond
	return u
}
func (u *updater) RETURN(cols ...string) *updater {
	u.Return = cols
	return u
}
func (u *updater) Args() (args []interface{}) {
	for _, c := range u.Changes {
		args = append(args, c.Val)
	}
	if u.Where != nil {
		args = u.Where.Args(args)
	}
	return args
}
func (u *updater) Compile() (query string, args []interface{}) {
	setParts := make([]string, len(u.Changes))
	for i, c := range u.Changes {
		var str string
		str, args = c.SQL(args)
		setParts[i] = str
	}
	query = fmt.Sprintf("UPDATE %s SET %s", u.Table, strings.Join(setParts, ","))
	if u.Where != nil {
		var str string
		str, args = u.Where.SQL(args)
		query = fmt.Sprintf("%s WHERE %s", query, str)
	}
	if len(u.Return) > 0 {
		query = fmt.Sprintf("%s RETURNING %s", query, strings.Join(u.Return, ","))
	}
	return query, args
}
