package main

import "fmt"

func main() {
	d := newDeck()
	d.shuffle()
	handSize := 3
	i := 0
	for len(d) > handSize {
		var hand deck
		hand, d = deal(d, handSize)
		hand.saveToFile(fmt.Sprintf("%s_%d", "hand", i))
		i++
	}
	d.saveToFile("remaining_deck")
	remainingD, err := readDeckFromFile("remaining_deck")
	if err != nil {
		panic(err)
	}
	remainingD.print()
}
