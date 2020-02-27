package main

import (
	"flag"
	"fmt"
)

// Command Line Flag Example
func main() {
	// go run flag.go -help
	// go run flag.go -aa=test -bb=9 -cc=true

	a := flag.String("aa", "-", "Input String Command Option Value")
	b := flag.Int("bb", 0, "Input Int Command Option Value")
	c := flag.Bool("cc", false, "Input Bool Command Option Value")

	flag.Parse()

	fmt.Println(*a, *b, *c)
}
