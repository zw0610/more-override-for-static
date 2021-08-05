package main

import (
	"fmt"
)

type Sounder interface {
	Sound()
	Foo()
}

type Animal struct {
	Sounder
}

func (a *Animal) Foo() {
	fmt.Println("foo from animal")
}

func (a *Animal) Sound() {
	a.Foo()
	fmt.Println("Pruttel")
}

type Duck struct {
	Animal
}

func (d *Duck) Foo() {
	fmt.Println("foo from duck")
}

func main() {
	// speak to interfaces not implementations
	var animal Sounder = &Animal{}
	animal.Sound()

	var duck Sounder = &Duck{}
	duck.Sound()
}