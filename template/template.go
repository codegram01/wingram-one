package template

import (
	"bytes"
	"fmt"

	"github.com/google/safehtml/template"
)

type Template struct {
	Templates map[string]*template.Template
}

func executeTemplate(templateName string, tmpl *template.Template, data any) ([]byte, error) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (t *Template) findTemplate(templateName string) (*template.Template, error) {
	tmpl := t.Templates[templateName]
	if tmpl == nil {
		return nil, fmt.Errorf("BUG: s.templates[%q] not found", templateName)
	}
	return tmpl, nil
}
