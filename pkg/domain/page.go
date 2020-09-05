package domain

type Page struct {
	Template PageTemplate
}

type PageTemplate interface {
	HtmlFilePath() string
}
