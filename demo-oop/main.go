package main

import (
	"demo-oop/contest"
	"demo-oop/zoo"
)

func main() {

	a2 := zoo.NewAnimal("frog")
	a2.Eat()
	// a2.Sleep()

	cat := zoo.NewCat()
	// cat.Sleep()
	// cat.Eat()

	fish := zoo.NewFish()
	// fish.Sleep()
	// fish.Eat()

	contestants := []contest.Eater{
		a2,
		cat,
		fish,
	}

	contest.RunEatContest(contestants)

}
