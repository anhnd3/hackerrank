package repeatedstring

// RepeatedString ...
func RepeatedString(s string, n int64) int64 {
	var count int64
	var length = int64(len(s))

	if length == 1 && s == "a" {
		return n
	}

	var i int64
	for i = 0; i < length; i++ {
		if s[i] == 'a' {
			count++
		}
	}

	count = count * int64(n/length)

	remainStr := s[:int64(n%length)]

	for i = 0; i < int64(len(remainStr)); i++ {
		if remainStr[i] == 'a' {
			count++
		}
	}

	return count
}
