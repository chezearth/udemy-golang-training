package main

import "fmt"

func main() {
	for i := 127000; i < 130000; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}
