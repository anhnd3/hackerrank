package birthdaychocolates

// Birthday ...
func Birthday(s []int32, d int32, m int32) int32 {
	length := int32(len(s))
	var index, counter int32

	for index < length-m+1 {
		tmpIdx1 := index
		var index2, sum int32
		for index2 < m {
			sum += s[tmpIdx1]
			tmpIdx1++
			index2++
		}
		if sum == d {
			counter++
		}

		index++
	}
	return counter
}
