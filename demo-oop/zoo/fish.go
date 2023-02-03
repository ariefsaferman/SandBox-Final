package zoo

type Fish struct {
	animal Animal
}

func NewFish() *Fish {
	return &Fish{
		animal: Animal{species: "fish"},
	}
}

func (fish *Fish) Sleep() {
	fish.animal.Sleep()
}

func (fish *Fish) Eat() {
	fish.animal.Eat()
}
