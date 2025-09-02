// 15. Array
package main

import "fmt"

func latihanArray() {
	var names [3]string

	names[0] = "Sulthan"
	names[1] = "Raghib"
	names[2] = "Fillah"

	// for i := 0; i < 3; i++ { // cara pertama
	// 	fmt.Println(names[i])
	// }

	// for i := 0; i < len(names); i++ { // menggunakan len untuk mendapatkan panjang array
	// 	fmt.Println(names[i])
	// }

	for index, value := range names { // foreach dalam golang
		fmt.Println("index", index, "=", value)
	}

	for _, value := range names { // foreach dalam golang jika ingin hanya nilai dan menghilangkan index
		fmt.Println(value)
	}

	// fmt.Println(names)

	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])

	name := [5]string{ // array
		"Sulthan",
		"Raghib",
		"Fillah",
		"Fauzan",
		"Rizki",
	}
	fmt.Println(name)

}
