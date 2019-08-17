package jumpingontheclouds

// JumpingOnClouds ...
func JumpingOnClouds(c []int32) int32 {
	var index, jump int32
	length := int32(len(c))
	for index < length-1 {
		if (index+2) < length && c[index+2] == 0 {
			index += 2
		} else {
			index++
		}
		jump++
	}
	return jump
}
