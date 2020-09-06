package main

import (
	"fmt"
	"github.com/GoosvandenBekerom/go-cms/pkg/bootstrap"
	"github.com/GoosvandenBekerom/go-cms/pkg/config"
	"log"
	"net/http"
)

func main() {
	cfg := config.Get()

	bootstrap.Pages()
	bootstrap.Api()

	log.Printf("Starting webserver on port %d\n", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), nil))
}
