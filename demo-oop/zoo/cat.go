package zoo

import "fmt"

type Cat struct {
	Animal
}

func NewCat() *Cat {
	return &Cat{
		Animal: Animal{
			species: "cat",
		},
	}
}

func (cat *Cat) Eat() {
	cat.Animal.Eat()
	cat.meow()

}

func (cat *Cat) meow() {
	fmt.Println("meow")
}
