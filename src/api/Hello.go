package api

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(Message{"Hello world!"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
