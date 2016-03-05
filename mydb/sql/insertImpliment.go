package sql

import (
	"fmt"
	"strings"
)

type inserter struct {
	Table  string
	Vals   []Change
	Return []string
}

func (u *inserter) COLS(cols ...string) *inserter {
	c := make([]Change, len(cols))
	for i, n := range cols {
		c[i].Col = n
	}
	u.Vals = c
	return u
}
func (u *inserter) TO(vals ...interface{}) *inserter {
	for i, v := range vals {
		u.Vals[i].Val = v
	}
	return u
}
func (u *inserter) RETURN(cols ...string) *inserter {
	u.Return = cols
	return u
}
func (u *inserter) Args() (args []interface{}) {
	for _, c := range u.Vals {
		args = append(args, c.Val)
	}
	return args
}
func (u *inserter) Compile() (query string, args []interface{}) {
	colParts := make([]string, len(u.Vals))
	setParts := make([]string, len(u.Vals))
	for i, c := range u.Vals {
		args = append(args, c.Val)
		colParts[i] = c.Col
		setParts[i] = fmt.Sprintf("$%d", len(args))
	}
	query = fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		u.Table, strings.Join(colParts, ","), strings.Join(setParts, ","),
	)
	if len(u.Return) > 0 {
		query = fmt.Sprintf("%s RETURNING %s", query, strings.Join(u.Return, ","))
	}
	return query, args
}
