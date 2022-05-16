package main

import (
	"net/http"
	"errors"
	"mime"
)

type Service struct {
	data map[string][]*Config 
	group map[string][]*ConfigGroup
}

func (ts *Service) addConfigHandler(w http.ResponseWriter, req *http.Request) {
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

	id := createId()
	ts.data[id] = rt

	renderJSON(w, rt)
}

func (ts *Service) addConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
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

	//rt, err := decodeBodyGroup(req.Body)
	rt, err := decodeConfigGroup(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	id := createId()
	ts.group[id] = rt

	renderJSON(w, rt)
}

func (ts *Service) getAllHandler(w http.ResponseWriter, req *http.Request) {
	allTasks :=  make(map[string][]*Config)
	allGroups := make(map[string][]*ConfigGroup)
	for s, v := range ts.data {
		allTasks[s]= v
	}
	for s, v := range ts.group {
		allGroups[s]= v
	}

	renderJSON(w, allTasks)
	renderJSON(w, allGroups)
}