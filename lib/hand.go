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
	highcards []int
	pairs     [][]int
	threes    [][]int
	fours     [][]int
	ranks     [][]int
	suits     [][]int
	straight  [][]int
}

func appendrank(ranks [][]int, i int, r Rank) [][]int {
	xs := ranks[r]
	if xs == nil {
		xs = make([]int, 0)
	}
	ranks[r] = append(xs, i)
	return ranks
}

func MakeCardRanking(xs Deck) CardRanking {
	cr := CardRanking{
		xs,
		make([]int, 0, 5),
		make([][]int, 0, 5),
		make([][]int, 0, 5),
		make([][]int, 0, 5),
		make([][]int, RANKS),
		make([][]int, SUITS),
		make([][]int, RANKS),
	}

	for i, x := range xs {
		cr.ranks = appendrank(cr.ranks, i, x.R)
		if x.R == ACE {
			cr.ranks = appendrank(cr.ranks, i, HIACE)
		}
	}

	for i, x := range xs {
		xs := cr.suits[x.S]
		if xs == nil {
			xs = make([]int, 0)
		}
		cr.suits[x.S] = append(xs, i)
	}

	for i := HIACE; i > ACE; i -= 1 {
		fmt.Println(i, len(cr.ranks[i]))
		if len(cr.ranks[i]) == 1 {
			cr.highcards = append(cr.highcards, cr.ranks[i][0])
		}
		if len(cr.ranks[i]) == 2 {
			cr.pairs = append(cr.pairs, []int{cr.ranks[i][0], cr.ranks[i][1]})
		}
		if len(cr.ranks[i]) == 3 {
			cr.threes = append(cr.threes, []int{cr.ranks[i][0], cr.ranks[i][1], cr.ranks[i][2]})
		}
		if len(cr.ranks[i]) == 4 {
			cr.fours = append(cr.fours, []int{cr.ranks[i][0], cr.ranks[i][1], cr.ranks[i][2], cr.ranks[i][3]})
		}
	}

	for i := HIACE; i > FOUR; i -= 1 {
		j := i
		if len(cr.ranks[i]) > 0 {
			for ; j > FOUR && len(cr.ranks[j]) > 0; j -= 1 {
				if i-4 == j {
					xs := cr.straight[i]
					if xs == nil {
						xs = make([]int, 0)
					}
					cr.straight[i] = append(xs, cr.ranks[i][0])
					i -= 1
					continue
				}
			}
			i = j
		}
	}
	return cr
}

func (cr CardRanking) String() string {
	var xs []string

	xs = append(xs, fmt.Sprintf("%v", cr.xs))

	var highcards []string
	for _, pos := range cr.highcards {
		highcards = append(highcards, fmt.Sprintf("%v", cr.xs[pos]))
	}
	xs = append(xs, strings.Join(highcards, ","))

	var pairs []string
	for _, p := range cr.pairs {
		pairs = append(pairs, fmt.Sprintf("%v %v", cr.xs[p[0]], cr.xs[p[1]]))
	}
	xs = append(xs, strings.Join(pairs, ","))

	var threes []string
	for _, p := range cr.threes {
		threes = append(threes, fmt.Sprintf("%v %v %v", cr.xs[p[0]], cr.xs[p[1]], cr.xs[p[2]]))
	}
	xs = append(xs, strings.Join(threes, ","))

	var fours []string
	for _, p := range cr.fours {
		fours = append(fours, fmt.Sprintf("%v %v %v %v", cr.xs[p[0]], cr.xs[p[1]], cr.xs[p[2]], cr.xs[p[3]]))
	}
	xs = append(xs, strings.Join(fours, ","))

	/*
			ranks     [][]int
			suits     [][]int
			straight  [][]int
		var straight []string
		for _, p := range cr.straight {
			straight = append(straight, fmt.Sprintf("%v", cr.xs[p[0]]))
			//, cr.xs[p[1]], cr.xs[p[2]], cr.xs[p[3]], cr.xs[p[4]]))
		}
		xs = append(xs, strings.Join(straight, ","))
	*/

	return strings.Join(xs, "\n")
}

func CalcHand(xs Deck) PokerHandDiscriptor {
	MakeCardRanking(xs)
	return PokerHandDiscriptor{}
}
