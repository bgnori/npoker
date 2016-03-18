package npoker

import (
	"fmt"
)

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
