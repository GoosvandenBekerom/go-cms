package main

import (
	"github.com/GoosvandenBekerom/go-cms/pkg/bootstrap"
	"log"
	"net/http"
)

func main() {
	bootstrap.Pages()
	bootstrap.Api()
	log.Println("Starting webserver on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
