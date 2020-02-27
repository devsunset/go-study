package main

import (
	"fmt"
	"os"
)

// Environment value Example
func main() {
	gopath := os.Getenv("GOPATH")
	fmt.Println("GOPATH : ", gopath)

	os.Setenv("TEST_ENV", "GO_TEST")
	fmt.Printf("TEST_ENV : %s", os.Getenv("TEST_ENV"))

	for i, env := range os.Environ() {
		fmt.Println(i, env)
	}
}
