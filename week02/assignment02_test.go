package main

import "fmt"

func main() {
	b := make([]string, 2, 4)
	b = append(b, "sat", "Mar", "Jup")

	fmt.Println(b)
	fmt.Println("lens:", len(b))
	fmt.Println("cap:", cap(b))
}	