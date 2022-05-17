package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Server defines a global variable holding the server instance
var Server *gin.Engine

func initServer() {
	Server = gin.Default()

	Server.SetFuncMap(template.FuncMap{
		"getDisplayLang": getDisplayLang,
	})
	Server.LoadHTMLGlob(templatesPath)

	Server.Static("/static/assets", assetsPath)

	Server.GET("/", serveIndex)
	Server.GET("/:id", serveDetail)
}

func serveIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "base.tmpl", gin.H{
		"page":      "index",
		"title":     "LeetCode 150 - xhu",
		"questions": Questions,
	})
}

func serveDetail(ctx *gin.Context) {
	var q question
	var err error

	defer func() {
		if err != nil {
			ctx.HTML(http.StatusInternalServerError, "base.tmpl", gin.H{
				"page": "error",
				"msg":  err.Error(),
			})
			return
		}

		ctx.HTML(http.StatusOK, "base.tmpl", gin.H{
			"page":     "detail",
			"title":    fmt.Sprintf("%d. %s - xhu", q.ID, q.Title),
			"question": q,
		})
	}()

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}

	q, err = Questions.FindByID(id)
	if err != nil {
		return
	}
}

func getDisplayLang(lang string) string {
	return strings.ToUpper(lang[:1]) + lang[1:]
}
