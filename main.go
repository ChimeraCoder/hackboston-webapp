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

var alice Person = Person{"Alice", Female}
var bob = Person{"Bob", Male}

func main() {

	log.Printf("Here is Alice: %+v", alice)
	log.Printf("Here is Bob: %+v", bob)

}
