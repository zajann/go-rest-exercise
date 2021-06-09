package main

import (
	"log"
	"net/http"

	"github.com/zajann/go-rest-exercise/myapp"
)

func main() {
	handler := myapp.NewHandler()
	log.Fatal(http.ListenAndServe(":9999", handler))
}
