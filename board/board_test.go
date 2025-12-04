package board

import (
	"fmt"
	"maps"
	"testing"
)

func TestBoard(t *testing.T) {
	b := make(Board)
	s := []string{"abcd"}
	b.Populate(s)
	want := make(map[complex128]rune)
	want[0+1i] = 'a'
	want[1+1i] = 'b'
	want[2+1i] = 'c'
	want[3+1i] = 'd'
	if !maps.Equal(b, want) {
		t.Errorf("Got %v, want %v", b, want)
	}
}

func ExampleBoard() {
	b := make(Board)
	s := []string{"ab", "cd", "ef"}
	b.Populate(s)
	fmt.Printf("%v", b)
	// Output:
	// map[(0+1i):101 (0+2i):99 (0+3i):97 (1+1i):102 (1+2i):100 (1+3i):98]
}

// func ExampleBoard_range() {
// 	b := make(Board)
// 	s := []string{"ab", "cd", "ef"}
// 	b.Populate(s)
// 	for k, v := range b {
// 		fmt.Printf("%v:%v, ", k, string(v))
// 	}
// 	// Output:
// 	// (1+3i):b, (0+2i):c, (1+2i):d, (0+1i):e, (1+1i):f, (0+3i):a,
// }

// func ExampleBoard_Print() {
// 	b := make(Board)
// 	s := []string{"abcd", "efgh"}
// 	b.Populate(s)
// 	b.Print()
// 	// Output:
// 	// abcd
// 	// efgh
// }
