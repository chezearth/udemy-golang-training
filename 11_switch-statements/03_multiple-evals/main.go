package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, er, Jenny")
	case "Marcus", "Mehdi":
		fmt.Println("Both of your names start with 'M'")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian/Sushant")
	}
}
