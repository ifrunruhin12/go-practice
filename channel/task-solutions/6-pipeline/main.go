package main

func main() {
	//3 stage pipeline
	genChan := make(chan int)
	squareChan := make(chan int)
	printChan := make(chan int)

	go func() {
		defer close(genChan)
		for i := 1; i <= 10; i++ {
			genChan <- i
		}
	}()
	go func() {
		defer close(squareChan)
		for v := range genChan {
			squareChan <- v * v
		}
	}()
	go func() {
		defer close(printChan)
		for v := range squareChan {
			printChan <- v
		}
	}()

	for v := range printChan {
		println(v)
	}
}
