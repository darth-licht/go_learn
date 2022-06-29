package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func ShowcaseChannels() {
	fmt.Println(">>>>Channels<<<<")
	defer fmt.Println(">>>>End Channels<<<<")
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go Worker{i}.process(c)
	}

	for i := 0; i < 20; i++ {
		c <- rand.Int()
		// use close(c) this to close the channel c
		// usually this is done only to signal to receivers that
		// there'll no longer be values coming through the channel
		// close(c)
		time.Sleep(time.Millisecond * 50)
	}

	showcaseBufferChannels()
	showcaseSelectChannels()
	showcaseTimeLimitChannels()
	showcaseCounterWithChannel()
}

func showcaseBufferChannels() {
	fmt.Println("----BufferChannels----")
	defer fmt.Println("----End BufferChannels----")
	// buffered channel with size 100
	// this will store unreceived data temporarily until receivers
	// receive them
	c := make(chan int, 100)
	for i := 0; i < 10; i++ {
		go Worker{i}.processSleeper(c)
	}

	for i := 0; i < 20; i++ {
		c <- i
		fmt.Printf("Buffer length %d\n", len(c))
		time.Sleep(time.Millisecond * 50)
	}
}

func showcaseSelectChannels() {
	fmt.Println("----Select Channels----")
	defer fmt.Println("----End Select Channels----")
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go Worker{i}.processSleeper(c)
	}

	for i := 0; i < 20; i++ {
		// use select to do something else when
		// no receiver is available
		select {
		case c <- i:
		// receiver was available
		default:
			// no receiver was available
			fmt.Printf("dropped %d\n\n", i)
		}
		time.Sleep(time.Millisecond * 50)
	}
}

func showcaseTimeLimitChannels() {
	fmt.Println("----Timed Channels----")
	defer fmt.Println("----End Timed Channels----")
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go Worker{i}.processTimedSleeper(c)
	}

	for i := 0; i < 20; i++ {
		// select can also be used to time out a channel
		// select will pick the first available channel regardless whether
		// were sending to or receiving from the channel
		// if multiple channels are available, one is picked randomly
		select {
		case c <- i:
		// receiver was available
		//case t := <-time.After(time.Millisecond): // valid also
		case <-time.After(time.Millisecond):
			// time.After returns a channel
			// timed out
			fmt.Printf("Timed out with %d\n", i)
		}
		time.Sleep(time.Millisecond * 50)
	}
}

func showcaseCounterWithChannel() {
	var counter = 0
	c := make(chan any)
	for i := 0; i < 20; i++ {
		go Worker{i}.signalIncrement(c)
	}

	for i := 0; i < 20; i++ {
		<-c
		counter++
		fmt.Printf("Counter is %d\n", counter)
	}
}

type Worker struct {
	id int
}

func (w Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
	}
}

func (w Worker) processSleeper(c chan int) {
	for {
		// data, ok := <-c // is also valid
		// (ok) signifies whether the channel is not closed
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		// this will make worker unavailable to receive data and thus
		// block the sender.
		time.Sleep(time.Millisecond * 600)
	}
}

func (w Worker) processTimedSleeper(c chan int) {
	for {
		select {
		case data := <-c:
			fmt.Printf("worker %d got %d\n", w.id, data)
		case <-time.After(time.Millisecond * 10):
			fmt.Println("Break time")
			time.Sleep(time.Second)
		}
	}
}

func (w Worker) signalIncrement(c chan any) {
	for i := 0; i < 5; i++ {
		select {
		case <-time.After(time.Millisecond * (500 * time.Duration(i))):
			c <- nil
		}
	}
}
