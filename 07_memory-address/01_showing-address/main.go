package main

import "fmt"

func main() {

	a := 43
	describer := "a's memory address (dec) -"

	fmt.Println("a -", a)
	fmt.Println("a's memory address (hex) -", &a)
	fmt.Printf("%s", describer)
	fmt.Printf(" %d\n", &a)

}
