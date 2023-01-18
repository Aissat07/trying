// ecrivons un programme qui nous affiche le nom du programme

package main

import (
	"fmt"
	"os"
)

func main() {
	Programmname := os.Args[0]
	fmt.Println(Programmname)
}
