package pages

import (
	"encoding/json"
)

type Page struct {
	Id   *int
	Path string
	TemplateType
	Template PageTemplate
}

func (p *Page) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id           *int `json:"Id,omitempty"`
		Path         string
		TemplateType string
		Template     PageTemplate
	}{
		Id:           p.Id,
		Path:         p.Path,
		TemplateType: p.TemplateType.String(),
		Template:     p.Template,
	})
}

func (p *Page) UnmarshalJSON(bytes []byte) error {
	placeholder := &struct {
		Id           *int `json:"Id,omitempty"`
		Path         string
		TemplateType string
		Template     json.RawMessage
	}{}
	if err := json.Unmarshal(bytes, placeholder); err != nil {
		return err
	}

	p.Id = placeholder.Id
	p.Path = placeholder.Path

	templateType, err := GetTemplateTypeFromString(placeholder.TemplateType)
	if err != nil {
		return err
	}
	p.TemplateType = templateType

	template, err := templateType.ParseJsonContent(placeholder.Template)
	if err != nil {
		return err
	}
	p.Template = template
	return nil
}
