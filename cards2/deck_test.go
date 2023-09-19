package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if d == nil {
		t.Errorf("Expected non nil value, found nil")
	}
}

func TestSaveToFileAndReadDeckFromFile(t *testing.T) {
	const fName = "_destTest"
	os.Remove(fName)

	d := newDeck()
	d.saveToFile(fName)

	d2, err := readDeckFromFile(fName)

	if err != nil {
		t.Errorf("Failed to read file %v", err)
	}

	if len(d) != len(d2) {
		t.Errorf("Excepcted length of read deck 12, found %v", len(d2))
	}
}
