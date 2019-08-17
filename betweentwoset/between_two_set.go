package betweentwoset

// GetTotalX ...
func GetTotalX(a []int32, b []int32) int32 {
	// Write your code here
	lcm, _ := GetLCMArr(a)
	var result int32
	min := b[0]

	for _, v := range b {
		if v < min {
			min = v
		}
	}

	tmpLCM := lcm
	for tmpLCM <= min {
		counter := 0
		for _, v := range b {
			if v%tmpLCM == 0 {
				counter++
			}
		}
		if counter == len(b) {
			result++
		}
		tmpLCM += lcm
	}

	return result
}

// GetGCD of 2 number
func GetGCD(a, b int32) int32 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// GetLCM of 2 number
func GetLCM(a, b int32) int32 {
	return (a * b) / GetGCD(a, b)
}

// GetLCMArr of array
func GetLCMArr(a []int32) (int32, error) {
	var lcm int32
	lcm = 1
	for _, v := range a {
		lcm = GetLCM(lcm, v)
	}
	return lcm, nil
}
