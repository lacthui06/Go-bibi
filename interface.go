package main

import "fmt"

type Shape interface {
	getPerimeter() uint
}

type Triangle struct {
	a uint
	b uint
	c uint
}
type Square struct {
	width uint
}

func (t Triangle) getPerimeter() uint {
	return t.a + t.b + t.c
}

//write a struct print slice
func (t Triangle) getSlice() []uint {
	return []uint{t.a, t.b, t.c}
}

func (s Square) getPerimeter() uint {
	return 4 * s.width
}

func main() {
	var s []Shape = []Shape{Triangle{1, 2, 3}, Square{10}}
	//s := []Shape{Triangle{1, 2, 3}, Square{10}}

	perimeterTotal := uint(0)
	// var perimeterTotal uint
	for _, vl := range s {
		perimeterTotal += vl.getPerimeter()
	}
	fmt.Println(perimeterTotal)

	t := Triangle{1, 2, 3}
	fmt.Println(t.getSlice())
}
