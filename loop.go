package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println("i love u")
	}
	a := 0
	for a < 3 {
		fmt.Println("Warning! Hacker Access***")
		a++
	}

	str := "hacker master"
	for _, char := range str {
		fmt.Printf("%c", char)
	}
}
