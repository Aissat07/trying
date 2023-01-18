// Erivons un programme qui nous permet de passer notre nom en paramètre et de l'afficher

package main

import (
	"fmt"
	"os"
)

func main() {
	MyName := os.Args[1]
	fmt.Println(MyName)
}
