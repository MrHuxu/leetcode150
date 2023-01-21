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
		"getDisplayLang": getDisplayLang,
	}).ParseFS(templatesFS, "base.tmpl", "index.tmpl", "detail.tmpl", "error.tmpl")
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func getDisplayLang(lang string) string {
	return strings.ToUpper(lang[:1]) + lang[1:]
}
