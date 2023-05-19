package hsd

import (
	"github.com/google/go-cmp/cmp"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestStringDistance(t *testing.T) {
	got := StringDistance("foo", "foh")
	want := 1
	if got != want {
		t.Fatalf("expected %d, but got %d", want, got)
	}
}

func TestStringDistance2(t *testing.T) {
	tests := []struct {
		name string
		lhs  string
		rhs  string
		want int
	}{
		{name: "lhs is longer than rhs", lhs: "foo", rhs: "fo", want: -1},
		{name: "lhs is shorter than rhs", lhs: "fo", rhs: "foo", want: -1},
		{name: "No diff", lhs: "foo", rhs: "foo", want: 0},
		{name: "1 diff", lhs: "foo", rhs: "foh", want: 1},
		{name: "2 diff", lhs: "foo", rhs: "fhh", want: 2},
		{name: "3 diff", lhs: "foo", rhs: "bar", want: 3},
		{name: "multibyte", lhs: "あいう", rhs: "あいえ", want: 1},
	}

	for _, tc := range tests {
		got := StringDistance(tc.lhs, tc.rhs)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected %v, but got %v", tc.name, tc.want, got)
		}
	}
}

func TestMain(m *testing.M) {
	log.Println("_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/")
	ret := m.Run()
	log.Println("_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/")
	os.Exit(ret)
}

func TestA(t *testing.T) {
	log.Println("test A")
}

func TestB(t *testing.T) {
	log.Println("test B")
}

func TestShort(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	log.Println("test Short")
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		lhs  int
		rhs  int
		want int
	}{
		{name: "test1", lhs: 0, rhs: 1, want: 1},
		{name: "test2", lhs: 1, rhs: -1, want: 0},
		{name: "test3", lhs: 2, rhs: 1, want: 3},
	}

	for _, tc := range tests {
		tc = tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := Add(tc.lhs, tc.rhs)
			if got != tc.want {
				t.Errorf("expected %d, but got %d", tc.want, got)
			}
		})
	}
}

func TestDoSomething(t *testing.T) {
	fns, err := filepath.Glob("testdata/*.dat")
	if err != nil {
		t.Fatal(err)
	}

	for _, fn := range fns {
		t.Log(fn)
		b, err := os.ReadFile(fn)
		if err != nil {
			t.Fatal(err)
		}

		got := doSomething(string(b))
		b, err = os.ReadFile(fn[:len(fn)-4] + ".out")
		if err != nil {
			t.Fatal(err)
		}
		want := string(b)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("diff: %s", diff)
		}
	}
}

func TestCreateProfile(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "profile.json")
	got, err := CreateProfile(filename)
	if err != nil {
		t.Fatal(err)
	}
	want := true
	if got != want {
		t.Fatalf("expected %v, but got %v", want, got)
	}
}

func TestCreateProfile2(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	err := printDatabaseURL()
	if err != nil {
		t.Fatalf("expected nil, but got %v", err)
	}
}

func FuzzCalc(f *testing.F) {
	f.Add(1, 2, "+")
	f.Fuzz(func(t *testing.T, v1, v2 int, ope string) {
		_, _ = Calc(v1, v2, ope)
	})
}
