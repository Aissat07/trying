/* Ecrivons un programme qui indique si le mot entré par l'utilisateur commence par une
voyelle ou non
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrer un mot de votre choix :> ")
	sc.Scan()
	text := sc.Text()
	runes := []rune(text)
	Mot := runes[0]
	if Mot == 'a' || Mot == 'e' || Mot == 'i' || Mot == 'o' || Mot == 'u' || Mot == 'y' {
		fmt.Println("Votre mot commmence par une VOYELLE !!")
	} else {
		fmt.Println("Votre mot commmence par une CONSONNE !!")
	}
}
