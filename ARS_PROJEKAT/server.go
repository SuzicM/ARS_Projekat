package main

import (
	"errors"
	"github.com/gorilla/mux"
	_ "mime"
	"net/http"
)

type Service struct {
	data  map[string][]*Config
	group map[string][]*ConfigGroup
}

func (ts *Service) getAllHandler(w http.ResponseWriter, req *http.Request) {
	allTasks := make(map[string][]*Config)
	allGroups := make(map[string][]*ConfigGroup)
	for s, v := range ts.data {
		allTasks[s] = v
	}
	for s, v := range ts.group {
		allGroups[s] = v
	}

	renderJSON(w, allTasks)
	renderJSON(w, allGroups)
}

func (ts Service) deleteConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	v, ok := ts.data[id]
	for _, s := range v {
		if ok && len(v) == 1 && s.Version == version {
			delete(ts.data, id)
			renderJSON(w, v)
		} else {
			err := errors.New("key not found")
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}
}

func (ts Service) deleteConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
	version := mux.Vars(req)["version"]
	id := mux.Vars(req)["id"]
	v, ok := ts.group[id]
	for _, s := range v {
		if !ok || len(s.Group) > 1 {
			if version == s.Version {
				delete(ts.group, id)
				renderJSON(w, v)
			}
		} else {
			err := errors.New("key not found")
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}
}
