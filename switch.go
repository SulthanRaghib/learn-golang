// 13. Switch
package main

import (
	"fmt"
	"runtime"
)

func latihanSwitch() {
	for i := 1; i <= 10; i++ {
		// if i == 1 {
		// 	fmt.Println("Satu")
		// } else if i == 2 {
		// 	fmt.Println("Dua")
		// } else if i == 3 {
		// 	fmt.Println("Tiga")
		// } else if i == 4 {
		// 	fmt.Println("Empat")
		// } else if i == 5 {
		// 	fmt.Println("Lima")
		// } else {
		// 	fmt.Println("Gak Tau")
		// }

		// Switch biasanya digunakan saat mengecek 1 variable saja
		switch i {
		case 1:
			fmt.Println("Satu")
		case 2:
			fmt.Println("Dua")
		case 3:
			fmt.Println("Tiga")
		case 4:
			fmt.Println("Empat")
		case 5:
			fmt.Println("Lima")
		default:
			fmt.Println("Gak Tau")
		}
	}

	sistemOperasi := runtime.GOOS // Mendapatkan nama sistem operasi yang digunakan
	switch sistemOperasi {
	case "windows":
		fmt.Println("Kamu menggunakan sistem operasi Windows")
	case "linux":
		fmt.Println("Kamu menggunakan sistem operasi Linux")
	case "darwin":
		fmt.Println("Kamu menggunakan sistem operasi MacOS")
	default:
		fmt.Println("Sistem operasi kamu tidak diketahui")
	}
}
