package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/MrHuxu/leetcode150/website/data"

	"github.com/gin-gonic/gin"
	"gopkg.in/russross/blackfriday.v2"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob(templatesPath)

	server.GET("/:id", serveDetail)
	server.GET("/", serveIndex)

	server.Run(fmt.Sprintf(":%d", port))
}

func serveIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "base.tmpl", gin.H(map[string]interface{}{
		"page":      "index",
		"title":     "LeetCode 150 - xhu",
		"questions": data.Questions,
	}))
}

func serveDetail(ctx *gin.Context) {
	var q data.Question
	var err error

	defer func() {
		if err != nil {
			ctx.HTML(http.StatusInternalServerError, "base.tmpl", gin.H(map[string]interface{}{
				"page": "error",
				"msg":  err.Error(),
			}))
			return
		}

		ctx.HTML(http.StatusOK, "base.tmpl", gin.H(map[string]interface{}{
			"page":     "detail",
			"title":    fmt.Sprintf("%d. %s - xhu", q.ID, q.Title),
			"question": q,
		}))
	}()

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}

	q, err = data.Questions.FindByID(id)
	if err != nil {
		return
	}

	if q.Solution, err = getSolution(q); err != nil {
		return
	}

	if q.Code, err = getCode(q); err != nil {
		return
	}
}

func getCode(q data.Question) (string, error) {
	var codeBytes []byte
	var err error
	switch q.Lang {
	case fileTypeGo:
		codeBytes, err = ioutil.ReadFile("../go/" + q.FolderName() + "/main.go")

	case fileTypeJavascript:
		codeBytes, err = ioutil.ReadFile("../go/" + q.FolderName() + "/main.js")

	default:
		err = errors.New("language of question undefined")
	}
	if err != nil {
		return "", err
	}

	arr := strings.Split(string(codeBytes), "// code\n")
	if len(arr) < 2 {
		return "", errors.New("code content is empty")
	}
	return arr[1], nil
}

func getSolution(q data.Question) (template.HTML, error) {
	solutionBytes, err := ioutil.ReadFile("../go/" + q.FolderName() + "/solution.md")
	if err != nil {
		return "", err
	}

	return template.HTML(blackfriday.Run(solutionBytes)), nil
}
