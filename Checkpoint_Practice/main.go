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
	//CountAlpha
	fmt.Println(two.CountAlpha("Hello world"))
	fmt.Println(two.CountAlpha("H e l l o"))
	fmt.Println(two.CountAlpha("H1e2l3l4o"))
	//CountCharacter
	fmt.Println(two.CountChar("Hello World", 'l'))
	fmt.Println(two.CountChar("5  balloons", 5))
	fmt.Println(two.CountChar("   ", ' '))
	fmt.Println(two.CountChar("The 7 deadly sins", '7'))

	fmt.Println(four.WeAreUnique("everyone", ""))

}
