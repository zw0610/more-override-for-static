package main

import (
	"fmt"
)

type Fooer interface {
	Foo()
}

type AnimalFooer struct {
	Fooer
}

func (f *AnimalFooer) Foo() {
	fmt.Println("foo from animal")
}

type DuckFooer struct {
	AnimalFooer
}

func (f *DuckFooer) Foo() {
	fmt.Println("foo from duck")
}

type Sounder interface {
	Sound()
}

type Animal struct {
	Sounder
	Fooer
}

func (a *Animal) Sound() {
	a.Foo()
	fmt.Println("Pruttel")
}

type Duck struct {
	Animal
}

func main() {
	// speak to interfaces not implementations
	var animal Sounder = &Animal{Fooer:&AnimalFooer{}}
	animal.Sound()

	var duck Sounder = &Duck{
		Animal{
			Fooer:   &DuckFooer{},
		},
	}
	duck.Sound()
}