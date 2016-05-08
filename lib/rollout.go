package npoker

import (
	"fmt"
	//"strings"
	"math/rand"
	"sync"
	"time"
)

type Result interface {
	String() string
}

type Runnable interface {
	Clone() Runnable
	Run(source rand.Source) Result
}

type Work interface {
	Lock()
	Unlock()
	Run() Result
}

type RX struct {
	xs      Deck
	board   Deck
	players []Deck
}

func NewRX(board Deck, players []Deck) Runnable {
	xs := BuildFullDeck()
	xs.Subtract(board)
	for _, p := range players {
		xs.Subtract(p)
	}
	return &RX{*xs, board, players}
}

func (x *RX) Clone() Runnable {
	return &RX{
		x.xs.Clone(),
		x.board.Clone(),
		x.players, //  Ugh!
	}
}

func (x *RX) Run(source rand.Source) Result {
	r := rand.New(source)
	x.xs.Shuffle(r)
	x.xs.ShrinkTo(5 - len(x.board))
	river := Join(x.board, x.xs)
	return MakeShowDown(river, x.players...)
}

type Trial struct {
	sync.Mutex
	source rand.Source
	r      Runnable
}

func NewTrial(s rand.Source, r Runnable) Work {
	return &Trial{
		source: s,
		r:      r,
	}
}

func (t *Trial) Run() Result {
	//defer t.Unlock()
	//t.Lock()
	r := t.r.Run(t.source)
	return r
}

func Seeder(count int, wch chan Work, r Runnable) {
	for i := 0; i < count; i++ {
		fmt.Printf("Seeder %d of %d \n", i, count)
		wch <- NewTrial(
			rand.NewSource(time.Now().UnixNano()),
			r.Clone(),
		)
	}
}

func RollOut(trial int, board Deck, players ...Deck) []int {
	stat := make([]int, len(players))

	wch := make(chan Work, 5)
	rch := make(chan Result, 5)

	var wg sync.WaitGroup

	rx := NewRX(board, players)

	go func() {
		for i := 0; i < trial; i++ {
			wch <- NewTrial(
				rand.NewSource(time.Now().UnixNano()),
				rx.Clone(),
			)
		}
		close(wch)
	}()

	for j := 0; j < 4; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				w, ok := <-wch
				if ok {
					//fmt.Printf("Running%d\n", w)
					r := w.Run()
					rch <- r
				} else {
					return
				}
			}
		}()
	}

	go func() {
		for {
			r, ok := <-rch
			if ok {
				sd := r.(*ShowDown)
				for _, w := range sd.Winners {
					stat[w] += 1
				}
			} else {
				return
			}
		}
	}()

	wg.Wait()
	close(rch)
	fmt.Printf("Done.\n %+v", stat)
	return stat
}
