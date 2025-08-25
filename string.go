// 06 Manipulasi String
package main

import "fmt"

func latihanString() {
	fmt.Println("Sulthan Raghib Fillah")           // mencetak nama
	fmt.Println("Sulthan Raghib" + " " + "Fillah") // menggabungkan string dengan tanda +
	fmt.Println(len("Sulthan Raghib Fillah"))      // menghitung jumlah panjang string
	fmt.Println("Sulthan Raghib Fillah"[0])        // mengambil karakter pertama (byte)
	fmt.Println("Sulthan Raghib Fillah"[8:14])     // ambil karakter index 8-13
	fmt.Println("Sulthan Raghib Fillah"[8:])       // ambil karakter index 8 sampai akhir
	fmt.Println("Sulthan Raghib Fillah"[:7])       // ambil karakter index 0-6
}
