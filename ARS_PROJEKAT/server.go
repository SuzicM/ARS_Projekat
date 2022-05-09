package main

import (
	"errors"
	"github.com/gorilla/mux"
	_ "mime"
	"net/http"
)

type Service struct {
	data map[string][]*Config
}

func (ts *Service) getAllHandler(w http.ResponseWriter, req *http.Request) {
	allTasks := make(map[string][]*Config)
	for s, v := range ts.data {
		allTasks[s] = v
	}

	renderJSON(w, allTasks)
}

func (ts Service) deleteConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	v, ok := ts.data[id]
	if ok && len(v) == 1 {
		delete(ts.data, id)
		renderJSON(w, v)
	} else {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (ts Service) deleteConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	v, ok := ts.data[id]
	if ok && len(v) > 1 {
		delete(ts.data, id)
		renderJSON(w, v)
	} else {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
