package data

import (
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Questions ...
var Questions QuestionList

func init() {
	bs, err := ioutil.ReadFile("data/data.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(bs, &Questions); err != nil {
		log.Fatal(err)
	}
}

// Question ...
type Question struct {
	ID         int    `json:"id"`
	Difficulty int    `json:"difficulty"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Lang       string `json:"lang"`

	Code     string
	Solution template.HTML
}

// FolderName ...
func (q Question) FolderName() string {
	return strconv.Itoa(q.ID) + "_" + strings.Replace(q.Slug, "-", "_", -1)
}

// QuestionList ...
type QuestionList []Question

// FindByID ...
func (ql QuestionList) FindByID(id int) (Question, error) {
	for _, q := range ql {
		if q.ID == id {
			return q, nil
		}
	}

	return Question{}, errors.New("question not found")
}
