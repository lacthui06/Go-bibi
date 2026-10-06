package main

import "fmt"

func change(temp *int) {
	*temp = 100
}

type Book struct {
	id    int
	title string
}

func (b *Book) setTitle(title string) {
	(*b).title = title
}
func (b Book) getTitle() string {
	return b.title
}

func Test(pointerSlice *[]*int) {
	values := *pointerSlice
	for _, vl := range values {
		*vl = 10
	}
}

func main() {
	//a := 10
	//change(&a)
	//fmt.Println(a)
	b := Book{1, "hello"}
	(&b).setTitle("nihao")
	fmt.Println(b.getTitle())

	x := 1
	y := 2
	z := 3
	ps := []*int{&x, &y, &z}
	Test(&ps)
	for _, vl := range ps {
		fmt.Println(*vl)
	}
}
