package main

import (
	config "go-demo/src/config"
	router "go-demo/src/router"
)

func main() {
	config.Configuration()
	router.ApiRouter()
}
