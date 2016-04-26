package npoker

import (
	"fmt"
	"testing"
)

func TestFIFO5(t *testing.T) {
	f := newFIFO5()
	f.Push(1)
	f.Push(2)
	f.Push(3)
	f.Push(4)
	f.Push(5)
	if len(f.xs) != 5 {
		t.Error("f.xs should be {1, 2, 3, 4, 5}, but got %+v", f.xs)
	}
	f.Push(6)
	if len(f.xs) != 5 {
		t.Error("f.xs should be {1, 2, 3, 4, 5}, but got %+v", f.xs)
	}
	if f.xs[0] != 2 {
		t.Error("f.xs should be {2, 3, 4, 5, 6}, but got %+v", f.xs)
	}
	if f.xs[1] != 3 {
		t.Error("f.xs should be {2, 3, 4, 5, 6}, but got %+v", f.xs)
	}
	if f.xs[2] != 4 {
		t.Error("f.xs should be {2, 3, 4, 5, 6}, but got %+v", f.xs)
	}
	if f.xs[3] != 5 {
		t.Error("f.xs should be {2, 3, 4, 5, 6}, but got %+v", f.xs)
	}
	if f.xs[4] != 6 {
		t.Error("f.xs should be {2, 3, 4, 5, 6}, but got %+v", f.xs)
	}
}

func TestNew(t *testing.T) {
	d := Deck{}
	cr := prepareCardRanking(d)
	if cr.xs == nil {
		t.Error("got nil")
	}
	if cr.highcards != nil {
		t.Error("got non-nil")
	}
	if cr.pairs != nil {
		t.Error("got non-nil")
	}
	if cr.threes != nil {
		t.Error("got non-nil")
	}
	if cr.fours != nil {
		t.Error("got non-nil")
	}
	if cr.straight != nil {
		t.Error("got non-nil")
	}
}

func TestCalcSuit(t *testing.T) {
	d := Deck{}
	cr := prepareCardRanking(d)
	if cr.xs == nil {
		t.Error("got nil")
	}
	if cr.highcards != nil {
		t.Error("got non-nil")
	}
	if cr.pairs != nil {
		t.Error("got non-nil")
	}
	if cr.threes != nil {
		t.Error("got non-nil")
	}
	if cr.fours != nil {
		t.Error("got non-nil")
	}
	if cr.straight != nil {
		t.Error("got non-nil")
	}

}

func TestCalcPairwise(t *testing.T) {
	d := Deck{}
	cr := prepareCardRanking(d)
	cr.calcPairwise()
	if cr.xs == nil {
		t.Error("got nil")
	}
	if cr.highcards == nil {
		t.Error("got nil")
	}
	if cr.pairs == nil {
		t.Error("got nil")
	}
	if cr.threes == nil {
		t.Error("got nil")
	}
	if cr.fours == nil {
		t.Error("got nil")
	}
	if cr.straight != nil {
		t.Error("got non-nil")
	}
}

func TestCalcStraight(t *testing.T) {
	d := Deck{}
	cr := prepareCardRanking(d)
	cr.calcPairwise()
	cr.calcStraight()
	if cr.xs == nil {
		t.Error("got nil")
	}
	if cr.highcards == nil {
		t.Error("got nil")
	}
	if cr.pairs == nil {
		t.Error("got nil")
	}
	if cr.threes == nil {
		t.Error("got nil")
	}
	if cr.fours == nil {
		t.Error("got nil")
	}
	if cr.straight == nil {
		t.Error("got nil for straight")
	}
}

func TestCalcStraightDeck001(t *testing.T) {
	d := Deck{
		Card{ACE, CLUBS},
		Card{TEN, DIAMONDS},
		Card{JACK, HEARTS},
		Card{KING, SPADES},
		Card{NINE, CLUBS},
		Card{EIGHT, DIAMONDS},
		Card{QUEEN, HEARTS},
	}
	cr := prepareCardRanking(d)
	cr.calcPairwise()
	cr.calcStraight()
	if cr.straight == nil {
		t.Error("got nil for straight")
	}
	if len(cr.straight) != int(RANKS) {
		t.Errorf("should be 15 entries, but %d with %+v", len(cr.straight), cr.straight)
	}
	if len(cr.straight[int(HIACE-HIACE)]) != 5 {
		t.Error("straight entry must have 5 index, but %+v", cr.straight)
	}
	if len(cr.straight[int(HIACE-KING)]) != 5 {
		t.Error("straight entry must have 5 index, but %+v", cr.straight)
	}
	if len(cr.straight[int(HIACE-QUEEN)]) != 5 {
		t.Error("straight entry must have 5 index, but %+v", cr.straight)
	}
}

