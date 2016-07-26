package npoker

import (
	"fmt"
	"testing"
)

func TestDeckString(t *testing.T) {
	expected := "As,4h"
	d := Deck{NewCard(ACE, SPADES), NewCard(FOUR, HEARTS)}
	got := fmt.Sprintf("%v", d)
	if expected != got {
		t.Errorf("'%s' is expected for %d, but got %v", expected, d, got)
	}
}

func TestDeckJSON(t *testing.T) {
	d := Deck{NewCard(ACE, SPADES), NewCard(FOUR, HEARTS)}
	expected := "\"A\u26604\u2665\""
	j, err := d.MarshalJSON()
	if err != nil {
		t.Errorf("'%s' is expected for %d, but error! %v", expected, d, err)
	}
	if expected != string(j) {
		t.Errorf("'%s' is expected for %d, but got %v", expected, d, j)
	}
}

func TestFullDeck(t *testing.T) {
	d := BuildFullDeck()
	if d.Length() != 52 {
		t.Errorf("full length is expected to be 52.")
	}
}

func TestDropDeck(t *testing.T) {
	d := Deck{NewCard(ACE, SPADES), NewCard(ACE, CLUBS)}
	d.Drop(NewCard(ACE, SPADES))
	if d.Length() != 1 {
		t.Errorf("expected is 1.")
	}
	expected := "Ac"
	got := fmt.Sprintf("%s", d)

	if expected != got {
		t.Errorf("'%s' is expected for %d, but got %v", expected, d, got)
	}
}

func TestJoinDeck(t *testing.T) {
	d := Deck{NewCard(ACE, SPADES), NewCard(ACE, CLUBS)}
	e := Deck{NewCard(KING, HEARTS), NewCard(KING, DIAMONDS)}
	f := Join(d, e)
	if f.Length() != 4 {
		t.Errorf("expected is 4.")
	}
	expected := "As,Ac,Kh,Kd"
	got := fmt.Sprintf("%s", f)

	if expected != got {
		t.Errorf("'%s' is expected for %d, but got %v", expected, f, got)
	}

}

func TestShuffle(t *testing.T) {
	d := BuildFullDeck()
	if d.Length() != 52 {
		t.Errorf("full length is expected to be 52.")
	}
	r := NewRand()
	r.SeedFromBytes(GetSeedFromURAND())
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

func BenchmarkShuffle(b *testing.B) {
	d := BuildFullDeck()
	r := NewRand()
	r.SeedFromBytes(GetSeedFromURAND())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Shuffle(r)
	}

}
