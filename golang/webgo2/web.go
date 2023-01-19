package main

import (
	"html/template"
	"log"
	"net/http"
)

// la variable de la ruta de los templates
var plantillas = template.Must(template.ParseGlob("templates/*"))

// creando la funcion a manejar
func index(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "index", nil)
	//fmt.Fprintf(w, "Hello Jeremy por formato")
}

func main() {
	// manejador de funciones y la ruta
	http.HandleFunc("/", index)

	// crear el servidor
	log.Println("Server Running at port:8000")
	http.ListenAndServe("localhost:8000", nil)
}
