package main

import (
	"fmt"
	"math/rand"
)

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

type NumberCard struct {
	Card
	Number int
}

func NewNumberCard(number int, color int, description string) NumberCard {
	return NumberCard{
		Card: Card{
			Cod:         fmt.Sprintf("#%d_%s", number, colorNames[color]),
			Color:       color,
			Description: description,
		},
		Number: number,
	}
}

func GetRandomCard() NumberCard {
	color := rand.Intn(4)
	cardNumber := rand.Intn(10)
	cod := fmt.Sprintf("#%d_%s", cardNumber, colorNames[color])

	return NumberCard{
		Card: Card{
			Cod:         cod,
			Color:       color,
			Description: "",
		},
		Number: cardNumber,
	}
}

func main() {
	fmt.Println("Hello, World!")
}
