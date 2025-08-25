// MANIPULASI STRING
package main

import "fmt"

func main() {
	fmt.Println("Sulthan Raghib Fillah")           // mencetak nama
	fmt.Println("Sulthan Raghib" + " " + "Fillah") // menggabungkan string dengan tanda +
	fmt.Println(len("Sulthan Raghib Fillah"))      // menghitung jumlah panjang string
	fmt.Println("Sulthan Raghib Fillah"[0])        // mengambil karakter pertama dari string (result:83) golang mengambil byte dari huruf tersebut
	fmt.Println("Sulthan Raghib Fillah"[8:14])     // mengambil karakter dari index 8 sampai 14
	fmt.Println("Sulthan Raghib Fillah"[8:])       // mengambil karakter dari index 8 sampai akhir
	fmt.Println("Sulthan Raghib Fillah"[8:])       // mengambil karakter dari index 8 sampai akhir
	fmt.Println("Sulthan Raghib Fillah"[:7])       // mengambil karakter dari index 0 sampai 7

}
