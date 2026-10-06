package main

import "fmt"

func main() {
	//mp := map[string]int{"a": 1}
	//mp1 := map[string][]int{"a": {1, 1, 1}}
	//mp1["b"] = []int{1, 2, 3}
	//mp1["b"] = []int{}
	//delete(mp1, "b")

	mp1 := map[string]bool{}
	mp1["b"] = false

	value, exist := mp1["b"]
	fmt.Println(value, exist)

	mp2 := map[uint]uint{}
	n := uint(100)
	for i := uint(1); i <= n; i++ {
		for j := uint(1); j <= 5; j++ {
			if i%j == 0 {
				mp2[j]++
			}
		}
	}
	fmt.Println(mp2)
}
