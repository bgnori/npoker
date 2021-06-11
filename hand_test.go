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
		t.Errorf("f.xs should be {1, 2, 3, 4, 5}, but got %+v", f.xs)
	}
	f.Push(6)
	if len(f.xs) != 5 {
		t.Errorf("f.xs should be {1, 2, 3, 4, 5}, but got %+v", f.xs)
	}
	if f.xs[0] != 2 {
		t.Errorf("f.xs should be {2, 3, 4, 5, 6}, but got %+v", f.xs)
	}
	if f.xs[1] != 3 {
		t.Errorf("f.xs should be {2, 3, 4, 5, 6}, but got %+v", f.xs)
	}
	if f.xs[2] != 4 {
		t.Errorf("f.xs should be {2, 3, 4, 5, 6}, but got %+v", f.xs)
	}
	if f.xs[3] != 5 {
		t.Errorf("f.xs should be {2, 3, 4, 5, 6}, but got %+v", f.xs)
	}
	if f.xs[4] != 6 {
		t.Errorf("f.xs should be {2, 3, 4, 5, 6}, but got %+v", f.xs)
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
		NewCard(ACE, CLUBS),
		NewCard(TEN, DIAMONDS),
		NewCard(JACK, HEARTS),
		NewCard(KING, SPADES),
		NewCard(NINE, CLUBS),
		NewCard(EIGHT, DIAMONDS),
		NewCard(QUEEN, HEARTS),
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
		t.Errorf("straight entry must have 5 index, but %+v", cr.straight)
	}
	if len(cr.straight[int(HIACE-KING)]) != 5 {
		t.Errorf("straight entry must have 5 index, but %+v", cr.straight)
	}
	if len(cr.straight[int(HIACE-QUEEN)]) != 5 {
		t.Errorf("straight entry must have 5 index, but %+v", cr.straight)
	}
}

func TestCalcFlushDeck001(t *testing.T) {
	d := Deck{
		NewCard(ACE, SPADES),
		NewCard(TEN, SPADES),
		NewCard(JACK, SPADES),
		NewCard(KING, SPADES),
		NewCard(NINE, SPADES),
		NewCard(EIGHT, SPADES),
		NewCard(SIX, SPADES),
	}
	cr := prepareCardRanking(d)
	cr.calcPairwise()
	cr.calcStraight()
}

func TestCalc(t *testing.T) {
}

