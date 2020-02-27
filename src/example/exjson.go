package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// JSON Example
func main() {
	man := Person{"Kang", 10}

	jsonbytes, err := json.Marshal(man)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonbytes))

	var woman Person
	err = json.Unmarshal(jsonbytes, &woman)
	if err != nil {
		panic(err)
	}
	fmt.Println(woman.Name, woman.Age)

}
