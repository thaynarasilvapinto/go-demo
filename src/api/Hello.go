package api

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func HelloWorld(w http.ResponseWriter, rq *http.Request) {

	message := Message{"Hello World."}
	resp, _ := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
