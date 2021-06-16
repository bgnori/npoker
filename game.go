package npoker

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type Line interface {
	/* set up */
	StartOfHand()
	SeatPlayer()
	SetBtn()
	PostSB()
	PostBB()
	DealAll()
	/* actions */
	Fold()
	Bet()
	Raise()
	Check()
	Call()
	/* phase */
	Preflop()
	Flop()
	Turn()
	River()
	Showdown()
	EndOfHand()
}


type LineEventID int

const (
	NULLEVENT LineEventID = iota
	STARTOFHAND
	SEATPLAYER
	SETBTN
	POSTSB
	POSTBB
	DEALX
	DEALALL
	FOLD
	BET
	RAISE
	CHECK
	CALL
	PREFLOP
	FLOP
	TURN
	RIVER
	SHOWDOWN
	ENDOFHAND
	LINEEVENTTYPES
)

type Mock struct {
	// implements Line
	foo LineEventID
}

func NewMock() *Mock {
	a := &Mock{}
	return a
}

func (m *Mock) StartOfHand(){
	m.foo = STARTOFHAND
}

func (m *Mock) SeatPlayer() {
	m.foo = SEATPLAYER
}

func (m *Mock) SetBtn() {
	m.foo = SETBTN
}

func (m *Mock) PostSB() {
	m.foo = POSTSB
}

func (m *Mock) PostBB() {
	m.foo = POSTBB
}

func (m *Mock) DealX(){
	m.foo = DEALX
}

func (m *Mock) DealAll(){
	m.foo = DEALX
}

//actions
func (m *Mock) Fold(){
	m.foo = FOLD
}

func (m *Mock) Bet(){
	m.foo = BET
}

func (m *Mock) Raise(){
	m.foo = RAISE
}

func (m *Mock) Check(){
	m.foo = CHECK
}

func (m *Mock) Call(){
	m.foo = CALL
}

//phase
func (m *Mock) Preflop(){
	m.foo = PREFLOP
}

func (m *Mock) Flop(){
	m.foo = FLOP
}

func (m *Mock) Turn(){
	m.foo = TURN
}

func (m *Mock) River(){
	m.foo = RIVER
}

func (m *Mock) Showdown(){
	m.foo = SHOWDOWN
}

func (m *Mock) EndOfHand(){
	m.foo = ENDOFHAND
}


type PSReader struct {
	line Line
	re *regexp.Regexp
	regExpStrings []string
	names map[string][]string
}

func (reader *PSReader)feed(input io.Reader){
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		reader.feedLine(scanner.Text())
	}
	if err:= scanner.Err(); err!= nil{
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func (reader *PSReader)add(pattern string)*PSReader {
	r := regexp.MustCompile(pattern)
	names := r.SubexpNames()
	reader.names[names[1]] = names[2:]
	reader.regExpStrings = append(reader.regExpStrings, pattern)
	return reader
}

func (reader *PSReader)endOfAdd(){
	reader.re = regexp.MustCompile(strings.Join(reader.regExpStrings, `|`))
}


func NewPSReader() *PSReader {
	g := &PSReader{}
	g.names = make(map[string][]string)
	g.regExpStrings = []string{}
	g.add(`(?P<StartOfHand>^PokerStars Hand #(?P<handNumber>\d+))`)
	g.add(`(?P<SeatPlayer>^Seat (?P<SeatNumber>\d): (?P<PlayerName>\w+))`)
	g.add(`(?P<PostSB>^\w+: posts small blind \d+)`)
	g.add(`(?P<PostBB>^\w+: posts big blind \d+)`)
	g.add(`(?P<DealX>^Dealt to (?P<PlayerName>\w+) \[\w\w \w\w\])`)
	g.add(`(?P<Fold>^\w+: folds)`)
	g.add(`(?P<Bet>^\w+: bets \d+)`)
	g.add(`(?P<Raise>^\w+: raises \d+)`)
	g.add(`(?P<Check>^\w+: checks)`)
	g.add(`(?P<Call>^\w+: calls \d+)`)
	g.add(`(?P<Preflop>^\*\*\* HOLE CARDS \*\*\*)`)
	g.add(`(?P<Flop>^\*\*\* FLOP \*\*\*)`)
	g.add(`(?P<Turn>^\*\*\* TURN \*\*\*)`)
	g.add(`(?P<River>^\*\*\* RIVER \*\*\*)`)
	g.add(`(?P<Showdown>^\*\*\* SHOW DOWN \*\*\*)`)
	g.add(`(?P<EndOfHand>^\*\*\* SUMMARY \*\*\*)`)
	g.endOfAdd()
	return g
}


func (reader *PSReader)Find(line string) map[string][]int{
	d := make(map[string][]int)
	matches := reader.re.FindStringSubmatchIndex(line)
	if len(matches) > 0{
		for i, name := range reader.re.SubexpNames() {
			d[name] = matches[2*i:2*i+2]
		}
	}
	return d
}


func (reader *PSReader)feedLine(line string){
	d := reader.Find(line)
	if len(d) < 1 {
		return
	}
	for key, value := range d {
		if key != "" && value[0] > -1  && value[1] > -1 {
			if  v, ok := reader.names[key]; ok {
				fmt.Println(key, "==>" ,value, v)
				method := reflect.ValueOf(reader.line).MethodByName(key)
				if method.IsValid() {
					method.Call([]reflect.Value{})
				}

			}
		}
	}
}


func (reader *PSReader)debug(){
}
