// convert hexadecimal to decimal

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func HexaDecimal(num string) int64 {
	Hexa, err := strconv.ParseInt(num, 16, 64) // cette ligne permet de convertir le nombre en int
	if err != nil {                            // cette ligne permet d'afficher un message d'erreur si l'utilisateur entre une donnee erronée
		panic(err)
	}
	return Hexa
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an Hexadecimal number: ")
	sc.Scan()
	msg := sc.Text()
	fmt.Println(HexaDecimal(msg))
}
