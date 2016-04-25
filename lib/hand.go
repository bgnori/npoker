package npoker

import (
	"fmt"
	"strings"
)

type PokerHands int

const (
	HighCard PokerHands = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

var pokerHandsString = []string{
	"high card",
	"one pair",
	"two pair",
	"three of a kind",
	"straight",
	"flush",
	"four of a kind",
	"straight flush",
}

type PokerHandDiscriptor struct {
	ph    PokerHands
	which []int
}

type CardRanking struct {
	xs        Deck
	highcards []Index
	pairs     [][]int
	threes    [][]int
	fours     [][]int
	straight  [][]int
	cards     [][]Index //rank, suite, index
}

func prepareCardRanking(d Deck) CardRanking {
	cr := CardRanking{
		xs:        d,
		highcards: nil,
		pairs:     nil,
		threes:    nil,
		fours:     nil,
		straight:  nil,
		cards:     nil,
	}

	cr.cards = make([][]Index, RANKS-1)
	for r := ACE; r < HIACE; r += 1 {
		cr.cards[r] = make([]Index, SUITS-1)
		for s := CLUBS; s < SUITS; s += 1 {
			cr.cards[r][s] = NullIndex
		}
	}
	for i, x := range d {
		cr.cards[x.R][x.S] = Index(i)
	}
	return cr

}

func MakeCardRanking(xs Deck) CardRanking {
	cr := prepareCardRanking(xs)
	cr.calcPairwise()
	cr.calcStraight()
	return cr
}

func (cr *CardRanking) calcPairwise() {
	cr.highcards = make([]Index, 0, 5)
	cr.pairs = make([][]int, 0, 5)
	cr.threes = make([][]int, 0, 5)
	cr.fours = make([][]int, 0, 5)
	/*
		for i := HIACE; i > ACE; i -= 1 {
			if len(cr.ranks[i]) >= 1 {
				cr.highcards = append(cr.highcards, cr.ranks[i][0])
			}
			if len(cr.ranks[i]) >= 2 {
				cr.pairs = append(cr.pairs, []int{cr.ranks[i][0], cr.ranks[i][1]})
			}
			if len(cr.ranks[i]) >= 3 {
				cr.threes = append(cr.threes, []int{cr.ranks[i][0], cr.ranks[i][1], cr.ranks[i][2]})
			}
			if len(cr.ranks[i]) >= 4 {
				cr.fours = append(cr.fours, []int{cr.ranks[i][0], cr.ranks[i][1], cr.ranks[i][2], cr.ranks[i][3]})
			}
		}
	*/

	for r := HIACE; r > ACE; r -= 1 {
		for s := CLUBS; s < SUITS; s += 1 {
			if cr.cards[r][s] != NullIndex {
				// UGH! Duplication
				cr.highcards = append(cr.highcards, cr.cards[r][s])
			}
		}

	}
}

type FIFO5 struct {
	xs []int
}

func newFIFO5() *FIFO5 {
	return &FIFO5{xs: make([]int, 0)}
}

func (f *FIFO5) Push(x int) {
	if len(f.xs) == 5 {
		copy(f.xs, f.xs[1:5])
		f.xs[4] = x
	} else {
		f.xs = append(f.xs, x)
	}
}

func (f *FIFO5) Empty() {
	f.xs = make([]int, 0)
}

func (f *FIFO5) CloneXS() []int {
	ys := make([]int, len(f.xs))
	copy(ys, f.xs)
	return ys
}

func (cr *CardRanking) calcStraightSub(head Rank, f *FIFO5, filled bool) (Rank, bool) {
	/*
		var start Rank
		start = 0
		if filled {
			start = 4
		}
			for i := start; i < 5; i++ {
				if len(cr.ranks[head-i]) > 0 {
					f.Push(cr.ranks[head-i][0])
				} else {
					//fail case
					f.Empty()
					return head - i - 1, false
				}
			}
	*/
	return head - 1, true

}

func (cr *CardRanking) calcStraight() {
	var next Rank
	f := newFIFO5()
	cr.straight = make([][]int, RANKS)
	found := false
	next = 0
	for i := HIACE; i >= FIVE; {
		next, found = cr.calcStraightSub(i, f, found)
		if found {
			cr.straight[int(HIACE-i)] = f.CloneXS()
		}
		i = next
	}
}

func (cr CardRanking) String() string {
	var xs []string

	xs = append(xs, fmt.Sprintf("%v", cr.xs))

	var highcards []string
	for _, pos := range cr.highcards {
		highcards = append(highcards, fmt.Sprintf("%v", cr.xs[pos]))
	}
	xs = append(xs, "highcards:")
	xs = append(xs, strings.Join(highcards, ","))

	var pairs []string
	for _, p := range cr.pairs {
		pairs = append(pairs, fmt.Sprintf("%v %v", cr.xs[p[0]], cr.xs[p[1]]))
	}
	xs = append(xs, "pairs:")
	xs = append(xs, strings.Join(pairs, ","))

	var threes []string
	for _, p := range cr.threes {
		threes = append(threes, fmt.Sprintf("%v %v %v", cr.xs[p[0]], cr.xs[p[1]], cr.xs[p[2]]))
	}
	xs = append(xs, "threes:")
	xs = append(xs, strings.Join(threes, ","))

	var fours []string
	for _, p := range cr.fours {
		fours = append(fours, fmt.Sprintf("%v %v %v %v", cr.xs[p[0]], cr.xs[p[1]], cr.xs[p[2]], cr.xs[p[3]]))
	}
	xs = append(xs, "fours:")
	xs = append(xs, strings.Join(fours, ","))

	var straight []string
	for _, p := range cr.straight {
		if len(p) > 0 {
			fmt.Println("%+v", p)
			straight = append(straight,
				fmt.Sprintf("%v %v %v %v %v",
					cr.xs[p[0]],
					cr.xs[p[1]],
					cr.xs[p[2]],
					cr.xs[p[3]],
					cr.xs[p[4]],
				))
		}
	}
	xs = append(xs, "straight:")
	xs = append(xs, strings.Join(straight, ","))

	/*
			var flash []string
			for _, p := range cr.suits {
				flash = append(flash, fmt.Sprintf("%v", p[0]))
					fmt.Sprintf("%v %v %v %v %v",
						cr.xs[p[0]],
						cr.xs[p[1]],
						cr.xs[p[2]],
						cr.xs[p[3]],
						cr.xs[p[4]],
					))
			}
		xs = append(xs, "flash:")
		xs = append(xs, strings.Join(flash, ","))
	*/

	return strings.Join(xs, "\n")
}

func CalcHand(xs Deck) PokerHandDiscriptor {
	MakeCardRanking(xs)
	return PokerHandDiscriptor{}
}
