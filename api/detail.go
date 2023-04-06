package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MrHuxu/leetcode150/data"
	"github.com/MrHuxu/leetcode150/templates"
)

// Detail ...
func Detail(w http.ResponseWriter, r *http.Request) {
	tmpl, err := templates.GetTemplate()
	if err != nil {
		log.Fatal(err)
	}

	var question *data.Question

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			tmpl.Execute(w, map[string]any{
				"page": "error",
				"msg":  err.Error(),
			})
			return
		}

		tmpl.Execute(w, map[string]any{
			"page":     "detail",
			"title":    fmt.Sprintf("%d. %s - xhu", question.ID, question.Title),
			"question": question,
		})
	}()

	idParams := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return
	}
	question, err = data.GetQuestion(id)
	if err != nil {
		return
	}
}
