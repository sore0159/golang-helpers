package myweb

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

var Log = log.Println

func SetLogger(f func(...interface{})) {
	Log = f
}

var TextValid = MakeValidator(15)

func MakeValidator(maxLen int) func(string) bool {
	return func(test string) bool {
		if test == "" || len(test) > maxLen {
			return false
		}
		for _, rn := range test {
			if !unicode.In(rn, unicode.L, unicode.N) {
				return false
			}
		}
		return true
	}
}

func GetInts(r *http.Request, varNames ...string) (ints map[string]int, ok bool) {
	m := make(map[string]int, len(varNames))
	badNames := []string{}
	for _, name := range varNames {
		varStr := r.FormValue(name)
		varI, err := strconv.Atoi(varStr)
		if err != nil {
			badNames = append(badNames, fmt.Sprintf("%s:%s", name, err))
		} else {
			m[name] = varI
		}
	}
	if len(badNames) == 0 {
		return m, true
	}
	Log("GetInts errors:", strings.Join(badNames, "||"))
	return nil, false
}

func GetIntsIf(r *http.Request, varNames ...string) (ints map[string]int) {
	m := make(map[string]int, len(varNames))
	for _, name := range varNames {
		varStr := r.FormValue(name)
		varI, err := strconv.Atoi(varStr)
		if err == nil {
			m[name] = varI
		}
	}
	return m
}

func GetIntsQuiet(r *http.Request, varNames ...string) (ints map[string]int, ok bool) {
	m := make(map[string]int, len(varNames))
	badNames := []string{}
	for _, name := range varNames {
		varStr := r.FormValue(name)
		varI, err := strconv.Atoi(varStr)
		if err != nil {
			badNames = append(badNames, fmt.Sprintf("%s:%s", name, err))
		} else {
			m[name] = varI
		}
	}
	if len(badNames) == 0 {
		return m, true
	}
	return nil, false
}
