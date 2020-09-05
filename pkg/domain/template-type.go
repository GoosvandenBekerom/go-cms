package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoosvandenBekerom/go-cms/pkg/domain/templates"
	"strings"
)

type TemplateType int

const (
	BASIC TemplateType = iota
	INFO  TemplateType = iota
)

// get string representation of "enum" value
func (t TemplateType) String() string {
	return [...]string{
		"BASIC",
		"INFO",
	}[t]
}

// find the correct template for this TemplateType and parse the raw json content into it
func (t TemplateType) ParseJsonContent(raw json.RawMessage) (PageTemplate, error) {
	switch t {
	case BASIC:
		var template templates.BasicTemplate
		err := json.Unmarshal(raw, &template)
		return template, err
	case INFO:
		var template templates.InfoTemplate
		err := json.Unmarshal(raw, &template)
		return template, err
	default:
		return nil, errors.New(fmt.Sprintf("unknown type received in ParseJsonContent: %s", t))
	}
}

// When reading TemplateType from the database parse the string value to its int representation
func (t *TemplateType) Scan(value interface{}) error {
	stringValue := value.(string)

	switch strings.ToUpper(stringValue) {
	case "BASIC":
		*t = BASIC
	case "INFO":
		*t = INFO
	default:
		return errors.New(fmt.Sprintf("unknown type received: %s", stringValue))
	}

	return nil
}

// When writing TemplateType to the database, use its string representation
func (t TemplateType) Value() (driver.Value, error) { return t.String(), nil }
