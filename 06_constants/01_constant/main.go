package main

import "fmt"

const p = "death & taxes"

func main() {
	const q = 42
	fmt.Printf("%T ", p)
	fmt.Println("p - ", p)
	fmt.Printf("%T ", q)
	fmt.Println("q - ", q)
}
