package main

import "fmt"

func main() {
	numeros := [1][6]int{{1, 2, 3, 4, 5}}
	for linha := range numeros {
		for elemento := range numeros[linha] {
			fmt.Println(elemento)
		}
	}

}
