package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	c := &a

	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

}
