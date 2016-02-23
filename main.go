package main

import (
	"fmt"
)

type Suit int

const (
	CLUBS Suit = iota
	DIAMOND
	HEART
	SPADE
)

// Spade
//U+2660 //&spades;

//Heart
//U+2665 //&hearts

//Diamond
//U+2666 //&diams;

//Clubs
//U+2663 //&clubs;
var suits = [...]string{
	"\u2663",
	"\u2666",
	"\u2665",
	"\u2660",
}

func (s Suit) String() string {
	return suits[s]
}

func main() {
	fmt.Println("Suits:", SPADE, HEART, DIAMOND, CLUBS)
}
