// Ecrivons un programme qui nous permet de calculer la somme et le produit de 2 nombres

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	num_1, _ := strconv.Atoi(os.Args[1])
	num_2, _ := strconv.Atoi(os.Args[2])

	som := num_1 + num_2
	product := num_1 * num_2

	fmt.Println("La somme est : ", som)
	fmt.Println("Le produit est :", product)
}
