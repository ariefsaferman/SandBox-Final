package contest

type Eater interface {
	Eat()
}

func RunEatContest(contestants []Eater) {
	for _, contestant := range contestants {
		contestant.Eat()
	}

}
