package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("out of the shapesCapacity range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if b.shapes == nil || len(b.shapes) == 0 {
		return nil, errors.New("shape by index doesn't exist")
	} else if len(b.shapes) < i+1 {
		return nil, errors.New("index went out of the range")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if b.shapes == nil || len(b.shapes) == 0 {
		return nil, errors.New("shape by index doesn't exist")
	} else if len(b.shapes) < i+1 {
		return nil, errors.New("index went out of the range")
	}
	s := b.shapes[i]
	if len(b.shapes) != i+1 {
		copy(b.shapes[i:], b.shapes[i+1:])
	}
	b.shapes = b.shapes[:len(b.shapes)-1]
	return s, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if b.shapes == nil || len(b.shapes) == 0 {
		return nil, errors.New("shape by index doesn't exist")
	} else if len(b.shapes) < i+1 {
		return nil, errors.New("index went out of the range")
	}
	s := b.shapes[i]
	b.shapes[i] = shape
	return s, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, v := range b.shapes {
		sum += v.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, v := range b.shapes {
		sum += v.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	notExist := true
	for i := len(b.shapes) - 1; i >= 0; i-- {
		if _, ok := b.shapes[i].(*Circle); ok {
			_, _ = b.ExtractByIndex(i)
			notExist = false
		}
	}
	if notExist {
		return errors.New("circles are not exist in the list")
	}
	return nil
}
