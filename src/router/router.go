package router

import (
	"fmt"
	"net/http"

	app "go-demo/src/api"

	"github.com/gorilla/mux"
)

func ApiRouter() {

	router := mux.NewRouter()
	router.HandleFunc("/hello", app.HelloWorld).Methods("GET")

	fmt.Println("Server is up and running on localhost:8080")
	http.ListenAndServe(":8080", router)
}
