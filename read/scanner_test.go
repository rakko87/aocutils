package read

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
	"testing"
)

func Test_scan(t *testing.T) {
	input := strings.NewReader("1\n2")
	want := []string{"1", "2"}

	got := make([]string, 0)
	for _, line := range Scan(input, bufio.ScanLines) {
		got = append(got, line)
	}
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func Example_scan() {
	input := strings.NewReader("1\n2")
	for i, line := range Scan(input, bufio.ScanLines) {
		fmt.Println(i, line)
	}
	// Output:
	// 0 1
	// 1 2
}
