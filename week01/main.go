package main

import "fmt"

func main() {
	//New line after it
	fmt.Println("Hello world!")
	//fmt.Printf("Hello world")

	// a constant cannot be declared using ":="
	//const s = "Go"
	//const i = 42

	//multiple typed constant are declared inside below block.
	//const(
	//s string = "Go"
	//i int = 42
	//b bool = true
	//)

	s := "Go"
	i := 42
	b := true

	fmt.Printf("Printing, %s!\n", s)
	fmt.Printf("Printing, %d!\n", i)
	fmt.Printf("Printing, %t!\n", b)

	//m, n, p := 10, 20, 30

	//fmt.Println(
	//"the caculation is :", m, "+", n, "=", p,
	//)
	//fmt.Printf(
	//"the calculation is : %d + %d = %d\n", m, n, p,
	//)
}
