// Ecrivons un programme qui prend l'annee actuel et note annee de naissance pour calculer notre age actuel et dans 5ans

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	AnneeActuel, _ := strconv.Atoi(os.Args[1])
	AnneeNaissance, _ := strconv.Atoi(os.Args[2])
	Age := AnneeActuel - AnneeNaissance
	DansCinqAns := Age + 5
	fmt.Println("Vous avez ", Age, "ans")
	fmt.Println("Dans ans vous aurez", DansCinqAns, "ans")
}
