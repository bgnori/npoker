package main

import (
	"fmt"
	"math/rand"
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
func SmokeParse() {
	SmokeParseWith("Ah")
	SmokeParseWith("5c")
	//SmokeParseWith("Xc")
}

type Deck []Card

func (d *Deck) Append(c Card) {
	*d = append([]Card(*d), c)
}

func (d Deck) String() string {
	var xs []string
	for _, c := range d {
		xs = append(xs, fmt.Sprintf("%v%v", c.R, c.S))
	}
	return strings.Join(xs, ",")
}
func (d *Deck) Length() int {
	return len([]Card(*d))
}

func (d *Deck) Shuffle() {
	n := d.Length()
	xs := []Card(*d)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		xs[i], xs[j] = xs[j], xs[i]
	}
}

func BuildDeck() *Deck {
	d := new(Deck)
	for s := CLUBS; s <= SPADES; s++ {
		for r := ACE; r <= KING; r++ {
			d.Append(Card{r, s})
		}
	}
	return d
}

func SmokeDeck() {
	d := Deck{Card{ACE, SPADES}, Card{FOUR, HEARTS}}
	fmt.Println(d)
	e := BuildDeck()
	fmt.Println(e)
	fmt.Println(e.Length())
	e.Shuffle()
	fmt.Println(e)
	e.Shuffle()
	fmt.Println(e)
}

func main() {
	SmokeCard()
	SmokeParse()
	SmokeDeck()
}
