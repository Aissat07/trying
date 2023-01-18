// Ecrivons un programme qui prend l'annee actuel et note annee de naissance pour calculer notre age actuel et dans 5ans

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	AnneeActuel := bufio.NewScanner(os.Stdin)
	fmt.Print(" Entrez l'année en cours : ")
	AnneeActuel.Scan()

	AnneeNaissance := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez votre année de naissance : ")
	AnneeNaissance.Scan()

	Actuel, _ := strconv.Atoi(AnneeActuel.Text())
	Naissance, _ := strconv.Atoi(AnneeNaissance.Text())
	Age := Actuel - Naissance

	DansCinqAns := Age + 5

	fmt.Println("Vous avez ", Age, "ans")

	fmt.Println("Dans ans vous aurez", DansCinqAns, "ans")

}
