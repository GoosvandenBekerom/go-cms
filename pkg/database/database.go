package database

import (
	. "github.com/GoosvandenBekerom/go-cms/pkg/domain"
	"github.com/GoosvandenBekerom/go-cms/pkg/templates"
)

func GetAllPages() []Page {
	return []Page{
		{
			Path: "/",
			Template: templates.BasicTemplate{
				Title:   "Homepage",
				Content: "This is the homepage using the basic template",
			},
		},
		{
			Path: "/test",
			Template: templates.BasicTemplate{
				Title:   "Test Page",
				Content: "This is a test page using the basic template",
			},
		},
	}
}
