package plusminus

import "fmt"

// PlusMinus Complete the plusMinus function below.
func PlusMinus(arr []int32) {
	var positive, negative, zeros float32
	length := float32(len(arr))

	for _, v := range arr {
		if v > 0 {
			positive++
		} else if v < 0 {
			negative++
		} else {
			zeros++
		}
	}

	fmt.Println(positive / length)
	fmt.Println(negative / length)
	fmt.Println(zeros / length)
}
