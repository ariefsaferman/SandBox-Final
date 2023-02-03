package main

import (
	f "fmt"
	m "math"
)

type calculate interface {
	area() float64
	surface() float64
}

type circular struct {
	diameter float64
}

type rectanguler struct {
	length, width float64
}

func (l circular) jariJari() float64 {
	return l.diameter / 2
}

func (l circular) area() float64 {
	return m.Pi * m.Pow(l.jariJari(), 2)
}

func (l circular) surface() float64 {
	return m.Pi * l.diameter
}

func main() {
	lingkaran := circular{20}

	f.Println(lingkaran.area())
	f.Println(lingkaran.surface())
}
