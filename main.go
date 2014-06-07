package main

import "log"

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

func (p Person) Greet() {
	switch p.Name {
	case "Alice":
		log.Println("Hey, Bob!")
	case "Bob":
		log.Println("Hey, Alice!")
	}
}

var alice Person = Person{"Alice", Female}
var bob = Person{"Bob", Male}

func main() {
	alice.Greet()
}
