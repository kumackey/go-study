package main

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}
	if post.Content != "Hello World!" {
		t.Error("Wrong contents, was expectiong 'Hello World!' but got", post.Content)
	}
}

func TestLongRunningTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode")
	}
	time.Sleep(5 * time.Second)
}

func TestParallelTest(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}

func TestParallel2Test(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestParallel3Test(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}

func BenchmarkUnmarchal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarchal("post.json")
	}
}
