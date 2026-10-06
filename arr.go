package main

import "fmt"

func main() {
	arr := [][2]int{{1, 2}, {2, 3}, {3, 5}}
	// matrix
	for _, value := range arr {
		fmt.Println(value)
	}
	/*fmt.Println("\n")
	for _, nested := range arr {
		for _, value := range nested {
			fmt.Println(value)
		}
	}*/
	mat := [][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	sl := mat[:2] //take row 2
	fmt.Println(sl)

	arr1 := []int{2, 4, 1, 5, 7, 8}
	//sl1 := arr1[1:3] //take index arr[3]
	sl1 := arr1[:4]
	fmt.Println(sl1, len(sl1), cap(sl1))

	str := []string{"nihao", "dajia"}
	for i := 0; i < 10; i++ {
		str = append(str, "yingle")
		fmt.Println(str, len(str), cap(str))
	}
	fmt.Println(str)
	/*cap()- capacity se tang gap doi sau moi lan cap phat dong
	nen begin len=2, cap=2 sau khi tang len=3 thi cap = 4

	make([]Kiểu_Dữ_Liệu, Chiều_Dài_Len, Sức_Chứa_Cap)
	sl := make([]int, 3, 5)
	Kết quả của sl lúc này: [0 0 0]*/
}
