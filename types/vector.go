package types

import (
	"math"
)

type Vector []int32

func (v Vector) Add(v2 Vector) Vector {
	ret := make(Vector, len(v))
	for idx, d := range v {
		ret[idx] = d + v2[idx]
	}
	return ret
}

func (v Vector) Sub(v2 Vector) Vector {
	ret := make(Vector, len(v))
	for idx, d := range v {
		ret[idx] = d - v2[idx]
	}
	return ret
}

func (v Vector) Distance(v2 Vector) float64 {
	d := int32(0)
	for idx, f := range v {
		diff := f - v2[idx]
		d += diff * diff
	}
	return math.Sqrt(float64(d))
}

func (v Vector) IsAdjacent(v2 Vector) bool {
	if v.Distance(v2) <= 1 {
		return true
	}
	return false
}
