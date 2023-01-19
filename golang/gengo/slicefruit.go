package main

import "fmt"

func main() {
	// creando un slice de frutas
	fruits := []string{"manzana", "pera", "guineo"}

	// el guion bajo se usa para ignorar algun valor

	for _, fruit := range fruits {
		// guardando el penultimo indice en una variable
		index := len(fruit) - 1
		// buscando las frutas que terminan con la letra "a" para imprimirla
		if fruit[index:] == "a" {
			fmt.Println(fruit)
		}
	}

}
