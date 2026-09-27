package main

import "fmt"

func init() {
	fmt.Println(farewell("Gopher"))
}

func farewell(name string) string {
	return "Goodbye, " + name
}
