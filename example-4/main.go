package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	items := int(1e5) // 100,0000 items to process
	start := time.Now()

	// do some work on a set of items
	c := doWork(items) // Kicks off

	// For every expected item, read from the channel
	for i := 0; i < int(items); i++ {

		// This Call blocks until there is an item writen to the channel (line # 41)
		// Take the bytes out of the channel and store in variable
		theBytes := <-c
		fmt.Printf("Recieved: %v | after %v \n ", theBytes, time.Now().Sub(start))
	}

	fmt.Println("total time taken ", time.Now().Sub(start))
}

// Go to the `network` N times (items), and queue that work up by writing to a channel
func doWork(items int) <-chan []byte {
	c := make(chan []byte, 10) // Buffered Channel - allow 10 sends before requiring a receive

	// Function which simulates expensive work by sleeping for a random amount of time, then writes some random bytes
	// to the channel above
	doSomeExpensiveWork := func(i int) {
		// Simulate Random Network Request Response times -> [0,2]
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)

		// create some random bytes for cool output
		c <- []byte{
			byte(i % (rand.Intn(10) + 1)),
			byte(i % (rand.Intn(10) + 1)),
			byte(i % (rand.Intn(10) + 1)),
		}
	}

	// For every item do some expensive work
	for i := 0; i < items; i++ {
		go doSomeExpensiveWork(i)
	}

	// return the channel being written to, back to the caller
	return c
}
