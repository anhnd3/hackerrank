package breakingtherecords

// BreakingRecords ...
func BreakingRecords(scores []int32) []int32 {
	var counterBest, counterWorst int32
	best := scores[0]
	worst := scores[0]

	for _, v := range scores {
		if v > best {
			best = v
			counterBest++
		}
		if v < worst {
			worst = v
			counterWorst++
		}
	}

	return []int32{counterBest, counterWorst}
}
