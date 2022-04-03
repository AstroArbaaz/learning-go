package main

import "fmt"

func main() {
	fmt.Println("printing and formating strings.")

	// using "Print" to print string on console instead of "Println"
	fmt.Print("this is from print method \n") 
	fmt.Print("this is from print method 2 \n")
	// this print method can be used for gaining information form datasets. 
	// And this doesn't print the two or more different prints in new line b default just like "println", you have to use "\n" at the end of the string

	// print different string on new line
	fmt.Println("line one \n") // "\n" means continue next string on new line
	fmt.Println("line two")

	// print and formating
	age := 23
	name := "Arbaaz"

	fmt.Println("My name is:", name,", and my age is:", age)

	// formated strings, using "Printf", %_ format specifiers
	fmt.Printf("My Name is %v, and I am %v years old. \n", name, age) // this will inject variable values into the string and use "\n" at the end of the string
	fmt.Printf("My Name is %q, and I am %q years old. \n", name, age) // this will inject variable values into the sting and put quotation marks arround it and also it doesnt work on int values.
	fmt.Printf("Age is a type of %T \n", age ) // this will give us the type of variable.
	fmt.Printf("You scored %f points \n", 22.56) // this will give us the output in float variable in proper float.
	fmt.Printf("You scored %0.2f points \n", 22.56) // this will give us the output in float variable in proper float but with in limited decimal points.

	// Sprintf ( save formatted Strings)
	var str = fmt.Sprintf("My Name is %v, and I am %v years old. \n", name, age)
	fmt.Println("the saved formatted string is:", str)
}