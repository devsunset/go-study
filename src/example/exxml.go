package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// XML Example
func main() {
	man := Person{"Kang", 10}

	xmlbytes, err := xml.Marshal(man)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(xmlbytes))

	var woman Person
	err = xml.Unmarshal(xmlbytes, &woman)
	if err != nil {
		panic(err)
	}
	fmt.Println(woman.Name, woman.Age)

}
