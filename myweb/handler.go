package myweb

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	Path []string
	App  interface{}
}

func NewHandler() *Handler {
	return &Handler{}
}

func MakeHandler(r *http.Request) *Handler {
	return &Handler{
		Path: strings.Split(r.URL.Path, "/"),
	}
}

func (h *Handler) Apply(w http.ResponseWriter, t *template.Template, tName string, app interface{}) {
	err := t.ExecuteTemplate(w, tName, app)
	if err != nil {
		Log("template apply error:", err)
		return
	}
}

func (h *Handler) SetApp(a interface{}) {
	h.App = a
}

func (h *Handler) DefaultApp() map[string]interface{} {
	m := map[string]interface{}{}
	h.App = m
	return m
}

func (v *Handler) LastFull() int {
	l := len(v.Path) - 1
	for ; l > 0; l-- {
		if v.Path[l] != "" {
			break
		}
	}
	return l
}

func (v *Handler) NewPath(n int) string {
	if len(v.Path) < n-1 {
		return strings.Join(v.Path, "/")
	}
	return strings.Join(v.Path[:n], "/")
}

func (v *Handler) Link(str string) string {
	r := []string{""}
	for _, part := range append(v.Path, strings.Split(str, "/")...) {
		if part == "" {
			continue
		} else if part == ".." {
			if len(r) > 1 {
				r = r[:len(r)-1]
			}
		} else {
			r = append(r, part)
		}
	}
	r = append(r, "")
	return strings.Join(r, "/")
}

func (v *Handler) IntAt(n int) (int, bool) {
	if n < 0 || len(v.Path)-1 < n {
		return 0, false
	}
	x, err := strconv.Atoi(v.Path[n])
	if err != nil {
		return 0, false
	}
	return x, true
}
