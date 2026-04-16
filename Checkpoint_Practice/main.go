package main

import (
	one "checkpoint_practice/piscine/level_1"
	two "checkpoint_practice/piscine/level_2"
	four "checkpoint_practice/piscine/level_4"
	"fmt"
)

func main() {
	//Level one
	one.Only1()
	fmt.Println()
	one.OnlyA()
	fmt.Println()
	one.OnlyB()
	fmt.Println()
	one.OnlyF()
	fmt.Println()
	one.OnlyZ()
	fmt.Println()

	//Level 2
	//CheckNumber
	fmt.Println(two.CheckNumber("Hello"))
	fmt.Println(two.CheckNumber("Hello1"))

	fmt.Println(four.WeAreUnique("everyone", ""))

}
