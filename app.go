package main

import (
	"log"

	httpserver "github.com/JeanVydes/apitestproject/pgk/httpserver"
)

func main() {
	httpserver.Start()
	log.Println("Server started")
}