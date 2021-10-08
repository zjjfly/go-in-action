package ch8

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

var JSON = `{
	"name": "Gopher",
 	"title": "programmer",
 	"contact": {
 		"home": "415.333.3333",
		"cell": "415.555.5555"
 	}
}`

func Test_decode(t *testing.T) {
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(c)
}

func Test_decode_to_map(t *testing.T) {
	var c = make(map[string]interface{})
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(c)
	fmt.Println("Name:", c["name"])
	fmt.Println("Title:", c["title"])
	fmt.Println("Contact")
	fmt.Println("H:", c["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", c["contact"].(map[string]interface{})["cell"])
}
