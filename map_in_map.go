// 19. Map in Map
package main

import "fmt"

func latihanMapInMap() {
	// mahasiswa := map[string]string{
	// 	"1001": "Sulthan",
	// 	"1002": "Raghib",
	// 	"1003": "Fillah",
	// }

	// fmt.Println(mahasiswa)

	mahasiswa := map[string]map[string]string{ // map di dalam map
		"1001": {
			"nama":    "Sulthan",
			"alamat":  "Jl. Merdeka No. 1",
			"jurusan": "Informatika",
		},
		"1002": {
			"nama":    "Raghib",
			"alamat":  "Jl. Merdeka No. 2",
			"jurusan": "Sistem Informasi",
		},
		"1003": {
			"nama":    "Fillah",
			"alamat":  "Jl. Merdeka No. 3",
			"jurusan": "Teknik Komputer",
		},
	}

	// fmt.Println(mahasiswa["1001"]["nama"]) // menampilkan nama dengan key

	delete(mahasiswa, "1003") // menghapus data dengan key 1003
	fmt.Println(mahasiswa)

	// for nim, mhs := range mahasiswa { // foreach dalam map
	// 	println("NIM:", nim)
	// 	for key, value := range mhs {
	// 		println(key+":", value)
	// 	}
	// 	println()
	// }

}
