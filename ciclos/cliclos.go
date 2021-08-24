package main

import "fmt"

func main() {
	// Ciclo for
	for i := 0; i < 10; i++ {
		fmt.Println("Hello world")
	}

	// Ciclo while [Simulacion]
	count := 0
	for count < 10 {
		fmt.Println("Hello world")
		count++
	}

	// Ciclo infinito
	count_infinity := 0
	for {
		/*
			if count_infinity == 2 {
				continue
			}
		*/

		fmt.Println(count_infinity)
		count_infinity++

		if count_infinity > 10 {
			break
		}
	}
}
