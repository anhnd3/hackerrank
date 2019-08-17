package sockmerchant

// SockMerchant ...
func SockMerchant(n int32, ar []int32) int32 {
	maps := make(map[int32]int32)

	var index int32
	for index = 0; index < n; index++ {
		if _, isExists := maps[ar[index]]; isExists {
			maps[ar[index]]++
		} else {
			maps[ar[index]] = 1
		}
	}

	var numOfPair int32
	for _, value := range maps {
		numOfPair += value / 2
	}

	return numOfPair
}
