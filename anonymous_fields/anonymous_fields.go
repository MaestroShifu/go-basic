package main

import "fmt"

/*
	Es como el manejo de la herencia de programacion orientada a objetos
*/

type Human struct {
	name string
}

func (this Human) hablar() string {
	return "bla bla bla"
}

type Tutor struct {
	Human
}

func (this Tutor) hablar() string { //  Sirve para sobre escribir una clase
	text := this.Human.hablar() // Sirve para llamar al padre
	return text + " " + "Hola X2"
}

func main() {
	tutor := Tutor{Human{"daniel"}}

	fmt.Println(tutor.name)
	fmt.Println(tutor.hablar())
}
