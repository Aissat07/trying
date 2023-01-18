/* Ecrivons un programme qui demande à l'utilisateur d'entrer deux entiers
et échange les deux nombres avant de les afficher
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Print("Entrez le premier entier : ")
	sc.Scan()
	FirstNum, _ := strconv.Atoi(sc.Text())

	fmt.Print("Entrer le deuxième entier : ")
	sc.Scan()
	SecondNum, _ := strconv.Atoi(sc.Text())
	FirstNum, SecondNum = SecondNum, FirstNum
	fmt.Println(FirstNum)
	fmt.Println(SecondNum)
}
