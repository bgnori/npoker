package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Suit int

const (
	CLUBS Suit = iota
	DIAMONDS
	HEARTS
	SPADES
	SUITS
)

// Spades
//U+2660 //&spades;

//Hearts
//U+2665 //&hearts

//Diamonds
//U+2666 //&diams;

//Clubs
//U+2663 //&clubs;
var suits = [...]string{
	"\u2663",
	"\u2666",
	"\u2665",
	"\u2660",
}

func (s Suit) String() string {
	return suits[s]
}

type Rank int

const (
	_ Rank = iota
	ACE
	DUCE
	TREY
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	HIACE
	RANKS
)

var ranks = "0A23456789TJQKA"

func (r Rank) String() string {
	return fmt.Sprintf("%c", ranks[r])
}

type Card struct {
	R Rank
	S Suit
}

func (c Card) String() string {
	return fmt.Sprintf("%v%v", c.R, c.S)
}

func (c Card) Ord() int {
	//2 <= x <=53
	x := 0
	switch c.R {
	case ACE:
		x += 14
	default:
		x += int(c.R)
	}
	return x + int(c.S)*int(KING)
}

func SmokeCard() {
	fmt.Println("Suits:", SPADES, HEARTS, DIAMONDS, CLUBS)
	fmt.Println("Ranks:", ACE, DUCE, TEN, JACK, QUEEN, KING)
	fmt.Println("Card:", Card{ACE, HEARTS})
}

func parse(s string) Card {
	i := strings.Index(ranks, s[0:1])
	j := strings.Index("cdhs", s[1:2])
	return Card{Rank(i), Suit(j)}
}

func SmokeParseWith(input string) {
	got := parse(input)
	fmt.Println(input, "==>", got)
}
func SmokeParse() {
	SmokeParseWith("Ah")
	SmokeParseWith("5c")
	//SmokeParseWith("Xc")
}

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

func SmokeDeck() {
	d := Deck{Card{ACE, SPADES}, Card{FOUR, HEARTS}}
	fmt.Println(d)
	e := BuildFullDeck()
	fmt.Println(e)
	fmt.Println(e.Length())
	e.Shuffle()
	fmt.Println(e)
	e.Shuffle()
	fmt.Println(e)
	f := BuildFullDeck()
	for i, c := range []Card(*f) {
		fmt.Println(i, c, c.Ord())
	}
	f.Drop(Card{ACE, HEARTS})
	fmt.Println(f.Length())
	fmt.Println(f)
	f.Shuffle()
	fmt.Println(f)

	g := Deck{Card{ACE, SPADES}, Card{ACE, CLUBS}}
	h := Deck{Card{KING, HEARTS}, Card{KING, DIAMONDS}}
	fmt.Println(Join(g, h))
}

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
		if len(cr.ranks[i]) == 3 {
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
					fmt.Println("Straight", i, j)
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

func CalcHand(xs Deck) PokerHandDiscriptor {
	MakeCardRanking(xs)
	return PokerHandDiscriptor{}
}

func SmokeCardRanking() {
	d := Deck{
		Card{TEN, CLUBS},
		Card{EIGHT, HEARTS},
		Card{JACK, DIAMONDS},
		Card{SEVEN, HEARTS},
		Card{NINE, SPADES},
		Card{ACE, SPADES},
		Card{TEN, HEARTS},
		Card{ACE, HEARTS},
		Card{ACE, CLUBS},
		Card{KING, CLUBS},
		Card{KING, SPADES},
		Card{KING, DIAMONDS},
		Card{ACE, DIAMONDS},
		Card{DUCE, HEARTS},
	}
	fmt.Printf("%v\n", d)
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)
}

func SmokeHand() {
	SmokeCardRanking()
}

func main() {
	SmokeCard()
	SmokeParse()
	SmokeDeck()
	SmokeHand()
}
