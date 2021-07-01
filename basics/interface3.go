package main

import (
	"fmt"
)

type IAnimal interface {
	Speak()
}

type AnimalType int

const (
	Duck AnimalType = iota
	Dog
	Tiger
)

type DuckClass struct{}

type DogClass struct{}

type TigerClass struct{}

func NewDuck() *DuckClass {
	return new(DuckClass)
}

func (d *DuckClass) Speak() {
	fmt.Println("Pack pack")
}

func NewDog() *DogClass {
	return new(DogClass)
}

func (d *DogClass) Speak() {
	fmt.Println("Wow wow")
}

func NewTiger() *TigerClass {
	return new(TigerClass)
}

func (t *TigerClass) Speak() {
	fmt.Println("Halum halum")
}

// Builder Mode
// return IAnimal type (any type that implements IAnimal's methods)
func New(t AnimalType) IAnimal {
	switch t {
	case Duck:
		return NewDuck()
	case Dog:
		return NewDog()
	case Tiger:
		return NewTiger()
	default:
		panic("Unknown animal type")
	}
}

func main() {
	animals := make([]IAnimal, 0)

	duck := New(Duck)
	dog := New(Dog)
	tiger := New(Tiger)

	animals = append(animals, duck, dog, tiger)

	for _, a := range animals {
		a.Speak()
	}
}
