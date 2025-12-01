package read

import (
	"bufio"
	"fmt"
	"io"
	"iter"
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
