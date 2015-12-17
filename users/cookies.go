package users

import (
	"net/http"
)

func (rg *Registry) setCookie(username string, w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "username",
		Value:    username,
		MaxAge:   30000000,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}

func (rg *Registry) unsetCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "username",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func (rg *Registry) checkCookie(r *http.Request) (User, bool) {
	cookie, err := r.Cookie("username")
	if err == http.ErrNoCookie {
		return User(""), false
	} else if err != nil {
		// Log(err)
		return User(""), false
	}
	return User(cookie.Value), true
}
