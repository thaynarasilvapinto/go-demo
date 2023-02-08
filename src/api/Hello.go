package api

import "fmt"

type Message struct {
	Message string `json:"message"`
}

func Teste() {
	fmt.Println("Teste")
}

// func HelloWorld(w http.ResponseWriter, rq *http.Request) {

// 	message := Message{"Hello World."}
// 	resp, _ := json.Marshal(message)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(resp)
// }
