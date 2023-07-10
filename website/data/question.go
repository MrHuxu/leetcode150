package data

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
var Questions QuestionList

// Question ...
type Question struct {
	ID         int    `json:"id"`
	Difficulty int    `json:"difficulty"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`

	Langs    []string
	Codes    map[string]string
	Document template.HTML

	Prev *Question
	Next *Question
}

// QuestionList ...
type QuestionList []Question

// Init ...
func Init() {
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

		if err := Questions[i].loadDocument(); err != nil {
			log.Fatal(err)
		}
		if err := Questions[i].loadCode(); err != nil {
			log.Fatal(err)
		}
	}
}

func (q *Question) loadDocument() error {
	bytes, err := ioutil.ReadFile(path.Join(documentDirectory, q.goFolderName()+".md"))
	if err != nil {
		return err
	}

	q.Document = template.HTML(blackfriday.Run(bytes))
	return nil
}

func (q *Question) loadCode() error {
	q.Codes = make(map[string]string)

	if goCode, err := getGoContent(path.Join(goCodeDirectory, q.goFolderName(), "main.go")); err != nil {
		return err
	} else if goCode != "" {
		q.Langs = append(q.Langs, langGo)
		q.Codes[langGo] = strings.Trim(goCode, "\n")
	}

	if rustCode, err := getRustContent(path.Join(rustCodeDirectory, "src", q.rustFileName())); err != nil {
		return err
	} else if rustCode != "" {
		q.Langs = append(q.Langs, langRust)
		q.Codes[langRust] = strings.Trim(rustCode, "\n")
	}

	if javaCode, err := getJavaContent(path.Join(javaCodeDirectory, "src/main/java", q.javaFileName())); err != nil {
		return err
	} else if javaCode != "" {
		q.Langs = append(q.Langs, langJava)
		q.Codes[langJava] = strings.Trim(javaCode, "\n")
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

	return extractGoCode(string(bytes))
}

func extractGoCode(text string) (string, error) {
	arr := strings.Split(text, "// code\n")
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

	return extractRustCode(string(bytes))
}

func extractRustCode(text string) (string, error) {
	arr := strings.Split(text, "struct Solution;\n\n")
	if len(arr) < 2 {
		return "", errors.New("code content is empty")
	}
	arr = strings.Split(arr[1], "\n#[test]")
	if len(arr) < 2 {
		return "", errors.New("code content is empty")
	}
	return arr[0], nil
}

func getJavaContent(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return extractJavaCode(string(bytes))
}

func extractJavaCode(text string) (string, error) {
	content := text
	if strings.Contains(content, "/*") {
		return content[strings.Index(content, "/*"):], nil
	} else if strings.Contains(content, "class Solution") {
		return content[strings.Index(content, "class Solution"):], nil
	}

	return "", errors.New("code content is empty")
}

func extractTypeScriptCode(text string) (string, error) {
	var startPos int
	if strings.Contains(text, "/*") {
		startPos = strings.Index(text, "/*")
	}
	if !strings.Contains("test(") {
		returrn text[startPos:], nil
	}
	return text[startPos:strings.Index(text, "test('")], nil
}

func extractPythonCode(text string) (string, error) {
	startPos := strings.Index(text, "class Solution")
	if strings.Contains(text, "# Definition") {
		startPos = strings.Index(text, "# Definition")
	}
	if !strings.Contains(text, "class TestSolution") {
		return text[startPos:], nil
	}
	return text[startPos:strings.Index(text, "class TestSolution")], nil
}

func (q Question) goFolderName() string {
	return strconv.Itoa(q.ID) + "_" + strings.Replace(q.Slug, "-", "_", -1)
}

func (q Question) rustFileName() string {
	return "question_" + strconv.Itoa(q.ID) + ".rs"
}

func (q Question) javaFileName() string {
	return "question_" + strconv.Itoa(q.ID) + "/Solution.java"
}

func (q Question) typeScriptFileName() string {
	return strconv.Itoa(q.ID) + ".ts"
}

func (q Question) pythonFileName() string {
	return strconv.Itoa(q.ID) + ".py"
}

// FindByID ...
func (q QuestionList) FindByID(id int) (*Question, error) {
	for _, q := range q {
		if q.ID == id {
			return &q, nil
		}
	}

	return nil, errors.New("question not found")
}
