package main

import "fmt"

type Person struct {
	name     string
	age      uint
	favSport Sport
}
type Sport struct {
	name     string
	position string
}

//Day la ham
func getName(p Person) string {
	return p.name
}

//Day la phuong thuc
func (p Person) getName() string {
	return p.name
}
func (p Person) setName(newName string) {
	p.name = newName
	fmt.Println(p)
}
func main() {
	p1 := Person{age: 18, name: "YingLe", favSport: []Sport{"Badminton", "Single Men"}}
	p1.name = "YingLe"
	fmt.Println(p1.getName())
	p1.setName("ChinYing")
	fmt.Println(p1.favSport[0].position)
}
