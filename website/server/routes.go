package server

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *server) registerRoutes() {
	s.GET("/:id", func(ctx *gin.Context) {
		solutionBytes, _ := ioutil.ReadFile("../problems/" + ctx.Param("id") + "/solution.go")
		explanationBytes, _ := ioutil.ReadFile("../problems/" + ctx.Param("id") + "/explanation.md")

		id := ctx.Param("id")
		if p, ok := problems[id]; !ok {
			ctx.Status(http.StatusNotFound)
		} else {
			explanation, algorithm := getExplanationAndAlgorithm(explanationBytes)
			ctx.Set("res", map[string]interface{}{
				"title": "LeetCode " + id + ": " + p.Title + " - xhu",
				"data": map[string]interface{}{
					"id":          id,
					"title":       p.Title,
					"slug":        p.Slug,
					"difficulty":  p.Difficulty,
					"solution":    string(solutionBytes),
					"explanation": explanation,
					"algorithm":   algorithm,
				},
			})
		}
	})

	s.GET("/", func(ctx *gin.Context) {
		ctx.Set("res", map[string]interface{}{
			"title": "LeetCode 150 - xhu",
		})
	})
}

func getExplanationAndAlgorithm(explanationBytes []byte) (string, string) {
	arr := strings.Split(string(explanationBytes), "# 解答")
	if len(arr) < 2 {
		return "", ""
	}
	return strings.TrimLeft(arr[0], "# 题意"), arr[1]
}
