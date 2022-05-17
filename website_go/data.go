package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/russross/blackfriday/v2"
)

// Questions defines a global variable holding question data list
var Questions questions

type question struct {
	ID         int    `json:"id"`
	Difficulty int    `json:"difficulty"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`

	Code     map[string]string
	Document template.HTML

	Prev *question
	Next *question
}

// questions ...
type questions []question

func initData() {
	bs, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(bs, &Questions); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(Questions); i++ {
		if i < len(Questions)-1 {
			Questions[i].Next = &(Questions[i+1])
		}
		if i > 0 {
			Questions[i].Prev = &(Questions[i-1])
		}

		if err = Questions[i].loadDocument(); err != nil {
			log.Fatal(err)
		}
		if err = Questions[i].loadCode(); err != nil {
			log.Fatal(err)
		}
	}
}

func (q *question) loadDocument() error {
	bytes, err := ioutil.ReadFile(path.Join(documentFilePath, q.goFolderName()+".md"))
	if err != nil {
		return err
	}

	q.Document = template.HTML(blackfriday.Run(bytes))
	return nil
}

func (q *question) loadCode() error {
	q.Code = make(map[string]string)

	if goCode, err := getGoContent(path.Join(goFilePath, q.goFolderName(), "main.go")); err != nil {
		return err
	} else if goCode != "" {
		q.Code["go"] = strings.Trim(goCode, "\n")
	}

	if rustCode, err := getRustContent(path.Join(rustFilePath, "src", q.rustFileName())); err != nil {
		return err
	} else if rustCode != "" {
		q.Code["rust"] = strings.Trim(rustCode, "\n")
	}

	return nil
}

func getGoContent(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return "", nil
	} else if err != nil {
		return "", err
	}

	arr := strings.Split(string(bytes), "// code\n")
	if len(arr) < 2 {
		return "", errors.New("code content is empty")
	}
	return arr[1], nil
}

func getRustContent(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return "", nil
	} else if err != nil {
		return "", err
	}

	arr := strings.Split(string(bytes), "struct Solution;\n\n")
	if len(arr) < 2 {
		return "", errors.New("code content is empty")
	}
	arr = strings.Split(arr[1], "\n#[test]")
	if len(arr) < 2 {
		return "", errors.New("code content is empty")
	}
	return arr[0], nil
}

func (q question) goFolderName() string {
	return strconv.Itoa(q.ID) + "_" + strings.Replace(q.Slug, "-", "_", -1)
}

func (q question) rustFileName() string {
	return "question_" + strconv.Itoa(q.ID) + ".rs"
}

// FindByID ...
func (q questions) FindByID(id int) (question, error) {
	for _, q := range q {
		if q.ID == id {
			return q, nil
		}
	}

	return question{}, errors.New("question not found")
}
