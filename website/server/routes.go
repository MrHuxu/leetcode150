package server

import (
	"github.com/gin-gonic/gin"
)

func (s *server) registerRoutes() {
	s.GET("/:id", func(ctx *gin.Context) {
		ctx.Set("res", map[string]interface{}{
			"title": "LeetCode " + ctx.Param("id") + " - xhu",
			"data":  map[string]interface{}{},
		})
	})

	s.GET("/", func(ctx *gin.Context) {
		ctx.Set("res", map[string]interface{}{
			"title": "LeetCode 150 - xhu",
			"data":  map[string]interface{}{},
		})
	})
}
