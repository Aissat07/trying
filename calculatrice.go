//  Programmons une calculatrice

package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operation string

	fmt.Print("Entrer un operateur (+, -, *, / ) : ")
	fmt.Scan(&operation)

	fmt.Print("Entrer le prémier nombre: ")
	fmt.Scan(&num1)

	fmt.Print("Entrer le deuxième nombre: ")
	fmt.Scan(&num2)

	switch operation {
	case "+":
		fmt.Print(num1, " + ", num2, " = ", num1+num2)
	case "-":
		fmt.Print(num1, " - ", num2, " = ", num1-num2)
	case "*":
		fmt.Print(num1, " * ", num2, " = ", num1*num2)
	case "/":
		fmt.Print(num1, " / ", num2, " = ", num1/num2)
	default:
		fmt.Println("Oprearation invalide !")
	}

}
