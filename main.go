package main

import "log"

type Person struct {
	Name string
	Sex  string
}

func main() {

	alice := Person{"Alice", "female"}
	bob := Person{"Bob", "male"}
	log.Printf("Here is Alice: %+v", alice)
	log.Printf("Here is Bob: %+v", bob)

}
