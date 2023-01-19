package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// creating a struct
type Users struct {
	Name   string
	Skills string
	Age    int
}

// creating a function
func index(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	//fmt.Fprintln(rw, "Hola Jeremy now")

	// usando la estructura
	user := Users{"Eli Marciano", "Programer", 36}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, user)
	}
}

func main() {
	// creating a route for handler
	http.HandleFunc("/", index)

	// creating a web server
	fmt.Println("Run Server: http://localhost:3000")
	fmt.Println("El servidor esta corriendo en el puerto 3000")
	http.ListenAndServe("localhost:3000", nil)

}
