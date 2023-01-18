// Ecrivons un programme qui demande un mot de passe à l'utilisateur et lui permet 3 essais

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	PassWord := os.Args[1]
	sc := bufio.NewScanner(os.Stdin)
	for i := 1; i <= 3; i++ {
		fmt.Print("Veuillez saissir votre mot de passe : ")
		sc.Scan()
		RealPass := sc.Text()

		if RealPass == PassWord {
			fmt.Println("Welcome !!!")
			break
		} else if RealPass != PassWord && i < 3 {
			fmt.Println("Vous avez echouer", i, "fois ")
			fmt.Println("Réessayer !!")
		} else {
			fmt.Print("Vous avez échouer 3 fois ")
			fmt.Println("Vous avez atteint la limit d'essai revenez plus tard !!")
		}

	}
}
