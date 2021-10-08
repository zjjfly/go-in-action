package ch6

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Test_unbuffered_channel(t *testing.T) {
	court := make(chan int)
	wg.Add(2)
	go play("Nodal", court)
	go play("Djokovic", court)
	court <- 1
	wg.Wait()
}

func play(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		court <- ball
	}
}

const (
	numberGoroutines = 4
	taskLoad         = 10
)

func Test_buffered_channel(t *testing.T) {
	tasks := make(chan string, taskLoad)
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}
	//关闭通道
	close(tasks)
	wg.Wait()
}

func worker(task chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-task
		if !ok {
			fmt.Printf("Worker %d : Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d : Started %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
