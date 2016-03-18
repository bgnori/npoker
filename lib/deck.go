package npoker

import (
	"fmt"
	"math/rand"
	"strings"
)

type Deck []Card

func (d *Deck) Append(c Card) {
	*d = append([]Card(*d), c)
}

func (d Deck) String() string {
	var xs []string
	for _, c := range []Card(d) {
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

func (d *Deck) Drop(c Card) {
	xs := []Card(*d)
	for i, x := range xs {
		if x == c {
			*d = Deck(append(xs[:i], xs[i+1:]...))
			return
		}
	}
	panic("card not found")
}

func BuildFullDeck() *Deck {
	d := new(Deck)
	for s := CLUBS; s <= SPADES; s++ {
		for r := ACE; r <= KING; r++ {
			d.Append(Card{r, s})
		}
	}
	return d
}

func Join(d Deck, e Deck) Deck {
	return Deck(append([]Card(d), []Card(e)...))
}
