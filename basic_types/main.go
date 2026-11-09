package main

import (
	"fmt"
	"basic_types/constants"
)

// var (
// 	a = 5
// 	b = 3
// 	c int
// ) // групповое или блочное объявление переменных

func main() {
	// var age int = 30 // полная форма
	// var age = 30 // тип определяется из значения
	// age := 30 // краткая форма
	// var x, y int = 1, 2
	// x, y := 1, 2 // краткая форма

	age := 35
	fmt.Println(age)
	age = 30
	fmt.Println(age)
	constants.PrintConst()
}