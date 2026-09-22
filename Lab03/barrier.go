// Barrier.go Template Code
// Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// --------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 21/9/2024
// Modified by: Feng Guanshuo
// Issues:
// The barrier is not implemented!
// --------------------------------------------

package main // The current program belongs to the main package

import (
	"context" // Used to create the context required by the semaphore
	"fmt"     // Used for printing output
	"sync"    // Provides Mutex and WaitGroup
	"time"    // Used for Sleep

	"golang.org/x/sync/semaphore" // Used for the Weighted Semaphore
)

// arrived keeps track of the number of goroutines
// that have completed Part A
var arrived int

// Place a barrier in this function -- use Mutex's and Semaphores
func doStuff(
	goNum int, // The ID of the current goroutine
	wg *sync.WaitGroup, // Pointer to the WaitGroup
	theLock *sync.Mutex, // Pointer to the Mutex
	sem *semaphore.Weighted, // Pointer to the Semaphore
	totalRoutines int, // Total number of goroutines
	ctx context.Context, // Context used by the semaphore
) bool {

	defer wg.Done() // Tell the WaitGroup that this goroutine has finished

	time.Sleep(time.Second) // Each goroutine waits for 1 second first

	fmt.Println("Part A", goNum) // Execute and print Part A

	// Barrier starts here

	theLock.Lock() // Lock arrived to prevent simultaneous modifications

	arrived++ // Record that another goroutine has completed Part A

	// Check if this is the last goroutine to arrive at the barrier
	if arrived == totalRoutines {

		// The last goroutine has arrived.
		// Release all semaphore permits.
		sem.Release(int64(totalRoutines))
	}

	theLock.Unlock() // Unlock the Mutex after updating arrived

	// Try to acquire one semaphore permit.
	// If the barrier is not open yet, this will wait.
	sem.Acquire(ctx, 1)

	// Release the semaphore permit after passing the barrier.
	// This allows other goroutines to continue.
	sem.Release(1)

	// Barrier ends here

	fmt.Println("PartB", goNum) // All goroutines pass the barrier before reaching here

	return true // Return true
}

// Main function
func main() {

	totalRoutines := 10 // Create a total of 10 goroutines

	var wg sync.WaitGroup // Create a WaitGroup

	wg.Add(totalRoutines) // Tell the WaitGroup to wait for 10 goroutines

	ctx := context.TODO() // Create a context

	var theLock sync.Mutex // Create a Mutex

	// Create a semaphore with a maximum of 10 permits
	sem := semaphore.NewWeighted(int64(totalRoutines))

	// Initially acquire all 10 semaphore permits.
	// This prevents any goroutine from passing the barrier.
	sem.Acquire(ctx, int64(totalRoutines))

	// Create 10 goroutines
	for i := range totalRoutines {

		// Start a goroutine to execute doStuff
		go doStuff(
			i,             // ID of the current goroutine
			&wg,           // Pass the address of the WaitGroup
			&theLock,      // Pass the address of the Mutex
			sem,           // Pass the Semaphore
			totalRoutines, // Pass the total number of goroutines
			ctx,           // Pass the context
		)
	}

	// The main function waits here
	// until all goroutines call wg.Done()
	wg.Wait()

	// Print this message after all goroutines have finished
	fmt.Println("All goroutines finished")
}
