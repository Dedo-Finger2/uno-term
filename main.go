package main

import "fmt"

var colorNames = []string{
	"YELLOW",
	"RED",
	"GREEN",
	"BLUE",
}

const (
	YELLOW = iota
	RED
	GREEN
	BLUE
)

type Card struct {
	Cod         string
	Color       int
	Description string
}

func main() {
	fmt.Println("Hello, World!")
}
