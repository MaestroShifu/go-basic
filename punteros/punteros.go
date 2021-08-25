package main

import "fmt"

func main() {
	/*
		- Es una direccion de memoria
		- En lugar del valor, tenemos la direccion e la que esta el valor
		- Sirve para acceder ymodificar un espacio en memoria

		Se declara de la siguiente forma
		*T => *int, *string, *Struct
		valor zero => nil
	*/

	var x, y *int
	entero := 5

	x = &entero // Acceder al valor de la direccion
	y = &entero

	*x = 6 // Nos da el valor que posee en ese espacio en la memoria

	fmt.Println("Direcciones de memoria")
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("Valor de la memoria")
	fmt.Println(*x)
	fmt.Println(*y)
}
