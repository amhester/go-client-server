package utils

type IVector interface {
	Add(v Vector) Vector
	Sub(v Vector) Vector
	Multiply(v Vector) Vector
	Divide(v Vector) Vector
}

type Vector struct {
	hash           string
	dimensionality int
	dimensions     []float64
}

func NewVector(vals ...float64) Vector {
	return Vector{
		hash:           "", //TODO: implement some spatial hash of the vals
		dimensionality: len(vals),
		dimensions:     vals,
	}
}

func (v Vector) Add(v2 Vector) Vector {
	if v.dimensionality != v2.dimensionality {
		return v
	}
	newDims := make([]float64, v.dimensionality)
	for idx, val := range v.dimensions {
		newDims[idx] = val + v2.dimensions[idx]
	}
	return Vector{
		hash:           "",
		dimensionality: v.dimensionality,
		dimensions:     newDims,
	}
}

func (v Vector) Sub(v2 Vector) Vector {
	if v.dimensionality != v2.dimensionality {
		return v
	}
	newDims := make([]float64, v.dimensionality)
	for idx, val := range v.dimensions {
		newDims[idx] = val - v2.dimensions[idx]
	}
	return Vector{
		hash:           "",
		dimensionality: v.dimensionality,
		dimensions:     newDims,
	}
}

func (v Vector) Multiply(v2 Vector) Vector {
	if v.dimensionality != v2.dimensionality {
		return v
	}
	newDims := make([]float64, v.dimensionality)
	for idx, val := range v.dimensions {
		newDims[idx] = val * v2.dimensions[idx]
	}
	return Vector{
		hash:           "",
		dimensionality: v.dimensionality,
		dimensions:     newDims,
	}
}

func (v Vector) Divide(v2 Vector) Vector {
	if v.dimensionality != v2.dimensionality {
		return v
	}
	newDims := make([]float64, v.dimensionality)
	for idx, val := range v.dimensions {
		newDims[idx] = val / oneIfZero64(v2.dimensions[idx])
	}
	return Vector{
		hash:           "",
		dimensionality: v.dimensionality,
		dimensions:     newDims,
	}
}

func hash(x, y, z float64) string {
	//hashChars := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	//boundsX := 268435456
	//boundsY := 268435456
	//boundsZ := 20480
	//h := ""
	//s := boundsX / 16

	return "" //h
}

func oneIfZero64(val float64) float64 {
	if val == 0 {
		return 1
	}
	return val
}
