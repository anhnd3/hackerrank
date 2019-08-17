package diagonaldifference

// DiagonalDifference ...
func DiagonalDifference(arr [][]int32) int32 {
	var diagonal1, diagonal2, i, j int32
	length := int32(len(arr))
	for i = 0; i < length; i++ {
		for j = 0; j < length; j++ {
			if i == j {
				diagonal1 += arr[i][j]
			}
			if i+j == length-1 {
				diagonal2 += arr[i][j]
			}
		}
	}

	if diagonal1 > diagonal2 {
		return diagonal1 - diagonal2
	}
	return diagonal2 - diagonal1
}
