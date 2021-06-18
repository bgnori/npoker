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
	Bare bool
}


func NewRecursiveHandler(name string, partial string, bare bool, children ...*RecursiveHandler) *RecursiveHandler {
	rh := &RecursiveHandler{}
	rh.Name = name
	rh.Partial = partial
	rh.Bare = bare
	rh.Children = children
	return rh
}

func (rh *RecursiveHandler) makeRegExp() string {
	t := make([]string, 0)
	for _, c := range rh.Children {
		t = append(t, c.makeRegExp())
	}

	if rh.Bare {
		return rh.Partial
	}
	if rh.Name != "" {
		return fmt.Sprintf(`(?P<%s>%s%s)`, rh.Name, rh.Partial, strings.Join(t, ""))
	}else {
		return fmt.Sprintf(`(%s%s)`, rh.Partial, strings.Join(t, ""))
	}
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
	playername := NewRecursiveHandler("PlayerName", `\w+`, false)
	amount := NewRecursiveHandler("Amount", ``, false,
			NewRecursiveHandler("TournamentChip", `\d+`, false),
			NewRecursiveHandler("", `|`, true),
			NewRecursiveHandler("RingChip", ``, false,
				NewRecursiveHandler("Currency", `\$`, false),
				NewRecursiveHandler("Float", `\d+\.\d+`, false)))
	allin := NewRecursiveHandler("Allin", `( and is all-in)?`, true)


	g := &PSReader{}
	g.handlers = make(map[string]*RecursiveHandler)
	g.add_handler(NewRecursiveHandler("StartOfHand", "^PokerStars Hand #", false,
		NewRecursiveHandler("HandNumber", `\d+`, false)))
	g.add_handler(NewRecursiveHandler("SetBtn", "^Table ", false))
	g.add_handler(NewRecursiveHandler("SeatPlayer", "^Seat ", false,
		NewRecursiveHandler("SeatNumber", `\d+`, false),
		NewRecursiveHandler("", `: `, true),
		playername))
	g.add_handler(NewRecursiveHandler("PostSB", "^", false,
			playername,
			NewRecursiveHandler("", `: posts small blind `, false),
			amount))
	g.add_handler(NewRecursiveHandler("PostBB", "^", false,
			playername,
			NewRecursiveHandler("", `: posts big blind `, false),
			amount))
	g.add_handler(NewRecursiveHandler("DealX", "^Dealt to ", false,
			playername,
			NewRecursiveHandler("", ` `, true),
			NewRecursiveHandler("HoleCard", `\[\w\w \w\w\]`, false)))
	g.add_handler(NewRecursiveHandler("Fold", "^", false,
			playername,
			NewRecursiveHandler("", `: folds`, false)))
	g.add_handler(NewRecursiveHandler("Bet", "^", false,
			playername,
			NewRecursiveHandler("", `: bets `, false),
			amount,
			allin))
	g.add_handler(NewRecursiveHandler("Raise", "^", false,
			playername,
			NewRecursiveHandler("", `: raises `, false),
			amount,
			allin))
	g.add_handler(NewRecursiveHandler("Check", "^", false,
			playername,
			NewRecursiveHandler("", `: checks`, false)))
	g.add_handler(NewRecursiveHandler("Call", "^", false,
			playername,
			NewRecursiveHandler("", `: calls `, false),
			amount,
			allin))
	g.add_handler(NewRecursiveHandler("Preflop", `^\*\*\* HOLE CARDS \*\*\*`, false))
	g.add_handler(NewRecursiveHandler("Flop", `^\*\*\* FLOP \*\*\*`, false))
	g.add_handler(NewRecursiveHandler("Turn", `^\*\*\* TURN \*\*\*`, false))
	g.add_handler(NewRecursiveHandler("River", `^\*\*\* RIVER \*\*\*`, false))
	g.add_handler(NewRecursiveHandler("Showdown", `^\*\*\* SHOW DOWN \*\*\*`, false))
	g.add_handler(NewRecursiveHandler("EndOfHand", `^\*\*\* SUMMARY \*\*\*`, false))

	g.endOfAdd()
	//fmt.Println(g.re)
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
	//fmt.Println(found)
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
