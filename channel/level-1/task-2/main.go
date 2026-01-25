package main

import "sync"

func main() {
	evenChan := make(chan int, 5)
	oddChan := make(chan int, 5)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer close(evenChan)
		defer close(oddChan)

		for i := 0; i <= 10; i++ {
			if i%2 == 0 {
				evenChan <- i
			} else {
				oddChan <- i
			}
		}
	}()

	go func() {
		defer wg.Done()
		for v := range evenChan {
			println("even:", v)
		}
	}()

	go func() {
		defer wg.Done()
		for v := range oddChan {
			println("odd:", v)
		}
	}()

	wg.Wait()
}
