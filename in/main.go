package main

import (
	"fmt"
	"encoding/json"
)

type Datum struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Version struct {
	Ref string `json:"ref,omitempty"`
}

type Output struct {
	Version  Version `json:"version,omitempty"`
	Metadata []Datum `json:"metadata,omitempty"`
}

func main() {
	output, err := json.Marshal(Output{})
	if err != nil {
		panic(err)
	}
	fmt.Print(string(output))
}
