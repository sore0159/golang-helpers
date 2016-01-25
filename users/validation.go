package users

import (
	"mule/myweb"
	"strings"
)

func (rg *Registry) validLogin(userName, password string) (valid bool, err error) {
	if !(myweb.TextValid(userName) && myweb.TextValid(password)) {
		return false, nil
	}
	pass, ok, err := rg.getPW(userName)
	if my, bad := Check(err, "valid login check failure", "username", userName, "password", password); bad {
		return false, my
	}
	if !ok {
		return false, nil
	}
	return pass == password, nil
}

func (rg *Registry) createUser(name, password string) (nameOk, passOk bool, err error) {
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
		nameOk, err = rg.nameFree(name)
		if my, bad := Check(err, "create user failure"); bad {
			return false, false, my
		}
	}
	if !(nameOk && passOk) {
		return nameOk, passOk, nil
	}
	err = rg.insertUser(name, password)
	if my, bad := Check(err, "create user failure"); bad {
		return false, false, my
	}
	return true, true, nil
}
