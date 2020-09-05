package templates

type InfoTemplate struct {
	Title   string
	Source  string
	Content string
}

func (_ InfoTemplate) HtmlFilePath() string {
	return "info.html"
}
