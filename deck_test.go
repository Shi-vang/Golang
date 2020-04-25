package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 160 {
		t.Errorf("wronge input fail")
	}
}

func TestDeckFileIo(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadDeck := newDeckFromFile("_decktesting")
	if len(loadDeck) != 49 {
		t.Errorf("pass file test")
	}

	os.Remove("_decktesting")
}
