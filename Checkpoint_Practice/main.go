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
	//printif
	fmt.Print(two.PrintIf("abcdefz"))
	fmt.Print(two.PrintIf("abc"))
	fmt.Print(two.PrintIf(""))
	fmt.Print(two.PrintIf("14"))
	//PrintIfNot
	fmt.Print(two.PrintIfNot("abcdefz"))
	fmt.Print(two.PrintIfNot("abc"))
	fmt.Print(two.PrintIfNot(""))
	fmt.Print(two.PrintIfNot("14"))
	//RectPerimeter
	fmt.Println(RectPerimeter(10, 2))
	fmt.Println(RectPerimeter(434343, 898989))
	fmt.Println(RectPerimeter(10, -2))


	fmt.Println(four.WeAreUnique("everyone", ""))

}
