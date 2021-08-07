package main

import (
	"log"

	"github.com/astrokir/http_rest_api/internal/app/apiserver"
)

func main() {
	server := apiserver.NewAPIServer()

	if err := server.StartAPIServer(); err != nil {
		log.Fatal(err)
	}
}
