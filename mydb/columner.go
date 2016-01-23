package mydb

import (
	"fmt"
)

type SQLColumner interface {
	SQLType() string
	SQLColumn(name string, pointer bool) (interface{}, bool)
}

func ColVals(cl SQLColumner, colNames []string) (valList []interface{}, err error) {
	valList = make([]interface{}, len(colNames))
	for i, name := range colNames {
		val, ok := cl.SQLColumn(name, false)
		if !ok {
			return nil, fmt.Errorf("SQLColumner type %s asked for bad val %s", cl.SQLType(), name)
		}
		valList[i] = val
	}
	return valList, nil
}

func ColPtrs(cl SQLColumner, colNames []string) (ptrList []interface{}, err error) {
	ptrList = make([]interface{}, len(colNames))
	for i, name := range colNames {
		ptr, ok := cl.SQLColumn(name, true)
		if !ok {
			return nil, fmt.Errorf("SQLColumner type %s asked for bad ptr %s", cl.SQLType(), name)
		}
		ptrList[i] = ptr
	}
	return ptrList, nil
}
