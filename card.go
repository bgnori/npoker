package npoker

import (
	"errors"
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

var suits = []rune{'\u2663', '\u2666', '\u2665', '\u2660'}

func (s Suit) String() string {
	return fmt.Sprintf("%c", suits[s])
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

type Card int

func NewCard(r Rank, s Suit) Card {
	return Card(int(s)*13 + int(r) - 1)
}

func (c Card) Rank() Rank {
	return Rank(c%13) + 1
}

func (c Card) Suit() Suit {
	return Suit(c / 13)
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Rank(), c.Suit())
}

func (c Card) Ord() int {
	return int(c)
}

func MatchRank(r rune) (Rank, error) {
	var found int
	if found = strings.IndexRune(ranks, r); found == -1 {
		return RANKS, errors.New(fmt.Sprintf("no such rune found: %#U", r))
	}
	return Rank(found), nil
}

func MatchSuit(r rune) (Suit, error) {
	if found := strings.IndexRune("cdhs", r); 0 <= found && found < 4 {
		return Suit(found), nil
	}
	if found := strings.IndexRune(string(suits), r); 0 <= found && found < 4 {
		return Suit(found), nil
	}
	return SUITS, errors.New(fmt.Sprintf("no such rune found: %#U", r))
}

func parse(s string) Card {
	i := strings.Index(ranks, s[0:1])
	j := strings.Index("cdhs", s[1:2])
	return NewCard(Rank(i), Suit(j))
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
