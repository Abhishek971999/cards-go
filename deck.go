package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// Create New Deck
func newDeck() deck{
	cards:=deck{}
	cardSuites:=[]string {"Spades","Diamonds","Hearts","Clubs"}
	cardValues:=[]string {"Ace","Two","Three","Four","Five","Six"}//and so on..
	for _,suite:=range cardSuites{
		for _,value:=range cardValues{
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

// Print Deck
func (cards deck) print(){
	for i,card :=range cards{
		fmt.Println(i,card)
	}
}

// Deal Deck
func deal(cards deck, handSize int)(deck,deck){
	return cards[:handSize],cards[handSize:]
}

// Convert Deck to String
func (cards deck) toString() string{
	return strings.Join([]string(cards),",")
}

//Save Deck to File
func (cards deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename,[]byte(cards.toString()),0666)
}

//Get a deck from a file
func newDeckFromFile(filename string) deck{
	bs,err:=ioutil.ReadFile(filename)
	if err!= nil{
		//If Error, log it and quit program
		os.Exit(1)
	}
	return deck(strings.Split(string(bs),","))//byte_slice->str->str_slice
}

//Shuffle Deck
func (cards deck) shuffle(){
	source:=rand.NewSource(time.Now().UnixNano())
	r:=rand.New(source)
	for i:=range cards{
		// newPos:=rand.Intn(len(cards)-1) //Inefficient due to same seed 
		newPos:=r.Intn(len(cards)-1) 
		cards[i],cards[newPos] = cards[newPos],cards[i]
	}
}