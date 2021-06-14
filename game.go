package npoker

import (
	"io"
	"bufio"
	"fmt"
	"os"
)

type Line interface {
	//set up
	startofHand()
	seatPlayer()
	setBtn()
	postSB()
	postBB()
	dealAll()
	//actions
	fold()
	bet()
	raise()
	check()
	call()
	//phase
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
}

func NewPSReader() *PSReader {
	return &PSReader{}
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

func (reader *PSReader)feedLine(line string){
	reader.line.startofHand()
}


func (reader *PSReader)debug(){
}
