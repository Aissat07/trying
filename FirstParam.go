// Ecrivons un programme qui affiche le premier argument

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println()
	} else {
		fmt.Println("Le premier argument est : ", os.Args[1])
	}
}
