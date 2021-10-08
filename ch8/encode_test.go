package ch8

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func Test_encode_map(t *testing.T) {
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Print(string(data))
}
