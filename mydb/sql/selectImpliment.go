package sql

import (
	"fmt"
	"strings"
)

type selector struct {
	Cols  []string
	Table string
	Where Condition
	Order string
}

func (s *selector) COLS(cols ...string) *selector {
	s.Cols = cols
	return s
}
func (s *selector) WHERE(cond Condition) *selector {
	s.Where = cond
	return s
}
func (s *selector) ORDER(ord string) *selector {
	s.Order = ord
	return s
}

func (s *selector) Compile() (query string, args []interface{}) {
	var colStr string
	if len(s.Cols) == 0 {
		colStr = "*"
	} else {
		colStr = strings.Join(s.Cols, ",")
	}
	query = fmt.Sprintf("SELECT %s FROM %s", colStr, s.Table)
	if s.Where != nil {
		var str string
		str, args = s.Where.SQL(args)
		query = fmt.Sprintf("%s WHERE %s", query, str)
	}
	if s.Order != "" {
		query = fmt.Sprintf("%s ORDER BY %s", query, s.Order)
	}
	return query, args
}
