package main

import "fmt"

func main() {
	var Buah = []string{"apel", "jeruk", "anggur", "melon"}

	fmt.Println("Jumlah Element	: ", len(Buah))

	//fmt.Println("Isi Element : ", Buah)

	var aBuah = append(Buah, "pepaya")

	fmt.Println("Isi Element : ", aBuah)

	var i = 0
	for {
		fmt.Println("Element ke - :", i, " ", aBuah[i])
		i++

		if i == 5 {
			break
		}
	}
}
