package main

import "fmt"

type Number interface {
	uint | float64 | int
}

func SUM[T Number](a T, b T) T {
	return a + b
}

func getValue[K comparable, V any](mp map[K]V) []V {
	values := []V{}
	for _, value := range mp {
		values = append(values, value)
	}

	return values
}

type genericSlice[T any] []T

func (g genericSlice[T]) Print() {
	for _, vl := range g {
		fmt.Println(vl)
	}
}

type genericStruct[T any, K any] struct {
	values T
}

func main() {
	mp := map[string]int{"a": 5, "b": 2, "c": 6, "d": 1}
	fmt.Println("Sum :", SUM(2, 2.2), "\nMap values :", getValue(mp))

	g := genericSlice[int]{1, 2, 3}
	g.Print()

	st := genericStruct[string, int]{values: "string"}
	/*type genericStruct[K any, T any] struct {
	    values T

		Vị trí 1: K nhận kiểu string.
		Vị trí 2: T nhận kiểu int.
		mà values T nhận là string nên sẽ bug
	}*/
	fmt.Println(st)
}
