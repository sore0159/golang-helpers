package myweb

import (
	"net/http"
	"strconv"
	"unicode"
)

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

func GetInts(r *http.Request, varNames ...string) (ints []int, ok bool) {
	ints = make([]int, len(varNames))
	for i, name := range varNames {
		varStr := r.FormValue(name)
		varI, err := strconv.Atoi(varStr)
		if err != nil {
			return nil, false
		}
		ints[i] = varI
	}
	return ints, true
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
