// une autre manière d'ecrire une calculatrice

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Choix un opérateur: 1.Add, 2.Sous, 3.Mul, 4.Div : ")
	sc.Scan()
	msg1 := sc.Text()
	choix, _ := strconv.ParseFloat(msg1, 64)

	fmt.Print("Entrer le Premier nombre : ")
	sc.Scan()
	msg2 := sc.Text()
	a, _ := strconv.ParseFloat(msg2, 64)

	fmt.Print("Entrer le Deuxième nombre : ")
	sc.Scan()
	msg3 := sc.Text()
	b, _ := strconv.ParseFloat(msg3, 64)

	if choix == 1 {
		cal := a + b
		fmt.Println(a, " + ", b, " = ", cal)
	} else if choix == 2 {
		cal := a - b
		fmt.Println(a, " - ", b, " = ", cal)
	} else if choix == 3 {
		cal := a * b
		fmt.Println(a, " * ", b, " = ", cal)
	} else if choix == 4 {
		if choix == 4 && b == 0 {
			fmt.Println("Operation invalide !!")
		} else {
			cal := a / b
			fmt.Println(a, " / ", b, " = ", cal)
		}
	}
}
