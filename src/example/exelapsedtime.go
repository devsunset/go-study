package main

import (
	"fmt"
	"time"
)

// Elapsed Time Example
func main() {
	startTime := time.Now()

	for i := 0; i < 1000; i++ {
		fmt.Println("Working...")
	}

	elapsedTime := time.Since(startTime)

	fmt.Printf("Elasped Time : %s", elapsedTime)
}
