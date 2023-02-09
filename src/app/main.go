package main

import (
	r "go-demo/src/router"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func main() {
	r.ApiRouter()
}
