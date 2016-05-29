package npoker

import (
	"fmt"
	"math/rand"
	"testing"
)

func deckhelper(t *testing.T, expected string, to_string Deck) {
	if expected != fmt.Sprintf("%v", to_string) {
		t.Errorf("'%s' is expected for %d, but got %v", expected, to_string, to_string)
	}
}

func TestDeckString(t *testing.T) {
	d := Deck{Card{ACE, SPADES}, Card{FOUR, HEARTS}}
	deckhelper(t, "A\u2660,4\u2665", d)
}

func TestFullDeck(t *testing.T) {
	d := BuildFullDeck()
	if d.Length() != 52 {
		t.Errorf("full length is expected to be 52.")
	}
}

func TestDropDeck(t *testing.T) {
	d := Deck{Card{ACE, SPADES}, Card{ACE, CLUBS}}
	d.Drop(Card{ACE, SPADES})
	if d.Length() != 1 {
		t.Errorf("expected is 1.")
	}
	deckhelper(t, "A\u2663", d)
}

func TestJoinDeck(t *testing.T) {
	d := Deck{Card{ACE, SPADES}, Card{ACE, CLUBS}}
	e := Deck{Card{KING, HEARTS}, Card{KING, DIAMONDS}}
	f := Join(d, e)
	if f.Length() != 4 {
		t.Errorf("expected is 4.")
	}
	deckhelper(t, "A\u2660,A\u2663,K\u2665,K\u2666", f)
}

func TestShuffle(t *testing.T) {
	d := BuildFullDeck()
	if d.Length() != 52 {
		t.Errorf("full length is expected to be 52.")
	}
	r := rand.New(rand.NewSource(0))
	d.Shuffle(r)
	if d.Length() != 52 {
		t.Errorf("full length is expected to be 52.")
	}
}

func TestShrinkToDeck(t *testing.T) {
	xs := BuildFullDeck()
	xs.ShrinkTo(10)
	if xs.Length() != 10 {
		t.Errorf("length is expected to be 10. but got %d", xs.Length())
	}

}

func TestCloneDeck(t *testing.T) {
	xs := BuildFullDeck()
	ys := xs.Clone()
	if ys.Length() != 52 {
		t.Errorf("length is expected to be 52, but got %d", xs.Length())
	}
}
