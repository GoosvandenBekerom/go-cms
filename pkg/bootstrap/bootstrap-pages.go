package bootstrap

import (
	"fmt"
	"github.com/GoosvandenBekerom/go-cms/pkg/database"
	"github.com/GoosvandenBekerom/go-cms/pkg/pages"
	"html/template"
	"log"
	"net/http"
	"os"
)

func Pages() {
	dbPages, err := database.GetAllPages()

	if err != nil {
		log.Fatal(err)
	}

	for _, page := range dbPages {
		handle(page)
	}
}

func handle(page pages.Page) {
	http.HandleFunc(page.Path, func(w http.ResponseWriter, r *http.Request) {
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
