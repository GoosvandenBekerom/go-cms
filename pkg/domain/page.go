package domain

type Page struct {
	Path     string
	Template PageTemplate
}

type PageTemplate interface {
	HtmlFilePath() string
}
