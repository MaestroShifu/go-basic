package main

import "fmt"

/*
	Operadores logicos

	== igual a
	!= diferente de
	< menor que
	> mayor que
	>= mayor o igual que
	<= mayor o igual que
	&& AND
	|| OR
*/

func main() {
	x := 10
	y := 12
	if x > y {
		fmt.Printf("%d es mayor que %d", x, y)
	} else if x < y {
		fmt.Printf("%d es mayor que %d", y, x)
		fmt.Println("Entro al else if")
	} else {
		fmt.Println("Son iguales")
	}
}
