package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	sharedRscOK()
}

func sharedRscOK() {
	var sharedRsc = make(map[string]interface{})
	var wg sync.WaitGroup
	var mu sync.Mutex
	c := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		c.L.Lock()
		for len(sharedRsc) == 0 {
			fmt.Println("Wait...")
			c.Wait()
		}

		fmt.Println("Reading...")
		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()

	// writes changes to sharedRsc
	time.Sleep(1 * time.Microsecond)
	fmt.Println("Writing...")
	c.L.Lock()
	sharedRsc["rsc1"] = "foo"
	c.Signal()
	c.L.Unlock()

	wg.Wait()
}

func sharedRscNG() {
	var sharedRsc = make(map[string]interface{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for len(sharedRsc) == 0 {
			fmt.Println("Wait...")
		}

		fmt.Println("Reading...")
		fmt.Println(sharedRsc["rsc1"])
	}()

	// writes changes to sharedRsc
	time.Sleep(1 * time.Microsecond)
	fmt.Println("Writing...")
	sharedRsc["rsc1"] = "foo"

	wg.Wait()
}

func livelock() {
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, "%v", dirName)
		atomic.AddInt32(dir, 1)
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ".Success!")
			return true
		}
		takeStep()
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }
	tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }

	_, _ = tryLeft, tryRight
}

func deadlock() {
	type value struct {
		mu    sync.Mutex
		value int
	}

	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(1 * time.Second)

		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}

func mutex() {
	var memmoryAccess sync.Mutex
	var data int
	go func() {
		memmoryAccess.Lock()
		data++
		memmoryAccess.Unlock()
	}()

	memmoryAccess.Lock()
	if data == 0 {
		fmt.Println("this value is 0")
	} else {
		fmt.Printf("this value is %d\n", data)
	}
	memmoryAccess.Unlock()
}