func TestSampleHand001(t *testing.T) {
	d := Deck{
		NewCard(FOUR, CLUBS),
		NewCard(TREY, CLUBS),
		NewCard(NINE, SPADES),
		NewCard(TEN, SPADES),
		NewCard(SIX, CLUBS),
		NewCard(DUCE, DIAMONDS),
		NewCard(QUEEN, HEARTS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%+v\n", phd)
	// nothing with QUEEN, TEN, NINE, SIX, FOUR high,
}

func TestSampleHand002(t *testing.T) {
	d := Deck{
		NewCard(SEVEN, DIAMONDS),
		NewCard(ACE, HEARTS),
		NewCard(FOUR, HEARTS),
		NewCard(TEN, SPADES),
		NewCard(NINE, CLUBS),
		NewCard(TREY, DIAMONDS),
		NewCard(JACK, DIAMONDS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// nothing with ACE, JACK, TEN, NINE, SEVEN high,
}

func TestSampleHand003(t *testing.T) {
	d := Deck{
		NewCard(EIGHT, SPADES),
		NewCard(ACE, DIAMONDS),
		NewCard(TREY, HEARTS),
		NewCard(SIX, DIAMONDS),
		NewCard(SIX, CLUBS),
		NewCard(JACK, CLUBS),
		NewCard(SEVEN, CLUBS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Pair of Six with ACE, JACK, EIGHT, high,
}

func TestSampleHand004(t *testing.T) {
	d := Deck{
		NewCard(SEVEN, CLUBS),
		NewCard(TEN, SPADES),
		NewCard(TEN, HEARTS),
		NewCard(EIGHT, SPADES),
		NewCard(FIVE, SPADES),
		NewCard(ACE, DIAMONDS),
		NewCard(SIX, SPADES),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Pair of Ten with ACE, EIGHT, SEVEN high
}

func TestSampleHand005(t *testing.T) {
	d := Deck{
		NewCard(JACK, CLUBS),
		NewCard(DUCE, CLUBS),
		NewCard(NINE, SPADES),
		NewCard(TEN, SPADES),
		NewCard(TREY, CLUBS),
		NewCard(EIGHT, CLUBS),
		NewCard(FOUR, SPADES),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// nothing with JACK, TEN, NINE, EIGHT, FOUR high,
}

func TestSampleHand006(t *testing.T) {
	d := Deck{
		NewCard(SIX, HEARTS),
		NewCard(SIX, CLUBS),
		NewCard(ACE, DIAMONDS),
		NewCard(FIVE, SPADES),
		NewCard(KING, HEARTS),
		NewCard(TEN, HEARTS),
		NewCard(SEVEN, CLUBS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// pair of SIX with ACE, KING, TEN high
}

func TestSampleHand007(t *testing.T) {
	d := Deck{
		NewCard(FOUR, HEARTS),
		NewCard(KING, SPADES),
		NewCard(QUEEN, SPADES),
		NewCard(TEN, HEARTS),
		NewCard(KING, CLUBS),
		NewCard(FOUR, SPADES),
		NewCard(SEVEN, HEARTS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Pairs of KING and FOUR with QUEEN high,
}

func TestSampleHand008(t *testing.T) {
	d := Deck{
		NewCard(EIGHT, SPADES),
		NewCard(KING, SPADES),
		NewCard(KING, HEARTS),
		NewCard(FOUR, SPADES),
		NewCard(SIX, SPADES),
		NewCard(JACK, DIAMONDS),
		NewCard(TREY, DIAMONDS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// pair of KING with JACK, EIGHT, SIX high
}

func TestSampleHand009(t *testing.T) {
	d := Deck{
		NewCard(EIGHT, SPADES),
		NewCard(EIGHT, DIAMONDS),
		NewCard(TEN, DIAMONDS),
		NewCard(SEVEN, CLUBS),
		NewCard(TEN, CLUBS),
		NewCard(FOUR, DIAMONDS),
		NewCard(JACK, DIAMONDS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Pairs of TEN and EIGHT with JACK high
}

func TestSampleHand010(t *testing.T) {
	d := Deck{
		NewCard(SEVEN, SPADES),
		NewCard(FIVE, SPADES),
		NewCard(FIVE, HEARTS),
		NewCard(NINE, CLUBS),
		NewCard(TREY, DIAMONDS),
		NewCard(FIVE, CLUBS),
		NewCard(TREY, CLUBS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// FullHouse, FIVE over TREY
}

func TestSampleHand011(t *testing.T) {
	d := Deck{
		NewCard(TEN, HEARTS),
		NewCard(JACK, SPADES),
		NewCard(JACK, DIAMONDS),
		NewCard(SEVEN, DIAMONDS),
		NewCard(JACK, HEARTS),
		NewCard(FOUR, HEARTS),
		NewCard(FOUR, CLUBS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// FullHouse, JACK over FOUR
}

func TestSampleHand012(t *testing.T) {
	d := Deck{
		NewCard(KING, DIAMONDS),
		NewCard(TREY, CLUBS),
		NewCard(TEN, SPADES),
		NewCard(ACE, DIAMONDS),
		NewCard(QUEEN, DIAMONDS),
		NewCard(JACK, DIAMONDS),
		NewCard(TEN, DIAMONDS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Straight Flush with ACE HIGH,
}

func TestSampleHand013(t *testing.T) {
	d := Deck{
		NewCard(JACK, CLUBS),
		NewCard(EIGHT, SPADES),
		NewCard(QUEEN, SPADES),
		NewCard(FIVE, CLUBS),
		NewCard(JACK, DIAMONDS),
		NewCard(NINE, HEARTS),
		NewCard(FIVE, HEARTS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Pairs of JACK and FIVE with QUEEN high
}

func TestSampleHand014(t *testing.T) {
	d := Deck{
		NewCard(DUCE, DIAMONDS),
		NewCard(ACE, SPADES),
		NewCard(SEVEN, HEARTS),
		NewCard(TEN, DIAMONDS),
		NewCard(ACE, DIAMONDS),
		NewCard(KING, SPADES),
		NewCard(KING, CLUBS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Pairs of ACE and KING with TEN high
}

func TestSampleHand015(t *testing.T) {
	d := Deck{
		NewCard(SIX, CLUBS),
		NewCard(SEVEN, DIAMONDS),
		NewCard(SEVEN, CLUBS),
		NewCard(QUEEN, CLUBS),
		NewCard(DUCE, CLUBS),
		NewCard(DUCE, SPADES),
		NewCard(SEVEN, SPADES),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// FullHouse, SEVEN over DUCE
}

func TestSampleHand016(t *testing.T) {
	d := Deck{
		NewCard(ACE, CLUBS),
		NewCard(SIX, HEARTS),
		NewCard(ACE, SPADES),
		NewCard(NINE, DIAMONDS),
		NewCard(SEVEN, SPADES),
		NewCard(SIX, SPADES),
		NewCard(QUEEN, DIAMONDS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Pairs of ACE and SIX with QUEEN high
}

func TestSampleHand017(t *testing.T) {
	d := Deck{
		NewCard(QUEEN, DIAMONDS),
		NewCard(ACE, SPADES),
		NewCard(TEN, SPADES),
		NewCard(JACK, SPADES),
		NewCard(JACK, HEARTS),
		NewCard(KING, SPADES),
		NewCard(QUEEN, SPADES),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// straight  with ACE High
}

func TestSampleHand018(t *testing.T) {
	d := Deck{
		NewCard(JACK, HEARTS),
		NewCard(KING, CLUBS),
		NewCard(ACE, CLUBS),
		NewCard(TEN, CLUBS),
		NewCard(QUEEN, CLUBS),
		NewCard(SIX, DIAMONDS),
		NewCard(JACK, CLUBS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// straight  with ACE High
}

func TestSampleHand019(t *testing.T) {
	d := Deck{
		NewCard(QUEEN, CLUBS),
		NewCard(TREY, DIAMONDS),
		NewCard(NINE, CLUBS),
		NewCard(FIVE, CLUBS),
		NewCard(FOUR, CLUBS),
		NewCard(TEN, CLUBS),
		NewCard(DUCE, CLUBS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Flush with Queen, TEN, NINE, FIVE, FOUR high
}

func TestSampleHand020(t *testing.T) {
	d := Deck{
		NewCard(ACE, SPADES),
		NewCard(JACK, SPADES),
		NewCard(DUCE, SPADES),
		NewCard(SIX, SPADES),
		NewCard(EIGHT, HEARTS),
		NewCard(EIGHT, DIAMONDS),
		NewCard(QUEEN, SPADES),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Flush with ACE, QUEEN, JACK, SIX, DUCE high
}

func TestSampleHand021(t *testing.T) {
	d := Deck{
		NewCard(TEN, SPADES),
		NewCard(EIGHT, SPADES),
		NewCard(DUCE, SPADES),
		NewCard(SIX, SPADES),
		NewCard(EIGHT, HEARTS),
		NewCard(EIGHT, DIAMONDS),
		NewCard(QUEEN, SPADES),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Flush with JACK, TEN, EIGHT, SIX, DUCE high
}

func TestSampleHand022(t *testing.T) {
	d := Deck{
		NewCard(SIX, HEARTS),
		NewCard(SIX, DIAMONDS),
		NewCard(DUCE, SPADES),
		NewCard(SIX, SPADES),
		NewCard(EIGHT, HEARTS),
		NewCard(EIGHT, DIAMONDS),
		NewCard(QUEEN, SPADES),
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
		NewCard(ACE, DIAMONDS),
		NewCard(ACE, HEARTS),
		NewCard(FOUR, DIAMONDS),
		NewCard(SIX, CLUBS),
		NewCard(EIGHT, HEARTS),
		NewCard(QUEEN, DIAMONDS),
		NewCard(ACE, SPADES),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	ph := CalcHand(d)
	fmt.Printf("%v\n", ph)
	// Three of Kind ACE with QUEEN, EIGHT high
}

func TestSampleHand024(t *testing.T) {
	d := Deck{
		NewCard(FIVE, CLUBS),
		NewCard(SEVEN, SPADES),
		NewCard(FOUR, DIAMONDS),
		NewCard(SIX, CLUBS),
		NewCard(EIGHT, HEARTS),
		NewCard(QUEEN, DIAMONDS),
		NewCard(ACE, SPADES),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)

	// straight, EIGHT high
}

func TestSampleHand025(t *testing.T) {
	d := Deck{
		NewCard(JACK, SPADES),
		NewCard(JACK, CLUBS),
		NewCard(JACK, DIAMONDS),
		NewCard(TREY, SPADES),
		NewCard(QUEEN, CLUBS),
		NewCard(QUEEN, DIAMONDS),
		NewCard(DUCE, CLUBS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// FullHouse, Jack over Queen,
}

func TestSampleHand026(t *testing.T) {
	d := Deck{
		NewCard(ACE, SPADES),
		NewCard(ACE, HEARTS),
		NewCard(JACK, DIAMONDS),
		NewCard(TREY, SPADES),
		NewCard(QUEEN, CLUBS),
		NewCard(QUEEN, DIAMONDS),
		NewCard(DUCE, CLUBS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	// Two pairs, Aces, Queens, with JACK High
}

func TestSampleHand027(t *testing.T) {
	d := Deck{
		NewCard(KING, DIAMONDS),
		NewCard(ACE, DIAMONDS),
		NewCard(JACK, DIAMONDS),
		NewCard(TREY, SPADES),
		NewCard(QUEEN, CLUBS),
		NewCard(QUEEN, DIAMONDS),
		NewCard(DUCE, CLUBS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	// A pair , with Ace King, Jack High
	fmt.Printf("%v\n", CalcHand(d))
}

func TestSampleHand028(t *testing.T) {
	d := Deck{
		NewCard(KING, DIAMONDS),
		NewCard(ACE, DIAMONDS),
		NewCard(JACK, DIAMONDS),
		NewCard(TREY, SPADES),
		NewCard(QUEEN, CLUBS),
		NewCard(QUEEN, DIAMONDS),
		NewCard(TEN, DIAMONDS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Straight Flush
}

func TestSampleHand029(t *testing.T) {
	d := Deck{
		NewCard(ACE, CLUBS),
		NewCard(ACE, HEARTS),
		NewCard(QUEEN, DIAMONDS),
		NewCard(TEN, HEARTS),
		NewCard(ACE, SPADES),
		NewCard(ACE, DIAMONDS),
		NewCard(FOUR, HEARTS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Quad
}

func TestSampleHand030(t *testing.T) {
	d := Deck{
		NewCard(SIX, CLUBS),
		NewCard(SIX, HEARTS),
		NewCard(SIX, SPADES),
		NewCard(SIX, DIAMONDS),
		NewCard(EIGHT, CLUBS),
		NewCard(EIGHT, HEARTS),
		NewCard(EIGHT, SPADES),
		NewCard(EIGHT, DIAMONDS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Quad x 2
}

func TestSampleHand031(t *testing.T) {
	d := Deck{
		NewCard(SIX, CLUBS),
		NewCard(SIX, HEARTS),
		NewCard(SIX, SPADES),
		NewCard(EIGHT, CLUBS),
		NewCard(EIGHT, HEARTS),
		NewCard(EIGHT, SPADES),
		NewCard(FOUR, CLUBS),
		NewCard(FOUR, DIAMONDS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// FullHouse. EIGHT over SIX
}

func TestSampleHand032(t *testing.T) {
	d := Deck{
		NewCard(ACE, CLUBS),
		NewCard(DUCE, HEARTS),
		NewCard(TREY, DIAMONDS),
		NewCard(FOUR, SPADES),
		NewCard(FIVE, CLUBS),
		NewCard(SIX, HEARTS),
		NewCard(SEVEN, DIAMONDS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Straight, Seven High
}

func TestSampleHand033(t *testing.T) {
	d := Deck{
		NewCard(ACE, CLUBS),
		NewCard(DUCE, HEARTS),
		NewCard(TREY, DIAMONDS),
		NewCard(FOUR, SPADES),
		NewCard(FIVE, CLUBS),
		NewCard(KING, HEARTS),
		NewCard(QUEEN, DIAMONDS),
		NewCard(JACK, SPADES),
		NewCard(TEN, CLUBS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
	// Straight, ACE High
}

func TestSampleHand034(t *testing.T) {
	d := Deck{
		NewCard(EIGHT, SPADES),
		NewCard(EIGHT, DIAMONDS),
		NewCard(TEN, HEARTS),
		NewCard(TEN, DIAMONDS),
		NewCard(SIX, SPADES),
		NewCard(SIX, CLUBS),
		NewCard(ACE, CLUBS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	// nothing with Queen high,
	phd := CalcHand(d)
	fmt.Printf("%v\n", phd)
}

func TestSampleHand000(t *testing.T) {
	d := Deck{
		NewCard(EIGHT, SPADES),
		NewCard(ACE, DIAMONDS),
		NewCard(TREY, HEARTS),
		NewCard(SIX, DIAMONDS),
		NewCard(SIX, CLUBS),
		NewCard(JACK, CLUBS),
		NewCard(SEVEN, CLUBS),
	}
	cr := MakeCardRanking(d)
	fmt.Printf("%v\n", cr)

	// nothing with Queen high,
}

func TestSampleGame001(t *testing.T) {
	PlayerOneHole := Deck{
		NewCard(JACK, SPADES),
		NewCard(JACK, CLUBS),
	}
	// FullHouse, Jack over Queen,

	PlayerTwoHole := Deck{
		NewCard(ACE, SPADES),
		NewCard(ACE, HEARTS),
	}
	// Two pairs, Aces, Queens, with JACK High

	PlayerThreeHole := Deck{
		NewCard(KING, DIAMONDS),
		NewCard(ACE, DIAMONDS),
	}
	// A pair , with Ace King, Jack High

	board := Deck{
		NewCard(JACK, DIAMONDS),
		NewCard(TREY, SPADES),
		NewCard(QUEEN, CLUBS),
		NewCard(QUEEN, DIAMONDS),
		NewCard(DUCE, CLUBS),
	}
	sd := MakeShowDown(board, PlayerOneHole, PlayerTwoHole, PlayerThreeHole)
	fmt.Printf("%+v\n", sd)

	for i, v := range DistrubuteChips(100, 10, 0, sd) {
		fmt.Printf("player %d got %d.\n", i, v)
	}
}

func TestSampleGame002(t *testing.T) {
	PlayerOneHole := Deck{
		NewCard(SIX, SPADES),
		NewCard(SIX, CLUBS),
	}

	PlayerTwoHole := Deck{
		NewCard(FIVE, SPADES),
		NewCard(FIVE, HEARTS),
	}

	PlayerThreeHole := Deck{
		NewCard(KING, DIAMONDS),
		NewCard(NINE, DIAMONDS),
	}

	board := Deck{
		NewCard(TEN, DIAMONDS),
		NewCard(TEN, SPADES),
		NewCard(QUEEN, CLUBS),
		NewCard(QUEEN, DIAMONDS),
		NewCard(ACE, CLUBS),
	}
	sd := MakeShowDown(board, PlayerOneHole, PlayerTwoHole, PlayerThreeHole)
	fmt.Printf("%+v\n", sd)

	for i, v := range DistrubuteChips(100, 10, 0, sd) {
		fmt.Printf("player %d got %d.\n", i, v)
	}
}

func TestSampleGame003(t *testing.T) {
	PlayerOneHole := Deck{
		NewCard(SIX, SPADES),
		NewCard(SIX, CLUBS),
	}

	PlayerTwoHole := Deck{
		NewCard(ACE, SPADES),
		NewCard(SEVEN, HEARTS),
	}

	PlayerThreeHole := Deck{
		NewCard(ACE, DIAMONDS),
		NewCard(NINE, DIAMONDS),
	}

	board := Deck{
		NewCard(TEN, DIAMONDS),
		NewCard(TEN, SPADES),
		NewCard(QUEEN, CLUBS),
		NewCard(QUEEN, DIAMONDS),
		NewCard(DUCE, CLUBS),
	}
	sd := MakeShowDown(board, PlayerOneHole, PlayerTwoHole, PlayerThreeHole)
	fmt.Printf("%+v\n", sd)

	for i, v := range DistrubuteChips(100, 10, 0, sd) {
		fmt.Printf("player %d got %d.\n", i, v)
	}
}
