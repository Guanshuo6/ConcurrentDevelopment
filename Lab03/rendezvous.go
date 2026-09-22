package main

import (
	"fmt"          // Used to print information
	"math/rand/v2" // Used to generate random numbers
	"sync"         // Provides WaitGroup
	"time"         // Used to make goroutines wait
)

// WorkWithRendezvous is the function executed by each goroutine
func WorkWithRendezvous(
	wg *sync.WaitGroup, // Pointer to WaitGroup used to wait for all goroutines
	Num int, // ID of the current goroutine
	barrier chan bool, // Channel used to implement the Rendezvous
) bool {

	// Notify the WaitGroup when this function finishes
	// to indicate that the current goroutine is complete
	defer wg.Done()

	// Create a variable of type time.Duration
	var X time.Duration

	// Generate a random integer between 0 and 4
	X = time.Duration(rand.IntN(5))

	// Wait for a random amount of time between 0 and 4 seconds
	// This allows different goroutines to reach Part A at different times
	time.Sleep(X * time.Second)

	// Print Part A
	fmt.Println("Part A", Num)

	// Rendezvous starts here

	// Tell the other goroutines:
	// "I have completed Part A and reached the Rendezvous"
	barrier <- true

	// Wait here
	// If not all goroutines have arrived yet,
	// the coordinator will not send a signal, so this goroutine waits
	<-barrier

	// Rendezvous ends here

	// Only after all goroutines have completed Part A
	// can they continue to Part B
	fmt.Println("PartB", Num)

	// Return true
	return true
}

func main() {

	// Create a WaitGroup
	// Used to wait for all goroutines to finish
	var wg sync.WaitGroup

	// There are 5 goroutines
	threadCount := 5

	// Tell the WaitGroup that 5 goroutines will be started
	wg.Add(threadCount)

	// Create a buffered channel
	//
	// The channel has a capacity of 5
	//
	// This channel is used to implement the Rendezvous
	barrier := make(chan bool, threadCount)

	// Create a coordinator goroutine
	//
	// The coordinator's tasks are:
	// 1. Wait for all 5 goroutines to complete Part A
	// 2. Allow all goroutines to continue to Part B
	go func() {

		// Wait for 5 goroutines to reach the Rendezvous
		for i := 0; i < threadCount; i++ {

			// Receive a signal from the channel
			//
			// Each true value means that one goroutine
			// has completed Part A
			<-barrier
		}

		// If the program reaches this point,
		// all 5 goroutines have completed Part A
		//
		// They can now continue to Part B

		// Send 5 signals to the 5 goroutines
		for i := 0; i < threadCount; i++ {

			// Send a signal
			//
			// Each true value allows one waiting goroutine
			// to continue
			barrier <- true
		}

	}()

	// Create 5 goroutines
	for N := range threadCount {

		// Start WorkWithRendezvous
		go WorkWithRendezvous(
			&wg,     // Pass the address of the WaitGroup
			N,       // Pass the current goroutine ID
			barrier, // Pass the Rendezvous channel
		)
	}

	// Wait for all 5 goroutines to finish
	//
	// Note:
	// The WaitGroup is responsible for waiting for the entire task to finish.
	// It does not implement the Rendezvous between Part A and Part B.
	wg.Wait()

	// Print this message after all goroutines have finished
	fmt.Println("All goroutines finished")
}
