package migratorybirds

// MigratoryBirds ...
func MigratoryBirds(arr []int32) int32 {
	mapBird := make(map[int32]int32)
	for _, v := range arr {
		if val, ok := mapBird[v]; ok {
			mapBird[v] = val + 1
		} else {
			mapBird[v] = 1
		}
	}

	var highestVal, lowestKey int32
	lowestKey = -1
	for key, val := range mapBird {
		if val > 1 {
			if val > highestVal {
				highestVal = val
				lowestKey = key
			} else if val == highestVal && key < lowestKey {
				lowestKey = key
			}
		}
	}

	return lowestKey
}
