package main

import "fmt"

/*
	Nos permite comunicar go rutines
*/

func main() {
	channel := make(chan string)

	go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)

	msg := <-channel //  Bloquea la rutina principal hasta que le llega informacion del canal
	fmt.Println("Imprimo lo que recibo del canal " + msg)
}
