package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Mehdi":
		fmt.Println("Wassup Mehdi")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}

}
