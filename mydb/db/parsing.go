package db

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

// Wrapper types to handle parsing db values/scans for go slices.
type (
	StringList []string
	IntList    []int
)

// I have no idea how robust this is, but it works for now
// NULL, for example, will parse the same as "NULL" -_-
func SqlArrayToParts(array string) []string {
	if array == "{}" {
		return []string{}
	}
	array = strings.Trim(array, "{}")
	var inQuotes, escaped bool
	var count int
	parts := []string{""}
	for _, char := range array {
		if escaped {
			if char == '"' || char == '\\' {
				parts[count] = fmt.Sprintf("%s%c", parts[count], char)
			} else {
				parts[count] = fmt.Sprintf("%s\\%c", parts[count], char)
			}
			escaped = false
		} else if char == '\\' {
			escaped = true
		} else if char == '"' {
			inQuotes = !inQuotes
		} else if !inQuotes && char == ',' {
			count += 1
			parts = append(parts, "")
		} else {
			parts[count] = fmt.Sprintf("%s%c", parts[count], char)
		}
	}
	return parts
}

func NewStringList() *StringList {
	sl := StringList([]string{})
	return &sl
}

func (sl *StringList) Scan(value interface{}) error {
	if value == nil {
		*sl = nil
		return nil
	}
	// A real party bunch up there
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Bad value scanned to stringlist:", value)
	}
	*sl = SqlArrayToParts(string(bytes))
	return nil
}
func (sl StringList) Value() (driver.Value, error) {
	if len(sl) == 0 {
		return "{}", nil
	}
	parts := make([]string, len(sl))
	for i, pt := range sl {
		str := strings.Replace(pt, `\`, `\\`, -1)
		str = strings.Replace(str, "\"", `\"`, -1)
		parts[i] = fmt.Sprintf("\"%s\"", str)
	}
	str := fmt.Sprintf("{%s}", strings.Join(parts, ",")) //, nil
	return str, nil
}

func NewIntList() *IntList {
	il := IntList([]int{})
	return &il
}

func (il *IntList) Scan(value interface{}) error {
	if value == nil {
		*il = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Bad value scanned to intlist: %v", value)
	}
	valStr := string(bytes)
	if valStr == "{}" {
		*il = []int{}
		return nil
	}
	valStr = strings.Trim(valStr, "{}")
	parts := strings.Split(valStr, ",")
	res := make([]int, len(parts))
	for i, xStr := range parts {
		x, err := strconv.Atoi(xStr)
		if err != nil {
			return fmt.Errorf("Bad value at index %d for scanned to intlist: %v", i, value)
		}
		res[i] = x
	}
	*il = res
	return nil
}

func (il IntList) Value() (driver.Value, error) {
	if len(il) == 0 {
		return "{}", nil
	}
	parts := make([]string, len(il))
	for i, pt := range il {
		parts[i] = strconv.Itoa(pt)
	}
	str := fmt.Sprintf("{%s}", strings.Join(parts, ",")) //, nil
	return str, nil
}
