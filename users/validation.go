package users

import (
	"mule/myweb"
	"strings"
)

func (rg *Registry) validLogin(userName, password string) (valid bool) {
	if !(myweb.TextValid(userName) && myweb.TextValid(password)) {
		return false
	}
	pass, ok := rg.getPW(userName)
	if !ok {
		return false
	}
	return pass == password
}

func (rg *Registry) createUser(name, password string) (nameOk, passOk, dbOk bool) {
	dbOk = true
	nameOk, passOk = myweb.TextValid(name), myweb.TextValid(password)
	if nameOk {
		lowerName := strings.ToLower(name)
		reserved := []string{"command", "img", "home", "yours", "static", "turn", "admin", "themule", "mule", "login", "logout"}
		for _, test := range reserved {
			if lowerName == test {
				nameOk = false
				break
			}
		}
	}
	if nameOk {
		nameOk, dbOk = rg.nameFree(name)
	}
	if !(nameOk && passOk && dbOk) {
		return
	}
	dbOk = rg.insertUser(name, password)
	return
}
