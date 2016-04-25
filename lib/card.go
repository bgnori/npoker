package npoker

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
	SUITS
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
	HIACE
	RANKS
)

var ranks = "0A23456789TJQKA"

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

func (c Card) Ord() int {
	//2 <= x <=53
	x := 0
	switch c.R {
	case ACE:
		x += 14
	default:
		x += int(c.R)
	}
	return x + int(c.S)*int(KING)
}

func parse(s string) Card {
	i := strings.Index(ranks, s[0:1])
	j := strings.Index("cdhs", s[1:2])
	return Card{Rank(i), Suit(j)}
}

func SuitPermZero() [][]Suit {
	return [][]Suit{}
}

func SuitPermOne() [][]Suit {
	return [][]Suit{
		[]Suit{CLUBS},
		[]Suit{DIAMONDS},
		[]Suit{HEARTS},
		[]Suit{SPADES},
	}
}

func SuitPermTwo() [][]Suit {
	return [][]Suit{
		[]Suit{CLUBS, DIAMONDS},
		[]Suit{CLUBS, HEARTS},
		[]Suit{CLUBS, SPADES},
		[]Suit{DIAMONDS, HEARTS},
		[]Suit{DIAMONDS, SPADES},
		[]Suit{HEARTS, SPADES},
	}
}

func SuitPermThree() [][]Suit {
	return [][]Suit{
		[]Suit{CLUBS, DIAMONDS, HEARTS},
		[]Suit{CLUBS, DIAMONDS, SPADES},
		[]Suit{CLUBS, HEARTS, SPADES},
		[]Suit{DIAMONDS, HEARTS, SPADES},
	}
}

func SuitPermFour() [][]Suit {
	return [][]Suit{[]Suit{CLUBS, DIAMONDS, HEARTS, SPADES}}
}
