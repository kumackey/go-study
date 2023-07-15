package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"sync"
)

func main() {
	multitask()
}

func multitask() {
	err := doMultiTask([]string{"a", "b", "c"})
	if err != nil {
		if errs, ok := err.(interface{ Unwrap() []error }); ok {
			for _, e := range errs.Unwrap() {
				log.Println(e)
			}
		} else {
			log.Println(err)
		}
	}
}

func doMultiTask(files []string) error {
	var mu sync.Mutex
	var wg sync.WaitGroup
	var errs []error
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()

			f, err := os.Open(file)
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
			} else {
				defer f.Close()
				// do something
			}
		}(file)
	}
	wg.Wait()

	return errors.Join(errs...)
}

func newfunc() {
	endpoint := "https://example.com"
	endpoint, err := url.JoinPath(endpoint, "api", "v1", "users")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(endpoint)
}

func classic() {
	endpoint := "https://example.com"
	u, err := url.Parse(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = path.Join(u.Path, "api", "v1", "users")
	end := u.String()
	fmt.Println(end)
}
