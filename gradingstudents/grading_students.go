package gradingstudents

// GradingStudents ...
func GradingStudents(grades []int32) []int32 {
	// Write your code here
	var finalGrade []int32
	for _, value := range grades {
		rounded := value
		if value >= 38 {
			lastDigit := value % 10
			remainDigit := value / 10

			if lastDigit < 5 {
				lastDigit = 5
			} else {
				lastDigit = 0
				remainDigit++
			}

			rounded = remainDigit*10 + lastDigit
			if rounded-value >= 3 {
				rounded = value
			}
		}
		finalGrade = append(finalGrade, rounded)
	}
	return finalGrade
}
