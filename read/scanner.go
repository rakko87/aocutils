package read

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"unicode/utf8"
)

func Scan(r io.Reader, f bufio.SplitFunc) iter.Seq2[int, string] {
	scanner := bufio.NewScanner(r)
	scanner.Split(f)
	return func(yield func(int, string) bool) {
		for i := 0; scanner.Scan(); i++ {
			if !yield(i, scanner.Text()) {
				return
			}
			if err := scanner.Err(); err != nil {
				yield(i, fmt.Sprintf("error reading input %v", err))
				return
			}
		}
	}
}

// Consume n items from iterator, discarding the results
func consumeSeq2[K any, V any](it iter.Seq2[K, V], n int) {
	for range it {
		n--
		if n <= 0 {
			return
		}
	}
}

func string2strings(s string) []string {
	out := make([]string, utf8.RuneCountInString(s))
	for i, c := range s {
		out[i] = string(c)
	}
	return out
}
