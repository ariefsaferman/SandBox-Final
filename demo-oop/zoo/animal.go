package zoo

import "fmt"

type Animal struct {
	species string
	hungry  int
	energy  int
}

func NewAnimal(species string) *Animal {
	return &Animal{
		species: species,
	}
}

func (animal *Animal) Eat() {
	animal.hungry--
	animal.energy++
	fmt.Println(animal.species, "eating...")
}

func (animal *Animal) Sleep() {
	animal.hungry++
	animal.energy++
	fmt.Println(animal.species, "sleep...")
}
