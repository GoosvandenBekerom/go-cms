package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoosvandenBekerom/go-cms/pkg/config"
	. "github.com/GoosvandenBekerom/go-cms/pkg/pages"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func connect() *pgxpool.Pool {
	dbConfig := config.Get().Database

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Dbname)

	connectionPool, err := pgxpool.Connect(context.Background(), connectionString)

	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err)
	}

	return connectionPool
}

func GetAllPages() ([]Page, error) {
	db := connect()
	defer db.Close()

	rows, err := db.Query(context.Background(), "select id, path, template_type, content from pages")

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while querying for pages: %s", err))
	}

	var pages []Page

	for rows.Next() {
		var id *int
		var path string
		var templateType TemplateType
		var jsonContent json.RawMessage
		err = rows.Scan(&id, &path, &templateType, &jsonContent)

		if err != nil {
			return nil, errors.New(fmt.Sprintf("error while scanning page row: %s", err))
		}

		content, err := templateType.ParseJsonContent(jsonContent)

		if err != nil {
			return nil, errors.New(fmt.Sprintf("error while parsing json content: %s", err))
		}

		pages = append(pages, Page{
			Id:           id,
			Path:         path,
			TemplateType: templateType,
			Template:     content,
		})
	}

	return pages, nil
}

func SaveNewPage(page Page) (*Page, error) {
	db := connect()
	defer db.Close()

	var id *int
	var path string
	var templateType TemplateType
	var jsonContent json.RawMessage
	err := db.QueryRow(context.Background(),
		"insert into pages (path, template_type, content) values ($1,$2,$3) returning id, path, template_type, content",
		page.Path, page.TemplateType, page.Template).Scan(&id, &path, &templateType, &jsonContent)

	if err != nil {
		if err, ok := err.(*pgconn.PgError); ok && err.Code == pgerrcode.UniqueViolation {
			if err.ConstraintName == "pages_path_uindex" {
				return nil, UserFriendlyDatabaseError{
					ErrorType:      UniqueConstraintViolation,
					FieldName:      "path",
					AttemptedValue: page.Path,
					UserFault:      true,
					RootError:      err,
				}
			}
		}
		return nil, UserFriendlyDatabaseError{
			UserFault: false,
			RootError: err,
		}
	}

	content, err := templateType.ParseJsonContent(jsonContent)

	if err != nil {
		return nil, UserFriendlyDatabaseError{
			UserFault: false,
			RootError: err,
		}
	}

	return &Page{
		Id:           id,
		Path:         path,
		TemplateType: templateType,
		Template:     content,
	}, nil
}
