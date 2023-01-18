// Ecrivons un programme qui indique la parite du nombre

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Veuillez saissir un nombre :> ")
	sc.Scan()
	nbr, _ := strconv.Atoi(sc.Text())

	if nbr%2 == 0 {
		fmt.Println("Le nombre est Pair !!")
	} else {
		fmt.Println("Le nombre est Impair !!")
	}
}
