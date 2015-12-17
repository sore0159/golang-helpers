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

func (rg *Registry) Create(w http.ResponseWriter, r *http.Request) (nameOk, passOk, dbOk bool) {
	name, password := r.FormValue("username"), r.FormValue("password")
	nameOk, passOk, dbOk = rg.createUser(name, password)
	if !(nameOk && passOk && dbOk) {
		return
	}
	rg.setCookie(name, w)
	return
}

func (rg *Registry) Login(w http.ResponseWriter, r *http.Request) (User, bool) {
	name, password := r.FormValue("username"), r.FormValue("password")
	if !rg.validLogin(name, password) {
		return User(""), false
	}
	rg.setCookie(name, w)
	return User(name), true
}

func (rg *Registry) Logout(w http.ResponseWriter, r *http.Request) {
	rg.unsetCookie(w)
}

func (rg *Registry) IsLoggedIn(w http.ResponseWriter, r *http.Request) (User, bool) {
	return rg.checkCookie(r)
}
