package main

import "fmt"

func main() {
	/*
		Los arrecglos no se puede redimencionar en tiempo de ejecucion
	*/

	// var arrecglo [10]int
	arrecglo := [3]int{1, 2, 3}

	arrecglo[2] = 20

	for i := 0; i < len(arrecglo); i++ {
		fmt.Println(arrecglo[i])
	}

	matriz := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matriz)
}
