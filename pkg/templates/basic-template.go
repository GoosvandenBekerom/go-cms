package templates

type BasicTemplate struct {
	Title   string
	Content string
}

func (_ BasicTemplate) HtmlFilePath() string {
	return "basic.html"
}
