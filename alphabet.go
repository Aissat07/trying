// Ecrivons l'alphabet de A à Z et de Z à A

package main

import "fmt"

func main() {
	fmt.Println("L'alphabet de A à Z")

	for i := 'a'; i <= 'z'; i++ {
		fmt.Print(string(i))
	}
	fmt.Println()

	fmt.Println("L'alphabet de Z à A")
	for j := 'z'; j >= 'a'; j-- {
		fmt.Print(string(j))
	}
	fmt.Println()
}
