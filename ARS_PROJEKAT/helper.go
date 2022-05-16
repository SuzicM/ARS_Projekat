package main

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/google/uuid"
)


type ConfigGroup struct {
	id string 
	Version     string   `json:"version"`
	Group []Config `json:"group"`
}

func decodeBodyGroup(r io.Reader) ([]*Config, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt []*Config
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return rt, nil
}

func decodeConfigGroup(r io.Reader) ([]*ConfigGroup, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt []*ConfigGroup
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return rt, nil
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func createId() string {
	return uuid.New().String()
}
