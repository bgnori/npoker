package npoker

import (
	"fmt"
	"testing"
)

func suitjsonhelper(t *testing.T, expected string, to_json Suit) {
	j, err := to_json.MarshalJSON()
	if err != nil {
		t.Errorf("'%s' is expected for %d got error, %v", expected, to_json, err)
	}
	if expected != string(j) {
		t.Errorf("'%s' is expected for %d", expected, to_json)
	}
}

func TestSuitsMarshallJson(t *testing.T) {
	suitjsonhelper(t, "\u2660", SPADES)
	suitjsonhelper(t, "\u2665", HEARTS)
	suitjsonhelper(t, "\u2666", DIAMONDS)
	suitjsonhelper(t, "\u2663", CLUBS)
}

func suithelper(t *testing.T, expected string, to_string Suit) {
	if expected != fmt.Sprintf("%v", to_string) {
		t.Errorf("'%s' is expected for %d", expected, to_string)
	}
}

func TestSuitsAscii(t *testing.T) {
	suithelper(t, "s", SPADES)
	suithelper(t, "h", HEARTS)
	suithelper(t, "d", DIAMONDS)
	suithelper(t, "c", CLUBS)
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

func cardjsonhelper(t *testing.T, expected string, to_json Card) {
	j, err := to_json.MarshalJSON()
	if err != nil {
		t.Errorf("'%s' is expected for %d got error, %v", expected, to_json, err)
	}
	if expected != string(j) {
		t.Errorf("'%s' is expected for %v but got %s", expected, to_json, string(j))
	}
}

func TestJsonCards(t *testing.T) {
	cardjsonhelper(t, "A\u2665", NewCard(ACE, HEARTS))
	cardjsonhelper(t, "4\u2666", NewCard(FOUR, DIAMONDS))
}

func cardhelper(t *testing.T, expected string, to_string Card) {
	got := fmt.Sprintf("%s", to_string)
	if expected != got {
		t.Errorf("'%s' is expected for %+v but got %s", expected, to_string, got)
	}
}

func TestCards(t *testing.T) {
	cardhelper(t, "Ah", NewCard(ACE, HEARTS))
	cardhelper(t, "4d", NewCard(FOUR, DIAMONDS))
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
