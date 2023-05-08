// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main() {

// 	fmt.Println("Race Condition")
// 	fmt.Println("CPUs :", runtime.NumCPU())
// 	fmt.Println("Goroutine :", runtime.NumGoroutine())

// 	counter := 0
// 	const gs = 100

// 	var wg sync.WaitGroup
// 	wg.Add(gs)

// 	var mu sync.Mutex

// 	for i := 0; i < gs; i++ {
// 		go func() {
// 			mu.Lock()
// 			v := counter
// 			//time.Sleep(time.Second)
// 			runtime.Gosched()
// 			v++
// 			counter = v
// 			mu.Unlock()
// 			wg.Done()

// 		}()
// 		fmt.Println("Goroutine :", runtime.NumGoroutine())
// 	}
// 	wg.Wait()
// 	fmt.Println("Goroutine :", runtime.NumGoroutine())
// 	fmt.Println("count", counter)

// }

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("Race Condition")
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutine :", runtime.NumGoroutine())

	var counter int64

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("counter", atomic.LoadInt64(&counter))
			//time.Sleep(time.Second)

			wg.Done()

		}()
		fmt.Println("Goroutine :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutine :", runtime.NumGoroutine())
	fmt.Println("count", counter)

}
