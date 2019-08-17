package hourglasssum

// HourglassSum ...
func HourglassSum(arr [][]int32) int32 {
	var maxHourglass int32
	var arrHourglass []int32

	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			tmpSum := arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] + arr[i][j] + arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1]
			arrHourglass = append(arrHourglass, tmpSum)
		}
	}

	maxHourglass = arrHourglass[0]
	for i := 1; i < len(arrHourglass); i++ {
		if arrHourglass[i] > maxHourglass {
			maxHourglass = arrHourglass[i]
		}
	}

	return maxHourglass
}
