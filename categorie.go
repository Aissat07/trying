// Ecrivons un programme qui demande notre age et affiche notre categorie

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	EntreUtilisateur := bufio.NewScanner(os.Stdin)
	fmt.Print("Veuillez entrez votre age : ")
	EntreUtilisateur.Scan()
	Age, _ := strconv.Atoi(EntreUtilisateur.Text())

	switch {
	case Age < 18:
		fmt.Println("CATEGORIE MINEUR")
	case Age < 45:
		fmt.Println("CATEGORIE MAJEUR")
	default:
		fmt.Println("CATEGORIE ADULTE")
	}
}
