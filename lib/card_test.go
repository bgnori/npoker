package npoker

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	fmt.Println("Suits:", SPADES, HEARTS, DIAMONDS, CLUBS)
	fmt.Println("Ranks:", ACE, DUCE, TEN, JACK, QUEEN, KING)
	fmt.Println("Card:", Card{ACE, HEARTS})
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
