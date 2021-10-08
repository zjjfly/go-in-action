package ch6

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	shutdown int64
	counter  int64
	wg       sync.WaitGroup
	mutex    sync.Mutex
)

func Test_concurrent(t *testing.T) {
	//设置逻辑处理器的数目
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println("")
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println("")
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

func Test_long_concurrent(t *testing.T) {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Start Goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

func printPrime(prefix string) {
	// 在函数退出时调用Done来通知main函数工作已经完成 37 defer wg.Done()
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
			fmt.Printf("%s:%d\n", prefix, outer)
		}
		fmt.Println("Completed", prefix)
	}
}

//并行,通过增大逻辑处理器数目
func Test_parallel(t *testing.T) {
	runtime.GOMAXPROCS(2)

	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println("")
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println("")
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

func Test_race_condition(t *testing.T) {
	wg.Add(2)
	go incCounter()
	go incCounter()

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter() {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		//给counter增加值的过程不是原子的,所以有竞态条件
		value := counter
		runtime.Gosched()
		value++
		counter = value
	}
}

func Test_atomic(t *testing.T) {
	wg.Add(2)
	go atomicIncCounter()
	go atomicIncCounter()

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func atomicIncCounter() {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}

func Test_store_load(t *testing.T) {
	wg.Add(2)
	go doWork("A")
	go doWork("B")
	time.Sleep(1 * time.Second)
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}

func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}

func Test_mutex(t *testing.T) {
	wg.Add(2)
	go mutexIncCounter()
	go mutexIncCounter()
	wg.Wait()
	fmt.Printf("Final Counter: %d\\n", counter)
}

func mutexIncCounter() {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}
