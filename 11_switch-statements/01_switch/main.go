package main

import "fmt"

func main() {
	switch "Mehdi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Mehdi":
		fmt.Println("Wassup Mehdi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
}
