package npoker

import (
//	"fmt"
//"strings"
)

func RollOut(trial int, board Deck, players ...Deck) []int {
	stat := make([]int, len(players))

	xs := BuildFullDeck()
	xs.Subtract(board)
	for _, p := range players {
		xs.Subtract(p)
	}

	for i := 0; i < trial; i++ {
		ys := xs.Clone()
		ys.Shuffle()
		ys.ShrinkTo(5 - len(board))

		river := Join(board, ys)

		sd := MakeShowDown(river, players...)
		for _, w := range sd.Winners {
			stat[w] += 1
		}
	}
	return stat
}
