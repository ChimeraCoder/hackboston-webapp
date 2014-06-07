package main

import (
	"fmt"
	"log"
	"net/http"
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

func (p *Person) Greet() {
	switch p.Name {
	case "Alice":
		log.Println("Hey, Bob!")
	case "Bob":
		log.Println("Hey, Alice!")
	}
}

var alice Person = Person{"Alice", Female}
var bob = Person{"Bob", Male}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/", serveHome)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Error listening, %v", err)
	}
}
