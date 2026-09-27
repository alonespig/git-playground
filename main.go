package main

import (
	"fmt"
	"strings"
)

func main() {
	// Temporary debug output for the revert exercise.
	fmt.Println("DEBUG: start")
	fmt.Println(greeting("Gopher"))
	fmt.Println(greeting("   "))
}

func greeting(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "guest"
	}
	return "Hello, " + name
}
