package main

import "fmt"

func isPalindromo(text string) {
	var textReverse string
	for i := len(text) -1 ; i >= 0 ; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	slice := []string{"hola", "que", "haces"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	for _, valor := range slice {
		fmt.Println(valor)
	}

	for i := range slice {
		fmt.Println(i)
	}

	isPalindromo("casas")
}