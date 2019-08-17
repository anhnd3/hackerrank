package hurdlerace

// HurdleRace ...
func HurdleRace(k int32, height []int32) int32 {
	var maxHeight int32

	for _, v := range height {
		if v > maxHeight {
			maxHeight = v
		}
	}

	if maxHeight <= k {
		return 0
	}
	return maxHeight - k
}
