package staircase

import "fmt"

// Staircase Complete the staircase function below.
func Staircase(n int32) {
	var i, j int32
	for i = n; i > 0; i-- {
		for j = 0; j < n; j++ {
			if j < i-1 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}

}
