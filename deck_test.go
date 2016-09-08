package npoker

import (
	"fmt"
	"testing"
)

func TestDeckString(t *testing.T) {
	expected := "As4h"
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

func TestMakeDeckFrom(t *testing.T) {
	d := Deck{NewCard(ACE, SPADES), NewCard(ACE, CLUBS)}
	e := Deck{NewCard(KING, HEARTS), NewCard(KING, DIAMONDS)}
	f := MakeDeckFrom(d, e)
	if f.Length() != 4 {
		t.Errorf("expected is 4.")
	}
	expected := "AsAcKhKd"
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

func TestCombOne(t *testing.T) {
	count := 0
	xs := BuildFullDeck()
	fmt.Println(xs)
	for i, x := range []Card(*xs) {
		fmt.Println(i, x)
	}
	for comb := range xs.CombOne() {
		count += 1
		fmt.Println(comb.Chosen)
		if comb.Chosen.Length() != 1 {
			t.Errorf("Chosen length is expected to be 1, but got %d,\n%s", comb.Chosen.Length(), comb.Chosen)
		}
		if comb.Rest.Length() != 51 {
			t.Errorf("Rest length is expected to be 51, but got %d,\n%s", comb.Rest.Length(), comb.Rest)
		}
	}
	if count != 52 {
		t.Errorf("# of generations is expected to be 52.")
	}
}

func TestCombTwo(t *testing.T) {
	count := 0
	xs := BuildFullDeck()

	fmt.Println(xs)
	for comb := range xs.CombTwo() {
		count += 1
		cards := []Card(comb.Chosen)
		if cards[0] == cards[1] {
			t.Errorf("found same Card in cards %s", cards)
		}
		fmt.Println(comb.Chosen)

		if comb.Chosen.Length() != 2 {
			t.Errorf("Chosen length is expected to be 2, but got %d", comb.Chosen.Length())
		}
		if comb.Rest.Length() != 50 {
			t.Errorf("Rest length is expected to be 50, but got %d", comb.Rest.Length())
		}
	}
	if count != 52*51/2 {
		t.Errorf("# of generations is expected to be 52 * 51 /2 = %d, but got %d.", 52*51/2, count)
	}
}

func TestCombThree(t *testing.T) {
	count := 0
	xs := BuildFullDeck()
	for comb := range xs.CombThree() {
		count += 1
		flop := []Card(comb.Chosen)
		if flop[0] == flop[1] || flop[1] == flop[2] || flop[0] == flop[2] {
			t.Errorf("found same Card in flop %s", flop)
		}

		if comb.Chosen.Length() != 3 {
			t.Errorf("Chosen length is expected to be 3, but got %d", comb.Chosen.Length())
		}
		if comb.Rest.Length() != 49 {
			t.Errorf("Rest length is expected to be 49, but got %d", comb.Rest.Length())
		}
	}
	if count != 52*51*50/(3*2*1) {
		t.Errorf("# of generations is expected to be 52 * 51 *50/(3*2*1) = %d, but got %d.", 52*51*50/(3*2*1), count)
	}
}
