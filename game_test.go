package npoker

import (
	"bytes"
	"testing"
	//"reflect"
	//"fmt"
)

func TestPSReader_FluffyStutt(t *testing.T) {
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



PokerStars Hand #174088919486:  Hold'em No Limit (50/100) - 2017/08/08 23:17:57 MSK [2017/08/08 16:17:57 ET]
Table 'Euphemia II' 6-max (Play Money) Seat #4 is the button
Seat 1: adevlupec (53600 in chips) 
Seat 2: Dette32 (10745 in chips) 
Seat 3: Drug08 (9586 in chips) 
Seat 4: FluffyStutt (11276 in chips) 
yanksea will be allowed to play after the button
adevlupec: posts small blind 50
Dette32: posts big blind 100
*** HOLE CARDS ***
Dealt to FluffyStutt [8s Qc]
Drug08: calls 100
FluffyStutt: calls 100
adevlupec: folds 
Dette32: checks 
*** FLOP *** [8c Qh 4c]
Dette32: checks 
Drug08: checks 
FluffyStutt: bets 332
Dette32: folds 
Drug08: calls 332
*** TURN *** [8c Qh 4c] [Ac]
Drug08: checks 
FluffyStutt: bets 963
Drug08: calls 963
*** RIVER *** [8c Qh 4c Ac] [Td]
Drug08: checks 
FluffyStutt: bets 9881 and is all-in
Drug08: folds 
Uncalled bet (9881) returned to FluffyStutt
FluffyStutt collected 2793 from pot
FluffyStutt: doesn't show hand 
*** SUMMARY ***
Total pot 2940 | Rake 147 
Board [8c Qh 4c Ac Td]
Seat 1: adevlupec (small blind) folded before Flop
Seat 2: Dette32 (big blind) folded on the Flop
Seat 3: Drug08 folded on the River
Seat 4: FluffyStutt (button) collected (2793)



`
	input := bytes.NewBufferString(sample)
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feed(input)
	reader.debug()
}

func TestPSReader_wikipedia01(t *testing.T) {
	/*
		from https://en.wikipedia.org/wiki/Hand_history
	*/
	sample := `PokerStars Hand #156605842579: Tournament #1618015625, $2.00+$2.00+$0.40 USD Hold'em No Limit - Level XVIII (300/600) - 2016/07/29 23:25:05 MSK [2016/07/29 16:25:05 ET]
Table '1618015625 40' 9-max Seat #7 is the button
Seat 1: nossoff (22844 in chips) 
Seat 2: Gr0z3LB (27895 in chips) 
Seat 3: yarikkkkk (7005 in chips) 
Seat 4: sigis1313 (18222 in chips) 
Seat 5: S21berlin (25318 in chips) 
Seat 6: Hero (8240 in chips) 
Seat 7: tommesk (8185 in chips) 
Seat 8: gergelyt3 (17912 in chips) 
Seat 9: endreooo (50289 in chips) 
nossoff: posts the ante 90
Gr0z3LB: posts the ante 90
yarikkkkk: posts the ante 90
sigis1313: posts the ante 90
S21berlin: posts the ante 90
Hero: posts the ante 90
tommesk: posts the ante 90
gergelyt3: posts the ante 90
endreooo: posts the ante 90
gergelyt3: posts small blind 300
endreooo: posts big blind 600
*** HOLE CARDS ***
Dealt to Hero [Jd Ac]
nossoff: raises 600 to 1200
Gr0z3LB: folds 
yarikkkkk: folds 
sigis1313: folds 
S21berlin: folds 
Hero: raises 6950 to 8150 and is all-in
tommesk: folds 
gergelyt3: folds 
endreooo: folds 
nossoff: calls 6950
*** FLOP *** [7h 2h 5s]
*** TURN *** [7h 2h 5s] [5d]
*** RIVER *** [7h 2h 5s 5d] [4d]
*** SHOW DOWN ***
nossoff: shows [Qd Ah] (a pair of Fives)
Hero: shows [Jd Ac] (a pair of Fives - lower kicker)
nossoff collected 18010 from pot
nossoff wins $1 for eliminating Hero and their own bounty increases by $1 to $6
Hero finished the tournament in 5028th place
*** SUMMARY ***
Total pot 18010 | Rake 0 
Board [7h 2h 5s 5d 4d]
Seat 1: nossoff showed [Qd Ah] and won (18010) with a pair of Fives
Seat 2: Gr0z3LB folded before Flop (didn't bet)
Seat 3: yarikkkkk folded before Flop (didn't bet)
Seat 4: sigis1313 folded before Flop (didn't bet)
Seat 5: S21berlin folded before Flop (didn't bet)
Seat 6: Hero showed [Jd Ac] and lost with a pair of Fives
Seat 7: tommesk (button) folded before Flop (didn't bet)
Seat 8: gergelyt3 (small blind) folded before Flop
Seat 9: endreooo (big blind) folded before Flop`
	input := bytes.NewBufferString(sample)
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feed(input)
	reader.debug()
}

func TestPSReader_wikipedia02(t *testing.T) {
	/*
		from https://en.wikipedia.org/wiki/Hand_history
	*/
	sample := `PokerStars Hand #156624936552:  Hold'em No Limit ($0.05/$0.10 USD) - 2016/07/30 4:59:02 ET
Table 'Shenzhen' 9-max Seat #3 is the button
Seat 3: Hero ($10.05 in chips) 
Seat 8: Alpha242 ($10.89 in chips) 
Hero: posts small blind $0.05
Alpha242: posts big blind $0.10
*** HOLE CARDS ***
Dealt to Hero [6h 7h]
Hero: raises $0.10 to $0.20
Alpha242: calls $0.10
*** FLOP *** [5s Jd Jc]
Alpha242: checks 
Hero: bets $0.20
Alpha242: raises $0.51 to $0.71
Hero: calls $0.51
*** TURN *** [5s Jd Jc] [2s]
Alpha242: bets $1
Hero: raises $1.60 to $2.60
Alpha242: folds 
Uncalled bet ($1.60) returned to Hero
Hero collected $3.65 from pot
Hero: doesn't show hand 
*** SUMMARY ***
Total pot $3.82 | Rake $0.17 
Board [5s Jd Jc 2s]
Seat 3: Hero (button) (small blind) collected ($3.65)
Seat 8: Alpha242 (big blind) folded on the Turn
`
	input := bytes.NewBufferString(sample)
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feed(input)
	reader.debug()
}

func TestPSReader_wikipedia03(t *testing.T) {
	/*
		from https://en.wikipedia.org/wiki/Hand_history
	*/
	sample := `PokerStars Hand #156632473469:  Hold'em No Limit ($0.05/$0.10 USD) - 2016/07/30 10:40:31 AT [2016/07/30 9:40:31 ET]
Table 'Eltigen' 9-max Seat #8 is the button
Seat 1: raika14 ($10 in chips)
Seat 2: mazi161 ($11 in chips)
Seat 3: Hero ($10.10 in chips)
Seat 4: 5amevang ($10.42 in chips)
Seat 6: dinqua ($10.15 in chips)
Seat 7: gragasinha ($5.55 in chips)
Seat 8: MilonTl ($15.46 in chips)
Seat 9: enchantressA ($10.30 in chips)
enchantressA: posts small blind $0.05
raika14: posts big blind $0.10
*** HOLE CARDS ***
Dealt to Hero [Jh Jc]
mazi161: folds
Hero: raises $0.20 to $0.30
5amevang: folds
dinqua: raises $0.60 to $0.90
gragasinha: folds
MilonTl: folds
enchantressA: folds
raika14: folds
Hero: calls $0.60
*** FLOP *** [7h 8d 5c]
Hero: checks
dinqua: bets $1.30
Hero: calls $1.30
*** TURN *** [7h 8d 5c] [8c]
Hero: checks
dinqua: bets $2.90
Hero: folds
Uncalled bet ($2.90) returned to dinqua
dinqua collected $4.35 from pot
*** SUMMARY ***
Total pot $4.55 | Rake $0.20
Board [7h 8d 5c 8c]
Seat 1: raika14 (big blind) folded before Flop
Seat 2: mazi161 folded before Flop (didn't bet)
Seat 3: Hero folded on the Turn
Seat 4: 5amevang folded before Flop (didn't bet)
Seat 6: dinqua collected ($4.35)
Seat 7: gragasinha folded before Flop (didn't bet)
Seat 8: MilonTl (button) folded before Flop (didn't bet)
Seat 9: enchantressA (small blind) folded before Flop`
	input := bytes.NewBufferString(sample)
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feed(input)
	reader.debug()
}

func TestPSReader_wikipedia04(t *testing.T) {
	/*
		from https://en.wikipedia.org/wiki/Hand_history
	*/
	sample := `PokerStars Game #27738502010: Tournament #160417133, $0.25+$0.00 Hold'em No Limit - Level XV (250/500) - 2009/05/02 13:32:38 ET
Table '160417133 3' 9-max Seat #8 is the button
Seat 1: LLC 4Eva (9182 in chips) 
Seat 2: 618shooter (25711 in chips) is sitting out
Seat 3: suposd2bRich (21475 in chips) 
Seat 4: ElT007 (60940 in chips) 
Seat 5: Orlando I (18044 in chips) 
Seat 6: ih82bcool2 (8338 in chips) 
Seat 7: kovilen007 (8353 in chips) 
Seat 8: GerKingTiger (4404 in chips) 
Seat 9: Phontaz (23553 in chips) 
LLC 4Eva: posts the ante 60
618shooter: posts the ante 60
suposd2bRich: posts the ante 60
ElT007: posts the ante 60
Orlando I: posts the ante 60
ih82bcool2: posts the ante 60
kovilen007: posts the ante 60
GerKingTiger: posts the ante 60
Phontaz: posts the ante 60
Phontaz: posts small blind 250
LLC 4Eva: posts big blind 500
*** HOLE CARDS ***
Dealt to ElT007 [Qd Qc]
618shooter: folds 
suposd2bRich: folds 
ElT007: raises 2000 to 2500
Orlando I: raises 15484 to 17984 and is all-in
ih82bcool2: folds 
kovilen007: calls 8293 and is all-in
GerKingTiger: folds 
Phontaz: calls 17734
LLC 4Eva: folds 
ElT007: raises 15484 to 33468
Phontaz: calls 5509 and is all-in
Uncalled bet (9975) returned to ElT007
*** FLOP *** [2d 2c 3c]
*** TURN *** [2d 2c 3c] [8h]
*** RIVER *** [2d 2c 3c 8h] [4d]
*** SHOW DOWN ***
Phontaz: shows [9s 9h] (two pair, Nines and Deuces)
ElT007: shows [Qd Qc] (two pair, Queens and Deuces)
618shooter has returned
ElT007 collected 11018 from side pot-2 
Orlando I: shows [5d 5h] (two pair, Fives and Deuces)
ElT007 collected 29073 from side pot-1 
kovilen007: shows [Kh As] (a pair of Deuces)
ElT007 collected 34212 from main pot
*** SUMMARY ***
Total pot 74303 Main pot 34212. Side pot-1 29073. Side pot-2 11018. | Rake 0 
Board [2d 2c 3c 8h 4d]
Seat 1: LLC 4Eva (big blind) folded before Flop
Seat 2: 618shooter folded before Flop (didn't bet)
Seat 3: suposd2bRich folded before Flop (didn't bet)
Seat 4: ElT007 showed [Qd Qc] and won (74303) with two pair, Queens and Deuces
Seat 5: Orlando I showed [5d 5h] and lost with two pair, Fives and Deuces
Seat 6: ih82bcool2 folded before Flop (didn't bet)
Seat 7: kovilen007 showed [Kh As] and lost with a pair of Deuces
Seat 8: GerKingTiger (button) folded before Flop (didn't bet)
Seat 9: Phontaz (small blind) showed [9s 9h] and lost with two pair, Nines and Deuces`
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
		t.Errorf("%v is expected but got %v", STARTOFHAND, mock.foo)
	}
}

func TestPSReaderfeedLine_SEATPLAYER(t *testing.T) {
	sample := `Seat 1: adevlupec (53368 in chips)`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != SEATPLAYER {
		t.Errorf("%v is expected but got %v", SEATPLAYER, mock.foo)
	}
}

func TestPSReaderfeedLine_SETBTN(t *testing.T) {
	sample := `Table 'Eltigen' 9-max Seat #8 is the button`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != SETBTN {
		t.Errorf("%v is expected but got %v", SETBTN, mock.foo)
	}
}

func TestPSReaderfeedLine_POSTSB(t *testing.T) {
	sample := `FluffyStutt: posts small blind 50`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != POSTSB {
		t.Errorf("%v is expected but got %v", POSTSB, mock.foo)
	}
}

func TestPSReaderfeedLine_POSTBB(t *testing.T) {
	sample := `adevlupec: posts big blind 100`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != POSTBB {
		t.Errorf("%v is expected but got %v", POSTBB, mock.foo)
	}
}

func TestPSReaderfeedLine_DEALX(t *testing.T) {
	sample := `Dealt to FluffyStutt [2h Ks]`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != DEALX {
		t.Errorf("%v is expected but got %v", DEALX, mock.foo)
	}
}

func TestPSReaderfeedLine_FOLD(t *testing.T) {
	sample := `FluffyStutt: folds`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != FOLD {
		t.Errorf("%v is expected but got %v", FOLD, mock.foo)
	}
}

func TestPSReaderfeedLine_intBET(t *testing.T) {
	//UGH! no testing data
	sample := `FluffyStutt: bets 100`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != BET {
		t.Errorf("%v is expected but got %v", BET, mock.foo)
	}
}

func TestPSReaderfeedLine_floatBET(t *testing.T) {
	//UGH! no testing data
	sample := `Hero: bets $0.20`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != BET {
		t.Errorf("%v is expected but got %v", BET, mock.foo)
	}
}

func TestPSReaderfeedLine_RAISE(t *testing.T) {
	//UGH! no testing data
	sample := `FluffyStutt: raises 200`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != RAISE {
		t.Errorf("%v is expected but got %v", RAISE, mock.foo)
	}
}

func TestPSReaderfeedLine_CHECK(t *testing.T) {
	sample := `adevlupec: checks`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != CHECK {
		t.Errorf("%v is expected but got %v", CHECK, mock.foo)
	}
}

func TestPSReaderfeedLine_CALL(t *testing.T) {
	sample := `Drug08: calls 100`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != CALL {
		t.Errorf("%v is expected but got %v", CALL, mock.foo)
	}
}

func TestPSReaderfeedLine_PREFLOP(t *testing.T) {
	sample := `*** HOLE CARDS ***`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != PREFLOP {
		t.Errorf("%v is expected but got %v", PREFLOP, mock.foo)
	}
}

func TestPSReaderfeedLine_FLOP(t *testing.T) {
	sample := `*** FLOP *** [8h 7s 8d]`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != FLOP {
		t.Errorf("%v is expected but got %v", FLOP, mock.foo)
	}
}

func TestPSReaderfeedLine_TURN(t *testing.T) {
	sample := `*** TURN *** [8h 7s 8d] [Th]`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != TURN {
		t.Errorf("%v is expected but got %v", TURN, mock.foo)
	}
}

func TestPSReaderfeedLine_RIVER(t *testing.T) {
	sample := `*** RIVER *** [8h 7s 8d Th] [2c]`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != RIVER {
		t.Errorf("%v is expected but got %v", RIVER, mock.foo)
	}
}

func TestPSReaderfeedLine_SHOWDOWN(t *testing.T) {
	sample := `*** SHOW DOWN ***`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != SHOWDOWN {
		t.Errorf("%v is expected but got %v", SHOWDOWN, mock.foo)
	}
}

func TestPSReaderfeedLine_ENDOFHAND(t *testing.T) {
	sample := `*** SUMMARY ***`
	reader := NewPSReader()
	mock := NewMock()
	reader.line = mock
	reader.feedLine(sample)
	if mock.foo != ENDOFHAND {
		t.Errorf("%v is expected but got %v", ENDOFHAND, mock.foo)
	}
}
