package bootstrap

import (
	"github.com/GoosvandenBekerom/go-cms/pkg/api"
	"net/http"
)

const RootPath = "/api"

func Api() {
	http.HandleFunc(RootPath+"/pages", api.HandlePagesRequest)
}
