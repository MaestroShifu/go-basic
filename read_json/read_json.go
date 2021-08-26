package main

import (
	"encoding/json"
	"net/http"
)

/*
	Los atributos con la primera letra en mayuscula significa que es publico
*/

type Curso struct {
	Title        string `json:"title"` // Definimos como se va a mostrar en el json
	NumeroVideos int    `json:"numeroVideos"`
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		cursos := Cursos{
			Curso{"Curso Go", 3},
			Curso{"Curso node", 3},
			Curso{"Curso python", 3},
		}

		json.NewEncoder(rw).Encode(cursos)
	})

	http.ListenAndServe(":8000", nil)
}
