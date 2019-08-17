package drawingbook

// PageCount ...
func PageCount(n int32, p int32) int32 {
	var fromStart, fromEnd, i int32

	for i = 0; i < n; i = i + 2 {
		if p == i || p == i+1 {
			break
		}
		fromStart++
	}

	if n%2 == 0 {
		n = n + 1
	}
	for i = n; i > 0; i = i - 2 {
		if p == i || p == i-1 {
			break
		}
		fromEnd++
	}

	if fromStart > fromEnd {
		return fromEnd
	}
	return fromStart
}
