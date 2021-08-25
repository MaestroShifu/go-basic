package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go my_name_slow("Daniel")
	fmt.Println("Que aburrido!!!")
	var wait string
	fmt.Scanln(&wait)
}

func my_name_slow(name string) {
	letters := strings.Split(name, "")
	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}
