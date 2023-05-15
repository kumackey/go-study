package main

import "testing"

func Escape() *int {
	s := []int{1, 2, 3, 4, 5}
	y := &s[0]
	return y
}

func NoEscape() int {
	s := []int{1, 2, 3, 4, 5}
	y := s[0]
	return y
}

func BenchmarkEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Escape()
	}
}

func BenchmarkNoEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoEscape()
	}
}
