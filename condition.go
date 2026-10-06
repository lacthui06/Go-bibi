package main

import "fmt"

func main() {
	a := 1

	t := uint(8)
	u := uint(a)

	if t < u {
		fmt.Printf("%v < %v is wrong\n", t, u)
	} else if t > u {
		fmt.Printf("%v > %v is corect\n", t, u)
	} else {
		fmt.Printf("%v = %v is corect\n", t, u)
	}

	switch a {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("fallthrough la nhay xuong dong")
	default:
		fmt.Println("nill")
	}

}
