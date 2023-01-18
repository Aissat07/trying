// Ecrivons un programme qui indique que nous sommes le weekend ou non

package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Il est Samedi")
	case time.Sunday:
		fmt.Println("Il est Dimanche")
	default:
		fmt.Println("Le Weekend est fini")
	}
}
