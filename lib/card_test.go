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
	if expected != fmt.Sprintf("%v", to_string) {
		t.Errorf("'%s' is expected for %d", expected, to_string)
	}
}

func TestCards(t *testing.T) {
	cardhelper(t, "A\u2665", Card{ACE, HEARTS})
	cardhelper(t, "4\u2666", Card{FOUR, DIAMONDS})
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
	parsehelper(t, Card{ACE, HEARTS}, "Ah")
	parsehelper(t, Card{FIVE, CLUBS}, "5c")
}

func cardordhelper(t *testing.T, expected int, to_string Card) {
	if expected != to_string.Ord() {
		t.Errorf("'%s' is expected for %d, but got %d.", expected, to_string, to_string.Ord())
	}
}

func TestOrd(t *testing.T) {
	cardordhelper(t, 14, Card{ACE, CLUBS})
	cardordhelper(t, 13, Card{KING, CLUBS})
	cardordhelper(t, 2, Card{DUCE, CLUBS})
	cardordhelper(t, 53, Card{ACE, SPADES})
	cardordhelper(t, 52, Card{KING, SPADES})
	cardordhelper(t, 41, Card{DUCE, SPADES})
}
