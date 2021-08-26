package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("public"))

	http.Handle("/public/", http.StripPrefix("/public/", fileServer))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Hay una nueva peticion")
	io.WriteString(res, "Hola mundo web")
}
