package main

import "sync"

const numProducers = 3

func main() {
	sharedChan := make(chan int, 15)
	done := make(chan struct{})
	var wg sync.WaitGroup

	go func() {
		for v := range sharedChan {
			println(v)
		}
		close(done)
	}()

	wg.Add(numProducers)

	for p := 0; p < numProducers; p++ {
		go func(id int) {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				sharedChan <- i
			}
		}(p)
	}

	wg.Wait()
	close(sharedChan)
	<-done
}
