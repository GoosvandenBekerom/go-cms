package database

import (
	"context"
	"encoding/json"
	"fmt"
	. "github.com/GoosvandenBekerom/go-cms/pkg/domain"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

const (
	host     string = "localhost"
	port     int    = 5432
	username string = "cms"
	password string = "cms"
	dbname   string = "cms"
)

func connect() *pgxpool.Pool {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)

	connectionPool, err := pgxpool.Connect(context.Background(), connectionString)

	if err != nil {
		log.Fatalf("Unable to connect to databse: %s", err)
	}

	return connectionPool
}

func GetAllPages() []Page {
	db := connect()
	defer db.Close()

	rows, err := db.Query(context.Background(), "select path, template_type, content from pages")

	if err != nil {
		log.Fatalf("Error while querying for pages: %s", err)
	}

	var pages []Page

	for rows.Next() {
		var path string
		var templateType TemplateType
		var jsonContent json.RawMessage
		err = rows.Scan(&path, &templateType, &jsonContent)

		if err != nil {
			log.Fatalf("error while scanning page row: %s", err)
		}

		content, err := templateType.ParseJsonContent(jsonContent)

		if err != nil {
			log.Fatalf("error while parsing json content: %s", err)
		}

		pages = append(pages, Page{
			Path:         path,
			TemplateType: templateType,
			Content:      content,
		})
	}

	return pages
}
