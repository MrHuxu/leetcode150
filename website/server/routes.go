package server

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func (s *server) registerRoutes() {
	s.GET("/:id", func(ctx *gin.Context) {
		solutionBytes, _ := ioutil.ReadFile("../problems/" + ctx.Param("id") + "/solution.go")
		explanationBytes, _ := ioutil.ReadFile("../problems/" + ctx.Param("id") + "/explanation.md")

		ctx.Set("res", map[string]interface{}{
			"title": "LeetCode " + ctx.Param("id") + " - xhu",
			"data": map[string]interface{}{
				"solution":    string(solutionBytes),
				"explanation": string(explanationBytes),
			},
		})
	})

	s.GET("/", func(ctx *gin.Context) {
		ctx.Set("res", map[string]interface{}{
			"title": "LeetCode 150 - xhu",
			"data":  map[string]interface{}{},
		})
	})
}
