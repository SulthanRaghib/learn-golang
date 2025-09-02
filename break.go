// 14. Break dan Continue
package main

import "fmt"

func latihanBreak() {
	for i := 1; i <= 100; i++ {
		// if i%2 == 0 {
		// 	continue // Skip the current iteration
		// }
		fmt.Println(i)
		if i == 50 {
			break // Stop the loop when i equals 50
		}
	}
}
