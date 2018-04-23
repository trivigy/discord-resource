package main

import (
	"encoding/json"
	"fmt"
)

type Version struct {
	Ref string `json:"ref,omitempty"`
}

func main() {
	output, err := json.Marshal([]Version{})
	if err != nil {
		panic(err)
	}
	fmt.Print(string(output))
}
