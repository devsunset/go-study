package main

import (
	"fmt"
	"os"
)

// Command Line Argument Example
func main() {
	// go run osargs.go a b c

	for i, v := range os.Args {
		fmt.Println(i, v)
	}

	fmt.Println(os.Args[0:])

	fmt.Println(os.Args[1])
}
