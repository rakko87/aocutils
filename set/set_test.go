package set

import (
	"fmt"
	"maps"
	"slices"
	_ "testing"
)

func Example_newSet() {
	// Output:
	//
}

func Example_newSet1() {
	fmt.Printf("%#v\n", newSet[float64]())
	fmt.Printf("%#v\n", newSet[string]())
	// Output:
	//set.Set[float64]{}
	// set.Set[string]{}
}

func Example_intersection() {
	l := make(Set[string])
	l.Add("in left")
	l.Add("in both")
	r := make(Set[string])
	r.Add("in right")
	r.Add("in both")
	fmt.Println(Intersection(l, r))
	// Output:
	// map[in both:{}]
}

func Example_pop() {
	m := make(Set[int])
	m.Add(1)
	m.Add(2)
	m.Add(3)
	got := m.Pop()
	fmt.Printf("got %T. Remaining length is %d\n", got, len(m))
	// Output:
	// got int. Remaining length is 2
}

func Example() {

	m := make(Set[string])
	m.Add("one")
	fmt.Printf("%#v\n", m.Keys())

	fmt.Println(m.Has("one"))
	fmt.Println(m.Has("two"))
	fmt.Println(m.Remove("one"))
	fmt.Printf("%#v\n", m.Keys())

	n := make(Set[string])
	n.Add("In n")
	m.Add("In m")
	union := Union(m, n)
	fmt.Printf("%#v\n", union.Keys())
	union = m.Union(n)
	fmt.Printf("%#v\n", union.Keys())

	fmt.Printf("%#v\n", ToSet([]int{1, 2, 3}).Keys())
	fmt.Printf("%#v\n", ToSet([]float32{1, 2, 3}).Keys())

	fmt.Println(splittedSets([]string{"In left", "In both"}, []string{"In right", "In both"}))
	s1, s2, s3 := splittedSets([]string{"In left", "In both"}, []string{"In right", "In both"})
	got := Union(Union(s1, s2), s3)
	got1 := slices.Sorted(maps.Keys(got))
	fmt.Printf("%#v\n", got1)
	fmt.Println(SetOrdered(got))
	// Output unordered:
	// []string{"one"}
	// true
	// false
	// true
	// []string(nil)
	// []string{"In m", "In n"}
	// []string{"In m", "In n"}
	// []int{1, 2, 3}
	// []float32{1, 2, 3}
	// map[In left:{}] map[In right:{}] map[In both:{}]
	// []string{"In both", "In left", "In right"}
	// [In both In left In right]
}
