package main

import (
	"encoding/json"
	"fmt"
)

type State struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Cities []string `json:"cities"`
}

func (s State) ToJson() (string, error) {
	stateBytes, error := json.Marshal(s)
	if error != nil {
		return "", error
	}
	stateJson := string(stateBytes)
	return stateJson, nil
}

func (s *State) JsonToStruct(jsonParam string) error {
	bytes := []byte(jsonParam)
	error := json.Unmarshal(bytes, s)
	if error != nil {
		return error
	}
	return nil
}

func main() {
	state := State{
		ID:   1,
		Name: "California",
		Cities: []string{
			"Escalon",
			"Belmont",
			"Corcoran",
			"Avenal",
		},
	}

	// creating json
	jsonCreated, error := state.ToJson()
	if error != nil {
		panic(error)
	}
	fmt.Println("JSON CREATED -> ", jsonCreated)

	// modified struct with other json
	jsonStr := `{"id":4,"name":"São Paulo","cities":["Bauru"]}`
	error = state.JsonToStruct(jsonStr)
	if error != nil {
		panic(error)
	}
	fmt.Println("STRUCT MODIFIED -> ", state)
}
