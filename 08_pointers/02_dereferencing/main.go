package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0xc420088008 (different each time it's run!)

	var b *int = &a
	fmt.Println(b)  // 0xc420088008
	fmt.Println(*b) // 43

}
