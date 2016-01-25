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

func (rg *Registry) checkCookie(r *http.Request) (User, bool, error) {
	cookie, err := r.Cookie("username")
	if my, bad := Check(err, "check cookie failure"); bad {
		if my.BaseIs(http.ErrNoCookie) {
			return User(""), false, nil
		}
		return User(""), false, my
	}
	return User(cookie.Value), true, nil
}
