package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MrHuxu/leetcode150/website/data"
	"github.com/MrHuxu/leetcode150/website/templates"
)

// Detail ...
func Detail(w http.ResponseWriter, r *http.Request) {
	tmpl, err := templates.GetTemplate()
	if err != nil {
		log.Fatal(err)
	}

	idParams := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		log.Fatal(err)
	}
	question, err := data.GetQuestion(id)
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, map[string]any{
		"page":     "detail",
		"title":    fmt.Sprintf("%d. %s - xhu", question.ID, question.Title),
		"question": question,
	})
}
