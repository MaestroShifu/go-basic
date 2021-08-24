package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 22
	// Entero a string
	edad_str := strconv.Itoa(edad)
	// String a entero
	edad_int, _ := strconv.Atoi(edad_str)
	fmt.Println("Mi edad es " + edad_str)
	fmt.Println(edad_int + 10)
}
