package main

import (
	"log"
	"net/http"

	"github.com/arthurvicencio/nestlink/pkg/api"
)

func main() {
	api.Register()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
