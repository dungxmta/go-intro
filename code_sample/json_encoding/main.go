package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Object struct {
	Name    string `json:"name"`
	Public  int    `json:"public"`
	private int    `json:"private"`
}

func newObject() *Object {
	return &Object{
		Name:    "a",
		Public:  1,
		private: 2,
	}
}

func main() {
	obj := newObject()
	fmt.Printf("Go type: %v\n", *obj)
	// Go type: {a 1 2}

	err := jsonDumps(obj)
	if err != nil {
		return
	}
	// Json dumps: {"name":"a","public":1}

	jStr := `{
		"name": "123",
		"public": 2,
		"private": 3
	}`
	err = jsonLoads(jStr)
	if err != nil {
		return
	}
	// Go type: {123 2 0}
	// Json dumps: {"name":"123","public":2}
}

func jsonDumps(obj *Object) error {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Print("Json dumps: ")
	fmt.Println(string(bytes))
	return nil
}

func jsonLoads(s string) (err error) {
	var obj Object
	err = json.Unmarshal([]byte(s), &obj)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Print("Go type: ")
	fmt.Println(obj)

	err = jsonDumps(&obj)
	return
}
