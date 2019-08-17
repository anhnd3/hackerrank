package electronicsshop

// GetMoneySpent ...
func GetMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var maxOfBudget int32
	maxOfBudget = -1
	for _, keyboard := range keyboards {
		for _, drive := range drives {
			tmpSum := keyboard + drive
			if tmpSum <= b && tmpSum > maxOfBudget {
				maxOfBudget = tmpSum
			}
		}
	}
	return maxOfBudget
}
