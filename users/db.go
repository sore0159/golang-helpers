package users

import (
	"database/sql"
	"fmt"
	"mule/mydb"
	"strings"
)

func GetRegistry() (*Registry, bool) {
	db, ok := mydb.LoadDB(DB_USER, DB_PASS, PADB_NAME)
	if !ok {
		return nil, false
	}
	return &Registry{db: db}, true
}

func (rg *Registry) Close() {
	rg.db.Close()
}

func (rg *Registry) nameFree(name string) (free, ok bool) {
	lowerName := strings.ToLower(name)
	query := "SELECT name FROM userinfo WHERE lower(name) = $1"
	var found string
	err := rg.db.QueryRow(query, lowerName).Scan(&found)
	if err == sql.ErrNoRows {
		return true, true
	} else if err != nil {
		//Log(err)
		return false, false
	}
	return false, true
}

func (rg *Registry) getPW(name string) (password string, ok bool) {
	err := rg.db.QueryRow("SELECT password FROM userinfo WHERE name = $1", name).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", false
		}
		//Log(err)
		return "", false
	}
	return password, true
}

func (rg *Registry) insertUser(name, pw string) bool {
	query := fmt.Sprintf("INSERT INTO userinfo (name, password) VALUES('%s', '%s')", name, pw)
	return mydb.Exec(rg.db, query)
}
