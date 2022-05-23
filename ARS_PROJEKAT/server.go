package main

import (
	"net/http"
	"errors"
	"mime"
	ps "github.com/SuzicM/ARS_Projekat/ARS_PROJEKAT/poststore"
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

	ts.store.AddConfig(rt)

	renderJSON(w, rt)
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
	
	ts.store.AddConfigGroup(rt)

	renderJSON(w, rt)
}

func (ts *postStore) getAllHandler(w http.ResponseWriter, req *http.Request) {
	allTasks, err1 :=  ts.store.GetAllConfigs()
	allGroups, err2 := ts.store.GetAllGroups()
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	renderJSON(w, allTasks)
	renderJSON(w, allGroups)
}