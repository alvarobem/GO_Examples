package main

import "fmt"

func main() {
	// need run both files used -> go run *.go
	builder := NewPersonBuilder()
	person := builder.Name("Alvaro").SurName("Berrocal").Age(25).Build()
	fmt.Println(person.walk())
}
