package data

import (
	"embed"
	"encoding/json"
)

//go:embed data.json
var dataFS embed.FS

func GetQuestions() (questions, error) {
	bytes, err := dataFS.ReadFile("data.json")
	if err != nil {
		return nil, err
	}

	var qs questions
	if err = json.Unmarshal(bytes, &qs); err != nil {
		return nil, err
	}
	return qs, nil
}
