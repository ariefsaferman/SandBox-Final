package main

import "fmt"

func main() {
	months := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	skipMonth := months[3:7]
	fmt.Println(skipMonth)

	newSlice := make([]string, 3, 10)

	newSlice[0] = "Kampung"
	newSlice[1] = "durian"

	fmt.Println(newSlice)

}
