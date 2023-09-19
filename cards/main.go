package main

import "fmt"

func main() {
	cards := newdeck()
	hand, remainingCards := cards.deal(2)

	cards = remainingCards

	fmt.Print("Hand: \n\n")
	hand.print()

	fmt.Print("deck : \n\n")
	cards.print()
}
