package main

import "log"

type Person struct {
	Name string
	Sex  string
}

var alice Person = Person{"Alice", "female"}
var bob = Person{"Bob", "male"}

func main() {

	log.Printf("Here is Alice: %+v", alice)
	log.Printf("Here is Bob: %+v", bob)

}
