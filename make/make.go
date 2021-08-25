package main

import "fmt"

func main() {
	// make(type, tamaño, capacidad)
	slice := make([]int, 0, 5)
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)

	fmt.Println(slice)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	copia := make([]int, len(slice))
	/*
		La funcion copy copia el minimo de elementos de una arrecglo
		copy(destino, fuente)
	*/
	copy(copia, slice)
	fmt.Println(copia)
}
