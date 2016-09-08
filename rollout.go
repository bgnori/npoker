package npoker

import (
	//"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Request struct {
	Source  string `json:"source"`
	Players []Deck `json:"players"`
	Board   []Deck `json:"board"`
	Trials  int    `json:"trials"`
	Seed    []byte `json:"seed"`
}

type Response struct {
}

type WorkSet struct {
	xs      Deck
	board   Deck
	players []Deck
}

type Summary struct {
	Req   Request `json:"request"`
	Count int     `json:"count"`
	Wins  []int   `json:"win"`
	Eqs   []int   `json:"eqs"`
	Start time.Time
	End   time.Time
}

func NewWorkSet(board []Deck, players []Deck) *WorkSet {
	xs := BuildFullDeck()
	b := Deck{}
	for _, street := range board {
		xs.Subtract(street)
		b = MakeDeckFrom(b, street)
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
	river := MakeDeckFrom(x.board, x.xs)
	return MakeShowDown(river, x.players...)
}

func (w *WorkSet) ByComb(summary *Summary) {
	//assurme empty board
	for flop := range w.xs.CombThree() {
		for turn := range flop.Rest.CombOne() {
			for river := range turn.Rest.CombOne() {
				board := MakeDeckFrom(flop.Chosen, turn.Chosen, river.Chosen)
				fmt.Println(board)
				summary.Add(MakeShowDown(board, w.players...))
			}
		}
	}
}

func NewSummary(req Request) *Summary {
	return &Summary{
		Req:   req,
		Count: 0,
		Wins:  make([]int, len(req.Players)),
		Eqs:   make([]int, len(req.Players)),
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
	xs = append(xs, fmt.Sprintf(" %d players, %d trials", len(summary.Req.Players), summary.Count))
	for i, v := range summary.Req.Players {
		s := fmt.Sprintf("player %d has %s, won %.3f, won eq of %.2f.",
			i, v, float64(summary.Wins[i])/float64(summary.Count), float64(summary.Eqs[i])/float64(summary.Count))
		xs = append(xs, s)
	}
	return strings.Join(xs, "\n")
}
