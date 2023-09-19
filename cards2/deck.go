package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	const of = " of "
	suites := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	values := []string{"Ace", "Two", "Three"}
	d := deck{}

	for _, suite := range suites {
		for _, val := range values {
			d = append(d, val+of+suite)
		}
	}
	return d
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func deckFromString(s string) deck {
	strSlice := strings.Split(s, ",")
	return deck(strSlice)
}

func (d deck) saveToFile(fileName string) error {
	deckStr := d.toString()
	err := os.WriteFile(fileName, []byte(deckStr), 0666)
	if err != nil {
		fmt.Println("Failed to write to ", fileName, err)
		return err
	}
	return nil
}

func readDeckFromFile(fileName string) (deck, error) {
	byteSlice, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Failed to readfile ", fileName, err)
		return nil, err
	}
	return deckFromString(string(byteSlice)), nil
}

func (d deck) shuffle() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	n := len(d)
	for i := range d {
		newPos := r.Intn(n)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
