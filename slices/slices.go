package main

import "fmt"

func main() {
	/*
		Es una especie de arrecglo que no necesita declarar el tamaño
	*/
	matriz := []int{}
	if matriz == nil {
		fmt.Println("Esta vacio!!")
	} else {
		// Esto posee un slice
		// Puntero al arrecglo
		// Longitud del arrecglo al que apunta
		// Capacidad
		fmt.Println(matriz)
		fmt.Println(len(matriz))
	}

	/*
		Pasar un array a slice
	*/
	arrecglo := [3]int{1, 2, 3}
	// arrecglo[startPos:endPos] -> Se le llama slicing
	slice := arrecglo[:2]
	fmt.Println(slice)
}
