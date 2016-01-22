package mydb

/*
func CheckNull(test []byte) bool {
	if len(test) == 0 {
		return true
	}
	nullBytes := []byte("NULL")
	if len(test) != len(nullBytes) {
		return false
	}
	for i, b := range test {
		if nullBytes[i] != b {
			return false
		}
	}
	return true
}
*/

/* TODO
import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

func NewIntSlice() *IntSlice {
	var x IntSlice = []int{}
	return &x
}

// Scan impliments Scanner for use in SQL queries
func (i *IntSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("bad value scanned to intslice:", value))
	}
	str := string(bytes)
	if str == "{}" {
		*i = IntSlice([]int{})
		return nil
	}
	parts := strings.Split(str, "\"")
	stuff := []int{}
	for _, thing := range parts {
		if thing != "" && thing != "{" && thing != "}" && thing != "," {
			x, err := strconv.Atoi(thing)
			if err != nil {
				return errors.New(fmt.Sprintf("couldn't inslice scan thing:", thing))
			}
			stuff = append(stuff, x)
		}
	}
	*i = stuff
	return nil
}

func (i *IntSlice) Read() []int {
	return []int(*i)
}

type StrSlice []string

func NewStrSlice() *StrSlice {
	var x StrSlice = []string{}
	return &x
}

// Scan impliments Scanner for use in SQL queries
func (s *StrSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("bad value scanned to strslice:", value))
	}
	str := string(bytes)
	if len(str) < 2 {
		return errors.New("strslice scan got too short str: " + str)
	}
	str = str[1 : len(str)-1]
	if str == "{}" {
		*s = StrSlice([]string{})
		return nil
	}
	parts := strings.Split(str, "\"")
	stuff := []string{}
	for _, thing := range parts {
		if thing != "" && thing != "{" && thing != "}" && thing != "," {
			stuff = append(stuff, thing)
		}
	}
	*s = stuff
	return nil
}

func (s *StrSlice) Read() []string {
	return []string(*s)
}
*/
