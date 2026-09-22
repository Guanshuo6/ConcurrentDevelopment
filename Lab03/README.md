Student name: Feng Guanshuo
Student number: C00278723
Git link:https://github.com/Guanshuo6/ConcurrentDevelopment
1. Barrier.go

This program creates 10 goroutines.

Each goroutine does:

Part A → Barrier → Part B

The barrier makes sure that all goroutines finish Part A before any goroutine starts Part B.

It uses:

sync.Mutex - protects the counter that records how many goroutines have arrived.
Semaphore - controls when goroutines can pass the barrier.
sync.WaitGroup - waits for all goroutines to finish.

2. Rendezvous.go

This program creates 5 goroutines.

Each goroutine waits for a random amount of time and then reaches:

Part A → Rendezvous → Part B

A channel is used for the rendezvous. All goroutines must reach the rendezvous before they can continue to Part B.

It uses:

Channel - implements the rendezvous.
sync.WaitGroup - waits for all goroutines to finish.

These tasks were completed by Edward and me helping each other based on the examples and videos provided by the teacher.