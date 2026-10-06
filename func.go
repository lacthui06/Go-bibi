package main

import (
	"fmt"
	"math"
)

func test(num1 int, num2 int) (int, string) {
	return num1 + num2, "hi"
}

/*
Ham bac cao

	func TenHamChinh(tenThamSo func(KiểuDữLiệuVào) KiểuDữLiệuRa) KiểuDữLiệuRaCủaHàmChính {
		// Gọi hàm được truyền vào
		return tenThamSo(giá_trị_nào_đó)
	}
*/
func callFunc(callable func(int) float64) float64 {
	return callable(4)
}
func sum(num int) float64 {
	return math.Pow(float64(num), 2)
}

// Ham bao dong
func getFunc(str string) func(string) string {
	return func(str2 string) string {
		return str + str2
	}
}

// Ham da bien _ Variadic func
func sum2(nums ...int) (sum int, result int) {
	for _, vl := range nums {
		sum += vl
		result = 2 * vl
	}
	return
}

func main() {
	vl, str := test(2, 3)
	fmt.Println(str, "Value :", vl)

	vl1 := callFunc(sum)
	fmt.Printf("Value of callback func: %v", vl1)
	//Ham an _ anonymous func
	value := callFunc(func(x int) float64 {
		return float64(x + 1)
	})
	fmt.Printf("\nValue of anonymous func: %v\n", value)

	str1 := getFunc("ni hao")
	result := str1("ying le")
	result2 := str1("bao bei")
	fmt.Println(result, result2)

	sum2, result3 := sum2([]int{1, 2, 3, 4, 5}...)
	fmt.Println("Sum of arr:", sum2, result3)
}
