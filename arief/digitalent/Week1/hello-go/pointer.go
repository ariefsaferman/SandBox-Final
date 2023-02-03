package main

import "fmt"

func main() {
	a := 20
	b := 50
	ptrA := &a
	ptrB := &b

	fmt.Println("==== before swap ====")
	fmt.Println("value A =", *ptrA)
	fmt.Println("value B =", *ptrB)

	swap(ptrA, ptrB)

	fmt.Println("==== after swap ====")
	fmt.Println("value A =", *ptrA)
	fmt.Println("value B =", *ptrB)

}

func swap(valueA *int, valueB *int) {
	temp := *valueA
	*valueA = *valueB
	*valueB = temp
}
