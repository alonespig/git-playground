package main

import (
	"fmt"
	"strings"
)

func main() {
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