func TestCalcFlushDeck001(t *testing.T) {
	d := Deck{
		Card{ACE, SPADES},
		Card{TEN, SPADES},
		Card{JACK, SPADES},
		Card{KING, SPADES},
		Card{NINE, SPADES},
		Card{EIGHT, SPADES},
		Card{SIX, SPADES},
	}
	cr := prepareCardRanking(d)
	cr.calcPairwise()
	cr.calcStraight()
}

func TestCalc(t *testing.T) {
}

func TestSampleHand001(t *testing.T) {
	d := Deck{
		Card{FOUR, CLUBS},
		Card{TREY, CLUBS},
		Card{NINE, SPADES},
		Card{TEN, SPADES},
		Card{SIX, CLUBS},
		Card{DUCE, DIAMONDS},
		Card{QUEEN, HEARTS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// nothing with QUEEN, TEN, NINE, SIX, FOUR high,
}

func TestSampleHand002(t *testing.T) {
	d := Deck{
		Card{SEVEN, DIAMONDS},
		Card{ACE, HEARTS},
		Card{FOUR, HEARTS},
		Card{TEN, SPADES},
		Card{NINE, CLUBS},
		Card{TREY, DIAMONDS},
		Card{JACK, DIAMONDS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// nothing with ACE, JACK, TEN, NINE, SEVEN high,
}

func TestSampleHand003(t *testing.T) {
	d := Deck{
		Card{EIGHT, SPADES},
		Card{ACE, DIAMONDS},
		Card{TREY, HEARTS},
		Card{SIX, DIAMONDS},
		Card{SIX, CLUBS},
		Card{JACK, CLUBS},
		Card{SEVEN, CLUBS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Pair of Six with ACE, JACK, EIGHT, high,
}

func TestSampleHand004(t *testing.T) {
	d := Deck{
		Card{SEVEN, CLUBS},
		Card{TEN, SPADES},
		Card{TEN, HEARTS},
		Card{EIGHT, SPADES},
		Card{FIVE, SPADES},
		Card{ACE, DIAMONDS},
		Card{SIX, SPADES},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Pair of Ten with ACE, EIGHT, SEVEN high
}

func TestSampleHand005(t *testing.T) {
	d := Deck{
		Card{JACK, CLUBS},
		Card{DUCE, CLUBS},
		Card{NINE, SPADES},
		Card{TEN, SPADES},
		Card{TREY, CLUBS},
		Card{EIGHT, CLUBS},
		Card{FOUR, SPADES},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// nothing with JACK, TEN, NINE, EIGHT, FOUR high,
}

func TestSampleHand006(t *testing.T) {
	d := Deck{
		Card{SIX, HEARTS},
		Card{SIX, CLUBS},
		Card{ACE, DIAMONDS},
		Card{FIVE, SPADES},
		Card{KING, HEARTS},
		Card{TEN, HEARTS},
		Card{SEVEN, CLUBS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// pair of SIX with ACE, KING, TEN high
}

func TestSampleHand007(t *testing.T) {
	d := Deck{
		Card{FOUR, HEARTS},
		Card{KING, SPADES},
		Card{QUEEN, SPADES},
		Card{TEN, HEARTS},
		Card{KING, CLUBS},
		Card{FOUR, SPADES},
		Card{SEVEN, HEARTS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Pairs of KING and FOUR with QUEEN high,
}

func TestSampleHand008(t *testing.T) {
	d := Deck{
		Card{EIGHT, SPADES},
		Card{KING, SPADES},
		Card{KING, HEARTS},
		Card{FOUR, SPADES},
		Card{SIX, SPADES},
		Card{JACK, DIAMONDS},
		Card{TREY, DIAMONDS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// pair of KING with JACK, EIGHT, SIX high
}

func TestSampleHand009(t *testing.T) {
	d := Deck{
		Card{EIGHT, SPADES},
		Card{EIGHT, DIAMONDS},
		Card{TEN, DIAMONDS},
		Card{SEVEN, CLUBS},
		Card{TEN, CLUBS},
		Card{FOUR, DIAMONDS},
		Card{JACK, DIAMONDS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Pairs of TEN and EIGHT with JACK high
}

func TestSampleHand010(t *testing.T) {
	d := Deck{
		Card{SEVEN, SPADES},
		Card{FIVE, SPADES},
		Card{FIVE, HEARTS},
		Card{NINE, CLUBS},
		Card{TREY, DIAMONDS},
		Card{FIVE, CLUBS},
		Card{TREY, CLUBS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// FullHouse, FIVE over TREY
}

func TestSampleHand011(t *testing.T) {
	d := Deck{
		Card{TEN, HEARTS},
		Card{JACK, SPADES},
		Card{JACK, DIAMONDS},
		Card{SEVEN, DIAMONDS},
		Card{JACK, HEARTS},
		Card{FOUR, HEARTS},
		Card{FOUR, CLUBS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// FullHouse, JACK over FOUR
}

func TestSampleHand012(t *testing.T) {
	d := Deck{
		Card{KING, DIAMONDS},
		Card{TREY, CLUBS},
		Card{TEN, SPADES},
		Card{ACE, DIAMONDS},
		Card{QUEEN, DIAMONDS},
		Card{JACK, DIAMONDS},
		Card{TEN, DIAMONDS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Straight Flush with ACE HIGH,
}

func TestSampleHand013(t *testing.T) {
	d := Deck{
		Card{JACK, CLUBS},
		Card{EIGHT, SPADES},
		Card{QUEEN, SPADES},
		Card{FIVE, CLUBS},
		Card{JACK, DIAMONDS},
		Card{NINE, HEARTS},
		Card{FIVE, HEARTS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Pairs of JACK and FIVE with QUEEN high
}

func TestSampleHand014(t *testing.T) {
	d := Deck{
		Card{DUCE, DIAMONDS},
		Card{ACE, SPADES},
		Card{SEVEN, HEARTS},
		Card{TEN, DIAMONDS},
		Card{ACE, DIAMONDS},
		Card{KING, SPADES},
		Card{KING, CLUBS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Pairs of ACE and KING with TEN high
}

func TestSampleHand015(t *testing.T) {
	d := Deck{
		Card{SIX, CLUBS},
		Card{SEVEN, DIAMONDS},
		Card{SEVEN, CLUBS},
		Card{QUEEN, CLUBS},
		Card{DUCE, CLUBS},
		Card{DUCE, SPADES},
		Card{SEVEN, SPADES},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// FullHouse, SEVEN over DUCE
}

func TestSampleHand016(t *testing.T) {
	d := Deck{
		Card{ACE, CLUBS},
		Card{SIX, HEARTS},
		Card{ACE, SPADES},
		Card{NINE, DIAMONDS},
		Card{SEVEN, SPADES},
		Card{SIX, SPADES},
		Card{QUEEN, DIAMONDS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Pairs of ACE and SIX with QUEEN high
}

func TestSampleHand017(t *testing.T) {
	d := Deck{
		Card{QUEEN, DIAMONDS},
		Card{ACE, SPADES},
		Card{TEN, SPADES},
		Card{JACK, SPADES},
		Card{JACK, HEARTS},
		Card{KING, SPADES},
		Card{QUEEN, SPADES},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// straight  with ACE High
}

func TestSampleHand018(t *testing.T) {
	d := Deck{
		Card{JACK, HEARTS},
		Card{KING, CLUBS},
		Card{ACE, CLUBS},
		Card{TEN, CLUBS},
		Card{QUEEN, CLUBS},
		Card{SIX, DIAMONDS},
		Card{JACK, CLUBS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// straight  with ACE High
}

func TestSampleHand019(t *testing.T) {
	d := Deck{
		Card{QUEEN, CLUBS},
		Card{TREY, DIAMONDS},
		Card{NINE, CLUBS},
		Card{FIVE, CLUBS},
		Card{FOUR, CLUBS},
		Card{TEN, CLUBS},
		Card{DUCE, CLUBS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Flush with Queen, TEN, NINE, FIVE, FOUR high
}

func TestSampleHand020(t *testing.T) {
	d := Deck{
		Card{ACE, SPADES},
		Card{JACK, SPADES},
		Card{DUCE, SPADES},
		Card{SIX, SPADES},
		Card{EIGHT, HEARTS},
		Card{EIGHT, DIAMONDS},
		Card{QUEEN, SPADES},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Flush with ACE, QUEEN, JACK, SIX, DUCE high
}

func TestSampleHand021(t *testing.T) {
	d := Deck{
		Card{TEN, SPADES},
		Card{EIGHT, SPADES},
		Card{DUCE, SPADES},
		Card{SIX, SPADES},
		Card{EIGHT, HEARTS},
		Card{EIGHT, DIAMONDS},
		Card{QUEEN, SPADES},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Flush with JACK, TEN, EIGHT, SIX, DUCE high
}

func TestSampleHand022(t *testing.T) {
	d := Deck{
		Card{SIX, HEARTS},
		Card{SIX, DIAMONDS},
		Card{DUCE, SPADES},
		Card{SIX, SPADES},
		Card{EIGHT, HEARTS},
		Card{EIGHT, DIAMONDS},
		Card{QUEEN, SPADES},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	//fmt.Printf("debug: %+v\n", phd.which)
	fmt.Printf("%v\n", phd)
	// FullHouse, SIX over EIGHT
}

func TestSampleHand023(t *testing.T) {
	d := Deck{
		Card{ACE, DIAMONDS},
		Card{ACE, HEARTS},
		Card{FOUR, DIAMONDS},
		Card{SIX, CLUBS},
		Card{EIGHT, HEARTS},
		Card{QUEEN, DIAMONDS},
		Card{ACE, SPADES},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("debug: %+v\n", phd.which)
	fmt.Printf("%v\n", phd)
	// Three of Kind ACE with QUEEN, EIGHT high
}

func TestSampleHand024(t *testing.T) {
	d := Deck{
		Card{FIVE, CLUBS},
		Card{SEVEN, SPADES},
		Card{FOUR, DIAMONDS},
		Card{SIX, CLUBS},
		Card{EIGHT, HEARTS},
		Card{QUEEN, DIAMONDS},
		Card{ACE, SPADES},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)

	// straight, EIGHT high
}

func TestSampleHand025(t *testing.T) {
	d := Deck{
		Card{JACK, SPADES},
		Card{JACK, CLUBS},
		Card{JACK, DIAMONDS},
		Card{TREY, SPADES},
		Card{QUEEN, CLUBS},
		Card{QUEEN, DIAMONDS},
		Card{DUCE, CLUBS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	// FullHouse, Jack over Queen,
}

func TestSampleHand026(t *testing.T) {
	d := Deck{
		Card{ACE, SPADES},
		Card{ACE, HEARTS},
		Card{JACK, DIAMONDS},
		Card{TREY, SPADES},
		Card{QUEEN, CLUBS},
		Card{QUEEN, DIAMONDS},
		Card{DUCE, CLUBS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	// Two pairs, Aces, Queens, with JACK High
}

func TestSampleHand027(t *testing.T) {
	d := Deck{
		Card{KING, DIAMONDS},
		Card{ACE, DIAMONDS},
		Card{JACK, DIAMONDS},
		Card{TREY, SPADES},
		Card{QUEEN, CLUBS},
		Card{QUEEN, DIAMONDS},
		Card{DUCE, CLUBS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	// A pair , with Ace King, Jack High
}

func TestSampleHand028(t *testing.T) {
	d := Deck{
		Card{KING, DIAMONDS},
		Card{ACE, DIAMONDS},
		Card{JACK, DIAMONDS},
		Card{TREY, SPADES},
		Card{QUEEN, CLUBS},
		Card{QUEEN, DIAMONDS},
		Card{TEN, DIAMONDS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Straight Flush
}

func TestSampleHand029(t *testing.T) {
	d := Deck{
		Card{ACE, CLUBS},
		Card{ACE, HEARTS},
		Card{QUEEN, DIAMONDS},
		Card{TEN, HEARTS},
		Card{ACE, SPADES},
		Card{ACE, DIAMONDS},
		Card{FOUR, HEARTS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Quad
}

func TestSampleHand000(t *testing.T) {
	d := Deck{
		Card{EIGHT, SPADES},
		Card{ACE, DIAMONDS},
		Card{TREY, HEARTS},
		Card{SIX, DIAMONDS},
		Card{SIX, CLUBS},
		Card{JACK, CLUBS},
		Card{SEVEN, CLUBS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	// nothing with Queen high,
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

func SmokeOnePair() {
	fmt.Println("SmokeOnePair")
	d := Deck{
		Card{ACE, CLUBS},
		Card{ACE, HEARTS},
		Card{QUEEN, DIAMONDS},
		Card{TEN, HEARTS},
		Card{EIGHT, SPADES},
		Card{FIVE, DIAMONDS},
		Card{FOUR, HEARTS},
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)
}

func SmokeTwoPair() {
	fmt.Println("SmokeTwoPair")
	d := Deck{}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)
}

func SmokeThreeOfAKind() {
	d := Deck{
		Card{ACE, CLUBS},
		Card{ACE, HEARTS},
		Card{QUEEN, DIAMONDS},
		Card{TEN, HEARTS},
		Card{EIGHT, SPADES},
		Card{ACE, DIAMONDS},
		Card{FOUR, HEARTS},
	}
	fmt.Println("SmokeThreeOfAKind")
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)
}

func SmokeFourOfAKind() {
	d := Deck{
		Card{ACE, CLUBS},
		Card{ACE, HEARTS},
		Card{QUEEN, DIAMONDS},
		Card{TEN, HEARTS},
		Card{ACE, SPADES},
		Card{ACE, DIAMONDS},
		Card{FOUR, HEARTS},
	}
	fmt.Println("SmokeFourOfAKind")
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)
}

func SmokeHand() {
	SmokeCardRanking()
	SmokeOnePair()
	SmokeTwoPair()
	SmokeThreeOfAKind()
	SmokeFourOfAKind()
}
