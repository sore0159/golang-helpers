package sql

import (
	"fmt"
	"strings"
)

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
