package npoker

import (
	"bufio"
	"fmt"
	"io"
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

func (m *Mock) StartOfHand() {
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

func (m *Mock) DealX() {
	m.foo = DEALX
}

func (m *Mock) DealAll() {
	m.foo = DEALX
}

//actions
func (m *Mock) Fold() {
	m.foo = FOLD
}

func (m *Mock) Bet() {
	m.foo = BET
}

func (m *Mock) Raise() {
	m.foo = RAISE
}

func (m *Mock) Check() {
	m.foo = CHECK
}

func (m *Mock) Call() {
	m.foo = CALL
}

//phase
func (m *Mock) Preflop() {
	m.foo = PREFLOP
}

func (m *Mock) Flop() {
	m.foo = FLOP
}

func (m *Mock) Turn() {
	m.foo = TURN
}

func (m *Mock) River() {
	m.foo = RIVER
}

func (m *Mock) Showdown() {
	m.foo = SHOWDOWN
}

func (m *Mock) EndOfHand() {
	m.foo = ENDOFHAND
}

type RecursiveHandler struct {
	Name     string
	Partial  string
	Children []*RecursiveHandler
}

func NewRecursiveHandler(name string, partial string, children ...*RecursiveHandler) *RecursiveHandler {
	rh := &RecursiveHandler{}
	rh.Name = name
	rh.Partial = partial
	rh.Children = children
	return rh
}

func (rh *RecursiveHandler) makeRegExp() string {
	t := make([]string, 0)
	for _, c := range rh.Children {
		t = append(t, c.makeRegExp())
	}

	return fmt.Sprintf(`(?P<%s>%s%s)`, rh.Name, rh.Partial, strings.Join(t, ""))
}

type PSReader struct {
	line     Line
	re       *regexp.Regexp
	handlers map[string]*RecursiveHandler
}

func (reader *PSReader) feed(input io.Reader) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		reader.feedLine(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func (reader *PSReader) add_handler(rh *RecursiveHandler) {
	reader.handlers[rh.Name] = rh
}

func (reader *PSReader) endOfAdd() {
	t := make([]string, 0)
	for _, v := range reader.handlers {
		t = append(t, v.makeRegExp())
	}
	reader.re = regexp.MustCompile(strings.Join(t, `|`))
}

func NewPSReader() *PSReader {
	//amount := `(?P<Amount>(?P<TournamentChip>\d+)|(?P<RingChip>(?P<currency>\$)|(?P<float>\d+\.\d+)))`
	//allin := `(?P<AndAllin> and is all-in)?`

	g := &PSReader{}
	g.handlers = make(map[string]*RecursiveHandler)
	g.add_handler(NewRecursiveHandler("StartOfHand", "^PokerStars Hand #",
		NewRecursiveHandler("HandNumber", `\d+`)))
	/*
		g.add(`(?P<StartOfHand>^PokerStars Hand #(?P<handNumber>\d+))`)
		g.add(`(?P<SeatPlayer>^Seat (?P<SeatNumber>\d): (?P<PlayerName>\w+))`)
		g.add(`(?P<PostSB>^\w+: posts small blind \d+)`)
		g.add(`(?P<PostBB>^\w+: posts big blind \d+)`)
		g.add(`(?P<DealX>^Dealt to (?P<PlayerName>\w+) \[\w\w \w\w\])`)
		g.add(`(?P<Fold>^\w+: folds)`)
		g.add(`(?P<Bet>^\w+: bets ` + amount + allin + `)`)
		g.add(`(?P<Raise>^\w+: raises \d+)`)
		g.add(`(?P<Check>^\w+: checks)`)
		g.add(`(?P<Call>^\w+: calls \d+)`)
		g.add(`(?P<Preflop>^\*\*\* HOLE CARDS \*\*\*)`)
		g.add(`(?P<Flop>^\*\*\* FLOP \*\*\*)`)
		g.add(`(?P<Turn>^\*\*\* TURN \*\*\*)`)
		g.add(`(?P<River>^\*\*\* RIVER \*\*\*)`)
		g.add(`(?P<Showdown>^\*\*\* SHOW DOWN \*\*\*)`)
		g.add(`(?P<EndOfHand>^\*\*\* SUMMARY \*\*\*)`)
	*/
	g.endOfAdd()
	fmt.Println(g.re)
	return g
}

func (reader *PSReader) Find(line string) map[string][]int {
	d := make(map[string][]int)
	matches := reader.re.FindStringSubmatchIndex(line)
	if len(matches) > 0 {
		for i, name := range reader.re.SubexpNames() {
			d[name] = matches[2*i : 2*i+2]
		}
	}
	return d
}

func (reader *PSReader) feedLine(line string) {
	found := reader.Find(line)
	fmt.Println(found)
	if len(found) < 1 {
		return
	}
	for key, value := range found {
		if key != "" && value[0] > -1 && value[1] > -1 {
			if v, ok := reader.handlers[key]; ok {
				fmt.Println(key, "==>", value, v)
				method := reflect.ValueOf(reader.line).MethodByName(key)
				if method.IsValid() {
					method.Call([]reflect.Value{})
				}

			}
		}
	}
}

func (reader *PSReader) debug() {
}
