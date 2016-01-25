package mydb

import (
	"fmt"
)

type SQLer interface {
	SQLTable() string
	SQLVal(name string) interface{}
	SQLPtr(name string) interface{}
}

func ColVals(cl SQLer, colNames []string) (valList []interface{}, err error) {
	valList = make([]interface{}, len(colNames))
	for i, name := range colNames {
		val := cl.SQLVal(name)
		if val == nil {
			return nil, fmt.Errorf("SQLColumner type %s asked for bad val %s", cl.SQLTable(), name)
		}
		valList[i] = val
	}
	return valList, nil
}

func ColPtrs(cl SQLer, colNames []string) (ptrList []interface{}, err error) {
	ptrList = make([]interface{}, len(colNames))
	for i, name := range colNames {
		ptr := cl.SQLPtr(name)
		if ptr == nil {
			return nil, fmt.Errorf("SQLColumner type %s asked for bad ptr %s", cl.SQLTable(), name)
		}
		ptrList[i] = ptr
	}
	return ptrList, nil
}
