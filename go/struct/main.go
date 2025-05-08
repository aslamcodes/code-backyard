package main

import (
	"fmt"
)

type person struct {
	age  int
	name string
}

func main() {
	var pers1 person

	pers1.age = 10
	pers1.name = "Yahya"

	fmt.Println(pers1.name)
	fmt.Println(pers1.age)
}
