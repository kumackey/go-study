package main

import "fmt"

func main() {
	result := generator()
	for i := 0; i < 5; i++ {
		fmt.Println(<-result)
	}

}

func generator() <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for {
			result <- 1
		}

	}()
	return result
}
