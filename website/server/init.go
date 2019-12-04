package server

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type problem struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Difficulty int    `json:"difficulty"`
}

var problems = make(map[string]*problem)

func init() {
	bytes, err := ioutil.ReadFile("./data.json")
	if err != nil {
		panic(err)
	}

	var ps []*problem
	if err := json.Unmarshal(bytes, &ps); err != nil {
		panic(err)
	}

	for _, p := range ps {
		problems[strconv.Itoa(p.ID)] = p
	}
}
