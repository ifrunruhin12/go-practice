package main

import (
	"fmt"
	"sync"
)

type worker struct {
	id        int
	processed int
}

const (
	numWorkers = 5
	numJobs    = 20
)

func main() {
	oniChan := make(chan int)
	var wg sync.WaitGroup

	allWorkers := make([]*worker, numWorkers)

	for i := 0; i < numWorkers; i++ {
		w := &worker{id: i + 1}
		allWorkers[i] = w
		wg.Add(1)
		go func(w *worker) {
			defer wg.Done()
			for v := range oniChan {
				fmt.Println("worker:", w.id, "processing job", v)
				w.processed++
			}
		}(w)
	}

	for i := 0; i < numJobs; i++ {
		oniChan <- i
	}

	close(oniChan)
	wg.Wait()

	fmt.Println("\nSummary:")
	for _, w := range allWorkers {
		fmt.Println("worker:", w.id, "processed", w.processed, "jobs")
	}
	fmt.Println("All jobs done!")
}
