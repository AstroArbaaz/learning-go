package main

import "fmt"

func main() {
	// strings variables
	var nameOne string = "Arbaaz" // variable with explicit value type 
	var nameTwo = "Michelle" // variable with no explicit valaue type
	var nameThree string // empty vairable with no value type

	nameThree = "Sudindra"

	// short-hand way of declaring variable
	nameFour := "Sonali" // Warning : you cant you this shorthand way of declaring a variable outside a function

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// integers variables
	var ageOne int = 10
	var ageTwo = 20
	ageThree := 30

	// caution these are whole numbers only. no decimals are allowed.

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	var numOne int8 = 25 // int8 is for assigning 8 bits of integer, Range -128 through 127
	var numTwo uint8 = 255 // uint is for not using a negative number

	// floats integer variables
	var scoreOne float32 = 25.88 // two type of floats 32 and 64 bit, float64 has higher preccession

	fmt.Println(numOne, numTwo, scoreOne)

}