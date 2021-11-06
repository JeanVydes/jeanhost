package main

import (
	"log"

	httpserver "github.com/JeanVydes/apitestproject/pgk/httpserver"
)

func main() {
	log.Println("Server started")
	httpserver.Start()
}