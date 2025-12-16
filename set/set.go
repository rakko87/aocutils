// Practice generics by creating a Set type
package set

import (
	"cmp"
	"fmt"
	"iter"
	"maps"
	"slices"
)

// 2025 november, practice on generics
// Set is just a map where we ignore the values and only care about the keys.
type Set[E comparable] map[E]unary

// semaphore. "Something exist"
type unary struct{}

func newSet[E comparable]() Set[E] {
	return make(Set[E])
}

// tranlations of set methods to map functionality
func (m Set[E]) Add(e E) {
	m[e] = unary{}
}

func (m Set[E]) Has(e E) bool {
	_, ok := m[e]
	return ok
}

func (m Set[E]) Remove(e E) (success bool) {
	delete(m, e)
	_, ok := m[e]
	return !ok
}

func (m Set[E]) Pop() E {
	var item E
	for item := range m {
		if m.Remove(item) {
			break
		}
	}
	return item
}

func (m Set[E]) Keys() []E {
	return slices.Collect(maps.Keys(m))
}

func (m Set[E]) Union(n Set[E]) (union Set[E]) {
	return Union(m, n)
}

// func (m Set[E]) String() string {
// 	s := make([]string, len(m))
// 	s[0] = fmt.Sprintf("%T", m)
// 	i := 1
// 	for k := range m {
// 		s[i] = fmt.Sprintf("%v", k)
// 		i++
// 	}
// 	return strings.Join(s, " ")
// }

func Union[E comparable](m, n Set[E]) (union Set[E]) {
	union = maps.Clone(m)
	maps.Copy(union, n)
	return union
}

func Intersection[E comparable](l, r Set[E]) Set[E] {
	common := make(Set[E])
	for k1 := range l {
		if _, ok := r[k1]; ok {
			// Key exists in both
			common[k1] = unary{}
		} // else // only in left
	}
	return common
}

// Convert slice to a set
func ToSet[S ~[]E, E comparable](s S) Set[E] {
	m := make(Set[E])
	for _, item := range s {
		m[item] = unary{}
	}
	return m
}

// Using the more constrained type cmp.Ordered
func SetOrdered[E cmp.Ordered](m Set[E]) []E {
	return slices.Sorted(maps.Keys(m))
}

// split two slices into left,right and union sets
func splittedSets[S ~[]E, E comparable](left, right S) (l, r, u Set[E]) {
	// Left, Right and union sets
	l, r, u = make(Set[E]), make(Set[E]), make(Set[E])

	// Find set for left and union
	for _, x1 := range left {
		onlyLeft := true
		for _, x2 := range right {
			if x1 == x2 {
				u[x1] = struct{}{}
				onlyLeft = false
				break
			}
		}
		// Found no x1 in right => x1 is only in left slice.
		if onlyLeft {
			l[x1] = struct{}{}
		}
	}

	// Find the only right set
	for _, x2 := range right {
		_, inBoth := u[x2]
		if !inBoth {
			r[x2] = struct{}{}
		}
	}

	return l, r, u
}

// something entirely different
func compareSets[S ~[]E, E comparable](a, b S) (both iter.Seq[E], A iter.Seq[E], B iter.Seq[E]) {
	// Sets
	onlyA := make(map[E]unary)
	onlyB := make(map[E]unary)
	union := make(map[E]unary)

	for _, k1 := range a {
		exists := slices.Index(b, k1) > -1
		if exists {
			union[k1] = unary{}
		} else {
			onlyA[k1] = unary{}
		}
	}

	for _, k2 := range b {
		_, ok := union[k2]
		if !ok {
			onlyB[k2] = unary{}
		}
	}
	return maps.Keys(union), maps.Keys(onlyA), maps.Keys(onlyB)
}

func oldSplitToSets() {
	fmt.Println("hello from old")
	a, b := []string{"in a", "in both"}, []string{"in b", "in both"}
	uu, aa, bb := compareSets(a, b)
	fmt.Println(
		slices.Collect(uu),
		slices.Collect(aa),
		slices.Collect(bb),
	)
}
