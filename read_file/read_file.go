package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file_data, err := ioutil.ReadFile("./hola.txt") // ./route_or_filename -> Esto es una ruta relativa depende de dodne se ejecuta el script
	if err != nil {
		fmt.Println("Hay un error")
	}

	fmt.Println(string(file_data))
}
