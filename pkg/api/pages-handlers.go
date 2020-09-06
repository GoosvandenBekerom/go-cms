package api

import (
	"encoding/json"
	"github.com/GoosvandenBekerom/go-cms/pkg/database"
	"github.com/GoosvandenBekerom/go-cms/pkg/pages"
	"github.com/GoosvandenBekerom/go-cms/pkg/validation"
	"io/ioutil"
	"log"
	"net/http"
)

func HandlePagesRequest(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		handleGetPages(writer)
	case http.MethodPost:
		handlePostPages(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleGetPages(response http.ResponseWriter) {
	pagesDb, err := database.GetAllPages()

	if err != nil {
		log.Printf("error while reading pages from database: %s\n", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	pagesJson, err := json.Marshal(pagesDb)

	if err != nil {
		log.Printf("error while json marshalling pages: %s\n", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	_, err = response.Write(pagesJson)

	if err != nil {
		log.Printf("error while writing json to response: %s\n", err)
		response.WriteHeader(http.StatusInternalServerError)
	}
}

func handlePostPages(response http.ResponseWriter, request *http.Request) {
	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Printf("error while reading request body: %s\n", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	var page pages.Page
	err = json.Unmarshal(requestBody, &page)

	if err != nil {
		log.Printf("unable to parse request to pages.Page: %s\n", err)
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte("unable to parse request body to a page"))
		return
	}

	err = validation.Validate(page)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte(err.Error()))
		return
	}

	savedPage, err := database.SaveNewPage(page)

	if err != nil {
		if err, ok := err.(database.UserFriendlyDatabaseError); ok {
			if err.UserFault {
				response.WriteHeader(http.StatusBadRequest)
			} else {
				response.WriteHeader(http.StatusInternalServerError)
				log.Println(err.RootError)
			}
			_, _ = response.Write([]byte(err.Error()))
			return
		}
		_, _ = response.Write([]byte("an unknown error occurred, try again later."))
		return
	}

	responseJson, err := json.Marshal(savedPage)

	if err != nil {
		log.Printf("error while json marshalling page: %s\n", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	_, err = response.Write(responseJson)

	if err != nil {
		log.Printf("error while writing json to response: %s\n", err)
		response.WriteHeader(http.StatusInternalServerError)
	}
}
