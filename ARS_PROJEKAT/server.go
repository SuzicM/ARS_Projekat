package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "fmt"
	"net/http"
	_ "net/http"
)

type Service struct {
	data map[string][]*Config
}

func (ts *Service) UpdateConfig(response http.ResponseWriter, request *http.Request) {
	var configStruct Config
	err := json.NewDecoder(request.Body).Decode(&configStruct)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		fmt.Println("Product Info - Updated")
		respondWithJSON(response, http.StatusOK, configStruct)
	}
}

func respondWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func respondWithError(response http.ResponseWriter, statusCode int, msg string) {
	respondWithJSON(response, statusCode, map[string]string{"error": msg})
}
