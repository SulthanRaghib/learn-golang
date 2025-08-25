// 07. Variabel dan Constant
package main

import "fmt"

func printVariables() {
	var hello string = "Hello World"
	fmt.Println(hello)
	fmt.Println(hello)
	fmt.Println(hello)

	hello = "Hello Go"
	fmt.Println(hello)
	fmt.Println(hello)
	fmt.Println(hello)

	var nama string   // mendeklarasikan variabel nama
	fmt.Println(nama) // default value dari string adalah kosong
	nama = "Raghib"   // menginisialisasi variabel nama
	fmt.Println(nama) // mencetak nilai variabel nama

	namaLengkap := "Sulthan Raghib Fillah" // mendeklarasikan dan menginisialisasi variabel namaLengkap
	fmt.Println(namaLengkap)

	nilai := 10 // mendeklarasikan dan menginisialisasi variabel nilai
	fmt.Println(nilai)

	// Constanta Variable
	const kampus string = "STT Terpadu Nurul Fikri"
	fmt.Println(kampus)
}
