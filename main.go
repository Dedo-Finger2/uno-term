package main

import "fmt"

const (
	YELLOW = iota
	RED
	GREEN
	BLUE
)

type Card struct {
	Cod         string
	Name        string
	Color       int
	Description string
}

func main() {
	fmt.Println("Hello, World!")
}
