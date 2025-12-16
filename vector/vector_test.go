package vector

import (
	"fmt"
	"slices"
	"testing"
)

func Test_vec_add(t *testing.T) {
	p1 := vec{162, 817, 812}
	p2 := vec{425, 690, 689}
	want := vec{587, 1507, 1501}
	got := Add(p1, p2)
	if !want.Equal(got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Example_negative() {
	p1 := vec{162, 817, 812}
	fmt.Println("-1*", p1, " = ", p1.Neg())
	// Output:
	// -1* [162 817 812]  =  [-162 -817 -812]
}

func Example_dist() {
	p1 := vec{162, 817, 812}
	p2 := vec{425, 690, 689}
	d := p1.Dist(p2)
	fmt.Printf("distance %v and %v is %.1f\n", p1, p2, d)
	// Output:
	// distance [162 817 812] and [425 690 689] is 316.9
}

func TestLerpVector(t *testing.T) {
	cases := []struct {
		scale float64
		want  vec
	}{
		{0, vec{200, 100, 20}},
		{1, vec{100, 254, 60}},
		{0.5, vec{150, 177, 40}}}
	for _, tt := range cases {
		got := LerpVector(vec{200, 100, 20}, vec{100, 254, 60}, tt.scale)
		want := tt.want
		// got := LerpVector(vec{200, 100, 20}, vec{100, 254, 60}, 0)
		// want := vec{200, 100, 20}
		if !slices.Equal(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	}
}
