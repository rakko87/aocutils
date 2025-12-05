package ranges

import (
	"fmt"
	_ "testing"
)

func Example() {
	r := R{1, 5}
	rr := R{3, 7}
	fmt.Println(
		r, "contains 1:", r.contains(1),
		rr, "contains 1:", rr.contains(1),
		"ranges intersect:", r.intersect(rr),
	)
	// Output:
	// 1-5 contains 1: true 3-7 contains 1: false ranges intersect: true
}
