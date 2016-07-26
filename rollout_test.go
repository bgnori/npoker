package npoker

import (
	"testing"
)

func BenchmarkRollOut(b *testing.B) {
	PlayerOneHole := Deck{
		NewCard(NINE, DIAMONDS),
		NewCard(SIX, DIAMONDS),
	}

	PlayerTwoHole := Deck{
		NewCard(FIVE, CLUBS),
		NewCard(FIVE, HEARTS),
	}

	PlayerThreeHole := Deck{
		NewCard(JACK, CLUBS),
		NewCard(EIGHT, CLUBS),
	}

	board := []Deck{
		Deck{
			NewCard(NINE, CLUBS),
			NewCard(SEVEN, CLUBS),
			NewCard(FIVE, DIAMONDS),
		},
		Deck{
			NewCard(FOUR, DIAMONDS),
		},
	}

	req := Request{
		Source:  "rollout test",
		Players: []Deck{PlayerOneHole, PlayerTwoHole, PlayerThreeHole},
		Board:   board,
		Trials:  1,
	}
	w := NewWorkSet(req.Board, req.Players)
	seed := GetSeedFromRAND()
	summary := NewSummary(req, seed)

	r := NewRand()
	r.SeedFromBytes(seed)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u := w.Clone()
		summary.Add(u.Run(r))
	}
}