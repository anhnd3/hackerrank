package newyearchaos

import "fmt"

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	var totalBribes, i int32

	expectedFirst := int32(1)
	expectedSecond := int32(2)
	expectedThird := int32(3)

	for i = 0; i < int32(len(q)); i++ {
		if q[i] == expectedFirst {
			expectedFirst = expectedSecond
			expectedSecond = expectedThird
			expectedThird++
		} else if q[i] == expectedSecond {
			totalBribes++
			expectedSecond = expectedThird
			expectedThird++
		} else if q[i] == expectedThird {
			totalBribes += 2
			expectedThird++
		} else {
			fmt.Println("Too chaotic")
			return
		}
	}

	fmt.Println(totalBribes)
}
