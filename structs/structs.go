package main

import "fmt"

/*
	type -> va a definir un nuevo tipo
	struct -> vamos a definir una estructura

	Las estructuras son mutables

	Nota: En los metodos se envian son copias del objeto, toca recibir un puntero para mutar el objeto

	"Usar punteros es mas economico"
*/

type User struct {
	edad             int
	nombre, apellido string
}

func (this User) nombre_completo() string { // Este es un metodo de una estructura
	return this.nombre + " " + this.apellido
}

func (this *User) set_name(name string) {
	this.nombre = name
}

func main() {
	person := User{edad: 24, nombre: "daniel", apellido: "santos"} // Otra dorma de declarar
	personTwo := User{24, "", ""}                                  // otra forma de declarar

	personThree := new(User) // retorna un puntero

	personThree.nombre = "daniel"
	personThree.apellido = "santos"

	personThree.set_name("marcos")

	fmt.Println(person)
	fmt.Println(personTwo)
	fmt.Println(personThree.nombre_completo())
}
