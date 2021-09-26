package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func generator(ctx context.Context, num int) <-chan int {
	out := make(chan int)

	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				if err := ctx.Err(); errors.Is(err, context.Canceled) {
					fmt.Println("canceled")
				} else if errors.Is(err, context.DeadlineExceeded) {
					fmt.Println("deadline")
				}
				break LOOP
			case out <- num:
			}
		}

		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

func main() {
	// doneチャネルがcloseされたらキャンセル
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	gen := generator(ctx, 1)

	wg.Add(1)

LOOP:
	for i := 0; i < 5; i++ {
		select {
		case result, ok := <-gen: // genから値を受信できた場合
			if ok {
				fmt.Println(result)
			} else {
				fmt.Println("timeout")
				break LOOP
			}
		}
	}
	cancel()

	wg.Wait()
}
