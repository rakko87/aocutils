package vector

import (
	"math"
	"slices"
)

type vec []float64

func (p vec) Equal(q vec) bool {
	for i := range p {
		if p[i] != q[i] {
			return false
		}
	}
	return true
}

func (p vec) Add(q vec) vec {
	out := slices.Clone(vec)
	for i := range q {
		out[i] += q[i]
	}
	return out
}

func (p vec) Sub(q vec) vec {
	return p.Add(q.Neg())
}

func (p vec) Neg() (negative vec) {
	for i := range p {
		negative[i] = -1 * p[i]
	}
	return negative
}

func (p vec) Mag() float64 {
	s := float64(0)
	for i := range p {
		s += p[i] * p[i]
	}
	return math.Sqrt(float64(s))
}

func (p vec) Dist(q vec) float64 {
	d := Add(p, q.Neg())
	return d.Mag()
}

func (p vec) Unit() vec {
	s := p.Mag()
	out := slices.Clone(p)
	for i := range out {
		out[i] /= s
	}
	return out
}

func Add(p1, p2 vec) vec {
	s := vec{}
	for i := range p1 {
		s[i] = p1[i] + p2[i]
	}
	return s
}

// Linear interpolation between two vectors.
// scale ∈ [0,1]
func LerpVector(a, b vec, scale float64) vec {
	diff := b.Sub(a)
	var scaleFactor float64 = diff.Magnitude() * scale
	lerp := a.Add(diff.Unit().Scale(scaleFactor))
	return lerp
}
