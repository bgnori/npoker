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
	startofHand()
	seatPlayer()
	setBtn()
	postSB()
	postBB()
	dealAll()
	/* actions */
	fold()
	bet()
	raise()
	check()
	call()
	/* phase */
	preflop()
	flop()
	turn()
	river()
	showdown()
	endofHand()
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
	return &Mock{}
}

func (m *Mock) startofHand(){
	m.foo = STARTOFHAND
}

func (m *Mock) seatPlayer() {
	m.foo = SEATPLAYER
}

func (m *Mock) setBtn() {
	m.foo = SETBTN
}

func (m *Mock) postSB() {
	m.foo = POSTSB
}

func (m *Mock) postBB() {
	m.foo = POSTBB
}

func (m *Mock) dealX(){
	m.foo = DEALX
}
func (m *Mock) dealAll(){
	m.foo = DEALX
}

//actions
func (m *Mock) fold(){
	m.foo = FOLD
}

func (m *Mock) bet(){
	m.foo = BET
}

func (m *Mock) raise(){
	m.foo = RAISE
}

func (m *Mock) check(){
	m.foo = CHECK
}

func (m *Mock) call(){
	m.foo = CALL
}

//phase
func (m *Mock) preflop(){
	m.foo = PREFLOP
}

func (m *Mock) flop(){
	m.foo = FLOP
}

func (m *Mock) turn(){
	m.foo = TURN
}

func (m *Mock) river(){
	m.foo = RIVER
}

func (m *Mock) showdown(){
	m.foo = SHOWDOWN
}

func (m *Mock) endofHand(){
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
	g.add(`(?P<startofHand>^PokerStars Hand #(?P<handNumber>\d+))`)
	g.add(`(?P<seatPlayer>^Seat (?P<SeatNumber>\d): \w+)`)
	g.add(`(?P<postSB>^%w+: posts small blind %d+)`)
	g.add(`(?P<postBB>^%w+: posts big blind %d+)`)
/*
(?P<dealAll>)
(?P<fold>)
(?P<bet>)
(?P<raise>)
(?P<check>)
(?P<call>)
(?P<preflop>)
(?P<flop>)
(?P<turn)
(?P<river)
(?P<showdown>)
(?P<endofHand>)```
*/
	g.endOfAdd()
	fmt.Println(g.names)
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
				fmt.Println(reflect.ValueOf(&reader.line))
				fmt.Println(reflect.ValueOf(&reader.line).NumMethod())
				fmt.Println(reflect.ValueOf(&reader.line).MethodByName(`startofHand`))
				method := reflect.ValueOf(reader.line).MethodByName(key)
				fmt.Println(method)
				if method.IsValid() {
					method.Call([]reflect.Value{})
				}

			}
		}
	}
}


func (reader *PSReader)debug(){
}
