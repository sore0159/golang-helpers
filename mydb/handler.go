package mydb

import (
	"fmt"
	"strconv"
	"strings"
)

// Modder is a type to make updating even easier!
// It works well with SQLHandler but doesn't need it
type Modder interface {
	GetMods() map[string]string
	TableName() string
	SQLID() []string
}

// A type that impliments Modder can use ModderQ to
// generate InsertQ() for implimenting Updater
func ModderQ(m Modder) string {
	mods := m.GetMods()
	if len(mods) == 0 {
		return ""
	}
	parts := make([]string, len(mods))
	query := "UPDATE " + m.TableName() + " SET "
	var i int
	for key, val := range mods {
		parts[i] = key + " = " + val
		i++
	}
	query += strings.Join(parts, ", ") + " WHERE "
	idstuff := m.SQLID()
	parts = make([]string, len(idstuff)/2)
	for i, _ := range parts {
		parts[i] = fmt.Sprintf("%s = %s", idstuff[2*i], idstuff[(2*i)+1])
	}
	query += strings.Join(parts, " AND ")
	return query
}

// SQLHandler is a helper for caching changes to
// structs for later mass UPDATE queries
// Add the SetX functions to a struct's Setters
// and impliment Modder to call ModderQ for
// UpdateQ to impliment Updater
type SQLHandler struct {
	Mods map[string]string
}

func NewSQLHandler() *SQLHandler {
	return &SQLHandler{
		Mods: map[string]string{},
	}
}

func (s *SQLHandler) GetMods() map[string]string {
	return s.Mods
}

func (s *SQLHandler) Commit() {
	s.Mods = map[string]string{}
}

func (s *SQLHandler) SetInt(key string, val int) {
	s.Mods[key] = strconv.Itoa(val)
}

// SetStr replaces ' in val with '' (per PostGreSQL)
// for safe insertion, and surrounds val with '
func (s *SQLHandler) SetStr(key, val string) {
	s.Mods[key] = fmt.Sprintf("'%s'", strings.Replace(val, "'", "''", -1))
}

func (s *SQLHandler) SetBool(key string, val bool) {
	if val {
		s.Mods[key] = "true"
	} else {
		s.Mods[key] = "false"
	}
}

// I haven't actually looked up or tested if this is
// the proper NULL literal syntax: TODO
func (s *SQLHandler) SetNull(key string) {
	s.Mods[key] = "NULL"
}

// SetEtc is for string representations of SQL data
// that you build yourself
func (s *SQLHandler) SetEtc(key, val string) {
	s.Mods[key] = val
}

/*   TODO
func (s *SQLHandler) SetIntSlice(key string, val []int) {
	str := "'{"
	parts := make([]string, len(val))
	for i, x := range val {
		parts[i] = strconv.Itoa(x)
	}
	str += strings.Join(parts, ",") + "}'"
	s.Mods[key] = str
}

func (s *SQLHandler) SetStrSlice(key string, val []string) {
	str := "'{"
	parts := make([]string, len(val))
	for i, x := range val {
		parts[i] = fmt.Sprintf("\"%s\"", strings.Replace(x, "'", "''", -1))
	}
	str += strings.Join(parts, ",") + "}'"
	s.Mods[key] = str
}
*/
