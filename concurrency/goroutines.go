package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func ShowcaseGoroutines() {
	fmt.Println(">>>>Start Goroutines<<<<")
	defer fmt.Println(">>>>End Goroutines<<<<")
	go process()
	go func() {
		time.Sleep(time.Millisecond * 500) // this is bad, don't do this!!!
		fmt.Println("processing2")
	}()
	// sleeping here cos ShowcaseGoroutines exits before the
	// goroutines execute. Channels are needed to fix this
	//time.Sleep(time.Second * 2) // this is bad, don't do this!!!
	var counter = 0

	for i := 0; i < 20; i++ {
		go func() {
			counter++
			fmt.Println(counter)
		}()
	}
	time.Sleep(time.Millisecond * 700) // this is bad, don't do this!!!
	//showcaseDeadlock()
	showcaseGoroutinesWithLock()
}

func showcaseGoroutinesWithLock() {
	fmt.Println("----Start Lock Goroutines----")
	defer fmt.Println("----End Lock Goroutines----")
	go process()
	go func() {
		time.Sleep(time.Millisecond * 500) // this is bad, don't do this!!!
		fmt.Println("processing2")
	}()
	var (
		counter = 0
		lock    sync.Mutex // default value is unlocked. this protects
		// shared var from concurrent access
	)

	for i := 0; i < 20; i++ {
		go func() {
			lock.Lock()
			// only one goroutine can be here at any given time
			defer lock.Unlock() // defer executes right-hand code when function exits
			counter++
			fmt.Println(counter)
		}()
	}
	time.Sleep(time.Millisecond * 700) // this is bad, don't do this!!!
	//showcaseDeadlock()
}

// showcaseDeadlock will cause a deadlock error
func showcaseDeadlock() {
	var lock sync.Mutex

	go func() { lock.Lock() }()
	time.Sleep(time.Millisecond * 10)
	lock.Lock()
}

func process() {
	time.Sleep(time.Millisecond * 600) // this is bad, don't do this!!!
	fmt.Println("processing")
}
