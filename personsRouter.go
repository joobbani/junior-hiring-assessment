package main

import (
	"net/http"
	"strings"
)

// PersonsRouter handles the persons route
func PersonsRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/people" {
		switch r.Method {
		case http.MethodGet:
			personsGetAll(w, r)
			return
		case http.MethodPost:
			personsPostOne(w, r)
			return
		case http.MethodHead:
			personsGetAll(w, r)
			return
		case http.MethodOptions:
			postOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/people/")

	name := path

	switch r.Method {
	case http.MethodGet:
		personsGetOne(w, r, name)
		return
	case http.MethodOptions:
		postOptionsResponse(w, []string{http.MethodGet, http.MethodPut, http.MethodHead, http.MethodOptions}, nil)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
