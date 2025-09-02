// 11 For
package main

import "fmt"

func latihanFor() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	fmt.Println("Perulangan selesai")

	for i := 10; i >= 1; i-- {
		fmt.Println("Perulangan ke-", i)
	}
}
