
package npoker

import (
	"testing"
	"bytes"
	//"reflect"
	//"fmt"
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

func TestPSReaderfeedLine_SETBTN(t *testing.T) {
	sample := `Seat 1: adevlupec (53368 in chips)`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != SETBTN{
		t.Errorf("%v is expected but got %v",  SETBTN, mock.foo)
	}
}


func TestPSReaderfeedLine_POSTSB(t *testing.T) {
	sample := `FluffyStutt: posts small blind 50`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != POSTSB{
		t.Errorf("%v is expected but got %v",  POSTSB, mock.foo)
	}
}


func TestPSReaderfeedLine_POSTBB(t *testing.T) {
	sample := `adevlupec: posts big blind 100`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != POSTBB{
		t.Errorf("%v is expected but got %v",  POSTBB, mock.foo)
	}
}


func TestPSReaderfeedLine_DEALX(t *testing.T) {
	sample := `Dealt to FluffyStutt [2h Ks]`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != DEALX{
		t.Errorf("%v is expected but got %v",  DEALX, mock.foo)
	}
}

func TestPSReaderfeedLine_FOLD(t *testing.T) {
	sample := `FluffyStutt: folds`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != FOLD{
		t.Errorf("%v is expected but got %v",  FOLD, mock.foo)
	}
}


func TestPSReaderfeedLine_BET(t *testing.T) {
	//UGH! no testing data
	sample := `FluffyStutt: bets 100`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != BET{
		t.Errorf("%v is expected but got %v",  BET, mock.foo)
	}
}


func TestPSReaderfeedLine_RAISE(t *testing.T) {
	//UGH! no testing data
	sample := `FluffyStutt: raises 200`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != RAISE{
		t.Errorf("%v is expected but got %v",  RAISE, mock.foo)
	}
}

func TestPSReaderfeedLine_CHECK(t *testing.T) {
	sample := `adevlupec: checks`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != CHECK{
		t.Errorf("%v is expected but got %v",  CHECK, mock.foo)
	}
}

func TestPSReaderfeedLine_CALL(t *testing.T) {
	sample := `Drug08: calls 100`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != CALL{
		t.Errorf("%v is expected but got %v",  CALL, mock.foo)
	}
}

func TestPSReaderfeedLine_PREFLOP(t *testing.T) {
	sample := `*** HOLE CARDS ***`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != PREFLOP{
		t.Errorf("%v is expected but got %v",  PREFLOP, mock.foo)
	}
}

func TestPSReaderfeedLine_FLOP(t *testing.T) {
	sample := `*** FLOP *** [8h 7s 8d]`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != FLOP{
		t.Errorf("%v is expected but got %v",  FLOP, mock.foo)
	}
}


func TestPSReaderfeedLine_TURN(t *testing.T) {
	sample := `*** TURN *** [8h 7s 8d] [Th]`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != TURN{
		t.Errorf("%v is expected but got %v",  TURN, mock.foo)
	}
}

func TestPSReaderfeedLine_RIVER(t *testing.T) {
	sample := `*** RIVER *** [8h 7s 8d Th] [2c]`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != RIVER{
		t.Errorf("%v is expected but got %v",  RIVER, mock.foo)
	}
}

func TestPSReaderfeedLine_SHOWDOWN(t *testing.T) {
	sample := `*** SHOW DOWN ***`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != SHOWDOWN{
		t.Errorf("%v is expected but got %v",  SHOWDOWN, mock.foo)
	}
}

func TestPSReaderfeedLine_ENDOFHAND(t *testing.T) {
	sample := `*** SUMMARY ***`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != ENDOFHAND{
		t.Errorf("%v is expected but got %v",  ENDOFHAND, mock.foo)
	}
}
