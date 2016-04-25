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
	cards     [][]Index //rank, suite, index
	highcards []Index
	pairs     [][][]Index // ranks nth (index, index)
	threes    [][][]Index // ranks nth (index, index, index)
	fours     [][][]Index // ranks nth (index, index, index, index)
	straight  [][]Index
}

func prepareCardRanking(d Deck) CardRanking {
	cr := CardRanking{
		xs:        d,
		cards:     nil,
		highcards: nil,
		pairs:     nil,
		threes:    nil,
		fours:     nil,
		straight:  nil,
	}

	cr.cards = make([][]Index, RANKS)
	for r := ACE; r < RANKS; r += 1 {
		cr.cards[r] = make([]Index, SUITS)
		for s := CLUBS; s < SUITS; s += 1 {
			cr.cards[r][s] = NullIndex
		}
	}
	for i, x := range d {
		cr.cards[x.R][x.S] = Index(i)
		if x.R == ACE {
			cr.cards[HIACE][x.S] = Index(i)
		}
	}
	return cr
}

func MakeCardRanking(xs Deck) CardRanking {
	cr := prepareCardRanking(xs)
	cr.calcPairwise()
	cr.calcStraight()
	return cr
}

func (cr *CardRanking) isSameRankWithLastItem(r Rank, xs []Index) bool {
	n := len(xs)
	if n == 0 {
		return false
	}
	return cr.xs[xs[n-1]].R == r
}

func (cr *CardRanking) calcPairwise() {
	cr.highcards = make([]Index, 0, 5)
	cr.pairs = make([][][]Index, RANKS)
	cr.threes = make([][][]Index, RANKS)
	cr.fours = make([][][]Index, RANKS)

	for r := HIACE; r > ACE; r -= 1 {
		for _, s := range SuitPermOne() {
			if !cr.isSameRankWithLastItem(r, cr.highcards) && cr.cards[r][s[0]] != NullIndex {
				cr.highcards = append(cr.highcards, cr.cards[r][s[0]])
			}
		}
		for _, s := range SuitPermTwo() {
			//fmt.Println(s)
			p := cr.cards[r][s[0]]
			q := cr.cards[r][s[1]]
			if p != NullIndex && q != NullIndex {
				if cr.pairs[r] == nil {
					cr.pairs[r] = make([][]Index, 0)
				}
				cr.pairs[r] = append(cr.pairs[r], []Index{p, q})
			}
		}
		for _, s := range SuitPermThree() {
			//fmt.Println(s)
			x := cr.cards[r][s[0]]
			y := cr.cards[r][s[1]]
			z := cr.cards[r][s[2]]
			if x != NullIndex && y != NullIndex && z != NullIndex {
				if cr.threes[r] == nil {
					cr.threes[r] = make([][]Index, 0)
				}
				cr.threes[r] = append(cr.threes[r], []Index{x, y, z})
			}
		}
		for _, s := range SuitPermFour() {
			//fmt.Println(s)
			a := cr.cards[r][s[0]]
			b := cr.cards[r][s[1]]
			c := cr.cards[r][s[2]]
			d := cr.cards[r][s[3]]
			if a != NullIndex && b != NullIndex && c != NullIndex && d != NullIndex {
				if cr.fours[r] == nil {
					cr.fours[r] = make([][]Index, 0)
				}
				cr.fours[r] = append(cr.fours[r], []Index{a, b, c, d})
			}
		}
	}
}

type FIFO5 struct {
	xs []Index
}

func newFIFO5() *FIFO5 {
	return &FIFO5{xs: make([]Index, 0)}
}

func (f *FIFO5) Push(x Index) {
	if len(f.xs) == 5 {
		copy(f.xs, f.xs[1:5])
		f.xs[4] = x
	} else {
		f.xs = append(f.xs, x)
	}
}

func (f *FIFO5) Empty() {
	f.xs = make([]Index, 0)
}

func (f *FIFO5) CloneXS() []Index {
	ys := make([]Index, len(f.xs))
	copy(ys, f.xs)
	return ys
}

func (cr *CardRanking) calcStraightSub(head Rank, f *FIFO5, filled bool) (Rank, bool) {
	var start Rank
	start = 0
	if filled {
		start = 4
	}
	for i := start; i < 5; i++ {
		found := NullIndex
		for s := CLUBS; found == NullIndex && s < SUITS; s++ {
			found = cr.cards[head-i][s]
		}

		if found != NullIndex {
			f.Push(found)
		} else {
			f.Empty()
			return head - i - 1, false
		}
	}
	return head - 1, true

}

func (cr *CardRanking) calcStraight() {
	var next Rank
	f := newFIFO5()
	cr.straight = make([][]Index, RANKS)
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

func (cr *CardRanking) findFlashOf(s Suit) (bool, []Index) {
	xs := make([]Index, 0)
	for r := HIACE; r >= DUCE; r-- {
		checked := cr.cards[r][s]
		if checked == NullIndex {
			continue
		}
		xs = append(xs, checked)
		if len(xs) >= 5 {
			return true, xs
		}
	}
	return false, nil
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
	for r, px := range cr.pairs {
		for _, p := range px {
			pairs = append(pairs, fmt.Sprintf("%v: %v %v", r, cr.xs[p[0]], cr.xs[p[1]]))
		}
	}
	xs = append(xs, "pairs:")
	xs = append(xs, strings.Join(pairs, ","))

	var threes []string
	for r, px := range cr.threes {
		for _, p := range px {
			threes = append(threes, fmt.Sprintf("%v: %v %v %v", r, cr.xs[p[0]], cr.xs[p[1]], cr.xs[p[2]]))
		}
	}
	xs = append(xs, "threes:")
	xs = append(xs, strings.Join(threes, ","))

	var fours []string
	for r, px := range cr.fours {
		for _, p := range px {
			fours = append(fours, fmt.Sprintf("%v: %v %v %v %v", r, cr.xs[p[0]], cr.xs[p[1]], cr.xs[p[2]], cr.xs[p[3]]))
		}
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

	var flash []string
	for _, s := range SuitPermOne() {
		if found, p := cr.findFlashOf(s[0]); found {
			flash = append(flash,
				fmt.Sprintf("%v %v %v %v %v",
					cr.xs[p[0]],
					cr.xs[p[1]],
					cr.xs[p[2]],
					cr.xs[p[3]],
					cr.xs[p[4]],
				))
		}
	}
	xs = append(xs, "flash:")
	xs = append(xs, strings.Join(flash, ","))

	return strings.Join(xs, "\n")
}

func CalcHand(xs Deck) PokerHandDiscriptor {
	MakeCardRanking(xs)
	return PokerHandDiscriptor{}
}
