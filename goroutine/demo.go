package main

import (
	"fmt"
	"time"
)

func count() {
	time.Sleep(2 * time.Second)
	fmt.Println("1")
}
func count1() {
	time.Sleep(4 * time.Second)
	fmt.Println("2")
}
func count2() {
	time.Sleep(6 * time.Second)
	fmt.Println("3")
}

func add(a int, b int, ch chan int) {
	ch <- a + b
}
func add1(a int, b int, ch chan int) {
	ch <- a + b
}
func add2(a int, b int, ch chan int) {
	ch <- a + b
}
func main() {
	ch := make(chan int)
	go add(1, 0, ch)
	x := <-ch
	fmt.Println(x)

	go add(1, 1, ch)

	fmt.Println(x)

	go add(1, 2, ch)
	x = <-ch
	x = <-ch
	fmt.Println(x)
	//time.Sleep(6 * time.Second)
	fmt.Println("Done")
}
