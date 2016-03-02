package sql

import (
	"fmt"
	"strings"
)

type selector struct {
	SelectCols []string
	Table      string
	Where      Condition
	Order      string
}

func (s *selector) FROM(table string) *selector {
	s.Table = table
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
	query = fmt.Sprintf("SELECT %s FROM %s", strings.Join(s.SelectCols, ","), s.Table)
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

type Condition interface {
	SQL(args []interface{}) (string, []interface{})
}

type Compare struct {
	Operator string
	ColName  string
	ColVal   interface{}
}

func (c Compare) SQL(args []interface{}) (string, []interface{}) {
	args = append(args, c.ColVal)
	count := len(args)
	return fmt.Sprintf("%s %s $%d", c.ColName, c.Operator, count), args
}

type In struct {
	ColName string
	InVals  []interface{}
}

func (s In) SQL(args []interface{}) (string, []interface{}) {
	inParts := make([]string, len(s.InVals))
	count := len(args)
	for i, val := range s.InVals {
		count += 1
		inParts[i] = fmt.Sprintf("$%d", count)
		args = append(args, val)
	}
	return fmt.Sprintf("%s IN (%s)", s.ColName, strings.Join(inParts, ",")),
		args
}

type Conjunction struct {
	Joiner string
	Parts  []Condition
}

func (c Conjunction) SQL(args []interface{}) (string, []interface{}) {
	strs := make([]string, len(c.Parts))
	var str string
	for i, p := range c.Parts {
		str, args = p.SQL(args)
		strs[i] = fmt.Sprintf("(%s)", str)
	}
	return strings.Join(strs, fmt.Sprintf(" %s ", c.Joiner)), args
}
