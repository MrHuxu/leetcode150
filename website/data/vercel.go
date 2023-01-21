package data

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"os"

	"github.com/russross/blackfriday/v2"
)

//go:embed data.json
var dataFS embed.FS

//go:embed documents
var documentsFS embed.FS

//go:embed solutions
var solutionsFS embed.FS

// GetQuestions ...
func GetQuestions() (QuestionList, error) {
	questions, err := loadQuestions()
	if err != nil {
		return nil, err
	}

	for i := range questions {
		if _, err := solutionsFS.Open(fmt.Sprintf("solutions/go/%s/main.go", questions[i].goFolderName())); err == nil {
			questions[i].Langs = append(questions[i].Langs, langGo)
		}
		if _, err := solutionsFS.Open(fmt.Sprintf("solutions/rust/src/%s", questions[i].rustFileName())); err == nil {
			questions[i].Langs = append(questions[i].Langs, langRust)
		}
		if _, err := solutionsFS.Open(fmt.Sprintf("solutions/java/src/main/java/%s", questions[i].javaFileName())); err == nil {
			questions[i].Langs = append(questions[i].Langs, langJava)
		}
	}

	return questions, err
}

// GetQuestion ...
func GetQuestion(id int) (*Question, error) {
	questions, err := loadQuestions()
	if err != nil {
		return nil, err
	}
	question, err := questions.FindByID(id)
	if err != nil {
		return nil, err
	}

	if id > 1 {
		question.Prev = &(questions[id-2])
	}
	if id < len(questions) {
		question.Next = &(questions[id])
	}

	bytes, err := documentsFS.ReadFile(fmt.Sprintf("documents/%s.md", question.goFolderName()))
	if err != nil {
		return nil, err
	}
	question.Document = template.HTML(blackfriday.Run(bytes))
	question.Codes = make(map[string]string)

	bytes, err = solutionsFS.ReadFile(fmt.Sprintf("solutions/go/%s/main.go", question.goFolderName()))
	if err == nil {
		question.Langs = append(question.Langs, langGo)
		goCode, err := extractGoCode(string(bytes))
		if err != nil {
			return nil, err
		}
		question.Codes[langGo] = goCode
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	bytes, err = solutionsFS.ReadFile(fmt.Sprintf("solutions/rust/src/%s", question.rustFileName()))
	if err == nil {
		question.Langs = append(question.Langs, langRust)
		rustCode, err := extractRustCode(string(bytes))
		if err != nil {
			return nil, err
		}
		question.Codes[langRust] = rustCode
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	bytes, err = solutionsFS.ReadFile(fmt.Sprintf("solutions/java/src/main/java/%s", question.javaFileName()))
	if err == nil {
		question.Langs = append(question.Langs, langJava)
		javaCode, err := extractJavaCode(string(bytes))
		if err != nil {
			return nil, err
		}
		question.Codes[langJava] = javaCode
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	return question, nil
}

func loadQuestions() (QuestionList, error) {
	bytes, err := dataFS.ReadFile("data.json")
	if err != nil {
		return nil, err
	}

	var questions QuestionList
	if err = json.Unmarshal(bytes, &questions); err != nil {
		return nil, err
	}
	return questions, nil
}
