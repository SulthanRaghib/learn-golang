// 16. Slice
package main

import "fmt"

func latihanSlice() {
	names := [5]string{
		"Sulthan",
		"Raghib",
		"Fillah",
		"Fauzan",
		"Rizki",
	}

	fmt.Println(names) // menampilkan semua isi array

	slice1 := names[0:3] // mengambil index 0 sampai 2 (3-1)
	fmt.Println(slice1)

	slice2 := names[:3] // mengambil index 0 sampai 2 (3-1)
	fmt.Println(slice2)

	slice3 := names[3:] // mengambil index 3 sampai akhir
	fmt.Println(slice3)

	names[3] = "Fauzan Fadhil" // mengubah isi array index ke 3

	fmt.Println(names)  // menampilkan semua isi array
	fmt.Println(slice1) // menampilkan slice1
	fmt.Println(slice2) // menampilkan slice2
	fmt.Println(slice3) // menampilkan slice3

	slice1[2] = "Ramadhan" // mengubah isi array index ke 4

	fmt.Println(names)  // menampilkan semua isi array
	fmt.Println(slice1) // menampilkan slice1
	fmt.Println(slice2) // menampilkan slice2
	fmt.Println(slice3) // menampilkan slice3

	// Hati-hati ketika membuat slice dari array
	// Jika array diubah, maka slice juga akan berubah
	// Jika slice diubah, maka array juga akan berubah

}
