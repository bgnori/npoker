package npoker

import (
	//"encoding/json"
	"fmt"
	"strings"
)

type Request struct {
	Source  string `json:"source"`
	Players []Deck `json:"players"`
	Board   []Deck `json:"board"`
	Trials  int    `json:"trials"`
}

type Response struct {
}

type WorkSet struct {
	xs      Deck
	board   Deck
	players []Deck
}

type Summary struct {
	players []Deck
	Count   int
	Wins    []int
	Eqs     []int
}

func NewWorkSet(board []Deck, players []Deck) *WorkSet {
	xs := BuildFullDeck()
	b := Deck{}
	for _, street := range board {
		xs.Subtract(street)
		b = Join(b, street)
	}
	for _, p := range players {
		xs.Subtract(p)
	}
	return &WorkSet{*xs, b, players}
}

func (x *WorkSet) Clone() *WorkSet {
	return &WorkSet{
		x.xs.Clone(),
		x.board.Clone(),
		x.players, //  Ugh!
	}
}

func (x *WorkSet) Run(r *Rand) *ShowDown {
	x.xs.Shuffle(r)
	x.xs.ShrinkTo(5 - len(x.board))
	river := Join(x.board, x.xs)
	return MakeShowDown(river, x.players...)
}

func NewSummary(players []Deck) *Summary {
	return &Summary{
		players: players,
		Count:   0,
		Wins:    make([]int, len(players)),
		Eqs:     make([]int, len(players)),
	}
}

func (summary *Summary) Add(sd *ShowDown) *Summary {
	for _, idx := range sd.Winners {
		summary.Wins[idx] += 1
	}
	xs := DistrubuteChips(1000, 1, 0, sd)
	for i, v := range xs {
		summary.Eqs[i] += v
	}
	summary.Count += 1
	return summary
}

func (summary *Summary) String() string {
	var xs []string
	xs = append(xs, fmt.Sprintf(" %d players, %d trials", len(summary.players), summary.Count))
	for i, v := range summary.players {
		s := fmt.Sprintf("player %d has %s, won %.3f, won eq of %.2f.",
			i, v, float64(summary.Wins[i])/float64(summary.Count), float64(summary.Eqs[i])/float64(summary.Count))
		xs = append(xs, s)
	}
	return strings.Join(xs, "\n")
}
