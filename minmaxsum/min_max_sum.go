package minmaxsum

import (
	"fmt"
	"math"
)

// MiniMaxSum Complete the miniMaxSum function below.
func MiniMaxSum(arr []int32) {
	var length = int32(len(arr))
	var i, j int32
	max := int64(math.MinInt64)
	min := int64(math.MaxInt64)

	for i = 0; i < length; i++ {
		var sum int64
		for j = 0; j < length; j++ {
			if j == i {
				continue
			}
			sum += int64(arr[j])
		}
		if sum > max {
			max = sum
		}
		if sum < min {
			min = sum
		}
	}
	fmt.Println(min, max)
}
