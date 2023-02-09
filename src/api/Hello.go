package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Message struct {
	Message string `json:"message"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("name")

	message := Message{
		fmt.Sprintf("Hello world %s!", name),
	}

	response, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
