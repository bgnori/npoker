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

type Experiment interface {
	Run(int)
	Report([]Deck)
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

type rollout struct {
	sync.WaitGroup
	trial    int
	runnable Runnable
	works    chan Work
	results  chan Result
	acc      interface{}
}

func NewRollOut(trial int, r Runnable) Experiment {
	return &rollout{
		trial:    trial,
		runnable: r,
		works:    make(chan Work, 5),
		results:  make(chan Result, trial), // Ugh!
	}
}

func (ro *rollout) Run(nworker int) {
	ro.Add(1)
	go func() {
		defer ro.Done()
		for i := 0; i < ro.trial; i++ {
			ro.works <- NewTrial(
				rand.NewSource(time.Now().UnixNano()),
				ro.runnable.Clone(),
			)
		}
		close(ro.works)
	}()

	for i := 0; i < nworker; i++ {
		ro.Add(1)
		go func() {
			defer ro.Done()
			for {
				w, ok := <-ro.works
				if ok {
					r := w.Run()
					ro.results <- r
				} else {
					return
				}
			}
		}()
	}

	ro.Wait()
	close(ro.results)
	return
}

func (ro *rollout) Report(players []Deck) {
	stat := make([]int, len(players))
	for r := range ro.results {
		r := r.(*ShowDown)
		for i, amount := range DistrubuteChips(1000, 1, 0, r) {
			stat[i] += amount
		}
	}
	for i, v := range stat {
		fmt.Printf("player %d has %s, won %d times.\n", i, players[i], v/1000.0)
	}
}
