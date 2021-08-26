package main

import (
	"fmt"

	"./dummy"
)

// Para poder correr GO111MODULE=off go run [file].go

func main() {
	fmt.Println(dummy.HelloWorld())
}
