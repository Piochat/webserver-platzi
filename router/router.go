package router

import (
	"net/http"
)

type Router struct {
	Rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{Rules: make(map[string]map[string]http.HandlerFunc)}
}

func (R *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, methodExist, exist := R.FindHandler(r.URL.Path, r.Method)

	if !exist {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, r)
}

func (r *Router) FindHandler(path, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.Rules[path]
	handler, methodExist := r.Rules[path][method]

	return handler, methodExist, exist
}
