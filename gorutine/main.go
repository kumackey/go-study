package main

import "fmt"

func main() {
	done := make(chan struct{})
	result := generator(done)
	for i := 0; i < 5; i++ {
		fmt.Println(<-result)
	}
	close(done)
}

func generator(done chan struct{}) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
	LOOP:
		for {
			select {
			case <-done:
				break LOOP
			case result <- 1:
			}
		}

	}()
	return result
}
