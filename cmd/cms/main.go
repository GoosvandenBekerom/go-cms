package main

import (
	"fmt"
	"github.com/GoosvandenBekerom/go-cms/pkg/database"
	"github.com/GoosvandenBekerom/go-cms/pkg/domain"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, page := range database.GetAllPages() {
		handle(page)
	}
	log.Println("Starting webserver on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(page domain.Page) {
	http.HandleFunc(page.Path, func(w http.ResponseWriter, r *http.Request) {
		t, parseError := template.ParseFiles(getHtmlTemplate(page.Content.HtmlFilePath()))

		if parseError != nil {
			log.Println(parseError)
		}

		execError := t.Execute(w, page.Content)

		if execError != nil {
			log.Println(execError)
		}
	})
}

func getHtmlTemplate(file string) string {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s/web/html/%s", wd, file)
}
