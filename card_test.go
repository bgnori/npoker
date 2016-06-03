package npoker

import (
	"fmt"
	"testing"
)

func suithelper(t *testing.T, expected string, to_string Suit) {
	if expected != fmt.Sprintf("%v", to_string) {
		t.Errorf("'%s' is expected for %d", expected, to_string)
	}
}

func TestSuits(t *testing.T) {
	suithelper(t, "\u2660", SPADES)
	suithelper(t, "\u2665", HEARTS)
	suithelper(t, "\u2666", DIAMONDS)
	suithelper(t, "\u2663", CLUBS)
}

func rankhelper(t *testing.T, expected string, to_string Rank) {
	if expected != fmt.Sprintf("%v", to_string) {
		t.Errorf("'%s' is expected for %d", expected, to_string)
	}
}

func TestRanks(t *testing.T) {
	rankhelper(t, "A", ACE)
	rankhelper(t, "2", DUCE)
	rankhelper(t, "3", TREY)
	rankhelper(t, "4", FOUR)
	rankhelper(t, "5", FIVE)
	rankhelper(t, "6", SIX)
	rankhelper(t, "7", SEVEN)
	rankhelper(t, "8", EIGHT)
	rankhelper(t, "9", NINE)
	rankhelper(t, "T", TEN)
	rankhelper(t, "J", JACK)
	rankhelper(t, "Q", QUEEN)
	rankhelper(t, "K", KING)
	rankhelper(t, "A", HIACE)
}

func cardhelper(t *testing.T, expected string, to_string Card) {
	if expected != fmt.Sprintf("%s", to_string) {
		t.Errorf("'%s' is expected for %v but got %s", expected, to_string, fmt.Sprintf("%s", to_string))
	}
}

func TestCards(t *testing.T) {
	cardhelper(t, "A\u2665", NewCard(ACE, HEARTS))
	cardhelper(t, "4\u2666", NewCard(FOUR, DIAMONDS))
}

func SmokeParseWith(input string) {
	got := parse(input)
	fmt.Println(input, "==>", got)
}

func parsehelper(t *testing.T, expected Card, input string) {
	if expected != parse(input) {
		t.Errorf("'%v' is expected for '%s'", expected, input)
	}
}

func TestParse(t *testing.T) {
	parsehelper(t, NewCard(ACE, HEARTS), "Ah")
	parsehelper(t, NewCard(FIVE, CLUBS), "5c")
}

func cardordhelper(t *testing.T, expected int, to_string Card) {
	if expected != to_string.Ord() {
		t.Errorf("'%d' is expected for %v, but got %d.", expected, to_string, to_string.Ord())
	}
}

func TestOrd(t *testing.T) {
	cardordhelper(t, 0, NewCard(ACE, CLUBS))
	cardordhelper(t, 12, NewCard(KING, CLUBS))
	cardordhelper(t, 1, NewCard(DUCE, CLUBS))
	cardordhelper(t, 39, NewCard(ACE, SPADES))
	cardordhelper(t, 51, NewCard(KING, SPADES))
	cardordhelper(t, 40, NewCard(DUCE, SPADES))
}
