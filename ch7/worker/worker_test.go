package worker

import (
	"log"
	"sync"
	"testing"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(300 * time.Millisecond)
}

func Test_worker(t *testing.T) {
	pool := New(10)
	var wg sync.WaitGroup
	wg.Add(100 * len(names))
	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				pool.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	pool.Shutdown()
}
