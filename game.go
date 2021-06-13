package npoker

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

func (reader *PSReader)feed(text string){
}

func (reader *PSReader)debug(){
}
