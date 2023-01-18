/*
Tout d’abord, écrivez un programme qui imprime les chiffres 1 à 100, avec les modifications suivantes :

Imprimer Fizz si le nombre est divisible par 3.
Imprimer Buzz si le nombre est divisible par 5.
Imprimer FizzBuzz si le nombre est divisible par 3 et par 5.
Imprimer le nombre si aucun des cas précédents ne correspond.


package main

import (
	"fmt"
)

func main() {
	for i := 0; i <    100; i++ {
		switch {
		case i%3 == 0:
			fmt.Print("Fizz")
		case i%5 == 0:
			fmt.Print("Buzz")
		case i%5 == 0, i%3 == 0:
			fmt.Print("FizzBuzz")
		default:
			fmt.Print(i)
		}
	}
}
*/

package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(num int) string {
	switch {
	case num%15 == 0:
		return "FizzBuzz"
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(num)
}

func main() {
	for num := 1; num <= 100; num++ {
		fmt.Print(fizzbuzz(num))
	}
}
