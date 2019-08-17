package comparethetriplets

// CompareTriplets Complete the compareTriplets function below.
func CompareTriplets(a []int32, b []int32) []int32 {
	length := int32(len(a))
	var i int32
	result := []int32{0, 0}
	for i = 0; i < length; i++ {
		if a[i] > b[i] {
			result[0]++
		} else if a[i] < b[i] {
			result[1]++
		}
	}

	return result
}
