package npoker

import (
	"fmt"
	"math/rand"
	"strings"
)

type EqCalc struct {
	xs      Deck
	board   Deck
	players []Deck
}

func NewEqCalc(board Deck, players []Deck) Runnable {
	xs := BuildFullDeck()
	xs.Subtract(board)
	for _, p := range players {
		xs.Subtract(p)
	}
	return &EqCalc{*xs, board, players}
}

func (x *EqCalc) Clone() Runnable {
	return &EqCalc{
		x.xs.Clone(),
		x.board.Clone(),
		x.players, //  Ugh!
	}
}

func (x *EqCalc) Run(source rand.Source) Result {
	r := rand.New(source)
	x.xs.Shuffle(r)
	x.xs.ShrinkTo(5 - len(x.board))
	river := Join(x.board, x.xs)
	return MakeShowDown(river, x.players...)
}

type EqSummarizer struct {
	calc  *EqCalc
	count int
	wins  []int
	eqs   []int
}

func NewEqSummarizer(r Runnable) Summarizer {
	c := r.(*EqCalc)
	return &EqSummarizer{
		calc:  c,
		count: 0,
		wins:  make([]int, len(c.players)),
		eqs:   make([]int, len(c.players)),
	}
}

func (x *EqSummarizer) Zero() Summary {
	return x
}

func (x *EqSummarizer) String() string {
	var xs []string
	xs = append(xs, fmt.Sprintf(" %d players, %d trials", len(x.calc.players), x.count))
	for i, v := range x.calc.players {
		s := fmt.Sprintf("player %d has %s, won %.3f, won eq of %.2f.",
			i, v, float64(x.wins[i])/float64(x.count), float64(x.eqs[i])/float64(x.count))
		xs = append(xs, s)
	}

	return strings.Join(xs, "\n")
}

func (x *EqSummarizer) Fold(s Summary, result Result) Summary {
	sd := result.(*ShowDown)
	summary := s.(*EqSummarizer)
	for _, idx := range sd.Winners {
		summary.wins[idx] += 1
	}
	xs := DistrubuteChips(1000, 1, 0, sd)
	for i, v := range xs {
		summary.eqs[i] += v
	}
	summary.count += 1
	return summary
}
