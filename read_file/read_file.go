package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func read_file() {
	file_data, err := ioutil.ReadFile("./hola.txt") // ./route_or_filename -> Esto es una ruta relativa depende de dodne se ejecuta el script
	if err != nil {
		fmt.Println("Hay un error")
	}
	fmt.Println(string(file_data))
}

func read_file_by_line() bool {
	file, err := os.Open("./hola.txts")
	defer func() { // Se ejecuta cuando retorna o hay un error en la ejecucion
		file.Close()
		fmt.Println("Se ejecuto el defer")
		r := recover() // Ayuda a quitar el panic y ayuda a continuar el ciclo
		fmt.Println(r)
	}()
	if err != nil {
		panic(err) // Forma de imprimir un error y saber la traza del error -  Genera un error y detiene el ciclo de la aplicacion
	}
	var i int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
		line := scanner.Text()
		fmt.Println(i)
		fmt.Println(line)
	}
	return true
}

func execute_read_file() {
	value := read_file_by_line()
	fmt.Println(value)
}

func main() {
	// read_file()
	execute_read_file()
	fmt.Println("No se va a llamar")
}
