// Package hsd is a package for distance calculation.
package hsd

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// StringDistance returns the distance between two strings.
func StringDistance(lhs, rhs string) int {
	return Distance([]rune(lhs), []rune(rhs))
}

// Distance returns the distance between two runes.
func Distance(a []rune, b []rune) int {
	dist := 0
	if len(a) != len(b) {
		return -1
	}
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}

func Add(a, b int) int {
	// very slow
	time.Sleep(3 * time.Second)
	return a + b
}

func doSomething(s string) string {
	return s
}

func CreateProfile(filename string) (bool, error) {
	fmt.Printf("creating profile: %s\n", filename)
	return true, nil
}

func printDatabaseURL() error {
	url := os.Getenv("DATABASE_URL")
	fmt.Printf("DATABASE_URL: %s\n", url)
	return nil
}

func Calc(v1, v2 int, ope string) (int, error) {
	switch ope {
	case "+":
		return v1 + v2, nil
	case "-":
		return v1 - v2, nil
	case "*":
		return v1 * v2, nil
	case "/":
		return v1 / v2, nil
	}
	return 0, errors.New("")
}
