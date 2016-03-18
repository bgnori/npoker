package npoker

import (
	"fmt"
	//	"testing"
)

func SmokeDeck() {
	d := Deck{Card{ACE, SPADES}, Card{FOUR, HEARTS}}
	fmt.Println(d)
	e := BuildFullDeck()
	fmt.Println(e)
	fmt.Println(e.Length())
	e.Shuffle()
	fmt.Println(e)
	e.Shuffle()
	fmt.Println(e)
	f := BuildFullDeck()
	for i, c := range []Card(*f) {
		fmt.Println(i, c, c.Ord())
	}
	f.Drop(Card{ACE, HEARTS})
	fmt.Println(f.Length())
	fmt.Println(f)
	f.Shuffle()
	fmt.Println(f)

	g := Deck{Card{ACE, SPADES}, Card{ACE, CLUBS}}
	h := Deck{Card{KING, HEARTS}, Card{KING, DIAMONDS}}
	fmt.Println(Join(g, h))
}
