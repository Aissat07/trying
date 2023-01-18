// Ecrivons un programme qui nous permet de calculer la somme et le produit de 2 nombres en utlisant des fonctions

package main

import (
	"fmt"
	"os"
	"strconv"
)

func Somme(a, b string) int {
	num_1, _ := strconv.Atoi(a)
	num_2, _ := strconv.Atoi(b)
	return num_1 + num_2
}

func Produit(a, b string) int {
	num_1, _ := strconv.Atoi(a)
	num_2, _ := strconv.Atoi(b)
	return num_1 * num_2
}

func main() {
	som := Somme(os.Args[1], os.Args[2])
	prod := Produit(os.Args[1], os.Args[2])

	fmt.Println(" La somme est : ", som)
	fmt.Println("Le produit est : ", prod)
}
