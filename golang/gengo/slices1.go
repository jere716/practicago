package main

import "fmt"

func main() {
	// slices
	slc := []int{3, 5, 7}

	// agregar elementos a un slices existente
	slc = append(slc, 9, 11)
	fmt.Println(slc)

	// eliminando elementos de un slice
	slc = slc[1:3]
	fmt.Println(slc)

	/*
		// se puede crear un for usando la longitud(len) del slice para indicar el limite
			for i:=0; len(slc); i++ {
				fmt.Println(slc[i])
			}
	*/

	// range devuelve el valor del indice, y el valor ocupado por ese indice
	for i, va := range slc {
		fmt.Println("index:", i, "values", va)
	}

	// podemos ignorar el indice y usar solo el valor
	for _, num := range slc {
		fmt.Println("valor:", num)
	}
}
