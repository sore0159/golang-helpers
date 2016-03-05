package sql

import (
	"fmt"
	"strings"
)

type Condition interface {
	SQL(args []interface{}) (string, []interface{})
	Args(args []interface{}) []interface{}
}

type Compare struct {
	Operator string
	ColName  string
	ColVal   interface{}
}

func (c Compare) Args(args []interface{}) []interface{} {
	return append(args, c.ColVal)
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

func (c In) Args(args []interface{}) []interface{} {
	return append(args, c.InVals...)
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

func MakeConjunction(join string, parts ...Condition) Condition {
	l := len(parts)
	if l == 0 {
		return nil
	} else if l == 1 {
		return parts[0]
	}
	return Conjunction{Joiner: join, Parts: parts}
}
func (c Conjunction) Args(args []interface{}) []interface{} {
	for _, p := range c.Parts {
		args = p.Args(args)
	}
	return args
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
