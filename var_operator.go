package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//byte = int8; rune = in32
	// var x "dtype" = or x := ...
	var x string = "hi"
	z := "01289"
	n := 4.5
	//y, err := strconv.Atoi(z)
	y, _ := strconv.ParseInt(z, 16, 0)
	fmt.Println(y)
	fmt.Printf("\t%s\"", x)
	fmt.Println(math.Sqrt(n))

	fmt.Printf("%T", x[0])
}
