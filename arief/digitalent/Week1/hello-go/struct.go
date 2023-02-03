package main

import f "fmt"

type Robot struct {
	name string
	tipe string
	seri int
}

func main() {

	var robot1 = Robot{name: "Ajax", tipe: "autobots"}
	// var robot2 = Robot{name: "Rebox", tipe: "coconot"}

	var ptrRobot1 *Robot = &robot1
	// var ptrRobot2 = &robot2

	f.Println(ptrRobot1.name)
	ptrRobot1.name = "Rebox"
	f.Println(robot1.name)
	// fmt.Println(ptrRobot2)

}
