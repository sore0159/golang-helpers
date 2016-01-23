package myweb

import (
	"errors"
	"fmt"
	"html/template"
)

func MakeMixer(tpdir string, fMap map[string]interface{}) func(...string) *template.Template {
	return func(fileNames ...string) *template.Template {
		names := make([]string, len(fileNames))
		for i, val := range fileNames {
			names[i] = tpdir + val + ".html"
		}
		return template.Must(template.New("").Funcs(template.FuncMap(fMap)).ParseFiles(names...))
	}
}

func TemplateDict(val ...interface{}) (map[string]interface{}, error) {
	if len(val)%2 != 0 {
		err := errors.New("template dict needs even args")
		return nil, err
	}
	d := make(map[string]interface{}, len(val)/2)
	for i := 0; i < len(val); i += 2 {
		str, ok := val[i].(string)
		if !ok {
			err := fmt.Errorf("bad template dict arg", val[i], ": need string keys")
			return nil, err
		}
		d[str] = val[i+1]
	}
	return d, nil
}
