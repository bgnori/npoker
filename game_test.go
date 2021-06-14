
package npoker

import (
	//"fmt"
	"testing"
	"bytes"
)

func TestPSReaderFull(t *testing.T) {
	/*
		from https://gist.github.com/evgeniyp/e2e6842e7c84881d7611869482c0c930
	*/
	sample := `PokerStars Hand #174088855475:  Hold'em No Limit (50/100) - 2017/08/08 23:16:30 MSK [2017/08/08 16:16:30 ET]
Table 'Euphemia II' 6-max (Play Money) Seat #3 is the button
Seat 1: adevlupec (53368 in chips) 
Seat 2: Dette32 (10845 in chips) 
Seat 3: Drug08 (9686 in chips) 
Seat 4: FluffyStutt (11326 in chips) 
FluffyStutt: posts small blind 50
adevlupec: posts big blind 100
*** HOLE CARDS ***
Dealt to FluffyStutt [2h Ks]
FluffyStutt said, \"nh\"
Dette32: calls 100
Drug08: calls 100
FluffyStutt: folds 
adevlupec: checks 
*** FLOP *** [8h 7s 8d]
adevlupec: checks 
Dette32: checks 
Drug08: checks 
*** TURN *** [8h 7s 8d] [Th]
adevlupec: checks 
Dette32: checks 
Drug08: checks 
*** RIVER *** [8h 7s 8d Th] [2c]
adevlupec: checks 
Dette32: checks 
Drug08: checks 
*** SHOW DOWN ***
adevlupec: shows [Qs Ts] (two pair, Tens and Eights)
Dette32: mucks hand 
Drug08: mucks hand 
adevlupec collected 332 from pot
*** SUMMARY ***
Total pot 350 | Rake 18 
Board [8h 7s 8d Th 2c]
Seat 1: adevlupec (big blind) showed [Qs Ts] and won (332) with two pair, Tens and Eights
Seat 2: Dette32 mucked [5s Kc]
Seat 3: Drug08 (button) mucked [4d 6h]
Seat 4: FluffyStutt (small blind) folded before Flop
`
	input := bytes.NewBufferString(sample)
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feed(input)
	reader.debug()
}


func TestPSReaderfeedLine_STARTOFHAND(t *testing.T) {
	sample := `PokerStars Hand #174088855475:  Hold'em No Limit (50/100) - 2017/08/08 23:16:30 MSK [2017/08/08 16:16:30 ET]
`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != STARTOFHAND {
		t.Errorf("%v is expected but got %v",  STARTOFHAND, mock.foo)
	}
}



func TestPSReaderfeedLine_SEATPLAYER(t *testing.T) {
	sample := `Seat 1: adevlupec (53368 in chips)`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != SEATPLAYER {
		t.Errorf("%v is expected but got %v",  SEATPLAYER, mock.foo)
	}
}


