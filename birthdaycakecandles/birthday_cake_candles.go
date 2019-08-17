package birthdaycakecandles

// BirthdayCakeCandles Complete the birthdayCakeCandles function below.
func BirthdayCakeCandles(ar []int32) int32 {
	var i, max int32
	length := int32(len(ar))
	var total int32
	for i = 0; i < length; i++ {
		if ar[i] > max {
			max = ar[i]
		}
	}

	for i = 0; i < length; i++ {
		if ar[i] == max {
			total++
		}
	}
	return total
}
