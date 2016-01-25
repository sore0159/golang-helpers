package users

import (
	"database/sql"
	"net/http"
)

// GetRegistry() (*Registry, bool)
type Registry struct {
	db *sql.DB
}

type User string

func (u User) String() string {
	return string(u)
}

func (rg *Registry) Create(w http.ResponseWriter, r *http.Request) (nameOk, passOk bool, err error) {
	name, password := r.FormValue("username"), r.FormValue("password")
	nameOk, passOk, err = rg.createUser(name, password)
	if !(nameOk && passOk && err == nil) {
		return
	}
	rg.setCookie(name, w)
	return
}

func (rg *Registry) Login(w http.ResponseWriter, r *http.Request) (User, bool, error) {
	name, password := r.FormValue("username"), r.FormValue("password")
	ok, err := rg.validLogin(name, password)
	if my, bad := Check(err, "Login failure", "name", name, "password", password); bad {
		return User(""), false, my
	}
	if !ok {
		return User(""), false, nil
	}
	rg.setCookie(name, w)
	return User(name), true, nil
}

func (rg *Registry) Logout(w http.ResponseWriter, r *http.Request) {
	rg.unsetCookie(w)
}

func (rg *Registry) IsLoggedIn(w http.ResponseWriter, r *http.Request) (User, bool, error) {
	return rg.checkCookie(r)
}
