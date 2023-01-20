package api

import (
	"log"
	"net/http"

	"github.com/MrHuxu/leetcode150/website/data"
	"github.com/MrHuxu/leetcode150/website/templates"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := templates.GetTemplate()
	if err != nil {
		log.Fatal(err)
	}

	questions, err := data.GetQuestions()
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, map[string]any{
		"page":      "index",
		"title":     "LeetCode 150 - xhu",
		"questions": questions,
	})
}
