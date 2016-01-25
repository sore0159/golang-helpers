package users

import (
	"database/sql"
	"mule/mydb"
	"strings"
)

func GetRegistry() (*Registry, error) {
	db, err := mydb.LoadDB(DB_USER, DB_PASS, PADB_NAME)
	if my, bad := Check(err, "get registry failure"); bad {
		return nil, my
	}
	return &Registry{db: db}, nil
}

func (rg *Registry) Close() {
	rg.db.Close()
}

func (rg *Registry) nameFree(name string) (free bool, err error) {
	lowerName := strings.ToLower(name)
	query := "SELECT name FROM userinfo WHERE lower(name) = $1"
	var found string
	err = rg.db.QueryRow(query, lowerName).Scan(&found)
	if err == sql.ErrNoRows {
		return true, nil
	} else if my, bad := Check(err, "namefree failure"); bad {
		return false, my
	}
	return false, nil
}

func (rg *Registry) getPW(name string) (password string, ok bool, err error) {
	err = rg.db.QueryRow("SELECT password FROM userinfo WHERE name = $1", name).Scan(&password)
	if err == sql.ErrNoRows {
		return "", false, nil
	} else if my, bad := Check(err, "get password failure", "name", name); bad {
		return "", false, my
	}
	return password, true, nil
}

func (rg *Registry) insertUser(name, pw string) error {
	query := "INSERT INTO userinfo (name, password) VALUES($1, $2)"
	return mydb.ExecCheck(rg.db.Exec(query, name, pw))
}
