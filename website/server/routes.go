package server

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *server) registerRoutes() {
	s.GET("/:id", func(ctx *gin.Context) {
		codeBytes, _ := ioutil.ReadFile("../problems/" + getProblemFolderName(ctx.Param("id")) + "/main.go")
		solutionBytes, _ := ioutil.ReadFile("../problems/" + getProblemFolderName(ctx.Param("id")) + "/solution.md")

		id := ctx.Param("id")
		if p, ok := problems[id]; !ok {
			ctx.Status(http.StatusNotFound)
		} else {
			explanation, algorithm := getExplanationAndAlgorithm(solutionBytes)
			ctx.Set("res", map[string]interface{}{
				"title": "LeetCode " + id + ": " + p.Title + " - xhu",
				"data": map[string]interface{}{
					"id":          id,
					"title":       p.Title,
					"slug":        p.Slug,
					"difficulty":  p.Difficulty,
					"code":        getCode(codeBytes),
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

func getProblemFolderName(id string) string {
	return id + "_" + strings.Replace(problems[id].Slug, "-", "_", -1)
}

func getCode(codeBytes []byte) string {
	arr := strings.Split(string(codeBytes), "// code\n")
	if len(arr) < 2 {
		return ""
	}
	return arr[1]
}

func getExplanationAndAlgorithm(solutionBytes []byte) (string, string) {
	arr := strings.Split(string(solutionBytes), "# 解答")
	if len(arr) < 2 {
		return "", ""
	}
	return strings.TrimLeft(arr[0], "# 题意"), arr[1]
}
