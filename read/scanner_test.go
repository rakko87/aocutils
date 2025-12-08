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

func Example_consumeSeq2() {
	input := strings.NewReader("1\n2")
	it := Scan(input, bufio.ScanLines)
	consumeSeq2(it, 1)
	for i, line := range it {
		fmt.Println(i, line)
	}
	// Output:
	// 0 2
}

func Example_string2strings() {
	got := string2strings("abcdefgh")
	fmt.Println(got)
	// Output:
	// [a b c d e f g h]

}
