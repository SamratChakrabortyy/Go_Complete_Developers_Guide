package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newdeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuite := range cardSuites {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuite)
		}
	}
	return cards
}

func (d deck) deal(handSize int) (deck, deck) {
	hand := d[:handSize]
	d = d[handSize:]
	return hand, d
}
