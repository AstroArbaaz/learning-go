// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, world")
// }

// How It Works:
// When you write a program in Go you will have a main package defined with a main func inside it. Packages are ways of grouping up related Go code together.
// The func keyword is how you define a function with a name and a body.
// With import "fmt" we are importing a package which contains the Println function that we use to print.

// seperating domain code with outside world;

package main

import "fmt"

func Hello() string {
	return "Hello, World"
}

func main() {
	fmt.Println(Hello())
}