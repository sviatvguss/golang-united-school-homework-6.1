package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t *Triangle) CalcPerimeter() float64 {
	return t.Side * 3
}

func (t *Triangle) CalcArea() float64 {
	return (t.Side * t.Side * math.Sqrt(3)) / 4
}
