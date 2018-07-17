package main

import "fmt"

func main() {
	greet("Jane")
	greet("john")
}

func greet(name string) {
	fmt.Println(name)
}
