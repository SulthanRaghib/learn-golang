// 08. Koversi Tipe Data
package main

import (
	"fmt"
	"strconv"
)

func convertTypes() {
	var x int32 = 10
	var y int64 = int64(x) // konversi dari int32 ke int64

	fmt.Println(y) // 10

	var z float64 = float64(y) // konversi dari int64 ke float64
	fmt.Println(z)             // 10

	var a int32 = int32(z) // konversi dari float64 ke int32
	fmt.Println(a)         // 10

	var b float64 = 3.3
	var c int32 = int32(b) // konversi dari float64 ke int32
	fmt.Println(c)         // 3

	var nilai string = "100"
	nilaiInt, _ := strconv.Atoi(nilai) // konversi dari string ke int
	fmt.Println(nilaiInt)              // 100

	nilaiString := strconv.Itoa(nilaiInt) // konversi dari int ke string
	fmt.Println(nilaiString)              // "100"
}
