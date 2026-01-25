// LEVEL 1: Channel to channel communication

package main

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "ping"
	}()

	go func() {
		msg := <-ch1
		println(msg)
		ch2 <- "pong"
	}()

	msg := <-ch2
	println(msg)
}
