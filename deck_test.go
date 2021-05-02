package main

import (
	"os"
	"testing"
)


func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d)!=24{
		t.Errorf("Expected deck length of 24 but got %v",len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktest")
	d := newDeck()
	d.saveToFile("_decktest")
	loadedDeck:=newDeckFromFile("_decktest")
	if len(loadedDeck)!=24{
		t.Errorf("Expected deck length of 24 but got %v",len(d))
	}
	os.Remove("_decktest")
}