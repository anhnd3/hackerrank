package climbingtheleaderboard

// ClimbingLeaderboard ...
func ClimbingLeaderboard(scores []int32, alices []int32) []int32 {
	// distinct score
	var distinctScore []int32
	mapScore := make(map[int32]bool)
	for _, score := range scores {
		if _, ok := mapScore[score]; !ok {
			mapScore[score] = true
			distinctScore = append(distinctScore, score)
		}
	}

	var i, lenAlices, lastedIndex, lenDistinctRank int32
	lenAlices = int32(len(alices))
	lenDistinctRank = int32(len(distinctScore))

	// init results array
	lastedIndex = lenDistinctRank - 1
	results := make([]int32, lenAlices)
	for i = 0; i < lenAlices; i++ {
		alice := alices[i]
		var currRank, index int32
		for index = lastedIndex; index >= 0; index-- {
			if currRank == 0 {
				if alice < distinctScore[index] {
					currRank = index + 2
					lastedIndex = index
					break
				} else if alice == distinctScore[index] {
					currRank = index + 1
					lastedIndex = index
					break
				}
			}
		}
		if currRank == 0 {
			currRank = 1
		}
		results[i] = currRank
	}

	return results
}
