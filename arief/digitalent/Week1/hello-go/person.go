package hellogo

type Person struct {
	name string
	age  int
}

func NewPerson(person *Person) {
	person = &Person{person.name, person.age}
}
