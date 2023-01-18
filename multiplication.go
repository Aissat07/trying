// table de Multiplication

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {

		fmt.Println("Entrez un nombre en argument :> ")
		return
	}
	nbr, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("nombre invalide")
		return
	}
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", nbr, i, nbr*i)
	}

}
