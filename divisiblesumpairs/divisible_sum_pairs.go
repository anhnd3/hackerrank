package divisiblesumpairs

// DivisibleSumPairs ...
func DivisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var counter, i, j int32

	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			sum := ar[i] + ar[j]
			if sum == k || sum%k == 0 {
				counter++
			}
		}
	}
	return counter
}
