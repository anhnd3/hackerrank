package appleandorange

import "fmt"

// CountApplesAndOranges ...
func CountApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var totalApples, totalOranges int32

	for _, value := range apples {
		tmp := a + value
		if tmp >= s && tmp <= t {
			totalApples++
		}
	}

	for _, value := range oranges {
		tmp := b + value
		if tmp >= s && tmp <= t {
			totalOranges++
		}
	}
	fmt.Println(totalApples)
	fmt.Println(totalOranges)
}
