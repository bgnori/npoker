package npoker

import (
	"io"
	"bufio"
	"fmt"
	"os"
)

type Line interface {
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

type PSReader struct {
	pass int
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
	fmt.Println(line)
}

func (reader *PSReader)debug(){
}
