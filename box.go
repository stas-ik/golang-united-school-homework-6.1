package golang_united_school_homework

import "errors"

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
	if b.shapesCapacity == len(b.shapes) {
		return errors.New("too much")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("not found")
	}
	for k, v := range b.shapes {
		if k == i {
			return v, nil
		}
	}
	return nil, errors.New("not found")
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if shape, err := b.GetByIndex(i); err != nil {
		return nil, err
	} else {
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return shape, nil
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if oldShape, err := b.GetByIndex(i); err != nil {
		return nil, err
	} else {
		b.shapes[i] = shape
		return oldShape, nil
	}
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var a float64
	for _, s := range b.shapes {
		a += s.CalcArea()
	}
	return a
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var p float64
	for _, s := range b.shapes {
		p += s.CalcPerimeter()
	}
	return p
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	count := 0
	var i interface{} = ""
	result := []Shape{}
	for _, r := range b.shapes {
		i = r
		_, ok := i.(*Circle)
		fmt.Println(r)
		if ok {
			count++
		} else {
			result = append(result, r)
		}
	}
	if count == 0 {
		return fmt.Errorf("circles are not exist in the list")
	}
	b.shapes = result
	return nil
}
