package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	stop2()
}

func d3() {
	before := time.Now()
	limit := make(chan struct{}, 20)
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		i := i
		go func() {
			limit <- struct{}{}
			defer wg.Done()
			u := fmt.Sprintf("http://example.com/id=%d", i)
			downloadJSON(u)
			<-limit
		}()
	}
	wg.Wait()
	fmt.Printf("%v\n", time.Since(before))
}

func d2() {
	before := time.Now()
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			u := fmt.Sprintf("http://example.com/id=%d", i)
			downloadJSON(u)
		}()
	}
	wg.Wait()
	fmt.Printf("%v\n", time.Since(before))
}

func downloadJSON(u string) {
	println(u)
	time.Sleep(time.Second * 6)
}

// 後処理が必要なケース
func stop2() {
	quit := make(chan string)
	ch := generatorBool2("hi", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-ch, i)
	}
	quit <- "Bye!"
	fmt.Printf("Generator says %s", <-quit)

}

func generatorBool2(msg string, quit chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- fmt.Sprintf("%s", msg):
			// nothing
			case <-quit:
				quit <- "See you"
				return
			}
		}
	}()
	return ch
}

// 停止処理の例
func stop() {
	quit := make(chan bool)
	ch := generatorBool("hi", quit)
	for i := rand.Intn(50); i >= 0; i-- {
		fmt.Println(<-ch, i)
	}
	quit <- true
}

func generatorBool(msg string, quit <-chan bool) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- fmt.Sprintf("%s", msg):
			// nothing
			case <-quit:
				fmt.Println("Grountine done")
				return
			}
		}
	}()
	return ch
}

// 全体をtimeout
func sampleGoroutineOverallTimeout() {
	ch := generator("hi")
	timeout := time.After(5 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("5s Timeout!")
			return
		}
	}
}

// timeout処理
func sampleGoroutineTimeout() {
	ch := generator("Hi")
	for i := 0; i < 10; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-time.After(time.Second):
			fmt.Print("Waited too long")
			return
		}
	}
}

func sampleGoroutineWithFanIn2() {
	ch := fanIn2(generator("HELLO"), generator("BYE"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

// 先着処理
func fanIn2(ch1, ch2 <-chan string) <-chan string {
	newCh := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				newCh <- s
			case s := <-ch2:
				newCh <- s
			}
		}
	}()
	return newCh
}

func sampleGoroutine3() {
	ch := fanIn(generator("hello"), generator("bye"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

// 合流処理
func fanIn(ch1, ch2 <-chan string) <-chan string {
	newCh := make(chan string)
	go func() {
		for {
			newCh <- <-ch1
		}
	}()
	go func() {
		for {
			newCh <- <-ch2
		}
	}()
	return newCh
}

func sampleGoroutine2() {
	ch := generator("hello")
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(1 * time.Second)
		}
	}()
	return ch
}

func download() {
	urls := []string{
		"https://www.google.com",
		"https://www.yahoo.co.jp",
		"https://www.bing.com",
	}

	ch := make(chan []byte)
	var wg sync.WaitGroup
	wg.Add(1)
	go downloadCSV(&wg, urls, ch)

	for b := range ch {
		r := csv.NewReader(bytes.NewReader(b))
		for {
			records, err := r.Read()
			if err != nil {
				log.Fatal(err)
			}
			insertRecords(records)
		}
	}
	wg.Wait()
}

func insertRecords(records []string) {
	//
}

func downloadCSV(wg *sync.WaitGroup, urls []string, ch chan []byte) {
	defer wg.Done()
	defer close(ch)
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Println("cannot download CSV:", err)
			continue
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			log.Println("cannot read content:", err)
			continue
		}
		resp.Body.Close()
		ch <- b

	}
}

func doSomething(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Second)
}

func sampleGoroutine() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go doSomething(&wg)
	}
	wg.Wait()
}
