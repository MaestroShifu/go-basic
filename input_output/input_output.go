package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var edad int
	var bandera bool

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Agrega tu nombre de usuario")
	nombre, err := reader.ReadString('\n')

	if err != nil {
		println(err)
		return
	}

	fmt.Println("Agrega tu edad")
	fmt.Scanf("%d\n", &edad)

	fmt.Println("Eres activo?")
	fmt.Scanf("%t\n", &bandera)

	/*
		Verbos para fmt

		- Para imprimir una estructura de datos %v
		- Para imprimir una estructura con las propiedades %+v
		- Para imprimir una representacion de la sintaxis de Go del valor %#v
		- Para imprimir el tipo de valor %T
		- Para imprimir un formato boleano %t
		- Para imprimir un entero %d
		- Para imprimir una representacion binaria %b
		- Para imprimir la letra que corresponda al entero %c
		- Para imprimir en decodificacion hexadecimal %x
		- Para imprimir punto flotante %f
		- Para imprimir punto flotante notacion cientifica %e o %E
		- Para imprimir cadenas simples %s
		- Para incluir doble comilla es %q
		- Para imprimir la salida de un apuntador %p
		-
	*/
	fmt.Printf("Mi nombre es: %s\nMi edad es: %d\nEstoy activo: %t", nombre, edad, bandera)

}
