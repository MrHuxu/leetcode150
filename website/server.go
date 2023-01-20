package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/MrHuxu/leetcode150/website/data"
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
		"questions": data.Questions,
	})
}

func serveDetail(ctx *gin.Context) {
	var q data.Question
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

	q, err = data.Questions.FindByID(id)
	if err != nil {
		return
	}
}

func getDisplayLang(lang string) string {
	return strings.ToUpper(lang[:1]) + lang[1:]
}
