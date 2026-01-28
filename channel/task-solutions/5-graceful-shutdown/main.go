package main

import (
	"time"
)

func main() {
	dataChan := make(chan int)
	done := make(chan struct{})

	// producer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		defer close(dataChan)

		i := 0
		for {
			select {
			case <-done:
				println("producer shutting down")
				return
			case <-ticker.C:
				dataChan <- i
				i++
			}
		}
	}()

	// consumer
	go func() {
		for {
			select {
			case <-done:
				println("consumer shutting down")
				return
			case v, ok := <-dataChan:
				if !ok {
					return
				}
				println("consumed:", v)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	close(done)

	// give goroutines time to exit (learning-only)
	time.Sleep(500 * time.Millisecond)
}
