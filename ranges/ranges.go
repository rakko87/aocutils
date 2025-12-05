package ranges

import "fmt"

type R struct{ fog, tog int }

func (r R) contains(v int) bool {
	return (v >= r.fog) && (v <= r.tog)
}

func (r R) intersect(rr R) bool {
	return r.contains(rr.fog) || r.contains(rr.tog)
}

func (r R) equal(rr R) bool {
	return (r.fog == rr.fog) && (r.tog == rr.tog)
}

func (r R) String() string {
	return fmt.Sprintf("%v-%v", r.fog, r.tog)
}
