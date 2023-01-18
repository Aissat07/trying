package main

import (
	"bufio"
	"fmt"
)

func main() {

	scanner := bufio.NewScanner(Stdin) // création du scanne capturant l'entrée de l'utilisateur
	fmt.Print("Entrez quelque chose: ")
	scanner.Scan()                      // lancement du scan
	entreeUtilisateur := scanner.Text() // stockage de données entré par un utilisateur
	fmt.Println(entreeUtilisateur)
}
