package main

import (
	"crypto/sha256"
	"fmt"
)

// Sha256 Example
func main() {
	s := "sha256 text"
	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
