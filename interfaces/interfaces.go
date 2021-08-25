package main

import "fmt"

/*
	- Solo sirve para declarar que una estructura cumple algunos parametros
	- Un tipo de dato
*/

type User interface {
	Permissions() int // 1 - 5
	Name() string
}

type Admin struct {
	name string
}

func (this Admin) Permissions() int {
	return 5
}

func (this Admin) Name() string {
	return this.name
}

type Editor struct {
	name string
}

func (this Editor) Permissions() int {
	return 2
}

func (this Editor) Name() string {
	return this.name
}

func Auth(user User) string {
	if user.Permissions() > 4 {
		return user.Name() + " tiene permisos de administrador"
	}
	return "No tiene permisos"
}

func main() {
	/* 	admin := Admin{"Daniel"}
	   	editor := Editor{"Pepito"} */

	usuarios := []User{Admin{"Admin1"}, Editor{"Editor1"}}

	for _, usuario := range usuarios {
		fmt.Println(Auth(usuario))
	}

	/* 	fmt.Println(Auth(admin))
	   	fmt.Println(Auth(editor)) */
}
