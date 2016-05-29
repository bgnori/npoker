package npoker

import (
	//"fmt"
	"math/rand"
	"sync"
	"time"
)

type Result interface {
	String() string
}

type Summary interface {
	String() string
}

type Summarizer interface {
	Zero() Summary
	Fold(Summary, Result) Summary
}

type Runnable interface {
	Clone() Runnable
	Run(source rand.Source) Result
}

type Seeded struct {
	source   rand.Source
	runnable Runnable
}

type Runner interface {
	Run()
	Summary() Summary
}

type runner struct {
	sync.WaitGroup
	sync.Mutex
	n          int
	worker     int
	runnable   Runnable
	summarizer Summarizer
	seeded     chan Seeded
	results    chan Result
	summary    Summary
}

func NewRunner(n int, worker int, r Runnable, s Summarizer) Runner {
	return &runner{
		n:          n,
		runnable:   r,
		summarizer: s,
		worker:     worker,
		seeded:     make(chan Seeded, 5),
		results:    make(chan Result, 100), // Ugh!
	}
}

func (ro *runner) Run() {
	ro.Add(1)
	go func() {
		defer ro.Done()
		for i := 0; i < ro.n; i++ {
			ro.seeded <- Seeded{
				rand.NewSource(time.Now().UnixNano()),
				ro.runnable.Clone(),
			}
		}
		//fmt.Println("closing seeder")
		close(ro.seeded)
	}()

	for i := 0; i < ro.worker; i++ {
		ro.Add(1)
		go func() {
			defer ro.Done()
			for s := range ro.seeded {
				//fmt.Printf("starting %+v\n", s.runnable)
				ro.results <- s.runnable.Run(s.source)
			}
		}()
	}

	ro.Add(1)
	go func() {
		defer ro.Done()
		x := ro.summarizer.Zero()
		for i := 0; i < ro.n; i++ {
			r := <-ro.results
			//fmt.Printf("got result %+v\n", r)
			x = ro.summarizer.Fold(x, r)
		}
		ro.Lock()
		//fmt.Println("time to write back result")
		ro.summary = x
		ro.Unlock()
	}()

	ro.Wait()
	close(ro.results)
	return
}

func (ro *runner) Summary() Summary {
	ro.Lock()
	s := ro.summary
	ro.Unlock()
	return s
}
