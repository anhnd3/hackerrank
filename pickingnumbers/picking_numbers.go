package pickingnumbers

import "fmt"

func pickingNumbers(a []int32) int32 {
	// Write your code here
	var highest, i int32
	counter := make([]int32, 100)

	for _, v := range a {
		counter[v]++
	}

	for i = 1; i < 100; i++ {
		tmpSum := counter[i] + counter[i-1]
		fmt.Println(tmpSum)
		if tmpSum > highest {
			highest = tmpSum
		}
	}

	return highest
}
