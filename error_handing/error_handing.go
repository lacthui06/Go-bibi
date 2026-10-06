package main

import (
	"fmt"
)

func divide(a int, b int) int {

	return a / b
}
func defferedFunc() {
	r := recover()
	if r == nil {
		fmt.Println("No error")
	} else {
		fmt.Println(r)
	}
}
func main() {
	defer defferedFunc()
	//panic("Program has bug")
	fmt.Println(divide(1, 0))
}
