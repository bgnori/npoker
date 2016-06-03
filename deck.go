package npoker

import (
	"fmt"
	"math/rand"
	"strings"
)

type Deck []Card

type Index int

const (
	NullIndex Index = -1
)

func (d *Deck) Append(c Card) {
	*d = append([]Card(*d), c)
}

func (d Deck) String() string {
	var xs []string
	for _, c := range []Card(d) {
		xs = append(xs, fmt.Sprintf("%v%v", c.Rank(), c.Suit()))
	}
	return strings.Join(xs, ",")
}
func (d *Deck) Length() int {
	return len([]Card(*d))
}

func (d *Deck) Shuffle(r *rand.Rand) {
	n := d.Length()
	xs := []Card(*d)
	for i := n - 1; i >= 0; i-- {
		j := r.Intn(i + 1)
		xs[i], xs[j] = xs[j], xs[i]
	}
}

func (d *Deck) Drop(c Card) {
	xs := []Card(*d)
	for i, x := range xs {
		if x == c {
			*d = Deck(append(xs[:i], xs[i+1:]...))
			return
		}
	}
	panic(fmt.Sprintf("card %v, not found in %+v", c, xs))
}

func (d *Deck) Subtract(e Deck) {
	for _, c := range e {
		d.Drop(c)
	}
}

func BuildFullDeck() *Deck {
	d := new(Deck)
	for s := CLUBS; s <= SPADES; s++ {
		for r := ACE; r <= KING; r++ {
			d.Append(NewCard(r, s))
		}
	}
	return d
}

func (d *Deck) Clone() Deck {
	e := make([]Card, d.Length())
	copy(e, []Card(*d))
	return Deck(e)
}

func (d *Deck) ShrinkTo(count int) {
	*d = []Card(*d)[0:count]
}

func Join(d Deck, e Deck) Deck {
	return Deck(append([]Card(d), []Card(e)...))
}
