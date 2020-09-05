package main

import (
	"fmt"
	"github.com/GoosvandenBekerom/go-cms/pkg/domain"
	"github.com/GoosvandenBekerom/go-cms/pkg/templates"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	handlePath("/", domain.Page{Template: templates.BasicTemplate{
		Title:   "Homepage",
		Content: "This is the homepage using the basic template",
	}})
	handlePath("/test", domain.Page{Template: templates.BasicTemplate{
		Title:   "Test Page",
		Content: "This is a test page using the basic template",
	}})
	log.Println("Starting webserver on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePath(path string, page domain.Page) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		t, parseError := template.ParseFiles(getHtmlTemplate(page.Template.HtmlFilePath()))

		if parseError != nil {
			log.Println(parseError)
		}

		execError := t.Execute(w, page.Template)

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
