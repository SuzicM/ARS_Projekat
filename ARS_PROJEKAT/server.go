package main

import (
	_ "encoding/json"
	"errors"
	_ "fmt"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
	_ "net/http"
)

type Service struct {
	data map[string][]*Config
}

func (ts *Service) UpdateConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
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
	ts.data[id] = append(ts.data[id], rt...)

	renderJSON(w, rt)
}

func (ts Service) getConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	v, ok := ts.data[id]
	if !ok || len(v) == 1 {
		renderJSON(w, v)
	} else {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (ts Service) getConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	v, ok := ts.data[id]
	if !ok || len(v) > 1 {
		renderJSON(w, v)
	} else {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
