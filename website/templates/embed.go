package templates

import (
	"embed"
	"html/template"
	"strings"
)

//go:embed *
var templatesFS embed.FS

// GetTemplate ...
func GetTemplate() (*template.Template, error) {
	tmpl, err := template.New("base.tmpl").Funcs(template.FuncMap{
		"getDisplayLang": GetDisplayLang,
	}).ParseFS(templatesFS, "base.tmpl", "index.tmpl", "detail.tmpl", "error.tmpl")
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

// GetDisplayLang ...
func GetDisplayLang(lang string) string {
	if lang == "typescript" {
		return "TypeScript"
	}
	return strings.ToUpper(lang[:1]) + lang[1:]
}
