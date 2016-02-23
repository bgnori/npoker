package main

import (
	"fmt"
	"strings"
)

type Suit int

const (
	CLUBS Suit = iota
	DIAMONDS
	HEARTS
	SPADES
)

// Spades
//U+2660 //&spades;

//Hearts
//U+2665 //&hearts

//Diamonds
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

type Rank int

const (
	_ Rank = iota
	ACE
	DUCE
	TREY
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

var ranks = "0A23456789TJQK"

func (r Rank) String() string {
	return fmt.Sprintf("%c", ranks[r])
}

type Card struct {
	R Rank
	S Suit
}

func (c Card) String() string {
	return fmt.Sprintf("%v%v", c.R, c.S)
}

func SmokeCard() {
	fmt.Println("Suits:", SPADES, HEARTS, DIAMONDS, CLUBS)
	fmt.Println("Ranks:", ACE, DUCE, TEN, JACK, QUEEN, KING)
	fmt.Println("Card:", Card{ACE, HEARTS})
}

func parse(s string) Card {
	i := strings.Index(ranks, s[0:1])
	j := strings.Index("cdhs", s[1:2])
	return Card{Rank(i), Suit(j)}
}

func SmokeParseWith(input string) {
	got := parse(input)
	fmt.Println(input, "==>", got)
}

func main() {
	SmokeCard()
	SmokeParseWith("Ah")
	SmokeParseWith("5c")
}
