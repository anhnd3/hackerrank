package countingvalleys

// CountingValleys ...
func CountingValleys(n int32, s string) int32 {
	var sum, count, i int32
	for i = 0; i < n; i++ {
		if s[i] == 'U' {
			sum++
			if sum == 0 {
				count++
			}
		} else {
			sum--
		}
	}
	return count
}
