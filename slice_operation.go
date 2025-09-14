// 17. Slice Manipulation
package main

import "fmt"

func latihanSliceManupulation() {
	// slice1 := []string{
	// 	"Sulthan",
	// 	"Raghib",
	// 	"Fillah",
	// }

	// fmt.Println(slice1)
	slice1 := make([]string, 3) // membuat slice dengan panjang 3 index
	slice1[0] = "Sulthan"
	slice1[1] = "Raghib"
	slice1[2] = "Fillah"

	slice2 := append(slice1, "Fauzan", "Abdurrahman") // menambahkan data ke slice1
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1[0] = "S" // mengubah data slice1 tanpa mempengaruhi slice2

	fmt.Println(slice1)
	fmt.Println(slice2) // slice2 tetap tidak berubah

	slice3 := make([]string, 1) // membuat slice dengan panjang 1 index

	copy(slice3, slice1) // menyalin data dari slice1 ke slice3

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
