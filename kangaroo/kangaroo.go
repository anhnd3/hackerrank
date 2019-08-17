package kangaroo

// Kangaroo ...
func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if x1 > x2 && v1 > v2 {
		return "NO"
	}

	if x2 > x1 && v2 > v1 {
		return "NO"
	}

	if x1 != x2 && v1 == v2 {
		return "NO"
	}

	if x1 == x2 && v1 != v2 {
		return "NO"
	}

	isSameLocation := false

	for !isSameLocation {
		tmpx1 := x1 + v1
		tmpx2 := x2 + v2
		if tmpx1 == tmpx2 {
			return "YES"
		}

		if x1 > x2 && tmpx2 > x1 {
			return "NO"
		}

		if x1 < x2 && tmpx1 > tmpx2 {
			return "NO"
		}

		x1 = tmpx1
		x2 = tmpx2
	}
	return "NO"
}
