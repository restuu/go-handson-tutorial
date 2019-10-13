package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := []byte(`{"id":1,"name":"Gopher"}`)
	data := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{}

	if err := json.Unmarshal(str, &data); err != nil {
		panic(err)
	}

	fmt.Println("ID: ", data.ID, "Name: ", data.Name)
}
