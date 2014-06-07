package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Create a type alias
// Normally this would just be called "Sex"
// The T (for "type") is added here just for clarity
type SexT int

const (
	Unknown SexT = iota
	Male
	Female
)

type Person struct {
	Name string
	Sex  SexT
}

func (p Person) Greet() string {
	switch p.Name {
	case "Alice":
		return "Hey, Bob!"
	case "Bob":
		return "Hey, Alice!"
	}
	return "Eve is eavesdropping!"
}

var alice Person = Person{"Alice", Female}
var bob = Person{"Bob", Male}

func serveHome(w http.ResponseWriter, r *http.Request) {
	t := template.New("base")
	s1, err := t.ParseFiles("templates/base.tmpl")
	if err != nil {
		// TODO don't panic!
		panic(err)
	}

	err = s1.ExecuteTemplate(w, "base", alice)
	if err != nil {
		// TODO don't panic!
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)
	http.Handle("/", r)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Error listening, %v", err)
	}
}
