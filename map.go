// 18. Map
package main

import "fmt"

func latihanMap() {
	mahasiswa := make(map[string]string)

	mahasiswa["1001"] = "Sulthan"
	mahasiswa["1002"] = "Raghib"
	mahasiswa["1003"] = "Fillah"

	fmt.Println(mahasiswa)

	fmt.Println(mahasiswa["1001"]) // menampilkan data dengan key 1001
	fmt.Println(mahasiswa["1002"]) // menampilkan data dengan key 1002
	fmt.Println(mahasiswa["1003"]) // menampilkan data dengan key 1003

	for nim, name := range mahasiswa { // foreach dalam map
		fmt.Println("NIM", nim, "Bernama", name)
	}

}
