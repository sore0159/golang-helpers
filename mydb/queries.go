package mydb

import (
	"errors"
	"fmt"
	"strings"
)

func UpdateQ(table string, cols, condCols []string) string {
	parts := make([]string, len(cols))
	for i, name := range cols {
		parts[i] = fmt.Sprintf("%s = $%d", name, i+1)
	}
	setStr := strings.Join(parts, ", ")
	count := len(cols)
	var whereStr string
	if len(condCols) > 0 {
		parts := make([]string, len(condCols))
		for i, name := range condCols {
			parts[i] = fmt.Sprintf("%s = $%d", name, count+i+1)
		}
		whereStr = fmt.Sprintf(" WHERE %s", strings.Join(parts, " AND "))
	}
	return fmt.Sprintf("UPDATE %s SET %s%s", table, setStr, whereStr)
}

func DeleteQA(table string, conditions C) (string, []interface{}, error) {
	whereStr, args, err := conditions.WhereStr(nil)
	if my, bad := Check(err, "DeleteQA fail", "table", table, "conditions", conditions); bad {
		return "", nil, my
	}
	query := fmt.Sprintf("DELETE FROM %s%s", table, whereStr)
	return query, args, nil
}

func SelectQA(table string, cols []string, conditions C) (string, []interface{}, error) {
	whereStr, args, err := conditions.WhereStr(nil)
	if my, bad := Check(err, "selectQA fail", "table", table, "conditions", conditions); bad {
		return "", nil, my
	}
	var colStr string
	if len(cols) > 0 {
		colStr = strings.Join(cols, ", ")
	} else {
		colStr = "*"
	}
	query := fmt.Sprintf("SELECT %s FROM %s%s", colStr, table, whereStr)
	return query, args, nil
}

func InsertQ(table string, cols, scanCols []string) string {
	parts := make([]string, len(cols))
	dollaParts := make([]string, len(cols))
	for i, col := range cols {
		parts[i] = col
		dollaParts[i] = fmt.Sprintf("$%d", i+1)
	}
	colStr := strings.Join(parts, ", ")
	dollaStr := strings.Join(dollaParts, ", ")
	var retStr string
	if len(scanCols) > 0 {
		retStr = fmt.Sprintf(" RETURNING %s", strings.Join(scanCols, ", "))
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)%s", table, colStr, dollaStr, retStr)
	return query
}

type C []interface{}

func (c C) Split() ([]string, []interface{}, error) {
	if len(c)%2 != 0 {
		return nil, nil, errors.New("conditions must be of form key, val")
	}
	hl := len(c) / 2
	keys := make([]string, hl)
	vals := make([]interface{}, hl)
	for i := 0; i < hl; i += 1 {
		key, ok := c[2*i].(string)
		if !ok {
			return nil, nil, errors.New("conditions must be of form key, val")
		}
		keys[i] = key
		vals[i] = c[2*i+1]
	}
	return keys, vals, nil
}

func (c C) WhereStr(curArgs []interface{}) (string, []interface{}, error) {
	if len(c) == 0 {
		return "", curArgs, nil
	}
	cols, vals, err := c.Split()
	if my, bad := Check(err, "wherestr generation failure"); bad {
		return "", curArgs, my
	}
	ln := len(curArgs)
	if curArgs == nil {
		curArgs = vals
	} else {
		curArgs = append(curArgs, vals...)
	}
	parts := make([]string, len(cols))
	for i, name := range cols {
		parts[i] = fmt.Sprintf("%s = $%d", name, i+ln+1)
	}
	return fmt.Sprintf(" WHERE %s", strings.Join(parts, " AND ")), curArgs, nil
}
