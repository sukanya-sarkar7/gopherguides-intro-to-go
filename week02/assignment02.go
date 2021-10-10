package main

import (
	"fmt"	
)
func main() {
	exp := map[int]string{
		1: "John",
		2: "Paul",
		3: "George",
		4: "Ringo",
	}
	fmt.Println("expected map :", exp)

	act := map[int]string{}
	for i := range exp{
		fmt.Println(i)
		act[i] = exp[i]
	}
	fmt.Println("actual map :", act)
}