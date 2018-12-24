// +build ignore

package main

import (
	"fmt"
	"time"
)

// p81
func main() {
	var c <-chan int
	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.")
	}
}
