package main

import (
	"net/http"
	"errors"
	"mime"
	ps "SuzicM/ARS_PROJEKAT/poststore"
	"github.com/gorilla/mux"
)

type postStore struct {
	store *ps.PostStore
}

func (ts *postStore) addConfigHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeBodyGroup(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	config , err := ts.store.AddConfig(rt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	renderJSON(w, config)
}

func (ts *postStore) addConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeConfigGroup(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, config := range rt.Group {
		if config.Version != rt.Version{
			config.Version=rt.Version
		}
	}
	
	configgroup , err := ts.store.AddConfigGroup(rt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	renderJSON(w, configgroup)
}

func (ts *postStore) getAllConfigsHandler(w http.ResponseWriter, req *http.Request) {
	allTasks, err :=  ts.store.GetAllConfigs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	renderJSON(w, allTasks)
}

func (ts *postStore) getAllGroupsHandler(w http.ResponseWriter, req *http.Request) {
	allTasks, err :=  ts.store.GetAllGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	renderJSON(w, allTasks)
}

func (ts *postStore) deleteConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	config, err := ts.store.DeleteConfig(id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, config)
}

func (ts *postStore) deleteConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	group, err := ts.store.DeleteConfigGroup(id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, group)
}

func (ts *postStore) getConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	config, err := ts.store.GetConfig(id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, config)
}

func (ts *postStore) getConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	group, err := ts.store.GetConfigGroup(id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, group)
}

func (ts *postStore) UpdateConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeBodyGroup(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := mux.Vars(req)
	id := v["id"]
	version := v["version"]

	configgroup, err := ts.store.UpdateConfigGroup(id,version,rt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	renderJSON(w, configgroup)
}

func (cs *postStore) getConfigByLabels(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	labels := mux.Vars(req)["labels"]
	group, err := cs.store.GetConfigFromGroupWithLabel(id, version, labels)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	renderJSON(w, group)
}